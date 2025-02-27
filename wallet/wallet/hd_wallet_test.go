package wallet_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	vgrand "code.vegaprotocol.io/vega/libs/rand"
	"code.vegaprotocol.io/vega/wallet/wallet"
	"code.vegaprotocol.io/vega/wallet/wallets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestRecoveryPhrase1 = "swing ceiling chaos green put insane ripple desk match tip melt usual shrug turkey renew icon parade veteran lens govern path rough page render"
	FirstPubKeyV1       = "30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371"
	FirstPubKeyV2       = "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0"
)

func TestHDWallet(t *testing.T) {
	t.Run("Creating wallet succeeds", testHDWalletCreateWalletSucceeds)
	t.Run("Importing wallet succeeds", testHDWalletImportingWalletSucceeds)
	t.Run("Importing wallet with invalid recovery phrase fails", testHDWalletImportingWalletWithInvalidRecoveryPhraseFails)
	t.Run("Importing wallet with unsupported key derivation version fails", testHDWalletImportingWalletWithUnsupportedKeyDerivationVersionFails)
	t.Run("Generating key pair succeeds", testHDWalletGeneratingKeyPairSucceeds)
	t.Run("Generating key pair on isolated wallet fails", testHDWalletGeneratingKeyPairOnIsolatedWalletFails)
	t.Run("Tainting key pair succeeds", testHDWalletTaintingKeyPairSucceeds)
	t.Run("Tainting key pair removes the key from the allowed keys", testHDWalletTaintingKeyPairRemovesTheKeyFromAllowedKeys)
	t.Run("Tainting the last key pair voids the permissions", testHDWalletTaintingLastKeyPairVoidsThePermissions)
	t.Run("Tainting key pair that is already tainted fails", testHDWalletTaintingKeyThatIsAlreadyTaintedFails)
	t.Run("Tainting unknown key pair fails", testHDWalletTaintingUnknownKeyFails)
	t.Run("Untainting key pair succeeds", testHDWalletUntaintingKeyPairSucceeds)
	t.Run("Untainting key pair that is not tainted fails", testHDWalletUntaintingKeyThatIsNotTaintedFails)
	t.Run("Untainting unknown key pair fails", testHDWalletUntaintingUnknownKeyFails)
	t.Run("Updating key pair metadata succeeds", testHDWalletUpdatingKeyPairMetaSucceeds)
	t.Run("Updating key pair metadata with unknown public key fails", testHDWalletUpdatingKeyPairMetaWithUnknownPublicKeyFails)
	t.Run("Describing public key succeeds", testHDWalletDescribingPublicKeysSucceeds)
	t.Run("Describing unknown public key fails", testHDWalletDescribingUnknownPublicKeysFails)
	t.Run("Listing public keys succeeds", testHDWalletListingPublicKeysSucceeds)
	t.Run("Listing key pairs succeeds", testHDWalletListingKeyPairsSucceeds)
	t.Run("Signing transaction request succeeds", testHDWalletSigningTxSucceeds)
	t.Run("Signing transaction request with tainted key fails", testHDWalletSigningTxWithTaintedKeyFails)
	t.Run("Signing transaction request with unknown key fails", testHDWalletSigningTxWithUnknownKeyFails)
	t.Run("Signing any message succeeds", testHDWalletSigningAnyMessageSucceeds)
	t.Run("Signing any message with tainted key fails", testHDWalletSigningAnyMessageWithTaintedKeyFails)
	t.Run("Signing any message with unknown key fails", testHDWalletSigningAnyMessageWithUnknownKeyFails)
	t.Run("Verifying any message succeeds", testHDWalletVerifyingAnyMessageSucceeds)
	t.Run("Verifying any message with unknown key fails", testHDWalletVerifyingAnyMessageWithUnknownKeyFails)
	t.Run("Marshaling wallet succeeds", testHDWalletMarshalingWalletSucceeds)
	t.Run("Marshaling isolated wallet succeeds", testHDWalletMarshalingIsolatedWalletSucceeds)
	t.Run("Unmarshaling wallet succeeds", testHDWalletUnmarshalingWalletSucceeds)
	t.Run("Getting wallet info succeeds", testHDWalletGettingWalletInfoSucceeds)
	t.Run("Getting isolated wallet info succeeds", testHDWalletGettingIsolatedWalletInfoSucceeds)
	t.Run("Isolating wallet succeeds", testHDWalletIsolatingWalletSucceeds)
	t.Run("Isolating wallet with tainted key pair fails", testHDWalletIsolatingWalletWithTaintedKeyPairFails)
	t.Run("Isolating wallet with non-existing key pair fails", testHDWalletIsolatingWalletWithNonExistingKeyPairFails)
	t.Run("Getting master key pair succeeds", testHDWalletGettingWalletMasterKeySucceeds)
	t.Run("Updating permissions with inconsistent setup fails", testHDWalletUpdatingPermissionsWithInconsistentSetupFails)
	t.Run("Updating permissions with empty permission fallback to default", testHDWalletUpdatingPermissionsWithEmptyPermissionFallbackToDefaults)
	t.Run("Updating public keys permissions without keys in wallet fails", testHDWalletUpdatingPermissionsWithoutKeysInWalletFails)
	t.Run("Updating public keys permissions with tainted keys only fails", testHDWalletUpdatingPermissionsWithTaintedKeysOnlyFails)
	t.Run("Revoking permissions succeeds", testHDWalletRevokingPermissionsSucceeds)
	t.Run("Purging permissions succeeds", testHDWalletPurgingPermissionsSucceeds)
}

