// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = InetNum{}

type InetNum struct {
	Object rpsl.Object
}

func (o InetNum) Class() string {
	return "inetnum"
}

func (o InetNum) Key() string {
	return *o.Object.GetFirst("inetnum")
}

func (o InetNum) Validate() error {
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o InetNum) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
	schema := `
        inetnum:          mandatory  single     primary/lookup key
        netname:          mandatory  single     lookup key
        descr:            optional   multiple
        country:          mandatory  multiple
        geofeed:          optional   single
        geoloc:           optional   single
        language:         optional   multiple
        org:              optional   single     inverse key
        sponsoring-org:   optional   single
        admin-c:          mandatory  multiple   inverse key
        tech-c:           mandatory  multiple   inverse key
        abuse-c:          optional   single     inverse key
        status:           mandatory  single
        assignment-size:  optional   single
        remarks:          optional   multiple
        notify:           optional   multiple   inverse key
        mnt-by:           mandatory  multiple   inverse key
        mnt-lower:        optional   multiple   inverse key
        mnt-routes:       optional   multiple   inverse key
        mnt-domains:      optional   multiple   inverse key
        mnt-irt:          optional   multiple   inverse key
        created:          generated  single
        last-modified:    generated  single
        source:           mandatory  single
	`

	return ensureSchema(schema, "inetnum", &o.Object, skipUnknownKeys, skipKeys)
}

func NewInetNum(object rpsl.Object) (*InetNum, error) {
	return NewInetNumWithOptions(object, false, make([]string, 0))
}

func NewInetNumWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*InetNum, error) {
	obj := NewInetNumUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewInetNumUnchecked(object rpsl.Object) InetNum {
	return InetNum{Object: object}
}
