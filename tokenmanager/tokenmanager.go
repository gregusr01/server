package tokenmanager

import "github.com/sirupsen/logrus"

// New creates and returns a Vela service capable of
// integrating with the token manager service.
//
// .
func New(s *Setup) (Service, error) {
	// validate the setup being provided
	//
	// https://pkg.go.dev/github.com/go-vela/server/scm?tab=doc#Setup.Validate
	err := s.Validate()
	if err != nil {
		return nil, err
	}

	logrus.Debug("creating tokenManager service from setup")

	return s.Tokenmanager()
}
