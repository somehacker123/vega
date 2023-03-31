-- +goose Up
ALTER TABLE markets ADD column quadratic_slippage_factor NUMERIC;
ALTER TABLE markets ADD column linear_slippage_factor NUMERIC;
ALTER TABLE markets_current ADD column quadratic_slippage_factor NUMERIC;
ALTER TABLE markets_current ADD column linear_slippage_factor NUMERIC;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_current_markets()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL AS
$$
BEGIN
    INSERT INTO markets_current(id,tx_hash,vega_time,instrument_id,tradable_instrument,decimal_places,fees,opening_auction,price_monitoring_settings,liquidity_monitoring_parameters,trading_mode,state,market_timestamps,position_decimal_places,lp_price_range, linear_slippage_factor, quadratic_slippage_factor)
    VALUES(NEW.id,NEW.tx_hash,NEW.vega_time,NEW.instrument_id,NEW.tradable_instrument,NEW.decimal_places,NEW.fees,NEW.opening_auction,NEW.price_monitoring_settings,NEW.liquidity_monitoring_parameters,NEW.trading_mode,NEW.state,NEW.market_timestamps,NEW.position_decimal_places,NEW.lp_price_range, NEW.linear_slippage_factor, NEW.quadratic_slippage_factor)
    ON CONFLICT(id) DO UPDATE SET
                                                           tx_hash=EXCLUDED.tx_hash,
                                                           instrument_id=EXCLUDED.instrument_id,
                                                           tradable_instrument=EXCLUDED.tradable_instrument,
                                                           decimal_places=EXCLUDED.decimal_places,
                                                           fees=EXCLUDED.fees,
                                                           opening_auction=EXCLUDED.opening_auction,
                                                           price_monitoring_settings=EXCLUDED.price_monitoring_settings,
                                                           liquidity_monitoring_parameters=EXCLUDED.liquidity_monitoring_parameters,
                                                           trading_mode=EXCLUDED.trading_mode,
                                                           state=EXCLUDED.state,
                                                           market_timestamps=EXCLUDED.market_timestamps,
                                                           position_decimal_places=EXCLUDED.position_decimal_places,
                                                           lp_price_range=EXCLUDED.lp_price_range,
                                                           linear_slippage_factor=EXCLUDED.linear_slippage_factor,
                                                           quadratic_slippage_factor=EXCLUDED.quadratic_slippage_factor,
                                                           vega_time=EXCLUDED.vega_time;
    RETURN NULL;
END;
$$;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_current_markets()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL AS
$$
BEGIN
    INSERT INTO markets_current(id,tx_hash,vega_time,instrument_id,tradable_instrument,decimal_places,fees,opening_auction,price_monitoring_settings,liquidity_monitoring_parameters,trading_mode,state,market_timestamps,position_decimal_places,lp_price_range)
    VALUES(NEW.id,NEW.tx_hash,NEW.vega_time,NEW.instrument_id,NEW.tradable_instrument,NEW.decimal_places,NEW.fees,NEW.opening_auction,NEW.price_monitoring_settings,NEW.liquidity_monitoring_parameters,NEW.trading_mode,NEW.state,NEW.market_timestamps,NEW.position_decimal_places,NEW.lp_price_range)
    ON CONFLICT(id) DO UPDATE SET
                                                           tx_hash=EXCLUDED.tx_hash,
                                                           instrument_id=EXCLUDED.instrument_id,
                                                           tradable_instrument=EXCLUDED.tradable_instrument,
                                                           decimal_places=EXCLUDED.decimal_places,
                                                           fees=EXCLUDED.fees,
                                                           opening_auction=EXCLUDED.opening_auction,
                                                           price_monitoring_settings=EXCLUDED.price_monitoring_settings,
                                                           liquidity_monitoring_parameters=EXCLUDED.liquidity_monitoring_parameters,
                                                           trading_mode=EXCLUDED.trading_mode,
                                                           state=EXCLUDED.state,
                                                           market_timestamps=EXCLUDED.market_timestamps,
                                                           position_decimal_places=EXCLUDED.position_decimal_places,
                                                           lp_price_range=EXCLUDED.lp_price_range,
                                                           vega_time=EXCLUDED.vega_time;
    RETURN NULL;
END;
$$;
-- +goose StatementEnd

ALTER TABLE markets DROP column IF EXISTS quadratic_slippage_factor;
ALTER TABLE markets DROP column IF EXISTS linear_slippage_factor;
ALTER TABLE markets_current DROP column IF EXISTS quadratic_slippage_factor;
ALTER TABLE markets_current DROP column IF EXISTS linear_slippage_factor;

