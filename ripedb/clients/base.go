// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeClientOptionsPartial struct {
	Endpoint *string
	Format   *bool
	Filter   *bool
	NoError  *bool
	Source   *string
}

type RipeClientOptions struct {
	Endpoint string
	Format   bool
	Filter   bool
	NoError  bool
	Source   string
}

type RipeDbClient interface {
	GetEndpoint() string
	GetSource() string
	GetFilter() bool
	GetFormat() bool
	GetNoError() bool

	SetEndpoint(endpoint string)
	SetSource(source string)
	SetFilter(filter bool)
	SetFormat(format bool)
	SetNoError(noError bool)

	GetObject(resource string, key string) (*rpsl.Object, error)
	CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error)
	UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error)
	DeleteObject(resource string, key string) (*rpsl.Object, error)

	GetResource(resource string, key string) (*models.Resource, error)
	PostResource(resource string, data models.Resource) (*models.Resource, error)
	PutResource(resource string, key string, data models.Resource) (*models.Resource, error)
	DeleteResource(resource string, key string) (*models.Resource, error)
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

func Lookup(c RipeDbClient, resource string, key string) (*models.Resource, error) {
	resp, err := c.GetResource(resource, key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
