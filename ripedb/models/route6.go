// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Route6 struct {
	Object rpsl.Object
}

func (o *Route6) Class() string {
	return "route6"
}

func (o *Route6) Key() string {
	return *o.Object.GetFirst("route6") + *o.Object.GetFirst("origin")
}

func NewRoute6(object rpsl.Object) (*Route6, error) {
	schema := `
        route6:         mandatory  single     primary/lookup key
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

	if err := ensureSchema(schema, "route6", &object); err != nil {
		return nil, err
	}

	obj := Route6{Object: object}
	return &obj, nil
}
