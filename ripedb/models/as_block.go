// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = AsBlock{}

type AsBlock struct {
	Object rpsl.Object
}

func (o AsBlock) Class() string {
	return "as-block"
}

func (o AsBlock) Key() string {
	return *o.Object.GetFirst("as-block")
}

func (o AsBlock) Validate() error {
	schema := `
        as-block:       mandatory  single     primary/lookup
        descr:          optional   multiple
        remarks:        optional   multiple
        org:            optional   multiple   inverse
        notify:         optional   multiple   inverse
        mnt-lower:      optional   multiple   inverse
        mnt-by:         mandatory  multiple   inverse
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "as-block", &o.Object)
}

func NewAsBlock(object rpsl.Object) (*AsBlock, error) {
	obj := NewAsBlockUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewAsBlockUnchecked(object rpsl.Object) AsBlock {
	return AsBlock{Object: object}
}
