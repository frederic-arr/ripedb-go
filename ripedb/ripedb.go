// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"github.com/frederic-arr/ripedb-go/ripedb/clients"
)

type RipeDbClient = clients.RipeDbClient
type RipeAnonymousClient = clients.RipeAnonymousClient
type RipePasswordClient = clients.RipePasswordClient

const (
	RIPE_TEST_ENDPOINT_INSECURE = "http://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT          = "https://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT_MTLS     = "https://rest-cert-test.db.ripe.net"
	RIPE_PROD_ENDPOINT_INSECURE = "http://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT          = "https://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT_MTLS     = "https://rest-cert.db.ripe.net"
)

func partialToOptions(input *clients.RipeClientOptionsPartial) clients.RipeClientOptions {
	opts := clients.RipeClientOptions{
		Endpoint: RIPE_PROD_ENDPOINT,
		Filter:   false,
		Format:   true,
		NoError:  false,
		Source:   "ripe",
	}

	partial := clients.RipeClientOptionsPartial{}
	if input != nil {
		partial = *input
	}

	if partial.Endpoint != nil {
		opts.Endpoint = *partial.Endpoint
	}

	if partial.Filter != nil {
		opts.Filter = *partial.Filter
	}

	if partial.Format != nil {
		opts.Format = *partial.Format
	}

	if partial.NoError != nil {
		opts.NoError = *partial.NoError
	}

	if partial.Source != nil {
		opts.Source = *partial.Source
	}

	return opts
}

func NewRipeAnonymousClient(opts *clients.RipeClientOptionsPartial) *RipeAnonymousClient {
	return &RipeAnonymousClient{
		Opts: partialToOptions(opts),
	}
}

func NewRipePasswordClient(user *string, password string, opts *clients.RipeClientOptionsPartial) *clients.RipePasswordClient {
	return &clients.RipePasswordClient{
		Opts:     partialToOptions(opts),
		User:     user,
		Password: password,
	}
}
