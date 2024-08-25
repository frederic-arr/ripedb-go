package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/ripedb-go/ripedb/resources"
)

type RipeDbClient interface {
	Get(resource string, key string) (*models.WhoisResponseModel, error)

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

const (
	RIPE_TEST_ENDPOINT_INSECURE = "http://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT          = "https://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT_MTLS     = "https://rest-cert-test.db.ripe.net"
	RIPE_PROD_ENDPOINT_INSECURE = "http://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT          = "https://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT_MTLS     = "https://rest-cert.db.ripe.net"
)

func gatherErrors(whoisResponse *models.WhoisResponseModel) []string {
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

func parseResponse(resp http.Response) (*models.WhoisResponseModel, error) {
	whoisResponse := &models.WhoisResponseModel{}
	json.NewDecoder(resp.Body).Decode(whoisResponse)

	// if resp.StatusCode != http.StatusOK || jsonErr != nil {
	// 	if jsonErr == nil {
	// 		errors := gatherErrors(whoisResponse)
	// 		return nil, fmt.Errorf("error: %v", errors)
	// 	}

	// 	return nil, jsonErr
	// }

	return whoisResponse, nil
}

func Lookup(c RipeDbClient, resource string, key string) (*models.WhoisResponseModel, error) {
	resp, err := c.Get(resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func findOne(c RipeDbClient, resource string, key string) (*models.WhoisObjectModel, error) {
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
