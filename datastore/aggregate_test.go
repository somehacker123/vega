package datastore

import (
	"fmt"
	"math/rand"
	"testing"
	"vega/msg"

	"github.com/stretchr/testify/assert"
)

type TestOrderAndTrades struct {
	order *Order
	trade *Trade
}

func generateRandomOrderAndTrade(price, size, timestamp uint64) *TestOrderAndTrades {
	orderId := fmt.Sprintf("%d", rand.Intn(1000000000000))
	tradeId := fmt.Sprintf("%d", rand.Intn(1000000000000))
	d := &TestOrderAndTrades{
		&Order{
			Order: msg.Order{
				Id:     orderId,
				Market: testMarket,
				Price:     price,
				Size:      size,
				Remaining: size,
				Timestamp: timestamp,
			},
		},
		&Trade{
			Trade: msg.Trade{
				Id:        tradeId,
				Price:     price,
				Market:    testMarket,
				Size:      size,
				Timestamp: timestamp,
			},
			OrderId: orderId,
		},
	}
	return d
}

func TestMemTradeStore_GetCandles(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	price := uint64(100)
	timestamp := uint64(0)
	for i := 0; i < 100; i++ {
		if rand.Intn(3) == 1{
			price--
		} else {
			price++
		}

		if rand.Intn(5) == 1 {
			timestamp++
		}
		size := uint64(rand.Intn(400) + 800)

		// simulate timestamp gap
		if i == 10 {
			i = 15
			timestamp += 5
		}
		d := generateRandomOrderAndTrade(price, size, timestamp)

		err := newOrderStore.Post(*d.order)
		assert.Nil(t, err)
		err = newTradeStore.Post(*d.trade)
		assert.Nil(t, err)
	}

	candles, err := newTradeStore.GetCandles(testMarket, 0, timestamp, 3)
	fmt.Printf("candles returned:\n")
	for idx, c := range candles.Candles {
		fmt.Printf("%d %+v\n", idx, *c)
	}
	assertCandleIsEmpty(t, candles.Candles[2])
	fmt.Println()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(candles.Candles))


	candles, err = newTradeStore.GetCandles(testMarket, 5, timestamp, 3)
	fmt.Printf("candles returned:\n")
	for idx, c := range candles.Candles {
		fmt.Printf("%d %+v\n", idx, *c)
	}
	fmt.Println()
	assert.Nil(t, err)
	assert.Equal(t, 8, len(candles.Candles))
	assertCandleIsEmpty(t, candles.Candles[0])

	candles, err = newTradeStore.GetCandles(testMarket, 5, timestamp, 2)
	fmt.Printf("candles returned:\n")
	for idx, c := range candles.Candles {
		fmt.Printf("%d %+v\n", idx, *c)
	}
	fmt.Println()
	assert.Nil(t, err)
	assert.Equal(t, 12, len(candles.Candles))
	assertCandleIsEmpty(t, candles.Candles[0])
	assertCandleIsEmpty(t, candles.Candles[1])

	candles, err = newTradeStore.GetCandles(testMarket, 10, timestamp, 2)
	fmt.Printf("candles returned:\n")
	for idx, c := range candles.Candles {
		fmt.Printf("%d %+v\n", idx, *c)
	}
	fmt.Println()
	assert.Nil(t, err)
	assert.Equal(t, 9, len(candles.Candles))

}

func assertCandleIsEmpty(t assert.TestingT, candle *msg.Candle) {
	assert.Equal(t, uint64(0), candle.Volume)
	assert.Equal(t, uint64(0), candle.High)
	assert.Equal(t, uint64(0), candle.Low)
	assert.Equal(t, uint64(0), candle.Open)
	assert.Equal(t, uint64(0), candle.Close)
}

func TestMemOrderStore_GetOrderBookDepth(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	price := uint64(100)
	timestamp := uint64(0)
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			timestamp++
			price++
		}
		o := &Order{
			msg.Order{
				Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
				Market: testMarket,
				Side: msg.Side_Sell,
				Price:     price,
				Size:      uint64(100),
				Remaining: uint64(100),
				Timestamp: timestamp,
			},
		}

		err := newOrderStore.Post(*o)
		assert.Nil(t, err)
	}

	timestamp = 0
	price = 100
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			timestamp++
			price++
		}
		o := &Order{
			msg.Order{
				Id:     fmt.Sprintf("%d", rand.Intn(1000000000000)),
				Market: testMarket,
				Side: msg.Side_Buy,
				Price:     price,
				Size:      uint64(100),
				Remaining: uint64(100),
				Timestamp: timestamp,
			},
		}

		err := newOrderStore.Post(*o)
		assert.Nil(t, err)
	}

	orderBookDepth, err := newOrderStore.GetOrderBookDepth(testMarket)
	assert.Nil(t, err)
	fmt.Printf("orderBookDepth for buy side:\n")
	for idx, priceLevel := range orderBookDepth.Buy {
		fmt.Printf("%d %+v\n", idx, *priceLevel)
	}
	assert.Equal(t, orderBookDepth.Buy[0].Price, uint64(150))
	assert.Equal(t, orderBookDepth.Buy[len(orderBookDepth.Buy)-1].Price, uint64(100))
	assert.Equal(t, orderBookDepth.Buy[0].CumulativeVolume, orderBookDepth.Buy[0].Volume)
	assert.Equal(t, orderBookDepth.Buy[len(orderBookDepth.Buy)-1].CumulativeVolume, uint64(100*100))

	fmt.Printf("orderBookDepth for sell side:\n")
	for idx, priceLevel := range orderBookDepth.Sell {
		fmt.Printf("%d %+v\n", idx, *priceLevel)
	}

	assert.Equal(t, orderBookDepth.Sell[0].Price, uint64(100))
	assert.Equal(t, orderBookDepth.Sell[len(orderBookDepth.Sell)-1].Price, uint64(150))

	assert.Equal(t, orderBookDepth.Sell[0].CumulativeVolume, orderBookDepth.Sell[0].Volume)
	assert.Equal(t, orderBookDepth.Sell[len(orderBookDepth.Sell)-1].CumulativeVolume, uint64(100*100))
}
