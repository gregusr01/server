// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package signingkeys

import (
	"time"

	"github.com/sirupsen/logrus"
)

// DeletePipeline deletes an existing pipeline from the database.
func (e *engine) DeleteExpiredKeys(tokenCleanupDuration time.Duration) error {

	logrus.Info("deleting expired keys")

	ts := time.Now().Add(-tokenCleanupDuration).Unix()

	//token struct - this should be added to library later

	//var tk string
	var sk signingKey

	// send query to the database
	return e.client.
		Table("signing_keys").
		Where("timestamp < ?", ts).
		Delete(&sk).
		Error
}
