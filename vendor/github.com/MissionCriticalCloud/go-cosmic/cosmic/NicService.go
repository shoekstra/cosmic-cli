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

type RemoveIpFromNicParams struct {
	p map[string]interface{}
}

func (p *RemoveIpFromNicParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RemoveIpFromNicParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new RemoveIpFromNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewRemoveIpFromNicParams(id string) *RemoveIpFromNicParams {
	p := &RemoveIpFromNicParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes secondary IP from the NIC.
func (s *NicService) RemoveIpFromNic(p *RemoveIpFromNicParams) (*RemoveIpFromNicResponse, error) {
	resp, err := s.cs.newRequest("removeIpFromNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveIpFromNicResponse
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

type RemoveIpFromNicResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type AddIpToNicParams struct {
	p map[string]interface{}
}

func (p *AddIpToNicParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	return u
}

func (p *AddIpToNicParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *AddIpToNicParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
}

// You should always use this function to get a new AddIpToNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewAddIpToNicParams(nicid string) *AddIpToNicParams {
	p := &AddIpToNicParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	return p
}

// Assigns secondary IP to NIC
func (s *NicService) AddIpToNic(p *AddIpToNicParams) (*AddIpToNicResponse, error) {
	resp, err := s.cs.newRequest("addIpToNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddIpToNicResponse
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

type AddIpToNicResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Id               string `json:"id,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Networkid        string `json:"networkid,omitempty"`
	Nicid            string `json:"nicid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
}

type ListNicsParams struct {
	p map[string]interface{}
}

func (p *ListNicsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListNicsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListNicsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNicsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListNicsParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
}

func (p *ListNicsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNicsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNicsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

// You should always use this function to get a new ListNicsParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewListNicsParams(virtualmachineid string) *ListNicsParams {
	p := &ListNicsParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// list the vm nics  IP to NIC
func (s *NicService) ListNics(p *ListNicsParams) (*ListNicsResponse, error) {
	var r ListNicsResponse
	for page := 2; ; page++ {
		var l ListNicsResponse
		resp, err := s.cs.newRequest("listNics", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.Nics = append(r.Nics, l.Nics...)

		if r.Count == len(r.Nics) {
			return &r, nil
		}

		p.SetPagesize(len(l.Nics))
		p.SetPage(page)
	}
}

type ListNicsResponse struct {
	Count int    `json:"count"`
	Nics  []*Nic `json:"nic"`
}

type Nic struct {
	Broadcasturi string `json:"broadcasturi,omitempty"`
	Gateway      string `json:"gateway,omitempty"`
	Id           string `json:"id,omitempty"`
	Ip6address   string `json:"ip6address,omitempty"`
	Ip6cidr      string `json:"ip6cidr,omitempty"`
	Ip6gateway   string `json:"ip6gateway,omitempty"`
	Ipaddress    string `json:"ipaddress,omitempty"`
	Isdefault    bool   `json:"isdefault,omitempty"`
	Isolationuri string `json:"isolationuri,omitempty"`
	Macaddress   string `json:"macaddress,omitempty"`
	Netmask      string `json:"netmask,omitempty"`
	Networkid    string `json:"networkid,omitempty"`
	Networkname  string `json:"networkname,omitempty"`
	Secondaryip  []struct {
		Id        string `json:"id,omitempty"`
		Ipaddress string `json:"ipaddress,omitempty"`
	} `json:"secondaryip,omitempty"`
	Traffictype      string `json:"traffictype,omitempty"`
	Type             string `json:"type,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
}

type UpdateVmNicIpParams struct {
	p map[string]interface{}
}

func (p *UpdateVmNicIpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	return u
}

func (p *UpdateVmNicIpParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *UpdateVmNicIpParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
}

// You should always use this function to get a new UpdateVmNicIpParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewUpdateVmNicIpParams(nicid string) *UpdateVmNicIpParams {
	p := &UpdateVmNicIpParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	return p
}

// Update the default Ip of a VM Nic
func (s *NicService) UpdateVmNicIp(p *UpdateVmNicIpParams) (*UpdateVmNicIpResponse, error) {
	resp, err := s.cs.newRequest("updateVmNicIp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVmNicIpResponse
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

type UpdateVmNicIpResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Bootmenutimeout       int64             `json:"bootmenutimeout,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Laststartdate         string            `json:"laststartdate,omitempty"`
	Laststartversion      string            `json:"laststartversion,omitempty"`
	Maintenancepolicy     string            `json:"maintenancepolicy,omitempty"`
	Manufacturerstring    string            `json:"manufacturerstring,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Optimisefor          string `json:"optimisefor,omitempty"`
	Ostypeid             int64  `json:"ostypeid,omitempty"`
	Password             string `json:"password,omitempty"`
	Passwordenabled      bool   `json:"passwordenabled,omitempty"`
	Project              string `json:"project,omitempty"`
	Projectid            string `json:"projectid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Publicipid           string `json:"publicipid,omitempty"`
	Restartrequired      bool   `json:"restartrequired,omitempty"`
	Rootdevicecontroller string `json:"rootdevicecontroller,omitempty"`
	Rootdeviceid         int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype       string `json:"rootdevicetype,omitempty"`
	Serviceofferingid    string `json:"serviceofferingid,omitempty"`
	Serviceofferingname  string `json:"serviceofferingname,omitempty"`
	Servicestate         string `json:"servicestate,omitempty"`
	State                string `json:"state,omitempty"`
	Tags                 []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}
