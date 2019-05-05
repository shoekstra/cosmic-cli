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

// StaticRoute embeds *cosmic.StaticRoute to allow additional fields.
type StaticRoute struct {
	*cosmic.StaticRoute
}

// StaticRoutes exists to provide helper methods for []*StaticRoute.
type StaticRoutes []*StaticRoute

// Sort will sort StaticRoutes by either the "cidr" or "nexthop" field.
func (srs StaticRoutes) Sort(sortBy string, reverseSort bool) {
	if h.Contains([]string{"cidr", "nexthop"}, sortBy) == false {
		fmt.Println("Invalid sort option provided, provide either \"cidr\" or \"nexthop\".")
		os.Exit(1)
	}

	switch {
	case strings.EqualFold(sortBy, "Cidr"):
		sort.SliceStable(srs, func(i, j int) bool {
			if reverseSort {
				return srs[i].Cidr > srs[j].Cidr
			}
			return srs[i].Cidr < srs[j].Cidr
		})
	case strings.EqualFold(sortBy, "Nexthop"):
		sort.SliceStable(srs, func(i, j int) bool {
			if reverseSort {
				return srs[i].Nexthop > srs[j].Nexthop
			}
			return srs[i].Nexthop < srs[j].Nexthop
		})
	}
}

// CreateVPCRoute loops through all configured *cosmic.CosmicClient objects and adds a new
// VPC static route if the provided VPC ID is found.
func CreateVPCRoute(clientMap map[string]*cosmic.CosmicClient, vpcID, nextHop string, cidr string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VPC.NewCreateStaticRouteParams(cidr, nextHop, vpcID)
			if _, err := clientMap[client].VPC.CreateStaticRoute(params); err != nil {
				if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
					return
				}
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
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
			return err
		}
	}

	return nil
}

// DeleteVPCRoute loops through all configured *cosmic.CosmicClient objects and removes an
// existing VPC static route if the provided VPC ID is found.
func DeleteVPCRoute(clientMap map[string]*cosmic.CosmicClient, id string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VPC.NewDeleteStaticRouteParams(id)
			if _, err := clientMap[client].VPC.DeleteStaticRoute(params); err != nil {
				if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
					return
				}
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
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
			return err
		}
	}

	return nil
}

// ListVPCRoutes returns a StaticRoutes object using all configured *cosmic.CosmicClient objects.
func ListVPCRoutes(clientMap map[string]*cosmic.CosmicClient, vpcID string) (StaticRoutes, error) {
	srs := []*StaticRoute{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	errChannel := make(chan error, 1)
	finished := make(chan bool, 1)

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			params := clientMap[client].VPC.NewListStaticRoutesParams()
			params.SetVpcid(vpcID)
			resp, err := clientMap[client].VPC.ListStaticRoutes(params)
			if err != nil {
				if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
					return
				}
				errChannel <- profileError{fmt.Sprintf("Error returned using profile \"%s\": %s", client, err)}
				return
			}

			for _, sr := range resp.StaticRoutes {
				srs = append(srs, &StaticRoute{
					StaticRoute: sr,
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

	return srs, nil
}
