package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestPersonSchema(t *testing.T) {
	data := "" +
		"person:          Test Person\n" +
		"mnt-by:          TEST-ROOT-MNT\n" +
		"address:         Somewhere in nowhere\n" +
		"phone:           +12 34 5678900\n" +
		"fax-no:          +12 34 5678900\n" +
		"e-mail:          bitbucket@ripe.net\n" +
		"nic-hdl:         AA1-TEST\n" +
		"remarks:         This is an automatically created object.\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewPerson(*object)
	if err != nil {
		t.Fatal(err)
	}
}
