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
	"github.com/go-vela/server/router/middleware/user"
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
		workers.POST("", user.Establish(), perm.MustPlatformAdmin(), middleware.Payload(), api.CreateWorker)
		workers.GET("", user.Establish(), api.GetWorkers)

		// Worker endpoints
		w := workers.Group("/:worker")
		{
			w.GET("", user.Establish(), worker.Establish(), api.GetWorker)
			w.PUT("", user.Establish(), perm.MustPlatformAdmin(), worker.Establish(), api.UpdateWorker)
			w.DELETE("", user.Establish(), perm.MustPlatformAdmin(), worker.Establish(), api.DeleteWorker)
		} // end of worker endpoints

		wa := workers.Group("/auth")
		{
			wa.POST("", middleware.Payload(), auth.MustValidToken(), auth.MustRegistration(), api.RegisterWorker)
			wa.GET("", auth.MustValidToken(), auth.MustAuth(), api.GetWorkers)

			ww := wa.Group("/:worker")
			{
				ww.GET("", auth.MustValidToken(), auth.MustAuth(), worker.EstablishWithAuth(), api.GetWorker)
				ww.PUT("", auth.MustValidToken(), auth.MustAuth(), worker.EstablishWithAuth(), api.UpdateWorker)
				ww.DELETE("", auth.MustValidToken(), auth.MustAuth(), worker.EstablishWithAuth(), api.DeleteWorker)
			}
		}
	} // end of workers endpoints
}
