package minter

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/go-vela/server/database"
	"github.com/golang-jwt/jwt/v4"
)

// ClientOpt represents a configuration option to initialize the secret client for Native.
type ClientOpt func(*client) error

// WithRegTokenDuration sets the token duration for the registration token in the tokenmanager client
func WithRegTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring token duration of %v for registration token", tokenDuration)
		c.config.RegTokenDuration = tokenDuration

		return nil
	}
}

// WithAuthTokenDuration sets the token duration for the authentication in the tokenmanager client
func WithAuthTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring token duration of %v for auth token", tokenDuration)
		c.config.AuthTokenDuration = tokenDuration

		return nil
	}
}

// WithPrivKey sets the private key that will be used for token signing in the tokenmanager client
func WithPrivKey(privKey *rsa.PrivateKey) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring private key for token signing")
		c.config.PrivKey = privKey

		return nil
	}
}

// WithPubKey sets the public key that will be used for token validation in the tokenmanager client
func WithPubKey(pubKey *rsa.PublicKey) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring public key for token validation")
		c.config.PubKey = pubKey

		return nil
	}
}

// WithSignMethod sets the token signing method in the tokenmanager client
func WithSignMethod(signingMethod jwt.SigningMethod) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token signing method")
		c.config.SignMethod = signingMethod

		return nil
	}
}

// WithDatabase sets the Vela database service in the tokenmanager client
func WithDatabase(d database.Service) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring database service for token manager")

		// check if the Vela database service provided is empty
		if d == nil {
			return fmt.Errorf("no Vela database service provided")
		}

		// set the Vela database service in the tokenmanager client
		c.Database = d

		return nil
	}
}

// WithInvalidTokenTTL sets the token time to live in the tokenmanager client
func WithInvalidTokenTTL(invalidTokenTTL time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring interval of %v for token cleanup duration", invalidTokenTTL)
		c.config.InvalidTokenTTL = invalidTokenTTL

		return nil
	}
}

// WithInvalidTokenTTL sets the signing key time to live in the tokenmanager client
func WithSigningKeyTTL(signingKeyTTL time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring interval for token cleanup duration")
		c.config.SigningKeyTTL = signingKeyTTL

		return nil
	}
}

// WithTokenCleanupTicker sets the token cleanup interval in the tokenmanager client
func WithTokenCleanupTicker(tokenCleanupTicker time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring interval of %v for token cleanup ticker", tokenCleanupTicker)
		c.config.TokenCleanupTicker = tokenCleanupTicker

		return nil
	}
}

// WithKeyCleanupTicker sets the key cleanup interval in the tokenmanager client
func WithKeyCleanupTicker(keyCleanupTicker time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring interval of %v for key cleanup ticker", keyCleanupTicker)
		c.config.KeyCleanupTicker = keyCleanupTicker

		return nil
	}
}

// WithPubKeyCache sets the cache in which public keys will be served
func WithPubKeyCache(pubKeycache map[string]*rsa.PublicKey) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring public key cache that will be used for token validation")
		c.config.PublicKeyCache = pubKeycache

		return nil
	}
}

// WithKid sets the key identifier that will be stamped in every token
func WithKid(kid string) ClientOpt {
	return func(c *client) error {
		c.Logger.Tracef("configuring %v as the public key identifier (KID)", kid)
		c.config.Kid = kid

		return nil
	}
}
