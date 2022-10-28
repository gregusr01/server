// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import (
	"database/sql"
	"errors"
)

// GetRepo gets a repo by ID from the database.
func (e *engine) GetInvalidToken(t string) error {
	e.logger.Tracef("getting token hash from the database")

	// variable to store query results
	type token struct {
		ts sql.NullString `sql:"token_hash"`
	}

	var tk token

	// send query to the database and store result in variable
	err := e.client.
		Table("invalid_tokens").
		Where("token_hash = ?", t).
		Take(&tk).
		Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil
		}
		return err
	}

	e.logger.Tracef("what we got back: %s", tk)

	// if we got something
	return errors.New("token hash found inside invalidation database")

}
