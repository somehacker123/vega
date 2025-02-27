package v1

import (
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	vgcrypto "code.vegaprotocol.io/vega/libs/crypto"
	vgrand "code.vegaprotocol.io/vega/libs/rand"
	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
)

const (
	LengthForSessionHashSeed = 10

	jwtBearer = "Bearer "
)

var ErrSessionNotFound = errors.New("session not found")

type auth struct {
	log *zap.Logger
	// sessionID -> wallet name
	sessions    map[string]string
	privKey     *rsa.PrivateKey
	pubKey      *rsa.PublicKey
	tokenExpiry time.Duration

	mu sync.Mutex
}

func NewAuth(log *zap.Logger, cfgStore RSAStore, tokenExpiry time.Duration) (Auth, error) { //revive:disable:unexported-return
	keys, err := cfgStore.GetRsaKeys()
	if err != nil {
		return nil, err
	}
	priv, err := jwt.ParseRSAPrivateKeyFromPEM(keys.Priv)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse private RSA key: %w", err)
	}
	pub, err := jwt.ParseRSAPublicKeyFromPEM(keys.Pub)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse public RSA key: %w", err)
	}

	return &auth{
		sessions:    map[string]string{},
		privKey:     priv,
		pubKey:      pub,
		log:         log,
		tokenExpiry: tokenExpiry,
	}, nil
}

type Claims struct {
	jwt.StandardClaims
	Session string
	Wallet  string
}

func (a *auth) NewSession(walletName string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	expiresAt := time.Now().Add(a.tokenExpiry)

	session := genSession()

	claims := &Claims{
		Session: session,
		Wallet:  walletName,
		StandardClaims: jwt.StandardClaims{
			// these are seconds
			ExpiresAt: jwt.NewTime((float64)(expiresAt.Unix())),
			Issuer:    "vega wallet",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	ss, err := token.SignedString(a.privKey)
	if err != nil {
		a.log.Error("unable to sign token", zap.Error(err))
		return "", err
	}

	a.sessions[session] = walletName
	return ss, nil
}

// VerifyToken returns the wallet name associated for this session.
func (a *auth) VerifyToken(token string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	claims, err := a.parseToken(token)
	if err != nil {
		return "", err
	}

	walletName, ok := a.sessions[claims.Session]
	if !ok {
		return "", ErrSessionNotFound
	}

	return walletName, nil
}

func (a *auth) Revoke(token string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	claims, err := a.parseToken(token)
	if err != nil {
		return "", err
	}

	w, ok := a.sessions[claims.Session]
	if !ok {
		return "", ErrSessionNotFound
	}
	delete(a.sessions, claims.Session)
	return w, nil
}

func (a *auth) RevokeAllToken() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.sessions = map[string]string{}
}

func (a *auth) parseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return a.pubKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't parse JWT token: %w", err)
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, ErrInvalidClaims
}

func extractToken(r *http.Request) (string, error) {
	token := strings.TrimSpace(r.Header.Get("Authorization"))
	if !strings.HasPrefix(token, jwtBearer) {
		return "", ErrInvalidOrMissingToken
	}
	return strings.TrimSpace(token[len(jwtBearer):]), nil
}

func genSession() string {
	return hex.EncodeToString(vgcrypto.Hash(vgrand.RandomBytes(LengthForSessionHashSeed)))
}

func writeError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	buf, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
