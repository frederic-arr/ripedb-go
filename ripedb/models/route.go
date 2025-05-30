// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Route{}

type Route struct {
	Object rpsl.Object
}

func (o Route) Class() string {
	return "route"
}

func (o Route) Key() string {
	return *o.Object.GetFirst("route") + *o.Object.GetFirst("origin")
}

func (o Route) Validate() error {
	schema := `
        route:          mandatory  single     primary/lookup key
        descr:          optional   multiple
        origin:         mandatory  single     primary/inverse key
        pingable:       optional   multiple
        ping-hdl:       optional   multiple   inverse key
        holes:          optional   multiple
        org:            optional   multiple   inverse key
        member-of:      optional   multiple   inverse key
        inject:         optional   multiple
        aggr-mtd:       optional   single
        aggr-bndry:     optional   single
        export-comps:   optional   single
        components:     optional   single
        remarks:        optional   multiple
        notify:         optional   multiple   inverse key
        mnt-lower:      optional   multiple   inverse key
        mnt-routes:     optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "route", &o.Object)
}

func NewRoute(object rpsl.Object) (*Route, error) {
	obj := NewRouteUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewRouteUnchecked(object rpsl.Object) Route {
	return Route{Object: object}
}
