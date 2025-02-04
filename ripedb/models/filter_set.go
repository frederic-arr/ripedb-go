// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = FilterSet{}

type FilterSet struct {
	Object rpsl.Object
}

func (o FilterSet) Class() string {
	return "filter-set"
}

func (o FilterSet) Key() string {
	return *o.Object.GetFirst("filter-set")
}

func (o FilterSet) Validate() error {
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

	return ensureSchema(schema, "filter-set", &o.Object)
}

func NewFilterSet(object rpsl.Object) (*FilterSet, error) {
	obj := NewFilterSetUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewFilterSetUnchecked(object rpsl.Object) FilterSet {
	return FilterSet{Object: object}
}
