// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipePasswordClient struct {
	Endpoint string
	Filter   bool
	Format   bool
	User     *string
	Password string
	Source   string
}

func (c *RipePasswordClient) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

func (c *RipePasswordClient) SetSource(source string) {
	c.Source = source
}

func (c *RipePasswordClient) SetFilter(filter bool) {
	c.Filter = filter
}

func (c *RipePasswordClient) SetFormat(format bool) {
	c.Format = format
}

func (c *RipePasswordClient) request(method string, resource string, key string, body io.Reader) (*models.Resource, error) {
	httpClient := &http.Client{}
	var path string
	if key == "" {
		path = fmt.Sprintf("%s/%s/%s", c.Endpoint, c.Source, resource)
	} else {
		path = fmt.Sprintf("%s/%s/%s/%s", c.Endpoint, c.Source, resource, url.PathEscape(key))
	}
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	// print body
	if body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		fmt.Println(buf.String())
	}

	if c.User != nil {
		req.SetBasicAuth(*c.User, c.Password)
	} else {
		q := req.URL.Query()
		q.Add("password", c.Password)
		req.URL.RawQuery = q.Encode()
	}

	if !c.Format {
		q := req.URL.Query()
		q.Add("unformatted", "")
		req.URL.RawQuery = q.Encode()
	}

	if !c.Filter {
		q := req.URL.Query()
		q.Add("unfiltered", "")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(*resp)
}

func (c *RipePasswordClient) GetObject(resource string, key string) (*rpsl.Object, error) {
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

func (c *RipePasswordClient) CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error) {
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
func (c *RipePasswordClient) UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error) {
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
func (c *RipePasswordClient) DeleteObject(resource string, key string) (*rpsl.Object, error) {
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

func (c *RipePasswordClient) GetResource(resource string, key string) (*models.Resource, error) {
	return c.request("GET", resource, key, nil)
}

func (c *RipePasswordClient) PostResource(resource string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("POST", resource, "", bytes.NewReader(body))
}

func (c *RipePasswordClient) PutResource(resource string, key string, data models.Resource) (*models.Resource, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.request("PUT", resource, key, bytes.NewReader(body))
}

func (c *RipePasswordClient) DeleteResource(resource string, key string) (*models.Resource, error) {
	return c.request("DELETE", resource, key, nil)
}
