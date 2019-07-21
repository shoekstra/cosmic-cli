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

package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic"
)

func newVPCRouteListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List VPC routes",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("reverse-sort", cmd.Flags().Lookup("reverse-sort"))
			viper.BindPFlag("show-id", cmd.Flags().Lookup("show-id"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
			viper.BindPFlag("vpc-id", cmd.Flags().Lookup("vpc-id"))
			viper.BindPFlag("vpc-name", cmd.Flags().Lookup("vpc-name"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runVPCRouteListCmd(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("reverse-sort", "", false, "reverse sort order")
	cmd.Flags().BoolP("show-id", "", false, "show VPC id in result")
	cmd.Flags().StringSliceP("filter", "f", nil, "filter results (supports regex)")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")
	cmd.Flags().StringP("sort-by", "s", "cidr", "field to sort by")
	cmd.Flags().StringP("vpc-id", "", "", "specify VPC id")
	cmd.Flags().StringP("vpc-name", "", "", "specify VPC name")

	cmd.Flags().MarkHidden("output") // Not in use yet.

	return cmd
}

func runVPCRouteListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	if err := validateVPCRouteListCmd(cfg); err != nil {
		return err
	}

	v, err := getVPC(cfg)
	if err != nil {
		return err
	}

	// Fetch list of routes and add the VPC name if the next hop is a private
	// gateway attached to a VPC.
	routes, err := cosmic.ListVPCRoutes(cosmic.NewAsyncClients(cfg), v.Id)
	if err != nil {
		return err
	}
	pgws, err := cosmic.ListVPCPrivateGateways(cosmic.NewAsyncClients(cfg))
	if err != nil {
		return err
	}
	for _, r := range routes {
		pgws := cosmic.PrivateGateways(pgws).FindByIPAddress(r.Nexthop)
		if len(pgws) > 0 {
			// Just in case we somehow got more than one VPC returned...
			vpcnames := []string{}
			for _, pgw := range pgws {
				vpcnames = append(vpcnames, pgw.Vpcname)
			}
			sort.Strings(vpcnames)
			r.Vpcname = strings.Join(vpcnames, ", ")
		} else {
			r.Vpcname = ""
		}
	}
	routes.Sort(cfg.SortBy, cfg.ReverseSort)

	// Print output
	fields := []string{"CIDR", "NextHop", "VPCName"}
	if cfg.ShowID {
		fields = append(fields, "ID")
	}
	printResult(cfg.Output, "static route", cfg.Filter, fields, routes)

	return nil
}

func validateVPCRouteListCmd(cfg *config.Config) error {
	cmd := newVPCRouteListCmd()

	if cfg.VPCID != "" && cfg.VPCName != "" {
		return errors.New("Cannot specify --vpc-id and --vpc-name together")
	}

	if cfg.VPCID == "" && cfg.VPCName == "" {
		cmd.Help()
		os.Exit(0)
	}

	return nil
}
