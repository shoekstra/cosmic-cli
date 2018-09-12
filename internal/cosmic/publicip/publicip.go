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

package publicip

import (
	"log"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
)

// Address embeds *cosmic.PublicIpAddress to allow additional fields.
type Address struct {
	*cosmic.PublicIpAddress
}

// Addresses exists to provide helper methods for []*Address.
type Addresses []*Address

// List returns a slice of *cosmic.PublicIpAddress objects using a *cosmic.CosmicClient object.
func List(client *cosmic.CosmicClient) ([]*cosmic.PublicIpAddress, error) {
	params := client.PublicIPAddress.NewListPublicIpAddressesParams()
	resp, err := client.PublicIPAddress.ListPublicIpAddresses(params)
	if err != nil {
		return resp.PublicIpAddresses, err
	}

	return resp.PublicIpAddresses, nil
}

// ListAll returns an Addresses object using all configured *cosmic.CosmicClient objects.
func ListAll(clientMap map[string]*cosmic.CosmicClient) Addresses {
	publicIPs := []*Address{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listpublicIPs, err := List(clientMap[client])
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, ip := range listpublicIPs {
				publicIPs = append(publicIPs, &Address{
					PublicIpAddress: ip,
				})
			}
		}(client)
	}
	wg.Wait()

	return publicIPs
}
