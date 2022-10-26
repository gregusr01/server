// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//nolint:dupl // ignore similar code with update.go
package token

import (
	"fmt"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/database"
	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

// InvalidateToken adds a token hash to the token_invalidate database.
func (e *engine) InvalidateToken(t string) error {
	e.logger.Tracef("Invalidating token")

	//any vaidation we can do on hash?

	// send query to the database
	return e.client.
		Table("invalid_tokens").
		Create(t).
		Error
}
