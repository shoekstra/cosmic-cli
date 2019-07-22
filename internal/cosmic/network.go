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
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
)

// Network embeds *cosmic.Network to allow additional fields.
type Network struct {
	*cosmic.Network
}

// Networks exists to provide helper methods for []*Network.
type Networks []*Network

// FindByID looks for a Network object by ID in Networks and returns it if it exists.
func (n Networks) FindByID(id string) ([]*Network, error) {
	r := []*Network{}
	for _, v := range n {
		if v.Id == id {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for network with id %s", id)
	}
	return r, nil
}

// FindByName looks for a Network object by name in Networks and returns it if it exists.
func (n Networks) FindByName(name string) ([]*Network, error) {
	r := []*Network{}
	for _, v := range n {
		if v.Name == name {
			r = append(r, v)
		}
	}
	if len(r) == 0 {
		return nil, fmt.Errorf("No match found for network with name %s", name)
	}
	if len(r) > 1 {
		return r, fmt.Errorf("More than one match found for network with name %s, use the --network-id option to specify the network", name)
	}
	return r, nil
}

// ListNetworks returns a Networks object using all configured *cosmic.CosmicClient objects.
func ListNetworks(clientMap map[string]*cosmic.CosmicClient) (Networks, error) {
	networks := []*Network{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].Network.NewListNetworksParams()
			resp, err := clientMap[client].Network.ListNetworks(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, n := range resp.Networks {
				networks = append(networks, &Network{
					Network: n,
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

	return networks, nil
}
