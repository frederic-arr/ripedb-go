// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Role{}

type Role struct {
	Object rpsl.Object
}

func (o Role) Class() string {
	return "role"
}

func (o Role) Key() string {
	return *o.Object.GetFirst("nic-hdl")
}

func (o Role) Validate() error {
	schema := `
        role:           mandatory  single     lookup key
        address:        mandatory  multiple
        phone:          optional   multiple
        fax-no:         optional   multiple
        e-mail:         mandatory  multiple   lookup key
        org:            optional   multiple   inverse key
        admin-c:        optional   multiple   inverse key
        tech-c:         optional   multiple   inverse key
        nic-hdl:        mandatory  single     primary/lookup key
        remarks:        optional   multiple
        notify:         optional   multiple   inverse key
        abuse-mailbox:  optional   single     inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-ref:        optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "role", &o.Object)
}

func NewRole(object rpsl.Object) (*Role, error) {
	obj := NewRoleUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewRoleUnchecked(object rpsl.Object) Role {
	return Role{Object: object}
}
