// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"crypto/rsa"
	"time"
)

// SigningKeyService represents the Vela interface for token manager
// functions with the supported Database backends.
//
//nolint:revive // ignore name stutter
type SigningKeyService interface {
	// Token Data Definition Language Functions
	//
	// https://en.wikipedia.org/wiki/Data_definition_language

	// CreateSigningKeyTable defines a function that creates the signing_keys table
	CreateSigningKeyTable(string) error

	// AddSigningKey defines a function that adds a signing key to the signing_keys table
	AddSigningKey(string, string, *rsa.PublicKey) error

	// GetSigningKey defines a function that returns a signing key given a key identifier
	GetSigningKey(string) (*rsa.PublicKey, error)

	// ListSigningKeys defines a function that returns all valid signing keys from the database
	ListSigningKeys() ([]signingKey, error)

	// DeleteExpiredKeys defines a function that deletes stale signing keys from the database
	DeleteExpiredKeys(time.Duration) error

	// UpdateKeyTTL defines a function that updates a key time to live in the signing_keys table
	UpdateKeyTTL(string) error
}
