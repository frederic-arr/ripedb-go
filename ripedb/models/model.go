package models

import (
	"fmt"

	"github.com/frederic-arr/rpsl-go"
)

type Model interface {
	Class() string
	Key() string
	Validate() error
	ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error
}

func ObjectToModelUnchecked(resource string, object rpsl.Object) Model {
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
	}

	return m
}

func ObjectToModel(resource string, object rpsl.Object) (Model, error) {
	return ObjectToModelWithOptions(resource, object, false, []string{})
}

func ObjectToModelWithOptions(resource string, object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (Model, error) {
	m := ObjectToModelUnchecked(resource, object)

	err := m.ValidateWithOptions(skipUnknownKeys, skipKeys)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ValidateObject(resource string, object rpsl.Object) error {
	return ValidateObjectWithOptions(resource, object, false, []string{})
}

func ValidateObjectWithOptions(resource string, object rpsl.Object, skipUnknownKeys bool, skipKeys []string) error {
	_, err := ObjectToModelWithOptions(resource, object, skipUnknownKeys, skipKeys)
	return err
}

func ValidateResource(resource string, data Resource) error {
	return ValidateResourceWithOptions(resource, data, false, []string{})
}

func ValidateResourceWithOptions(resource string, data Resource, skipUnknownKeys bool, skipKeys []string) error {
	if len(data.Objects.Object) != 1 {
		return fmt.Errorf("no object found")
	}

	obj, err := ModelObjectToRpslObject(&data.Objects.Object[0])
	if err != nil {
		return err
	}

	return ValidateObjectWithOptions(resource, *obj, skipUnknownKeys, skipKeys)
}
