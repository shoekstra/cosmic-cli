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

package network

import (
	"log"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	// . "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

// List returns a slice of *cosmic.Network objects using a *cosmic.CosmicClient object.
func List(client *cosmic.CosmicClient) ([]*cosmic.Network, error) {
	params := client.Network.NewListNetworksParams()
	resp, err := client.Network.ListNetworks(params)
	if err != nil {
		return resp.Networks, err
	}

	return resp.Networks, nil
}

func ListAll(clientMap map[string]*cosmic.CosmicClient, filter, sortBy string, reverseSort bool) []*cosmic.Network {
	networks := []*cosmic.Network{}
	var wg sync.WaitGroup
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listNetworks, err := List(clientMap[client])
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, n := range listNetworks {
				networks = append(networks, n)
			}
		}(client)
	}
	wg.Wait()

	// networks = Sort(networks, sortBy, reverseSort)

	return networks
}