func testHDWalletCreateWalletSucceeds(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, recoveryPhrase, err := wallet.NewHDWallet(name)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, recoveryPhrase)
	assert.NotNil(t, w)
	assert.Equal(t, uint32(2), w.KeyDerivationVersion())
}

func testHDWalletImportingWalletSucceeds(t *testing.T) {
	tcs := []struct {
		name        string
		version     uint32
		expectedKey string
	}{
		{
			name:        "version 1",
			version:     1,
			expectedKey: FirstPubKeyV1,
		}, {
			name:        "version 2",
			version:     2,
			expectedKey: FirstPubKeyV2,
		},
	}

	whitespaces := " \t\n\r"

	genWhiteSpace := func() string {
		var ws string
		whiteSpaceIndex := rand.Intn(len(whitespaces))
		numChar := rand.Intn(5) + 1
		for i := 0; i < numChar; i++ {
			ws += string(whitespaces[whiteSpaceIndex])
		}
		return ws
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// We mess up the recovery phrase structure to ensure whitespaces
			// are not an actual problem.
			words := strings.Split(TestRecoveryPhrase1, " ")
			tamperedRecoveryPhrase := "  "
			for _, word := range words {
				tamperedRecoveryPhrase += word + genWhiteSpace()
			}

			// when
			w, err := wallet.ImportHDWallet(name, tamperedRecoveryPhrase, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)
			assert.Equal(tt, tc.version, w.KeyDerivationVersion())

			// when
			kp, err := w.GenerateKeyPair(nil)

			// then
			require.NoError(tt, err)
			assert.Equal(tt, tc.expectedKey, kp.PublicKey())
		})
	}
}

func testHDWalletImportingWalletWithInvalidRecoveryPhraseFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, "vladimir harkonnen doesn't like trees", tc.version)

			// then
			require.ErrorIs(tt, err, wallet.ErrInvalidRecoveryPhrase)
			assert.Nil(tt, w)
		})
	}
}

func testHDWalletImportingWalletWithUnsupportedKeyDerivationVersionFails(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 3)

	// then
	require.ErrorIs(t, err, wallet.NewUnsupportedWalletVersionError(3))
	assert.Nil(t, w)
}

func testHDWalletGeneratingKeyPairSucceeds(t *testing.T) {
	tcs := []struct {
		name       string
		version    uint32
		publicKey  string
		privateKey string
	}{
		{
			name:       "version 1",
			version:    1,
			publicKey:  "30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371",
			privateKey: "1bbd4efb460d0bf457251e866697d5d2e9b58c5dcb96a964cd9cfff1a712a2b930ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371",
		}, {
			name:       "version 2",
			version:    2,
			publicKey:  "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			privateKey: "0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			meta := []wallet.Metadata{{Key: "env", Value: "test"}}
			expectedMeta := []wallet.Metadata{meta[0], {
				Key:   "name",
				Value: "Key 1",
			}}

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair(meta)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)
			assert.Equal(tt, expectedMeta, kp.Metadata())
			assert.Equal(tt, tc.publicKey, kp.PublicKey())
			assert.Equal(tt, tc.privateKey, kp.PrivateKey())
		})
	}
}

func testHDWalletGeneratingKeyPairOnIsolatedWalletFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			isolatedWallet, err := w.IsolateWithKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, isolatedWallet)

			// when
			keyPair, err := isolatedWallet.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.ErrorIs(tt, err, wallet.ErrIsolatedWalletCantGenerateKeys)
			require.Nil(tt, keyPair)
		})
	}
}

func testHDWalletTaintingKeyPairSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			pubKey, err := w.DescribePublicKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, pubKey)
			assert.True(tt, pubKey.IsTainted())
		})
	}
}

