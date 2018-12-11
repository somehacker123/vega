package datastore

import (
	"errors"
	"fmt"
	"sync"
	"vega/log"
	"vega/filters"
	"vega/msg"
	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
)

type TradeStore interface {
	Subscribe(trades chan<- []msg.Trade) uint64
	Unsubscribe(id uint64) error

	// Post adds a trade to the store, adds
	// to queue the operation to be committed later.
	Post(trade *msg.Trade) error

	// Commit typically saves any operations that are queued to underlying storage,
	// if supported by underlying storage implementation.
	Commit() error

	// Close can be called to clean up and close any storage
	// connections held by the underlying storage mechanism.
	Close()

	// GetByMarket retrieves trades for a given market.
	GetByMarket(market string, params *filters.TradeQueryFilters) ([]*msg.Trade, error)
	// GetByMarketAndId retrieves a trade for a given market and id.
	GetByMarketAndId(market string, id string) (*msg.Trade, error)
	// GetBuMarketAndOrderId retrieves trades relating to the given market and order id.
	GetByMarketAndOrderId(market string, orderId string) ([]*msg.Trade, error)
	// GetByParty retrieves trades for a given party.
	GetByParty(party string, params *filters.TradeQueryFilters) ([]*msg.Trade, error)
	// GetByPartyAndId retrieves a trade for a given party and id.
	GetByPartyAndId(party string, id string) (*msg.Trade, error)
	// GetMarkPrice returns the current market price.
	GetMarkPrice(market string) (uint64, error)
	
	// GetTradesBySideBuckets retrieves a map of market name to market buckets.
	GetTradesBySideBuckets(party string) map[string]*MarketBucket
}

// badgerTradeStore is a package internal data struct that implements the TradeStore interface.
type badgerTradeStore struct {
	badger *badgerStore
	buffer []msg.Trade

	subscribers map[uint64] chan<- []msg.Trade
	subscriberId uint64
	mu sync.Mutex
}

// NewTradeStore is used to initialise and create a TradeStore, this implementation is currently
// using the badger k-v persistent storage engine under the hood. The caller will specify a dir to
// use as the storage location on disk for any stored files.
func NewTradeStore(dir string) TradeStore {
	db, err := badger.Open(customBadgerOptions(dir))
	if err != nil {
		log.Fatalf(err.Error())
	}
	bs := badgerStore{db: db}
	return &badgerTradeStore{badger: &bs,  buffer: make([]msg.Trade, 0), subscribers: make(map[uint64] chan<- []msg.Trade)}
}

// Subscribe to a channel of new or updated trades. The subscriber id will be returned as a uint64 value
// and must be retained for future reference and to unsubscribe.
func (ts *badgerTradeStore) Subscribe(trades chan<- []msg.Trade) uint64 {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.subscriberId = ts.subscriberId+1
	ts.subscribers[ts.subscriberId] = trades

	log.Debugf("TradeStore -> Subscribe: Trade subscriber added: %d", ts.subscriberId)
	return ts.subscriberId
}

// Unsubscribe from an trades channel. Provide the subscriber id you wish to stop receiving new events for.
func (ts *badgerTradeStore) Unsubscribe(id uint64) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if len(ts.subscribers) == 0 {
		log.Debugf("TradeStore -> Unsubscribe: No subscribers connected")
		return nil
	}

	if _, exists := ts.subscribers[id]; exists {
		delete(ts.subscribers, id)
		log.Debugf("TradeStore -> Unsubscribe: Subscriber removed: %v", id)
		return nil
	}
	return errors.New(fmt.Sprintf("TradeStore subscriber does not exist with id: %d", id))
}

// Post adds an trade to the badger store, adds
// to queue the operation to be committed later.
func (ts *badgerTradeStore) Post(trade *msg.Trade) error {
	// with badger we always buffer for future batch insert via Commit()
	ts.addToBuffer(*trade)
	return nil
}

// Commit saves any operations that are queued to badger store, and includes all updates.
// It will also call notify() to push updated data to any subscribers.
func (ts *badgerTradeStore) Commit() error {
	if len(ts.buffer) == 0 {
		return nil
	}

	ts.mu.Lock()
	items := ts.buffer
	ts.buffer = make([]msg.Trade, 0)
	ts.mu.Unlock()

	err := ts.writeBatch(items)
	if err != nil {
		return err
	}
	err = ts.notify(items)
	if err != nil {
		return err
	}
	return nil
}

