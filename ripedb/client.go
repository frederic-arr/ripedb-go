package ripedb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RipeDbClient interface {
	Get(resource string, key string) (*WhoisResponseModel, error)

	GetAsBlock(key string) (*AsBlockModel, error)
	GetAsSet(key string) (*AsSetModel, error)
	GetAutNum(key string) (*AutNumModel, error)
	GetDomain(key string) (*DomainModel, error)
	GetFilterSet(key string) (*FilterSetModel, error)
	GetInetRtr(key string) (*InetRtrModel, error)
	GetInet6Num(key string) (*Inet6NumModel, error)
	GetInetNum(key string) (*InetNumModel, error)
	GetIrt(key string) (*IrtModel, error)
	GetKeyCert(key string) (*KeyCertModel, error)
	GetMntner(key string) (*MntnerModel, error)
	GetOrganisation(key string) (*OrganisationModel, error)
	GetPeeringSet(key string) (*PeeringSetModel, error)
	GetPerson(key string) (*PersonModel, error)
	GetRole(key string) (*RoleModel, error)
	GetRoute(key string) (*RouteModel, error)
	GetRouteSet(key string) (*RouteSetModel, error)
	GetRoute6(key string) (*Route6Model, error)
	GetRtrSet(key string) (*RtrSetModel, error)
}

const (
	RIPE_TEST_ENDPOINT_INSECURE = "http://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT          = "https://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT_MTLS     = "https://rest-cert-test.db.ripe.net"
	RIPE_PROD_ENDPOINT_INSECURE = "http://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT          = "https://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT_MTLS     = "https://rest-cert.db.ripe.net"
)

type RipeAnonymousClient struct {
	Endpoint string
}

func gatherErrors(whoisResponse *WhoisResponseModel) []string {
	errors := []string{}
	if whoisResponse.ErrorMessages != nil {
		for _, errorMessage := range whoisResponse.ErrorMessages.ErrorMessage {
			if errorMessage.Text != nil {
				errors = append(errors, *errorMessage.Text)
			}
		}
	}

	return errors
}

func parseResponse(resp http.Response) (*WhoisResponseModel, error) {
	whoisResponse := &WhoisResponseModel{}
	jsonErr := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if resp.StatusCode != http.StatusOK || jsonErr != nil {
		if jsonErr == nil {
			errors := gatherErrors(whoisResponse)
			return nil, fmt.Errorf("error: %v", errors)
		}

		return nil, jsonErr
	}

	return whoisResponse, nil
}

func Lookup(c RipeDbClient, resource string, key string) (*WhoisObjectModel, error) {
	resp, err := c.Get(resource, key)
	if err != nil {
		return nil, err
	}

	if resp.Objects == nil || resp.Objects.Object == nil || len(resp.Objects.Object) == 0 {
		return nil, fmt.Errorf("no objects found")
	}

	if len(resp.Objects.Object) > 1 {
		return nil, fmt.Errorf("more than one object found")
	}

	return &resp.Objects.Object[0], nil
}

func (c *RipeAnonymousClient) Get(resource string, key string) (*WhoisResponseModel, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ripe/%s/%s", c.Endpoint, resource, url.PathEscape(key)), nil)
	req.Header.Add("Accept", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(*resp)
}

func (c *RipeAnonymousClient) GetAsBlock(key string) (*AsBlockModel, error) {
	obj, err := Lookup(c, "as-block", key)
	if err != nil {
		return nil, err
	}

	return asBlock(obj)
}

func (c *RipeAnonymousClient) GetAsSet(key string) (*AsSetModel, error) {
	obj, err := Lookup(c, "as-set", key)
	if err != nil {
		return nil, err
	}

	return asSet(obj)
}

func (c *RipeAnonymousClient) GetAutNum(key string) (*AutNumModel, error) {
	obj, err := Lookup(c, "aut-num", key)
	if err != nil {
		return nil, err
	}

	return autNum(obj)
}

func (c *RipeAnonymousClient) GetDomain(key string) (*DomainModel, error) {
	obj, err := Lookup(c, "domain", key)
	if err != nil {
		return nil, err
	}

	return domain(obj)
}

func (c *RipeAnonymousClient) GetFilterSet(key string) (*FilterSetModel, error) {
	obj, err := Lookup(c, "filter-set", key)
	if err != nil {
		return nil, err
	}

	return filterSet(obj)
}

func (c *RipeAnonymousClient) GetInetRtr(key string) (*InetRtrModel, error) {
	obj, err := Lookup(c, "inet-rtr", key)
	if err != nil {
		return nil, err
	}

	return inetRtr(obj)
}

func (c *RipeAnonymousClient) GetInet6Num(key string) (*Inet6NumModel, error) {
	obj, err := Lookup(c, "inet6num", key)
	if err != nil {
		return nil, err
	}

	return inet6Num(obj)
}

func (c *RipeAnonymousClient) GetInetNum(key string) (*InetNumModel, error) {
	obj, err := Lookup(c, "inetnum", key)
	if err != nil {
		return nil, err
	}

	return inetNum(obj)
}

func (c *RipeAnonymousClient) GetIrt(key string) (*IrtModel, error) {
	obj, err := Lookup(c, "irt", key)
	if err != nil {
		return nil, err
	}

	return irt(obj)
}

func (c *RipeAnonymousClient) GetKeyCert(key string) (*KeyCertModel, error) {
	obj, err := Lookup(c, "key-cert", key)
	if err != nil {
		return nil, err
	}

	return keyCert(obj)
}

func (c *RipeAnonymousClient) GetMntner(key string) (*MntnerModel, error) {
	obj, err := Lookup(c, "mntner", key)
	if err != nil {
		return nil, err
	}

	return mntner(obj)
}

func (c *RipeAnonymousClient) GetOrganisation(key string) (*OrganisationModel, error) {
	obj, err := Lookup(c, "organisation", key)
	if err != nil {
		return nil, err
	}

	return organisation(obj)
}

func (c *RipeAnonymousClient) GetPeeringSet(key string) (*PeeringSetModel, error) {
	obj, err := Lookup(c, "peering-set", key)
	if err != nil {
		return nil, err
	}

	return peeringSet(obj)
}

func (c *RipeAnonymousClient) GetPerson(key string) (*PersonModel, error) {
	obj, err := Lookup(c, "person", key)
	if err != nil {
		return nil, err
	}

	return person(obj)
}

func (c *RipeAnonymousClient) GetRole(key string) (*RoleModel, error) {
	obj, err := Lookup(c, "role", key)
	if err != nil {
		return nil, err
	}

	return role(obj)
}

func (c *RipeAnonymousClient) GetRoute(key string) (*RouteModel, error) {
	obj, err := Lookup(c, "route", key)
	if err != nil {
		return nil, err
	}

	return route(obj)
}

func (c *RipeAnonymousClient) GetRouteSet(key string) (*RouteSetModel, error) {
	obj, err := Lookup(c, "route-set", key)
	if err != nil {
		return nil, err
	}

	return routeSet(obj)
}

func (c *RipeAnonymousClient) GetRoute6(key string) (*Route6Model, error) {
	obj, err := Lookup(c, "route6", key)
	if err != nil {
		return nil, err
	}

	return route6(obj)
}

func (c *RipeAnonymousClient) GetRtrSet(key string) (*RtrSetModel, error) {
	obj, err := Lookup(c, "rtr-set", key)
	if err != nil {
		return nil, err
	}

	return rtrSet(obj)
}
