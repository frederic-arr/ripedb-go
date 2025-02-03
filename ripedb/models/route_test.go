// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestRouteSchema(t *testing.T) {
	data := "" +
		"route:           10.11.11.0/24\n" +
		"origin:          AS101111\n" +
		"descr:           route\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewRoute(*object)
	if err != nil {
		t.Fatal(err)
	}
}
