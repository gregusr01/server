package tokenmanager

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	TokenType string
	Iat       time.Time
	Exp       time.Time
	Sub       string
}

// New creates and returns a Vela service capable of
// integrating with the configured scm provider.
//
// Currently the following scm providers are supported:
//
// * Github
// .
func New(s *Setup) (Service, error) {
	// validate the setup being provided
	//
	// https://pkg.go.dev/github.com/go-vela/server/scm?tab=doc#Setup.Validate
	//err := s.Validate()
	// if err != nil {
	// 	return nil, err
	// }

	logrus.Debug("creating tokenManager service from setup")


	return s.Tokenmanager()
}

func MintToken(ctx context.Context, tokenType, hostname string) (string, error) {

	//pull priv key from postgres
	var tk *jwt.Token
	switch tokenType {
	case "Registration":

		// check DB for registration token existence
		tk = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Unix(),
			"sub":       hostname,
		})
	case "Auth":
		tk = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Unix(),
			"sub":       hostname,
		})
	default:
		return "", errors.New("invalid token type")
	}
	token, err := tk.SignedString(signKey)
	if err != nil {
		return "", err //wrap error
	}
	return token, nil
}

// validateToken validates a token using the public key
func ValidateToken(ctx context.Context, token string) (*AuthClaims, error) {

	//pull pub key from postgres (or use priv key and derive pub)

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, errors.New("failed parsing: " + err.Error())
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	//validate exp time

	//validate not part of invalidationDB

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

func InvalidateToken(ctx context.Context, token string) error{

  //drop in invalidation db

}
