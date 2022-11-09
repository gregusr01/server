// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

// TokenService represents the Vela interface for token manager
// functions with the supported Database backends.
//
//nolint:revive // ignore name stutter
type SigningKeyService interface {
	// Token Data Definition Language Functions
	//
	// https://en.wikipedia.org/wiki/Data_definition_language

	// CreateTokenTable defines a function that creates the invalid_tokens table.
	CreateSigningKeyTable(string) error

	// InvalidateToken defines a function that adds a token hash to the invalid_tokens table
	AddSigningKey(string, string, *rsa.PublicKey) error

	GetSigningKey(string) error

	ListKeys() error

	//UpdateKeyTTL(string) error

	//DeleteExpiredKeys() error
}
