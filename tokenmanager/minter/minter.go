package minter

import (
	"crypto/rsa"
	"time"

	"github.com/go-vela/server/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type config struct {
	PrivKey              *rsa.PrivateKey   //Private Key Used to Sign Token
	PubKey               *rsa.PublicKey    //Public Key Used to Validate Token
	SignMethod           jwt.SigningMethod //Token SigningMethod
	Kid                  string
	PublicKeyCache       map[string]*rsa.PublicKey
	RegTokenDuration     time.Duration
	AuthTokenDuration    time.Duration
	InvalidTokenTTL time.Duration
	TokenCleanupTicker  time.Duration
	SigningKeyTTL   time.Duration
	KeyCleanupTicker    time.Duration
}

type client struct {
	config   *config
	Database database.Service
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
