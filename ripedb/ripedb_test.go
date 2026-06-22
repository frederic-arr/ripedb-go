// Copyright (c) The RPSL Go Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb_test

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"testing"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

func TestIntegration(t *testing.T) {
	files, err := os.ReadDir("tests/data")
	if err != nil {
		t.Fatalf("unable to read directory: %v\nDid you run scripts/download-dumps.sh?", err)
	}

	datasets := make([]string, 0)
	for _, file := range files {
		datasets = append(datasets, file.Name())
	}

	for _, dataset := range datasets {
		t.Run(dataset, func(t *testing.T) {
			t.Parallel()

			data, err := os.Open("tests/data/" + dataset)
			if err != nil {
				t.Fatalf("unable to read file: %v", err)
			}

			objects, err := rpsl.ParseManyFromReader(data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(objects) == 0 {
				t.Fatalf(`parseObjectsFromBytes => length of %v, want > 0`, len(objects))
			}

			maxProcs := runtime.GOMAXPROCS(0)
			sem := make(chan struct{}, maxProcs)
			var wg sync.WaitGroup
			errCh := make(chan error, len(objects))

			for idx, object := range objects {
				wg.Add(1)

				go func() {
					defer wg.Done()
					sem <- struct{}{}
					defer func() { <-sem }()

					if object.Exists("as-block") {
						_, err = models.NewAsBlockWithOptions(object, false, []string{})
					} else if object.Exists("as-set") {
						_, err = models.NewAsSetWithOptions(object, false, []string{})
					} else if object.Exists("aut-num") {
						_, err = models.NewAutNumWithOptions(object, false, []string{})
					} else if object.Exists("domain") {
						_, err = models.NewDomainWithOptions(object, false, []string{})
					} else if object.Exists("filter-set") {
						_, err = models.NewFilterSetWithOptions(object, false, []string{})
					} else if object.Exists("inet-rtr") {
						_, err = models.NewInetRtrWithOptions(object, false, []string{})
					} else if object.Exists("inet6num") {
						_, err = models.NewInet6NumWithOptions(object, false, []string{})
					} else if object.Exists("inetnum") {
						_, err = models.NewInetNumWithOptions(object, false, []string{})
					} else if object.Exists("irt") {
						_, err = models.NewIrtWithOptions(object, false, []string{"auth"})
					} else if object.Exists("key-cert") {
						_, err = models.NewKeyCertWithOptions(object, false, []string{})
					} else if object.Exists("mntner") {
						_, err = models.NewMntnerWithOptions(object, false, []string{"auth"})
					} else if object.Exists("organisation") {
						_, err = models.NewOrganisationWithOptions(object, false, []string{})
					} else if object.Exists("peering-set") {
						_, err = models.NewPeeringSetWithOptions(object, false, []string{})
					} else if object.Exists("person") {
						_, err = models.NewPersonWithOptions(object, false, []string{})
					} else if object.Exists("role") {
						_, err = models.NewRoleWithOptions(object, false, []string{})
					} else if object.Exists("route-set") {
						_, err = models.NewRouteSetWithOptions(object, false, []string{})
					} else if object.Exists("route") {
						_, err = models.NewRouteWithOptions(object, false, []string{})
					} else if object.Exists("route6") {
						_, err = models.NewRoute6WithOptions(object, false, []string{})
					} else if object.Exists("rtr-set") {
						_, err = models.NewRtrSetWithOptions(object, false, []string{})
					} else {
						errCh <- fmt.Errorf(
							"unknown object type %v=%v at %s:%d",
							object.Attributes[0].Name,
							object.Attributes[0].Value,
							dataset,
							idx,
						)
					}

					if err != nil {
						errCh <- fmt.Errorf(
							"error while parsing %v=%v at %s:%d: %v",
							object.Attributes[0].Name,
							object.Attributes[0].Value,
							dataset,
							idx,
							err,
						)
					}
				}()
			}
		})
	}
}
