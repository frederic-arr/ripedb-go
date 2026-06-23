// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type RipeClientOptions struct {
	// The endpoint of the RIPE Database RESTful API.
	Endpoint *string

	// The database where the queries should be made. This is equivalent to the `source` field of the objects.
	Source *string

	Format *bool

	UserAgent *string

	// A filtering process restricts some data from the default query response. This applies to email contact data.
	// By default, the filter is disabled, you can enable it with this flag
	Filter *bool

	// If set to `true`, L7 errors (HTTP or RIPE) are not handled
	IgnoreError *bool

	// API key for the basic authentication protocol.
	// You cannot use API key Authentication along with any other authentication protocol.
	ApiKey *string

	// PEM-encoded client certificate for TLS authentication. Both `certificate` and `key` must be
	// provided. The `endpoint` field must be set appropriately if you are not using the default production
	// API. You cannot use X.509 Client Authentication along with any other authentication protocol.
	Certificate *[]byte

	// PEM-encoded client certificate key for TLS authentication. Both `certificate` and `key` must be
	// provided. The `endpoint` field must be set appropriately if you are not using the default production
	// API. You cannot use X.509 Client Authentication and Basic authentication at the same time.
	Key *[]byte

	// Exit with an error if any warning message is encountered.
	ExitOnWarning *bool

	// Exit with an error if any info message is encountered.
	ExitOnInfo *bool

	// Exit with an error if any unknown severity message is encountered.
	ExitOnUnknown *bool

	// Performs a dry run.
	DryRun *bool

	// Disables color in the terminal
	NoColor *bool
}

type ripeClientOptions struct {
	Endpoint      string
	Format        bool
	Filter        bool
	IgnoreError   bool
	Source        string
	UserAgent     string
	ExitOnWarning bool
	ExitOnInfo    bool
	ExitOnUnknown bool
	DryRun        bool
	NoColor       bool
}

type RipeClient struct {
	opts    ripeClientOptions
	request *func(method string, resource string, key string, body io.Reader) (*models.Resource, error)
}

func (c *RipeClient) GetEndpoint() string {
	return c.opts.Endpoint
}

func (c *RipeClient) GetSource() string {
	return c.opts.Source
}

func (c *RipeClient) GetFilter() bool {
	return c.opts.Filter
}

func (c *RipeClient) GetFormat() bool {
	return c.opts.Format
}

func (c *RipeClient) GetIgnoreError() bool {
	return c.opts.IgnoreError
}

func (c *RipeClient) GetUserAgent() string {
	return c.opts.UserAgent
}

func (c *RipeClient) GetExitOnWarning() bool {
	return c.opts.ExitOnWarning
}

func (c *RipeClient) GetExitOnInfo() bool {
	return c.opts.ExitOnInfo
}

func (c *RipeClient) GetExitOnUnknown() bool {
	return c.opts.ExitOnUnknown
}

func (c *RipeClient) GetDryRun() bool {
	return c.opts.DryRun
}

func (c *RipeClient) GetNoColor() bool {
	return c.opts.NoColor
}

func (c *RipeClient) SetEndpoint(endpoint string) {
	c.opts.Endpoint = endpoint
}

func (c *RipeClient) SetSource(source string) {
	c.opts.Source = source
}

func (c *RipeClient) SetFilter(filter bool) {
	c.opts.Filter = filter
}

func (c *RipeClient) SetFormat(format bool) {
	c.opts.Format = format
}

func (c *RipeClient) SetignoreError(ignoreError bool) {
	c.opts.IgnoreError = ignoreError
}

func (c *RipeClient) SetUserAgent(userAgent string) {
	c.opts.UserAgent = userAgent
}

func (c *RipeClient) SetExitOnWarning(exitOnWarning bool) {
	c.opts.ExitOnWarning = exitOnWarning
}

func (c *RipeClient) SetExitOnInfo(exitOnInfo bool) {
	c.opts.ExitOnInfo = exitOnInfo
}

func (c *RipeClient) SetExitOnUnknown(exitOnUnknown bool) {
	c.opts.ExitOnUnknown = exitOnUnknown
}

func (c *RipeClient) SetDryRun(dryRun bool) {
	c.opts.DryRun = dryRun
}

func (c *RipeClient) SetNoColor(noColor bool) {
	c.opts.NoColor = noColor
}

func (c *RipeClient) GetObject(resource string, key string) (*rpsl.Object, error) {
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

func (c *RipeClient) CreateObject(resource string, object *rpsl.Object) (*rpsl.Object, error) {
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
func (c *RipeClient) UpdateObject(resource string, key string, object *rpsl.Object) (*rpsl.Object, error) {
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
func (c *RipeClient) DeleteObject(resource string, key string) (*rpsl.Object, error) {
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

func (c *RipeClient) GetResource(resource string, key string) (*models.Resource, error) {
	return (*c.request)("GET", resource, key, nil)
}

func (c *RipeClient) PostResource(resource string, data models.Resource) (*models.Resource, error) {
	timestampKeys := []string{"created", "last-modified", "dry-run"}
	data.RemoveKeys(timestampKeys)

	err := models.ValidateResourceWithOptions(resource, data, false, timestampKeys)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return (*c.request)("POST", resource, "", bytes.NewReader(body))
}

func (c *RipeClient) PutResource(resource string, key string, data models.Resource) (*models.Resource, error) {
	timestampKeys := []string{"created", "last-modified", "dry-run"}
	data.RemoveKeys(timestampKeys)

	err := models.ValidateResourceWithOptions(resource, data, false, timestampKeys)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return (*c.request)("PUT", resource, key, bytes.NewReader(body))
}

func (c *RipeClient) DeleteResource(resource string, key string) (*models.Resource, error) {
	return (*c.request)("DELETE", resource, key, nil)
}
