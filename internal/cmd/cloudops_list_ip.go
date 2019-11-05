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
	"fmt"
	"net"
	"os"

	"github.com/shoekstra/cosmic-cli/internal/config"
	"github.com/shoekstra/cosmic-cli/internal/cosmic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCloudOpsListIPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ip IPADDRESS",
		Short: "List IP details",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Bind local flags in the PreRun stage to not overwrite bindings in other commands.
			viper.BindPFlag("profile", cmd.Flags().Lookup("profile"))
			viper.BindPFlag("output", cmd.Flags().Lookup("output"))
			viper.BindPFlag("show-mac-address", cmd.Flags().Lookup("show-mac-address"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCloudOpsListIPCmd(args); err != nil {
				printErr(err)
				os.Exit(1)
			}
		},
	}

	// Add local flags.
	cmd.Flags().BoolP("show-mac-address", "", false, "show MAC address in result")
	cmd.Flags().StringP("output", "o", "table", "specify output type")
	cmd.Flags().StringP("profile", "p", "", "specify profile(s) to use")

	cmd.Flags().MarkHidden("output") // Not in use yet.

	return cmd
}

func runCloudOpsListIPCmd(args []string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// Validate args.
	ip, err := validateCloudOpsListIPArgs(args)
	if err != nil {
		return err
	}

	// Filter results so we only return the IP we're looking for.
	cfg.Filter = []string{fmt.Sprintf("%s=^%s$", "ipaddress", ip)}

	ips, err := cosmic.ListIP(cosmic.NewAsyncClients(cfg), ip)
	if err != nil {
		return err
	}

	// Print output
	fields := []string{"IPAddress", "NetMask", "VPCName", "Zonename"}
	if cfg.ShowMACAddress {
		fields = append(fields, "MACAddress")
	}
	// Some fields are populated only in certain conditions, so we only want to
	// have these columns if there is data to print.
	fNetworkName := false
	fVirtualMachineName := false
	for _, ip := range ips {
		if ip.Networkname != "" {
			fNetworkName = true
		}
		if ip.Virtualmachinename != "" {
			fVirtualMachineName = true
		}
	}
	if fNetworkName {
		fields = append(fields, "NetworkName")
	}
	if fVirtualMachineName {
		fields = append(fields, "VirtualMachineName")
	}

	printResult(cfg.Output, "IP Addresses", cfg.Filter, fields, ips)

	return nil
}

func validateCloudOpsListIPArgs(args []string) (string, error) {
	if len(args) != 1 {
		cmd := newCloudOpsListIPCmd()
		cmd.Help()
		os.Exit(0)
	}

	ip := net.ParseIP(args[0])
	if ip == nil {
		return "", fmt.Errorf("%s is not a valid IP address", args[0])
	}
	return ip.String(), nil
}
