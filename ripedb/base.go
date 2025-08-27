// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
)

const (
	RIPE_TEST_ENDPOINT_INSECURE = "http://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT          = "https://rest-test.db.ripe.net"
	RIPE_TEST_ENDPOINT_MTLS     = "https://rest-cert-test.db.ripe.net"
	RIPE_PROD_ENDPOINT_INSECURE = "http://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT          = "https://rest.db.ripe.net"
	RIPE_PROD_ENDPOINT_MTLS     = "https://rest-cert.db.ripe.net"
)

func partialToOptions(input *RipeClientOptions, defaultEndpoint string) ripeClientOptions {
	opts := ripeClientOptions{
		Endpoint: defaultEndpoint,
		Filter:   false,
		Format:   true,
		NoError:  false,
		Source:   "ripe",
	}

	partial := RipeClientOptions{}
	if input != nil {
		partial = *input
	}

	if partial.Endpoint != nil {
		opts.Endpoint = *partial.Endpoint
	}

	if partial.Filter != nil {
		opts.Filter = *partial.Filter
	}

	if partial.Format != nil {
		opts.Format = *partial.Format
	}

	if partial.NoError != nil {
		opts.NoError = *partial.NoError
	}

	if partial.Source != nil {
		opts.Source = *partial.Source
	}

	return opts
}

func gatherErrors(whoisResponse *models.Resource) []string {
	errors := []string{}
	if whoisResponse.ErrorMessages != nil {
		for _, errorMessage := range whoisResponse.ErrorMessages.ErrorMessage {
			if errorMessage.Text != nil {
				msg := *errorMessage.Text
				args := make([]interface{}, len(errorMessage.Args))
				for i, arg := range errorMessage.Args {
					args[i] = arg.Value
				}
				msg = fmt.Sprintf(msg, args...)

				errors = append(errors, msg)
			}
		}
	}

	return errors
}

func parseResponse(resp http.Response, noError bool) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if err != nil {
		return nil, err
	}

	if !noError && resp.StatusCode != http.StatusOK {
		errors := gatherErrors(whoisResponse)
		return nil, fmt.Errorf("ripedb-go request error: %v", errors)
	}

	return whoisResponse, nil
}

func Lookup(c RipeClient, resource string, key string) (*models.Resource, error) {
	resp, err := c.GetResource(resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
