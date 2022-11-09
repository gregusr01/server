package minter

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// MintToken checks the DB for existing registration token, creates and returns a new one
// if one does not exist
func (c *client) MintToken(ctx context.Context, tokenType, hostname string) (string, error) {

	//pull priv key from postgres
	var tk *jwt.Token

	switch tokenType {
	case "Registration":

		// check DB for registration token existence
		tk = jwt.NewWithClaims(c.config.SignMethod, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Add(c.config.RegTokenDuration).Unix(),
			"sub":       hostname,
		})
	case "Auth":
		tk = jwt.NewWithClaims(c.config.SignMethod, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Add(c.config.AuthTokenDuration).Unix(),
			"sub":       hostname,
		})
	default:
		return "", errors.New("invalid token type")
	}

	tk.Header["kid"] = c.config.Kid

	token, err := tk.SignedString(c.config.PrivKey)
	if err != nil {
		return "", err //wrap error
	}
	return token, nil
}
