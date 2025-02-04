// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type FilterSet struct {
	Object rpsl.Object
}

func (o FilterSet) Class() string {
	return "filter-set"
}

func (o FilterSet) Key() string {
	return *o.Object.GetFirst("filter-set")
}

func NewFilterSet(object rpsl.Object) (*FilterSet, error) {
	schema := `
        filter-set:     mandatory  single     primary/lookup key
        descr:          optional   multiple
        filter:         optional   single
        mp-filter:      optional   single
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

	if err := ensureSchema(schema, "filter-set", &object); err != nil {
		return nil, err
	}

	obj := FilterSet{Object: object}
	return &obj, nil
}
