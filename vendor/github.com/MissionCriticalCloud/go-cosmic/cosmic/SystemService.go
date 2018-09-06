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

type ListApisParams struct {
	p map[string]interface{}
}

func (p *ListApisParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *ListApisParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

// You should always use this function to get a new ListApisParams instance,
// as then you are sure you have configured all required params
func (s *SystemService) NewListApisParams() *ListApisParams {
	p := &ListApisParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists all available apis on the server, provided by the Api Discovery plugin
func (s *SystemService) ListApis(p *ListApisParams) (*ListApisResponse, error) {
	resp, err := s.cs.newRequest("listApis", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListApisResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListApisResponse struct {
	Count int    `json:"count"`
	Apis  []*Api `json:"api"`
}

type Api struct {
	Description      string `json:"description,omitempty"`
	Groupdescription string `json:"groupdescription,omitempty"`
	Groupname        string `json:"groupname,omitempty"`
	Isasync          bool   `json:"isasync,omitempty"`
	Name             string `json:"name,omitempty"`
	Params           []struct {
		Description string `json:"description,omitempty"`
		Length      int    `json:"length,omitempty"`
		Name        string `json:"name,omitempty"`
		Related     string `json:"related,omitempty"`
		Required    bool   `json:"required,omitempty"`
		Since       string `json:"since,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"params,omitempty"`
	Related  string `json:"related,omitempty"`
	Response []struct {
		Description string   `json:"description,omitempty"`
		Name        string   `json:"name,omitempty"`
		Response    []string `json:"response,omitempty"`
		Type        string   `json:"type,omitempty"`
	} `json:"response,omitempty"`
	Since string `json:"since,omitempty"`
	Type  string `json:"type,omitempty"`
}

type ListCapacityParams struct {
	p map[string]interface{}
}

func (p *ListCapacityParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["fetchlatest"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchlatest", vv)
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["sortby"]; found {
		u.Set("sortby", v.(string))
	}
	if v, found := p.p["type"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("type", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListCapacityParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListCapacityParams) SetFetchlatest(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fetchlatest"] = v
}

func (p *ListCapacityParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListCapacityParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListCapacityParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListCapacityParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListCapacityParams) SetSortby(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortby"] = v
}

func (p *ListCapacityParams) SetType(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListCapacityParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListCapacityParams instance,
// as then you are sure you have configured all required params
func (s *SystemService) NewListCapacityParams() *ListCapacityParams {
	p := &ListCapacityParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all the system wide capacities.
func (s *SystemService) ListCapacity(p *ListCapacityParams) (*ListCapacityResponse, error) {
	var r ListCapacityResponse
	for page := 2; ; page++ {
		var l ListCapacityResponse
		resp, err := s.cs.newRequest("listCapacity", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.Capacity = append(r.Capacity, l.Capacity...)

		if r.Count == len(r.Capacity) {
			return &r, nil
		}

		p.SetPagesize(len(l.Capacity))
		p.SetPage(page)
	}
}

type ListCapacityResponse struct {
	Count    int         `json:"count"`
	Capacity []*Capacity `json:"capacity"`
}

type Capacity struct {
	Capacitytotal       int64  `json:"capacitytotal,omitempty"`
	Capacityused        int64  `json:"capacityused,omitempty"`
	Clusterid           string `json:"clusterid,omitempty"`
	Clustername         string `json:"clustername,omitempty"`
	Percentageallocated string `json:"percentageallocated,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Podname             string `json:"podname,omitempty"`
	Type                int    `json:"type,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type GetCloudIdentifierParams struct {
	p map[string]interface{}
}

func (p *GetCloudIdentifierParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *GetCloudIdentifierParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

// You should always use this function to get a new GetCloudIdentifierParams instance,
// as then you are sure you have configured all required params
func (s *SystemService) NewGetCloudIdentifierParams(userid string) *GetCloudIdentifierParams {
	p := &GetCloudIdentifierParams{}
	p.p = make(map[string]interface{})
	p.p["userid"] = userid
	return p
}

// Retrieves a cloud identifier.
func (s *SystemService) GetCloudIdentifier(p *GetCloudIdentifierParams) (*GetCloudIdentifierResponse, error) {
	resp, err := s.cs.newRequest("getCloudIdentifier", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetCloudIdentifierResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type GetCloudIdentifierResponse struct {
	Cloudidentifier string `json:"cloudidentifier,omitempty"`
	Signature       string `json:"signature,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
