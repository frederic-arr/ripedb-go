// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"encoding/json"
	"fmt"
	"log/slog"
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
		Endpoint:      defaultEndpoint,
		Filter:        false,
		Format:        true,
		IgnoreError:   false,
		Source:        "ripe",
		UserAgent:     "ripedb-go (https://github.com/frederic-arr/ripedb-go)",
		ExitOnWarning: false,
		ExitOnInfo:    false,
		ExitOnUnknown: false,
		DryRun:        false,
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

	if partial.IgnoreError != nil {
		opts.IgnoreError = *partial.IgnoreError
	}

	if partial.Source != nil {
		opts.Source = *partial.Source
	}

	if partial.UserAgent != nil {
		opts.UserAgent = *partial.UserAgent
	}

	if partial.ExitOnWarning != nil {
		opts.ExitOnWarning = *partial.ExitOnWarning
	}

	if partial.ExitOnInfo != nil {
		opts.ExitOnInfo = *partial.ExitOnInfo
	}

	if partial.ExitOnUnknown != nil {
		opts.ExitOnUnknown = *partial.ExitOnUnknown
	}

	if partial.DryRun != nil {
		opts.DryRun = *partial.DryRun
	}

	if partial.NoColor != nil {
		opts.NoColor = *partial.NoColor
	}

	return opts
}

func gatherErrors(whoisResponse *models.Resource, opts *ripeClientOptions) ([]string, []string, []string, []string) {
	errors := []string{}
	warnings := []string{}
	infos := []string{}
	unknown := []string{}
	if whoisResponse.ErrorMessages != nil {
		for _, errorMessage := range whoisResponse.ErrorMessages.ErrorMessage {
			if errorMessage.Text != nil {
				msg := *errorMessage.Text
				args := make([]interface{}, len(errorMessage.Args))
				for i, arg := range errorMessage.Args {
					args[i] = arg.Value
				}
				msg = fmt.Sprintf(msg, args...)
				if errorMessage.Severity != nil {
					if *errorMessage.Severity == "Info" {
						infos = append(infos, msg)
					} else if *errorMessage.Severity == "Warning" {
						warnings = append(warnings, msg)
					} else if *errorMessage.Severity == "Error" {
						errors = append(errors, msg)
					} else if !opts.ExitOnUnknown {
						unknown = append(unknown, msg)
					}
				} else {
					unknown = append(unknown, msg)
				}
			}
		}
	}

	return errors, warnings, infos, unknown
}

func parseResponse(resp http.Response, opts *ripeClientOptions) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(whoisResponse)
	slog.Debug("HTTP response", "body", string(jsonBytes))

	if err != nil {
		return nil, err
	}

	errors, warnings, infos, unknown := gatherErrors(whoisResponse, opts)

	for _, msg := range infos {
		slog.Info(msg)
	}

	for _, msg := range warnings {
		slog.Warn(msg)
	}

	for _, msg := range errors {
		slog.Error(msg)
	}

	for _, msg := range unknown {
		slog.Info(msg)
	}

	if len(errors) > 0 && !opts.IgnoreError {
		return nil, fmt.Errorf("ripedb-go request error: %v", errors)
	}

	if len(warnings) > 0 && opts.ExitOnWarning {
		return nil, fmt.Errorf("ripedb-go request error: %v", warnings)
	}

	if len(infos) > 0 && opts.ExitOnInfo {
		return nil, fmt.Errorf("ripedb-go request error: %v", infos)
	}

	if len(unknown) > 0 && opts.ExitOnUnknown {
		return nil, fmt.Errorf("ripedb-go request error: %v", unknown)
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
