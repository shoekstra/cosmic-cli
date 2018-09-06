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
	"strings"
)

type CreatePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *CreatePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreatePrivateGatewayParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *CreatePrivateGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreatePrivateGatewayParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *CreatePrivateGatewayParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *CreatePrivateGatewayParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreatePrivateGatewayParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
}

func (p *CreatePrivateGatewayParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

// You should always use this function to get a new CreatePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreatePrivateGatewayParams(ipaddress string, networkid string, vpcid string) *CreatePrivateGatewayParams {
	p := &CreatePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddress"] = ipaddress
	p.p["networkid"] = networkid
	p.p["vpcid"] = vpcid
	return p
}

// Creates a private gateway
func (s *VPCService) CreatePrivateGateway(p *CreatePrivateGatewayParams) (*CreatePrivateGatewayResponse, error) {
	resp, err := s.cs.newRequest("createPrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePrivateGatewayResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreatePrivateGatewayResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Account            string `json:"account,omitempty"`
	Aclid              string `json:"aclid,omitempty"`
	Cidr               string `json:"cidr,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Id                 string `json:"id,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Networkid          string `json:"networkid,omitempty"`
	Networkname        string `json:"networkname,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Sourcenatsupported bool   `json:"sourcenatsupported,omitempty"`
	State              string `json:"state,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
}

type DeletePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *DeletePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePrivateGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeletePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeletePrivateGatewayParams(id string) *DeletePrivateGatewayParams {
	p := &DeletePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Private gateway
func (s *VPCService) DeletePrivateGateway(p *DeletePrivateGatewayParams) (*DeletePrivateGatewayResponse, error) {
	resp, err := s.cs.newRequest("deletePrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePrivateGatewayResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeletePrivateGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListPrivateGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListPrivateGatewaysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListPrivateGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListPrivateGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListPrivateGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPrivateGatewaysParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListPrivateGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListPrivateGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPrivateGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListPrivateGatewaysParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListPrivateGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPrivateGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPrivateGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListPrivateGatewaysParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListPrivateGatewaysParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

// You should always use this function to get a new ListPrivateGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListPrivateGatewaysParams() *ListPrivateGatewaysParams {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetPrivateGatewayByID(id string, opts ...OptionFunc) (*PrivateGateway, int, error) {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPrivateGateways(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.PrivateGateways[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for PrivateGateway UUID: %s!", id)
}

// List private gateways
func (s *VPCService) ListPrivateGateways(p *ListPrivateGatewaysParams) (*ListPrivateGatewaysResponse, error) {
	var r ListPrivateGatewaysResponse
	for page := 2; ; page++ {
		var l ListPrivateGatewaysResponse
		resp, err := s.cs.newRequest("listPrivateGateways", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.PrivateGateways = append(r.PrivateGateways, l.PrivateGateways...)

		if r.Count == len(r.PrivateGateways) {
			return &r, nil
		}

		p.SetPagesize(len(l.PrivateGateways))
		p.SetPage(page)
	}
}

type ListPrivateGatewaysResponse struct {
	Count           int               `json:"count"`
	PrivateGateways []*PrivateGateway `json:"privategateway"`
}

type PrivateGateway struct {
	Account            string `json:"account,omitempty"`
	Aclid              string `json:"aclid,omitempty"`
	Cidr               string `json:"cidr,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Id                 string `json:"id,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Networkid          string `json:"networkid,omitempty"`
	Networkname        string `json:"networkname,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Sourcenatsupported bool   `json:"sourcenatsupported,omitempty"`
	State              string `json:"state,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
}

type CreateStaticRouteParams struct {
	p map[string]interface{}
}

func (p *CreateStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	if v, found := p.p["nexthop"]; found {
		u.Set("nexthop", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreateStaticRouteParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *CreateStaticRouteParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
}

func (p *CreateStaticRouteParams) SetNexthop(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nexthop"] = v
}

func (p *CreateStaticRouteParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

// You should always use this function to get a new CreateStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateStaticRouteParams(cidr string, nexthop string, vpcid string) *CreateStaticRouteParams {
	p := &CreateStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["cidr"] = cidr
	p.p["nexthop"] = nexthop
	p.p["vpcid"] = vpcid
	return p
}

// Creates a static route
func (s *VPCService) CreateStaticRoute(p *CreateStaticRouteParams) (*CreateStaticRouteResponse, error) {
	resp, err := s.cs.newRequest("createStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStaticRouteResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateStaticRouteResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Account   string `json:"account,omitempty"`
	Cidr      string `json:"cidr,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Id        string `json:"id,omitempty"`
	Nexthop   string `json:"nexthop,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	State     string `json:"state,omitempty"`
	Tags      []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcid string `json:"vpcid,omitempty"`
}

type DeleteStaticRouteParams struct {
	p map[string]interface{}
}

func (p *DeleteStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStaticRouteParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteStaticRouteParams(id string) *DeleteStaticRouteParams {
	p := &DeleteStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a static route
func (s *VPCService) DeleteStaticRoute(p *DeleteStaticRouteParams) (*DeleteStaticRouteResponse, error) {
	resp, err := s.cs.newRequest("deleteStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStaticRouteResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteStaticRouteResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListStaticRoutesParams struct {
	p map[string]interface{}
}

func (p *ListStaticRoutesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["nexthop"]; found {
		u.Set("nexthop", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListStaticRoutesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListStaticRoutesParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *ListStaticRoutesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListStaticRoutesParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
}

func (p *ListStaticRoutesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStaticRoutesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListStaticRoutesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStaticRoutesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListStaticRoutesParams) SetNexthop(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nexthop"] = v
}

func (p *ListStaticRoutesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStaticRoutesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStaticRoutesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListStaticRoutesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListStaticRoutesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

// You should always use this function to get a new ListStaticRoutesParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListStaticRoutesParams() *ListStaticRoutesParams {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetStaticRouteByID(id string, opts ...OptionFunc) (*StaticRoute, int, error) {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStaticRoutes(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.StaticRoutes[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for StaticRoute UUID: %s!", id)
}

// Lists all static routes
func (s *VPCService) ListStaticRoutes(p *ListStaticRoutesParams) (*ListStaticRoutesResponse, error) {
	var r ListStaticRoutesResponse
	for page := 2; ; page++ {
		var l ListStaticRoutesResponse
		resp, err := s.cs.newRequest("listStaticRoutes", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.StaticRoutes = append(r.StaticRoutes, l.StaticRoutes...)

		if r.Count == len(r.StaticRoutes) {
			return &r, nil
		}

		p.SetPagesize(len(l.StaticRoutes))
		p.SetPage(page)
	}
}

type ListStaticRoutesResponse struct {
	Count        int            `json:"count"`
	StaticRoutes []*StaticRoute `json:"staticroute"`
}

type StaticRoute struct {
	Account   string `json:"account,omitempty"`
	Cidr      string `json:"cidr,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Id        string `json:"id,omitempty"`
	Nexthop   string `json:"nexthop,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	State     string `json:"state,omitempty"`
	Tags      []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcid string `json:"vpcid,omitempty"`
}

type CreateVPCParams struct {
	p map[string]interface{}
}

func (p *CreateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["sourcenatlist"]; found {
		u.Set("sourcenatlist", v.(string))
	}
	if v, found := p.p["start"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("start", vv)
	}
	if v, found := p.p["syslogserverlist"]; found {
		u.Set("syslogserverlist", v.(string))
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVPCParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateVPCParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *CreateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateVPCParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVPCParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVPCParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *CreateVPCParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateVPCParams) SetSourcenatlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatlist"] = v
}

func (p *CreateVPCParams) SetStart(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["start"] = v
}

func (p *CreateVPCParams) SetSyslogserverlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["syslogserverlist"] = v
}

func (p *CreateVPCParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

func (p *CreateVPCParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new CreateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCParams(cidr string, displaytext string, name string, vpcofferingid string, zoneid string) *CreateVPCParams {
	p := &CreateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["cidr"] = cidr
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["vpcofferingid"] = vpcofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a VPC
func (s *VPCService) CreateVPC(p *CreateVPCParams) (*CreateVPCResponse, error) {
	resp, err := s.cs.newRequest("createVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateVPCResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Account     string `json:"account,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	Created     string `json:"created,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     []struct {
		Account                     string `json:"account,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Id                          string `json:"id,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Ip6gateway                  string `json:"ip6gateway,omitempty"`
		Ipexclusionlist             string `json:"ipexclusionlist,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Ispersistent                bool   `json:"ispersistent,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Name                        string `json:"name,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Service                     []struct {
			Capability []struct {
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
				Value                      string `json:"value,omitempty"`
			} `json:"capability,omitempty"`
			Name     string `json:"name,omitempty"`
			Provider []struct {
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Id                           string   `json:"id,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Specifyipranges  bool   `json:"specifyipranges,omitempty"`
		State            string `json:"state,omitempty"`
		Strechedl2subnet bool   `json:"strechedl2subnet,omitempty"`
		Subdomainaccess  bool   `json:"subdomainaccess,omitempty"`
		Tags             []struct {
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Traffictype       string   `json:"traffictype,omitempty"`
		Type              string   `json:"type,omitempty"`
		Vlan              string   `json:"vlan,omitempty"`
		Vpcid             string   `json:"vpcid,omitempty"`
		Vpcname           string   `json:"vpcname,omitempty"`
		Zoneid            string   `json:"zoneid,omitempty"`
		Zonename          string   `json:"zonename,omitempty"`
		Zonesnetworkspans []string `json:"zonesnetworkspans,omitempty"`
	} `json:"network,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Redundantvpcrouter bool   `json:"redundantvpcrouter,omitempty"`
	Restartrequired    bool   `json:"restartrequired,omitempty"`
	Service            []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Sourcenatlist    string `json:"sourcenatlist,omitempty"`
	State            string `json:"state,omitempty"`
	Syslogserverlist string `json:"syslogserverlist,omitempty"`
	Tags             []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcofferingdisplaytext string `json:"vpcofferingdisplaytext,omitempty"`
	Vpcofferingid          string `json:"vpcofferingid,omitempty"`
	Vpcofferingname        string `json:"vpcofferingname,omitempty"`
	Zoneid                 string `json:"zoneid,omitempty"`
	Zonename               string `json:"zonename,omitempty"`
}

type DeleteVPCParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCParams(id string) *DeleteVPCParams {
	p := &DeleteVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a VPC
func (s *VPCService) DeleteVPC(p *DeleteVPCParams) (*DeleteVPCResponse, error) {
	resp, err := s.cs.newRequest("deleteVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteVPCResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type RestartVPCParams struct {
	p map[string]interface{}
}

func (p *RestartVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RestartVPCParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *RestartVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new RestartVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewRestartVPCParams(id string) *RestartVPCParams {
	p := &RestartVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restarts a VPC
func (s *VPCService) RestartVPC(p *RestartVPCParams) (*RestartVPCResponse, error) {
	resp, err := s.cs.newRequest("restartVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartVPCResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type RestartVPCResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Account     string `json:"account,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	Created     string `json:"created,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     []struct {
		Account                     string `json:"account,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Id                          string `json:"id,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Ip6gateway                  string `json:"ip6gateway,omitempty"`
		Ipexclusionlist             string `json:"ipexclusionlist,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Ispersistent                bool   `json:"ispersistent,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Name                        string `json:"name,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Service                     []struct {
			Capability []struct {
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
				Value                      string `json:"value,omitempty"`
			} `json:"capability,omitempty"`
			Name     string `json:"name,omitempty"`
			Provider []struct {
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Id                           string   `json:"id,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Specifyipranges  bool   `json:"specifyipranges,omitempty"`
		State            string `json:"state,omitempty"`
		Strechedl2subnet bool   `json:"strechedl2subnet,omitempty"`
		Subdomainaccess  bool   `json:"subdomainaccess,omitempty"`
		Tags             []struct {
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Traffictype       string   `json:"traffictype,omitempty"`
		Type              string   `json:"type,omitempty"`
		Vlan              string   `json:"vlan,omitempty"`
		Vpcid             string   `json:"vpcid,omitempty"`
		Vpcname           string   `json:"vpcname,omitempty"`
		Zoneid            string   `json:"zoneid,omitempty"`
		Zonename          string   `json:"zonename,omitempty"`
		Zonesnetworkspans []string `json:"zonesnetworkspans,omitempty"`
	} `json:"network,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Redundantvpcrouter bool   `json:"redundantvpcrouter,omitempty"`
	Restartrequired    bool   `json:"restartrequired,omitempty"`
	Service            []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Sourcenatlist    string `json:"sourcenatlist,omitempty"`
	State            string `json:"state,omitempty"`
	Syslogserverlist string `json:"syslogserverlist,omitempty"`
	Tags             []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcofferingdisplaytext string `json:"vpcofferingdisplaytext,omitempty"`
	Vpcofferingid          string `json:"vpcofferingid,omitempty"`
	Vpcofferingname        string `json:"vpcofferingname,omitempty"`
	Zoneid                 string `json:"zoneid,omitempty"`
	Zonename               string `json:"zonename,omitempty"`
}

type UpdateVPCParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sourcenatlist"]; found {
		u.Set("sourcenatlist", v.(string))
	}
	if v, found := p.p["syslogserverlist"]; found {
		u.Set("syslogserverlist", v.(string))
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	return u
}

func (p *UpdateVPCParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateVPCParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVPCParams) SetSourcenatlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatlist"] = v
}

func (p *UpdateVPCParams) SetSyslogserverlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["syslogserverlist"] = v
}

func (p *UpdateVPCParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

// You should always use this function to get a new UpdateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCParams(id string) *UpdateVPCParams {
	p := &UpdateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a VPC
func (s *VPCService) UpdateVPC(p *UpdateVPCParams) (*UpdateVPCResponse, error) {
	resp, err := s.cs.newRequest("updateVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UpdateVPCResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Account     string `json:"account,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	Created     string `json:"created,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     []struct {
		Account                     string `json:"account,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Id                          string `json:"id,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Ip6gateway                  string `json:"ip6gateway,omitempty"`
		Ipexclusionlist             string `json:"ipexclusionlist,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Ispersistent                bool   `json:"ispersistent,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Name                        string `json:"name,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Service                     []struct {
			Capability []struct {
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
				Value                      string `json:"value,omitempty"`
			} `json:"capability,omitempty"`
			Name     string `json:"name,omitempty"`
			Provider []struct {
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Id                           string   `json:"id,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Specifyipranges  bool   `json:"specifyipranges,omitempty"`
		State            string `json:"state,omitempty"`
		Strechedl2subnet bool   `json:"strechedl2subnet,omitempty"`
		Subdomainaccess  bool   `json:"subdomainaccess,omitempty"`
		Tags             []struct {
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Traffictype       string   `json:"traffictype,omitempty"`
		Type              string   `json:"type,omitempty"`
		Vlan              string   `json:"vlan,omitempty"`
		Vpcid             string   `json:"vpcid,omitempty"`
		Vpcname           string   `json:"vpcname,omitempty"`
		Zoneid            string   `json:"zoneid,omitempty"`
		Zonename          string   `json:"zonename,omitempty"`
		Zonesnetworkspans []string `json:"zonesnetworkspans,omitempty"`
	} `json:"network,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Redundantvpcrouter bool   `json:"redundantvpcrouter,omitempty"`
	Restartrequired    bool   `json:"restartrequired,omitempty"`
	Service            []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Sourcenatlist    string `json:"sourcenatlist,omitempty"`
	State            string `json:"state,omitempty"`
	Syslogserverlist string `json:"syslogserverlist,omitempty"`
	Tags             []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcofferingdisplaytext string `json:"vpcofferingdisplaytext,omitempty"`
	Vpcofferingid          string `json:"vpcofferingid,omitempty"`
	Vpcofferingname        string `json:"vpcofferingname,omitempty"`
	Zoneid                 string `json:"zoneid,omitempty"`
	Zonename               string `json:"zonename,omitempty"`
}

type CreateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["secondaryserviceofferingid"]; found {
		u.Set("secondaryserviceofferingid", v.(string))
	}
	if v, found := p.p["servicecapabilitylist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].key", i), k)
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["serviceproviderlist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].service", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].provider", i), vv)
			i++
		}
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	return u
}

func (p *CreateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVPCOfferingParams) SetSecondaryserviceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secondaryserviceofferingid"] = v
}

func (p *CreateVPCOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicecapabilitylist"] = v
}

func (p *CreateVPCOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateVPCOfferingParams) SetServiceproviderlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceproviderlist"] = v
}

func (p *CreateVPCOfferingParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

// You should always use this function to get a new CreateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCOfferingParams(displaytext string, name string, supportedservices []string) *CreateVPCOfferingParams {
	p := &CreateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["supportedservices"] = supportedservices
	return p
}

// Creates VPC offering
func (s *VPCService) CreateVPCOffering(p *CreateVPCOfferingParams) (*CreateVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("createVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateVPCOfferingResponse struct {
	JobID                        string `json:"jobid,omitempty"`
	Created                      string `json:"created,omitempty"`
	Displaytext                  string `json:"displaytext,omitempty"`
	Id                           string `json:"id,omitempty"`
	Isdefault                    bool   `json:"isdefault,omitempty"`
	Name                         string `json:"name,omitempty"`
	Secondaryserviceofferingid   string `json:"secondaryserviceofferingid,omitempty"`
	Secondaryserviceofferingname string `json:"secondaryserviceofferingname,omitempty"`
	Service                      []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
}

type DeleteVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCOfferingParams(id string) *DeleteVPCOfferingParams {
	p := &DeleteVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes VPC offering
func (s *VPCService) DeleteVPCOffering(p *DeleteVPCOfferingParams) (*DeleteVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteVPCOfferingResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type UpdateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVPCOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

// You should always use this function to get a new UpdateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCOfferingParams(id string) *UpdateVPCOfferingParams {
	p := &UpdateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates VPC offering
func (s *VPCService) UpdateVPCOffering(p *UpdateVPCOfferingParams) (*UpdateVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UpdateVPCOfferingResponse struct {
	JobID                        string `json:"jobid,omitempty"`
	Created                      string `json:"created,omitempty"`
	Displaytext                  string `json:"displaytext,omitempty"`
	Id                           string `json:"id,omitempty"`
	Isdefault                    bool   `json:"isdefault,omitempty"`
	Name                         string `json:"name,omitempty"`
	Secondaryserviceofferingid   string `json:"secondaryserviceofferingid,omitempty"`
	Secondaryserviceofferingname string `json:"secondaryserviceofferingname,omitempty"`
	Service                      []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
}

type ListVPCOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListVPCOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	return u
}

func (p *ListVPCOfferingsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListVPCOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVPCOfferingsParams) SetIsdefault(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdefault"] = v
}

func (p *ListVPCOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVPCOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVPCOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVPCOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVPCOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVPCOfferingsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

// You should always use this function to get a new ListVPCOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCOfferingsParams() *ListVPCOfferingsParams {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVPCOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VPCOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VPCOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingByName(name string, opts ...OptionFunc) (*VPCOffering, int, error) {
	id, count, err := s.GetVPCOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVPCOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingByID(id string, opts ...OptionFunc) (*VPCOffering, int, error) {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVPCOfferings(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.VPCOfferings[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for VPCOffering UUID: %s!", id)
}

// Lists VPC offerings
func (s *VPCService) ListVPCOfferings(p *ListVPCOfferingsParams) (*ListVPCOfferingsResponse, error) {
	var r ListVPCOfferingsResponse
	for page := 2; ; page++ {
		var l ListVPCOfferingsResponse
		resp, err := s.cs.newRequest("listVPCOfferings", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.VPCOfferings = append(r.VPCOfferings, l.VPCOfferings...)

		if r.Count == len(r.VPCOfferings) {
			return &r, nil
		}

		p.SetPagesize(len(l.VPCOfferings))
		p.SetPage(page)
	}
}

type ListVPCOfferingsResponse struct {
	Count        int            `json:"count"`
	VPCOfferings []*VPCOffering `json:"vpcoffering"`
}

type VPCOffering struct {
	Created                      string `json:"created,omitempty"`
	Displaytext                  string `json:"displaytext,omitempty"`
	Id                           string `json:"id,omitempty"`
	Isdefault                    bool   `json:"isdefault,omitempty"`
	Name                         string `json:"name,omitempty"`
	Secondaryserviceofferingid   string `json:"secondaryserviceofferingid,omitempty"`
	Secondaryserviceofferingname string `json:"secondaryserviceofferingname,omitempty"`
	Service                      []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
}

type ListVPCsParams struct {
	p map[string]interface{}
}

func (p *ListVPCsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVPCsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVPCsParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *ListVPCsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListVPCsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVPCsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListVPCsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVPCsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVPCsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVPCsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVPCsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVPCsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVPCsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVPCsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVPCsParams) SetRestartrequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["restartrequired"] = v
}

func (p *ListVPCsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVPCsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *ListVPCsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVPCsParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

func (p *ListVPCsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListVPCsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCsParams() *ListVPCsParams {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVPCs(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VPCs[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VPCs {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCByName(name string, opts ...OptionFunc) (*VPC, int, error) {
	id, count, err := s.GetVPCID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVPCByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCByID(id string, opts ...OptionFunc) (*VPC, int, error) {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVPCs(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.VPCs[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for VPC UUID: %s!", id)
}

// Lists VPCs
func (s *VPCService) ListVPCs(p *ListVPCsParams) (*ListVPCsResponse, error) {
	var r ListVPCsResponse
	for page := 2; ; page++ {
		var l ListVPCsResponse
		resp, err := s.cs.newRequest("listVPCs", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.VPCs = append(r.VPCs, l.VPCs...)

		if r.Count == len(r.VPCs) {
			return &r, nil
		}

		p.SetPagesize(len(l.VPCs))
		p.SetPage(page)
	}
}

type ListVPCsResponse struct {
	Count int    `json:"count"`
	VPCs  []*VPC `json:"vpc"`
}

type VPC struct {
	Account     string `json:"account,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	Created     string `json:"created,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     []struct {
		Account                     string `json:"account,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Id                          string `json:"id,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Ip6gateway                  string `json:"ip6gateway,omitempty"`
		Ipexclusionlist             string `json:"ipexclusionlist,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Ispersistent                bool   `json:"ispersistent,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Name                        string `json:"name,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Service                     []struct {
			Capability []struct {
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
				Value                      string `json:"value,omitempty"`
			} `json:"capability,omitempty"`
			Name     string `json:"name,omitempty"`
			Provider []struct {
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Id                           string   `json:"id,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Specifyipranges  bool   `json:"specifyipranges,omitempty"`
		State            string `json:"state,omitempty"`
		Strechedl2subnet bool   `json:"strechedl2subnet,omitempty"`
		Subdomainaccess  bool   `json:"subdomainaccess,omitempty"`
		Tags             []struct {
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Traffictype       string   `json:"traffictype,omitempty"`
		Type              string   `json:"type,omitempty"`
		Vlan              string   `json:"vlan,omitempty"`
		Vpcid             string   `json:"vpcid,omitempty"`
		Vpcname           string   `json:"vpcname,omitempty"`
		Zoneid            string   `json:"zoneid,omitempty"`
		Zonename          string   `json:"zonename,omitempty"`
		Zonesnetworkspans []string `json:"zonesnetworkspans,omitempty"`
	} `json:"network,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Project            string `json:"project,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Redundantvpcrouter bool   `json:"redundantvpcrouter,omitempty"`
	Restartrequired    bool   `json:"restartrequired,omitempty"`
	Service            []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Sourcenatlist    string `json:"sourcenatlist,omitempty"`
	State            string `json:"state,omitempty"`
	Syslogserverlist string `json:"syslogserverlist,omitempty"`
	Tags             []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Vpcofferingdisplaytext string `json:"vpcofferingdisplaytext,omitempty"`
	Vpcofferingid          string `json:"vpcofferingid,omitempty"`
	Vpcofferingname        string `json:"vpcofferingname,omitempty"`
	Zoneid                 string `json:"zoneid,omitempty"`
	Zonename               string `json:"zonename,omitempty"`
}
