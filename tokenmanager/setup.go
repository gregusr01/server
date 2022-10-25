package secret

import (
  	"github.com/golang-jwt/jwt"
)

type Setup struct {
	// tokenmanager Configuration

	// specifies the database service to use for the client
	Database database.Service

  //Private Key Used to Sign Token
  PrivKey *rsa.PrivateKey

  //Public Key Used to Validate Token
  PubKey *rsa.PublicKey

  //Token SigningMethod
  SignMethod SigningMethod

	// specifies the token duration to use for the registration token
	RegTokenDuration time.Duration

  // specifies the token duration to use for the worker authentication token
  AuthTokenDuration time.Duration
}
