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

package sqlstore_test

import (
	"testing"
	"time"

	"code.vegaprotocol.io/vega/datanode/entities"
	"code.vegaprotocol.io/vega/datanode/sqlstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNetworkLimits(t *testing.T) {
	ctx, rollback := tempTransaction(t)
	defer rollback()

	bs := sqlstore.NewBlocks(connectionSource)
	block := addTestBlock(t, ctx, bs)
	block2 := addTestBlock(t, ctx, bs)
	nls := sqlstore.NewNetworkLimits(connectionSource)

	nl := entities.NetworkLimits{
		VegaTime:                 block.VegaTime,
		CanProposeMarket:         true,
		CanProposeAsset:          false,
		ProposeMarketEnabled:     false,
		ProposeAssetEnabled:      true,
		GenesisLoaded:            false,
		ProposeMarketEnabledFrom: time.Now().Truncate(time.Millisecond),
		ProposeAssetEnabledFrom:  time.Now().Add(time.Second * 10).Truncate(time.Millisecond),
	}

	err := nls.Add(ctx, nl)
	require.NoError(t, err)

	nl2 := nl
	nl2.VegaTime = block2.VegaTime
	err = nls.Add(ctx, nl2)
	require.NoError(t, err)

	fetched, err := nls.GetLatest(ctx)
	require.NoError(t, err)
	assert.Equal(t, nl2, fetched)
}
