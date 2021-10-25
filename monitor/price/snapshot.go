package price

import (
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
)

func internalBoundToPriceBoundType(b *bound) *types.PriceBound {
	return &types.PriceBound{
		Active:     b.Active,
		UpFactor:   b.UpFactor,
		DownFactor: b.DownFactor,
		Trigger:    b.Trigger.DeepClone(),
	}
}

func priceBoundTypeToInternal(pb *types.PriceBound) *bound {
	return &bound{
		Active:     pb.Active,
		UpFactor:   pb.UpFactor,
		DownFactor: pb.DownFactor,
		Trigger:    pb.Trigger.DeepClone(),
	}
}

func mapToKeyDecimalPair(m map[int64]num.Decimal) []*types.KeyDecimalPair {
	dm := make([]*types.KeyDecimalPair, 0, len(m))

	for k, v := range m {
		dm = append(dm, &types.KeyDecimalPair{
			Key: k,
			Val: v,
		})
	}

	return dm
}

func keyDecimalPairToMap(dms []*types.KeyDecimalPair) map[int64]num.Decimal {
	m := map[int64]num.Decimal{}

	for _, dm := range dms {
		m[dm.Key] = dm.Val
	}

	return m
}

func wrappedDecimalFromDecimal(d num.Decimal) num.WrappedDecimal {
	uit, _ := num.UintFromDecimal(d)
	return num.NewWrappedDecimal(uit, d)
}

func (e *Engine) restoreBounds(pbs []*types.PriceBound) {
	e.bounds = make([]*bound, 0, len(pbs))
	for _, pb := range pbs {
		e.bounds = append(e.bounds, priceBoundTypeToInternal(pb))
	}
}

func (e *Engine) serialiseBounds() []*types.PriceBound {
	bounds := make([]*types.PriceBound, 0, len(e.bounds))
	for _, b := range e.bounds {
		bounds = append(bounds, internalBoundToPriceBoundType(b))
	}

	return bounds
}

func (e *Engine) restorePriceRanges(prs []*types.PriceRangeCache) {
	for _, pr := range prs {
		e.priceRangesCache[priceBoundTypeToInternal(pr.Bound)] = priceRange{
			MinPrice:       wrappedDecimalFromDecimal(pr.Range.Min),
			MaxPrice:       wrappedDecimalFromDecimal(pr.Range.Max),
			ReferencePrice: pr.Range.Ref,
		}
	}
}

func (e Engine) serialisePriceRanges() []*types.PriceRangeCache {
	prc := make([]*types.PriceRangeCache, 0, len(e.priceRangesCache))
	for bound, priceRange := range e.priceRangesCache {
		prc = append(prc, &types.PriceRangeCache{
			Bound: internalBoundToPriceBoundType(bound),
			Range: &types.PriceRange{
				Min: priceRange.MinPrice.Original(),
				Max: priceRange.MaxPrice.Original(),
				Ref: priceRange.ReferencePrice,
			},
		})
	}
	return prc
}

func (e Engine) Changed() bool {
	return e.stateChanged
}

func (e *Engine) GetState() *types.PriceMonitor {
	pm := &types.PriceMonitor{
		Initialised:         e.initialised,
		FPHorizons:          mapToKeyDecimalPair(e.fpHorizons),
		Now:                 e.now,
		Update:              e.update,
		Bounds:              e.serialiseBounds(),
		PriceRangeCache:     e.serialisePriceRanges(),
		PriceRangeCacheTime: e.priceRangeCacheTime,
		RefPriceCache:       mapToKeyDecimalPair(e.refPriceCache),
		RefPriceCacheTime:   e.refPriceCacheTime,
	}

	e.stateChanged = false

	return pm
}

func (e *Engine) RestoreState(pm *types.PriceMonitor) {
	e.initialised = pm.Initialised
	e.fpHorizons = keyDecimalPairToMap(pm.FPHorizons)
	e.now = pm.Now
	e.update = pm.Update
	e.priceRangeCacheTime = pm.PriceRangeCacheTime
	e.refPriceCache = keyDecimalPairToMap(pm.RefPriceCache)
	e.refPriceCacheTime = pm.RefPriceCacheTime
	e.restoreBounds(pm.Bounds)
	e.restorePriceRanges(pm.PriceRangeCache)
}
