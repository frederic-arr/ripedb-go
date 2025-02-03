// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestIrtSchema(t *testing.T) {
	data := "" +
		"irt:             IRT-EXAMPLE-CERT\n" +
		"address:         example\n" +
		"address:         IT department - DI group\n" +
		"address:         CH-0000 Kanton 1\n" +
		"address:         Switzerland\n" +
		"phone:           +41 00 0000 000\n" +
		"fax-no:          +41 00 0000 000\n" +
		"e-mail:          cert@example.ch\n" +
		"admin-c:         EXAMPLE513\n" +
		"tech-c:          EXAMPLE513\n" +
		"remarks:         Emergency telephone number: +41 00 000 00 00\n" +
		"remarks:         Timezone: GMT+1/GMT+2 with DST\n" +
		"remarks:         http://example.ch/security\n" +
		"irt-nfy:         cert@example.ch\n" +
		"notify:          cert@example.ch\n" +
		"notify:          extip@example.ch\n" +
		"mnt-by:          EXAMPLE-MNT\n" +
		"auth:            MD5-PW# Filtered\n" +
		"created:         2011-05-05T08:00:54Z\n" +
		"last-modified:   2015-07-29T15:29:57Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewIrt(*object)
	if err != nil {
		t.Fatal(err)
	}
}
