package minter

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/go-vela/server/database"
)

func (c *client) InvalidateToken(ctx context.Context, token string) error {

	//hash token
	th := sha1.Sum([]byte(token))

	//drop in invalidation db
	if err := database.FromContext(ctx).InvalidateToken(string(th[:])); err != nil {
		retErr := fmt.Errorf("unable to invalidate token: %w", err)

		return retErr
	}

	return nil
}
