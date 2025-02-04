// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/frederic-arr/rpsl-go"
)

type AutNum struct {
	Object rpsl.Object
}

func (o AutNum) Class() string {
	return "aut-num"
}

func (o AutNum) Key() string {
	return *o.Object.GetFirst("aut-num")
}

func NewAutNum(object rpsl.Object) (*AutNum, error) {
	schema := `
        aut-num:         mandatory  single     primary/lookup
        as-name:         mandatory  single
        descr:           optional   multiple
        member-of:       optional   multiple   inverse
        import-via:      optional   multiple
        import:          optional   multiple
        mp-import:       optional   multiple
        export-via:      optional   multiple
        export:          optional   multiple
        mp-export:       optional   multiple
        default:         optional   multiple
        mp-default:      optional   multiple
        remarks:         optional   multiple
        org:             optional   single     inverse
        sponsoring-org:  optional   single     inverse
        admin-c:         mandatory  multiple   inverse
        tech-c:          mandatory  multiple   inverse
        abuse-c:         optional   single     inverse
        status:          generated  single
        notify:          optional   multiple   inverse
        mnt-by:          mandatory  multiple   inverse
        created:         generated  single
        last-modified:   generated  single
        source:          mandatory  single
	`

	if err := ensureSchema(schema, "aut-num", &object); err != nil {
		return nil, err
	}

	obj := AutNum{Object: object}
	return &obj, nil
}
