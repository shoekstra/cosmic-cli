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

package route

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

// Create adds a new VPC static route if the VPC exists.
func Create(client *cosmic.CosmicClient, vpcID, cidr, nextHop string) error {
	params := client.VPC.NewCreateStaticRouteParams(cidr, nextHop, vpcID)
	if _, err := client.VPC.CreateStaticRoute(params); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil
		}
		return err
	}

	return nil
}

// CreateAll loops through all configured *cosmic.CosmicClient objects and adds a new VPC static route
// if the provided VPC ID is found.
func CreateAll(clientMap map[string]*cosmic.CosmicClient, vpcID, nextHop string, cidr string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			err := Create(clientMap[client], vpcID, cidr, nextHop)
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}
		}(client)
	}
	wg.Wait()

	return nil
}

// Delete removes an existing VPC static route if the VPC exists.
func Delete(client *cosmic.CosmicClient, id string) error {
	params := client.VPC.NewDeleteStaticRouteParams(id)
	if _, err := client.VPC.DeleteStaticRoute(params); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil
		}
		return err
	}

	return nil
}

// DeleteAll loops through all configured *cosmic.CosmicClient objects and removes an existing VPC
// static route if the provided VPC ID is found.
func DeleteAll(clientMap map[string]*cosmic.CosmicClient, id string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			err := Delete(clientMap[client], id)
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}
		}(client)
	}
	wg.Wait()

	return nil
}

// List returns a slice of *cosmic.StaticRoute objects using a *cosmic.CosmicClient object.
func List(client *cosmic.CosmicClient, vpcID string) ([]*cosmic.StaticRoute, error) {
	params := client.VPC.NewListStaticRoutesParams()
	params.SetVpcid(vpcID)
	resp, err := client.VPC.ListStaticRoutes(params)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("entity does not exist")) {
			return nil, nil
		}
		return nil, err
	}

	if err != nil {
		return resp.StaticRoutes, err
	}

	return resp.StaticRoutes, nil
}

// ListAll returns a StaticRoutes object using all configured *cosmic.CosmicClient objects.
func ListAll(clientMap map[string]*cosmic.CosmicClient, vpcID string) StaticRoutes {
	srs := []*StaticRoute{}
	wg := sync.WaitGroup{}
	wg.Add(len(clientMap))

	for client := range clientMap {
		go func(client string) {
			defer wg.Done()

			listStaticRoutes, err := List(clientMap[client], vpcID)
			if err != nil {
				log.Fatalf("Error returned using profile \"%s\": %s", client, err)
			}

			for _, sr := range listStaticRoutes {
				srs = append(srs, &StaticRoute{
					StaticRoute: sr,
				})
			}
		}(client)
	}
	wg.Wait()

	return srs
}
