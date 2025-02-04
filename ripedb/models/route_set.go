// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type RouteSet struct {
	Object rpsl.Object
}

func (o RouteSet) Class() string {
	return "route-set"
}

func (o RouteSet) Key() string {
	return *o.Object.GetFirst("route-set")
}

func NewRouteSet(object rpsl.Object) (*RouteSet, error) {
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

	if err := ensureSchema(schema, "route-set", &object); err != nil {
		return nil, err
	}

	obj := RouteSet{Object: object}
	return &obj, nil
}
