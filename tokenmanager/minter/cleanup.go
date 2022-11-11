package minter

import (
	"time"

	"github.com/sirupsen/logrus"
)

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) CleanInvalidTokens() {
	logrus.Info("Minter:CleanInvalidTokens method called")
	for {
		ticker := time.NewTicker(c.config.TickerInterval) //double check me
		for range ticker.C {
			logrus.Info("Cleaning Invalid Tokens")
			if err := c.Database.DeleteInvalidTokens(c.config.TokenCleanupInterval); err != nil {
				logrus.Info("Error cleaning invalid token database", err)
			}
		}
	}
}

func (c *client) CleanExpiredSigningKeys() {
	logrus.Info("Minter:DeleteExpiredKeys method called")
	for {
		ticker := time.NewTicker(c.config.TickerInterval) //double check me
		for range ticker.C {
			logrus.Info("Cleaning Invalid Tokens")
			if err := c.Database.DeleteExpiredKeys(c.config.TokenCleanupInterval); err != nil {
				logrus.Info("Error cleaning invalid token database", err)
			}
		}
	}
}
