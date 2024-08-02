package ripedb

import "errors"

type commonAttributes struct {
	Org          []string
	AdminC       []string
	TechC        []string
	Remarks      []string
	Notify       []string
	MntBy        []string
	Created      *string
	LastModified *string
	Source       string
}

type AsBlockModel struct {
	*commonAttributes
	AsBlock  string
	Descr    []string
	MntLower []string
}

type AsSetModel struct {
	*commonAttributes
	AsSet     string
	Descr     []string
	Members   []string
	MbrsByRef []string
	MntLower  []string
}

type AutNumModel struct {
	*commonAttributes
	AutNum        string
	AsName        string
	Descr         []string
	MemberOf      []string
	ImportVia     []string
	Import        []string
	MpImport      []string
	ExportVia     []string
	Export        []string
	MpExport      []string
	Default       []string
	MpDefault     []string
	SponsoringOrg *string
	AbuseC        *string
	Status        *string
}

type DomainModel struct {
	*commonAttributes
	Domain  string
	Descr   []string
	ZoneC   []string
	NServer []string
	DsRData []string
}

type FilterSetModel struct {
	*commonAttributes
	FilterSet string
	Descr     []string
	filter    *string
	MpFilter  *string
	MntLower  []string
}

type InetRtrModel struct {
	*commonAttributes
	InetRtr   string
	Descr     []string
	Alias     []string
	LocalAs   []string
	IfAddr    []string
	Interface []string
	Peer      []string
	MpPeer    []string
	MemberOf  []string
}

type Inet6NumModel struct {
	*commonAttributes
	Inet6Num       string
	NetName        string
	Descr          []string
	Country        []string
	Geofeed        *string
	Geoloc         *string
	Language       []string
	SponsoringOrg  *string
	AbuseC         *string
	Status         string
	AssignmentSize *string
	MntLower       []string
	MntRoutes      []string
	MntDomains     []string
	MntIrt         []string
}

type InetNumModel struct {
	*commonAttributes
	InetNum        string
	NetName        string
	Descr          []string
	Country        []string
	Geofeed        *string
	Geoloc         *string
	Language       []string
	SponsoringOrg  *string
	AbuseC         *string
	Status         string
	AssignmentSize *string
	MntLower       []string
	MntRoutes      []string
	MntDomains     []string
	MntIrt         []string
}

type IrtModel struct {
	*commonAttributes
	Irt          string
	Address      []string
	Phone        []string
	FaxNo        []string
	Email        []string
	AbuseMailbox []string
	Signature    []string
	Encryption   []string
	Auth         []string
	IrtNfy       []string
	MntRef       []string
}

type KeyCertModel struct {
	*commonAttributes
	KeyCert  string
	Method   *string
	Owner    []string
	Fingerpr *string
	Certif   []string
}

type MntnerModel struct {
	*commonAttributes
	Mntner       string
	Descr        []string
	UpdTo        []string
	MntNfy       []string
	Auth         []string
	AbuseMailbox []string
	MntRef       []string
}

type OrganisationModel struct {
	*commonAttributes
	Organisation string
	OrgName      string
	OrgType      string
	Descr        []string
	Address      []string
	Country      *string
	Phone        []string
	FaxNo        []string
	Email        []string
	Geoloc       *string
	Language     []string
	AbuseC       *string
	RefNfy       []string
	MntRef       []string
	AbuseMailbox []string
}

type PeeringSetModel struct {
	*commonAttributes
	PeeringSet string
	Descr      []string
	Peering    []string
	MpPeering  []string
	MntLower   []string
}

type PersonModel struct {
	*commonAttributes
	Person  string
	Address []string
	Phone   []string
	FaxNo   []string
	Email   []string
	NicHdl  string
	MntRef  []string
}

type RoleModel struct {
	*commonAttributes
	Role         string
	Address      []string
	Phone        []string
	FaxNo        []string
	Email        []string
	NicHdl       string
	AbuseMailbox *string
	MntRef       []string
}

type RouteModel struct {
	*commonAttributes
	Route       string
	Descr       []string
	Origin      string
	Pingable    []string
	PingHdl     []string
	Holes       []string
	MemberOf    []string
	Inject      []string
	AggrMtd     *string
	AggrBndry   *string
	ExportComps *string
	Components  *string
	MntLower    []string
	MntRoutes   []string
}

type RouteSetModel struct {
	*commonAttributes
	RouteSet  string
	Descr     []string
	Members   []string
	MpMembers []string
	MbrsByRef []string
	MntLower  []string
}

type Route6Model struct {
	*commonAttributes
	Route6      string
	Descr       []string
	Origin      string
	Pingable    []string
	PingHdl     []string
	Holes       []string
	MemberOf    []string
	Inject      []string
	AggrMtd     *string
	AggrBndry   *string
	ExportComps *string
	Components  *string
	MntLower    []string
	MntRoutes   []string
}

