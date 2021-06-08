package wallet_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"code.vegaprotocol.io/vega/wallet"
	"code.vegaprotocol.io/vega/wallet/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerDeprecated(t *testing.T) {
	t.Run("sign any - success", testSignAnySuccess)
	t.Run("sign tx - success", testSignTxSuccess)
	t.Run("sign tx - failure key tainted", testSignTxFailure)
}

func testSignTxSuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return(name, nil)

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet(name, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, tok)

	key, err := h.GenerateKeypair(tok, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	message := "hello world."

	keyBytes, _ := hex.DecodeString(key)

	signedBundle, err := h.SignTx(tok, base64.StdEncoding.EncodeToString([]byte(message)), key)
	require.NoError(t, err)

	// verify signature then
	alg, err := crypto.NewSignatureAlgorithm(crypto.Ed25519)
	require.NoError(t, err)

	v, err := alg.Verify(keyBytes, signedBundle.Tx, signedBundle.Sig.Sig)
	require.NoError(t, err)
	assert.True(t, v)
}

func testSignAnySuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return(name, nil)

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet(name, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, tok)

	key, err := h.GenerateKeypair(tok, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	message := "@myTwitterHandle"

	keyBytes, _ := hex.DecodeString(key)

	signature, err := h.SignAny(
		tok, base64.StdEncoding.EncodeToString([]byte(message)), key)
	require.NoError(t, err)

	// verify signature then
	alg, err := crypto.NewSignatureAlgorithm(crypto.Ed25519)
	require.NoError(t, err)

	v, err := alg.Verify(keyBytes, []byte(message), signature)
	require.NoError(t, err)
	assert.True(t, v)
}

func testSignTxFailure(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return(name, nil)

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet(name, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, tok)

	key, err := h.GenerateKeypair(tok, passphrase)
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// taint the key
	err = h.TaintKey(tok, key, passphrase)
	require.NoError(t, err)

	message := "hello world."
	_, err = h.SignTx(tok, base64.StdEncoding.EncodeToString([]byte(message)), key)
	assert.EqualError(t, err, wallet.ErrPubKeyIsTainted.Error())
}
