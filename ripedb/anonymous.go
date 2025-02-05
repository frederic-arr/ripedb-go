// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
)

func newAnonymousClient(opts *RipeClientOptions) (*RipeClient, error) {
	fullOpts := partialToOptions(opts, RIPE_PROD_ENDPOINT)

	request := func(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
		httpClient := &http.Client{}
		req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/%s/%s", fullOpts.Endpoint, fullOpts.Source, resource, url.PathEscape(key)), body)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Accept", "application/json")
		if !fullOpts.Format {
			q := req.URL.Query()
			q.Add("unformatted", "")
			req.URL.RawQuery = q.Encode()
		}

		if !fullOpts.Filter {
			q := req.URL.Query()
			q.Add("unfiltered", "")
			req.URL.RawQuery = q.Encode()
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		return parseResponse(*resp, fullOpts.NoError)
	}

	return &RipeClient{
		opts:    partialToOptions(opts, RIPE_PROD_ENDPOINT_MTLS),
		request: &request,
	}, nil
}
