// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestRouteSetSchema(t *testing.T) {
	data := "" +
		"route-set:       RS-EXAMPLE2\n" +
		"descr:           community for routing domain EXAMPLE2\n" +
		"mbrs-by-ref:     EXAMPLE-MNT\n" +
		"tech-c:          JS1-TEST\n" +
		"admin-c:         JS1-TEST\n" +
		"notify:          abc@example.com\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"created:         1970-01-01T00:00:00Z\n" +
		"last-modified:   2001-09-22T09:34:03Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewRouteSet(*object)
	if err != nil {
		t.Fatal(err)
	}
}
