package validators_test

import (
	"context"
	"fmt"
	"testing"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

	"github.com/stretchr/testify/assert"
)

func addNodes(top *testTop, number int) {
	tmPubKeys := make([]string, 0, number)

	for i := 0; i < number; i++ {
		tmPubKeys = append(tmPubKeys, fmt.Sprintf("tm-pub-key-%d", i))
	}

	top.UpdateValidatorSet(tmPubKeys)

	ctx := context.Background()

	for i := 0; i < number; i++ {
		top.AddNodeRegistration(ctx, &commandspb.NodeRegistration{
			Id:              fmt.Sprintf("vega-master-pubkey-%d", i),
			ChainPubKey:     tmPubKeys[0],
			VegaPubKey:      fmt.Sprintf("vega-key-%d", i),
			EthereumAddress: fmt.Sprintf("eth-address-%d", i),
		})
	}
}

func TestTopologyCheckpoint(t *testing.T) {
	t.Run("test checkpoint success", testTopologyCheckpointSuccess)
	t.Run("test checkpoint uses relative block height", testTopologyCheckpointUsesRelativeBlockHeight)
}

func testTopologyCheckpointSuccess(t *testing.T) {
	top := getTestTopWithDefaultValidator(t)
	defer top.ctrl.Finish()

	ctx := context.Background()
	addNodes(top, 2)

	kr1 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       10,
		NewPubKey:         "new-vega-key",
		CurrentPubKeyHash: hashKey("vega-key-0"),
	}
	err := top.AddKeyRotate(ctx, "vega-master-pubkey-0", 5, kr1)
	assert.NoError(t, err)

	kr2 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       11,
		NewPubKey:         "new-vega-key-1",
		CurrentPubKeyHash: hashKey("vega-key-1"),
	}
	err = top.AddKeyRotate(ctx, "vega-master-pubkey-1", 5, kr2)
	assert.NoError(t, err)

	pkrs := top.GetAllPendingKeyRotations()
	assert.Len(t, pkrs, 2)

	ckp, err := top.Checkpoint()
	assert.NotEmpty(t, ckp)
	assert.NoError(t, err)

	newTop := getTestTopWithDefaultValidator(t)
	defer newTop.ctrl.Finish()

	addNodes(newTop, 2)
	newTop.Load(ctx, ckp)

	newPkrs := newTop.GetAllPendingKeyRotations()
	assert.Len(t, newPkrs, 2)
	assert.Equal(t, pkrs, newPkrs)
}

func testTopologyCheckpointUsesRelativeBlockHeight(t *testing.T) {
	top := getTestTopWithDefaultValidator(t)
	defer top.ctrl.Finish()

	ctx := context.Background()
	addNodes(top, 2)

	kr1 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       105,
		NewPubKey:         "new-vega-key",
		CurrentPubKeyHash: hashKey("vega-key-0"),
	}
	err := top.AddKeyRotate(ctx, "vega-master-pubkey-0", 5, kr1)
	assert.NoError(t, err)

	kr2 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       115,
		NewPubKey:         "new-vega-key-1",
		CurrentPubKeyHash: hashKey("vega-key-1"),
	}
	err = top.AddKeyRotate(ctx, "vega-master-pubkey-1", 5, kr2)
	assert.NoError(t, err)

	pkrs := top.GetAllPendingKeyRotations()
	assert.Len(t, pkrs, 2)

	ckp, err := top.Checkpoint()
	assert.NotEmpty(t, ckp)
	assert.NoError(t, err)

	newTop := getTestTopWithDefaultValidator(t)
	defer newTop.ctrl.Finish()

	addNodes(newTop, 2)

	var newNetworkBlockHeight uint64 = 100

	// set current block height to newNetworkBlockHeight
	newTop.BeginBlock(ctx, newNetworkBlockHeight)

	newTop.Load(ctx, ckp)

	newPkrs := newTop.GetAllPendingKeyRotations()
	assert.Len(t, newPkrs, 2)

	assert.Equal(t, pkrs[0].BlockHeight+newNetworkBlockHeight, newPkrs[0].BlockHeight)
	assert.Equal(t, pkrs[1].BlockHeight+newNetworkBlockHeight, newPkrs[1].BlockHeight)
}
