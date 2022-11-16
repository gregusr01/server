package minter

import (
	"crypto/rsa"
	"time"

	"github.com/go-vela/server/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type config struct {
	//Private Key Used to Sign Token
	PrivKey              *rsa.PrivateKey
	//Public Key Used to Validate Token
	PubKey               *rsa.PublicKey
	//Token SigningMethod
	SignMethod           jwt.SigningMethod
	//Key identifier
	Kid                  string
	//Cache for valid public signing keys
	PublicKeyCache       map[string]*rsa.PublicKey
	//Validity duration for registration tokens
	RegTokenDuration     time.Duration
	//Validity duration for authentication tokens
	AuthTokenDuration    time.Duration
	//Time to live for invalid tokens in database
	InvalidTokenTTL 		 time.Duration
	//Interval to run token cleanup function
	TokenCleanupTicker   time.Duration
	//Time to live for signing keys in database
	SigningKeyTTL   		 time.Duration
	//Interval to run key cleanup function
	KeyCleanupTicker     time.Duration
}

type client struct {
	config   *config
	Database database.Service
	// https://pkg.go.dev/github.com/sirupsen/logrus#Entry
	Logger   *logrus.Entry
}

// New returns a tokenmanager implementation.
//
//nolint:revive // ignore returning unexported client
func New(opts ...ClientOpt) (*client, error) {

	// create new GitHub client
	c := new(client)

	// create new fields
	c.config = new(config)

	// create new fields
	c.Database = *new(database.Service)

	// create new logger for the client
	//
	// https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#StandardLogger
	logger := logrus.StandardLogger()

	// create new logger for the client
	//
	// https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#NewEntry
	c.Logger = logrus.NewEntry(logger).WithField("tokenmanager", c.Driver())

	// apply all provided configuration options

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
