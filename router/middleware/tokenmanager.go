// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/tokenmanager"
)

// Queue is a middleware function that initializes the queue and
// attaches to the context of every http.Request.
func Tokenmanager(t tokenmanager.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenmanager.WithGinContext(c, t)
		c.Next()
	}
}
