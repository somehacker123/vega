// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package future

import (
	"context"

	"code.vegaprotocol.io/vega/core/types"
)

const (
	// PriceMoveMid used to indicate that the mid price has moved.
	PriceMoveMid = 1

	// PriceMoveBestBid used to indicate that the best bid price has moved.
	PriceMoveBestBid = 2

	// PriceMoveBestAsk used to indicate that the best ask price has moved.
	PriceMoveBestAsk = 4

	// PriceMoveAll used to indicate everything has moved.
	PriceMoveAll = PriceMoveMid + PriceMoveBestBid + PriceMoveBestAsk
)

type OrderReferenceCheck types.Order

func (o OrderReferenceCheck) HasMoved(changes uint8) bool {
	return (o.PeggedOrder.Reference == types.PeggedReferenceMid &&
		changes&PriceMoveMid > 0) ||
		(o.PeggedOrder.Reference == types.PeggedReferenceBestBid &&
			changes&PriceMoveBestBid > 0) ||
		(o.PeggedOrder.Reference == types.PeggedReferenceBestAsk &&
			changes&PriceMoveBestAsk > 0)
}

func (m *Market) checkForReferenceMoves(
	ctx context.Context, orderUpdates []*types.Order, forceUpdate bool,
) {
	if m.as.InAuction() {
		return
	}

	// will be set to non-nil if a peg is missing
	_, _, err := m.getBestStaticPricesDecimal()

	newBestBid, _ := m.getBestStaticBidPrice()
	newBestAsk, _ := m.getBestStaticAskPrice()
	newMidBuy, _ := m.getStaticMidPrice(types.SideBuy)
	newMidSell, _ := m.getStaticMidPrice(types.SideSell)

	// Look for a move
	var changes uint8
	if !forceUpdate {
		if newMidBuy.NEQ(m.lastMidBuyPrice) || newMidSell.NEQ(m.lastMidSellPrice) {
			changes |= PriceMoveMid
		}
		if newBestBid.NEQ(m.lastBestBidPrice) {
			changes |= PriceMoveBestBid
		}
		if newBestAsk.NEQ(m.lastBestAskPrice) {
			changes |= PriceMoveBestAsk
		}
	} else {
		changes = PriceMoveAll
	}

	// now we can start all special order repricing...
	if err == nil {
		minLpPrice, maxLpPrice := m.computeValidLPVolumeRange(newBestBid, newBestAsk)
		orderUpdates = m.repriceAllSpecialOrders(ctx, changes, orderUpdates, minLpPrice, maxLpPrice)
	} else {
		// we won't be able to reprice here
		m.stopAllSpecialOrders(ctx, orderUpdates)
		orderUpdates = nil
	}

	// Update the last price values
	// no need to clone the prices, they're not used in calculations anywhere in this function
	m.lastMidBuyPrice = newMidBuy
	m.lastMidSellPrice = newMidSell
	m.lastBestBidPrice = newBestBid
	m.lastBestAskPrice = newBestAsk

	// now we had new orderUpdates while processing those,
	// that would means someone got distressed, so some order
	// got uncrossed, so we need to check all these again.
	// we do not use the forceUpdate ffield here as it's
	// not required that prices moved though
	if len(orderUpdates) > 0 {
		m.checkForReferenceMoves(ctx, orderUpdates, false)
	}
}
