package ripedb

type ObjectMessageArgValueModel struct {
	Value string `json:"value"`
}

type WhoisLinkModel struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

type StatusOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AttributeModel struct {
	Name             string          `json:"name"`
	Value            interface{}     `json:"value,omitempty"`
	Link             *WhoisLinkModel `json:"link,omitempty"`
	ReferencedType   *string         `json:"referenced-type,omitempty"`
	Error            *string         `json:"$$error,omitempty"`
	Info             *string         `json:"$$info,omitempty"`
	Invalid          *bool           `json:"$$invalid,omitempty"`
	ID               *string         `json:"$$id,omitempty"`
	Comment          *string         `json:"comment,omitempty"`
	Success          *string         `json:"$$success,omitempty"`
	StatusOptionList []StatusOption  `json:"$$statusOptionList,omitempty"`
	Hidden           *bool           `json:"$$hidden,omitempty"`
	Disable          *bool           `json:"$$disable,omitempty"`
	HashKey          *string         `json:"$$hashKey,omitempty"`
	Meta             *Meta           `json:"$$meta,omitempty"`
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

type ObjectMessageModel struct {
	Attribute *AttributeModel              `json:"attribute,omitempty"`
	Severity  *string                      `json:"severity,omitempty"`
	Text      *string                      `json:"text,omitempty"`
	Args      []ObjectMessageArgValueModel `json:"args,omitempty"`
	PlainText *string                      `json:"plainText,omitempty"`
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

type WhoisObjectModel struct {
	Type           *string         `json:"type,omitempty"`
	Link           *WhoisLinkModel `json:"link,omitempty"`
	Source         Source          `json:"source"`
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
	Attribute []AttributeModel `json:"attribute"`
}

type Attributes struct {
	Attribute []AttributeModel `json:"attribute"`
}

type ObjectMessages struct {
	ObjectMessage []ObjectMessageModel `json:"objectmessage"`
}

type ResourceHolder struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type WhoisResponseModel struct {
	Link               *WhoisLinkModel `json:"link,omitempty"`
	ErrorMessages      *ErrorMessages  `json:"errormessages,omitempty"`
	Service            *Service        `json:"service,omitempty"`
	Parameters         *Parameters     `json:"parameters,omitempty"`
	Objects            *Objects        `json:"objects,omitempty"`
	TermsAndConditions interface{}     `json:"terms-and-conditions,omitempty"`
	Version            *Version        `json:"version,omitempty"`
}

type ErrorMessages struct {
	ErrorMessage []ObjectMessageModel `json:"errormessage"`
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
	Object []WhoisObjectModel `json:"object"`
}

type MntByModel struct {
	Auth  []string `json:"auth,omitempty"`
	Key   string   `json:"key"`
	Mine  *bool    `json:"mine,omitempty"`
	Type  string   `json:"type"`
	IsNew *bool    `json:"isNew,omitempty"`
}
