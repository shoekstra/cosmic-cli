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

package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/shoekstra/cosmic-cli/internal/config"
	"github.com/shoekstra/cosmic-cli/internal/cosmic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newVPCRouteAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add CIDR[,CIDR,CIDR] via NEXTHOP",
		Short: "Add VPC routes",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("vpc-id", cmd.Flags().Lookup("vpc-id"))
			viper.BindPFlag("vpc-name", cmd.Flags().Lookup("vpc-name"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runVPCRouteAddCmd(args); err != nil {
				printErr(err)
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

func runVPCRouteAddCmd(args []string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// Validate args.
	if err := validateVPCRouteAddArgs(args); err != nil {
		return err
	}

	// Validate the config.
	if err := validateVPCRouteAddCmd(cfg); err != nil {
		return err
	}

	// Get a list of existing routes.
	v, err := getVPC(cfg)
	if err != nil {
		return err
	}
	routes, err := cosmic.ListVPCRoutes(cosmic.NewAsyncClients(cfg), v.Id)
	if err != nil {
		return err
	}

	// Add routes to VPC.
	nextHop := args[2]
	newCidrs := []string{}
Loop:
	for _, cidr := range strings.Split(args[0], ",") {
		for _, r := range routes {
			// Don't try to add the route if it exists.
			if cidr == r.Cidr {
				fmt.Printf("Route already exists cidr:%s, nexthop:%s \n", r.Cidr, r.Nexthop)
				continue Loop
			}
		}
		newCidrs = append(newCidrs, cidr)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(newCidrs))

	for _, cidr := range newCidrs {
		go func(cidr string) error {
			defer wg.Done()

			fmt.Printf("Creating route cidr:%s, nexthop:%s ... \n", cidr, nextHop)
			if err := cosmic.CreateVPCRoute(cosmic.NewAsyncClients(cfg), v.Id, nextHop, cidr); err != nil {
				return err
			}

			return nil
		}(cidr)
	}
	wg.Wait()

	return nil
}

func validateVPCRouteAddArgs(args []string) error {
	if len(args) == 0 {
		cmd := newVPCRouteAddCmd()
		cmd.Help()
		os.Exit(0)
	}

	if len(args) != 3 {
		return errors.New("Incorrect number of parameters passed, this command expects \"<network> via <nexthop>\"")
	}

	if !strings.EqualFold(args[1], "via") {
		return errors.New("Invalid parameters passed, this command expects \"<network> via <nexthop>\"")
	}

	for _, c := range strings.Split(args[0], ",") {
		if _, _, err := net.ParseCIDR(c); err != nil {
			return fmt.Errorf("%s is not a valid network CIDR", c)
		}
	}

	if ip := net.ParseIP(args[2]); ip == nil {
		return fmt.Errorf("%s is not a valid IP address", args[2])
	}

	return nil
}

func validateVPCRouteAddCmd(cfg *config.Config) error {
	cmd := newVPCRouteAddCmd()

	if cfg.VPCID != "" && cfg.VPCName != "" {
		return errors.New("Cannot specify --vpc-id and --vpc-name together")
	}

	if cfg.VPCID == "" && cfg.VPCName == "" {
		cmd.Help()
		os.Exit(0)
	}

	return nil
}