type RtrSetModel struct {
	*commonAttributes
	RtrSet   string
	Descr    []string
	Members  []string
	MpMembs  []string
	MmbByRef []string
	MntLower []string
}

func commonAttributesFromObject(obj *WhoisObjectModel) *commonAttributes {
	attrs := commonAttributes{}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "org":
			attrs.Org = append(attrs.Org, value)
		case "admin-c":
			attrs.AdminC = append(attrs.AdminC, value)
		case "tech-c":
			attrs.TechC = append(attrs.TechC, value)
		case "remarks":
			attrs.Remarks = append(attrs.Remarks, value)
		case "notify":
			attrs.Notify = append(attrs.Notify, value)
		case "mnt-by":
			attrs.MntBy = append(attrs.MntBy, value)
		case "created":
			attrs.Created = &value
		case "last-modified":
			attrs.LastModified = &value
		case "source":
			attrs.Source = value
		}
	}

	return &attrs
}

func asBlock(obj *WhoisObjectModel) (*AsBlockModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "as-block" {
		return nil, errors.New("object is not an as-block")
	}

	model := AsBlockModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "as-block":
			model.AsBlock = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}

func asSet(obj *WhoisObjectModel) (*AsSetModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "as-set" {
		return nil, errors.New("object is not an as-set")
	}

	model := AsSetModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "as-set":
			model.AsSet = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "members":
			model.Members = append(model.Members, value)
		case "mbrs-by-ref":
			model.MbrsByRef = append(model.MbrsByRef, value)
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}

func autNum(obj *WhoisObjectModel) (*AutNumModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "aut-num" {
		return nil, errors.New("object is not an aut-num")
	}

	model := AutNumModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "aut-num":
			model.AutNum = value
		case "as-name":
			model.AsName = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "member-of":
			model.MemberOf = append(model.MemberOf, value)
		case "import":
			model.Import = append(model.Import, value)
		case "import-via":
			model.ImportVia = append(model.ImportVia, value)
		case "mp-import":
			model.MpImport = append(model.MpImport, value)
		case "export":
			model.Export = append(model.Export, value)
		case "export-via":
			model.ExportVia = append(model.ExportVia, value)
		case "mp-export":
			model.MpExport = append(model.MpExport, value)
		case "default":
			model.Default = append(model.Default, value)
		case "mp-default":
			model.MpDefault = append(model.MpDefault, value)
		case "sponsoring-org":
			model.SponsoringOrg = &value
		case "abuse-c":
			model.AbuseC = &value
		case "status":
			model.Status = &value
		}
	}

	return &model, nil
}

func domain(obj *WhoisObjectModel) (*DomainModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "domain" {
		return nil, errors.New("object is not a domain")
	}

	model := DomainModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "domain":
			model.Domain = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "zone-c":
			model.ZoneC = append(model.ZoneC, value)
		case "nserver":
			model.NServer = append(model.NServer, value)
		case "ds-rdata":
			model.DsRData = append(model.DsRData, value)
		}
	}

	return &model, nil
}

func filterSet(obj *WhoisObjectModel) (*FilterSetModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "filter-set" {
		return nil, errors.New("object is not a filter-set")
	}

	model := FilterSetModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "filter-set":
			model.FilterSet = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "filter":
			model.filter = &value
		case "mp-filter":
			model.MpFilter = &value
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}

func inetRtr(obj *WhoisObjectModel) (*InetRtrModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "inet-rtr" {
		return nil, errors.New("object is not a inet-rtr")
	}

	model := InetRtrModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "inet-rtr":
			model.InetRtr = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "alias":
			model.Alias = append(model.Alias, value)
		case "local-as":
			model.LocalAs = append(model.LocalAs, value)
		case "ifaddr":
			model.IfAddr = append(model.IfAddr, value)
		case "interface":
			model.Interface = append(model.Interface, value)
		case "peer":
			model.Peer = append(model.Peer, value)
		case "mp-peer":
			model.MpPeer = append(model.MpPeer, value)
		case "member-of":
			model.MemberOf = append(model.MemberOf, value)
		}
	}

	return &model, nil
}

