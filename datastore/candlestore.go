package datastore

import (
	"errors"
	"fmt"
	"time"

	"vega/msg"

	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"vega/log"
	"sync"
)

type candleStore struct {
	persistentStore *badger.DB

	subscribers map[uint64] map[string]chan msg.Candle
	buffer map[string]msg.Candle
	subscriberId uint64
	mu sync.Mutex
}

func NewCandleStore(dir string) CandleStore {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return &candleStore{persistentStore: db, buffer: make(map[string]msg.Candle)}
}

func (c *candleStore) Subscribe(internalTransport map[string]chan msg.Candle) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subscribers == nil {
		log.Debugf("CandleStore -> Subscribe: Creating subscriber chan map")
		c.subscribers = make(map[uint64] map[string]chan msg.Candle)
	}

	c.subscriberId = c.subscriberId+1
	c.subscribers[c.subscriberId] = internalTransport
	log.Debugf("CandleStore -> Subscribe: Candle subscriber added: %d", c.subscriberId)
	return c.subscriberId
}

func (c *candleStore) Unsubscribe(id uint64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subscribers == nil || len(c.subscribers) == 0 {
		log.Debugf("CandleStore -> Unsubscribe: No subscribers connected")
		return nil
	}

	if _, exists := c.subscribers[id]; exists {
		delete(c.subscribers, id)
		log.Debugf("CandleStore -> Unsubscribe: Subscriber removed: %v", id)
		return nil
	}
	return errors.New(fmt.Sprintf("CandleStore subscriber does not exist with id: %d", id))
}

func (c *candleStore) Notify() error {

	if c.subscribers == nil || len(c.subscribers) == 0 {
		log.Debugf("CandleStore -> Notify: No subscribers connected")
		return nil
	}

	if c.buffer == nil {
		// Only publish when we have items
		log.Debugf("CandleStore -> Notify: No new candle")
		return nil
	}

	c.mu.Lock()
	intervalsToCandlesMap := c.buffer
	c.mu.Unlock()

	// update candle for each interval for each subscriber
	for id, internalTransport := range c.subscribers {
		for interval, candleForUpdate := range intervalsToCandlesMap {
			select {
			case internalTransport[interval] <- candleForUpdate:
				log.Debugf("Candle updated for interval: ", interval)
				break
			default:
				log.Infof("Candles state could not been updated for subscriber %d", id)
			}
		}
	}
	return nil
}

func (c *candleStore) QueueEvent(candle msg.Candle, interval string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subscribers == nil || len(c.subscribers) == 0 {
		log.Debugf("CandleStore -> queueEvent: No subscribers connected")
		return nil
	}

	c.buffer[interval] = candle

	log.Debugf("CandleStore -> queueEvent: Adding candle to buffer of intervals at: %s", interval)
	return nil
}

func (c *candleStore) Close() {
	defer c.persistentStore.Close()
}

func (c *candleStore) GetCandles(market string, sinceTimestamp uint64, interval string) []*msg.Candle {
	fetchKey := generateFetchKey(market, sinceTimestamp, interval)
	it :=  c.persistentStore.NewTransaction(false).NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()
	prefix := []byte(fmt.Sprintf("M:%s_I:%s_T:", market, interval))

	fmt.Printf("prefix %s\n", fmt.Sprintf("M:%s_I:%s_T:", market, interval))
	fmt.Printf("fetchkey %s\n", string(fetchKey))

	var candles []*msg.Candle
	for it.Seek(fetchKey); it.ValidForPrefix(prefix); it.Next() {
		item := it.Item()
		value, err := item.ValueCopy(nil)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		var newCandle msg.Candle
		if err := proto.Unmarshal(value, &newCandle); err != nil {
			fmt.Printf(err.Error())
			continue
		}
		candles = append(candles, &newCandle)
	}

	return candles
}

func (c *candleStore) GenerateCandles(trade *msg.Trade) error {

	//given trade generate appropriate timestamps and badger keys for each interval
	candleTimestamps, badgerKeys := generateCandleParamsForTimestamp(trade.Market, trade.Timestamp)

	// for each trade generate candle keys and run update on each record
	txn := c.persistentStore.NewTransaction(true)
	for interval, badgerKey := range badgerKeys {

		item, err := txn.Get(badgerKey)

		// if key does not exist, insert candle for this timestamp
		if err == badger.ErrKeyNotFound {
			fmt.Printf("KEY DOES NOT EXIST, %s\n", badgerKey)
			candleTimestamp := candleTimestamps[interval]
			candle := NewCandle(uint64(candleTimestamp), trade.Price, trade.Size)
			candleBuf, err := proto.Marshal(candle)
			if err != nil {
				return err
			}

			if err = txn.Set(badgerKey, candleBuf); err != nil {
				return err
			}
			fmt.Printf("Candle inserted %+v\n", candle)
		}

		// if key exists, update candle with this trade
		if err == nil {
			// umarshal fetched candle
			var candleForUpdate msg.Candle
			itemCopy, err := item.ValueCopy(nil)
			proto.Unmarshal(itemCopy, &candleForUpdate)
			fmt.Printf("Candle fetched %+v\n", candleForUpdate)

			// update fetched candle with new trade
			UpdateCandle(&candleForUpdate, trade)

			// marshal candle
			candleBuf, err := proto.Marshal(&candleForUpdate)
			if err != nil {
				return err
			}

			// push candle to badger
			if err = txn.Set(badgerKey, candleBuf); err != nil {
				return err
			}
			fmt.Printf("Candle updated and inserted %+v\n", candleForUpdate)
		}
	}

	if err := txn.Commit(); err != nil {
		return err
	}
	fmt.Printf("All good for trade %+v\n", trade)
	return nil
}


