// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import "github.com/go-vela/types/constants"

const (
	// CreatePostgresTable represents a query to create the Postgres repos table.
	CreatePostgresTable = `
CREATE TABLE
IF NOT EXISTS
invalid_tokens (
	UNIQUE(token_hash)    VARCHAR(250)
);`

	// CreateSqliteTable represents a query to create the Sqlite repos table.
	CreateSqliteTable = `
CREATE TABLE
IF NOT EXISTS
invalid_tokens (
	UNIQUE(token_hash)    TEXT
);`
)

// CreateRepoTable creates the token table in the database.
func (e *engine) CreateInvalidTokenTable(driver string) error {
	e.logger.Tracef("creating repos table in the database")

	// handle the driver provided to create the table
	switch driver {
	case constants.DriverPostgres:
		// create the token table for Postgres
		return e.client.Exec(CreatePostgresTable).Error
	default:
		// create the token table for Sqlite
		return e.client.Exec(CreateSqliteTable).Error
	}
}
