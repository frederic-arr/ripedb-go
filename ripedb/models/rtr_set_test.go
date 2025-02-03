// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestRtrSetSchema(t *testing.T) {
	data := "" +
		"rtr-set:         rtrs-TEST\n" +
		"descr:           TEST NETWORK\n" +
		"mbrs-by-ref:     TEST-MNT\n" +
		"org:             ORG-EXAMPLE-TEST\n" +
		"tech-c:          JS1-TEST\n" +
		"admin-c:         JS1-TEST\n" +
		"notify:          noc@example.com\n" +
		"mnt-by:          TEST-MNT\n" +
		"created:         2019-11-04T00:11:46Z\n" +
		"last-modified:   2024-03-08T09:50:16Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewRtrSet(*object)
	if err != nil {
		t.Fatal(err)
	}
}
