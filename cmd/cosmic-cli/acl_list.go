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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic"
)

func newACLListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List ACLs",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("reverse-sort", cmd.Flags().Lookup("reverse-sort"))
			viper.BindPFlag("show-description", cmd.Flags().Lookup("show-description"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
			viper.BindPFlag("vpc-id", cmd.Flags().Lookup("vpc-id"))
			viper.BindPFlag("vpc-name", cmd.Flags().Lookup("vpc-name"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runACLListCmd(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("reverse-sort", "", false, "reverse sort order")
	cmd.Flags().BoolP("show-description", "", false, "show ACL description in result")
	cmd.Flags().StringSliceP("filter", "f", nil, "filter results (supports regex)")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")
	cmd.Flags().StringP("sort-by", "s", "vpcname", "field to sort by")
	cmd.Flags().StringP("vpc-id", "", "", "specify VPC id")
	cmd.Flags().StringP("vpc-name", "", "", "specify VPC name")

	return cmd
}

func runACLListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	if err := validateACLListCmd(cfg); err != nil {
		return err
	}

	acls, err := cosmic.ListACLs(cosmic.NewAsyncClients(cfg))
	if err != nil {
		return err
	}
	acls.Sort(cfg.SortBy, cfg.ReverseSort)

	// Print output
	fields := []string{"ID", "Name", "VPCName", "ZoneName"}
	if cfg.ShowDescription {
		fields = append(fields, "Description")
	}
	// Filter results if --vpc-id or --vpc-name flags are used
	if cfg.VPCID != "" {
		fields = append(fields, "VPCID")
		cfg.Filter = []string{fmt.Sprintf("%s=^%s$", "vpcid", cfg.VPCID)}
	}
	if cfg.VPCName != "" {
		cfg.Filter = []string{fmt.Sprintf("%s=^%s$", "vpcname", cfg.VPCName)}
	}
	printResult(cfg.Output, "ACL", cfg.Filter, fields, acls)

	return nil
}

func validateACLListCmd(cfg *config.Config) error {
	if cfg.VPCID != "" && cfg.VPCName != "" {
		return errors.New("Cannot specify --vpc-id and --vpc-name together")
	}

	return nil
}
