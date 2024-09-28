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
}

func (c *RipePasswordClient) request(method string, source string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	var path string
	if key == "" {
		path = fmt.Sprintf("%s/%s/%s", c.Endpoint, source, resource)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s", c.Endpoint, source, resource, url.PathEscape(key))
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

func (c *RipePasswordClient) Get(source string, resource string, key string) (*models.Resource, error) {
	return c.request("GET", source, resource, key, nil)
}

func (c *RipePasswordClient) Post(source string, resource string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("POST", source, resource, "", bytes.NewReader(body))
}

func (c *RipePasswordClient) Put(source string, resource string, key string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("PUT", source, resource, key, bytes.NewReader(body))
}

func (c *RipePasswordClient) Delete(source string, resource string, key string) (*models.Resource, error) {
	return c.request("DELETE", source, resource, key, nil)
}

func (c *RipePasswordClient) GetAsBlock(source string, key string) (*resources.AsBlockModel, error) {
	obj, err := findOne(c, source, "as-block", key)
	if err != nil {
		return nil, err
	}

	return resources.AsBlockFromModel(obj)
}

func (c *RipePasswordClient) GetAsSet(source string, key string) (*resources.AsSetModel, error) {
	obj, err := findOne(c, source, "as-set", key)
	if err != nil {
		return nil, err
	}

	return resources.AsSetFromModel(obj)
}

func (c *RipePasswordClient) GetAutNum(source string, key string) (*resources.AutNumModel, error) {
	obj, err := findOne(c, source, "aut-num", key)
	if err != nil {
		return nil, err
	}

	return resources.AutNumFromModel(obj)
}

func (c *RipePasswordClient) GetDomain(source string, key string) (*resources.DomainModel, error) {
	obj, err := findOne(c, source, "domain", key)
	if err != nil {
		return nil, err
	}

	return resources.DomainFromModel(obj)
}

func (c *RipePasswordClient) GetFilterSet(source string, key string) (*resources.FilterSetModel, error) {
	obj, err := findOne(c, source, "filter-set", key)
	if err != nil {
		return nil, err
	}

	return resources.FilterSetFromModel(obj)
}

func (c *RipePasswordClient) GetInetRtr(source string, key string) (*resources.InetRtrModel, error) {
	obj, err := findOne(c, source, "inet-rtr", key)
	if err != nil {
		return nil, err
	}

	return resources.InetRtrFromModel(obj)
}

func (c *RipePasswordClient) GetInet6Num(source string, key string) (*resources.Inet6NumModel, error) {
	obj, err := findOne(c, source, "inet6num", key)
	if err != nil {
		return nil, err
	}

	return resources.Inet6NumFromModel(obj)
}

func (c *RipePasswordClient) GetInetNum(source string, key string) (*resources.InetNumModel, error) {
	obj, err := findOne(c, source, "inetnum", key)
	if err != nil {
		return nil, err
	}

	return resources.InetNumFromModel(obj)
}

func (c *RipePasswordClient) GetIrt(source string, key string) (*resources.IrtModel, error) {
	obj, err := findOne(c, source, "irt", key)
	if err != nil {
		return nil, err
	}

	return resources.IrtFromModel(obj)
}

func (c *RipePasswordClient) GetKeyCert(source string, key string) (*resources.KeyCertModel, error) {
	obj, err := findOne(c, source, "key-cert", key)
	if err != nil {
		return nil, err
	}

	return resources.KeyCertFromModel(obj)
}

func (c *RipePasswordClient) GetMntner(source string, key string) (*resources.MntnerModel, error) {
	obj, err := findOne(c, source, "mntner", key)
	if err != nil {
		return nil, err
	}

	return resources.MntnerFromModel(obj)
}

func (c *RipePasswordClient) GetOrganisation(source string, key string) (*resources.OrganisationModel, error) {
	obj, err := findOne(c, source, "organisation", key)
	if err != nil {
		return nil, err
	}

	return resources.OrganisationFromModel(obj)
}

func (c *RipePasswordClient) GetPeeringSet(source string, key string) (*resources.PeeringSetModel, error) {
	obj, err := findOne(c, source, "peering-set", key)
	if err != nil {
		return nil, err
	}

	return resources.PeeringSetFromModel(obj)
}

func (c *RipePasswordClient) GetPerson(source string, key string) (*resources.PersonModel, error) {
	obj, err := findOne(c, source, "person", key)
	if err != nil {
		return nil, err
	}

	return resources.PersonFromModel(obj)
}

func (c *RipePasswordClient) GetRole(source string, key string) (*resources.RoleModel, error) {
	obj, err := findOne(c, source, "role", key)
	if err != nil {
		return nil, err
	}

	return resources.RoleFromModel(obj)
}

func (c *RipePasswordClient) GetRoute(source string, key string) (*resources.RouteModel, error) {
	obj, err := findOne(c, source, "route", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteFromModel(obj)
}

func (c *RipePasswordClient) GetRouteSet(source string, key string) (*resources.RouteSetModel, error) {
	obj, err := findOne(c, source, "route-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteSetFromModel(obj)
}

func (c *RipePasswordClient) GetRoute6(source string, key string) (*resources.Route6Model, error) {
	obj, err := findOne(c, source, "route6", key)
	if err != nil {
		return nil, err
	}

	return resources.Route6FromModel(obj)
}

func (c *RipePasswordClient) GetRtrSet(source string, key string) (*resources.RtrSetModel, error) {
	obj, err := findOne(c, source, "rtr-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RtrSetFromModel(obj)
}
