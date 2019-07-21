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
	"github.com/spf13/cobra"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/cosmic"
)

func newACLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acl",
		Short: "ACL subcommands",
	}

	// Add subcommands.
	cmd.AddCommand(newACLListCmd())

	// Add subgroups.
	cmd.AddCommand(newACLRuleCmd())

	return cmd
}

func getACL(cfg *config.Config) ([]*cosmic.ACL, error) {
	acl := []*cosmic.ACL{}
	// var err error

	acls, err := cosmic.ListACLs(cosmic.NewAsyncClients(cfg))
	if err != nil {
		return acls, err
	}

	switch {
	case cfg.ACLID != "":
		acl, err = acls.FindByID(cfg.ACLID)
	case cfg.ACLName != "":
		acl, err = acls.FindByName(cfg.ACLName)
	case cfg.InstanceID != "":
		vms, e := cosmic.ListVMs(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		vm, e := vms.FindByID(cfg.InstanceID)
		if e != nil {
			return nil, e
		}
		nets, e := cosmic.ListNetworks(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		net, e := nets.FindByID(vm[0].Nic[0].Networkid)
		if e != nil {
			return nil, e
		}
		acl, err = acls.FindByID(net[0].Aclid)
	case cfg.InstanceName != "":
		vms, e := cosmic.ListVMs(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		vm, e := vms.FindByName(cfg.InstanceName)
		if e != nil {
			return nil, e
		}
		nets, e := cosmic.ListNetworks(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		net, e := nets.FindByID(vm[0].Nic[0].Networkid)
		if e != nil {
			return nil, e
		}
		acl, err = acls.FindByID(net[0].Aclid)
	case cfg.NetworkID != "":
		nets, e := cosmic.ListNetworks(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		net, e := nets.FindByID(cfg.NetworkID)
		if e != nil {
			return nil, e
		}
		acl, err = acls.FindByID(net[0].Aclid)
	case cfg.NetworkName != "":
		nets, e := cosmic.ListNetworks(cosmic.NewAsyncClients(cfg))
		if e != nil {
			return nil, e
		}
		net, e := nets.FindByName(cfg.NetworkName)
		if e != nil {
			return nil, e
		}
		acl, err = acls.FindByID(net[0].Aclid)
	}

	return acl, nil
}
