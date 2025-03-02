// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Organisation{}

type Organisation struct {
	Object rpsl.Object
}

func (o Organisation) Class() string {
	return "organisation"
}

func (o Organisation) Key() string {
	return *o.Object.GetFirst("organisation")
}

func (o Organisation) Validate() error {
	schema := `
        organisation:     mandatory  single     primary/lookup key
        org-name:         mandatory  single     lookup key
        org-type:         mandatory  single
        descr:            optional   multiple
        remarks:          optional   multiple
        address:          mandatory  multiple
        country:          optional   single
        phone:            optional   multiple
        fax-no:           optional   multiple
        e-mail:           mandatory  multiple   lookup key
        geoloc:           optional   single
        language:         optional   multiple
        org:              optional   multiple   inverse key
        admin-c:          optional   multiple   inverse key
        tech-c:           optional   multiple   inverse key
        abuse-c:          optional   single     inverse key
        ref-nfy:          optional   multiple   inverse key
        mnt-ref:          mandatory  multiple   inverse key
        notify:           optional   multiple   inverse key
        mnt-by:           mandatory  multiple   inverse key
        created:          generated  single
        last-modified:    generated  single
        source:           mandatory  single
	`

	return ensureSchema(schema, "organisation", &o.Object)
}

func NewOrganisation(object rpsl.Object) (*Organisation, error) {
	obj := NewOrganisationUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewOrganisationUnchecked(object rpsl.Object) Organisation {
	return Organisation{Object: object}
}
