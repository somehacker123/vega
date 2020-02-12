package wallet

import (
	"encoding/json"
	"testing"

	"code.vegaprotocol.io/vega/wallet/crypto"

	"github.com/stretchr/testify/assert"
)

func TestMarshalWallet(t *testing.T) {
	w := New("jeremy")
	w.Keypairs = append(w.Keypairs, NewKeypair(crypto.NewEd25519(), []byte{1, 2, 3, 4}, []byte{4, 3, 2, 1}))
	expected := `{"Owner":"jeremy","Keypairs":[{"Pub":"01020304","Priv":"04030201","Algorithm":"ed25519","Tainted":false,"Meta":null}]}`
	m, err := json.Marshal(&w)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(m))
}

func TestUnMarshalWallet(t *testing.T) {
	w := Wallet{}
	marshalled := `{"Owner":"jeremy","Keypairs":[{"Pub":"01020304","Priv":"04030201","Algorithm":"ed25519","Tainted":false,"Meta":null}]}`
	err := json.Unmarshal([]byte(marshalled), &w)
	assert.NoError(t, err)
	assert.Len(t, w.Keypairs, 1)
	assert.Equal(t, []byte{1, 2, 3, 4}, w.Keypairs[0].pubBytes)
	assert.Equal(t, []byte{4, 3, 2, 1}, w.Keypairs[0].privBytes)
	assert.Equal(t, "ed25519", w.Keypairs[0].Algorithm.Name())
}

func TestUnMarshalWalletErrorInvalidAlgorithm(t *testing.T) {
	w := Wallet{}
	marshalled := `{"Owner":"jeremy","Keypairs":[{"Pub":"01020304","Priv":"04030201","Algorithm":"notanalgorithm","Tainted":false,"Meta":null}]}`
	err := json.Unmarshal([]byte(marshalled), &w)
	assert.EqualError(t, err, crypto.ErrUnsupportedSignatureAlgorithm.Error())
}
