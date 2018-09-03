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
	. "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

type VPC struct {
	*cosmic.VPC
	SourceNatIP string
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

// List returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func ListAll(clientMap map[string]*cosmic.CosmicClient, filter, sortBy string, reverseSort bool) []*VPC {
	var VPCs []*VPC
	var wg sync.WaitGroup
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listVPCs, err := List(clientMap[client])
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, vpc := range listVPCs {
				VPCs = append(VPCs, &VPC{
					VPC: vpc,
				})
			}
		}(client)
	}
	wg.Wait()

	VPCs = Sort(VPCs, sortBy, reverseSort)

	return VPCs
}

// Sort returns a sorted slice of []*VPC objects.
func Sort(VPCs []*VPC, sortBy string, reverseSort bool) []*VPC {
	if Contains([]string{"cidr", "name", "vpcofferingname", "zonename"}, sortBy) == false {
		fmt.Println("Invalid sort option provided, provide either \"cidr\", \"name\", \"vpcofferingname\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Cidr"):
		sort.SliceStable(VPCs, func(i, j int) bool {
			if reverseSort {
				return VPCs[i].Cidr > VPCs[j].Cidr
			}
			return VPCs[i].Cidr < VPCs[j].Cidr
		})
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(VPCs, func(i, j int) bool {
			if reverseSort {
				return VPCs[i].Name > VPCs[j].Name
			}
			return VPCs[i].Name < VPCs[j].Name
		})
	case strings.EqualFold(sortBy, "Vpcofferingname"):
		sort.SliceStable(VPCs, func(i, j int) bool {
			if reverseSort {
				return VPCs[i].Vpcofferingname > VPCs[j].Vpcofferingname
			}
			return VPCs[i].Vpcofferingname < VPCs[j].Vpcofferingname
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(VPCs, func(i, j int) bool {
			if reverseSort {
				return VPCs[i].Zonename > VPCs[j].Zonename
			}
			return VPCs[i].Zonename < VPCs[j].Zonename
		})
	}

	return VPCs
}
