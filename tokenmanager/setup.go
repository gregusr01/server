package tokenmanager

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/go-vela/server/database"
	"github.com/go-vela/server/tokenmanager/minter"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type Setup struct {
	// tokenmanager Configuration

	Driver string

	// specifies the database service to use for the client
	Database database.Service

	//Private Key Used to Sign Token
	PrivKey *rsa.PrivateKey

	//Public Key Used to Validate Token
	PubKey *rsa.PublicKey

	//Key ID for the signing key
	Kid string

	PublicKeyCache map[string]*rsa.PublicKey

	//Token SigningMethod
	SignMethod jwt.SigningMethod

	// specifies the token duration to use for the registration token
	RegTokenDuration time.Duration

	// specifies the token duration to use for the worker authentication token
	AuthTokenDuration time.Duration

	// specifies the interval for cleanup
	TokenCleanupInterval time.Duration

	// specifies the interval for the ticker
	TickerInterval time.Duration
}

// Tokenmanager creates and returns a Vela tokenmanager service
func (s *Setup) Tokenmanager() (Service, error) {
	logrus.Trace("creating tokenManger from setup")

	switch s.Driver {
	case "minter":
		return s.Minter()
	default:
		// handle an invalid tokenmanager driver being provided
		return nil, fmt.Errorf("invalid tokenmanger driver provided: %s", s.Driver)
	}
}

// Minter creates and returns a Vela service capable of
// integrating with the minter service.
func (s *Setup) Minter() (Service, error) {
	logrus.Trace("creating token manager client from setup")

	if s.PrivKey == nil {
		logrus.Trace("you failed")
	}

	// create new minter service
	//
	// https://pkg.go.dev/github.com/go-vela/server/scm/github?tab=doc#New
	return minter.New(
		minter.WithPrivKey(s.PrivKey),
		minter.WithPubKey(s.PubKey),
		minter.WithSignMethod(s.SignMethod),
		minter.WithRegTokenDuration(s.RegTokenDuration),
		minter.WithAuthTokenDuration(s.AuthTokenDuration),
		minter.WithDatabase(s.Database),
		minter.WithTokenCleanupInterval(s.TokenCleanupInterval),
		minter.WithTickerInterval(s.TickerInterval),
		minter.WithKid(s.Kid),
		minter.WithPubKeyCache(s.PublicKeyCache),
	)
}