func testHDWalletTaintingKeyPairRemovesTheKeyFromAllowedKeys(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	assert.NotNil(t, w)

	// when
	kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

	// then
	require.NoError(t, err)
	assert.NotNil(t, kp1)

	// when
	kp2, err := w.GenerateKeyPair([]wallet.Metadata{})

	// then
	require.NoError(t, err)
	assert.NotNil(t, kp1)

	// given
	hostname := vgrand.RandomStr(5)

	// when
	err = w.UpdatePermissions(hostname, wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access: wallet.ReadAccess,
			AllowedKeys: []string{
				kp1.PublicKey(),
				kp2.PublicKey(),
			},
		},
	})

	// then
	require.NoError(t, err)

	// when
	err = w.TaintKey(kp1.PublicKey())

	// then
	require.NoError(t, err)

	// when
	pubKey1, err := w.DescribePublicKey(kp1.PublicKey())

	// then
	require.NoError(t, err)
	assert.NotNil(t, pubKey1)
	assert.True(t, pubKey1.IsTainted())

	// when
	permissions := w.Permissions(hostname)

	// then
	require.NotContains(t, permissions.PublicKeys.AllowedKeys, kp1.PublicKey())

	// when
	err = w.TaintKey(kp2.PublicKey())

	// then
	require.NoError(t, err)

	// when
	pubKey2, err := w.DescribePublicKey(kp2.PublicKey())

	// then
	require.NoError(t, err)
	assert.NotNil(t, pubKey2)
	assert.True(t, pubKey2.IsTainted())

	// when
	permissions = w.Permissions(hostname)

	// then
	require.Equal(t, permissions.PublicKeys, wallet.NoPublicKeysPermission())
}

func testHDWalletTaintingLastKeyPairVoidsThePermissions(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	assert.NotNil(t, w)

	// when
	kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

	// then
	require.NoError(t, err)
	assert.NotNil(t, kp1)

	// given
	hostname := vgrand.RandomStr(5)

	// when
	err = w.UpdatePermissions(hostname, wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access: wallet.ReadAccess,
			AllowedKeys: []string{
				kp1.PublicKey(),
			},
		},
	})

	// then
	require.NoError(t, err)

	// when
	err = w.TaintKey(kp1.PublicKey())

	// then
	require.NoError(t, err)

	// when
	pubKey1, err := w.DescribePublicKey(kp1.PublicKey())

	// then
	require.NoError(t, err)
	assert.NotNil(t, pubKey1)
	assert.True(t, pubKey1.IsTainted())

	// when
	permissions := w.Permissions(hostname)

	// then
	require.Equal(t, permissions.PublicKeys, wallet.NoPublicKeysPermission())
}

func testHDWalletTaintingKeyThatIsAlreadyTaintedFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			assert.ErrorIs(tt, err, wallet.ErrPubKeyAlreadyTainted)

			// when
			pubKey, err := w.DescribePublicKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, pubKey)
			assert.True(tt, pubKey.IsTainted())
		})
	}
}

func testHDWalletTaintingUnknownKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			err = w.TaintKey("vladimirharkonnen")

			// then
			assert.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
		})
	}
}

func testHDWalletUntaintingKeyPairSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			pubKey, err := w.DescribePublicKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, pubKey)
			assert.True(tt, pubKey.IsTainted())

			// when
			err = w.UntaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			pubKey, err = w.DescribePublicKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, pubKey)
			assert.False(tt, pubKey.IsTainted())
		})
	}
}

func testHDWalletUntaintingKeyThatIsNotTaintedFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.UntaintKey(kp.PublicKey())

			// then
			assert.ErrorIs(tt, err, wallet.ErrPubKeyNotTainted)

			// when
			pubKey, err := w.DescribePublicKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, pubKey)
			assert.False(tt, pubKey.IsTainted())
		})
	}
}

func testHDWalletUntaintingUnknownKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			err = w.UntaintKey("vladimirharkonnen")

			// then
			assert.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
		})
	}
}

func testHDWalletUpdatingKeyPairMetaSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			meta := []wallet.Metadata{{Key: "primary", Value: "yes"}}

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			updatedMeta, err := w.AnnotateKey(kp.PublicKey(), meta)

			// then
			require.NoError(tt, err)
			assert.Equal(tt, []wallet.Metadata{{Key: "primary", Value: "yes"}, {Key: "name", Value: "Key 1"}}, updatedMeta)
		})
	}
}

func testHDWalletUpdatingKeyPairMetaWithUnknownPublicKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			meta := []wallet.Metadata{{Key: "primary", Value: "yes"}}

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			_, err = w.AnnotateKey("somekey", meta)

			// then
			require.Error(tt, err, wallets.ErrWalletDoesNotExists)
		})
	}
}

func testHDWalletDescribingPublicKeysSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			pubKey, err := w.DescribePublicKey(kp1.PublicKey())

			// then
			require.NoError(tt, err)
			assert.Equal(tt, kp1.PublicKey(), pubKey.Key())
			assert.Equal(tt, kp1.Metadata(), pubKey.Metadata())
			assert.Equal(tt, kp1.IsTainted(), pubKey.IsTainted())
			assert.Equal(tt, kp1.AlgorithmName(), pubKey.AlgorithmName())
			assert.Equal(tt, kp1.AlgorithmVersion(), pubKey.AlgorithmVersion())
		})
	}
}

func testHDWalletDescribingUnknownPublicKeysFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			pubKey, err := w.DescribePublicKey("vladimirharkonnen")

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
			assert.Empty(tt, pubKey)
		})
	}
}

func testHDWalletListingPublicKeysSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			kp2, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp2)

			// when
			keys := w.ListPublicKeys()

			// then
			assert.Len(tt, keys, 2)
			assert.Equal(tt, keys[0].Key(), kp1.PublicKey())
			assert.Equal(tt, keys[1].Key(), kp2.PublicKey())
		})
	}
}

func testHDWalletListingKeyPairsSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			kp2, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp2)

			// when
			keys := w.ListKeyPairs()

			// then
			assert.Equal(tt, keys, []wallet.KeyPair{kp1, kp2})
		})
	}
}

func testHDWalletSigningTxSucceeds(t *testing.T) {
	tcs := []struct {
		name      string
		version   uint32
		signature string
	}{
		{
			name:      "version 1",
			version:   1,
			signature: "3849965c2f327f0b148e3b122cdc89a17fa07611e2a4178b1605dea5442ab7cfadb35d0b0ef527522f6477a5633b8f22d3b2d1e619d306111b7851a9d6100d02",
		}, {
			name:      "version 2",
			version:   2,
			signature: "4ad1fcd911f18d0df24de692376e5beac2700322e2ab5083bcf59fd17e0a21ffd64c88e4ba79162a7d46abd9ed0a81817c1648c8d7e93ed1b1d13499b12adb08",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			signature, err := w.SignTx(kp.PublicKey(), data)

			// then
			require.NoError(tt, err)
			assert.Equal(tt, kp.AlgorithmVersion(), signature.Version)
			assert.Equal(tt, kp.AlgorithmName(), signature.Algo)
			assert.Equal(tt, tc.signature, signature.Value)
		})
	}
}

func testHDWalletSigningTxWithTaintedKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			signature, err := w.SignTx(kp.PublicKey(), data)

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyIsTainted)
			assert.Nil(tt, signature)
		})
	}
}

func testHDWalletSigningTxWithUnknownKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			signature, err := w.SignTx("vladimirharkonnen", data)

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
			assert.Empty(tt, signature)
		})
	}
}

func testHDWalletSigningAnyMessageSucceeds(t *testing.T) {
	tcs := []struct {
		name      string
		version   uint32
		signature []byte
	}{
		{
			name:      "version 1",
			version:   1,
			signature: []byte{0x38, 0x49, 0x96, 0x5c, 0x2f, 0x32, 0x7f, 0xb, 0x14, 0x8e, 0x3b, 0x12, 0x2c, 0xdc, 0x89, 0xa1, 0x7f, 0xa0, 0x76, 0x11, 0xe2, 0xa4, 0x17, 0x8b, 0x16, 0x5, 0xde, 0xa5, 0x44, 0x2a, 0xb7, 0xcf, 0xad, 0xb3, 0x5d, 0xb, 0xe, 0xf5, 0x27, 0x52, 0x2f, 0x64, 0x77, 0xa5, 0x63, 0x3b, 0x8f, 0x22, 0xd3, 0xb2, 0xd1, 0xe6, 0x19, 0xd3, 0x6, 0x11, 0x1b, 0x78, 0x51, 0xa9, 0xd6, 0x10, 0xd, 0x2},
		}, {
			name:      "version 2",
			version:   2,
			signature: []byte{0x4a, 0xd1, 0xfc, 0xd9, 0x11, 0xf1, 0x8d, 0xd, 0xf2, 0x4d, 0xe6, 0x92, 0x37, 0x6e, 0x5b, 0xea, 0xc2, 0x70, 0x3, 0x22, 0xe2, 0xab, 0x50, 0x83, 0xbc, 0xf5, 0x9f, 0xd1, 0x7e, 0xa, 0x21, 0xff, 0xd6, 0x4c, 0x88, 0xe4, 0xba, 0x79, 0x16, 0x2a, 0x7d, 0x46, 0xab, 0xd9, 0xed, 0xa, 0x81, 0x81, 0x7c, 0x16, 0x48, 0xc8, 0xd7, 0xe9, 0x3e, 0xd1, 0xb1, 0xd1, 0x34, 0x99, 0xb1, 0x2a, 0xdb, 0x8},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			signature, err := w.SignAny(kp.PublicKey(), data)

			// then
			require.NoError(tt, err)
			assert.Equal(tt, tc.signature, signature)
		})
	}
}

