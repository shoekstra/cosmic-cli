//
// Copyright © 2018 Stephen Hoekstra
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

package vpc

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	h "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

// VPC embeds *cosmic.VPC to allow additional fields.
type VPC struct {
	*cosmic.VPC
	Sourcenatip string
}

// VPCs exists to provide helper methods for []*VPC.
type VPCs []*VPC

// Sort returns a sorted slice of []*VPC objects.
func (vpcs VPCs) Sort(sortBy string, reverseSort bool) {
	if h.Contains([]string{"cidr", "name", "vpcofferingname", "zonename"}, sortBy) == false {
		fmt.Println("Invalid sort option provided, provide either \"cidr\", \"name\", \"vpcofferingname\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Cidr"):
		sort.SliceStable(vpcs, func(i, j int) bool {
			if reverseSort {
				return vpcs[i].Cidr > vpcs[j].Cidr
			}
			return vpcs[i].Cidr < vpcs[j].Cidr
		})
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(vpcs, func(i, j int) bool {
			if reverseSort {
				return vpcs[i].Name > vpcs[j].Name
			}
			return vpcs[i].Name < vpcs[j].Name
		})
	case strings.EqualFold(sortBy, "Vpcofferingname"):
		sort.SliceStable(vpcs, func(i, j int) bool {
			if reverseSort {
				return vpcs[i].Vpcofferingname > vpcs[j].Vpcofferingname
			}
			return vpcs[i].Vpcofferingname < vpcs[j].Vpcofferingname
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(vpcs, func(i, j int) bool {
			if reverseSort {
				return vpcs[i].Zonename > vpcs[j].Zonename
			}
			return vpcs[i].Zonename < vpcs[j].Zonename
		})
	}
}

// List returns a slice of *VPC objects using a *cosmic.CosmicClient object.
func List(client *cosmic.CosmicClient) ([]*cosmic.VPC, error) {
	params := client.VPC.NewListVPCsParams()
	resp, err := client.VPC.ListVPCs(params)
	if err != nil {
		return resp.VPCs, err
	}

	return resp.VPCs, nil
}

// GetByID returns a *cosmic.VPC object using a *cosmic.CosmicClient object.
func GetByID(client *cosmic.CosmicClient, id string) (*cosmic.VPC, int, error) {
	resp, count, err := client.VPC.GetVPCByID(id)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil, 0, nil
		}
		return resp, count, err
	}

	return resp, count, nil
}

// GetByName returns a *cosmic.VPC object using a *cosmic.CosmicClient object.
func GetByName(client *cosmic.CosmicClient, name string) (*cosmic.VPC, int, error) {
	resp, count, err := client.VPC.GetVPCByName(name)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil, 0, nil
		}
		return resp, count, err
	}

	return resp, count, nil
}

// ListAll returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func ListAll(clientMap map[string]*cosmic.CosmicClient) VPCs {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listVPCs, err := List(clientMap[client])
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, vpc := range listVPCs {
				vpcs = append(vpcs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	return vpcs
}

// GetAllByID returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func GetAllByID(clientMap map[string]*cosmic.CosmicClient, id string) ([]*VPC, error) {
	var VPCs []*VPC
	var wg sync.WaitGroup
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, err := GetByID(clientMap[client], id)
			if err != nil {
				// return nil, err
			}
			if count == 1 {
				VPCs = append(VPCs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	if len(VPCs) == 0 {
		return nil, fmt.Errorf("No match found for VPC with id %s", id)
	}

	if len(VPCs) > 1 {
		return VPCs, fmt.Errorf("More than one match found for VPC with id %s, use the --vpc-id option to specify the VPC", id)
	}

	return VPCs, nil
}

// GetAllByName returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func GetAllByName(clientMap map[string]*cosmic.CosmicClient, name string) ([]*VPC, error) {
	var VPCs []*VPC
	var wg sync.WaitGroup
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, err := GetByName(clientMap[client], name)
			if err != nil {
				// return nil, err
			}
			if count == 1 {
				VPCs = append(VPCs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	if len(VPCs) == 0 {
		return nil, fmt.Errorf("No match found for VPC with name %s", name)
	}

	if len(VPCs) > 1 {
		return VPCs, fmt.Errorf("More than one match found for VPC with name %s, use the --vpc-id option to specify the VPC", name)
	}

	return VPCs, nil
}
