// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package tokenmanager

import (
	"context"

	"github.com/go-vela/server/tokenmanager/minter"
)

// Service represents the interface for Vela integrating
// with the different supported Queue backends.
type Service interface {
	// Service Interface Functions
	Driver() string

	// MintToken defines a function that mints a token for a subject of a specific type
	MintToken(context.Context, string, string) (string, error)

	// ValidateToken defines a function that validates a token
	ValidateToken(context.Context, string) (*minter.AuthClaims, error)

	// InvalidateToken defines a function that invalidates a token by adding it to the invalidation table in the database
	InvalidateToken(context.Context, string) error
}
