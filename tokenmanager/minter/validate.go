package minter

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type AuthClaims struct {
	TokenType string
	Iat       time.Time
	Exp       time.Time
	Sub       string
}

// ValidateToken validates a token using the public key
func (c *client) ValidateToken(ctx context.Context, token string) (*AuthClaims, error) {

	//pull kid from token header

	//pull pub key from c.config.PubKeyCache[KID]

	//if key not exist in PubKeyCache
		// ???get individual key from DB???

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
	hasher := sha1.New()
	hasher.Write([]byte(token))
	sth := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	logrus.Info("STH: ", sth)

	if err = c.Database.GetInvalidToken(sth); err != nil {
		retErr := fmt.Errorf("unable to call token invalidation db: %w", err)

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
