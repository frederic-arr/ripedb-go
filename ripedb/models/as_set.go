// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type AsSet struct {
	Object rpsl.Object
}

func (o AsSet) Class() string {
	return "as-set"
}

func (o AsSet) Key() string {
	return *o.Object.GetFirst("as-set")
}

func NewAsSet(object rpsl.Object) (*AsSet, error) {
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

	if err := ensureSchema(schema, "as-set", &object); err != nil {
		return nil, err
	}

	obj := AsSet{Object: object}
	return &obj, nil
}