func (c *candleStore) GenerateEmptyCandles(market string, timestamp uint64) error {

	// generate keys for this timestamp
	candleTimestamp, candleKeys := generateCandleParamsForTimestamp(market, timestamp)

	// if key does not exist seek most recent values, create empty candle with those close value and insert
	txn := c.persistentStore.NewTransaction(true)

	// for all candle intervals
	for interval, key := range candleKeys {

		// if key does not exist, seek most recent value
		_, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {

			prefixForMostRecent := append([]byte(string(key)[:len(string(key))-19]), 0xFF)
			options := badger.DefaultIteratorOptions
			options.Reverse = true
			it := txn.NewIterator(options)
			it.Seek(prefixForMostRecent)
			item := it.Item()
			it.Close()
			value, err := item.ValueCopy(nil)
			fmt.Printf("previousKey %+v\n", string(item.KeyCopy(nil)))
			if err != nil {
				return err
			}

			// extract close price from previous candle
			var previousCandle msg.Candle
			proto.Unmarshal(value, &previousCandle)

			fmt.Printf("previousCandle %+v\n", previousCandle)

			// generate new candle with extracted close price
			candleTimestamp := candleTimestamp[interval]
			newCandle := NewCandle(uint64(candleTimestamp), previousCandle.Close, 0)
			candleBuf, err := proto.Marshal(newCandle)
			if err != nil {
				return err
			}

			fmt.Printf("newCandle %+v\n", newCandle)

			// push new candle to the
			if err := txn.Set(key, candleBuf); err != nil {
				return err
			}
			//fmt.Printf("inserted\n")
		}
		//if present do nothing
		//fmt.Printf("candle for %s is present at key %s\n", interval, string(key))
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func NewCandle(timestamp, openPrice, size uint64) *msg.Candle {
	//TODO: get candle form pool of candles
	return &msg.Candle{Timestamp: timestamp, Open: openPrice, Low: openPrice, High: openPrice, Close:openPrice, Volume: size}
}

func UpdateCandle(candle *msg.Candle, trade *msg.Trade) {
	// always overwrite close price
	candle.Close = trade.Price
	// set minimum
	if trade.Price < candle.Low {
		candle.Low = trade.Price
	}
	// set maximum
	if trade.Price > candle.High {
		candle.High = trade.Price
	}
	candle.Volume += trade.Size
}

func generateCandleParamsForTimestamp(market string, timestamp uint64) (map[string]int64, map[string][]byte) {
	keys := make(map[string][]byte)
	timestamps := getMapOfIntervalsToTimestamps(int64(timestamp))

	for key, val := range timestamps {
		keys[key] = []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, key, val))
	}

	return timestamps,keys
}

func getMapOfIntervalsToTimestamps(timestamp int64) map[string]int64 {
	timestamps := make(map[string]int64)
	seconds := timestamp / int64(time.Second)
	nano := timestamp % seconds
	t := time.Unix(int64(seconds), nano)
	// round floor
	timestamps["1m"] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location()).UnixNano()
	timestamps["5m"] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/5)*5, 0, 0, t.Location()).UnixNano()
	timestamps["15m"] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/15)*15, 0, 0, t.Location()).UnixNano()
	timestamps["1h"] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location()).UnixNano()
	timestamps["6h"] = time.Date(t.Year(), t.Month(), t.Day(), (t.Hour()/6)*6, 0, 0, 0, t.Location()).UnixNano()
	timestamps["1d"] = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano()
	return timestamps
}

func generateFetchKey(market string, sinceTimestsamp uint64, interval string) []byte {
	seconds := sinceTimestsamp / uint64(time.Second)
	nano := sinceTimestsamp % seconds
	t := time.Unix(int64(seconds), int64(nano))
	fmt.Printf("\n\n")
	// round floor
	switch interval {
	case "1m":
		roundedToMinute := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedToMinute.UnixNano()))
	case "5m":
		roundedTo5Minutes := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/5)*5, 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedTo5Minutes.UnixNano()))
	case "15m":
		roundedTo15Minutes := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/15)*15, 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedTo15Minutes.UnixNano()))
	case "1h":
		roundedTo1Hour := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedTo1Hour.UnixNano()))
	case "6h":
		roundedTo6Hour := time.Date(t.Year(), t.Month(), t.Day(), (t.Hour()/6)*6, 0, 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedTo6Hour.UnixNano()))
	case "1d":
		roundedToDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		return []byte(fmt.Sprintf("M:%s_I:%s_T:%d", market, interval, roundedToDay.UnixNano()))
	}
	return nil
}