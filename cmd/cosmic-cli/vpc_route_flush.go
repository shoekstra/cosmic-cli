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

package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/client"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/vpc/route"
)

func newVPCRouteFlushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flush",
		Short: "Flush VPC routes",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("vpc-id", cmd.Flags().Lookup("vpc-id"))
			viper.BindPFlag("vpc-name", cmd.Flags().Lookup("vpc-name"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runVPCRouteFlushCmd(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")
	cmd.Flags().StringP("vpc-id", "", "", "specify VPC id")
	cmd.Flags().StringP("vpc-name", "", "", "specify VPC name")

	return cmd
}

func runVPCRouteFlushCmd(args []string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// Validate the config.
	if err := validateVPCRouteFlushCmd(cfg); err != nil {
		return err
	}

	// Get a list of existing routes.
	v, err := getVPC(cfg)
	if err != nil {
		return err
	}
	routes, err := route.List(client.NewAsyncClientMap(cfg), v.Id)
	if err != nil {
		return err
	}

	// Delete routes from VPC.
	wg := sync.WaitGroup{}
	wg.Add(len(routes))

	for _, r := range routes {
		go func(r *route.StaticRoute) error {
			defer wg.Done()

			fmt.Printf("Deleting route cidr:%s, nexthop:%s ... \n", r.Cidr, r.Nexthop)
			if err := route.Delete(client.NewAsyncClientMap(cfg), r.Id); err != nil {
				return err
			}

			return nil
		}(r)
	}
	wg.Wait()

	return nil
}

func validateVPCRouteFlushCmd(cfg *config.Config) error {
	cmd := newVPCRouteFlushCmd()

	if cfg.VPCID != "" && cfg.VPCName != "" {
		return errors.New("Cannot specify --vpc-id and --vpc-name together")
	}

	if cfg.VPCID == "" && cfg.VPCName == "" {
		cmd.Help()
		os.Exit(0)
	}

	return nil
}
