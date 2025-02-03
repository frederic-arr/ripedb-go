// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestAsBlockSchema(t *testing.T) {
	data := "" +
		"as-block:        AS1 - AS4294967295\n" +
		"descr:           Whole as block\n" +
		"org:             ORG-TT1-TEST\n" +
		"mnt-by:          TEST-ROOT-MNT\n" +
		"mnt-lower:       TEST-DBM-MNT\n" +
		"remarks:         This is an automatically created object.\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewAsBlock(*object)
	if err != nil {
		t.Fatal(err)
	}
}
