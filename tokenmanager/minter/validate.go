package minter

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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
	//parse and validate given token
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		//pull kid from token header
		kid, err := getKid(token)
		if err != nil {
			return nil, err
		}

		//pull pub key from c.config.PubKeyCache[KID]
		if k, ok := c.config.PublicKeyCache[kid]; ok {
			return k, nil
		}

		logrus.Infof("KID %v not part of cache, checking database", kid)

		//check db for signing key
		k, err := c.Database.GetSigningKey(kid)
		if err != nil {
			return nil, err
		}

		return k, nil
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

	//check if hash present in invalidation db
	if err = c.Database.GetInvalidToken(sth); err != nil {
		return nil, fmt.Errorf("unable to call token invalidation db: %w", err)
	}

	return parseAuthClaims(tkn)
}

// parseAuthClaims parses jwtAuthN claims post signature validation.
func parseAuthClaims(token *jwt.Token) (*AuthClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not create token claims")
	}

	//set issued at time
	iatTime, ok := claims["iat"].(float64)
	if !ok {
		return nil, errors.New("iat claim is of invalid type")
	}

	//set expiration time
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

// Kid returns the key-identifier (kid) value of the key that signed the token
// or an error if the kid is not present.
func getKid(t string) (string, error) {
	//get token header
	parts := strings.Split(t, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token length")
	}

	//decode header string
	b, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	//unmarshal claims
	var claims map[string]string
	if err := json.Unmarshal(b, &claims); err != nil {
		return "", err
	}

	//retrieve kid
	kid, ok := claims["kid"]
	if !ok {
		return "", errors.New("missing kid")
	}
	return kid, nil
}
