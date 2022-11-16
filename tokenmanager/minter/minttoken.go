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

	c.Logger.Tracef("Minting %s token for %s", tokenType, hostname)

	//set token claims
	switch tokenType {
	case "Registration":
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

	// set kid header
	tk.Header["kid"] = c.config.Kid

	//sign token with configured private signing key
	token, err := tk.SignedString(c.config.PrivKey)
	if err != nil {
		return "", err //wrap error
	}
	return token, nil
}
