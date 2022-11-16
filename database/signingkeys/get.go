// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

// GetSigningKey pulls a specific signing key from the database given a key identifier.
func (e *engine) GetSigningKey(k string) (*rsa.PublicKey, error) {
	e.logger.Tracef("attempting to retrieve key %s from database", k)

	var sk signingKey

	// send query to the database and store result in variable
	err := e.client.
		Table("signing_keys").
		Where("kid = ?", k).
		Take(&sk).
		Error
	if err != nil {
		return nil, err
	}

	e.logger.Tracef("what we got back: %s", sk)

	//decode public key
	unB64, err := base64.StdEncoding.DecodeString(sk.PublicKey.String)

	if err != nil {
		logrus.Info("unable to decode public key", err)
	}

	//parse into pub key format
	pubKey, err := x509.ParsePKCS1PublicKey(unB64)
	if err != nil {
		logrus.Info("unable to parse public key string", err)
	}

	return pubKey, nil
}
