// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package tokenmanager

import (
	"context"

	"github.com/gin-gonic/gin"
)

// key defines the key type for storing
// the tokenmanger Service in the context.
const key = "tokenmanager"

// FromContext retrieves the tokenmanager Service from the context.Context.
func FromContext(c context.Context) Service {
	// get tokenmanger value from context.Context
	v := c.Value(key)
	if v == nil {
		return nil
	}

	// cast tokenmanger value to expected Service type
	s, ok := v.(Service)
	if !ok {
		return nil
	}

	return s
}

// FromGinContext retrieves the tokenmanager Service from the gin.Context.
func FromGinContext(c *gin.Context) Service {
	// get tokenmanger value from gin.Context
	//
	// https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc#Context.Get
	v, ok := c.Get(key)
	if !ok {
		return nil
	}

	// cast tokenmanger value to expected Service type
	s, ok := v.(Service)
	if !ok {
		return nil
	}

	return s
}

// WithContext inserts the tokenmanager Service into the context.Context.
func WithContext(c context.Context, s Service) context.Context {
	// set the tokenmanger Service in the context.Context
	//
	// https://pkg.go.dev/context?tab=doc#WithValue
	//
	//nolint:staticcheck,revive // ignore using string with context value
	return context.WithValue(c, key, s)
}

// WithGinContext inserts the tokenmanager Service into the gin.Context.
func WithGinContext(c *gin.Context, s Service) {
	// set the tokenmanger Service in the gin.Context
	//
	// https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc#Context.Set
	c.Set(key, s)
}
