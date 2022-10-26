// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import (
	"github.com/go-vela/types/library"
)

// RepoService represents the Vela interface for repo
// functions with the supported Database backends.
//
//nolint:revive // ignore name stutter
type TokenService interface {
	// Token Data Definition Language Functions
	//
	// https://en.wikipedia.org/wiki/Data_definition_language

	// CreateTokenTable defines a function that creates the invalid_tokens table.
	CreateInvalidTokenTable(string) error

	// InvalidateToken defines a function that adds a token hash to the invalid_tokens table
	InvalidateToken(string) error

	// GetInvalidToken defines a function that gets a token hash from the invalid_tokens table
	GetInvalidToken(string) error
}
