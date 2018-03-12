package main

import (
	"fmt"

	"proto"
	"vega"
)

func main() {
	app := vega.New()
	app.CreateMarket("BTC/DEC18")

	app.SubmitOrder(msg.Order{
		Market:    "BTC/DEC18",
		Party:     "A",
		Side:      msg.Side_Buy,
		Price:     100,
		Size:      50,
		Remaining: 50,
		Type:      msg.Order_GTC,
		Timestamp: 0,
	})

	app.SubmitOrder(msg.Order{
		Market:    "BTC/DEC18",
		Party:     "B",
		Side:      msg.Side_Buy,
		Price:     102,
		Size:      44,
		Remaining: 44,
		Type:      msg.Order_GTC,
		Timestamp: 0,
	})

	app.SubmitOrder(msg.Order{
		Market:    "BTC/DEC18",
		Party:     "B",
		Side:      msg.Side_Buy,
		Price:     99,
		Size:      42,
		Remaining: 42,
		Type:      msg.Order_GTC,
		Timestamp: 0,
	})

	res, _ :=
		app.SubmitOrder(msg.Order{
		Market:    "BTC/DEC18",
		Party:     "D",
		Side:      msg.Side_Sell,
		Price:     110,
		Size:      100,
		Remaining: 100,
		Type:      msg.Order_GTC,
		Timestamp: 0,
	})

	app.SubmitOrder(msg.Order{
		Market:    "BTC/DEC18",
		Party:     "C",
		Side:      msg.Side_Sell,
		Price:     98,
		Size:      120,
		Remaining: 120,
		Type:      msg.Order_GTC,
		Timestamp: 0,
	})

	app.DeleteOrder(res.Order.Id)

	fmt.Println(app.GetMarketData("BTC/DEC18"))
}
