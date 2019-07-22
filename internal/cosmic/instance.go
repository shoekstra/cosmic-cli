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
	h "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

// VirtualMachine embeds *cosmic.VirtualMachine to allow additional fields.
type VirtualMachine struct {
	*cosmic.VirtualMachine
	Networkname string
	Vpcname     string
}

// VirtualMachines exists to provide helper methods for []*VirtualMachine.
type VirtualMachines []*VirtualMachine

// FindByID looks for a VirtualMachine object by ID in VirtualMachines and returns it if it exists.
func (vms VirtualMachines) FindByID(id string) ([]*VirtualMachine, error) {
	r := []*VirtualMachine{}
	for _, v := range vms {
		if v.Id == id {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for instance with id %s", id)
	}
	return r, nil
}

// FindByName looks for a VirtualMachine object by name in VirtualMachines and returns it if it exists.
func (vms VirtualMachines) FindByName(name string) ([]*VirtualMachine, error) {
	r := []*VirtualMachine{}
	for _, v := range vms {
		if v.Name == name {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for instance with name %s", name)
	}
	if len(r) > 1 {
		return r, fmt.Errorf("More than one match found for instance with name %s, use the --instance-id option to specify the instance", name)
	}
	return r, nil
}

// Sort will sort VirtualMachines by either the "ipaddress", "name" or "zonename" field.
func (vms VirtualMachines) Sort(sortBy string, reverseSort bool) {
	if h.Contains([]string{"ipaddress", "name", "zonename"}, sortBy) == false {
		fmt.Println("Invalid sort option provided, provide either \"ipaddress\", \"name\" or \"zonename\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Ipaddress"):
		sort.SliceStable(vms, func(i, j int) bool {
			if reverseSort {
				return vms[i].Nic[0].Ipaddress > vms[j].Nic[0].Ipaddress
			}
			return vms[i].Nic[0].Ipaddress < vms[j].Nic[0].Ipaddress
		})
	case strings.EqualFold(sortBy, "Name"):
		sort.SliceStable(vms, func(i, j int) bool {
			if reverseSort {
				return vms[i].Name > vms[j].Name
			}
			return vms[i].Name < vms[j].Name
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(vms, func(i, j int) bool {
			if reverseSort {
				return vms[i].Zonename > vms[j].Zonename
			}
			return vms[i].Zonename < vms[j].Zonename
		})
	}
}

// ListVMs returns a VirtualMachines object using all configured *cosmic.CosmicClient objects.
func ListVMs(clientMap map[string]*cosmic.CosmicClient) (VirtualMachines, error) {
	vms := []*VirtualMachine{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VirtualMachine.NewListVirtualMachinesParams()
			resp, err := clientMap[client].VirtualMachine.ListVirtualMachines(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, vm := range resp.VirtualMachines {
				vms = append(vms, &VirtualMachine{
					VirtualMachine: vm,
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

	return vms, nil
}