// GetByMarket retrieves trades for a given market. Provide optional query filters to
// refine the data set further (if required), any errors will be returned immediately.
func (ts *badgerTradeStore) GetByMarket(market string, queryFilters *filters.TradeQueryFilters) ([]*msg.Trade, error) {
	var result []*msg.Trade

	if queryFilters == nil {
		queryFilters = &filters.TradeQueryFilters{}
	}

	txn := ts.badger.db.NewTransaction(false)
	defer txn.Discard()

	filter := TradeFilter{queryFilter: queryFilters}
	descending := filter.queryFilter.HasLast()
	it := ts.badger.getIterator(txn, descending)
	defer it.Close()

	marketPrefix, validForPrefix := ts.badger.marketPrefix(market, descending)
	for it.Seek(marketPrefix); it.ValidForPrefix(validForPrefix); it.Next() {
		item := it.Item()
		tradeBuf, _ := item.ValueCopy(nil)
		var trade msg.Trade
		if err := proto.Unmarshal(tradeBuf, &trade); err != nil {
			log.Errorf("unmarshal failed: %s", err.Error())
			return nil, err
		}
		if filter.apply(&trade) {
			result = append(result, &trade)
		}
		if filter.isFull() {
			break
		}
	}
	return result, nil
}

// GetByMarketAndId retrieves a trade for a given market and id, any errors will be returned immediately.
func (ts *badgerTradeStore) GetByMarketAndId(market string, Id string) (*msg.Trade, error) {
	var trade msg.Trade

	txn := ts.badger.db.NewTransaction(false)
	defer txn.Discard()

	marketKey := ts.badger.tradeMarketKey(market, Id)
	item, err := txn.Get(marketKey)
	if err != nil {
		return nil, err
	}
	tradeBuf, _ := item.ValueCopy(nil)
	if err := proto.Unmarshal(tradeBuf, &trade); err != nil {
		log.Errorf("Failed to unmarshal %s", err.Error())
		return nil, err
	}
	return &trade, err
}

// GetByParty retrieves trades for a given party. Provide optional query filters to
// refine the data set further (if required), any errors will be returned immediately.
func (ts *badgerTradeStore) GetByParty(party string, queryFilters *filters.TradeQueryFilters) ([]*msg.Trade, error) {
	var result []*msg.Trade
	
	if queryFilters == nil {
		queryFilters = &filters.TradeQueryFilters{}
	}

	txn := ts.badger.db.NewTransaction(false)
	defer txn.Discard()
	
	filter := TradeFilter{queryFilter: queryFilters}
	descending := filter.queryFilter.HasLast()
	it := ts.badger.getIterator(txn, descending)
	defer it.Close()

	partyPrefix, validForPrefix := ts.badger.partyPrefix(party, descending)
	for it.Seek(partyPrefix); it.ValidForPrefix(validForPrefix); it.Next() {
		marketKeyItem := it.Item()
		marketKey, _ := marketKeyItem.ValueCopy(nil)
		tradeItem, err := txn.Get(marketKey)
		if err != nil {
			log.Errorf("trade with key %s does not exist in store", string(marketKey))
			return nil, err
		}
		tradeBuf, _ := tradeItem.ValueCopy(nil)
		var trade msg.Trade
		if err := proto.Unmarshal(tradeBuf, &trade); err != nil {
			log.Errorf("unmarshal failed %s", err.Error())
			return nil, err
		}
		if filter.apply(&trade) {
			result = append(result, &trade)
		}
		if filter.isFull() {
			break
		}
	}

	return result, nil
}

