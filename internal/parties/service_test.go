package parties

import (
	"context"
	"testing"

	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/parties/newmocks"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testService struct {
	Service
	log   *logging.Logger
	ctx   context.Context
	cfunc context.CancelFunc
	ctrl  *gomock.Controller
	store *newmocks.MockPartyStore
}

func getTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	store := newmocks.NewMockPartyStore(ctrl)
	log := logging.NewLoggerFromEnv("dev")
	ctx, cfunc := context.WithCancel(context.Background())
	svc, err := NewPartyService(
		NewDefaultConfig(log),
		store,
	)
	assert.NoError(t, err)
	return &testService{
		Service: svc,
		log:     log,
		ctx:     ctx,
		cfunc:   cfunc,
		ctrl:    ctrl,
		store:   store,
	}
}

func TestPartyService_CreateParty(t *testing.T) {
	svc := getTestService(t)
	defer svc.Finish()
	p := &types.Party{Name: "Christina"}

	svc.store.EXPECT().Post(p).Times(1).Return(nil)

	assert.NoError(t, svc.CreateParty(svc.ctx, p))
}

func TestPartyService_GetAll(t *testing.T) {
	svc := getTestService(t)
	defer svc.Finish()

	expected := []*types.Party{
		{Name: "Edd"},
		{Name: "Barney"},
		{Name: "Ramsey"},
		{Name: "Jeremy"},
	}

	svc.store.EXPECT().GetAll().Times(1).Return(expected, nil)

	parties, err := svc.GetAll(svc.ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, parties)
}

func TestPartyService_GetByName(t *testing.T) {
	svc := getTestService(t)
	defer svc.Finish()

	expect := &types.Party{
		Name: "Candida",
	}
	svc.store.EXPECT().GetByName(expect.Name).Times(1).Return(expect, nil)

	party, err := svc.GetByName(svc.ctx, expect.Name)
	assert.NoError(t, err)
	assert.Equal(t, expect, party)
}

func (t *testService) Finish() {
	t.log.Sync()
	t.cfunc()
	t.ctrl.Finish()
}
