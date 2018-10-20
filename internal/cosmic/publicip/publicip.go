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
	"fmt"
	"sync"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
)

// profileError represents a profile config error.
type profileError struct {
	message string
}

// Error returns the profile error message.
func (e profileError) Error() string {
	return e.message
}

// Address embeds *cosmic.PublicIpAddress to allow additional fields.
type Address struct {
	*cosmic.PublicIpAddress
}

// Addresses exists to provide helper methods for []*Address.
type Addresses []*Address

// List returns an Addresses object using all configured *cosmic.CosmicClient objects.
func List(clientMap map[string]*cosmic.CosmicClient) (Addresses, error) {
	publicips := []*Address{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].PublicIPAddress.NewListPublicIpAddressesParams()
			resp, err := clientMap[client].PublicIPAddress.ListPublicIpAddresses(params)
			if err != nil {
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, ip := range resp.PublicIpAddresses {
				publicips = append(publicips, &Address{
					PublicIpAddress: ip,
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

	return publicips, nil
}
