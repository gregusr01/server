// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-vela/server/database"
	"github.com/go-vela/server/tokenmanager"
	"github.com/golang-jwt/jwt"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
	"github.com/google/uuid"
)

// helper function to setup the queue from the CLI arguments.
func setupTokenManger(c *cli.Context, d database.Service) (tokenmanager.Service, error) {

	logrus.Debug("Creating tokenManger for server worker authentication")

	keyLoc := c.String("signing-key-path")
	var k *rsa.PrivateKey
	var err error

	switch keyLoc {
	case "":
		logrus.Trace("no key defined generating one internally")
		k, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			logrus.Trace("error generating key pair")
			return nil, errors.New("unable to generate a new key pair")
		}
		logrus.Tracef("successfully generated private key")
	default:

		logrus.Info("using private key from command line")
		//read key from file
		pemBytes, err := ioutil.ReadFile(keyLoc)
		if err != nil {
			logrus.Trace("error reading key from location")
			return nil, errors.New("unable to read file for key pair generation")
		}

		//decode pem
		block, _ := pem.Decode(pemBytes)
		if block == nil {
			logrus.Trace("error decoding key")
			return nil, errors.New("unable decode file into pem")
		}
		k, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			logrus.Trace("error parsing key")
			return nil, errors.New("unable to parse the key from the block")
		}
		logrus.Tracef("successfully loaded key from file")
	}

	pk := &k.PublicKey

	//generate Kid value
	kid := fmt.Sprintf("%v", uuid.New())

	//add public key to database
	if err := d.AddSigningKey(kid, c.String("server-addr"), pk); err != nil {
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
		TokenTickerInterval:  time.Minute * 5,
		KeyCleanupInterval:   time.Hour * 5,
		KeyTickerInterval:    time.Hour * 1,
		TokenCleanupInterval: time.Minute * 15,
	}

	return tokenmanager.New(_manager)
}
