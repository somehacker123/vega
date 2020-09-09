package events

import (
	"context"

	types "code.vegaprotocol.io/vega/proto"
)

type Acc struct {
	*Base
	a types.Account
}

func NewAccountEvent(ctx context.Context, a types.Account) *Acc {
	return &Acc{
		Base: newBase(ctx, AccountEvent),
		a:    a,
	}
}

func (a *Acc) Account() types.Account {
	return a.a
}

func (a Acc) Proto() types.Account {
	return a.a
}

func (a Acc) StreamMessage() *types.BusEvent {
	return &types.BusEvent{
		ID:   a.traceID,
		Type: a.et.ToProto(),
		Event: &types.BusEvent_Account{
			Account: &a.a,
		},
	}
}
