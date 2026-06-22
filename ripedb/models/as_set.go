// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = AsSet{}

type AsSet struct {
	Object rpsl.Object
}

func (o AsSet) Class() string {
	return "as-set"
}

func (o AsSet) Key() string {
	return *o.Object.GetFirst("as-set")
}

func (o AsSet) Validate() error {
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o AsSet) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
	schema := `
        as-set:         mandatory  single     primary/lookup key
        descr:          optional   multiple
        members:        optional   multiple
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

	return ensureSchema(schema, "as-set", &o.Object, skipUnknownKeys, skipKeys)
}

func NewAsSet(object rpsl.Object) (*AsSet, error) {
	return NewAsSetWithOptions(object, false, make([]string, 0))
}

func NewAsSetWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*AsSet, error) {
	obj := NewAsSetUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewAsSetUnchecked(object rpsl.Object) AsSet {
	return AsSet{Object: object}
}
