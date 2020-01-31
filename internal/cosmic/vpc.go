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

// VPC embeds *cosmic.VPC to allow additional fields.
type VPC struct {
	*cosmic.VPC
	Sourcenatip string
}

// VPCs exists to provide helper methods for []*VPC.
type VPCs []*VPC

// FindByID looks for a VPC object by ID in VPCs and returns it if it exists.
func (v VPCs) FindByID(id string) ([]*VPC, error) {
	r := []*VPC{}
	for _, i := range v {
		if i.Id == id {
			r = append(r, i)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for vpc with id %s", id)
	}
	return r, nil
}

// FindByName looks for a VPC object by name in VPCs and returns it if it exists.
func (v VPCs) FindByName(name string) ([]*VPC, error) {
	r := []*VPC{}
	for _, i := range v {
		if i.Name == name {
			r = append(r, i)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for vpc with name %s", name)
	}
	if len(r) > 1 {
		return r, fmt.Errorf("More than one match found for vpc with name %s, use the --vpc-id option to specify the vpc", name)
	}
	return r, nil
}

// Sort will sort VPCs by either the "cidr", "name", "vpcofferingname" or "zonename" field.
func (v VPCs) Sort(sortBy string, reverseSort bool) {
	if !h.Contains([]string{"cidr", "name", "vpcofferingname", "zonename"}, sortBy) {
		fmt.Println("Invalid sort option provided, provide either \"cidr\", \"name\", \"vpcofferingname\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Cidr"):
		sort.SliceStable(v, func(i, j int) bool {
			if reverseSort {
				return v[i].Cidr > v[j].Cidr
			}
			return v[i].Cidr < v[j].Cidr
		})
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(v, func(i, j int) bool {
			if reverseSort {
				return v[i].Name > v[j].Name
			}
			return v[i].Name < v[j].Name
		})
	case strings.EqualFold(sortBy, "Vpcofferingname"):
		sort.SliceStable(v, func(i, j int) bool {
			if reverseSort {
				return v[i].Vpcofferingname > v[j].Vpcofferingname
			}
			return v[i].Vpcofferingname < v[j].Vpcofferingname
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(v, func(i, j int) bool {
			if reverseSort {
				return v[i].Zonename > v[j].Zonename
			}
			return v[i].Zonename < v[j].Zonename
		})
	}
}

// VPCGetByID returns a *cosmic.VPC object using a *cosmic.CosmicClient object.
func VPCGetByID(client *cosmic.CosmicClient, id string) (*cosmic.VPC, int, error) {
	resp, count, err := client.VPC.GetVPCByID(id)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil, 0, nil
		}
		return resp, count, err
	}

	return resp, count, nil
}

// VPCGetByName returns a *cosmic.VPC object using a *cosmic.CosmicClient object.
func VPCGetByName(client *cosmic.CosmicClient, name string) (*cosmic.VPC, int, error) {
	resp, count, err := client.VPC.GetVPCByName(name)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil, 0, nil
		}
		return resp, count, err
	}

	return resp, count, nil
}

// ListVPCs returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func ListVPCs(clientMap map[string]*cosmic.CosmicClient) (VPCs, error) {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VPC.NewListVPCsParams()
			resp, err := clientMap[client].VPC.ListVPCs(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, vpc := range resp.VPCs {
				vpcs = append(vpcs, &VPC{
					VPC: vpc,
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

	return vpcs, nil
}

// VPCGetAllByID returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func VPCGetAllByID(clientMap map[string]*cosmic.CosmicClient, id string) ([]*VPC, error) {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, _ := VPCGetByID(clientMap[client], id)
			if count == 1 {
				vpcs = append(vpcs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	if len(vpcs) == 0 {
		return nil, fmt.Errorf("No match found for VPC with id %s", id)
	}

	if len(vpcs) > 1 {
		return vpcs, fmt.Errorf("More than one match found for VPC with id %s, use the --vpc-id option to specify the VPC", id)
	}

	return vpcs, nil
}

// VPCGetAllByName returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func VPCGetAllByName(clientMap map[string]*cosmic.CosmicClient, name string) ([]*VPC, error) {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, _ := VPCGetByName(clientMap[client], name)
			if count == 1 {
				vpcs = append(vpcs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	if len(vpcs) == 0 {
		return nil, fmt.Errorf("No match found for VPC with name %s", name)
	}

	if len(vpcs) > 1 {
		return vpcs, fmt.Errorf("More than one match found for VPC with name %s, use the --vpc-id option to specify the VPC", name)
	}

	return vpcs, nil
}