// GetByPartyAndId retrieves a trade for a given party and id.
func (ts *badgerTradeStore) GetByPartyAndId(party string, Id string) (*msg.Trade, error) {
	var trade msg.Trade
	err := ts.badger.db.View(func(txn *badger.Txn) error {
		partyKey := ts.badger.tradePartyKey(party, Id)
		marketKeyItem, err := txn.Get(partyKey)
		if err != nil {
			return err
		}
		marketKey, err := marketKeyItem.ValueCopy(nil)
		if err != nil {
			return err
		}
		tradeItem, err := txn.Get(marketKey)
		if err != nil {
			return err
		}

		tradeBuf, err := tradeItem.ValueCopy(nil)
		if err != nil {
			return err
		}
		if err := proto.Unmarshal(tradeBuf, &trade); err != nil {
			log.Errorf("unmarshal failed %s", err.Error())
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &trade, nil
}

// GetBuMarketAndOrderId retrieves trades relating to the given market and order id.
func (ts *badgerTradeStore) GetByMarketAndOrderId(market string, orderId string) ([]*msg.Trade, error) {
	var result []*msg.Trade

	queryFilters := &filters.TradeQueryFilters{}
	txn := ts.badger.db.NewTransaction(false)
	defer txn.Discard()

	filter := TradeFilter{queryFilter: queryFilters}
	descending := filter.queryFilter.HasLast()
	it := ts.badger.getIterator(txn, descending)
	defer it.Close()
	
	marketPrefix, validForPrefix := ts.badger.marketPrefix(market, descending)
	for it.Seek(marketPrefix); it.ValidForPrefix(validForPrefix); it.Next() {
		item := it.Item()
		tradeBuf, _ := item.ValueCopy(nil)
		var trade msg.Trade
		if err := proto.Unmarshal(tradeBuf, &trade); err != nil {
			log.Errorf("unmarshal failed: %s", err.Error())
		}
		// We don't support OR in filter query so we need to do the below logic to support related trades given orderId
		if filter.apply(&trade) && (trade.BuyOrder == orderId || trade.SellOrder == orderId) {
			result = append(result, &trade)
		}
	}

	return result, nil
}

// Close our connection to the badger database
// ensuring errors will be returned up the stack.
func (ts *badgerTradeStore) Close() {
	ts.badger.db.Close()
}

// GetMarkPrice returns the current market price, for a requested market.
func (ts *badgerTradeStore) GetMarkPrice(market string) (uint64, error) {
	f := &filters.TradeQueryFilters{}
	recentTrade, err := ts.GetByMarket(market, f)
	if err != nil {
		return 0, err
	}
	if len(recentTrade) == 0 {
		return 0, errors.New("no trades available when getting market price")
	}
	return recentTrade[0].Price, nil
}

// add a trade to the write-batch/notify buffer.
func (ts *badgerTradeStore) addToBuffer(t msg.Trade) {
	ts.mu.Lock()
	ts.buffer = append(ts.buffer, t)
	ts.mu.Unlock()
}

// notify any subscribers of trade updates.
func (ts *badgerTradeStore) notify(items []msg.Trade) error {
	if len(items) == 0 {
		return nil
	}
	if len(ts.subscribers) == 0 {
		log.Debugf("TradeStore -> Notify: No subscribers connected")
		return nil
	}

	var ok bool
	for id, sub := range ts.subscribers {
		select {
		case sub <- items:
			ok = true
			break
		default:
			ok = false
		}
		if ok{
			log.Debugf("TradeStore -> send on channel success for subscriber %d", id)
		} else {
			log.Infof("TradeStore -> channel could not been updated for subscriber %d", id)
		}
	}
	return nil
}

// writeBatch flushes a batch of trades to the underlying badger store.
func (ts *badgerTradeStore) writeBatch(batch []msg.Trade) error {
	wb := ts.badger.db.NewWriteBatch()
	defer wb.Cancel()

	insertBatchAtomically := func() error {
		for idx := range batch {
			tradeBuf, err := proto.Marshal(&batch[idx])
			if err != nil {
				log.Errorf("marshal failed %s", err.Error())
			}
			marketKey := ts.badger.tradeMarketKey(batch[idx].Market, batch[idx].Id)
			idKey := ts.badger.tradeIdKey(batch[idx].Id)
			buyerPartyKey := ts.badger.tradePartyKey(batch[idx].Buyer, batch[idx].Id)
			sellerPartyKey := ts.badger.tradePartyKey(batch[idx].Seller, batch[idx].Id)
			if err := wb.Set(marketKey, tradeBuf, 0); err != nil {
				return err
			}
			if err := wb.Set(idKey, marketKey, 0); err != nil {
				return err
			}
			if err := wb.Set(buyerPartyKey, marketKey, 0); err != nil {
				return err
			}
			if err := wb.Set(sellerPartyKey, marketKey, 0); err != nil {
				return err
			}
		}
		return nil
	}

	if err := insertBatchAtomically(); err == nil {
		if err := wb.Flush(); err != nil {
			// todo: can we handle flush errors in a similar way to below?
			log.Errorf("failed to flush batch: %s", err)
		}
	} else {
		wb.Cancel()
		// todo: retry mechanism, also handle badger txn too large errors
		log.Errorf("failed to insert trade batch atomically, %s", err)
	}

	return nil
}


// TradeFilter is the trade specific filter query data holder. It includes the raw filters
// and helper methods that are used internally to apply and track filter state.
type TradeFilter struct {
	queryFilter *filters.TradeQueryFilters
	skipped uint64
	found uint64
}

func (f *TradeFilter) apply(trade *msg.Trade) (include bool) {
	if f.queryFilter.First == nil && f.queryFilter.Last == nil && f.queryFilter.Skip == nil {
		include = true
	} else {

		if f.queryFilter.HasFirst() && f.found < *f.queryFilter.First {
			include = true
		}
		if f.queryFilter.HasLast() && f.found < *f.queryFilter.Last {
			include = true
		}
		if f.queryFilter.HasSkip() && f.skipped < *f.queryFilter.Skip {
			f.skipped++
			return false
		}
	}

	if !applyTradeFilters(trade, f.queryFilter) {
		return false
	}

	// if item passes the filter, increment the found counter
	if include {
		f.found++
	}
	return include
}

func (f *TradeFilter) isFull() bool {
	if f.queryFilter.HasLast() && f.found == *f.queryFilter.Last {
		return true
	}
	return false
}