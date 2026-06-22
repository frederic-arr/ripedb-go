package ripedb_test

import (
	"strings"
	"testing"
	"time"

	"github.com/frederic-arr/ripedb-go/ripedb"
	"github.com/frederic-arr/ripedb-go/ripedb/models"
)

func TestIntegrationLive(t *testing.T) {
	userAgent := "ripedb-go-e2e (https://github.com/frederic-arr/ripedb-go)"
	client, err := ripedb.NewRipeClient(&ripedb.RipeClientOptions{
		UserAgent: &userAgent,
	})

	if err != nil {
		t.Fatalf("unable to create RIPE client: %v", err)
	}

	urls := []string{
		"as-block/AS139610 - AS139610",
		"as-set/AS-CERN",
		"aut-num/AS12333",
		"domain/103.160.in-addr.arpa",
		// filter-set
		// inet-rtr
		"inet6num/2001:1458::/31",
		"inetnum/128.141.0.0 - 128.141.255.255",
		"irt/IRT-CERN-CERT",
		"key-cert/PGPKEY-41B35C52",
		"mntner/CERN-MNT",
		"organisation/ORG-CEOf1-RIPE",
		"peering-set/AS29222:PRNG-CIXP",
		"person/AL6939-RIPE",
		"role/CERN513",
		"route/128.141.0.0/16AS513",
		"route6/2001:1458::/32AS513",
		"route-set/RS-KROOT-CERN",
		// rtr-set
	}

	for _, url := range urls {
		t.Logf("testing %v", url)
		parts := strings.SplitN(url, "/", 2)
		res, err := client.GetResource(parts[0], parts[1])
		if err != nil {
			t.Fatalf("unable to create RIPE client: %v", err)
		}

		err = models.ValidateResourceWithOptions(parts[0], *res, false, []string{"auth"})
		if err != nil {
			t.Fatalf("unable to validate resource %s: %v", url, err)
		}

		time.Sleep(time.Second * 1)
	}
}

// func TestIntegration(t *testing.T) {
// 	files, err := os.ReadDir("tests/data")
// 	if err != nil {
// 		t.Fatalf("unable to read directory: %v\nDid you run scripts/download-dumps.sh?", err)
// 	}

// 	datasets := make([]string, 0, len(files))
// 	for _, file := range files {
// 		datasets = append(datasets, file.Name())
// 	}

// 	for _, dataset := range datasets {
// 		t.Run(dataset, func(t *testing.T) {
// 			t.Parallel()

// 			data, err := os.Open("tests/data/" + dataset)
// 			if err != nil {
// 				t.Fatalf("unable to read file: %v", err)
// 			}
// 			defer func() {
// 				if err := data.Close(); err != nil {
// 					t.Fatalf("failed to close file: %v", err)
// 				}
// 			}()

// 			reader := rpsl.NewReader(data)

// 			const maxWorkers = 8
// 			const channelBuffer = 100
// 			objectsCh := make(chan *rpsl.Object, channelBuffer)
// 			var wg sync.WaitGroup
// 			errCh := make(chan error, maxWorkers)

// 			for w := 0; w < maxWorkers; w++ {
// 				wg.Add(1)
// 				go func() {
// 					defer wg.Done()
// 					for object := range objectsCh {
// 						if err := processObject(object, dataset); err != nil {
// 							errCh <- err
// 						}
// 					}
// 				}()
// 			}

// 			var parseErr error
// 			idx := 0
// 			for {
// 				object, err := reader.Next()
// 				if err != nil {
// 					parseErr = err
// 					break
// 				}
// 				objectsCh <- &object
// 				idx++
// 			}
// 			close(objectsCh)
// 			wg.Wait()
// 			close(errCh)

// 			if err != nil && !errors.Is(err, io.EOF) {
// 				t.Fatalf("reader error in %s: %v", dataset, parseErr)
// 			}

// 			if idx == 0 {
// 				t.Fatalf("no objects parsed from %s, want > 0", dataset)
// 			}

// 			for e := range errCh {
// 				t.Fatal(e)
// 			}
// 		})
// 	}
// }

// func processObject(object *rpsl.Object, dataset string) error {
// 	var err error
// 	if object.Exists("as-block") {
// 		_, err = models.NewAsBlockWithOptions(*object, false, []string{})
// 	} else if object.Exists("as-set") {
// 		_, err = models.NewAsSetWithOptions(*object, false, []string{})
// 	} else if object.Exists("aut-num") {
// 		_, err = models.NewAutNumWithOptions(*object, false, []string{})
// 	} else if object.Exists("domain") {
// 		_, err = models.NewDomainWithOptions(*object, false, []string{})
// 	} else if object.Exists("filter-set") {
// 		_, err = models.NewFilterSetWithOptions(*object, false, []string{})
// 	} else if object.Exists("inet-rtr") {
// 		_, err = models.NewInetRtrWithOptions(*object, false, []string{})
// 	} else if object.Exists("inet6num") {
// 		_, err = models.NewInet6NumWithOptions(*object, false, []string{})
// 	} else if object.Exists("inetnum") {
// 		_, err = models.NewInetNumWithOptions(*object, false, []string{})
// 	} else if object.Exists("irt") {
// 		_, err = models.NewIrtWithOptions(*object, false, []string{"auth"})
// 	} else if object.Exists("key-cert") {
// 		_, err = models.NewKeyCertWithOptions(*object, false, []string{})
// 	} else if object.Exists("mntner") {
// 		_, err = models.NewMntnerWithOptions(*object, false, []string{"auth"})
// 	} else if object.Exists("organisation") {
// 		_, err = models.NewOrganisationWithOptions(*object, false, []string{})
// 	} else if object.Exists("peering-set") {
// 		_, err = models.NewPeeringSetWithOptions(*object, false, []string{})
// 	} else if object.Exists("person") {
// 		_, err = models.NewPersonWithOptions(*object, false, []string{})
// 	} else if object.Exists("role") {
// 		_, err = models.NewRoleWithOptions(*object, false, []string{})
// 	} else if object.Exists("route-set") {
// 		_, err = models.NewRouteSetWithOptions(*object, false, []string{})
// 	} else if object.Exists("route") {
// 		_, err = models.NewRouteWithOptions(*object, false, []string{})
// 	} else if object.Exists("route6") {
// 		_, err = models.NewRoute6WithOptions(*object, false, []string{})
// 	} else if object.Exists("rtr-set") {
// 		_, err = models.NewRtrSetWithOptions(*object, false, []string{})
// 	} else {
// 		return fmt.Errorf(
// 			"unknown object type %v=%v at %s",
// 			object.Attributes[0].Name,
// 			object.Attributes[0].Value,
// 			dataset,
// 		)
// 	}
// 	if err != nil {
// 		return fmt.Errorf(
// 			"error while parsing %v=%v at %s: %w",
// 			object.Attributes[0].Name,
// 			object.Attributes[0].Value,
// 			dataset,
// 			err,
// 		)
// 	}
// 	return nil
// }
