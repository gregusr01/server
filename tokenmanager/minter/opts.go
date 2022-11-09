package minter

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/go-vela/server/database"
	"github.com/golang-jwt/jwt/v4"
)

// ClientOpt represents a configuration option to initialize the secret client for Native.
type ClientOpt func(*client) error

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithRegTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token duration for registration token")

		// set the token duration in the vault client
		c.config.RegTokenDuration = tokenDuration

		return nil
	}
}

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithAuthTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token duration for auth token")

		// set the token duration in the vault client
		c.config.AuthTokenDuration = tokenDuration

		return nil
	}
}

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithPrivKey(privKey *rsa.PrivateKey) ClientOpt {
	return func(c *client) error {

		c.Logger.Trace("configuring private key for token signing")

		if privKey == nil {
			c.Logger.Trace("no private key found")
			return errors.New("this failed")
		}

		// set the token duration in the vault client
		c.config.PrivKey = privKey

		return nil
	}
}

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithPubKey(pubKey *rsa.PublicKey) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring public key for token validation")

		// set the token duration in the vault client
		c.config.PubKey = pubKey

		return nil
	}
}

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithSignMethod(signingMethod jwt.SigningMethod) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token signing method")

		// set the token duration in the vault client
		c.config.SignMethod = signingMethod

		return nil
	}
}

// WithDatabase sets the Vela database service in the secret client for Native.
func WithDatabase(d database.Service) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring database service in native secret client")

		// check if the Vela database service provided is empty
		if d == nil {
			return fmt.Errorf("no Vela database service provided")
		}

		// set the Vela database service in the secret client
		c.Database = d

		return nil
	}
}

// WithTokenCleanupInterval sets the token interval in the secret client for Vault.
func WithTokenCleanupInterval(tokenCleanupInterval time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring interval for token cleanup duration")

		// set the token duration in the vault client
		c.config.TokenCleanupInterval = tokenCleanupInterval

		return nil
	}
}

// WithTokenCleanupInterval sets the token interval in the secret client for Vault.
func WithTickerInterval(tickerInterval time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring interval for ticker")

		// set the ticker interval in config
		c.config.TickerInterval = tickerInterval

		return nil
	}
}

// WithPubKeyCache sets the cache in which public keys will be served
func WithPubKeyCache(pubKeycache map[string]*rsa.PublicKey) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring public key for token validation")

		// set the token duration in the vault client
		c.config.PubKeyCache = pubKeycache

		return nil
	}
}

// WithKid sets the key identifier that will be stamped in every token
func WithKid(pubKeycache string) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring public key for token validation")

		// set the token duration in the vault client
		c.config.Kid = pubKeycache

		return nil
	}
}
