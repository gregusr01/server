// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import (
	"database/sql"
	"time"
)

// DeleteInvalidTokens deletes stale invalid tokens from the database.
func (e *engine) DeleteInvalidTokens(invalidTokenTTL time.Duration) error {

	// Initializing token clean up interval
	ts := time.Now().Add(-invalidTokenTTL).Unix()

	//token struct - this should be added to library later
	type token struct {
		TokenHash sql.NullString `sql:"token_hash"`
		Timestamp sql.NullInt64  `sql:"timestamp"`
	}

	var tk token

	// send query to the database
	return e.client.
		Table("invalid_tokens").
		Where("timestamp < ?", ts).
		Delete(&tk).
		Error
}
