package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type AsBlock struct {
	Object rpsl.Object
}

func (o *AsBlock) Class() string {
	return "as-block"
}

func (o *AsBlock) Key() string {
	return *o.Object.GetFirst("as-block")
}

func NewAsBlock(object rpsl.Object) (*AsBlock, error) {
	schema := `
        as-block:       mandatory  single     primary/lookup
        descr:          optional   multiple
        remarks:        optional   multiple
        org:            optional   multiple   inverse
        notify:         optional   multiple   inverse
        mnt-lower:      optional   multiple   inverse
        mnt-by:         mandatory  multiple   inverse
        created:        generated  single
        last-modified:  generated  single
        source:         mandatory  single
	`

	if err := ensureSchema(schema, "as-block", &object); err != nil {
		return nil, err
	}

	keycert := AsBlock{Object: object}
	return &keycert, nil
}
