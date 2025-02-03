// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Mntner struct {
	Object rpsl.Object
}

func (o *Mntner) Class() string {
	return "mntner"
}

func (o *Mntner) Key() string {
	return *o.Object.GetFirst("mntner")
}

func NewMntner(object rpsl.Object) (*Mntner, error) {
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

	if err := ensureSchema(schema, "mntner", &object); err != nil {
		return nil, err
	}

	obj := Mntner{Object: object}
	return &obj, nil
}
