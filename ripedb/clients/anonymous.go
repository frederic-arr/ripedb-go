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
}

func (c *RipeAnonymousClient) request(method string, source string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/%s/%s", c.Endpoint, source, resource, url.PathEscape(key)), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
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

	res, err := parseResponse(*resp)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from URL `%s`: %w", req.URL.String(), err)
	}
	return res, nil
}

func (c *RipeAnonymousClient) Get(source string, resource string, key string) (*models.Resource, error) {
	return c.request("GET", source, resource, key, nil)
}

func (c *RipeAnonymousClient) Post(source string, resource string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot create resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) Put(source string, resource string, key string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot update resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) Delete(source string, resource string, key string) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot delete resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) GetAsBlock(source string, key string) (*resources.AsBlockModel, error) {
	obj, err := findOne(c, source, "as-block", key)
	if err != nil {
		return nil, err
	}

	return resources.AsBlockFromModel(obj)
}

func (c *RipeAnonymousClient) GetAsSet(source string, key string) (*resources.AsSetModel, error) {
	obj, err := findOne(c, source, "as-set", key)
	if err != nil {
		return nil, err
	}

	return resources.AsSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetAutNum(source string, key string) (*resources.AutNumModel, error) {
	obj, err := findOne(c, source, "aut-num", key)
	if err != nil {
		return nil, err
	}

	return resources.AutNumFromModel(obj)
}

func (c *RipeAnonymousClient) GetDomain(source string, key string) (*resources.DomainModel, error) {
	obj, err := findOne(c, source, "domain", key)
	if err != nil {
		return nil, err
	}

	return resources.DomainFromModel(obj)
}

func (c *RipeAnonymousClient) GetFilterSet(source string, key string) (*resources.FilterSetModel, error) {
	obj, err := findOne(c, source, "filter-set", key)
	if err != nil {
		return nil, err
	}

	return resources.FilterSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetInetRtr(source string, key string) (*resources.InetRtrModel, error) {
	obj, err := findOne(c, source, "inet-rtr", key)
	if err != nil {
		return nil, err
	}

	return resources.InetRtrFromModel(obj)
}

func (c *RipeAnonymousClient) GetInet6Num(source string, key string) (*resources.Inet6NumModel, error) {
	obj, err := findOne(c, source, "inet6num", key)
	if err != nil {
		return nil, err
	}

	return resources.Inet6NumFromModel(obj)
}

func (c *RipeAnonymousClient) GetInetNum(source string, key string) (*resources.InetNumModel, error) {
	obj, err := findOne(c, source, "inetnum", key)
	if err != nil {
		return nil, err
	}

	return resources.InetNumFromModel(obj)
}

func (c *RipeAnonymousClient) GetIrt(source string, key string) (*resources.IrtModel, error) {
	obj, err := findOne(c, source, "irt", key)
	if err != nil {
		return nil, err
	}

	return resources.IrtFromModel(obj)
}

func (c *RipeAnonymousClient) GetKeyCert(source string, key string) (*resources.KeyCertModel, error) {
	obj, err := findOne(c, source, "key-cert", key)
	if err != nil {
		return nil, err
	}

	return resources.KeyCertFromModel(obj)
}

func (c *RipeAnonymousClient) GetMntner(source string, key string) (*resources.MntnerModel, error) {
	obj, err := findOne(c, source, "mntner", key)
	if err != nil {
		return nil, err
	}

	return resources.MntnerFromModel(obj)
}

func (c *RipeAnonymousClient) GetOrganisation(source string, key string) (*resources.OrganisationModel, error) {
	obj, err := findOne(c, source, "organisation", key)
	if err != nil {
		return nil, err
	}

	return resources.OrganisationFromModel(obj)
}

func (c *RipeAnonymousClient) GetPeeringSet(source string, key string) (*resources.PeeringSetModel, error) {
	obj, err := findOne(c, source, "peering-set", key)
	if err != nil {
		return nil, err
	}

	return resources.PeeringSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetPerson(source string, key string) (*resources.PersonModel, error) {
	obj, err := findOne(c, source, "person", key)
	if err != nil {
		return nil, err
	}

	return resources.PersonFromModel(obj)
}

func (c *RipeAnonymousClient) GetRole(source string, key string) (*resources.RoleModel, error) {
	obj, err := findOne(c, source, "role", key)
	if err != nil {
		return nil, err
	}

	return resources.RoleFromModel(obj)
}

func (c *RipeAnonymousClient) GetRoute(source string, key string) (*resources.RouteModel, error) {
	obj, err := findOne(c, source, "route", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteFromModel(obj)
}

func (c *RipeAnonymousClient) GetRouteSet(source string, key string) (*resources.RouteSetModel, error) {
	obj, err := findOne(c, source, "route-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RouteSetFromModel(obj)
}

func (c *RipeAnonymousClient) GetRoute6(source string, key string) (*resources.Route6Model, error) {
	obj, err := findOne(c, source, "route6", key)
	if err != nil {
		return nil, err
	}

	return resources.Route6FromModel(obj)
}

func (c *RipeAnonymousClient) GetRtrSet(source string, key string) (*resources.RtrSetModel, error) {
	obj, err := findOne(c, source, "rtr-set", key)
	if err != nil {
		return nil, err
	}

	return resources.RtrSetFromModel(obj)
}
