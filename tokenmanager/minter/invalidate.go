package minter

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/sirupsen/logrus"
)

// InvalidateToken uses the client context to create a database connection
// and adds a string hash of the token to the invalid_tokens table
func (c *client) InvalidateToken(ctx context.Context, token string) error {

	//hash token
	hasher := sha1.New()
	hasher.Write([]byte(token))
	sth := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	c.Logger.Tracef("Adding token hash %s to the invalidation db", sth)

	//add to invalidation db
	if err := c.Database.InvalidateToken(sth); err != nil {
		return fmt.Errorf("unable to invalidate token: %w", err)
	}

	return nil
}
