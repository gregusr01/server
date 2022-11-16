// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//nolint:dupl // ignore similar code with update.go
package token

import (
	"database/sql"
	"time"
)

// InvalidateToken adds a token hash to the token_invalidate database.
func (e *engine) InvalidateToken(t string) error {
	e.logger.Tracef("Invalidating token")

	type token struct {
		TokenHash sql.NullString `sql:"token_hash"`
		Timestamp sql.NullInt64  `sql:"timestamp"`
	}

	tk := token{
		TokenHash: sql.NullString{String: t, Valid: true},
		Timestamp: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// send query to the database
	return e.client.
		Table("invalid_tokens").
		Create(&tk).
		Error
}
