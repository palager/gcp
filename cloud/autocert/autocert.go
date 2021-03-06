// Copyright 2017 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package autocert provides an autocert.Cache implementation that stores
// the certificate cache in a Google Cloud Storage bucket.
package autocert // import "github.com/palager/gcp/cloud/autocert"

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/palager/upspin/cloud/https"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// NewCache returns an https.AutocertCache that stores the certificate cache in
// the provided Google Cloud Storage bucket using the given file prefix.
func NewCache(bucket, prefix string) (https.AutocertCache, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithScopes(storage.ScopeFullControl))
	if err != nil {
		return nil, err
	}
	return autocertCache{
		b:      client.Bucket(bucket),
		server: prefix + "-",
	}, nil
}

type autocertCache struct {
	b      *storage.BucketHandle
	server string
}

func (cache autocertCache) Get(ctx context.Context, name string) ([]byte, error) {
	r, err := cache.b.Object(cache.server + name).NewReader(ctx)
	if err == storage.ErrObjectNotExist {
		return nil, https.ErrAutocertCacheMiss
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func (cache autocertCache) Put(ctx context.Context, name string, data []byte) error {
	// TODO(ehg) Do we need to add contentType="text/plain; charset=utf-8"?
	w := cache.b.Object(cache.server + name).NewWriter(ctx)
	if _, err := w.Write(data); err != nil {
		log.Printf("https: writing letsencrypt cache: %s %v", name, err)
		return err
	}
	if err := w.Close(); err != nil {
		log.Printf("https: writing letsencrypt cache: %s %v", name, err)
		return err
	}
	return nil
}

func (cache autocertCache) Delete(ctx context.Context, name string) error {
	return cache.b.Object(cache.server + name).Delete(ctx)
}
