// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"github.com/frederic-arr/ripedb-go/ripedb/clients"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

type RipeDbClient = clients.RipeDbClient
type RipeAnonymousClient = clients.RipeAnonymousClient

type AsBlockModel = resources.AsBlockModel
type AsSetModel = resources.AsSetModel
type AutNumModel = resources.AutNumModel
type DomainModel = resources.DomainModel
type FilterSetModel = resources.FilterSetModel
type InetRtrModel = resources.InetRtrModel
type Inet6NumModel = resources.Inet6NumModel
type InetNumModel = resources.InetNumModel
type IrtModel = resources.IrtModel
type KeyCertModel = resources.KeyCertModel
type MntnerModel = resources.MntnerModel
type OrganisationModel = resources.OrganisationModel
type PeeringSetModel = resources.PeeringSetModel
type PersonModel = resources.PersonModel
type RoleModel = resources.RoleModel
type RouteModel = resources.RouteModel
type RouteSetModel = resources.RouteSetModel
type Route6Model = resources.Route6Model
type RtrSetModel = resources.RtrSetModel

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
