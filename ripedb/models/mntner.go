// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Mntner{}

type Mntner struct {
	Object rpsl.Object
}

func (o Mntner) Class() string {
	return "mntner"
}

func (o Mntner) Key() string {
	return *o.Object.GetFirst("mntner")
}

func (o Mntner) Validate() error {
	schema := `
        mntner:         mandatory  single     primary/lookup key
        descr:          optional   multiple
        org:            optional   multiple   inverse key
        admin-c:        mandatory  multiple   inverse key
        tech-c:         optional   multiple   inverse key
        upd-to:         mandatory  multiple   inverse key
        mnt-nfy:        optional   multiple   inverse key
        auth:           mandatory  multiple   inverse key
        remarks:        optional   multiple
        notify:         optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-ref:        optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "mntner", &o.Object)
}

func NewMntner(object rpsl.Object) (*Mntner, error) {
	obj := NewMntnerUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewMntnerUnchecked(object rpsl.Object) Mntner {
	return Mntner{Object: object}
}
