package tokenmanager

import (
  	"github.com/golang-jwt/jwt"
)

type Setup struct {
	// tokenmanager Configuration

  Driver string

	// specifies the database service to use for the client
	//Database database.Service

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


// Github creates and returns a Vela tokenmanager service
func (s *Setup) Tokenmanager() (Service, error) {
	logrus.Trace("creating tokenManger from setup")

  switch s.Driver {
  case "minter":
    // handle the Github scm driver being provided
    //
    // https://pkg.go.dev/github.com/go-vela/server/scm?tab=doc#Setup.Github
    return s.Minter()
  default:
    // handle an invalid scm driver being provided
    return nil, fmt.Errorf("invalid tokenmanger driver provided: %s", s.Driver)
  }
}


// Github creates and returns a Vela service capable of
// integrating with a Github scm system.
func (s *Setup) Minter() (Service, error) {
	logrus.Trace("creating token manager client from setup")

	// create new Github scm service
	//
	// https://pkg.go.dev/github.com/go-vela/server/scm/github?tab=doc#New
	return Minter.New(
    minter.WithPrivKey(s.PrivKey),
    minter.WithPubKey(s.PubKey),
    minter.WithSignMethod(s.SignMethod),
    minter.WithRegTokenDuration(s.RegTokenDuration),
    minter.WithAuthTokenDuration(s.AuthTokenDuration),
    minter.WithDatabase(s.Database),
	)
}