func testHDWalletSigningAnyMessageWithTaintedKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			err = w.TaintKey(kp.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			signature, err := w.SignAny(kp.PublicKey(), data)

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyIsTainted)
			assert.Empty(tt, signature)
		})
	}
}

func testHDWalletSigningAnyMessageWithUnknownKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			signature, err := w.SignAny("vladimirharkonnen", data)

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
			assert.Empty(tt, signature)
		})
	}
}

func testHDWalletVerifyingAnyMessageSucceeds(t *testing.T) {
	tcs := []struct {
		name      string
		version   uint32
		signature []byte
	}{
		{
			name:      "version 1",
			version:   1,
			signature: []byte{0x38, 0x49, 0x96, 0x5c, 0x2f, 0x32, 0x7f, 0xb, 0x14, 0x8e, 0x3b, 0x12, 0x2c, 0xdc, 0x89, 0xa1, 0x7f, 0xa0, 0x76, 0x11, 0xe2, 0xa4, 0x17, 0x8b, 0x16, 0x5, 0xde, 0xa5, 0x44, 0x2a, 0xb7, 0xcf, 0xad, 0xb3, 0x5d, 0xb, 0xe, 0xf5, 0x27, 0x52, 0x2f, 0x64, 0x77, 0xa5, 0x63, 0x3b, 0x8f, 0x22, 0xd3, 0xb2, 0xd1, 0xe6, 0x19, 0xd3, 0x6, 0x11, 0x1b, 0x78, 0x51, 0xa9, 0xd6, 0x10, 0xd, 0x2},
		}, {
			name:      "version 2",
			version:   2,
			signature: []byte{0x4a, 0xd1, 0xfc, 0xd9, 0x11, 0xf1, 0x8d, 0xd, 0xf2, 0x4d, 0xe6, 0x92, 0x37, 0x6e, 0x5b, 0xea, 0xc2, 0x70, 0x3, 0x22, 0xe2, 0xab, 0x50, 0x83, 0xbc, 0xf5, 0x9f, 0xd1, 0x7e, 0xa, 0x21, 0xff, 0xd6, 0x4c, 0x88, 0xe4, 0xba, 0x79, 0x16, 0x2a, 0x7d, 0x46, 0xab, 0xd9, 0xed, 0xa, 0x81, 0x81, 0x7c, 0x16, 0x48, 0xc8, 0xd7, 0xe9, 0x3e, 0xd1, 0xb1, 0xd1, 0x34, 0x99, 0xb1, 0x2a, 0xdb, 0x8},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			verified, err := w.VerifyAny(kp.PublicKey(), data, tc.signature)

			// then
			require.NoError(tt, err)
			assert.True(tt, verified)
		})
	}
}

func testHDWalletVerifyingAnyMessageWithUnknownKeyFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)
			data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit.")
			sig := []byte{0xd5, 0xc4, 0x9e, 0xfd, 0x13, 0x73, 0x9b, 0xdd, 0x36, 0x81, 0x75, 0xcc, 0x59, 0xc8, 0xbe, 0xe1, 0x20, 0x25, 0xe4, 0xb9, 0x14, 0x7a, 0x22, 0xbb, 0xa4, 0x84, 0xef, 0x7e, 0xe7, 0x2f, 0x55, 0x13, 0x5f, 0x52, 0x55, 0xad, 0x90, 0x35, 0x67, 0x6c, 0x91, 0x9d, 0xbb, 0x91, 0x21, 0x1f, 0x98, 0x53, 0xcc, 0x68, 0xe, 0x58, 0x5b, 0x4c, 0x26, 0xd7, 0xea, 0x20, 0x1, 0x50, 0x6c, 0x41, 0xcb, 0x3}

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			signature, err := w.VerifyAny("vladimirharkonnen", data, sig)

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
			assert.Empty(tt, signature)
		})
	}
}

func testHDWalletMarshalingWalletSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
		result  string
	}{
		{
			name:    "version 1",
			version: 1,
			result:  `{"version":1,"node":"PjI6zxEu4dtcTu92dYlB/2Da+rvSpg7KzvmLMQ9wv6i6n75/ftik1rPYiZ/nTfBzqVttvNnoswyldTjPCjV5kw==","id":"9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e","keys":[{"index":1,"public_key":"30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","private_key":"1bbd4efb460d0bf457251e866697d5d2e9b58c5dcb96a964cd9cfff1a712a2b930ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","meta":[{"key":"name","value":"Key 1"}],"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}],"permissions":{}}`,
		}, {
			name:    "version 2",
			version: 2,
			result:  `{"version":2,"node":"PjI6zxEu4dtcTu92dYlB/2Da+rvSpg7KzvmLMQ9wv6i6n75/ftik1rPYiZ/nTfBzqVttvNnoswyldTjPCjV5kw==","id":"9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e","keys":[{"index":1,"public_key":"b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","private_key":"0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","meta":[{"key":"name","value":"Key 1"}],"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}],"permissions":{}}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := "test"

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			m, err := json.Marshal(&w)

			// then
			assert.NoError(tt, err)
			assert.Equal(tt, tc.result, string(m))
		})
	}
}

func testHDWalletMarshalingIsolatedWalletSucceeds(t *testing.T) {
	tcs := []struct {
		name      string
		version   uint32
		marshaled string
	}{
		{
			name:      "version 1",
			version:   1,
			marshaled: `{"version":1,"id":"9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e","keys":[{"index":1,"public_key":"30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","private_key":"1bbd4efb460d0bf457251e866697d5d2e9b58c5dcb96a964cd9cfff1a712a2b930ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","meta":[{"key":"name","value":"Key 1"}],"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}],"permissions":{}}`,
		}, {
			name:      "version 2",
			version:   2,
			marshaled: `{"version":2,"id":"9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e","keys":[{"index":1,"public_key":"b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","private_key":"0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","meta":[{"key":"name","value":"Key 1"}],"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}],"permissions":{}}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := "test"

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp)

			// when
			isolatedWallet, err := w.IsolateWithKey(kp.PublicKey())

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, isolatedWallet)

			// when
			m, err := json.Marshal(&isolatedWallet)

			// then
			assert.NoError(tt, err)
			assert.Equal(tt, tc.marshaled, string(m))
		})
	}
}

func testHDWalletUnmarshalingWalletSucceeds(t *testing.T) {
	tcs := []struct {
		name                string
		marshaled           string
		expectedVersion     uint32
		expectedPublicKey   string
		expectedPrivateKey  string
		expectedPermissions wallet.Permissions
	}{
		{
			name:                "version 1",
			expectedVersion:     1,
			marshaled:           `{"version":1,"node":"CZ13XhuFZ8K7TxNTAdKmMXh+OIVX6TFxTToXgnAqGlcO5eTY/5AVqZkWRIU3zfr8hvE7i2yIYAB6HT28ibi1fg==","keys":[{"index":1,"public_key":"30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","private_key":"1bbd4efb460d0bf457251e866697d5d2e9b58c5dcb96a964cd9cfff1a712a2b930ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371","meta":null,"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}]}`,
			expectedPublicKey:   "30ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371",
			expectedPrivateKey:  "1bbd4efb460d0bf457251e866697d5d2e9b58c5dcb96a964cd9cfff1a712a2b930ebce58d94ad37c4ff6a9014c955c20e12468da956163228cc7ec9b98d3a371",
			expectedPermissions: wallet.DefaultPermissions(),
		}, {
			name:                "version 2 without permissions",
			expectedVersion:     2,
			marshaled:           `{"version":2,"node":"CZ13XhuFZ8K7TxNTAdKmMXh+OIVX6TFxTToXgnAqGlcO5eTY/5AVqZkWRIU3zfr8hvE7i2yIYAB6HT28ibi1fg==","keys":[{"index":1,"public_key":"b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","private_key":"0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","meta":null,"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}]}`,
			expectedPublicKey:   "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			expectedPrivateKey:  "0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			expectedPermissions: wallet.DefaultPermissions(),
		}, {
			name:               "version 2 with permissions",
			expectedVersion:    2,
			marshaled:          `{"version":2,"node":"CZ13XhuFZ8K7TxNTAdKmMXh+OIVX6TFxTToXgnAqGlcO5eTY/5AVqZkWRIU3zfr8hvE7i2yIYAB6HT28ibi1fg==","keys":[{"index":1,"public_key":"b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","private_key":"0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0","meta":null,"tainted":false,"algorithm":{"name":"vega/ed25519","version":1}}],"permissions":{"vega.xyz":{"publicKeys":{"access":"read","allowedKeys":["b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0"]}}}}`,
			expectedPublicKey:  "b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			expectedPrivateKey: "0bfdfb4a04e22d7252a4f24eb9d0f35a82efdc244cb0876d919361e61f6f56a2b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
			expectedPermissions: wallet.Permissions{
				PublicKeys: wallet.PublicKeysPermission{
					Access: "read",
					AllowedKeys: []string{
						"b5fd9d3c4ad553cb3196303b6e6df7f484cf7f5331a572a45031239fd71ad8a0",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			w := wallet.HDWallet{}

			// when
			err := json.Unmarshal([]byte(tc.marshaled), &w)

			// then
			require.NoError(tt, err)
			assert.Equal(tt, tc.expectedVersion, w.KeyDerivationVersion())
			keyPairs := w.ListKeyPairs()
			assert.Len(tt, keyPairs, 1)
			assert.Equal(tt, tc.expectedPublicKey, keyPairs[0].PublicKey())
			assert.Equal(tt, tc.expectedPrivateKey, keyPairs[0].PrivateKey())
			assert.Equal(tt, uint32(1), keyPairs[0].AlgorithmVersion())
			assert.Equal(tt, "vega/ed25519", keyPairs[0].AlgorithmName())
			assert.False(tt, keyPairs[0].IsTainted())
			assert.Nil(tt, keyPairs[0].Metadata())
			assert.NotEmpty(tt, w.ID())
			assert.Equal(tt, tc.expectedPermissions, w.Permissions("vega.xyz"))
		})
	}
}

func testHDWalletGettingWalletInfoSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			require.NotNil(tt, w)
			assert.Equal(tt, "9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e", w.ID())
			assert.Equal(tt, "HD wallet", w.Type())
			assert.Equal(tt, tc.version, w.KeyDerivationVersion())
		})
	}
}

func testHDWalletGettingIsolatedWalletInfoSucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			isolatedWallet, err := w.IsolateWithKey(kp1.PublicKey())

			// then
			require.NoError(tt, err)
			require.NotNil(tt, isolatedWallet)
			assert.Equal(tt, "9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e", w.ID())
			assert.Equal(tt, "HD wallet", w.Type())
			assert.Equal(tt, tc.version, w.KeyDerivationVersion())
		})
	}
}

