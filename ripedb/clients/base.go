package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

type RipeDbClient interface {
	Get(source string, resource string, key string) (*models.Resource, error)
	Post(source string, resource string, data models.Resource) (*models.Resource, error)
	Put(source string, resource string, key string, data models.Resource) (*models.Resource, error)
	Delete(source string, resource string, key string) (*models.Resource, error)

	GetAsBlock(source string, key string) (*resources.AsBlockModel, error)
	GetAsSet(source string, key string) (*resources.AsSetModel, error)
	GetAutNum(source string, key string) (*resources.AutNumModel, error)
	GetDomain(source string, key string) (*resources.DomainModel, error)
	GetFilterSet(source string, key string) (*resources.FilterSetModel, error)
	GetInetRtr(source string, key string) (*resources.InetRtrModel, error)
	GetInet6Num(source string, key string) (*resources.Inet6NumModel, error)
	GetInetNum(source string, key string) (*resources.InetNumModel, error)
	GetIrt(source string, key string) (*resources.IrtModel, error)
	GetKeyCert(source string, key string) (*resources.KeyCertModel, error)
	GetMntner(source string, key string) (*resources.MntnerModel, error)
	GetOrganisation(source string, key string) (*resources.OrganisationModel, error)
	GetPeeringSet(source string, key string) (*resources.PeeringSetModel, error)
	GetPerson(source string, key string) (*resources.PersonModel, error)
	GetRole(source string, key string) (*resources.RoleModel, error)
	GetRoute(source string, key string) (*resources.RouteModel, error)
	GetRouteSet(source string, key string) (*resources.RouteSetModel, error)
	GetRoute6(source string, key string) (*resources.Route6Model, error)
	GetRtrSet(source string, key string) (*resources.RtrSetModel, error)
}

const (
	RIPE_TEST_ENDPOINT_INSECURE = "http://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT          = "https://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT_MTLS     = "https://rest-cert-test.db.ripe.net"
	RIPE_PROD_ENDPOINT_INSECURE = "http://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT          = "https://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT_MTLS     = "https://rest-cert.db.ripe.net"
)

func gatherErrors(whoisResponse *models.Resource) []string {
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

func parseResponse(resp http.Response) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if resp.StatusCode != http.StatusOK || err != nil {
		if err == nil {
			errors := gatherErrors(whoisResponse)
			return nil, fmt.Errorf("error: %v", errors)
		}

		return nil, err
	}

	return whoisResponse, nil
}

func Lookup(c RipeDbClient, source string, resource string, key string) (*models.Resource, error) {
	resp, err := c.Get(source, resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func findOne(c RipeDbClient, source string, resource string, key string) (*models.Object, error) {
	resp, err := c.Get(source, resource, key)
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
