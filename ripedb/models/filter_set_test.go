// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestFilterSetSchema(t *testing.T) {
	data := "" +
		"filter-set:      AS1234:fltr-filterlist\n" +
		"descr:           Filterlist for AS1234 ingress\n" +
		"filter:          AS1234:fltr-rfc1918 OR fltr-bogons\n" +
		"admin-c:         JS1-TEST\n" +
		"tech-c:          JS1-TEST\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"created:         2008-09-02T00:16:54Z\n" +
		"last-modified:   2021-07-15T08:57:29Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewFilterSet(*object)
	if err != nil {
		t.Fatal(err)
	}
}
