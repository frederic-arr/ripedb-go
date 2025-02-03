// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestAutNumSchema(t *testing.T) {
	data := "" +
		"aut-num:        AS101111\n" +
		"as-name:        Test-AS-2\n" +
		"descr:          AS for 10.11.11.0 - 10.11.11.255\n" +
		"org:            ORG-TT1-TEST\n" +
		"admin-c:        JS1-TEST\n" +
		"tech-c:         JS1-TEST\n" +
		"mnt-by:         EXAMPLE-MNT\n" +
		"created:        2002-04-08T12:43:46Z\n" +
		"last-modified:  2014-02-24T13:15:13Z\n" +
		"source:         TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewAutNum(*object)
	if err != nil {
		t.Fatal(err)
	}
}
