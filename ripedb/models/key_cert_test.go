// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestKeyCertSchema(t *testing.T) {
	data := "" +
		"key-cert:        PGPKEY-10B066BE\n" +
		"method:          PGP\n" +
		"owner:           Example Person <example@example.net>\n" +
		"fingerpr:        D4A8 4C43 8A5E 3117 6FBA 8159 2649 F26D 10B0 66BE\n" +
		"certif:          -----BEGIN PGP PUBLIC KEY BLOCK-----\n" +
		"certif:          Version: GnuPG/MacGPG2 v2.0.17 (Darwin)\n" +
		"mnt-by:          RIPE-DBM-STARTUP-MNT\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewKeyCert(*object)
	if err != nil {
		t.Fatal(err)
	}
}
