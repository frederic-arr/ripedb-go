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
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o Person) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
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

	return ensureSchema(schema, "person", &o.Object, skipUnknownKeys, skipKeys)
}

func NewPerson(object rpsl.Object) (*Person, error) {
	return NewPersonWithOptions(object, false, make([]string, 0))
}

func NewPersonWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*Person, error) {
	obj := NewPersonUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewPersonUnchecked(object rpsl.Object) Person {
	return Person{Object: object}
}
