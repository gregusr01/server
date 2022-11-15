package tokenmanager

import "github.com/sirupsen/logrus"

// New creates and returns a Vela service capable of
// integrating with the token manager service.
//
// .
func New(s *Setup) (Service, error) {
	validate the setup being provided

	https://pkg.go.dev/github.com/go-vela/server/scm?tab=doc#Setup.Validate
	err := s.Validate()
	if err != nil {
		return nil, err
	}

	logrus.Debug("creating tokenManager service from setup")

	return s.Tokenmanager()
}

func (s *Setup) Validate() error {
	logrus.Trace("validating tokenManger setup")

	// verify a secret driver was provided
	if len(s.Driver) == 0 {
		return fmt.Errorf("no tokenManger driver provided")
	}

	if s.Database == nil {
		return fmt.Errorf("no database service provided for tokenManger")
	}

	if s.PrivKey == nil {
		return fmt.Errorf("no private key provided for tokenManger")
	}

	if s.PubKey == nil {
		return fmt.Errorf("no public key provided for tokenManger")
	}

	if s.PublicKeyCache == nil {
		return fmt.Errorf("no public key cache provided for tokenManger")
	}

	if len(s.Kid) == 0 {
		return fmt.Errorf("no kid provided for tokenMangager")
	}

	if s.SignMethod == nil {
		return fmt.Errorf("no public key cache provided for tokenManger")
	}

	if s.RegTokenDuration.String() == "0s" {
		return fmt.Errorf("no registration token duration provided for tokenManger")
	}

	if s.AuthTokenDuration.String() == "0s" {
		return fmt.Errorf("no authentication token duration provided for tokenManger")
	}

	if s.TokenTickerInterval.String() == "0s" {
		return fmt.Errorf("no token cleanup interval provided for tokenManger")
	}

	if s.TokenCleanupInterval.String() == "0s" {
		return fmt.Errorf("no invalid token TTL provided for tokenManger")
	}

	if s.TokenCleanupInterval < s.AuthTokenDuration {
		return fmt.Errorf("Token cleanup TTL for tokenManager too short, must be larger than configured AuthTokenDuration")
	}

	if s.TokenCleanupInterval < s.RegTokenDuration {

	}

	if s.KeyCleanupInterval.String() == "0s" {
		return fmt.Errorf("no signing key TTL provided for tokenManger")
	}

	if s.KeyTickerInterval.String() == "0s" {
		return fmt.Errorf("no key cleanup interval provided for tokenManger")
	}

	if s.TokenCleanupInterval.String() == "0s" {
		return fmt.Errorf("no invalid token TTL provided for tokenManger")
	}

	// setup is valid
	return nil
}
