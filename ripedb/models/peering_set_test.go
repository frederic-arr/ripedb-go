// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestPeeringSetSchema(t *testing.T) {
	data := "" +
		"peering-set:     AS1234:Test\n" +
		"descr:           Test Peers\n" +
		"remarks:         Test Collector\n" +
		"peering:         AS1234 127.0.0.1 at 127.0.0.1\n" +
		"peering:         AS1234 127.0.0.1 at 127.0.0.1\n" +
		"remarks:         Networks\n" +
		"peering:         AS5678 127.0.0.2 at 127.0.0.2\n" +
		"peering:         AS5678 127.0.0.2 at 127.0.0.2\n" +
		"admin-c:         JS1-TEST\n" +
		"tech-c:          JS1-TEST\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"org:             ORG-ABC-TEST\n" +
		"created:         2008-06-20T11:13:48Z\n" +
		"last-modified:   2008-07-16T00:15:39Z\n" +
		"source:          TEST\n"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewPeeringSet(*object)
	if err != nil {
		t.Fatal(err)
	}
}
