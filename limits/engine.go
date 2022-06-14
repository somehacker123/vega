// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package limits

import (
	"context"
	"fmt"
	"time"

	"code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
)

type Engine struct {
	log    *logging.Logger
	cfg    Config
	broker Broker

	timeService TimeService
	blockCount  uint16

	// are these action possible?
	canProposeMarket, canProposeAsset, bootstrapFinished bool

	// Settings from the genesis state
	proposeMarketEnabled, proposeAssetEnabled         bool
	proposeMarketEnabledFrom, proposeAssetEnabledFrom time.Time
	bootstrapBlockCount                               uint16

	genesisLoaded bool

	// snapshot state
	lss *limitsSnapshotState
}

type Broker interface {
	Send(event events.Event)
}

// TimeService provide the time of the vega node using the tm time.
//go:generate go run github.com/golang/mock/mockgen -destination mocks/time_service_mock.go -package mocks code.vegaprotocol.io/vega/limits TimeService
type TimeService interface {
	GetTimeNow() time.Time
}

func New(log *logging.Logger, cfg Config, tm TimeService, broker Broker) *Engine {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())

	return &Engine{
		log:         log,
		cfg:         cfg,
		lss:         &limitsSnapshotState{changed: true},
		broker:      broker,
		timeService: tm,
	}
}

// UponGenesis load the limits from the genesis state.
func (e *Engine) UponGenesis(ctx context.Context, rawState []byte) (err error) {
	e.log.Debug("Entering limits.Engine.UponGenesis")
	defer func() {
		if err != nil {
			e.log.Debug("Failure in limits.Engine.UponGenesis", logging.Error(err))
		} else {
			e.log.Debug("Leaving limits.Engine.UponGenesis without error")
		}
		e.genesisLoaded = true
		e.lss.changed = true
	}()

	state, err := LoadGenesisState(rawState)
	if err != nil && err != ErrNoLimitsGenesisState {
		e.log.Error("unable to load genesis state",
			logging.Error(err))
		return err
	}

	if err == ErrNoLimitsGenesisState {
		defaultState := DefaultGenesisState()
		state = &defaultState
	}

	// set enabled by default if not genesis state
	if state == nil {
		e.proposeAssetEnabled = true
		e.proposeMarketEnabled = true
		return nil
	}

	e.proposeAssetEnabled = state.ProposeAssetEnabled
	e.proposeMarketEnabled = state.ProposeMarketEnabled
	e.bootstrapBlockCount = state.BootstrapBlockCount

	e.log.Info("loaded limits genesis state",
		logging.String("state", fmt.Sprintf("%#v", *state)))

	e.sendEvent(ctx)
	return nil
}

func (e *Engine) OnLimitsProposeMarketEnabledFromUpdate(ctx context.Context, date string) error {
	// we already can propose market, nothing else to be done
	if e.canProposeMarket {
		return nil
	}
	// already validated by the netparams
	// no need to check it again, this is a valid date
	if len(date) <= 0 {
		e.proposeMarketEnabledFrom = time.Time{}
	} else {
		t, _ := time.Parse(time.RFC3339, date)
		if e.timeService.GetTimeNow().Before(t) {
			// only if the date is in the future
			e.proposeMarketEnabledFrom = t
		}
	}
	e.sendEvent(ctx)

	return nil
}

func (e *Engine) OnLimitsProposeAssetEnabledFromUpdate(ctx context.Context, date string) error {
	// we already can propose assets, nothing can be changes anymore
	if e.canProposeAsset {
		return nil
	}

	// already validated by the netparams
	// no need to check it again, this is a valid date
	if len(date) <= 0 {
		e.proposeAssetEnabledFrom = time.Time{}
	} else {
		t, _ := time.Parse(time.RFC3339, date)
		if e.timeService.GetTimeNow().Before(t) {
			// only if the date is in the future
			e.proposeAssetEnabledFrom = t
		}
	}
	e.sendEvent(ctx)

	return nil
}

func (e *Engine) OnTick(ctx context.Context, t time.Time) {
	if !e.genesisLoaded || (e.bootstrapFinished && e.canProposeAsset && e.canProposeMarket) {
		return
	}

	if !e.bootstrapFinished {
		e.blockCount++
		if e.blockCount > e.bootstrapBlockCount {
			e.log.Info("bootstraping period finished, transactions are now allowed")
			e.bootstrapFinished = true
			e.sendEvent(ctx)
		}
		e.lss.changed = true
	}

	if !e.canProposeMarket && e.bootstrapFinished && e.proposeMarketEnabled && t.After(e.proposeMarketEnabledFrom) {
		e.log.Info("all required conditions are met, proposing markets is now allowed")
		e.canProposeMarket = true
		e.lss.changed = true
		e.sendEvent(ctx)
	}
	if !e.canProposeAsset && e.bootstrapFinished && e.proposeAssetEnabled && t.After(e.proposeAssetEnabledFrom) {
		e.log.Info("all required conditions are met, proposing assets is now allowed")
		e.canProposeAsset = true
		e.lss.changed = true
		e.sendEvent(ctx)
	}
}

func (e *Engine) CanProposeMarket() bool {
	return e.canProposeMarket
}

func (e *Engine) CanProposeAsset() bool {
	return e.canProposeAsset
}

func (e *Engine) CanTrade() bool {
	return e.canProposeAsset && e.canProposeMarket
}

func (e *Engine) BootstrapFinished() bool {
	return e.bootstrapFinished
}

func (e *Engine) sendEvent(ctx context.Context) {
	limits := vega.NetworkLimits{
		CanProposeMarket:     e.canProposeMarket,
		CanProposeAsset:      e.canProposeAsset,
		BootstrapFinished:    e.bootstrapFinished,
		ProposeMarketEnabled: e.proposeMarketEnabled,
		ProposeAssetEnabled:  e.proposeAssetEnabled,
		BootstrapBlockCount:  uint32(e.bootstrapBlockCount),
		GenesisLoaded:        e.genesisLoaded,
	}

	if !e.proposeMarketEnabledFrom.IsZero() {
		limits.ProposeMarketEnabledFrom = e.proposeAssetEnabledFrom.UnixNano()
	}

	if !e.proposeAssetEnabledFrom.IsZero() {
		limits.ProposeAssetEnabledFrom = e.proposeAssetEnabledFrom.UnixNano()
	}

	event := events.NewNetworkLimitsEvent(ctx, &limits)
	e.broker.Send(event)
}
