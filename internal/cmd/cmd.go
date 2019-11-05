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
	"regexp"

	"github.com/spf13/cobra"
)

// NewCosmicCLICmd creates the `cosmic-cli` command and its subcommands.
func NewCosmicCLICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cosmic-cli",
		Short: "A CLI interface to manage Cosmic Cloud resources",
		Long: `cosmic-cli is a CLI interface to manage Cosmic Cloud resources.

It aims to simplify administration of Cosmic Cloud resources by providing single commands for
actions that may require multiple API calls, whilst running commands against multiple API endpoints
in parallel.`,
		DisableAutoGenTag: true,
	}

	// Add subcommands.
	cmd.AddCommand(newDocsCmd())
	cmd.AddCommand(newVersionCmd())

	// Add subgroups.
	cmd.AddCommand(newACLCmd())
	cmd.AddCommand(newCloudOpsCmd())
	cmd.AddCommand(newInstanceCmd())
	cmd.AddCommand(newVPCCmd())

	return cmd
}

// printErr prints the error after santizing the output.
func printErr(err error) {
	s := err.Error()
	reapi := regexp.MustCompile(`apiKey=([aA0-zZ9%-]+)`)
	s = fmt.Sprint(reapi.ReplaceAllString(s, "apiKey=**redacted**"))
	resig := regexp.MustCompile(`signature=([aA0-zZ9%-]+):`)
	s = fmt.Sprint(resig.ReplaceAllString(s, "signature=**redacted**:"))

	fmt.Println(s)
}
