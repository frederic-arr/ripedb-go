// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type PeeringSet struct {
	Object rpsl.Object
}

func (o *PeeringSet) Class() string {
	return "peering-set"
}

func (o *PeeringSet) Key() string {
	return *o.Object.GetFirst("peering-set")
}

func NewPeeringSet(object rpsl.Object) (*PeeringSet, error) {
	schema := `
        peering-set:     mandatory  single     primary/lookup key
        descr:           optional   multiple
        peering:         optional   multiple
        mp-peering:      optional   multiple
        remarks:         optional   multiple
        org:             optional   multiple   inverse key
        tech-c:          mandatory  multiple   inverse key
        admin-c:         mandatory  multiple   inverse key
        notify:          optional   multiple   inverse key
        mnt-by:          mandatory  multiple   inverse key
        mnt-lower:       optional   multiple   inverse key
        created:         generated  single
        last-modified:   generated  single
        source:          mandatory  single
	`

	if err := ensureSchema(schema, "peering-set", &object); err != nil {
		return nil, err
	}

	obj := PeeringSet{Object: object}
	return &obj, nil
}
