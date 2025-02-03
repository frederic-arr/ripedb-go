// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"fmt"
	"strings"

	"github.com/frederic-arr/rpsl-go"
)

type ObjectMessageArgValue struct {
	Value string `json:"value"`
}

type Link struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

type StatusOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Attribute struct {
	Name             string         `json:"name"`
	Value            interface{}    `json:"value,omitempty"`
	Link             *Link          `json:"link,omitempty"`
	ReferencedType   *string        `json:"referenced-type,omitempty"`
	Error            *string        `json:"$$error,omitempty"`
	Info             *string        `json:"$$info,omitempty"`
	Invalid          *bool          `json:"$$invalid,omitempty"`
	ID               *string        `json:"$$id,omitempty"`
	Comment          *string        `json:"comment,omitempty"`
	Success          *string        `json:"$$success,omitempty"`
	StatusOptionList []StatusOption `json:"$$statusOptionList,omitempty"`
	Hidden           *bool          `json:"$$hidden,omitempty"`
	Disable          *bool          `json:"$$disable,omitempty"`
	HashKey          *string        `json:"$$hashKey,omitempty"`
	Meta             *Meta          `json:"$$meta,omitempty"`
}

type Meta struct {
	Idx        *int     `json:"$$idx,omitempty"`
	Mandatory  *bool    `json:"$$mandatory,omitempty"`
	Multiple   *bool    `json:"$$multiple,omitempty"`
	PrimaryKey *bool    `json:"$$primaryKey,omitempty"`
	Refs       []string `json:"$$refs,omitempty"`
	Searchable *bool    `json:"$$searchable,omitempty"`
	IsEnum     *bool    `json:"$$isEnum,omitempty"`
	IsLir      *bool    `json:"$$isLir,omitempty"`
	Disable    *bool    `json:"$$disable,omitempty"`
	Short      *string  `json:"$$short,omitempty"`
}

type ObjectMessage struct {
	Attribute *Attribute              `json:"attribute,omitempty"`
	Severity  *string                 `json:"severity,omitempty"`
	Text      *string                 `json:"text,omitempty"`
	Args      []ObjectMessageArgValue `json:"args,omitempty"`
	PlainText *string                 `json:"plainText,omitempty"`
}

type AbuseCModel struct {
	Key     string  `json:"key"`
	Email   string  `json:"email"`
	Suspect *bool   `json:"suspect,omitempty"`
	OrgID   *string `json:"org-id,omitempty"`
}

type Version struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type Object struct {
	Type           *string         `json:"type,omitempty"`
	Link           *Link           `json:"link,omitempty"`
	Source         *Source         `json:"source,omitempty"`
	PrimaryKey     *PrimaryKey     `json:"primary-key,omitempty"`
	Attributes     Attributes      `json:"attributes"`
	ObjectMessages *ObjectMessages `json:"objectmessages,omitempty"`
	ResourceHolder *ResourceHolder `json:"resource-holder,omitempty"`
	AbuseContact   *AbuseCModel    `json:"abuse-contact,omitempty"`
	Managed        *bool           `json:"managed,omitempty"`
	Version        *Version        `json:"version,omitempty"`
}

type Source struct {
	ID string `json:"id"`
}

type PrimaryKey struct {
	Attribute []Attribute `json:"attribute"`
}

type Attributes struct {
	Attribute []Attribute `json:"attribute"`
}

type ObjectMessages struct {
	ObjectMessage []ObjectMessage `json:"objectmessage"`
}

type ResourceHolder struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Resource struct {
	Link               *Link          `json:"link,omitempty"`
	ErrorMessages      *ErrorMessages `json:"errormessages,omitempty"`
	Service            *Service       `json:"service,omitempty"`
	Parameters         *Parameters    `json:"parameters,omitempty"`
	Objects            *Objects       `json:"objects,omitempty"`
	TermsAndConditions interface{}    `json:"terms-and-conditions,omitempty"`
	Version            *Version       `json:"version,omitempty"`
}

type ErrorMessages struct {
	ErrorMessage []ObjectMessage `json:"errormessage"`
}

type Service struct {
	Name string `json:"name"`
}

type Parameters struct {
	InverseLookup map[string]interface{} `json:"inverse-lookup"`
	TypeFilters   map[string]interface{} `json:"type-filters"`
	Flags         map[string]interface{} `json:"flags"`
	QueryStrings  QueryStrings           `json:"query-strings"`
	Sources       map[string]interface{} `json:"sources"`
}

type QueryStrings struct {
	QueryString []QueryString `json:"query-string"`
}

type QueryString struct {
	Value string `json:"value"`
}

type Objects struct {
	Object []Object `json:"object"`
}

func (m *Resource) FindOne() (*Object, error) {
	if m.Objects == nil || m.Objects.Object == nil || len(m.Objects.Object) == 0 {
		return nil, fmt.Errorf("no objects found")
	}

	if len(m.Objects.Object) > 1 {
		return nil, fmt.Errorf("more than one object found")
	}

	return &m.Objects.Object[0], nil
}

func ModelObjectToRpslObject(m *Object) (*rpsl.Object, error) {
	obj := rpsl.Object{
		Attributes: make([]rpsl.Attribute, len(m.Attributes.Attribute)),
	}

	for i, attr := range m.Attributes.Attribute {
		value, ok := attr.Value.(string)
		if !ok {
			return nil, fmt.Errorf("attribute value is not a string")
		}

		attr := rpsl.Attribute{
			Name:  attr.Name,
			Value: value,
		}

		obj.Attributes[i] = attr
	}

	return &obj, nil
}

func NewObjectFromRpslObject(o *rpsl.Object) Object {
	attributes := make([]Attribute, len(o.Attributes))
	for i, attr := range o.Attributes {
		attributes[i] = Attribute{
			Name:  attr.Name,
			Value: attr.Value,
		}
	}

	return Object{
		Attributes: Attributes{
			Attribute: attributes,
		},
	}
}

func NewResourceFromObject(o *Object) Resource {
	return Resource{
		Objects: &Objects{
			Object: []Object{*o},
		},
	}
}

func NewResourceFromRpslObject(o *rpsl.Object) Resource {
	obj := NewObjectFromRpslObject(o)
	return NewResourceFromObject(&obj)
}

func ensureSchema(schema string, class string, object *rpsl.Object) error {
	err := object.EnsureClass(class)
	if err != nil {
		return err
	}

	schema = strings.TrimSpace(schema)
	lines := strings.Split(schema, "\n")
	keys := make(map[string]bool)
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			return fmt.Errorf("empty line in schema")
		}

		isMandatory := strings.Contains(line, " mandatory ")
		isOptional := strings.Contains(line, " optional ")
		isGenerated := strings.Contains(line, " generated ")
		isSingle := strings.Contains(line, "single")
		isMultiple := strings.Contains(line, " multiple ")

		parts := strings.SplitN(line, ":", 2)
		attr := strings.TrimSpace(parts[0])

		var err error
		if isMandatory && isSingle {
			err = object.EnsureOne(attr)
		} else if isMandatory && isMultiple {
			err = object.EnsureAtLeastOne(attr)
		} else if (isOptional || isGenerated) && isSingle {
			err = object.EnsureAtMostOne(attr)
		}

		keys[attr] = true

		if err != nil {
			return err
		}
	}

	for _, attr := range object.Attributes {
		if _, ok := keys[attr.Name]; !ok {
			return fmt.Errorf("attribute %s not found in schema", attr.Name)
		}
	}

	return nil
}
