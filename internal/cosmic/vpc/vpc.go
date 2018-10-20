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
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	h "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

// profileError represents a profile config error.
type profileError struct {
	message string
}

// Error returns the profile error message.
func (e profileError) Error() string {
	return e.message
}

// VPC embeds *cosmic.VPC to allow additional fields.
type VPC struct {
	*cosmic.VPC
	Sourcenatip string
}

// VPCs exists to provide helper methods for []*VPC.
type VPCs []*VPC

// Sort will sort VPCs by either the "cidr", "name", "vpcofferingname" or "zonename" field.
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

// List returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func List(clientMap map[string]*cosmic.CosmicClient) (VPCs, error) {
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

// GetAllByID returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func GetAllByID(clientMap map[string]*cosmic.CosmicClient, id string) ([]*VPC, error) {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, err := GetByID(clientMap[client], id)
			if err != nil {
				// return nil, err
			}
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

// GetAllByName returns a slice of *VPC objects using all configured *cosmic.CosmicClient objects.
func GetAllByName(clientMap map[string]*cosmic.CosmicClient, name string) ([]*VPC, error) {
	vpcs := []*VPC{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			vpc, count, err := GetByName(clientMap[client], name)
			if err != nil {
				// return nil, err
			}
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
