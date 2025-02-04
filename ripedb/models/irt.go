// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Irt{}

type Irt struct {
	Object rpsl.Object
}

func (o Irt) Class() string {
	return "irt"
}

func (o Irt) Key() string {
	return *o.Object.GetFirst("irt")
}

func (o Irt) Validate() error {
	schema := `
        irt:            mandatory  single     primary/lookup key
        address:        mandatory  multiple
        phone:          optional   multiple
        fax-no:         optional   multiple
        e-mail:         mandatory  multiple   lookup key
        signature:      optional   multiple
        encryption:     optional   multiple
        org:            optional   multiple   inverse key
        admin-c:        mandatory  multiple   inverse key
        tech-c:         mandatory  multiple   inverse key
        auth:           mandatory  multiple   inverse key
        remarks:        optional   multiple
        irt-nfy:        optional   multiple   inverse key
        notify:         optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-ref:        optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "irt", &o.Object)
}

func NewIrt(object rpsl.Object) (*Irt, error) {
	obj := NewIrtUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewIrtUnchecked(object rpsl.Object) Irt {
	return Irt{Object: object}
}
