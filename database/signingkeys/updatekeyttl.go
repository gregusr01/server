// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//nolint:dupl // ignore similar code with update.go
package signingkeys

import "time"

// InvalidateToken adds a token hash to the token_invalidate database.
func (e *engine) UpdateKeyTTL(kid string) error {
	e.logger.Tracef("updating ttl for known key")

	//any vaidation we can do on hash?

	// send query to the database
	return e.client.
		Table("signing_keys").
		Where("kid = ?", kid).
		Update("timestamp", time.Now().Unix()).
		Error
}
