// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = InetRtr{}

type InetRtr struct {
	Object rpsl.Object
}

func (o InetRtr) Class() string {
	return "inet-rtr"
}

func (o InetRtr) Key() string {
	return *o.Object.GetFirst("inet-rtr")
}

func (o InetRtr) Validate() error {
	schema := `
        inet-rtr:       mandatory    single       primary/lookup key
        descr:          optional     multiple
        alias:          optional     multiple
        local-as:       mandatory    single       inverse key
        ifaddr:         mandatory    multiple     inverse key
        interface:      optional     multiple
        peer:           optional     multiple
        mp-peer:        optional     multiple
        member-of:      optional     multiple     inverse key
        remarks:        optional     multiple
        org:            optional     multiple     inverse key
        admin-c:        mandatory    multiple     inverse key
        tech-c:         mandatory    multiple     inverse key
        notify:         optional     multiple     inverse key
        mnt-by:         mandatory    multiple     inverse key
        created:        generated    single
        last-modified:  generated    single
        source:         mandatory    single
	`

	return ensureSchema(schema, "as-block", &o.Object)
}

func NewInetRtr(object rpsl.Object) (*InetRtr, error) {
	obj := NewInetRtrUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewInetRtrUnchecked(object rpsl.Object) InetRtr {
	return InetRtr{Object: object}
}
