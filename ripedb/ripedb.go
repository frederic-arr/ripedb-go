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

func NewRipeAnonymousClient() *RipeAnonymousClient {
	return &RipeAnonymousClient{
		Endpoint: RIPE_PROD_ENDPOINT,
		Filter:   false,
		Format:   true,
		Source:   "ripe",
	}
}

func NewRipePasswordClient(user *string, password string) *clients.RipePasswordClient {
	return &clients.RipePasswordClient{
		Endpoint: RIPE_PROD_ENDPOINT,
		Filter:   false,
		Format:   true,
		User:     user,
		Password: password,
		Source:   "ripe",
	}
}
