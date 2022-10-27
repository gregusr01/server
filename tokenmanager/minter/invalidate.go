package minter

import (
	"context"
	"crypto/sha1"
	)

func (c *client) InvalidateToken(ctx context.Context, token string) error {

	//hash token
	th := sha1.Sum([]byte(token))

	//drop in invalidation db
	err = database.FromContext(c).InvalidateToken(th)
	if err != nil {
		retErr := fmt.Errorf("unable to invalidate token: %w", err)

		return retErr
	}

	return nil
}
