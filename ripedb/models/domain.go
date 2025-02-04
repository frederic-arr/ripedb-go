// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Domain struct {
	Object rpsl.Object
}

func (o Domain) Class() string {
	return "domain"
}

func (o Domain) Key() string {
	return *o.Object.GetFirst("domain")
}

func NewDomain(object rpsl.Object) (*Domain, error) {
	schema := `
        domain:           mandatory      single       primary/lookup
        descr:            optional       multiple
        org:              optional       multiple     inverse
        admin-c:          mandatory      multiple     inverse
        tech-c:           mandatory      multiple     inverse
        zone-c:           mandatory      multiple     inverse
        nserver:          mandatory      multiple     inverse
        ds-rdata:         optional       multiple     inverse
        remarks:          optional       multiple
        notify:           optional       multiple     inverse
        mnt-by:           mandatory      multiple     inverse
        created:          generated      single
        last-modified:    generated      single
        source:           mandatory      single
	`

	if err := ensureSchema(schema, "domain", &object); err != nil {
		return nil, err
	}

	obj := Domain{Object: object}
	return &obj, nil
}
