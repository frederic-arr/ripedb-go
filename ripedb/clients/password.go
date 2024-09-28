// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

type RipePasswordClient struct {
	Endpoint string
	Format   bool
	User     *string
	Password string
	Source   string
}

func (c *RipePasswordClient) request(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	var path string
	if key == "" {
		path = fmt.Sprintf("%s/%s/%s", c.Endpoint, c.Source, resource)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s", c.Endpoint, c.Source, resource, url.PathEscape(key))
	}
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if c.User != nil {
		req.SetBasicAuth(*c.User, c.Password)
	} else {
		q := req.URL.Query()
		q.Add("password", c.Password)
		req.URL.RawQuery = q.Encode()
	}

	if !c.Format {
		q := req.URL.Query()
		q.Add("unformatted", "")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(*resp)
}

func (c *RipePasswordClient) Get(resource string, key string) (*models.Resource, error) {
	return c.request("GET", resource, key, nil)
}

func (c *RipePasswordClient) Post(resource string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("POST", resource, "", bytes.NewReader(body))
}

func (c *RipePasswordClient) Put(resource string, key string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("PUT", resource, key, bytes.NewReader(body))
}

func (c *RipePasswordClient) Delete(resource string, key string) (*models.Resource, error) {
	return c.request("DELETE", resource, key, nil)
}

func (c *RipePasswordClient) GetAsBlock(key string) (*resources.AsBlockModel, error) {
	obj, err := findOne(c, "as-block", key)
	if err != nil {
		return nil, err
	}

	return resources.AsBlockFromModel(obj)
}

func (c *RipePasswordClient) GetAsSet(key string) (*resources.AsSetModel, error) {
	obj, err := findOne(c, "as-set", key)
	if err != nil {
		return nil, err
	}

	return resources.AsSetFromModel(obj)
}

func (c *RipePasswordClient) GetAutNum(key string) (*resources.AutNumModel, error) {
	obj, err := findOne(c, "aut-num", key)
	if err != nil {
		return nil, err
	}

	return resources.AutNumFromModel(obj)
}

func (c *RipePasswordClient) GetDomain(key string) (*resources.DomainModel, error) {
	obj, err := findOne(c, "domain", key)
	if err != nil {
		return nil, err
	}

	return resources.DomainFromModel(obj)
}

func (c *RipePasswordClient) GetFilterSet(key string) (*resources.FilterSetModel, error) {
	obj, err := findOne(c, "filter-set", key)
	if err != nil {
		return nil, err
	}

	return resources.FilterSetFromModel(obj)
}

func (c *RipePasswordClient) GetInetRtr(key string) (*resources.InetRtrModel, error) {
	obj, err := findOne(c, "inet-rtr", key)
	if err != nil {
		return nil, err
	}

	return resources.InetRtrFromModel(obj)
}

func (c *RipePasswordClient) GetInet6Num(key string) (*resources.Inet6NumModel, error) {
	obj, err := findOne(c, "inet6num", key)
	if err != nil {
		return nil, err
	}

	return resources.Inet6NumFromModel(obj)
}

func (c *RipePasswordClient) GetInetNum(key string) (*resources.InetNumModel, error) {
	obj, err := findOne(c, "inetnum", key)
	if err != nil {
		return nil, err
	}

	return resources.InetNumFromModel(obj)
}

func (c *RipePasswordClient) GetIrt(key string) (*resources.IrtModel, error) {
	obj, err := findOne(c, "irt", key)
	if err != nil {
		return nil, err
	}

	return resources.IrtFromModel(obj)
}

func (c *RipePasswordClient) GetKeyCert(key string) (*resources.KeyCertModel, error) {
	obj, err := findOne(c, "key-cert", key)
	if err != nil {
		return nil, err
	}

	return resources.KeyCertFromModel(obj)
}

func (c *RipePasswordClient) GetMntner(key string) (*resources.MntnerModel, error) {
	obj, err := findOne(c, "mntner", key)
	if err != nil {
		return nil, err
	}

	return resources.MntnerFromModel(obj)
}

func (c *RipePasswordClient) GetOrganisation(key string) (*resources.OrganisationModel, error) {
	obj, err := findOne(c, "organisation", key)
	if err != nil {
		return nil, err
	}

	return resources.OrganisationFromModel(obj)
}

func (c *RipePasswordClient) GetPeeringSet(key string) (*resources.PeeringSetModel, error) {
	obj, err := findOne(c, "peering-set", key)
	if err != nil {
		return nil, err
	}

	return resources.PeeringSetFromModel(obj)
}

func (c *RipePasswordClient) GetPerson(key string) (*resources.PersonModel, error) {
	obj, err := findOne(c, "person", key)
	if err != nil {
		return nil, err
	}

	return resources.PersonFromModel(obj)
}

func (c *RipePasswordClient) GetRole(key string) (*resources.RoleModel, error) {
	obj, err := findOne(c, "role", key)
	if err != nil {
		return nil, err
	}

	return resources.RoleFromModel(obj)
}

func (c *RipePasswordClient) GetRoute(key string) (*resources.RouteModel, error) {
	obj, err := findOne(c, "route", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteFromModel(obj)
}

func (c *RipePasswordClient) GetRouteSet(key string) (*resources.RouteSetModel, error) {
	obj, err := findOne(c, "route-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteSetFromModel(obj)
}

func (c *RipePasswordClient) GetRoute6(key string) (*resources.Route6Model, error) {
	obj, err := findOne(c, "route6", key)
	if err != nil {
		return nil, err
	}

	return resources.Route6FromModel(obj)
}

func (c *RipePasswordClient) GetRtrSet(key string) (*resources.RtrSetModel, error) {
	obj, err := findOne(c, "rtr-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RtrSetFromModel(obj)
}
