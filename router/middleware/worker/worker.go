// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/database"
	"github.com/go-vela/server/router/middleware/auth"
	"github.com/go-vela/server/util"
	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

// Retrieve gets the worker in the given context.
func Retrieve(c *gin.Context) *library.Worker {
	return FromContext(c)
}

// Establish sets the worker in the given context.
func Establish() gin.HandlerFunc {
	return func(c *gin.Context) {
		wParam := util.PathParameter(c, "worker")
		if len(wParam) == 0 {
			retErr := fmt.Errorf("no worker parameter provided")
			util.HandleError(c, http.StatusBadRequest, retErr)

			return
		}

		logrus.Debugf("Reading worker %s", wParam)

		w, err := database.FromContext(c).GetWorker(wParam)
		if err != nil {
			retErr := fmt.Errorf("unable to read worker %s: %w", wParam, err)
			util.HandleError(c, http.StatusNotFound, retErr)

			return
		}

		ToContext(c, w)
		c.Next()
	}
}

// Establish sets the worker in the given context using auth token.
func EstablishWithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		wParam := util.PathParameter(c, "worker")
		if len(wParam) == 0 {
			retErr := fmt.Errorf("no worker parameter provided")
			util.HandleError(c, http.StatusBadRequest, retErr)

			return
		}

		logrus.Debug("Comparing provided worker from token", wParam)

		claims := auth.FromContext(c)

		if claims.Sub != wParam {
			retErr := fmt.Errorf("Worker %s is not authorized to perform this request for %s", claims.Sub, wParam)
			util.HandleError(c, http.StatusUnauthorized, retErr)

			return
		}

		logrus.Debugf("Reading worker %s", wParam)

		w, err := database.FromContext(c).GetWorker(wParam)
		if err != nil {
			retErr := fmt.Errorf("unable to read worker %s: %w", wParam, err)
			util.HandleError(c, http.StatusNotFound, retErr)

			return
		}

		ToContext(c, w)
		c.Next()
	}
}
