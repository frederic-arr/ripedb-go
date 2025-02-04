package models

import (
	"fmt"

	"github.com/frederic-arr/rpsl-go"
)

type Model interface {
	Class() string
	Key() string
	Validate() error
}

func ObjectToModelUnchecked(resource string, object rpsl.Object) (*Model, error) {
	var m Model
	switch resource {
	case "as-block":
		m = NewAsBlockUnchecked(object)
	case "as-set":
		m = NewAsSetUnchecked(object)
	case "aut-num":
		m = NewAutNumUnchecked(object)
	case "domain":
		m = NewDomainUnchecked(object)
	case "filter-set":
		m = NewFilterSetUnchecked(object)
	case "inet-rtr":
		m = NewInetRtrUnchecked(object)
	case "inet6num":
		m = NewInet6NumUnchecked(object)
	case "inetnum":
		m = NewInetNumUnchecked(object)
	case "irt":
		m = NewIrtUnchecked(object)
	case "key-cert":
		m = NewKeyCertUnchecked(object)
	case "mntner":
		m = NewMntnerUnchecked(object)
	case "organisation":
		m = NewOrganisationUnchecked(object)
	case "peering-set":
		m = NewPeeringSetUnchecked(object)
	case "person":
		m = NewPersonUnchecked(object)
	case "role":
		m = NewRoleUnchecked(object)
	case "route-set":
		m = NewRouteSetUnchecked(object)
	case "route":
		m = NewRouteUnchecked(object)
	case "route6":
		m = NewRoute6Unchecked(object)
	case "rtr-set":
		m = NewRtrSetUnchecked(object)
	default:
		return nil, fmt.Errorf("unknown resource type")
	}

	return &m, nil
}

func ObjectToModel(resource string, object rpsl.Object) (*Model, error) {
	m, err := ObjectToModelUnchecked(resource, object)
	if err != nil {
		return nil, err
	}

	err = (*m).Validate()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ValidateObject(resource string, object rpsl.Object) error {
	_, err := ObjectToModel(resource, object)
	return err
}

func ValidateResource(resource string, data Resource) error {
	if len(data.Objects.Object) != 1 {
		return fmt.Errorf("no object found")
	}

	obj, err := ModelObjectToRpslObject(&data.Objects.Object[0])
	if err != nil {
		return err
	}

	return ValidateObject(resource, *obj)
}
