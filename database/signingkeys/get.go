// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"database/sql"
	"errors"
)

// GetInvalidToken checks for an existing token from the database.
func (e *engine) GetSigningKey(k string) (*rsa.PublicKey, error) {
	e.logger.Tracef("attempting to retrieve key %s from database", k)

	//var tk string
	var sk string

	// send query to the database and store result in variable
	err := e.client.
		Table("signing_keys").
		Where("kid = ?", k).
		Take(&sk).
		Error
	if err != nil {
		return signingKey{}, err
	}

	e.logger.Tracef("what we got back: %s", sk)

	//decode public key
	unB64, err := base64.StdEncoding.DecodeString(v.PublicKey)

	if err != nil {
		logrus.Info("unable to decode public key", err)
	}

	//parse into pub key format
	pubKeys[v.Kid] := x509.ParsePKCS1PublicKey(unB64)

	// if we got something
	return sk, nil
}
