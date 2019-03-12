// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Storeserver-gcp is a wrapper for a store implementation that presents it as an
// HTTP interface to Google Cloud Storage.
package main // import "github.com/palager/gcp/cmd/storeserver-gcp"

import (
	"flag"

	cloudLog "github.com/palager/gcp/cloud/log"
	"github.com/palager/upspin/log"
	"github.com/palager/upspin/metric"
	"github.com/palager/upspin/serverutil/storeserver"

	"github.com/palager/gcp/cloud/gcpmetric"
	"github.com/palager/gcp/cloud/https"

	// Storage on GCS.
	_ "github.com/palager/gcp/cloud/storage/gcs"
)

const (
	serverName    = "storeserver"
	samplingRatio = 1    // report all metrics
	maxQPS        = 1000 // unlimited metric reports per second
)

func main() {
	project := flag.String("project", "", "GCP `project` name")

	ready := storeserver.Main()

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
