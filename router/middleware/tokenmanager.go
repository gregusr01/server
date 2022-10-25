// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/tokenmanger"
)

// Queue is a middleware function that initializes the queue and
// attaches to the context of every http.Request.
func Tokenmanger(t tokenmanger.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenmanger.WithGinContext(c, t)
		c.Next()
	}
}
