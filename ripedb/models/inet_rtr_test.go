// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestInetRtrSchema(t *testing.T) {
	data := "" +
		"inet-rtr:        abc.example.com\n" +
		"local-as:        AS1234\n" +
		"ifaddr:          127.0.0.1 masklen 0\n" +
		"org:             ORG-ABC-TEST\n" +
		"ifaddr:          192.168.2.1 masklen 0\n" +
		"admin-c:         JS1-TEST\n" +
		"tech-c:          JS1-TEST\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"created:         2021-06-10T20:45:07Z\n" +
		"last-modified:   2021-06-10T20:45:07Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewInetRtr(*object)
	if err != nil {
		t.Fatal(err)
	}
}
