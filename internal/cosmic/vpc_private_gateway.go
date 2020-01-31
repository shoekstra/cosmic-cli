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

// PrivateGateway embeds *cosmic.PrivateGateway to allow additional fields.
type PrivateGateway struct {
	*cosmic.PrivateGateway
	Vpccidr string
	Vpcname string
}

// PrivateGateways exists to provide helper methods for []*PrivateGateway.
type PrivateGateways []*PrivateGateway

// FindByIPAddress looks for a PrivateGateway object by ID in PrivateGateways and returns it if it exists.
func (p PrivateGateways) FindByIPAddress(ip string) []*PrivateGateway {
	pgws := []*PrivateGateway{}
	for _, pgw := range p {
		if pgw.Ipaddress == ip {
			pgws = append(pgws, pgw)
		}
	}
	return pgws
}

// Sort will sort PrivateGateways by either the "cidr" or "name" field.
func (p PrivateGateways) Sort(sortBy string, reverseSort bool) {
	if !h.Contains([]string{"cidr", "ipaddress", "vpccidr", "vpcname", "zonename"}, sortBy) {
		fmt.Println("Invalid sort option provided, provide either \"cidr\", \"ipaddress\", \"vpccidr\", \"vpcname\" or \"zonename\".")
		os.Exit(1)
	}
	switch {
	case strings.EqualFold(sortBy, "Cidr"):
		sort.SliceStable(p, func(i, j int) bool {
			if reverseSort {
				return p[i].Cidr > p[j].Cidr
			}
			return p[i].Cidr < p[j].Cidr
		})
	case strings.EqualFold(sortBy, "Ipaddress"):
		sort.SliceStable(p, func(i, j int) bool {
			if reverseSort {
				return p[i].Ipaddress > p[j].Ipaddress
			}
			return p[i].Ipaddress < p[j].Ipaddress
		})
	case strings.EqualFold(sortBy, "Vpccidr"):
		sort.SliceStable(p, func(i, j int) bool {
			if reverseSort {
				return p[i].Vpccidr > p[j].Vpccidr
			}
			return p[i].Vpccidr < p[j].Vpccidr
		})
	case strings.EqualFold(sortBy, "Vpcname"):
		sort.SliceStable(p, func(i, j int) bool {
			if reverseSort {
				return p[i].Vpcname > p[j].Vpcname
			}
			return p[i].Vpcname < p[j].Vpcname
		})
	case strings.EqualFold(sortBy, "Zonename"):
		sort.SliceStable(p, func(i, j int) bool {
			if reverseSort {
				return p[i].Zonename > p[j].Zonename
			}
			return p[i].Zonename < p[j].Zonename
		})
	}
}

// ListVPCPrivateGateways returns a PrivateGateways object using all configured *cosmic.CosmicClient objects.
func ListVPCPrivateGateways(clientMap map[string]*cosmic.CosmicClient) (PrivateGateways, error) {
	pgws := []*PrivateGateway{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	VPCs, err := ListVPCs(clientMap)
	if err != nil {
		return nil, err
	}

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VPC.NewListPrivateGatewaysParams()
			resp, err := clientMap[client].VPC.ListPrivateGateways(params)
			if err != nil {
				if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
					return
				}
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, pgw := range resp.PrivateGateways {
				v, _ := VPCs.FindByID(pgw.Vpcid)
				pgws = append(pgws, &PrivateGateway{
					PrivateGateway: pgw,
					Vpccidr:        v[0].Cidr,
					Vpcname:        v[0].Name,
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

	return pgws, nil
}
