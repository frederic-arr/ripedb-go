package ripedb

import (
	"github.com/frederic-arr/ripedb-go/ripedb/clients"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

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

func NewRipeAnonymousClient() *RipeAnonymousClient {
	return &RipeAnonymousClient{
		Endpoint: clients.RIPE_PROD_ENDPOINT,
		Format:   true,
	}
}
