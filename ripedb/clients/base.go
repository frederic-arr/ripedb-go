// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeDbClient interface {
	SetEndpoint(endpoint string)
	SetSource(source string)
	SetFilter(filter bool)
	SetFormat(format bool)

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
				for _, arg := range errorMessage.Args {
					msg = strings.Replace(msg, "%v", arg.Value, 1)
				}

				errors = append(errors, msg)
			}
		}
	}

	return errors
}

func parseResponse(resp http.Response) (*models.Resource, error) {
	whoisResponse := &models.Resource{}
	err := json.NewDecoder(resp.Body).Decode(whoisResponse)

	if resp.StatusCode != http.StatusOK || err != nil {
		if err == nil {
			errors := gatherErrors(whoisResponse)
			return nil, fmt.Errorf("ripedb-go request error: %v", errors)
		}

		return nil, err
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
