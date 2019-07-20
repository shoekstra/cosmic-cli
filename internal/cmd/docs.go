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
	"os"

	"github.com/spf13/cobra"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/docs"
)

func newDocsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "docs",
		Short:  "Generate documentation as markdown",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runDocsCmd(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func runDocsCmd(args []string) error {
	path := "./markdown"
	if len(args) > 0 {
		path = args[0]
	}
	if err := docs.GenMarkdownTree(NewCosmicCLICmd(), path); err != nil {
		return err
	}

	return nil
}
