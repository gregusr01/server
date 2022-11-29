// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package auth

import (
	"context"

	"github.com/go-vela/types/library"
)

const key = "auth_token"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Worker associated with this context.
func FromContext(c context.Context) *library.AuthClaims {
	value := c.Value(key)
	if value == nil {
		return nil
	}
	cl, ok := value.(*library.AuthClaims)
	if !ok {
		return nil
	}
	return cl
}

// ToContext adds the Worker to this context if it supports
// the Setter interface.
func ToContext(c Setter, cl *library.AuthClaims) {
	c.Set(key, cl)
}
