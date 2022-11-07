// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

import (
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/database"
	//"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
  "time"
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

  ts := time.Now().Unix().Add(-time.Minute * 1)

	// send query to the database
	return e.client.
		Table("invalid_tokens").
		Delete("*").
    Where("timestamp < ?", ts).
		Error
}
