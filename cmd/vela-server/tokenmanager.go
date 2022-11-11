// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/go-vela/server/database"
	"github.com/go-vela/server/tokenmanager"
	"github.com/golang-jwt/jwt"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// helper function to setup the queue from the CLI arguments.
func setupTokenManger(c *cli.Context, d database.Service) (tokenmanager.Service, error) {

	logrus.Debug("Creating tokenManger for server worker authentication")

	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.Trace("error generating key pair")
	}
	pk := &k.PublicKey
	//generate Kid value
	kid := fmt.Sprintf("test%v", time.Now().Unix())
	//add public key to database
	if err := d.AddSigningKey(kid, "testserver", pk); err != nil {
		logrus.Trace("error adding signing key")
	}
	//build list of public keys (pull from database)
	pubKeyCache := map[string]*rsa.PublicKey{kid: pk}

	_manager := &tokenmanager.Setup{
		Driver:               "minter",
		Database:             d,
		PrivKey:              k,
		PubKey:               pk,
		PublicKeyCache:       pubKeyCache,
		Kid:                  kid,
		SignMethod:           jwt.SigningMethodRS256,
		RegTokenDuration:     time.Minute * 10,
		AuthTokenDuration:    time.Minute * 10,
		TokenTickerInterval:  time.Minute * 1,
		KeyCleanupInterval:   time.Minute * 5,
		KeyTickerInterval:    time.Minute * 1,
		TokenCleanupInterval: time.Minute * 5, //THIS TIME MUUUUUST BE EQUAL TO OR LONGER THAN THE DURATION OF THE TOKENS!!!!!!!!!!!!
	}

	return tokenmanager.New(_manager)
}
