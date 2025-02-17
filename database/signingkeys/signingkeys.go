// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type (
	// config represents the settings required to create the engine that implements the SigningKeyService interface.
	config struct {
		// specifies to skip creating tables and indexes for the token engine
		SkipCreation bool
	}

	// engine represents the signing key functionality that implements the SigningKeyService interface.
	engine struct {
		// engine configuration settings used in signing key functions
		config *config

		// gorm.io/gorm database client used in signing key functions
		//
		// https://pkg.go.dev/gorm.io/gorm#DB
		client *gorm.DB

		// sirupsen/logrus logger used in signing key functions
		//
		// https://pkg.go.dev/github.com/sirupsen/logrus#Entry
		logger *logrus.Entry
	}
)

// New creates and returns a Vela service for integrating with signing keys in the database.
//
//nolint:revive // ignore returning unexported engine
func New(opts ...EngineOpt) (*engine, error) {
	// create new signing key engine
	e := new(engine)

	// create new fields
	e.client = new(gorm.DB)
	e.config = new(config)
	e.logger = new(logrus.Entry)

	// apply all provided configuration options
	for _, opt := range opts {
		err := opt(e)
		if err != nil {
			return nil, err
		}
	}

	// check if we should skip creating database objects
	if e.config.SkipCreation {
		e.logger.Warning("skipping creation of signing keys table and indexes in the database")

		return e, nil
	}

	// create the Signing Key table
	err := e.CreateSigningKeyTable(e.client.Config.Dialector.Name())
	if err != nil {
		return nil, fmt.Errorf("unable to create %s table: %w", "signing keys", err)
	}

	return e, nil
}
