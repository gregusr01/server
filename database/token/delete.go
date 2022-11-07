// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import (
	"time"

	"github.com/sirupsen/logrus"
)

// DeletePipeline deletes an existing pipeline from the database.
func (e *engine) DeleteInvalidTokens() error {
	// e.logger.WithFields(logrus.Fields{
	// 	"pipeline": p.GetCommit(),
	// }).Tracef("deleting pipeline %s from the database", p.GetCommit())

	// cast the library type to database type
	//
	// https://pkg.go.dev/github.com/go-vela/types/database#PipelineFromLibrary
	//pipeline := database.PipelineFromLibrary(p)

	ts := time.Now().Add(-time.Minute * 1).Unix()

	logrus.Info("TS: ", ts)

	// send query to the database
	return e.client.
		Table("invalid_tokens").
		Delete("timestamp < ?", ts).
		Error
}
