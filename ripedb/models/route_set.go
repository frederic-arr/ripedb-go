// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = RouteSet{}

type RouteSet struct {
	Object rpsl.Object
}

func (o RouteSet) Class() string {
	return "route-set"
}

func (o RouteSet) Key() string {
	return *o.Object.GetFirst("route-set")
}

func (o RouteSet) Validate() error {
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o RouteSet) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
	schema := `
        route-set:      mandatory  single     primary/lookup key
        descr:          optional   multiple
        members:        optional   multiple
        mp-members:     optional   multiple
        mbrs-by-ref:    optional   multiple   inverse key
        remarks:        optional   multiple
        org:            optional   multiple   inverse key
        tech-c:         mandatory  multiple   inverse key
        admin-c:        mandatory  multiple   inverse key
        notify:         optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-lower:      optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "route-set", &o.Object, skipUnknownKeys, skipKeys)
}

func NewRouteSet(object rpsl.Object) (*RouteSet, error) {
	return NewRouteSetWithOptions(object, false, make([]string, 0))
}

func NewRouteSetWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*RouteSet, error) {
	obj := NewRouteSetUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewRouteSetUnchecked(object rpsl.Object) RouteSet {
	return RouteSet{Object: object}
}
