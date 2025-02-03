// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeX509Client struct {
	Opts RipeClientOptions
	Key  string
	Cert string
}

func (c *RipeX509Client) SetEndpoint(endpoint string) {
	c.Opts.Endpoint = endpoint
}

func (c *RipeX509Client) SetSource(source string) {
	c.Opts.Source = source
}

func (c *RipeX509Client) SetFilter(filter bool) {
	c.Opts.Filter = filter
}

func (c *RipeX509Client) SetFormat(format bool) {
	c.Opts.Format = format
}

func (c *RipeX509Client) SetNoError(noError bool) {
	c.Opts.NoError = noError
}

func (c *RipeX509Client) request(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
	cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	var path string
	if key == "" {
		path = fmt.Sprintf("%s/%s/%s", c.Opts.Endpoint, c.Opts.Source, resource)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s", c.Opts.Endpoint, c.Opts.Source, resource, url.PathEscape(key))
	}
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

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
	return parseResponse(*resp, c.Opts.NoError)
}

func (c *RipeX509Client) GetObject(resource string, key string) (*rpsl.Object, error) {
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

func (c *RipeX509Client) CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error) {
	data := models.NewResourceFromRpslObject(object)
	res, err := c.PostResource(resource, data)
	if err != nil {
		return nil, err
	}

	obj, err := res.FindOne()
	if err != nil {
		return nil, err
	}

	return models.ModelObjectToRpslObject(obj)
}
func (c *RipeX509Client) UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error) {
	data := models.NewResourceFromRpslObject(object)
	res, err := c.PutResource(resource, key, data)
	if err != nil {
		return nil, err
	}

	obj, err := res.FindOne()
	if err != nil {
		return nil, err
	}

	return models.ModelObjectToRpslObject(obj)
}
func (c *RipeX509Client) DeleteObject(resource string, key string) (*rpsl.Object, error) {
	res, err := c.DeleteResource(resource, key)
	if err != nil {
		return nil, err
	}

	obj, err := res.FindOne()
	if err != nil {
		return nil, err
	}

	return models.ModelObjectToRpslObject(obj)
}

func (c *RipeX509Client) GetResource(resource string, key string) (*models.Resource, error) {
	return c.request("GET", resource, key, nil)
}

func (c *RipeX509Client) PostResource(resource string, data models.Resource) (*models.Resource, error) {
	err := ValidateResource(resource, data)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("POST", resource, "", bytes.NewReader(body))
}

func (c *RipeX509Client) PutResource(resource string, key string, data models.Resource) (*models.Resource, error) {
	err := ValidateResource(resource, data)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("PUT", resource, key, bytes.NewReader(body))
}

func (c *RipeX509Client) DeleteResource(resource string, key string) (*models.Resource, error) {
	return c.request("DELETE", resource, key, nil)
}
