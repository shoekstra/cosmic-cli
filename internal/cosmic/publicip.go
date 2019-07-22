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

// PublicIPAddress embeds *cosmic.PublicIpAddress to allow additional fields.
type PublicIPAddress struct {
	*cosmic.PublicIpAddress
}

// PublicIPAddresses exists to provide helper methods for []*PublicIPAddress.
type PublicIPAddresses []*PublicIPAddress

// ListPublicIPAddresses returns a PublicIPAddresses object using all configured *cosmic.CosmicClient objects.
func ListPublicIPAddresses(clientMap map[string]*cosmic.CosmicClient) (PublicIPAddresses, error) {
	publicips := []*PublicIPAddress{}
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
				publicips = append(publicips, &PublicIPAddress{
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
