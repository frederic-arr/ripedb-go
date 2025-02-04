// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type InetRtr struct {
	Object rpsl.Object
}

func (o InetRtr) Class() string {
	return "inet-rtr"
}

func (o InetRtr) Key() string {
	return *o.Object.GetFirst("inet-rtr")
}

func NewInetRtr(object rpsl.Object) (*InetRtr, error) {
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

	if err := ensureSchema(schema, "inet-rtr", &object); err != nil {
		return nil, err
	}

	obj := InetRtr{Object: object}
	return &obj, nil
}
