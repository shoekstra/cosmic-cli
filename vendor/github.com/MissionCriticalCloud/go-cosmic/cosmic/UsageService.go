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

type AddTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *AddTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["isolationmethod"]; found {
		u.Set("isolationmethod", v.(string))
	}
	if v, found := p.p["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (p *AddTrafficTypeParams) SetIsolationmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolationmethod"] = v
}

func (p *AddTrafficTypeParams) SetKvmnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kvmnetworklabel"] = v
}

func (p *AddTrafficTypeParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddTrafficTypeParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *AddTrafficTypeParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *AddTrafficTypeParams) SetXennetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["xennetworklabel"] = v
}

// You should always use this function to get a new AddTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams {
	p := &AddTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["traffictype"] = traffictype
	return p
}

// Adds traffic type to a physical network
func (s *UsageService) AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("addTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficTypeResponse
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

type AddTrafficTypeResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Id                string `json:"id,omitempty"`
	Kvmnetworklabel   string `json:"kvmnetworklabel,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Traffictype       string `json:"traffictype,omitempty"`
	Xennetworklabel   string `json:"xennetworklabel,omitempty"`
}

type DeleteTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *DeleteTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteTrafficTypeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams {
	p := &DeleteTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes traffic type of a physical network
func (s *UsageService) DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("deleteTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficTypeResponse
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

type DeleteTrafficTypeResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListTrafficTypesParams struct {
	p map[string]interface{}
}

func (p *ListTrafficTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListTrafficTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTrafficTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTrafficTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTrafficTypesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

// You should always use this function to get a new ListTrafficTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UsageService) GetTrafficTypeID(keyword string, physicalnetworkid string, opts ...OptionFunc) (string, int, error) {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["physicalnetworkid"] = physicalnetworkid

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListTrafficTypes(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.TrafficTypes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.TrafficTypes {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists traffic types of a given physical network.
func (s *UsageService) ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error) {
	var r ListTrafficTypesResponse
	for page := 2; ; page++ {
		var l ListTrafficTypesResponse
		resp, err := s.cs.newRequest("listTrafficTypes", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.TrafficTypes = append(r.TrafficTypes, l.TrafficTypes...)

		if r.Count == len(r.TrafficTypes) {
			return &r, nil
		}

		p.SetPagesize(len(l.TrafficTypes))
		p.SetPage(page)
	}
}

type ListTrafficTypesResponse struct {
	Count        int            `json:"count"`
	TrafficTypes []*TrafficType `json:"traffictype"`
}

type TrafficType struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
	Id                           string   `json:"id,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
	Servicelist                  []string `json:"servicelist,omitempty"`
	State                        string   `json:"state,omitempty"`
}
