// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
)

func newApiKeyClient(opts *RipeClientOptions) (*RipeClient, error) {
	fullOpts := partialToOptions(opts, RIPE_PROD_ENDPOINT)

	request := func(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
		httpClient := &http.Client{}
		var path string
		if key == "" {
			path = fmt.Sprintf("%s/%s/%s", fullOpts.Endpoint, fullOpts.Source, resource)
		} else {
			path = fmt.Sprintf("%s/%s/%s/%s", fullOpts.Endpoint, fullOpts.Source, resource, url.PathEscape(key))
		}
		req, err := http.NewRequest(method, path, body)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", fullOpts.UserAgent)
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", *opts.ApiKey))
		req.Header.Set("Accept", "application/json")

		if method == "POST" || method == "PUT" {
			req.Header.Set("Content-Type", "application/json")
		}

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

		defer func() {
			if err := resp.Body.Close(); err != nil {
				slog.Error("failed to close HTTP client", "error", err)
			}
		}()
		return parseResponse(*resp, fullOpts.NoError)
	}

	return &RipeClient{
		opts:    partialToOptions(opts, RIPE_PROD_ENDPOINT_MTLS),
		request: &request,
	}, nil
}
