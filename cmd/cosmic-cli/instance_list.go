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
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/instance"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/network"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic/vpc"
)

func newInstanceListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List instances",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("reverse-sort", cmd.Flags().Lookup("reverse-sort"))
			viper.BindPFlag("show-host", cmd.Flags().Lookup("show-host"))
			viper.BindPFlag("show-id", cmd.Flags().Lookup("show-id"))
			viper.BindPFlag("show-network", cmd.Flags().Lookup("show-network"))
			viper.BindPFlag("show-service-offering", cmd.Flags().Lookup("show-service-offering"))
			viper.BindPFlag("show-template", cmd.Flags().Lookup("show-template"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runInstanceListCmd(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("reverse-sort", "", false, "reverse sort order")
	cmd.Flags().BoolP("show-host", "", false, "show hypervisor hostname in result")
	cmd.Flags().BoolP("show-id", "", false, "show instance id in result")
	cmd.Flags().BoolP("show-network", "", false, "show network info in result")
	cmd.Flags().BoolP("show-service-offering", "", false, "show instance service offering in result")
	cmd.Flags().BoolP("show-template", "", false, "show instance template name in result")
	cmd.Flags().StringP("filter", "f", "", "filter results (supports regex)")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile to limit results")
	cmd.Flags().StringP("sort-by", "s", "name", "field to sort by")

	return cmd
}

func runInstanceListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	instances := instance.ListAll(client.NewAsyncClientMap(cfg))
	instances.Sort(cfg.SortBy, cfg.ReverseSort)

	if cfg.ShowNetwork {
		networks := network.ListAll(client.NewAsyncClientMap(cfg))
		vpcs := vpc.ListAll(client.NewAsyncClientMap(cfg))

		for _, i := range instances {
			vpcid := ""
			for _, n := range networks {
				if n.Id == i.Nic[0].Networkid {
					i.Networkname = n.Name
					vpcid = n.Vpcid
					break
				}
			}

			for _, v := range vpcs {
				if v.Id == vpcid {
					i.Vpcname = v.Name
					break
				}
			}
		}
	}

	// Print table
	fields := []string{"Name", "InstanceName", "State", "IPAddress", "ZoneName"}
	if cfg.ShowID {
		fields = append(fields, "ID")
	}
	if cfg.ShowHost {
		fields = append(fields, "Hostname")
	}
	if cfg.ShowServiceOffering {
		fields = append(fields, "ServiceOfferingName")
	}
	if cfg.ShowTemplate {
		fields = append(fields, "TemplateName")
	}
	if cfg.ShowNetwork {
		fields = append(fields, "NetworkName")
		fields = append(fields, "VPCName")
	}
	printResult(cfg.Output, cfg.Filter, "instance", fields, instances)

	return nil
}
