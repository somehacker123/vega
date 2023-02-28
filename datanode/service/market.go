// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package service

import (
	"context"
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/libs/num"
)

var nilPagination = entities.OffsetPagination{}

type MarketStore interface {
	Upsert(ctx context.Context, market *entities.Market) error
	GetByID(ctx context.Context, marketID string) (entities.Market, error)
	GetAll(ctx context.Context, pagination entities.OffsetPagination) ([]entities.Market, error)
	GetAllPaged(ctx context.Context, marketID string, pagination entities.CursorPagination, includeSettled bool) ([]entities.Market, entities.PageInfo, error)
}

type Markets struct {
	store     MarketStore
	cache     map[entities.MarketID]*entities.Market
	cacheLock sync.RWMutex
	sf        map[entities.MarketID]num.Decimal
}

func NewMarkets(store MarketStore) *Markets {
	return &Markets{
		store: store,
		cache: make(map[entities.MarketID]*entities.Market),
		sf:    map[entities.MarketID]num.Decimal{},
	}
}

func (m *Markets) Initialise(ctx context.Context) error {
	m.cacheLock.Lock()
	defer m.cacheLock.Unlock()

	all, err := m.store.GetAll(ctx, entities.OffsetPagination{})
	if err != nil {
		return err
	}
	for i := 0; i < len(all); i++ {
		m.cache[all[i].ID] = &all[i]
		m.sf[all[i].ID] = num.DecimalFromFloat(10).Pow(num.DecimalFromInt64(int64(all[i].PositionDecimalPlaces)))
	}
	return nil
}

func (m *Markets) Upsert(ctx context.Context, market *entities.Market) error {
	if err := m.store.Upsert(ctx, market); err != nil {
		return err
	}
	m.cacheLock.Lock()
	m.cache[market.ID] = market
	if market.State == entities.MarketStateSettled || market.State == entities.MarketStateRejected {
		// a settled or rejected market can be safely removed from this map.
		delete(m.sf, market.ID)
	} else {
		// just in case this gets updated, or the market is new.
		m.sf[market.ID] = num.DecimalFromFloat(10).Pow(num.DecimalFromInt64(int64(market.PositionDecimalPlaces)))
	}
	m.cacheLock.Unlock()
	return nil
}

func (m *Markets) GetByID(ctx context.Context, marketID string) (entities.Market, error) {
	m.cacheLock.RLock()
	defer m.cacheLock.RUnlock()

	data, ok := m.cache[entities.MarketID(marketID)]
	if !ok {
		return entities.Market{}, fmt.Errorf("no such market: %v", marketID)
	}
	return *data, nil
}

func (m *Markets) GetMarketScalingFactor(ctx context.Context, marketID string) (num.Decimal, bool) {
	m.cacheLock.RLock()
	defer m.cacheLock.RUnlock()
	pf, ok := m.sf[entities.MarketID(marketID)]
	return pf, ok
}

func (m *Markets) GetAll(ctx context.Context, pagination entities.OffsetPagination) ([]entities.Market, error) {
	if pagination != nilPagination {
		return m.store.GetAll(ctx, pagination)
	}

	m.cacheLock.RLock()
	defer m.cacheLock.RUnlock()

	data := make([]entities.Market, 0, len(m.cache))
	for _, v := range m.cache {
		data = append(data, *v)
	}
	return data, nil
}

func (m *Markets) GetAllPaged(ctx context.Context, marketID string, pagination entities.CursorPagination, includeSettled bool) ([]entities.Market, entities.PageInfo, error) {
	return m.store.GetAllPaged(ctx, marketID, pagination, includeSettled)
}
