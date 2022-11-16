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

	// specifies the driver to use for the token manager client
	Driver string

	// specifies the database service to use for the client
	Database database.Service

	//Private Key Used to Sign Token
	PrivKey *rsa.PrivateKey

	//Public Key Used to Validate Token
	PubKey *rsa.PublicKey

	//Key ID for the signing key
	Kid string

	//cache to hold all valid public signing keys of vela servers
	PublicKeyCache map[string]*rsa.PublicKey

	//Token SigningMethod
	SignMethod jwt.SigningMethod

	// specifies the token duration to use for the registration token
	RegTokenDuration time.Duration

	// specifies the token duration to use for the worker authentication token
	AuthTokenDuration time.Duration

	// specifies the time to live for the invalid tokens in the db
	InvalidTokenTTL time.Duration

	// specifies the interval for the token cleanup ticker
	TokenCleanupTicker time.Duration

	// specifies the time to live for the signing keys in the db
	SigningKeyTTL time.Duration

	// specifies the interval for the key cleanup ticker
	KeyCleanupTicker time.Duration
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
		minter.WithInvalidTokenTTL(s.InvalidTokenTTL),
		minter.WithTokenCleanupTicker(s.TokenCleanupTicker),
		minter.WithSigningKeyTTL(s.SigningKeyTTL),
		minter.WithKeyCleanupTicker(s.KeyCleanupTicker),
		minter.WithKid(s.Kid),
		minter.WithPubKeyCache(s.PublicKeyCache),
	)
}

func (s *Setup) Validate() error {
	logrus.Trace("validating tokenManger setup")

	// verify a driver was provided
	if len(s.Driver) == 0 {
		return fmt.Errorf("no tokenManger driver provided")
	}

	// verify a database was provided
	if s.Database == nil {
		return fmt.Errorf("no database service provided for tokenManger")
	}

	// verify signing private key was provided
	if s.PrivKey == nil {
		return fmt.Errorf("no private key provided for tokenManger")
	}

	// verify a public key was provided
	if s.PubKey == nil {
		return fmt.Errorf("no public key provided for tokenManger")
	}

	// verify a cache of public keys was provided
	if s.PublicKeyCache == nil {
		return fmt.Errorf("no public key cache provided for tokenManger")
	}

	// verify a signing key ID was provided
	if len(s.Kid) == 0 {
		return fmt.Errorf("no kid provided for tokenMangager")
	}

	// verify a signing method was provided
	if s.SignMethod == nil {
		return fmt.Errorf("no public key cache provided for tokenManger")
	}

	// verify a duration for the registation token was provided
	if s.RegTokenDuration.String() == "0s" {
		return fmt.Errorf("no registration token duration provided for tokenManger")
	}

	// verify a duration for the authentication token was provided
	if s.AuthTokenDuration.String() == "0s" {
		return fmt.Errorf("no authentication token duration provided for tokenManger")
	}

	// verify an interval to run the token cleanup job was provided
	if s.TokenCleanupTicker.String() == "0s" {
		return fmt.Errorf("no token cleanup interval provided for tokenManger")
	}

	// verify a time to live for the invalid tokens in the database was provided
	if s.InvalidTokenTTL.String() == "0s" {
		return fmt.Errorf("no invalid token TTL provided for tokenManger")
	}

	// verify the token cleanup interval is longer than the authentication token duration to avoid token replays
	if s.InvalidTokenTTL > s.AuthTokenDuration {
		return fmt.Errorf("Token cleanup TTL for tokenManager too short, must be larger than configured AuthTokenDuration")
	}

	// verify the token cleanup interval is longer than the registration token duration to avoid token replays
	if s.InvalidTokenTTL > s.RegTokenDuration {
		return fmt.Errorf("Token cleanup TTL for tokenManager too short, must be larger than configured RegTokenDuration")
	}

	// verify a time to live for the expired keys in the database was provided
	if s.SigningKeyTTL.String() == "0s" {
		return fmt.Errorf("no signing key TTL provided for tokenManger")
	}

	// verify an interval to run the key cleanup job was provided
	if s.KeyCleanupTicker.String() == "0s" {
		return fmt.Errorf("no key cleanup interval provided for tokenManger")
	}

	// setup is valid
	return nil
}
