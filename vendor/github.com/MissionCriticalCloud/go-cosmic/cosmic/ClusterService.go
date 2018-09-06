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

type AddClusterParams struct {
	p map[string]interface{}
}

func (p *AddClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *AddClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *AddClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *AddClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *AddClusterParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddClusterParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddClusterParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddClusterParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddClusterParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new AddClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams {
	p := &AddClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clustername"] = clustername
	p.p["clustertype"] = clustertype
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new cluster
func (s *ClusterService) AddCluster(p *AddClusterParams) (*AddClusterResponse, error) {
	resp, err := s.cs.newRequest("addCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddClusterResponse struct {
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
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
	} `json:"capacity,omitempty"`
	Clustertype           string `json:"clustertype,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Hypervisortype        string `json:"hypervisortype,omitempty"`
	Id                    string `json:"id,omitempty"`
	Managedstate          string `json:"managedstate,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Name                  string `json:"name,omitempty"`
	Podid                 string `json:"podid,omitempty"`
	Podname               string `json:"podname,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
}

type DedicateClusterParams struct {
	p map[string]interface{}
}

func (p *DedicateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *DedicateClusterParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicateClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *DedicateClusterParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

// You should always use this function to get a new DedicateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams {
	p := &DedicateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["domainid"] = domainid
	return p
}

// Dedicate an existing cluster
func (s *ClusterService) DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error) {
	resp, err := s.cs.newRequest("dedicateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateClusterResponse
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

type DedicateClusterResponse struct {
	JobID           string `json:"jobid,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Accountname     string `json:"accountname,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Clusterid       string `json:"clusterid,omitempty"`
	Clustername     string `json:"clustername,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Domainname      string `json:"domainname,omitempty"`
	Id              string `json:"id,omitempty"`
}

type DeleteClusterParams struct {
	p map[string]interface{}
}

func (p *DeleteClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDeleteClusterParams(id string) *DeleteClusterParams {
	p := &DeleteClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a cluster.
func (s *ClusterService) DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteClusterResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type UpdateClusterParams struct {
	p map[string]interface{}
}

func (p *UpdateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	return u
}

func (p *UpdateClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *UpdateClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *UpdateClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *UpdateClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *UpdateClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateClusterParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
}

// You should always use this function to get a new UpdateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewUpdateClusterParams(id string) *UpdateClusterParams {
	p := &UpdateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing cluster
func (s *ClusterService) UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error) {
	resp, err := s.cs.newRequest("updateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateClusterResponse struct {
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
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
	} `json:"capacity,omitempty"`
	Clustertype           string `json:"clustertype,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Hypervisortype        string `json:"hypervisortype,omitempty"`
	Id                    string `json:"id,omitempty"`
	Managedstate          string `json:"managedstate,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Name                  string `json:"name,omitempty"`
	Podid                 string `json:"podid,omitempty"`
	Podname               string `json:"podname,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
}

type ListClustersParams struct {
	p map[string]interface{}
}

func (p *ListClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListClustersParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *ListClustersParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *ListClustersParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListClustersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListClustersParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
}

func (p *ListClustersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListClustersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListClustersParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListClustersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersParams() *ListClustersParams {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClusters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Clusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Clusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error) {
	id, count, err := s.GetClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClusters(p)
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
		return l.Clusters[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for Cluster UUID: %s!", id)
}

// Lists clusters.
func (s *ClusterService) ListClusters(p *ListClustersParams) (*ListClustersResponse, error) {
	var r ListClustersResponse
	for page := 2; ; page++ {
		var l ListClustersResponse
		resp, err := s.cs.newRequest("listClusters", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.Clusters = append(r.Clusters, l.Clusters...)

		if r.Count == len(r.Clusters) {
			return &r, nil
		}

		p.SetPagesize(len(l.Clusters))
		p.SetPage(page)
	}
}

type ListClustersResponse struct {
	Count    int        `json:"count"`
	Clusters []*Cluster `json:"cluster"`
}

type Cluster struct {
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
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
	} `json:"capacity,omitempty"`
	Clustertype           string `json:"clustertype,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Hypervisortype        string `json:"hypervisortype,omitempty"`
	Id                    string `json:"id,omitempty"`
	Managedstate          string `json:"managedstate,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Name                  string `json:"name,omitempty"`
	Podid                 string `json:"podid,omitempty"`
	Podname               string `json:"podname,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
}

type ReleaseDedicatedClusterParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

// You should always use this function to get a new ReleaseDedicatedClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams {
	p := &ReleaseDedicatedClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Release the dedication for cluster
func (s *ClusterService) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedClusterResponse
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

type ReleaseDedicatedClusterResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListDedicatedClustersParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	return u
}

func (p *ListDedicatedClustersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDedicatedClustersParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListDedicatedClustersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListDedicatedClustersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDedicatedClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDedicatedClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDedicatedClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListDedicatedClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListDedicatedClustersParams() *ListDedicatedClustersParams {
	p := &ListDedicatedClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists dedicated clusters.
func (s *ClusterService) ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error) {
	var r ListDedicatedClustersResponse
	for page := 2; ; page++ {
		var l ListDedicatedClustersResponse
		resp, err := s.cs.newRequest("listDedicatedClusters", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.DedicatedClusters = append(r.DedicatedClusters, l.DedicatedClusters...)

		if r.Count == len(r.DedicatedClusters) {
			return &r, nil
		}

		p.SetPagesize(len(l.DedicatedClusters))
		p.SetPage(page)
	}
}

type ListDedicatedClustersResponse struct {
	Count             int                 `json:"count"`
	DedicatedClusters []*DedicatedCluster `json:"dedicatedcluster"`
}

type DedicatedCluster struct {
	Accountid       string `json:"accountid,omitempty"`
	Accountname     string `json:"accountname,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Clusterid       string `json:"clusterid,omitempty"`
	Clustername     string `json:"clustername,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Domainname      string `json:"domainname,omitempty"`
	Id              string `json:"id,omitempty"`
}
