// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestRoute6Schema(t *testing.T) {
	data := "" +
		"route6:          2001:67c:2e8::/48\n" +
		"descr:           RIPE-NCC\n" +
		"origin:          AS3333\n" +
		"mnt-lower:       TEST-DBM-MNT\n" +
		"mnt-routes:      TEST-DBM-MNT\n" +
		"mnt-by:          TEST-DBM-MNT\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewRoute6(*object)
	if err != nil {
		t.Fatal(err)
	}
}
