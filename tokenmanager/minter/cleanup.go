package minter

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) CleanInvalidTokens(ctx context.Context) {
	for {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			if err := c.Database.DeleteInvalidTokens(); err != nil {
				logrus.Info("Error cleaning invalid token database", err)
			}
		}
	}
}
