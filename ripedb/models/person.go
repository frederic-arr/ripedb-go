// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Person struct {
	Object rpsl.Object
}

func (o *Person) Class() string {
	return "person"
}

func (o *Person) Key() string {
	return *o.Object.GetFirst("nic-hdl")
}

func NewPerson(object rpsl.Object) (*Person, error) {
	schema := `
		person:           mandatory  single     lookup key
		address:          mandatory  multiple
		phone:            mandatory  multiple
		fax-no:           optional   multiple
		e-mail:           optional   multiple   lookup key
		org:              optional   multiple   inverse key
		nic-hdl:          mandatory  single     primary/lookup key
		remarks:          optional   multiple
		notify:           optional   multiple   inverse key
		mnt-by:           mandatory  multiple   inverse key
		mnt-ref:          optional   multiple   inverse key
		created:          generated  single
		last-modified:    generated  single
		source:           mandatory  single
	`

	if err := ensureSchema(schema, "person", &object); err != nil {
		return nil, err
	}

	obj := Person{Object: object}
	return &obj, nil
}
