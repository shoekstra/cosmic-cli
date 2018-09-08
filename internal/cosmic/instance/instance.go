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

package instance

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

// VirtualMachine embeds *cosmic.VirtualMachine to allow additional fields.
type VirtualMachine struct {
	*cosmic.VirtualMachine
	Networkname string
	Vpcname     string
}

// List returns a slice of *cosmic.VirtualMachine objects using a *cosmic.CosmicClient object.
func List(client *cosmic.CosmicClient) ([]*cosmic.VirtualMachine, error) {
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	resp, err := client.VirtualMachine.ListVirtualMachines(params)
	if err != nil {
		return resp.VirtualMachines, err
	}

	return resp.VirtualMachines, nil
}

// ListAll returns a slice of *VirtualMachine objects using all configured *cosmic.CosmicClient
// objects.
func ListAll(clientMap map[string]*cosmic.CosmicClient, filter, sortBy string, reverseSort bool) []*VirtualMachine {
	VirtualMachines := []*VirtualMachine{}
	var wg sync.WaitGroup
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listVirtualMachines, err := List(clientMap[client])
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, vm := range listVirtualMachines {
				VirtualMachines = append(VirtualMachines, &VirtualMachine{
					VirtualMachine: vm,
				})
			}
		}(client)
	}
	wg.Wait()

	VirtualMachines = Sort(VirtualMachines, sortBy, reverseSort)

	return VirtualMachines
}

// Sort returns a sorted []*cosmic.VirtualMachine slice.
func Sort(VirtualMachines []*VirtualMachine, sortBy string, reverseSort bool) []*VirtualMachine {
	if h.Contains([]string{"ipaddress", "name", "zonename"}, sortBy) == false {
		fmt.Println("Invalid sort option provided, provide either \"ipaddress\", \"name\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Ipaddress"):
		sort.SliceStable(VirtualMachines, func(i, j int) bool {
			if reverseSort {
				return VirtualMachines[i].Nic[0].Ipaddress > VirtualMachines[j].Nic[0].Ipaddress
			}
			return VirtualMachines[i].Nic[0].Ipaddress < VirtualMachines[j].Nic[0].Ipaddress
		})
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(VirtualMachines, func(i, j int) bool {
			if reverseSort {
				return VirtualMachines[i].Name > VirtualMachines[j].Name
			}
			return VirtualMachines[i].Name < VirtualMachines[j].Name
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(VirtualMachines, func(i, j int) bool {
			if reverseSort {
				return VirtualMachines[i].Zonename > VirtualMachines[j].Zonename
			}
			return VirtualMachines[i].Zonename < VirtualMachines[j].Zonename
		})
	}

	return VirtualMachines
}
