package minter

import (
	"time"
)

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) CleanInvalidTokens() {
	c.Logger.Tracef("Starting CleanInvalidTokens function on loop.  Will run every %v purging stale invalid tokens older than %v", c.config.TokenCleanupTicker, c.config.InvalidTokenTTL)
	for {
		ticker := time.NewTicker(c.config.TokenCleanupTicker) //double check me
		for range ticker.C {
			c.Logger.Tracef("Cleaning Stale Invalid Tokens")
			if err := c.Database.DeleteInvalidTokens(c.config.InvalidTokenTTL); err != nil {
				c.Logger.Warning("Error cleaning invalid token database", err)
			}
		}
	}
}

// CleanInvalidTokens cleans old entries to the signing key db
func (c *client) CleanExpiredSigningKeys() {
	c.Logger.Tracef("Starting CleanExpiredSigningKeys function on loop.  Will run every %v purging stale keys older than %v", c.config.KeyCleanupTicker, c.config.SigningKeyTTL)
	for {
		ticker := time.NewTicker(c.config.KeyCleanupTicker) //double check me
		for range ticker.C {
			c.Logger.Tracef("Cleaning Stale Signing Keys")
			if err := c.Database.DeleteExpiredKeys(c.config.SigningKeyTTL); err != nil {
				c.Logger.Warning("Error cleaning signing key database", err)
			}
		}
	}
}
