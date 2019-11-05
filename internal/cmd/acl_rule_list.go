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
	"os"

	"github.com/shoekstra/cosmic-cli/internal/config"
	"github.com/shoekstra/cosmic-cli/internal/cosmic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newACLRuleListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List rules in an ACL",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("acl-id", cmd.Flags().Lookup("acl-id"))
			viper.BindPFlag("acl-name", cmd.Flags().Lookup("acl-name"))
			viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
			viper.BindPFlag("instance-id", cmd.Flags().Lookup("instance-id"))
			viper.BindPFlag("instance-name", cmd.Flags().Lookup("instance-name"))
			viper.BindPFlag("network-id", cmd.Flags().Lookup("network-id"))
			viper.BindPFlag("network-name", cmd.Flags().Lookup("network-name"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("reverse-sort", cmd.Flags().Lookup("reverse-sort"))
			viper.BindPFlag("show-acl-id", cmd.Flags().Lookup("show-acl-id"))
			viper.BindPFlag("show-acl-name", cmd.Flags().Lookup("show-acl-name"))
			viper.BindPFlag("show-id", cmd.Flags().Lookup("show-id"))
			viper.BindPFlag("show-rule-number", cmd.Flags().Lookup("show-rule-number"))
			viper.BindPFlag("sort-by", cmd.Flags().Lookup("sort-by"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runACLRuleListCmd(); err != nil {
				printErr(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("reverse-sort", "", false, "reverse sort order")
	cmd.Flags().BoolP("show-acl-id", "", false, "show ACL id in result")
	cmd.Flags().BoolP("show-acl-name", "", false, "show ACL name in result")
	cmd.Flags().BoolP("show-id", "", false, "show ACL rule id in result")
	cmd.Flags().BoolP("show-rule-number", "", false, "show ACL rule number in result")
	cmd.Flags().StringSliceP("filter", "f", nil, "filter results (supports regex)")
	cmd.Flags().StringP("acl-id", "", "", "specify ACL id")
	cmd.Flags().StringP("acl-name", "", "", "specify ACL name")
	cmd.Flags().StringP("instance-id", "", "", "specify instance id")
	cmd.Flags().StringP("instance-name", "", "", "specify instance name")
	cmd.Flags().StringP("network-id", "", "", "specify network id")
	cmd.Flags().StringP("network-name", "", "", "specify network name")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")
	cmd.Flags().StringP("sort-by", "s", "number", "field to sort by")

	cmd.Flags().MarkHidden("output") // Not in use yet.

	return cmd
}

func runACLRuleListCmd() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	if err := validateACLRuleListCmd(cfg); err != nil {
		return err
	}

	// Get ACLs based on options
	acls, err := getACL(cfg)
	if err != nil {
		return err
	}

	// Get ACL rules:
	// We loop over acls because if we provided an instance name of id, it may have multiple NICs/ACLs.
	rules := cosmic.ACLRules{}
	for _, acl := range acls {
		r, err := cosmic.ListACLRules(cosmic.NewAsyncClients(cfg), acl.Id)
		if err != nil {
			return err
		}
		for _, v := range r {
			rules = append(rules, &cosmic.ACLRule{
				Aclname:    v.Aclname,
				NetworkACL: v.NetworkACL,
				Vpcname:    v.Vpcname,
			})
		}
	}
	rules.Sort(cfg.SortBy, cfg.ReverseSort)

	// Print output
	fields := []string{"Action", "CidrList", "EndPort", "Icmpcode", "Icmptype", "Protocol", "StartPort", "TrafficType"}
	if cfg.ShowACLID {
		fields = append(fields, "ACLID")
	}
	if cfg.ShowACLName {
		fields = append(fields, "ACLName")
	}
	if cfg.ShowID {
		fields = append(fields, "ID")
	}
	if cfg.ShowRuleNumber {
		fields = append(fields, "Number")
	}
	printResult(cfg.Output, "ACL rule", cfg.Filter, fields, rules)

	return nil
}

func validateACLRuleListCmd(cfg *config.Config) error {
	switch {
	case cfg.ACLID != "" && cfg.ACLName != "":
		return errors.New("Cannot specify --acl-id and --acl-name together")
	case cfg.ACLID != "" && cfg.InstanceID != "":
		return errors.New("Cannot specify --acl-id and --instance-id together")
	case cfg.ACLID != "" && cfg.InstanceName != "":
		return errors.New("Cannot specify --acl-id and --instance-name together")
	case cfg.ACLID != "" && cfg.NetworkID != "":
		return errors.New("Cannot specify --acl-id and --network-id together")
	case cfg.ACLID != "" && cfg.NetworkName != "":
		return errors.New("Cannot specify --acl-id and --network-name together")
	case cfg.ACLName != "" && cfg.InstanceID != "":
		return errors.New("Cannot specify --acl-name and --instance-id together")
	case cfg.ACLName != "" && cfg.InstanceName != "":
		return errors.New("Cannot specify --acl-name and --instance-name together")
	case cfg.ACLName != "" && cfg.NetworkID != "":
		return errors.New("Cannot specify --acl-name and --network-id together")
	case cfg.ACLName != "" && cfg.NetworkName != "":
		return errors.New("Cannot specify --acl-name and --network-name together")
	case cfg.InstanceID != "" && cfg.InstanceName != "":
		return errors.New("Cannot specify --instance-id and --instance-name together")
	case cfg.InstanceID != "" && cfg.NetworkID != "":
		return errors.New("Cannot specify --instance-id and --network-id together")
	case cfg.InstanceID != "" && cfg.NetworkName != "":
		return errors.New("Cannot specify --instance-id and --network-name together")
	case cfg.InstanceName != "" && cfg.NetworkID != "":
		return errors.New("Cannot specify --instance-name and --network-id together")
	case cfg.InstanceName != "" && cfg.NetworkName != "":
		return errors.New("Cannot specify --instance-name and --network-name together")
	case cfg.NetworkID != "" && cfg.NetworkName != "":
		return errors.New("Cannot specify --network-id and --network-name together")
	}

	if cfg.ACLID == "" &&
		cfg.ACLName == "" &&
		cfg.InstanceID == "" &&
		cfg.InstanceName == "" &&
		cfg.NetworkID == "" &&
		cfg.NetworkName == "" {
		cmd := newACLRuleListCmd()
		cmd.Help()
		os.Exit(0)
	}

	return nil
}
