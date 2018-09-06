//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cosmic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type LdapCreateAccountParams struct {
	p map[string]interface{}
}

func (p *LdapCreateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LdapCreateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *LdapCreateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *LdapCreateAccountParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *LdapCreateAccountParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *LdapCreateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LdapCreateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *LdapCreateAccountParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *LdapCreateAccountParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *LdapCreateAccountParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new LdapCreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLdapCreateAccountParams(accounttype int, username string) *LdapCreateAccountParams {
	p := &LdapCreateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	p.p["username"] = username
	return p
}

// Creates an account from an LDAP user
func (s *AuthenticationService) LdapCreateAccount(p *LdapCreateAccountParams) (*LdapCreateAccountResponse, error) {
	resp, err := s.cs.newRequest("ldapCreateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapCreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LdapCreateAccountResponse struct {
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Cputotal                  int64             `json:"cputotal,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Iptotal                   int64             `json:"iptotal,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Memorytotal               int64             `json:"memorytotal,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Networktotal              int64             `json:"networktotal,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Primarystoragetotal       int64             `json:"primarystoragetotal,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Projecttotal              int64             `json:"projecttotal,omitempty"`
	Receivedbytes             int64             `json:"receivedbytes,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Secondarystoragetotal     int64             `json:"secondarystoragetotal,omitempty"`
	Sentbytes                 int64             `json:"sentbytes,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Snapshottotal             int64             `json:"snapshottotal,omitempty"`
	State                     string            `json:"state,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Templatetotal             int64             `json:"templatetotal,omitempty"`
	User                      []struct {
		Account             string `json:"account,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Created             string `json:"created,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Email               string `json:"email,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Id                  string `json:"id,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Username            string `json:"username,omitempty"`
	} `json:"user,omitempty"`
	Vmavailable     string `json:"vmavailable,omitempty"`
	Vmlimit         string `json:"vmlimit,omitempty"`
	Vmrunning       int    `json:"vmrunning,omitempty"`
	Vmstopped       int    `json:"vmstopped,omitempty"`
	Vmtotal         int64  `json:"vmtotal,omitempty"`
	Volumeavailable string `json:"volumeavailable,omitempty"`
	Volumelimit     string `json:"volumelimit,omitempty"`
	Volumetotal     int64  `json:"volumetotal,omitempty"`
	Vpcavailable    string `json:"vpcavailable,omitempty"`
	Vpclimit        string `json:"vpclimit,omitempty"`
	Vpctotal        int64  `json:"vpctotal,omitempty"`
}

type ListDomainLdapLinkParams struct {
	p map[string]interface{}
}

func (p *ListDomainLdapLinkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *ListDomainLdapLinkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

// You should always use this function to get a new ListDomainLdapLinkParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewListDomainLdapLinkParams(domainid string) *ListDomainLdapLinkParams {
	p := &ListDomainLdapLinkParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	return p
}

// list link of domain to group or OU in ldap
func (s *AuthenticationService) ListDomainLdapLink(p *ListDomainLdapLinkParams) (*ListDomainLdapLinkResponse, error) {
	resp, err := s.cs.newRequest("listDomainLdapLink", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainLdapLinkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDomainLdapLinkResponse struct {
	Count          int               `json:"count"`
	DomainLdapLink []*DomainLdapLink `json:"domainldaplink"`
}

type DomainLdapLink struct {
	Accountid   string `json:"accountid,omitempty"`
	Accounttype int    `json:"accounttype,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Ldapenabled bool   `json:"ldapenabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type LinkDomainToLdapParams struct {
	p map[string]interface{}
}

func (p *LinkDomainToLdapParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["admin"]; found {
		u.Set("admin", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *LinkDomainToLdapParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *LinkDomainToLdapParams) SetAdmin(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["admin"] = v
}

func (p *LinkDomainToLdapParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LinkDomainToLdapParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *LinkDomainToLdapParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

// You should always use this function to get a new LinkDomainToLdapParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLinkDomainToLdapParams(accounttype int, domainid string, name string, authenticationType string) *LinkDomainToLdapParams {
	p := &LinkDomainToLdapParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	p.p["domainid"] = domainid
	p.p["name"] = name
	p.p["type"] = authenticationType
	return p
}

// link an existing cloudstack domain to group or OU in ldap
func (s *AuthenticationService) LinkDomainToLdap(p *LinkDomainToLdapParams) (*LinkDomainToLdapResponse, error) {
	resp, err := s.cs.newRequest("linkDomainToLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LinkDomainToLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LinkDomainToLdapResponse struct {
	Accountid   string `json:"accountid,omitempty"`
	Accounttype int    `json:"accounttype,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Ldapenabled bool   `json:"ldapenabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type AddLdapConfigurationParams struct {
	p map[string]interface{}
}

func (p *AddLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *AddLdapConfigurationParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *AddLdapConfigurationParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

// You should always use this function to get a new AddLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewAddLdapConfigurationParams(hostname string, port int) *AddLdapConfigurationParams {
	p := &AddLdapConfigurationParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["port"] = port
	return p
}

// Add a new Ldap Configuration
func (s *AuthenticationService) AddLdapConfiguration(p *AddLdapConfigurationParams) (*AddLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("addLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddLdapConfigurationResponse struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type DeleteLdapConfigurationParams struct {
	p map[string]interface{}
}

func (p *DeleteLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	return u
}

func (p *DeleteLdapConfigurationParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

// You should always use this function to get a new DeleteLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewDeleteLdapConfigurationParams(hostname string) *DeleteLdapConfigurationParams {
	p := &DeleteLdapConfigurationParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	return p
}

// Remove an Ldap Configuration
func (s *AuthenticationService) DeleteLdapConfiguration(p *DeleteLdapConfigurationParams) (*DeleteLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("deleteLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteLdapConfigurationResponse struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type ListLdapConfigurationsParams struct {
	p map[string]interface{}
}

func (p *ListLdapConfigurationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *ListLdapConfigurationsParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *ListLdapConfigurationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLdapConfigurationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLdapConfigurationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLdapConfigurationsParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

// You should always use this function to get a new ListLdapConfigurationsParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewListLdapConfigurationsParams() *ListLdapConfigurationsParams {
	p := &ListLdapConfigurationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all LDAP configurations
func (s *AuthenticationService) ListLdapConfigurations(p *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error) {
	var r ListLdapConfigurationsResponse
	for page := 2; ; page++ {
		var l ListLdapConfigurationsResponse
		resp, err := s.cs.newRequest("listLdapConfigurations", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.LdapConfigurations = append(r.LdapConfigurations, l.LdapConfigurations...)

		if r.Count == len(r.LdapConfigurations) {
			return &r, nil
		}

		p.SetPagesize(len(l.LdapConfigurations))
		p.SetPage(page)
	}
}

type ListLdapConfigurationsResponse struct {
	Count              int                  `json:"count"`
	LdapConfigurations []*LdapConfiguration `json:"ldapconfiguration"`
}

type LdapConfiguration struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type ImportLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ImportLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	return u
}

func (p *ImportLdapUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ImportLdapUsersParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *ImportLdapUsersParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *ImportLdapUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ImportLdapUsersParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *ImportLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ImportLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ImportLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ImportLdapUsersParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

// You should always use this function to get a new ImportLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewImportLdapUsersParams(accounttype int) *ImportLdapUsersParams {
	p := &ImportLdapUsersParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	return p
}

// Import LDAP users
func (s *AuthenticationService) ImportLdapUsers(p *ImportLdapUsersParams) (*ImportLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("importLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ImportLdapUsersResponse struct {
	Domain    string `json:"domain,omitempty"`
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Principal string `json:"principal,omitempty"`
	Username  string `json:"username,omitempty"`
}

type ListLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ListLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listtype"]; found {
		u.Set("listtype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLdapUsersParams) SetListtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listtype"] = v
}

func (p *ListLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewListLdapUsersParams() *ListLdapUsersParams {
	p := &ListLdapUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all LDAP Users
func (s *AuthenticationService) ListLdapUsers(p *ListLdapUsersParams) (*ListLdapUsersResponse, error) {
	var r ListLdapUsersResponse
	for page := 2; ; page++ {
		var l ListLdapUsersResponse
		resp, err := s.cs.newRequest("listLdapUsers", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.LdapUsers = append(r.LdapUsers, l.LdapUsers...)

		if r.Count == len(r.LdapUsers) {
			return &r, nil
		}

		p.SetPagesize(len(l.LdapUsers))
		p.SetPage(page)
	}
}

type ListLdapUsersResponse struct {
	Count     int         `json:"count"`
	LdapUsers []*LdapUser `json:"ldapuser"`
}

type LdapUser struct {
	Domain    string `json:"domain,omitempty"`
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Principal string `json:"principal,omitempty"`
	Username  string `json:"username,omitempty"`
}

type LoginParams struct {
	p map[string]interface{}
}

func (p *LoginParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LoginParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *LoginParams) SetDomainId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainId"] = v
}

func (p *LoginParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *LoginParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new LoginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLoginParams(password string, username string) *LoginParams {
	p := &LoginParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Logs a user into the CloudStack. A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Login(p *LoginParams) (*LoginResponse, error) {
	resp, err := s.cs.newRequest("login", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LoginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LoginResponse struct {
	Account    string `json:"account,omitempty"`
	Domainid   string `json:"domainid,omitempty"`
	Domainname string `json:"domainname,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	Registered string `json:"registered,omitempty"`
	Sessionkey string `json:"sessionkey,omitempty"`
	Timeout    int    `json:"timeout,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
	Type       string `json:"type,omitempty"`
	Userid     string `json:"userid,omitempty"`
	Username   string `json:"username,omitempty"`
}

type LogoutParams struct {
	p map[string]interface{}
}

func (p *LogoutParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LogoutParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLogoutParams() *LogoutParams {
	p := &LogoutParams{}
	p.p = make(map[string]interface{})
	return p
}

// Logs out the user
func (s *AuthenticationService) Logout(p *LogoutParams) (*LogoutResponse, error) {
	resp, err := s.cs.newRequest("logout", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LogoutResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LogoutResponse struct {
	Description string `json:"description,omitempty"`
}
