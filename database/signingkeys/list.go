// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"database/sql"
)

//signingKey struct - this should be added to library later
type signingKey struct {
	Kid        sql.NullString `sql:"kid"`
	PublicKey  sql.NullString `sql:"public_key"`
	ServerName sql.NullString `sql:"server_name"`
	Timestamp  sql.NullInt64  `sql:"timestamp"`
}

// ListSigningKeys returns all valid signing keys from the database.
func (e *engine) ListSigningKeys() ([]signingKey, error) {
	e.logger.Tracef("retrieving list of signing keys from database")

	var sk []signingKey

	// send query to the database and store result in variable
	err := e.client.
		Table("signing_keys").
		Find(&sk).
		Error
	if err != nil {
		return nil, err
	}

	e.logger.Tracef("what we got back: %s", sk)

	return sk, nil
}