func inet6Num(obj *WhoisObjectModel) (*Inet6NumModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "inet6num" {
		return nil, errors.New("object is not a inet6num")
	}

	model := Inet6NumModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "inet6num":
			model.Inet6Num = value
		case "netname":
			model.NetName = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "country":
			model.Country = append(model.Country, value)
		case "geofeed":
			model.Geofeed = &value
		case "geoloc":
			model.Geoloc = &value
		case "language":
			model.Language = append(model.Language, value)
		case "sponsoring-org":
			model.SponsoringOrg = &value
		case "abuse-c":
			model.AbuseC = &value
		case "status":
			model.Status = value
		case "assignment-size":
			model.AssignmentSize = &value
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		case "mnt-routes":
			model.MntRoutes = append(model.MntRoutes, value)
		case "mnt-domains":
			model.MntDomains = append(model.MntDomains, value)
		case "mnt-irt":
			model.MntIrt = append(model.MntIrt, value)
		}
	}

	return &model, nil
}

func inetNum(obj *WhoisObjectModel) (*InetNumModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "inetnum" {
		return nil, errors.New("object is not a inetnum")
	}

	model := InetNumModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "inet6num":
			model.InetNum = value
		case "netname":
			model.NetName = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "country":
			model.Country = append(model.Country, value)
		case "geofeed":
			model.Geofeed = &value
		case "geoloc":
			model.Geoloc = &value
		case "language":
			model.Language = append(model.Language, value)
		case "sponsoring-org":
			model.SponsoringOrg = &value
		case "abuse-c":
			model.AbuseC = &value
		case "status":
			model.Status = value
		case "assignment-size":
			model.AssignmentSize = &value
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		case "mnt-routes":
			model.MntRoutes = append(model.MntRoutes, value)
		case "mnt-domains":
			model.MntDomains = append(model.MntDomains, value)
		case "mnt-irt":
			model.MntIrt = append(model.MntIrt, value)
		}
	}

	return &model, nil
}

func irt(obj *WhoisObjectModel) (*IrtModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "irt" {
		return nil, errors.New("object is not a irt")
	}

	model := IrtModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "irt":
			model.Irt = value
		case "address":
			model.Address = append(model.Address, value)
		case "phone":
			model.Phone = append(model.Phone, value)
		case "fax-no":
			model.FaxNo = append(model.FaxNo, value)
		case "email":
			model.Email = append(model.Email, value)
		case "abuse-mailbox":
			model.AbuseMailbox = append(model.AbuseMailbox, value)
		case "signature":
			model.Signature = append(model.Signature, value)
		case "encryption":
			model.Encryption = append(model.Encryption, value)
		case "auth":
			model.Auth = append(model.Auth, value)
		case "irt-nfy":
			model.IrtNfy = append(model.IrtNfy, value)
		case "mnt-ref":
			model.MntRef = append(model.MntRef, value)
		}
	}

	return &model, nil
}

func keyCert(obj *WhoisObjectModel) (*KeyCertModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "key-cert" {
		return nil, errors.New("object is not a key-cert")
	}

	model := KeyCertModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "key-cert":
			model.KeyCert = value
		case "method":
			model.Method = &value
		case "owner":
			model.Owner = append(model.Owner, value)
		case "fingerpr":
			model.Fingerpr = &value
		case "certif":
			model.Certif = append(model.Certif, value)
		}
	}

	return &model, nil
}

func mntner(obj *WhoisObjectModel) (*MntnerModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "mntner" {
		return nil, errors.New("object is not a mntner")
	}

	model := MntnerModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "mntner":
			model.Mntner = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "upd-to":
			model.UpdTo = append(model.UpdTo, value)
		case "mnt-nfy":
			model.MntNfy = append(model.MntNfy, value)
		case "auth":
			model.Auth = append(model.Auth, value)
		case "abuse-mailbox":
			model.AbuseMailbox = append(model.AbuseMailbox, value)
		case "mnt-ref":
			model.MntRef = append(model.MntRef, value)
		}
	}

	return &model, nil
}

func organisation(obj *WhoisObjectModel) (*OrganisationModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "organisation" {
		return nil, errors.New("object is not a organisation")
	}

	model := OrganisationModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "organisation":
			model.Organisation = value
		case "org-name":
			model.OrgName = value
		case "org-type":
			model.OrgType = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "address":
			model.Address = append(model.Address, value)
		case "country":
			model.Country = &value
		case "phone":
			model.Phone = append(model.Phone, value)
		case "fax-no":
			model.FaxNo = append(model.FaxNo, value)
		case "email":
			model.Email = append(model.Email, value)
		case "geoloc":
			model.Geoloc = &value
		case "language":
			model.Language = append(model.Language, value)
		case "abuse-c":
			model.AbuseC = &value
		case "ref-nfy":
			model.RefNfy = append(model.RefNfy, value)
		case "mnt-ref":
			model.MntRef = append(model.MntRef, value)
		case "abuse-mailbox":
			model.AbuseMailbox = append(model.AbuseMailbox, value)
		}
	}

	return &model, nil
}

