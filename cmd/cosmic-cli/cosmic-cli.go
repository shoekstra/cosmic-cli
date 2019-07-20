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
)

// NewCosmicCLICmd creates the `cosmic-cli` command and its subcommands.
func NewCosmicCLICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cosmic-cli",
		Short: "A CLI interface to manage Cosmic Cloud resources",
		Long: `
cosmic-cli is a CLI interface to manage Cosmic Cloud resources.

It aims to simplify administration of Cosmic Cloud resources by providing single commands for
actions that may require multiple API calls, whilst running commands against multiple API endpoints
in parallel.`,
	}

	// Add subcommands.
	cmd.AddCommand(newVersionCmd())

	// Add subgroups.
	cmd.AddCommand(newACLCmd())
	cmd.AddCommand(newCloudOpsCmd())
	cmd.AddCommand(newInstanceCmd())
	cmd.AddCommand(newVPCCmd())

	return cmd
}

func main() {
	cmd := NewCosmicCLICmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
