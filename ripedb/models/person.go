// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = Person{}

type Person struct {
	Object rpsl.Object
}

func (o Person) Class() string {
	return "person"
}

func (o Person) Key() string {
	return *o.Object.GetFirst("nic-hdl")
}

func (o Person) Validate() error {
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

	return ensureSchema(schema, "person", &o.Object)
}

func NewPerson(object rpsl.Object) (*Person, error) {
	obj := NewPersonUnchecked(object)
	if err := obj.Validate(); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewPersonUnchecked(object rpsl.Object) Person {
	return Person{Object: object}
}