func peeringSet(obj *WhoisObjectModel) (*PeeringSetModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "peering-set" {
		return nil, errors.New("object is not a peering-set")
	}

	model := PeeringSetModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "peering-set":
			model.PeeringSet = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "peering":
			model.Peering = append(model.Peering, value)
		case "mp-peering":
			model.MpPeering = append(model.MpPeering, value)
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}

func person(obj *WhoisObjectModel) (*PersonModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "person" {
		return nil, errors.New("object is not a person")
	}

	model := PersonModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "person":
			model.Person = value
		case "address":
			model.Address = append(model.Address, value)
		case "phone":
			model.Phone = append(model.Phone, value)
		case "fax-no":
			model.FaxNo = append(model.FaxNo, value)
		case "email":
			model.Email = append(model.Email, value)
		case "nic-hdl":
			model.NicHdl = value
		case "mnt-ref":
			model.MntRef = append(model.MntRef, value)
		}
	}

	return &model, nil
}

func role(obj *WhoisObjectModel) (*RoleModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "role" {
		return nil, errors.New("object is not a role")
	}

	model := RoleModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "role":
			model.Role = value
		case "address":
			model.Address = append(model.Address, value)
		case "phone":
			model.Phone = append(model.Phone, value)
		case "fax-no":
			model.FaxNo = append(model.FaxNo, value)
		case "email":
			model.Email = append(model.Email, value)
		case "nic-hdl":
			model.NicHdl = value
		case "abuse-mailbox":
			model.AbuseMailbox = &value
		case "mnt-ref":
			model.MntRef = append(model.MntRef, value)
		}
	}

	return &model, nil
}

func route(obj *WhoisObjectModel) (*RouteModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "route" {
		return nil, errors.New("object is not a route")
	}

	model := RouteModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "route":
			model.Route = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "origin":
			model.Origin = value
		case "pingable":
			model.Pingable = append(model.Pingable, value)
		case "ping-hdl":
			model.PingHdl = append(model.PingHdl, value)
		case "holes":
			model.Holes = append(model.Holes, value)
		case "member-of":
			model.MemberOf = append(model.MemberOf, value)
		case "inject":
			model.Inject = append(model.Inject, value)
		case "aggr-mtd":
			model.AggrMtd = &value
		case "aggr-bndry":
			model.AggrBndry = &value
		case "export-comps":
			model.ExportComps = &value
		case "components":
			model.Components = &value
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		case "mnt-routes":
			model.MntRoutes = append(model.MntRoutes, value)
		}
	}

	return &model, nil
}

func routeSet(obj *WhoisObjectModel) (*RouteSetModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "route-set" {
		return nil, errors.New("object is not a route-set")
	}

	model := RouteSetModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "route-set":
			model.RouteSet = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "members":
			model.Members = append(model.Members, value)
		case "mp-members":
			model.MpMembers = append(model.MpMembers, value)
		case "mbrs-by-ref":
			model.MbrsByRef = append(model.MbrsByRef, value)
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}

func route6(obj *WhoisObjectModel) (*Route6Model, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "route6" {
		return nil, errors.New("object is not a route6")
	}

	model := Route6Model{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "route6":
			model.Route6 = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "origin":
			model.Origin = value
		case "pingable":
			model.Pingable = append(model.Pingable, value)
		case "ping-hdl":
			model.PingHdl = append(model.PingHdl, value)
		case "holes":
			model.Holes = append(model.Holes, value)
		case "member-of":
			model.MemberOf = append(model.MemberOf, value)
		case "inject":
			model.Inject = append(model.Inject, value)
		case "aggr-mtd":
			model.AggrMtd = &value
		case "aggr-bndry":
			model.AggrBndry = &value
		case "export-comps":
			model.ExportComps = &value
		case "components":
			model.Components = &value
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		case "mnt-routes":
			model.MntRoutes = append(model.MntRoutes, value)
		}
	}

	return &model, nil
}

func rtrSet(obj *WhoisObjectModel) (*RtrSetModel, error) {
	if obj == nil {
		return nil, errors.New("object is nil")
	}

	if obj.Type == nil || *obj.Type != "rtr-set" {
		return nil, errors.New("object is not a rtr-set")
	}

	model := RtrSetModel{commonAttributes: commonAttributesFromObject(obj)}
	for _, attr := range obj.Attributes.Attribute {
		value := attr.Value.(string)
		switch attr.Name {
		case "rtr-set":
			model.RtrSet = value
		case "descr":
			model.Descr = append(model.Descr, value)
		case "members":
			model.Members = append(model.Members, value)
		case "mp-membs":
			model.MpMembs = append(model.MpMembs, value)
		case "mmb-by-ref":
			model.MmbByRef = append(model.MmbByRef, value)
		case "mnt-lower":
			model.MntLower = append(model.MntLower, value)
		}
	}

	return &model, nil
}
