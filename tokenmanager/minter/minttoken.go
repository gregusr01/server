package minter

func (c *client) MintToken(ctx context.Context, tokenType, hostname string) (string, error) {

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
