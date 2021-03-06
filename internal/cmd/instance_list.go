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
	"os"

	"github.com/shoekstra/cosmic-cli/internal/config"
	"github.com/shoekstra/cosmic-cli/internal/cosmic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			viper.BindPFlag("show-version", cmd.Flags().Lookup("show-version"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runInstanceListCmd(); err != nil {
				printErr(err)
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
	cmd.Flags().BoolP("show-version", "", false, "show instance version in result")
	cmd.Flags().StringSliceP("filter", "f", nil, "filter results (supports regex)")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")
	cmd.Flags().StringP("sort-by", "s", "name", "field to sort by")

	cmd.Flags().MarkHidden("output") // Not in use yet.

	return cmd
}

func runInstanceListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	instances, err := cosmic.ListVMs(cosmic.NewAsyncClients(cfg))
	if err != nil {
		return err
	}
	instances.Sort(cfg.SortBy, cfg.ReverseSort)

	if cfg.ShowNetwork {
		networks, err := cosmic.ListNetworks(cosmic.NewAsyncClients(cfg))
		if err != nil {
			return err
		}
		vpcs, err := cosmic.ListVPCs(cosmic.NewAsyncClients(cfg))
		if err != nil {
			return err
		}
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

	// Print output
	fields := []string{"Name", "InstanceName", "State", "IPAddress", "ZoneName"}
	if cfg.ShowID {
		fields = append(fields, "ID")
	}
	if cfg.ShowHost {
		fields = append(fields, "Hostname")
	}
	if cfg.ShowNetwork {
		fields = append(fields, "NetworkName")
		fields = append(fields, "VPCName")
	}
	if cfg.ShowServiceOffering {
		fields = append(fields, "ServiceOfferingName")
	}
	if cfg.ShowTemplate {
		fields = append(fields, "TemplateName")
	}
	if cfg.ShowVersion {
		fields = append(fields, "Version")
	}
	printResult(cfg.Output, "instance", cfg.Filter, fields, instances)

	return nil
}
