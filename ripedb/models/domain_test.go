// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestDomainSchema(t *testing.T) {
	data := "" +
		"domain:         3.2.1.in-addr.arpa\n" +
		"nserver:        dns.example.com\n" +
		"nserver:        dns2.example.com\n" +
		"admin-c:        JS1-TEST\n" +
		"tech-c:         JS1-TEST\n" +
		"zone-c:         JS1-TEST\n" +
		"mnt-by:         EXAMPLE-MNT\n" +
		"created:        2020-01-04T16:05:06Z\n" +
		"last-modified:  2020-01-04T16:05:06Z\n" +
		"source:         TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewDomain(*object)
	if err != nil {
		t.Fatal(err)
	}
}
