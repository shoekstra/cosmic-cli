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
	"github.com/spf13/cobra"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic"
)

func newVPCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpc",
		Short: "VPC subcommands",
	}

	// Add subcommands.
	cmd.AddCommand(newVPCListCmd())

	// Add subgroups.
	cmd.AddCommand(newVPCPrivateGatewayCmd())
	cmd.AddCommand(newVPCRouteCmd())

	return cmd
}

func getVPC(cfg *config.Config) (*cosmic.VPC, error) {
	var err error
	vpcs := []*cosmic.VPC{}
	if cfg.VPCID != "" {
		vpcs, err = cosmic.VPCGetAllByID(
			cosmic.NewAsyncClients(cfg),
			cfg.VPCID,
		)
	}
	if cfg.VPCName != "" {
		vpcs, err = cosmic.VPCGetAllByName(
			cosmic.NewAsyncClients(cfg),
			cfg.VPCName,
		)
	}
	if err != nil {
		return nil, err
	}

	return vpcs[0], nil
}
