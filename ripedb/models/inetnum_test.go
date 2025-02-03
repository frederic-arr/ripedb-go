// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestInetNumSchema(t *testing.T) {
	data := "" +
		"inetnum:         128.0.0.0 - 128.0.7.255\n" +
		"netname:         NL-RIPENCC-TS-19930901\n" +
		"org:             ORG-EIPB1-TEST\n" +
		"descr:           BA TESTING\n" +
		"country:         NL\n" +
		"admin-c:         AA1-TEST\n" +
		"tech-c:          AA2-TEST\n" +
		"status:          ALLOCATED PA\n" +
		"mnt-by:          TEST-NCC-HM-MNT\n" +
		"mnt-lower:       TEST-DBM-MNT\n" +
		"mnt-routes:      TEST-DBM-MNT\n" +
		"mnt-domains:     TEST-DBM-MNT\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewInetNum(*object)
	if err != nil {
		t.Fatal(err)
	}
}
