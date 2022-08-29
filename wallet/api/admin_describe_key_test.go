package api_test

import (
	"context"
	"fmt"
	"testing"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	vgrand "code.vegaprotocol.io/vega/libs/rand"
	"code.vegaprotocol.io/vega/wallet/api"
	"code.vegaprotocol.io/vega/wallet/api/mocks"
	"code.vegaprotocol.io/vega/wallet/wallet"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdminDescribeKey(t *testing.T) {
	t.Run("Describing a key with invalid params fails", testDescribingKeyWithInvalidParamsFails)
	t.Run("Describing a key with valid params succeeds", testDescribingKeyWithValidParamsSucceeds)
	t.Run("Describing a key from wallet that does not exists fails", testDescribingKeyFromWalletThatDoesNotExistsFails)
	t.Run("Getting internal error during wallet verification fails", testGettingInternalErrorDuringWalletVerificationFails)
	t.Run("Getting internal error during wallet retrieval fails", testGettingInternalErrorDuringWalletRetrievalFails)
	t.Run("Describing a key that does not exists fails", testDescribingKeyThatDoesNotExistsFails)
}

func testDescribingKeyWithInvalidParamsFails(t *testing.T) {
	tcs := []struct {
		name          string
		params        interface{}
		expectedError error
	}{
		{
			name:          "with nil params",
			params:        nil,
			expectedError: api.ErrParamsRequired,
		}, {
			name:          "with wrong type of params",
			params:        "test",
			expectedError: api.ErrParamsDoNotMatch,
		}, {
			name: "with empty name",
			params: api.DescribeKeyParams{
				Wallet:     "",
				Passphrase: vgrand.RandomStr(5),
				PublicKey:  "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			},
			expectedError: api.ErrWalletIsRequired,
		}, {
			name: "with empty passphrase",
			params: api.DescribeKeyParams{
				Wallet:     vgrand.RandomStr(5),
				Passphrase: "",
				PublicKey:  "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			},
			expectedError: api.ErrPassphraseIsRequired,
		}, {
			name: "with empty public key",
			params: api.DescribeKeyParams{
				Wallet:     vgrand.RandomStr(5),
				Passphrase: vgrand.RandomStr(5),
				PublicKey:  "",
			},
			expectedError: api.ErrPublicKeyIsRequired,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			ctx, _ := contextWithTraceID()

			// setup
			handler := newDescribeKeyHandler(tt)
			// -- unexpected calls
			handler.walletStore.EXPECT().WalletExists(gomock.Any(), gomock.Any()).Times(0)
			handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
			handler.walletStore.EXPECT().GetWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

			// when
			result, errorDetails := handler.handle(t, ctx, tc.params)

			// then
			require.Empty(tt, result)
			assertInvalidParams(tt, errorDetails, tc.expectedError)
		})
	}
}

func testDescribingKeyWithValidParamsSucceeds(t *testing.T) {
	// given
	ctx := context.Background()
	passphrase := vgrand.RandomStr(5)
	name := vgrand.RandomStr(5)
	expectedWallet, firstKey := walletWithKey(t)

	// setup
	handler := newDescribeKeyHandler(t)
	// -- expected calls
	handler.walletStore.EXPECT().WalletExists(ctx, name).Times(1).Return(true, nil)
	handler.walletStore.EXPECT().GetWallet(ctx, name, passphrase).Times(1).Return(expectedWallet, nil)
	// -- unexpected calls
	handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
	handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

	// when
	result, errorDetails := handler.handle(t, ctx, api.DescribeKeyParams{
		Wallet:     name,
		Passphrase: passphrase,
		PublicKey:  firstKey.PublicKey(),
	})

	// then
	require.Nil(t, errorDetails)
	assert.Equal(t, api.DescribeKeyResult{
		PublicKey: firstKey.PublicKey(),
		Algorithm: wallet.Algorithm{
			Name:    firstKey.AlgorithmName(),
			Version: firstKey.AlgorithmVersion(),
		},
		Metadata:  firstKey.Metadata(),
		IsTainted: firstKey.IsTainted(),
	}, result)
}

func testDescribingKeyFromWalletThatDoesNotExistsFails(t *testing.T) {
	// given
	ctx := context.Background()
	passphrase := vgrand.RandomStr(5)
	name := vgrand.RandomStr(5)

	// setup
	handler := newDescribeKeyHandler(t)
	// -- expected calls
	handler.walletStore.EXPECT().WalletExists(ctx, name).Times(1).Return(false, nil)
	// -- unexpected calls
	handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().GetWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
	handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

	// when
	result, errorDetails := handler.handle(t, ctx, api.DescribeKeyParams{
		Wallet:     name,
		Passphrase: passphrase,
		PublicKey:  vgrand.RandomStr(5),
	})

	// then
	require.NotNil(t, errorDetails)
	assert.Empty(t, result)
	assertInvalidParams(t, errorDetails, api.ErrWalletDoesNotExist)
}

func testGettingInternalErrorDuringWalletVerificationFails(t *testing.T) {
	// given
	ctx := context.Background()
	passphrase := vgrand.RandomStr(5)
	name := vgrand.RandomStr(5)

	// setup
	handler := newDescribeKeyHandler(t)
	// -- expected calls
	handler.walletStore.EXPECT().WalletExists(ctx, name).Times(1).Return(false, assert.AnError)
	// -- unexpected calls
	handler.walletStore.EXPECT().GetWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
	handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

	// when
	result, errorDetails := handler.handle(t, ctx, api.DescribeKeyParams{
		Wallet:     name,
		Passphrase: passphrase,
		PublicKey:  vgrand.RandomStr(5),
	})

	// then
	require.NotNil(t, errorDetails)
	assert.Empty(t, result)
	assertInternalError(t, errorDetails, fmt.Errorf("could not verify the wallet existence: %w", assert.AnError))
}

func testGettingInternalErrorDuringWalletRetrievalFails(t *testing.T) {
	// given
	ctx := context.Background()
	passphrase := vgrand.RandomStr(5)
	name := vgrand.RandomStr(5)

	// setup
	handler := newDescribeKeyHandler(t)
	// -- expected calls
	handler.walletStore.EXPECT().WalletExists(ctx, name).Times(1).Return(true, nil)
	handler.walletStore.EXPECT().GetWallet(ctx, name, passphrase).Times(1).Return(nil, assert.AnError)
	// -- unexpected calls
	handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
	handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

	// when
	result, errorDetails := handler.handle(t, ctx, api.DescribeKeyParams{
		Wallet:     name,
		Passphrase: passphrase,
		PublicKey:  vgrand.RandomStr(5),
	})

	// then
	require.NotNil(t, errorDetails)
	assert.Empty(t, result)
	assertInternalError(t, errorDetails, fmt.Errorf("could not retrieve the wallet: %w", assert.AnError))
}

func testDescribingKeyThatDoesNotExistsFails(t *testing.T) {
	// given
	ctx := context.Background()
	passphrase := vgrand.RandomStr(5)
	expectedWallet, _ := walletWithKey(t)

	// setup
	handler := newDescribeKeyHandler(t)
	// -- expected calls
	handler.walletStore.EXPECT().WalletExists(ctx, expectedWallet.Name()).Times(1).Return(true, nil)
	handler.walletStore.EXPECT().GetWallet(ctx, expectedWallet.Name(), passphrase).Times(1).Return(expectedWallet, nil)
	// -- unexpected calls
	handler.walletStore.EXPECT().SaveWallet(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	handler.walletStore.EXPECT().ListWallets(gomock.Any()).Times(0)
	handler.walletStore.EXPECT().DeleteWallet(gomock.Any(), gomock.Any()).Times(0)

	// when
	result, errorDetails := handler.handle(t, ctx, api.DescribeKeyParams{
		Wallet:     expectedWallet.Name(),
		Passphrase: passphrase,
		PublicKey:  vgrand.RandomStr(5),
	})

	// then
	require.NotNil(t, errorDetails)
	assert.Empty(t, result)
	assertInvalidParams(t, errorDetails, api.ErrPublicKeyDoesNotExist)
}

type describeKeyHandler struct {
	*api.DescribeKey
	ctrl        *gomock.Controller
	walletStore *mocks.MockWalletStore
	pipeline    *mocks.MockPipeline
}

func (h *describeKeyHandler) handle(t *testing.T, ctx context.Context, params interface{}) (api.DescribeKeyResult, *jsonrpc.ErrorDetails) {
	t.Helper()

	rawResult, err := h.Handle(ctx, params)
	if rawResult != nil {
		result, ok := rawResult.(api.DescribeKeyResult)
		if !ok {
			t.Fatal("DescribeKey handler result is not a DescribeKeyResult")
		}
		return result, err
	}
	return api.DescribeKeyResult{}, err
}

func newDescribeKeyHandler(t *testing.T) *describeKeyHandler {
	t.Helper()

	ctrl := gomock.NewController(t)
	walletStore := mocks.NewMockWalletStore(ctrl)

	return &describeKeyHandler{
		DescribeKey: api.NewDescribeKey(walletStore),
		ctrl:        ctrl,
		walletStore: walletStore,
	}
}