func testHDWalletIsolatingWalletSucceeds(t *testing.T) {
	walletName := vgrand.RandomStr(5)
	tcs := []struct {
		name    string
		version uint32
		wallet  string
	}{
		{
			name:    "version 1",
			version: 1,
			wallet:  fmt.Sprintf("%s.30ebce58.isolated", walletName),
		}, {
			name:    "version 2",
			version: 2,
			wallet:  fmt.Sprintf("%s.b5fd9d3c.isolated", walletName),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// when
			w, err := wallet.ImportHDWallet(walletName, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			isolatedWallet, err := w.IsolateWithKey(kp1.PublicKey())

			// then
			require.NoError(tt, err)
			require.NotNil(tt, isolatedWallet)
			assert.Equal(tt, tc.wallet, isolatedWallet.Name())
		})
	}
}

func testHDWalletIsolatingWalletWithTaintedKeyPairFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			kp1, err := w.GenerateKeyPair([]wallet.Metadata{})

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, kp1)

			// when
			err = w.TaintKey(kp1.PublicKey())

			// then
			require.NoError(tt, err)

			// when
			isolatedWallet, err := w.IsolateWithKey(kp1.PublicKey())

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyIsTainted)
			require.Nil(tt, isolatedWallet)
		})
	}
}

func testHDWalletIsolatingWalletWithNonExistingKeyPairFails(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)

			// then
			require.NoError(tt, err)
			assert.NotNil(tt, w)

			// when
			isolatedWallet, err := w.IsolateWithKey("0xdeadbeef")

			// then
			require.ErrorIs(tt, err, wallet.ErrPubKeyDoesNotExist)
			require.Nil(tt, isolatedWallet)
		})
	}
}

func testHDWalletGettingWalletMasterKeySucceeds(t *testing.T) {
	tcs := []struct {
		name    string
		version uint32
	}{
		{
			name:    "version 1",
			version: 1,
		}, {
			name:    "version 2",
			version: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			name := vgrand.RandomStr(5)

			// when
			w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, tc.version)
			require.NoError(tt, err)
			require.NotNil(tt, w)

			masterKeyPair, err := w.MasterKey()

			// then
			require.NoError(tt, err)
			assert.Equal(tt, "9df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e", masterKeyPair.PublicKey())
			assert.Equal(tt, "3e323acf112ee1db5c4eef76758941ff60dafabbd2a60ecacef98b310f70bfa89df682a3c87d90567f260566a9c223ccbbb7529c38340cf163b8fe199dbf0f2e", masterKeyPair.PrivateKey())
			assert.Equal(tt, "HD wallet", w.Type())
		})
	}
}

