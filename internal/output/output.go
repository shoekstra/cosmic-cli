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

package output

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	. "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

func PrintTable(cosmicType string, fields []string, output interface{}) {
	slice := InterfaceSlice(output)

	if len(slice) == 0 {
		fmt.Printf("Found 0 %s.\n", cosmicType)
		return
	}

	sort.Strings(fields)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader(fields)

	for _, s := range slice {
		row := []string{}
		val := reflect.Indirect(reflect.ValueOf(s))
		for i := 0; i < val.NumField(); i++ {
			if ContainsNoSpaces(fields, val.Type().Field(i).Name) {
				row = append(row, fmt.Sprintf("%v", val.Field(i).Interface()))
			}
		}
		table.Append(row)
	}

	table.Render()
	fmt.Printf("Found %d %s.\n", len(slice), cosmicType)
}

func Print(outputType, cosmicType string, fields []string, output interface{}) {
	switch {
	case strings.EqualFold(outputType, "table"):
		PrintTable(cosmicType, fields, output)
	}
}
