package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Role struct {
	Object rpsl.Object
}

func (o *Role) Class() string {
	return "role"
}

func (o *Role) Key() string {
	return *o.Object.GetFirst("nic-hdl")
}

func NewRole(object rpsl.Object) (*Role, error) {
	schema := `
        role:           mandatory  single     lookup key
        address:        mandatory  multiple
        phone:          optional   multiple
        fax-no:         optional   multiple
        e-mail:         mandatory  multiple   lookup key
        org:            optional   multiple   inverse key
        admin-c:        optional   multiple   inverse key
        tech-c:         optional   multiple   inverse key
        nic-hdl:        mandatory  single     primary/lookup key
        remarks:        optional   multiple
        notify:         optional   multiple   inverse key
        abuse-mailbox:  optional   single     inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-ref:        optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	if err := ensureSchema(schema, "role", &object); err != nil {
		return nil, err
	}

	role := Role{Object: object}
	return &role, nil
}
