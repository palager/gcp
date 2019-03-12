// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Dirserver is a wrapper for a directory implementation that presents it as an
// HTTP interface that stores its data in a GCS implementation
// of the Store service.
package main // import "github.com/palager/gcp/cmd/dirserver-gcp"

import (
	"flag"

	cloudLog "github.com/palager/gcp/cloud/log"
	"github.com/palager/upspin/log"
	"github.com/palager/upspin/metric"
	"github.com/palager/upspin/serverutil/dirserver"

	"github.com/palager/gcp/cloud/gcpmetric"
	"github.com/palager/gcp/cloud/https"

	// TODO: Which of these are actually needed?

	// Load useful packers
	_ "github.com/palager/upspin/pack/ee"
	_ "github.com/palager/upspin/pack/eeintegrity"
	_ "github.com/palager/upspin/pack/plain"

	// Load required transports
	_ "github.com/palager/upspin/transports"
)

const (
	serverName    = "dirserver"
	samplingRatio = 1    // report all metrics
	maxQPS        = 1000 // unlimited metric reports per second
)

func main() {
	project := flag.String("project", "", "GCP `project` name")

	ready := dirserver.Main()

	if *project != "" {
		cloudLog.Connect(*project, serverName)
		svr, err := gcpmetric.NewSaver(*project, samplingRatio, maxQPS, "serverName", serverName)
		if err != nil {
			log.Fatalf("Can't start a metric saver for GCP project %q: %s", *project, err)
		} else {
			metric.RegisterSaver(svr)
		}
	}

	https.ListenAndServe(ready, serverName)
}
