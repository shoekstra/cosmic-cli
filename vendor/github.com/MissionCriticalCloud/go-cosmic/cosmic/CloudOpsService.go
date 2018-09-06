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
	"net/url"
	"strconv"
)

type ListHAWorkersParams struct {
	p map[string]interface{}
}

func (p *ListHAWorkersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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

func (p *ListHAWorkersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListHAWorkersParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListHAWorkersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListHAWorkersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListHAWorkersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListHAWorkersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListHAWorkersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListHAWorkersParams instance,
// as then you are sure you have configured all required params
func (s *CloudOpsService) NewListHAWorkersParams() *ListHAWorkersParams {
	p := &ListHAWorkersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all HA workers
func (s *CloudOpsService) ListHAWorkers(p *ListHAWorkersParams) (*ListHAWorkersResponse, error) {
	var r ListHAWorkersResponse
	for page := 2; ; page++ {
		var l ListHAWorkersResponse
		resp, err := s.cs.newRequest("listHAWorkers", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.HAWorkers = append(r.HAWorkers, l.HAWorkers...)

		if r.Count == len(r.HAWorkers) {
			return &r, nil
		}

		p.SetPagesize(len(l.HAWorkers))
		p.SetPage(page)
	}
}

type ListHAWorkersResponse struct {
	Count     int         `json:"count"`
	HAWorkers []*HAWorker `json:"haworker"`
}

type HAWorker struct {
	Created              string `json:"created,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Domainname           string `json:"domainname,omitempty"`
	Hypervisor           string `json:"hypervisor,omitempty"`
	Id                   int64  `json:"id,omitempty"`
	Managementservername string `json:"managementservername,omitempty"`
	State                string `json:"state,omitempty"`
	Step                 string `json:"step,omitempty"`
	Taken                string `json:"taken,omitempty"`
	Type                 string `json:"type,omitempty"`
	Virtualmachineid     string `json:"virtualmachineid,omitempty"`
	Virtualmachinename   string `json:"virtualmachinename,omitempty"`
	Virtualmachinestate  string `json:"virtualmachinestate,omitempty"`
}

type ListWhoHasThisIpParams struct {
	p map[string]interface{}
}

func (p *ListWhoHasThisIpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["uuid"]; found {
		u.Set("uuid", v.(string))
	}
	return u
}

func (p *ListWhoHasThisIpParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListWhoHasThisIpParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListWhoHasThisIpParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListWhoHasThisIpParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListWhoHasThisIpParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListWhoHasThisIpParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListWhoHasThisIpParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListWhoHasThisIpParams) SetUuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["uuid"] = v
}

// You should always use this function to get a new ListWhoHasThisIpParams instance,
// as then you are sure you have configured all required params
func (s *CloudOpsService) NewListWhoHasThisIpParams(ipaddress string) *ListWhoHasThisIpParams {
	p := &ListWhoHasThisIpParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddress"] = ipaddress
	return p
}

// Lists all for this IP address
func (s *CloudOpsService) ListWhoHasThisIp(p *ListWhoHasThisIpParams) (*ListWhoHasThisIpResponse, error) {
	var r ListWhoHasThisIpResponse
	for page := 2; ; page++ {
		var l ListWhoHasThisIpResponse
		resp, err := s.cs.newRequest("listWhoHasThisIp", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.WhoHasThisIp = append(r.WhoHasThisIp, l.WhoHasThisIp...)

		if r.Count == len(r.WhoHasThisIp) {
			return &r, nil
		}

		p.SetPagesize(len(l.WhoHasThisIp))
		p.SetPage(page)
	}
}

type ListWhoHasThisIpResponse struct {
	Count        int             `json:"count"`
	WhoHasThisIp []*WhoHasThisIp `json:"whohasthisip"`
}

type WhoHasThisIp struct {
	Associatednetworkname string `json:"associatednetworkname,omitempty"`
	Associatednetworkuuid string `json:"associatednetworkuuid,omitempty"`
	Broadcasturi          string `json:"broadcasturi,omitempty"`
	Created               string `json:"created,omitempty"`
	Domainname            string `json:"domainname,omitempty"`
	Domainuuid            string `json:"domainuuid,omitempty"`
	Ipaddress             string `json:"ipaddress,omitempty"`
	Macaddress            string `json:"macaddress,omitempty"`
	Mode                  string `json:"mode,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	Networkname           string `json:"networkname,omitempty"`
	Networkuuid           string `json:"networkuuid,omitempty"`
	State                 string `json:"state,omitempty"`
	Uuid                  string `json:"uuid,omitempty"`
	Virtualmachinename    string `json:"virtualmachinename,omitempty"`
	Virtualmachinetype    string `json:"virtualmachinetype,omitempty"`
	Virtualmachineuuid    string `json:"virtualmachineuuid,omitempty"`
	Vpcname               string `json:"vpcname,omitempty"`
	Vpcuuid               string `json:"vpcuuid,omitempty"`
}

type ListWhoHasThisMacParams struct {
	p map[string]interface{}
}

func (p *ListWhoHasThisMacParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["macaddress"]; found {
		u.Set("macaddress", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["uuid"]; found {
		u.Set("uuid", v.(string))
	}
	return u
}

func (p *ListWhoHasThisMacParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListWhoHasThisMacParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListWhoHasThisMacParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListWhoHasThisMacParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListWhoHasThisMacParams) SetMacaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["macaddress"] = v
}

func (p *ListWhoHasThisMacParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListWhoHasThisMacParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListWhoHasThisMacParams) SetUuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["uuid"] = v
}

// You should always use this function to get a new ListWhoHasThisMacParams instance,
// as then you are sure you have configured all required params
func (s *CloudOpsService) NewListWhoHasThisMacParams() *ListWhoHasThisMacParams {
	p := &ListWhoHasThisMacParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all for this MAC address
func (s *CloudOpsService) ListWhoHasThisMac(p *ListWhoHasThisMacParams) (*ListWhoHasThisMacResponse, error) {
	var r ListWhoHasThisMacResponse
	for page := 2; ; page++ {
		var l ListWhoHasThisMacResponse
		resp, err := s.cs.newRequest("listWhoHasThisMac", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.WhoHasThisMac = append(r.WhoHasThisMac, l.WhoHasThisMac...)

		if r.Count == len(r.WhoHasThisMac) {
			return &r, nil
		}

		p.SetPagesize(len(l.WhoHasThisMac))
		p.SetPage(page)
	}
}

type ListWhoHasThisMacResponse struct {
	Count         int              `json:"count"`
	WhoHasThisMac []*WhoHasThisMac `json:"whohasthismac"`
}

type WhoHasThisMac struct {
	Associatednetworkname string `json:"associatednetworkname,omitempty"`
	Associatednetworkuuid string `json:"associatednetworkuuid,omitempty"`
	Broadcasturi          string `json:"broadcasturi,omitempty"`
	Created               string `json:"created,omitempty"`
	Domainname            string `json:"domainname,omitempty"`
	Domainuuid            string `json:"domainuuid,omitempty"`
	Ipaddress             string `json:"ipaddress,omitempty"`
	Macaddress            string `json:"macaddress,omitempty"`
	Mode                  string `json:"mode,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	Networkname           string `json:"networkname,omitempty"`
	Networkuuid           string `json:"networkuuid,omitempty"`
	State                 string `json:"state,omitempty"`
	Uuid                  string `json:"uuid,omitempty"`
	Virtualmachinename    string `json:"virtualmachinename,omitempty"`
	Virtualmachinetype    string `json:"virtualmachinetype,omitempty"`
	Virtualmachineuuid    string `json:"virtualmachineuuid,omitempty"`
	Vpcname               string `json:"vpcname,omitempty"`
	Vpcuuid               string `json:"vpcuuid,omitempty"`
}
