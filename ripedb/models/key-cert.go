package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type KeyCert struct {
	Object rpsl.Object
}

func (o *KeyCert) Class() string {
	return "key-cert"
}

func (o *KeyCert) Key() string {
	return *o.Object.GetFirst("key-cert")
}

func NewKeyCert(object rpsl.Object) (*KeyCert, error) {
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

	if err := ensureSchema(schema, "key-cert", &object); err != nil {
		return nil, err
	}

	keycert := KeyCert{Object: object}
	return &keycert, nil
}
