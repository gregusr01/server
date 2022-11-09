// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//nolint:dupl // ignore similar code with update.go
package signingkeys

import (
	"database/sql"
	"time"
)

// InvalidateToken adds a token hash to the token_invalidate database.
func (e *engine) AddSigningKey(k, pk, s string) error {
	e.logger.Tracef("Adding new public key to signing key database")

	sk := signingKey{
		Kid: sql.NullString{String: k, Valid: true},
		PublicKey: sql.NullString{String: pk, Valid: true},
		ServerName: sql.NullString{String: s, Valid: true},
		Timestamp: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// send query to the database
	return e.client.
		Table("signing_keys").
		Create(&sk).
		Error
}
