package minter

func (c *client) MintToken(ctx context.Context, tokenType, hostname string) (string, error) {

	//pull priv key from postgres
	var tk *jwt.Token
	switch tokenType {
	case "Registration":

		// check DB for registration token existence
		tk = jwt.NewWithClaims(c.config.SignMethod, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Unix()+c.config.RegTokenDuration,
			"sub":       hostname,
		})
	case "Auth":
		tk = jwt.NewWithClaims(c.config.SignMethod, jwt.MapClaims{
			"tokenType": tokenType,
			"iat":       time.Now().Unix(),
			"exp":       time.Now().Unix()+c.config.AuthTokenDuration,
			"sub":       hostname,
		})
	default:
		return "", errors.New("invalid token type")
	}
	token, err := tk.SignedString(c.config.PrivKey)
	if err != nil {
		return "", err //wrap error
	}
	return token, nil
}
