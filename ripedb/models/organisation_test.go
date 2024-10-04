package models

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestOrganisationSchema(t *testing.T) {
	data := "" +
		"organisation:    ORG-TT1-TEST\n" +
		"org-name:        ORG TEST\n" +
		"org-type:        RIR\n" +
		"address:         Somewhere in nowhere\n" +
		"phone:           +12 34 567 8900\n" +
		"fax-no:          +12 34 567 8900\n" +
		"e-mail:          bitbucket@ripe.net\n" +
		"admin-c:         AA1-TEST\n" +
		"tech-c:          AA2-TEST\n" +
		"abuse-c:         AA2-TEST\n" +
		"ref-nfy:         bitbucket@ripe.net\n" +
		"mnt-ref:         TEST-DBM-MNT\n" +
		"notify:          bitbucket@ripe.net\n" +
		"mnt-by:          TEST-ROOT-MNT\n" +
		"remarks:         This is an automatically created object.\n" +
		"created:         2002-04-08T12:43:46Z\n" +
		"last-modified:   2014-02-24T13:15:13Z\n" +
		"source:          TEST"

	object, err := rpsl.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewOrganisation(*object)
	if err != nil {
		t.Fatal(err)
	}
}
