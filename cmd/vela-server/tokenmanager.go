// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/server/tokenmanager"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// helper function to setup the queue from the CLI arguments.
func setupTokenManger(c *cli.Context, d database.Service) (queue.Service, error) {

	//BELOW IS QUEUE EXAMPLE FOR REFFERENCE...

	// logrus.Debug("Creating queue client from CLI configuration")
	//
	// // queue configuration
	// _setup := &queue.Setup{
	// 	Driver:  c.String("queue.driver"),
	// 	Address: c.String("queue.addr"),
	// 	Cluster: c.Bool("queue.cluster"),
	// 	Routes:  c.StringSlice("queue.routes"),
	// 	Timeout: c.Duration("queue.pop.timeout"),
	// }
	//
	// // setup the queue
	// //
	// // https://pkg.go.dev/github.com/go-vela/server/queue?tab=doc#New
	// return queue.New(_setup)




}
