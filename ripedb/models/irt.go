package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type Irt struct {
	Object rpsl.Object
}

func (o *Irt) Class() string {
	return "irt"
}

func (o *Irt) Key() string {
	return *o.Object.GetFirst("irt")
}

func NewIrt(object rpsl.Object) (*Irt, error) {
	schema := `
        irt:            mandatory  single     primary/lookup key
        address:        mandatory  multiple
        phone:          optional   multiple
        fax-no:         optional   multiple
        e-mail:         mandatory  multiple   lookup key
        signature:      optional   multiple
        encryption:     optional   multiple
        org:            optional   multiple   inverse key
        admin-c:        mandatory  multiple   inverse key
        tech-c:         mandatory  multiple   inverse key
        auth:           mandatory  multiple   inverse key
        remarks:        optional   multiple
        irt-nfy:        optional   multiple   inverse key
        notify:         optional   multiple   inverse key
        mnt-by:         mandatory  multiple   inverse key
        mnt-ref:        optional   multiple   inverse key
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	if err := ensureSchema(schema, "irt", &object); err != nil {
		return nil, err
	}

	keycert := Irt{Object: object}
	return &keycert, nil
}
