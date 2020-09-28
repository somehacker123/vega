package validators_test

import (
	"testing"

	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/validators"
	"code.vegaprotocol.io/vega/validators/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto"
)

func tmTestPubKey() testPubKey {
	return testPubKey{bytes: []byte("test-pub-key")}
}

type testTop struct {
	*validators.Topology
	ctrl   *gomock.Controller
	wallet *mocks.MockWallet
}

func getTestTop(t *testing.T) *testTop {
	ctrl := gomock.NewController(t)
	wallet := mocks.NewMockWallet(ctrl)

	top := validators.NewTopology(logging.NewTestLogger(), validators.NewDefaultConfig(), wallet, true)

	return &testTop{
		Topology: top,
		ctrl:     ctrl,
		wallet:   wallet,
	}
}

func TestValidatorTopology(t *testing.T) {
	t.Run("add node registration - success", testAddNodeRegistrationSuccess)
	t.Run("add node registration - failure", testAddNodeRegistrationFailure)
	t.Run("get len ", testGetLen)
	t.Run("exists", testExists)
}

func testAddNodeRegistrationSuccess(t *testing.T) {
	top := getTestTop(t)
	defer top.ctrl.Finish()
	top.UpdateValidatorSet([][]byte{
		tmTestPubKey().Bytes(),
	})

	nr := types.NodeRegistration{
		ChainPubKey: tmTestPubKey().bytes,
		PubKey:      []byte("vega-key"),
	}
	err := top.AddNodeRegistration(&nr)
	assert.NoError(t, err)
}

func testAddNodeRegistrationFailure(t *testing.T) {
	top := getTestTop(t)
	defer top.ctrl.Finish()
	top.UpdateValidatorSet([][]byte{
		tmTestPubKey().Bytes(),
	})

	nr := types.NodeRegistration{
		ChainPubKey: tmTestPubKey().bytes,
		PubKey:      []byte("vega-key"),
	}
	err := top.AddNodeRegistration(&nr)
	assert.NoError(t, err)

	nr = types.NodeRegistration{
		ChainPubKey: tmTestPubKey().bytes,
		PubKey:      []byte("vega-key-2"),
	}
	err = top.AddNodeRegistration(&nr)
	assert.Error(t, err)
}

func testGetLen(t *testing.T) {
	top := getTestTop(t)
	defer top.ctrl.Finish()
	top.UpdateValidatorSet([][]byte{
		tmTestPubKey().Bytes(),
	})

	// first len is 0
	assert.Equal(t, 0, top.Len())

	nr := types.NodeRegistration{
		ChainPubKey: tmTestPubKey().bytes,
		PubKey:      []byte("vega-key"),
	}
	err := top.AddNodeRegistration(&nr)
	assert.NoError(t, err)

	assert.Equal(t, 1, top.Len())
}

func testExists(t *testing.T) {
	top := getTestTop(t)
	defer top.ctrl.Finish()
	top.UpdateValidatorSet([][]byte{
		tmTestPubKey().Bytes(),
	})

	assert.False(t, top.Exists([]byte("vega-key")))

	nr := types.NodeRegistration{
		ChainPubKey: tmTestPubKey().bytes,
		PubKey:      []byte("vega-key"),
	}
	err := top.AddNodeRegistration(&nr)
	assert.NoError(t, err)

	assert.True(t, top.Exists([]byte("vega-key")))

}

type testPubKey struct {
	addr  crypto.Address
	bytes []byte
}

func (t testPubKey) Address() crypto.Address { return t.addr }

func (t testPubKey) Bytes() []byte                           { return t.bytes }
func (t testPubKey) VerifyBytes(msg []byte, sig []byte) bool { return true }
func (t testPubKey) Equals(crypto.PubKey) bool               { return false }
func (t testPubKey) Type() string                            { return "test-pk" }
