package minter

import (
	"fmt"

	"github.com/go-vela/server/database"
)

// ClientOpt represents a configuration option to initialize the secret client for Native.
type ClientOpt func(*client) error

// WithTokenDuration sets the token duration in the secret client for Vault.
func WithRegTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token duration in vault secret client")

		// set the token duration in the vault client
		c.config.RegTokenDuration = tokenDuration

		return nil
	}
}


// WithTokenDuration sets the token duration in the secret client for Vault.
func WithAuthTokenDuration(tokenDuration time.Duration) ClientOpt {
	return func(c *client) error {
		c.Logger.Trace("configuring token duration in vault secret client")

		// set the token duration in the vault client
		c.config.AuthTokenDuration = tokenDuration

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
