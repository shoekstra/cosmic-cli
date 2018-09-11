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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/client"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/publicip"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/vpc"
)

func newVPCListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List VPCs",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("reverse-sort", cmd.Flags().Lookup("reverse-sort"))
			viper.BindPFlag("show-id", cmd.Flags().Lookup("show-id"))
			viper.BindPFlag("show-snat", cmd.Flags().Lookup("show-snat"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runVPCListCmd(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("reverse-sort", "", false, "reverse sort order")
	cmd.Flags().BoolP("show-id", "", false, "show VPC id in result")
	cmd.Flags().BoolP("show-snat", "", false, "show VPC Source NAT IP in result")
	cmd.Flags().StringP("filter", "f", "", "filter results (supports regex)")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile to limit results")
	cmd.Flags().StringP("sort-by", "s", "name", "field to sort by")

	return cmd
}

func runVPCListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	vpcs := vpc.ListAll(client.NewAsyncClientMap(cfg))
	vpcs.Sort(cfg.SortBy, cfg.ReverseSort)

	if cfg.ShowSNAT {
		publicIPs := publicip.ListAll(client.NewAsyncClientMap(cfg))
		for _, p := range publicIPs {
			if p.Issourcenat == false {
				continue
			}

			for _, v := range vpcs {
				if v.Id == p.Vpcid {
					v.Sourcenatip = p.Ipaddress
				}
			}
		}
	}

	// Print table
	fields := []string{"Name", "CIDR", "VPCOfferingName", "ZoneName"}
	if cfg.ShowID {
		fields = append(fields, "ID")
	}
	if cfg.ShowSNAT {
		fields = append(fields, "SourceNATIP")
	}
	printResult(cfg.Output, cfg.Filter, "VPC", fields, vpcs)

	return nil
}
