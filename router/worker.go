// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/api"
	"github.com/go-vela/server/router/middleware"
	"github.com/go-vela/server/router/middleware/auth"
	"github.com/go-vela/server/router/middleware/perm"
	"github.com/go-vela/server/router/middleware/worker"
)

// WorkerHandlers is a function that extends the provided base router group
// with the API handlers for worker functionality.
//
// POST   /api/v1/users
// GET    /api/v1/workers
// GET    /api/v1/workers/:worker
// PUT    /api/v1/workers/:worker
// DELETE /api/v1/workers/:worker .
func WorkerHandlers(base *gin.RouterGroup) {
	// Workers endpoints
	workers := base.Group("/workers")
	{
		workers.POST("", perm.MustPlatformAdmin(), middleware.Payload(), api.CreateWorker)
		workers.POST("/register", perm.MustPlatformAdmin(), middleware.Payload(), auth.MustValidToken(), auth.MustRegistration(), api.RegisterWorker)
		workers.GET("", api.GetWorkers)

		// Worker endpoints
		w := workers.Group("/:worker")
		{
			w.GET("", worker.Establish(), api.GetWorker)
			w.PUT("", perm.MustPlatformAdmin(), worker.Establish(), api.UpdateWorker)
			w.PUT("/auth", perm.MustPlatformAdmin(), worker.Establish(), auth.MustValidToken(), auth.MustAuth(), api.UpdateWorker)

			w.DELETE("", perm.MustPlatformAdmin(), worker.Establish(), api.DeleteWorker)
			w.DELETE("/auth", perm.MustPlatformAdmin(), worker.Establish(), auth.MustValidToken(), auth.MustAuth(), api.DeleteWorker)

		} // end of worker endpoints
	} // end of workers endpoints
}
