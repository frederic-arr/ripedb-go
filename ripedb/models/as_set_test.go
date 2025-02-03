// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestAsSetSchema(t *testing.T) {
	data := "" +
		"as-set:         AS1234:AS-TEST\n" +
		"admin-c:        AA1-TEST\n" +
		"tech-c:         AA1-TEST\n" +
		"mnt-by:         TEST-DBM-MNT\n" +
		"created:        2025-02-03T03:00:25Z\n" +
		"last-modified:  2025-02-03T03:00:25Z\n" +
		"remarks:        Managed by RIPE automation\n" +
		"remarks:        Production Network\n" +
		"members:        AS1234\n" +
		"members:        AS5678:AS-TEST\n" +
		"source:         TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewAsSet(*object)
	if err != nil {
		t.Fatal(err)
	}
}
