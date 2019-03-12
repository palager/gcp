// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command upspinserver-gcp is a combined DirServer and StoreServer for use on
// stand-alone machines. It provides the production implementations of the
// dir and store servers (dir/server and store/server) with support for storage
// in Google Cloud.
package main // import "github.com/palager/gcp/cmd/upspinserver-gcp"

import (
	"github.com/palager/gcp/cloud/https"

	"github.com/palager/upspin/serverutil/upspinserver"

	// Storage on GCS.
	_ "github.com/palager/gcp/cloud/storage/gcs"
)

func main() {
	ready := upspinserver.Main()
	https.ListenAndServe(ready, "upspinserver")
}
