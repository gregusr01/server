package minter

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
