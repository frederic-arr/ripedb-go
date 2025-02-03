// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeAnonymousClient struct {
	Opts RipeClientOptions
}

func (c *RipeAnonymousClient) SetEndpoint(endpoint string) {
	c.Opts.Endpoint = endpoint
}

func (c *RipeAnonymousClient) SetSource(source string) {
	c.Opts.Source = source
}

func (c *RipeAnonymousClient) SetFilter(filter bool) {
	c.Opts.Filter = filter
}

func (c *RipeAnonymousClient) SetFormat(format bool) {
	c.Opts.Format = format
}

func (c *RipeAnonymousClient) SetNoError(noError bool) {
	c.Opts.NoError = noError
}

func (c *RipeAnonymousClient) request(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/%s/%s", c.Opts.Endpoint, c.Opts.Source, resource, url.PathEscape(key)), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if !c.Opts.Format {
		q := req.URL.Query()
		q.Add("unformatted", "")
		req.URL.RawQuery = q.Encode()
	}

	if !c.Opts.Filter {
		q := req.URL.Query()
		q.Add("unfiltered", "")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := parseResponse(*resp, c.Opts.NoError)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from URL `%s`: %w", req.URL.String(), err)
	}
	return res, nil
}

func (c *RipeAnonymousClient) GetObject(resource string, key string) (*rpsl.Object, error) {
	res, err := c.GetResource(resource, key)
	if err != nil {
		return nil, err
	}

	obj, err := res.FindOne()
	if err != nil {
		return nil, err
	}

	return models.ModelObjectToRpslObject(obj)
}

func (c *RipeAnonymousClient) CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error) {
	return nil, fmt.Errorf("cannot create resources on anonymous endpoint")
}
func (c *RipeAnonymousClient) UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error) {
	return nil, fmt.Errorf("cannot update resources on anonymous endpoint")
}
func (c *RipeAnonymousClient) DeleteObject(resource string, key string) (*rpsl.Object, error) {
	return nil, fmt.Errorf("cannot delete resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) GetResource(resource string, key string) (*models.Resource, error) {
	return c.request("GET", resource, key, nil)
}

func (c *RipeAnonymousClient) PostResource(resource string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot create resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) PutResource(resource string, key string, data models.Resource) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot update resources on anonymous endpoint")
}

func (c *RipeAnonymousClient) DeleteResource(resource string, key string) (*models.Resource, error) {
	return nil, fmt.Errorf("cannot delete resources on anonymous endpoint")
}