func testHDWalletUpdatingPermissionsWithInconsistentSetupFails(t *testing.T) {
	// when
	w, err := wallet.ImportHDWallet(vgrand.RandomStr(5), TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	kp1, err := w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	kp2, err := w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	taintedKey, err := w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	err = w.TaintKey(taintedKey.PublicKey())

	// then
	require.NoError(t, err)

	// given
	randomKey := vgrand.RandomStr(5)

	tcs := []struct {
		name          string
		permissions   wallet.Permissions
		expectedError error
	}{
		{
			name: "with non-existing keys on the wallet",
			permissions: wallet.Permissions{
				PublicKeys: wallet.PublicKeysPermission{
					Access:      wallet.ReadAccess,
					AllowedKeys: []string{kp1.PublicKey(), kp2.PublicKey(), randomKey},
				},
			},
			expectedError: fmt.Errorf("this allowed key %s does not exist on wallet", randomKey),
		}, {
			name: "with allowed keys and no access",
			permissions: wallet.Permissions{
				PublicKeys: wallet.PublicKeysPermission{
					Access:      wallet.NoAccess,
					AllowedKeys: []string{kp1.PublicKey()},
				},
			},
			expectedError: wallet.ErrCannotSetAllowedKeysWithNoAccess,
		}, {
			name: "with tainted key",
			permissions: wallet.Permissions{
				PublicKeys: wallet.PublicKeysPermission{
					Access:      wallet.ReadAccess,
					AllowedKeys: []string{taintedKey.PublicKey()},
				},
			},
			expectedError: fmt.Errorf("this allowed key %s is tainted", taintedKey.PublicKey()),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// when
			err = w.UpdatePermissions("vega.xyz", tc.permissions)

			// then
			require.EqualError(tt, err, fmt.Sprintf("inconsistent permissions setup: %v", tc.expectedError))
		})
	}
}

func testHDWalletUpdatingPermissionsWithEmptyPermissionFallbackToDefaults(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	_, err = w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	err = w.UpdatePermissions("vega.xyz", wallet.Permissions{})

	// then
	require.NoError(t, err)

	// when
	permissions := w.Permissions("vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), permissions)
}

func testHDWalletUpdatingPermissionsWithoutKeysInWalletFails(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	err = w.UpdatePermissions("vega.xyz", wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access: wallet.ReadAccess,
		},
	})

	// then
	require.ErrorIs(t, err, wallet.ErrWalletDoesNotHaveKeys)

	// when
	permissions := w.Permissions("vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), permissions)
}

func testHDWalletUpdatingPermissionsWithTaintedKeysOnlyFails(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	kp, err := w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	err = w.TaintKey(kp.PublicKey())

	// then
	require.NoError(t, err)

	// when
	err = w.UpdatePermissions("vega.xyz", wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access: wallet.ReadAccess,
		},
	})

	// then
	require.ErrorIs(t, err, wallet.ErrAllKeysInWalletAreTainted)

	// when
	permissions := w.Permissions("vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), permissions)
}

func testHDWalletRevokingPermissionsSucceeds(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	_, err = w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	err = w.UpdatePermissions("vega.xyz", wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access:      wallet.ReadAccess,
			AllowedKeys: []string{w.ListKeyPairs()[0].PublicKey()},
		},
	})

	// then
	require.NoError(t, err)

	// given
	tokenPermissions := wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access:      wallet.ReadAccess,
			AllowedKeys: []string{w.ListKeyPairs()[0].PublicKey()},
		},
	}

	// when
	err = w.UpdatePermissions("token.vega.xyz", tokenPermissions)

	// then
	require.NoError(t, err)

	// when
	w.RevokePermissions("vega.xyz")

	// when
	revokedPerms := w.Permissions("vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), revokedPerms)

	// when
	tokenPerms := w.Permissions("token.vega.xyz")

	// then
	assert.Equal(t, tokenPermissions, tokenPerms)
}

func testHDWalletPurgingPermissionsSucceeds(t *testing.T) {
	// given
	name := vgrand.RandomStr(5)

	// when
	w, err := wallet.ImportHDWallet(name, TestRecoveryPhrase1, 2)

	// then
	require.NoError(t, err)
	require.NotNil(t, w)

	// when
	_, err = w.GenerateKeyPair(nil)

	// then
	require.NoError(t, err)

	// when
	err = w.UpdatePermissions("vega.xyz", wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access:      wallet.ReadAccess,
			AllowedKeys: []string{w.ListKeyPairs()[0].PublicKey()},
		},
	})

	// then
	require.NoError(t, err)

	// when
	err = w.UpdatePermissions("token.vega.xyz", wallet.Permissions{
		PublicKeys: wallet.PublicKeysPermission{
			Access:      wallet.ReadAccess,
			AllowedKeys: []string{w.ListKeyPairs()[0].PublicKey()},
		},
	})

	// then
	require.NoError(t, err)

	// when
	w.PurgePermissions()

	// when
	perms := w.Permissions("vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), perms)

	// when
	perms = w.Permissions("token.vega.xyz")

	// then
	assert.Equal(t, wallet.DefaultPermissions(), perms)
}
