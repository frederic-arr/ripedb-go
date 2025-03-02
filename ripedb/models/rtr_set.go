// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = RtrSet{}

type RtrSet struct {
	Object rpsl.Object
}

func (o RtrSet) Class() string {
	return "rtr-set"
}

func (o RtrSet) Key() string {
	return *o.Object.GetFirst("rtr-set")
}

func (o RtrSet) Validate() error {
	schema := `
        rtr-set:        mandatory  single     primary/lookup key
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

	return ensureSchema(schema, "rtr-set", &o.Object)
}

func NewRtrSet(object rpsl.Object) (*RtrSet, error) {
	obj := NewRtrSetUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewRtrSetUnchecked(object rpsl.Object) RtrSet {
	return RtrSet{Object: object}
}
