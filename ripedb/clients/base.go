// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
	"github.com/frederic-arr/rpsl-go"
)

type RipeDbClient interface {
	SetEndpoint(endpoint string)
	SetSource(source string)
	SetFilter(filter bool)
	SetFormat(format bool)

	GetObject(resource string, key string) (*rpsl.Object, error)
	CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error)
	UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error)
	DeleteObject(resource string, key string) (*rpsl.Object, error)

	Get(resource string, key string) (*models.Resource, error)
	Post(resource string, data models.Resource) (*models.Resource, error)
	Put(resource string, key string, data models.Resource) (*models.Resource, error)
	Delete(resource string, key string) (*models.Resource, error)

	GetAsBlock(key string) (*resources.AsBlockModel, error)
	GetAsSet(key string) (*resources.AsSetModel, error)
	GetAutNum(key string) (*resources.AutNumModel, error)
	GetDomain(key string) (*resources.DomainModel, error)
	GetFilterSet(key string) (*resources.FilterSetModel, error)
	GetInetRtr(key string) (*resources.InetRtrModel, error)
	GetInet6Num(key string) (*resources.Inet6NumModel, error)
	GetInetNum(key string) (*resources.InetNumModel, error)
	GetIrt(key string) (*resources.IrtModel, error)
	GetKeyCert(key string) (*resources.KeyCertModel, error)
	GetMntner(key string) (*resources.MntnerModel, error)
	GetOrganisation(key string) (*resources.OrganisationModel, error)
	GetPeeringSet(key string) (*resources.PeeringSetModel, error)
	GetPerson(key string) (*resources.PersonModel, error)
	GetRole(key string) (*resources.RoleModel, error)
	GetRoute(key string) (*resources.RouteModel, error)
	GetRouteSet(key string) (*resources.RouteSetModel, error)
	GetRoute6(key string) (*resources.Route6Model, error)
	GetRtrSet(key string) (*resources.RtrSetModel, error)
}

func gatherErrors(whoisResponse *models.Resource) []string {
	errors := []string{}
	if whoisResponse.ErrorMessages != nil {
		for _, errorMessage := range whoisResponse.ErrorMessages.ErrorMessage {
			if errorMessage.PlainText != nil {
				errors = append(errors, *errorMessage.PlainText)
			}
		}
	}

	return errors
}

func parseResponse(resp http.Response) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if resp.StatusCode != http.StatusOK || err != nil {
		if err == nil {
			errors := gatherErrors(whoisResponse)
			return nil, fmt.Errorf("ripedb-go request error: %v", errors)
		}

		return nil, err
	}

	return whoisResponse, nil
}

func Lookup(c RipeDbClient, resource string, key string) (*models.Resource, error) {
	resp, err := c.Get(resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func findOne(c RipeDbClient, resource string, key string) (*models.Object, error) {
	resp, err := c.Get(resource, key)
	if err != nil {
		return nil, err
	}

	return resp.FindOne()
}
