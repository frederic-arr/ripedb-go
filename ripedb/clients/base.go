// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeClientOptionsPartial struct {
	Endpoint *string
	Format   *bool
	Filter   *bool
	NoError  *bool
	Source   *string
}

type RipeClientOptions struct {
	Endpoint string
	Format   bool
	Filter   bool
	NoError  bool
	Source   string
}

type RipeDbClient interface {
	SetEndpoint(endpoint string)
	SetSource(source string)
	SetFilter(filter bool)
	SetFormat(format bool)
	SetNoError(noError bool)

	GetObject(resource string, key string) (*rpsl.Object, error)
	CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error)
	UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error)
	DeleteObject(resource string, key string) (*rpsl.Object, error)

	GetResource(resource string, key string) (*models.Resource, error)
	PostResource(resource string, data models.Resource) (*models.Resource, error)
	PutResource(resource string, key string, data models.Resource) (*models.Resource, error)
	DeleteResource(resource string, key string) (*models.Resource, error)
}

func gatherErrors(whoisResponse *models.Resource) []string {
	errors := []string{}
	if whoisResponse.ErrorMessages != nil {
		for _, errorMessage := range whoisResponse.ErrorMessages.ErrorMessage {
			if errorMessage.Text != nil {
				msg := *errorMessage.Text
				args := make([]interface{}, len(errorMessage.Args))
				for i, arg := range errorMessage.Args {
					args[i] = arg.Value
				}
				msg = fmt.Sprintf(msg, args...)

				errors = append(errors, msg)
			}
		}
	}

	return errors
}

func parseResponse(resp http.Response, noError bool) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if err != nil {
		return nil, err
	}

	if !noError && resp.StatusCode != http.StatusOK {
		errors := gatherErrors(whoisResponse)
		return nil, fmt.Errorf("ripedb-go request error: %v", errors)
	}

	return whoisResponse, nil
}

func Lookup(c RipeDbClient, resource string, key string) (*models.Resource, error) {
	resp, err := c.GetResource(resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ValidateResource(resource string, data models.Resource) error {
	if len(data.Objects.Object) != 1 {
		return fmt.Errorf("no object found")
	}

	obj, err := models.ModelObjectToRpslObject(&data.Objects.Object[0])
	if err != nil {
		return err
	}

	_, err = Validate(resource, *obj)
	return err
}

func Validate(resource string, object rpsl.Object) (models.Model, error) {
	switch resource {
	case "as-block":
		m, err := models.NewAsBlock(object)
		return m, err
	case "as-set":
		m, err := models.NewAsSet(object)
		return m, err
	case "aut-num":
		m, err := models.NewAutNum(object)
		return m, err
	case "domain":
		m, err := models.NewDomain(object)
		return m, err
	case "filter-set":
		m, err := models.NewFilterSet(object)
		return m, err
	case "inet-rtr":
		m, err := models.NewInetRtr(object)
		return m, err
	case "inet6num":
		m, err := models.NewInet6Num(object)
		return m, err
	case "inetnum":
		m, err := models.NewInetNum(object)
		return m, err
	case "irt":
		m, err := models.NewIrt(object)
		return m, err
	case "key-cert":
		m, err := models.NewKeyCert(object)
		return m, err
	case "mntner":
		m, err := models.NewMntner(object)
		return m, err
	case "organisation":
		m, err := models.NewOrganisation(object)
		return m, err
	case "peering-set":
		m, err := models.NewPeeringSet(object)
		return m, err
	case "person":
		m, err := models.NewPerson(object)
		return m, err
	case "role":
		m, err := models.NewRole(object)
		return m, err
	case "route-set":
		m, err := models.NewRouteSet(object)
		return m, err
	case "route":
		m, err := models.NewRoute(object)
		return m, err
	case "route6":
		m, err := models.NewRoute6(object)
		return m, err
	case "rtr-set":
		m, err := models.NewRtrSet(object)
		return m, err
	}

	return nil, fmt.Errorf("unknown resource")
}
