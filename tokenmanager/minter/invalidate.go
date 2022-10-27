package minter

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/go-vela/server/database"
)

func (c *client) InvalidateToken(ctx context.Context, token string) error {

	//hash token
	// th := sha1.Sum([]byte(token))

	// sth := fmt.Sprintf("%v", th)

	hasher := sha1.New()
	hasher.Write([]byte(token))
	sth := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	//drop in invalidation db
	if err := database.FromContext(ctx).InvalidateToken(sth); err != nil {
		retErr := fmt.Errorf("unable to invalidate token: %w", err)

		return retErr
	}

	return nil
}
