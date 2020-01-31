//
// Copyright © 2019 Stephen Hoekstra
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
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	h "github.com/shoekstra/cosmic-cli/internal/helper"
)

// ACL embeds *cosmic.NetworkACLList to allow additional fields.
type ACL struct {
	*cosmic.NetworkACLList
	Vpcname  string
	Zonename string
}

// ACLs exists to provide helper methods for []*ACL.
type ACLs []*ACL

// FindByID looks for an ACL object by ID in ACLs and returns it if it exists.
func (a ACLs) FindByID(id string) ([]*ACL, error) {
	r := []*ACL{}
	for _, v := range a {
		if v.Id == id {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for ACL with id %s", id)
	}
	return r, nil
}

// FindByName looks for an ACL object by name in ACLs and returns it if it exists.
func (a ACLs) FindByName(name string) ([]*ACL, error) {
	r := []*ACL{}
	for _, v := range a {
		if v.Name == name {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for ACL with name %s", name)
	}
	if len(r) > 1 {
		return r, fmt.Errorf("More than one match found for ACL with name %s, use the --acl-id option to specify the ACL", name)
	}
	return r, nil
}

// Sort will sort ACLs by either the "name", "vpcname" or "zonename" field.
func (a ACLs) Sort(sortBy string, reverseSort bool) {
	if !h.Contains([]string{"name", "vpcname", "zonename"}, sortBy) {
		fmt.Println("Invalid sort option provided, provide either \"name\", \"vpcname\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(a, func(i, j int) bool {
			if reverseSort {
				return a[i].Name > a[j].Name
			}
			return a[i].Name < a[j].Name
		})
	case strings.EqualFold(sortBy, "Vpcname"):
		sort.SliceStable(a, func(i, j int) bool {
			if reverseSort {
				return a[i].Vpcname > a[j].Vpcname
			}
			return a[i].Vpcname < a[j].Vpcname
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(a, func(i, j int) bool {
			if reverseSort {
				return a[i].Zonename > a[j].Zonename
			}
			return a[i].Zonename < a[j].Zonename
		})
	}
}

// ListACLs returns a ACLs object using all configured *cosmic.CosmicClient objects.
func ListACLs(clientMap map[string]*cosmic.CosmicClient) (ACLs, error) {
	acls := []*ACL{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			// Zonename isn't returned in *cosmic.ListNetworkACLListsResponse so we need to fetch it
			zoneparams := clientMap[client].Zone.NewListZonesParams()
			zoneresp, err := clientMap[client].Zone.ListZones(zoneparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			zonename := zoneresp.Zones[0].Name

			// Fetch VPCs so we can translate VPC IDs to names
			vpcparams := clientMap[client].VPC.NewListVPCsParams()
			vpcresp, err := clientMap[client].VPC.ListVPCs(vpcparams)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			// Fetch ACLs
			params := clientMap[client].NetworkACL.NewListNetworkACLListsParams()
			resp, err := clientMap[client].NetworkACL.ListNetworkACLLists(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, acl := range resp.NetworkACLLists {
				vpcname := ""
				for _, v := range vpcresp.VPCs {
					if v.Id == acl.Vpcid {
						vpcname = v.Name
						break
					}
				}
				acls = append(acls, &ACL{
					NetworkACLList: acl,
					Vpcname:        vpcname,
					Zonename:       zonename,
				})
			}
		}(client)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return nil, err
		}
	}

	return acls, nil
}

// ACLRule embeds *cosmic.NetworkACLRule to allow additional fields.
type ACLRule struct {
	*cosmic.NetworkACL
	Aclname  string
	Vpcname  string
	Zonename string
}

// ACLRules exists to provide helper methods for []*ACLRule.
type ACLRules []*ACLRule

// Sort will sort ACLs by either the "name", "vpcname" or "zonename" field.
func (rules ACLRules) Sort(sortBy string, reverseSort bool) {
	if !h.Contains([]string{"action", "cidrlist", "endport", "number", "startport"}, sortBy) {
		fmt.Println("Invalid sort option provided, provide either \"action\", \"cidrlist\", \"endport\", \"number\" or \"startport\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Action"):
		sort.SliceStable(rules, func(i, j int) bool {
			if reverseSort {
				return rules[i].Action > rules[j].Action
			}
			return rules[i].Action < rules[j].Action
		})
	case strings.EqualFold(sortBy, "Cidrlist"):
		sort.SliceStable(rules, func(i, j int) bool {
			if reverseSort {
				return rules[i].Cidrlist > rules[j].Cidrlist
			}
			return rules[i].Cidrlist < rules[j].Cidrlist
		})
	case strings.EqualFold(sortBy, "Endport"):
		sort.SliceStable(rules, func(i, j int) bool {
			if reverseSort {
				return rules[i].Endport > rules[j].Endport
			}
			return rules[i].Endport < rules[j].Endport
		})
	case strings.EqualFold(sortBy, "Number"):
		sort.SliceStable(rules, func(i, j int) bool {
			if reverseSort {
				return rules[i].Number > rules[j].Number
			}
			return rules[i].Number < rules[j].Number
		})
	case strings.EqualFold(sortBy, "Startport"):
		sort.SliceStable(rules, func(i, j int) bool {
			if reverseSort {
				return rules[i].Startport > rules[j].Startport
			}
			return rules[i].Startport < rules[j].Startport
		})
	}
}

// ListACLRules returns a ACLRules object using all configured *cosmic.CosmicClient objects.
func ListACLRules(clientMap map[string]*cosmic.CosmicClient, aclid string) (ACLRules, error) {
	acls := []*ACLRule{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].NetworkACL.NewListNetworkACLsParams()
			params.SetAclid(aclid)
			resp, err := clientMap[client].NetworkACL.ListNetworkACLs(params)
			if err != nil {
				if strings.Contains(err.Error(), fmt.Sprintf("does not have permission")) {
					return
				}
				if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
					return
				}
				if strings.Contains(err.Error(), fmt.Sprintf("Unable to find VPC associated with acl")) {
					return
				}
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}
			a, _, _ := clientMap[client].NetworkACL.GetNetworkACLListByID(aclid)
			for _, acl := range resp.NetworkACLs {
				acls = append(acls, &ACLRule{
					NetworkACL: acl,
					Aclname:    a.Name,
				})
			}
		}(client)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return nil, err
		}
	}

	return acls, nil
}
