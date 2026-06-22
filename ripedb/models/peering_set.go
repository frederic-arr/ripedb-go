// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = PeeringSet{}

type PeeringSet struct {
	Object rpsl.Object
}

func (o PeeringSet) Class() string {
	return "peering-set"
}

func (o PeeringSet) Key() string {
	return *o.Object.GetFirst("peering-set")
}

func (o PeeringSet) Validate() error {
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o PeeringSet) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
	schema := `
        peering-set:     mandatory   single     primary/lookup key
        descr:           optional    multiple
        peering:         conditional multiple
        mp-peering:      conditional multiple
        remarks:         optional    multiple
        org:             optional    multiple   inverse key
        tech-c:          mandatory   multiple   inverse key
        admin-c:         mandatory   multiple   inverse key
        notify:          optional    multiple   inverse key
        mnt-by:          mandatory   multiple   inverse key
        mnt-lower:       optional    multiple   inverse key
        created:         generated   single
        last-modified:   generated   single
        source:          mandatory   single
	`

	return ensureSchema(schema, "peering-set", &o.Object, skipUnknownKeys, skipKeys)
}

func NewPeeringSet(object rpsl.Object) (*PeeringSet, error) {
	return NewPeeringSetWithOptions(object, false, make([]string, 0))
}

func NewPeeringSetWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*PeeringSet, error) {
	obj := NewPeeringSetUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewPeeringSetUnchecked(object rpsl.Object) PeeringSet {
	return PeeringSet{Object: object}
}
