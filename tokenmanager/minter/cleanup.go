package minter

import (
	"time"

	"github.com/sirupsen/logrus"
)

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) CleanInvalidTokens() {
	logrus.Info("Minter:CleanInvalidTokens Function Called")
	for {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			logrus.Info("Cleaning Invalid Tokens")
			if err := c.Database.DeleteInvalidTokens(); err != nil {
				logrus.Info("Error cleaning invalid token database", err)
			}
		}
	}
}
