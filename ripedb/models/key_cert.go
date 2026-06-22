// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

var _ Model = KeyCert{}

type KeyCert struct {
	Object rpsl.Object
}

func (o KeyCert) Class() string {
	return "key-cert"
}

func (o KeyCert) Key() string {
	return *o.Object.GetFirst("key-cert")
}

func (o KeyCert) Validate() error {
	return o.ValidateWithOptions(false, make([]string, 0))
}

func (o KeyCert) ValidateWithOptions(skipUnknownKeys bool, skipKeys []string) error {
	schema := `
        key-cert:       mandatory  single     primary/lookup key
        method:         generated  single
        owner:          generated  multiple
        fingerpr:       generated  single     inverse key
        certif:         mandatory  multiple
        org:            optional   multiple   inverse key
        remarks:        optional   multiple
        notify:         optional   multiple   inverse key
        admin-c:        optional   multiple   inverse key
        tech-c:         optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	return ensureSchema(schema, "key-cert", &o.Object, skipUnknownKeys, skipKeys)
}

func NewKeyCert(object rpsl.Object) (*KeyCert, error) {
	return NewKeyCertWithOptions(object, false, make([]string, 0))
}

func NewKeyCertWithOptions(object rpsl.Object, skipUnknownKeys bool, skipKeys []string) (*KeyCert, error) {
	obj := NewKeyCertUnchecked(object)
	if err := obj.ValidateWithOptions(skipUnknownKeys, skipKeys); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewKeyCertUnchecked(object rpsl.Object) KeyCert {
	return KeyCert{Object: object}
}
