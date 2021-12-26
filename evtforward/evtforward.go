package evtforward

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/metrics"
	"code.vegaprotocol.io/vega/txn"

	"github.com/golang/protobuf/proto"
)

var (
	// ErrEvtAlreadyExist we have already handled this event.
	ErrEvtAlreadyExist = errors.New("event already exist")
	// ErrPubKeyNotAllowlisted this pubkey is not part of the allowlist.
	ErrPubKeyNotAllowlisted = errors.New("pubkey not allowlisted")
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/time_service_mock.go -package mocks code.vegaprotocol.io/vega/evtforward TimeService
type TimeService interface {
	GetTimeNow() time.Time
	NotifyOnTick(f func(context.Context, time.Time))
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/commander_mock.go -package mocks code.vegaprotocol.io/vega/evtforward Commander
type Commander interface {
	Command(ctx context.Context, cmd txn.Command, payload proto.Message, f func(error))
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/validator_topology_mock.go -package mocks code.vegaprotocol.io/vega/evtforward ValidatorTopology
type ValidatorTopology interface {
	SelfNodeID() string
	AllNodeIDs() []string
}

// EvtForwarder receive events from the blockchain queue
// and will try to send them to the vega chain.
// this will select a node in the network to forward the event.
type EvtForwarder struct {
	log  *logging.Logger
	cfg  Config
	cmd  Commander
	self string

	evtsmu    sync.Mutex
	ackedEvts map[string]*commandspb.ChainEvent
	evts      map[string]tsEvt

	mu               sync.RWMutex
	bcQueueAllowlist atomic.Value // this is actually an map[string]struct{}
	currentTime      time.Time
	nodes            []string

	top  ValidatorTopology
	efss *efSnapshotState
}

type tsEvt struct {
	ts  time.Time // timestamp of the block when the event has been added
	evt *commandspb.ChainEvent
}

// New creates a new instance of the event forwarder.
func New(log *logging.Logger, cfg Config, cmd Commander, time TimeService, top ValidatorTopology) *EvtForwarder {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())
	var allowlist atomic.Value
	allowlist.Store(buildAllowlist(cfg))
	evtf := &EvtForwarder{
		cfg:              cfg,
		log:              log,
		cmd:              cmd,
		nodes:            []string{},
		self:             top.SelfNodeID(),
		currentTime:      time.GetTimeNow(),
		ackedEvts:        map[string]*commandspb.ChainEvent{},
		evts:             map[string]tsEvt{},
		top:              top,
		bcQueueAllowlist: allowlist,
		efss: &efSnapshotState{
			changed:    true,
			hash:       []byte{},
			serialised: []byte{},
		},
	}
	evtf.updateValidatorsList()
	time.NotifyOnTick(evtf.onTick)
	return evtf
}

func buildAllowlist(cfg Config) map[string]struct{} {
	allowlist := make(map[string]struct{}, len(cfg.BlockchainQueueAllowlist))
	for _, v := range cfg.BlockchainQueueAllowlist {
		allowlist[v] = struct{}{}
	}
	return allowlist
}

// ReloadConf updates the internal configuration of the collateral engine.
func (e *EvtForwarder) ReloadConf(cfg Config) {
	e.log.Info("reloading configuration")
	if e.log.GetLevel() != cfg.Level.Get() {
		e.log.Info("updating log level",
			logging.String("old", e.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		e.log.SetLevel(cfg.Level.Get())
	}

	e.cfg = cfg
	// update the allowlist
	e.log.Info("evtforward allowlist updated",
		logging.Reflect("list", cfg.BlockchainQueueAllowlist))
	e.bcQueueAllowlist.Store(buildAllowlist(cfg))
}

// Ack will return true if the event is newly acknowledge
// if the event already exist and was already acknowledge this will return false.
func (e *EvtForwarder) Ack(evt *commandspb.ChainEvent) bool {
	res := "ok"
	defer func() {
		metrics.EvtForwardInc("ack", res)
	}()

	e.evtsmu.Lock()
	defer e.evtsmu.Unlock()

	key, err := e.getEvtKey(evt)
	if err != nil {
		e.log.Error("could not get event key", logging.Error(err))
		return false
	}
	_, ok, acked := e.getEvt(key)
	if ok && acked {
		e.log.Error("event already acknowledged",
			logging.String("evt", evt.String()),
		)
		res = "alreadyacked"
		// this was already acknowledged, nothing to be done, return false
		return false
	}
	if ok {
		// exists but was not acknowledged
		// we just remove it from the non-acked table
		delete(e.evts, key)
	}

	// now add it to the acknowledged evts
	e.ackedEvts[key] = evt
	e.efss.changed = true
	e.log.Info("new event acknowledged", logging.String("event", evt.String()))
	return true
}

func (e *EvtForwarder) isAllowlisted(pubkey string) bool {
	allowlist := e.bcQueueAllowlist.Load().(map[string]struct{})
	_, ok := allowlist[pubkey]
	return ok
}

// Forward will forward an ChainEvent to the tendermint network
// we expect the pubkey to be an ed25519 pubkey hex encoded.
func (e *EvtForwarder) Forward(ctx context.Context, evt *commandspb.ChainEvent, pubkey string) error {
	res := "ok"
	defer func() {
		metrics.EvtForwardInc("forward", res)
	}()

	if e.log.GetLevel() <= logging.DebugLevel {
		e.log.Debug("new event received to be forwarded",
			logging.String("event", evt.String()),
		)
	}

	// check if the sender of the event is whitelisted
	if !e.isAllowlisted(pubkey) {
		res = "pubkeynotallowed"
		return ErrPubKeyNotAllowlisted
	}

	e.evtsmu.Lock()
	defer e.evtsmu.Unlock()

	key, err := e.getEvtKey(evt)
	if err != nil {
		return err
	}
	_, ok, ack := e.getEvt(key)
	if ok {
		e.log.Error("event already processed",
			logging.String("evt", evt.String()),
			logging.Bool("acknowledged", ack),
		)
		res = "dupevt"
		return ErrEvtAlreadyExist
	}

	e.evts[key] = tsEvt{ts: e.currentTime, evt: evt}
	if e.isSender(evt) {
		// we are selected to send the event, let's do it.
		e.send(ctx, evt)
	}
	return nil
}

// ForwardFromSelf will forward event seen by the node itself, not from
// an external service like the eef for example
func (e *EvtForwarder) ForwardFromSelf(evt *commandspb.ChainEvent) {
	if e.log.GetLevel() <= logging.DebugLevel {
		e.log.Debug("new event received to be forwarded",
			logging.String("event", evt.String()),
		)
	}

	e.evtsmu.Lock()
	defer e.evtsmu.Unlock()

	key, err := e.getEvtKey(evt)
	if err != nil {
		// no way this error would be badly formatted
		// it is sent by the node, a badly formatted event
		// would mean a code bug
		e.log.Panic("invalid event to be forwarded",
			logging.String("event", evt.String()),
			logging.Error(err),
		)
	}

	_, ok, ack := e.getEvt(key)
	if ok {
		e.log.Error("event already processed",
			logging.String("evt", evt.String()),
			logging.Bool("acknowledged", ack),
		)
	}

	e.evts[key] = tsEvt{ts: e.currentTime, evt: evt}
}

func (e *EvtForwarder) updateValidatorsList() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.self = e.top.SelfNodeID()
	e.nodes = e.top.AllNodeIDs()
	sort.SliceStable(e.nodes, func(i, j int) bool {
		return e.nodes[i] < e.nodes[j]
	})
}

// getEvt assumes the lock is acquired before being called.
func (e *EvtForwarder) getEvt(key string) (evt *commandspb.ChainEvent, ok bool, acked bool) {
	if evt, ok = e.ackedEvts[key]; ok {
		return evt, true, true
	}

	if tsEvt, ok := e.evts[key]; ok {
		return tsEvt.evt, true, false
	}

	return nil, false, false
}

func (e *EvtForwarder) send(ctx context.Context, evt *commandspb.ChainEvent) {
	if e.log.GetLevel() <= logging.DebugLevel {
		e.log.Debug("trying to send event",
			logging.String("event", evt.String()),
		)
	}

	// error doesn't matter here
	e.cmd.Command(ctx, txn.ChainEventCommand, evt, nil)
}

func (e *EvtForwarder) isSender(evt *commandspb.ChainEvent) bool {
	key, err := e.makeEvtHashKey(evt)
	if err != nil {
		e.log.Error("could not marshal event", logging.Error(err))
		return false
	}
	h := e.hash(key) + uint64(e.currentTime.Unix())

	e.mu.RLock()
	if len(e.nodes) <= 0 {
		e.mu.RUnlock()
		return false
	}
	node := e.nodes[h%uint64(len(e.nodes))]
	e.mu.RUnlock()

	return node == e.self
}

func (e *EvtForwarder) onTick(ctx context.Context, t time.Time) {
	e.currentTime = t

	// get an updated list of validators from the topology
	e.updateValidatorsList()

	e.mu.RLock()
	retryRate := e.cfg.RetryRate.Duration
	e.mu.RUnlock()

	e.evtsmu.Lock()
	defer e.evtsmu.Unlock()

	// try to send all event that are not acknowledged at the moment
	for k, evt := range e.evts {
		// do we need to try to forward the event again?
		if evt.ts.Add(retryRate).Before(t) {
			// set next retry
			e.evts[k] = tsEvt{ts: t, evt: evt.evt}
			if e.isSender(evt.evt) {
				// we are selected to send the event, let's do it.
				e.send(ctx, evt.evt)
			}
		}
	}
}

func (e *EvtForwarder) getEvtKey(evt *commandspb.ChainEvent) (string, error) {
	mevt, err := e.marshalEvt(evt)
	if err != nil {
		return "", fmt.Errorf("invalid event: %w", err)
	}

	return string(crypto.Hash(mevt)), nil
}

func (e *EvtForwarder) marshalEvt(evt *commandspb.ChainEvent) ([]byte, error) {
	pbuf := proto.Buffer{}
	pbuf.Reset()
	pbuf.SetDeterministic(true)
	if err := pbuf.Marshal(evt); err != nil {
		return nil, err
	}
	return pbuf.Bytes(), nil
}

func (e *EvtForwarder) makeEvtHashKey(evt *commandspb.ChainEvent) ([]byte, error) {
	// deterministic marshal of the event
	pbuf, err := e.marshalEvt(evt)
	if err != nil {
		return nil, err
	}
	return pbuf, nil
}

func (e *EvtForwarder) hash(key []byte) uint64 {
	h := fnv.New64a()
	h.Write(key)

	return h.Sum64()
}
