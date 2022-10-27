package minter

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	TokenType string
	Iat       time.Time
	Exp       time.Time
	Sub       string
}

// validateToken validates a token using the public key
func (c *client) ValidateToken(ctx context.Context, token string) (*AuthClaims, error) {

	//pull pub key from postgres (or use priv key and derive pub)

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return c.config.PubKey, nil
	})
	if err != nil {
		return nil, errors.New("failed parsing: " + err.Error())
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	//validate not part of invalidationDB
	//hash token
	th := sha1.Sum([]byte(token))

	if err = database.FromContext(c).GetInvalidToken(th); err != nil {
		retErr := fmt.Errorf("unable to call token inalidation db: %w", err)

		return nil, retErr
	}

	return parseAuthClaims(tkn)
}

// parseAuthClaims parses jwtAuthN claims post signature validation.
func parseAuthClaims(token *jwt.Token) (*AuthClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not create token claims")
	}

	iatTime, ok := claims["iat"].(float64)
	if !ok {
		return nil, errors.New("iat claim is of invalid type")
	}

	expTime, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("exp claim is of invalid type")
	}

	return &AuthClaims{
		TokenType: claims["tokenType"].(string),
		Iat:       time.Unix(int64(iatTime), 0),
		Exp:       time.Unix(int64(expTime), 0),
		Sub:       claims["sub"].(string),
	}, nil
}
