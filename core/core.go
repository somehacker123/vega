package core

import (
	"vega/matching"
)

type Vega struct {
	config  Config
	markets map[string]*matching.OrderBook
	orders  map[string]*matching.OrderEntry
}

type Config struct {
	Matching matching.Config
}

func New(config Config) *Vega {
	return &Vega{
		config:  config,
		markets: make(map[string]*matching.OrderBook),
		orders:  make(map[string]*matching.OrderEntry),
	}
}

func DefaultConfig() Config {
	return Config{
		Matching: matching.DefaultConfig(),
	}
}
