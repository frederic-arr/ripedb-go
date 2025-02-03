// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestMntnerSchema(t *testing.T) {
	data := "" +
		"mntner:          TEST-ROOT-MNT\n" +
		"descr:           Mntner for TEST-DBM-MNT\n" +
		"admin-c:         AA1-TEST\n" +
		"tech-c:          AA2-TEST\n" +
		"upd-to:          bitbucket@ripe.net\n" +
		"mnt-nfy:         bitbucket@ripe.net\n" +
		"auth:            MD5-PW# Filtered\n" +
		"notify:          bitbucket@ripe.net\n" +
		"mnt-by:          TEST-ROOT-MNT\n" +
		"remarks:         This is an automatically created object.\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST# Filtered\n"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewMntner(*object)
	if err != nil {
		t.Fatal(err)
	}
}
