// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import "github.com/go-vela/types/constants"

const (
	// CreatePostgresTable represents a query to create the Postgres signing_keys table.
	CreatePostgresTable = `
CREATE TABLE
IF NOT EXISTS
signing_keys (
	kid    VARCHAR(250),
	public_key    VARCHAR(500),
	server_name    VARCHAR(250),
	timestamp    BIGINT,
	UNIQUE(kid)
);`

	// CreateSqliteTable represents a query to create the Sqlite signing_keys table.
	CreateSqliteTable = `
CREATE TABLE
IF NOT EXISTS
signing_keys (
	kid           TEXT,
	public_key    TEXT,
	server_name    TEXT,
	timestamp    BIGINT,
	UNIQUE(kid)
);`
)

// CreateInvalidTokenTable creates the signing_keys table in the database.
func (e *engine) CreateSigningKeyTable(driver string) error {
	e.logger.Tracef("creating signing_keys table in the database")

	// handle the driver provided to create the table
	switch driver {
	case constants.DriverPostgres:
		// create the signing_keys table for Postgres
		return e.client.Exec(CreatePostgresTable).Error
	default:
		// create the signing_keys table for Sqlite
		return e.client.Exec(CreateSqliteTable).Error
	}
}
