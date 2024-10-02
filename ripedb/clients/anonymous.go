// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

type RipeAnonymousClient struct {
	Endpoint string
	Format   bool
	Filter   bool
	Source   string
}

func (c *RipeAnonymousClient) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

func (c *RipeAnonymousClient) SetSource(source string) {
	c.Source = source
}

func (c *RipeAnonymousClient) SetFilter(filter bool) {
	c.Filter = filter
}

func (c *RipeAnonymousClient) SetFormat(format bool) {
	c.Format = format
}

func (c *RipeAnonymousClient) request(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/%s/%s", c.Endpoint, c.Source, resource, url.PathEscape(key)), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if !c.Format {
		q := req.URL.Query()
		q.Add("unformatted", "")
		req.URL.RawQuery = q.Encode()
	}

	if !c.Filter {
		q := req.URL.Query()
		q.Add("unfiltered", "")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := parseResponse(*resp)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from URL `%s`: %w", req.URL.String(), err)
	}
	return res, nil
}

func (c *RipeAnonymousClient) Get(resource string, key string) (*models.Resource, error) {
	return c.request("GET", resource, key, nil)
}

func (c *RipeAnonymousClient) Post(resource string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot create resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) Put(resource string, key string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot update resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) Delete(resource string, key string) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot delete resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) GetAsBlock(key string) (*resources.AsBlockModel, error) {
	obj, err := findOne(c, "as-block", key)
	if err != nil {
		return nil, err
	}

	return resources.AsBlockFromModel(obj)
}

func (c *RipeAnonymousClient) GetAsSet(key string) (*resources.AsSetModel, error) {
	obj, err := findOne(c, "as-set", key)
	if err != nil {
		return nil, err
	}

	return resources.AsSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetAutNum(key string) (*resources.AutNumModel, error) {
	obj, err := findOne(c, "aut-num", key)
	if err != nil {
		return nil, err
	}

	return resources.AutNumFromModel(obj)
}

func (c *RipeAnonymousClient) GetDomain(key string) (*resources.DomainModel, error) {
	obj, err := findOne(c, "domain", key)
	if err != nil {
		return nil, err
	}

	return resources.DomainFromModel(obj)
}

func (c *RipeAnonymousClient) GetFilterSet(key string) (*resources.FilterSetModel, error) {
	obj, err := findOne(c, "filter-set", key)
	if err != nil {
		return nil, err
	}

	return resources.FilterSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetInetRtr(key string) (*resources.InetRtrModel, error) {
	obj, err := findOne(c, "inet-rtr", key)
	if err != nil {
		return nil, err
	}

	return resources.InetRtrFromModel(obj)
}

func (c *RipeAnonymousClient) GetInet6Num(key string) (*resources.Inet6NumModel, error) {
	obj, err := findOne(c, "inet6num", key)
	if err != nil {
		return nil, err
	}

	return resources.Inet6NumFromModel(obj)
}

func (c *RipeAnonymousClient) GetInetNum(key string) (*resources.InetNumModel, error) {
	obj, err := findOne(c, "inetnum", key)
	if err != nil {
		return nil, err
	}

	return resources.InetNumFromModel(obj)
}

func (c *RipeAnonymousClient) GetIrt(key string) (*resources.IrtModel, error) {
	obj, err := findOne(c, "irt", key)
	if err != nil {
		return nil, err
	}

	return resources.IrtFromModel(obj)
}

func (c *RipeAnonymousClient) GetKeyCert(key string) (*resources.KeyCertModel, error) {
	obj, err := findOne(c, "key-cert", key)
	if err != nil {
		return nil, err
	}

	return resources.KeyCertFromModel(obj)
}

func (c *RipeAnonymousClient) GetMntner(key string) (*resources.MntnerModel, error) {
	obj, err := findOne(c, "mntner", key)
	if err != nil {
		return nil, err
	}

	return resources.MntnerFromModel(obj)
}

func (c *RipeAnonymousClient) GetOrganisation(key string) (*resources.OrganisationModel, error) {
	obj, err := findOne(c, "organisation", key)
	if err != nil {
		return nil, err
	}

	return resources.OrganisationFromModel(obj)
}

func (c *RipeAnonymousClient) GetPeeringSet(key string) (*resources.PeeringSetModel, error) {
	obj, err := findOne(c, "peering-set", key)
	if err != nil {
		return nil, err
	}

	return resources.PeeringSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetPerson(key string) (*resources.PersonModel, error) {
	obj, err := findOne(c, "person", key)
	if err != nil {
		return nil, err
	}

	return resources.PersonFromModel(obj)
}

func (c *RipeAnonymousClient) GetRole(key string) (*resources.RoleModel, error) {
	obj, err := findOne(c, "role", key)
	if err != nil {
		return nil, err
	}

	return resources.RoleFromModel(obj)
}

func (c *RipeAnonymousClient) GetRoute(key string) (*resources.RouteModel, error) {
	obj, err := findOne(c, "route", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteFromModel(obj)
}

func (c *RipeAnonymousClient) GetRouteSet(key string) (*resources.RouteSetModel, error) {
	obj, err := findOne(c, "route-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetRoute6(key string) (*resources.Route6Model, error) {
	obj, err := findOne(c, "route6", key)
	if err != nil {
		return nil, err
	}

	return resources.Route6FromModel(obj)
}

func (c *RipeAnonymousClient) GetRtrSet(key string) (*resources.RtrSetModel, error) {
	obj, err := findOne(c, "rtr-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RtrSetFromModel(obj)
}
