package nodewallet_test

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"code.vegaprotocol.io/vega/config/encoding"
	"code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/nodewallet"
	"code.vegaprotocol.io/vega/nodewallet/eth"
	"code.vegaprotocol.io/vega/nodewallet/eth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	rootDirPath = "/tmp/vegatests/nodewallet/"
)

func rootDir() string {
	path := filepath.Join(rootDirPath, crypto.RandomStr(10))
	os.MkdirAll(path, os.ModePerm)
	return path
}

func TestNodeWallet(t *testing.T) {
	t.Run("is supported fail", testIsSupportedFail)
	t.Run("is supported success", testIsSupportedSuccess)
	t.Run("test init success as new node wallet", testInitSuccess)
	t.Run("test init failure as new node wallet", testInitFailure)
	t.Run("test devInit success", testDevInitSuccess)
	t.Run("verify success", testVerifySuccess)
	t.Run("verify failure", testVerifyFailure)
	t.Run("new failure invalid store path", testNewFailureInvalidStorePath)
	t.Run("new failure missing required wallets", testNewFailureMissingRequiredWallets)
	t.Run("new failure invalidPassphrase", testNewFailureInvalidPassphrase)
	t.Run("import new wallet", testImportNewWallet)
}

func testIsSupportedFail(t *testing.T) {
	err := nodewallet.IsSupported("yolocoin")
	assert.EqualError(t, err, "unsupported chain wallet yolocoin")
}

func testIsSupportedSuccess(t *testing.T) {
	err := nodewallet.IsSupported("vega")
	assert.NoError(t, err)
}

func testInitSuccess(t *testing.T) {
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	err := nodewallet.Init(filePath, "somepassphrase")
	assert.NoError(t, err)

	assert.NoError(t, os.RemoveAll(rootDir))
}

func testInitFailure(t *testing.T) {
	filePath := filepath.Join("/invalid/path/", "nodewalletstore")

	err := nodewallet.Init(filePath, "somepassphrase")
	assert.EqualError(t, err, "open /invalid/path/nodewalletstore: no such file or directory")
}

func testDevInitSuccess(t *testing.T) {
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	err := nodewallet.DevInit(filePath, rootDir, "somepassphrase")
	require.NoError(t, err)

	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: rootDir,
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()

	ethclt.EXPECT().ChainID(gomock.Any()).Times(1).Return(big.NewInt(42), nil)
	nw, err := nodewallet.New(logging.NewTestLogger(), cfg, "somepassphrase", ethclt)

	require.NoError(t, err)
	assert.NotNil(t, nw)

	w, ok := nw.Get(nodewallet.Ethereum)
	assert.NotNil(t, w)
	assert.True(t, ok)
	assert.Equal(t, string(nodewallet.Ethereum), w.Chain())

	w1, ok := nw.Get(nodewallet.Vega)
	assert.NotNil(t, w1)
	assert.True(t, ok)
	assert.Equal(t, string(nodewallet.Vega), w1.Chain())

	assert.NoError(t, os.RemoveAll(rootDir))
}

func testVerifySuccess(t *testing.T) {
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	// no error to generate
	err := nodewallet.DevInit(filePath, rootDir, "somepassphrase")
	assert.NoError(t, err)

	// try to instantiate a wallet from that
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: rootDir,
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()
	ethclt.EXPECT().ChainID(gomock.Any()).Times(1).Return(big.NewInt(42), nil)

	err = nodewallet.Verify(cfg, "somepassphrase", ethclt)
	assert.NoError(t, err)
	assert.NoError(t, os.RemoveAll(rootDir))
}

func testVerifyFailure(t *testing.T) {
	// create a random non existing path
	filePath := filepath.Join("/", crypto.RandomStr(10), "somewallet")
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: "",
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()

	err := nodewallet.Verify(cfg, "somepassphrase", ethclt)
	assert.Error(t, err)
}

func testNewFailureInvalidStorePath(t *testing.T) {
	// create a random non existing path
	filePath := filepath.Join("/", crypto.RandomStr(10), "somewallet")
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: "",
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()

	nw, err := nodewallet.New(logging.NewTestLogger(), cfg, "somepassphrase", ethclt)
	assert.Error(t, err)
	assert.Nil(t, nw)
}

func testNewFailureMissingRequiredWallets(t *testing.T) {
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	// no error to generate
	err := nodewallet.Init(filePath, "somepassphrase")
	assert.NoError(t, err)

	// try to instantiate a wallet from that
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: rootDir,
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()

	nw, err := nodewallet.New(logging.NewTestLogger(), cfg, "somepassphrase", ethclt)
	require.NoError(t, err)

	assert.EqualError(t, nw.EnsureRequireWallets(),
		"missing required wallet for vega chain",
	)
	assert.NoError(t, os.RemoveAll(rootDir))

}

func testImportNewWallet(t *testing.T) {
	ethDir := rootDir()
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	// no error to generate
	err := nodewallet.DevInit(filePath, rootDir, "somepassphrase")
	assert.NoError(t, err)

	// try to instantiate a wallet from that
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: rootDir,
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()
	ethclt.EXPECT().ChainID(gomock.Any()).Times(2).Return(big.NewInt(42), nil)

	nw, err := nodewallet.New(logging.NewTestLogger(), cfg, "somepassphrase", ethclt)
	require.NoError(t, err)
	assert.NotNil(t, nw)

	// now generate an eth wallet
	path, err := eth.DevInit(ethDir, "ethpassphrase")
	require.NoError(t, err)
	assert.NotEmpty(t, path)

	// import this new wallet
	err = nw.Import(string(nodewallet.Ethereum), "somepassphrase", "ethpassphrase", path)
	require.NoError(t, err)

	assert.NoError(t, os.RemoveAll(rootDir))
	assert.NoError(t, os.RemoveAll(ethDir))
}
func testNewFailureInvalidPassphrase(t *testing.T) {
	rootDir := rootDir()
	filePath := filepath.Join(rootDir, "nodewalletstore")

	// no error to generate
	err := nodewallet.Init(filePath, "somepassphrase")
	assert.NoError(t, err)

	// try to instantiate a wallet from that
	cfg := nodewallet.Config{
		Level:          encoding.LogLevel{},
		StorePath:      filePath,
		DevWalletsPath: rootDir,
	}

	ctrl := gomock.NewController(t)
	ethclt := mocks.NewMockETHClient(ctrl)
	defer ctrl.Finish()

	nw, err := nodewallet.New(logging.NewTestLogger(), cfg, "notthesamepassphrase", ethclt)
	assert.EqualError(t, err, "unable to load nodewalletsore: unable to decrypt store file (cipher: message authentication failed)")
	assert.Nil(t, nw)
	assert.NoError(t, os.RemoveAll(rootDir))
}
