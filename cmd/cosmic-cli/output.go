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
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	h "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

func filterMatch(obj interface{}, filter string) bool {
	if filter == "" {
		return true
	}

	filterField, filterString := filterSplit(filter)

	value := ""
	val := reflect.Indirect(reflect.ValueOf(obj))
	switch {
	case strings.EqualFold(filterField, "ipaddress"):
		v := val.FieldByName("Nic")
		if v.IsValid() == true {
			value = fmt.Sprintf("%v", v.Index(0).FieldByName("Ipaddress"))
			break
		}
		fallthrough
	default:
		// First try to read the field directly; if it exists set the value and break.
		fn := strings.Title(strings.ToLower(filterField))
		valueField := val.FieldByName(fn)
		if valueField.IsValid() {
			value = fmt.Sprintf("%v", valueField.Interface())
			break
		}

		// If we can't read the field directly, we'll then loop through all field names and set
		// value if we find a matching field name.
		for i := 0; i < val.NumField(); i++ {
			typeField := val.Type().Field(i)

			if strings.EqualFold(filterField, typeField.Name) {
				valueField := val.Field(i)
				value = fmt.Sprintf("%v", valueField.Interface())
				break
			}
		}

		// At this point we'll assume the field doesn't exist and return 0 matches; if it has
		// some funky key then probably better to add as an exception to the switch statement.
		return false
	}

	match, _ := regexp.MatchString(strings.ToLower(filterString), strings.ToLower(value))
	if match == true {
		return true
	}

	return false
}

func filterOutput(result interface{}, filter string) interface{} {
	if filter == "" {
		return result
	}

	slice := h.InterfaceSlice(result)

	if len(slice) == 0 {
		return result
	}

	for i := 0; i < len(slice); i++ {
		if filterMatch(slice[i], filter) == false {
			slice = append(slice[:i], slice[i+1:]...)
			i-- // -1 as the slice just got shorter.
		}
	}

	return slice
}

func filterSplit(filter string) (field, value string) {
	if validFilter, _ := regexp.MatchString("=", filter); validFilter == false {
		fmt.Println("Invalid filter string passed, filters should be in the form of \"field=value\".")
		os.Exit(1)
	}

	split := strings.Split(filter, "=")
	filterField := strings.TrimSpace(split[0])
	filterString := strings.TrimSpace(split[1])

	return filterField, filterString
}

func printTable(cosmicType string, fields []string, result interface{}) {
	slice := h.InterfaceSlice(result)

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
		for _, f := range fields {
			fn := strings.Title(strings.ToLower(f))
			// We have some exceptions where the field name does not exist on the reflected object.
			switch fn {
			// *cosmic.VirtualMachine does not contain a "ipaddress" field so we need to manually
			// add the primary NIC IP to our table.
			case "Ipaddress":
				if cosmicType == "instance" {
					row = append(row, fmt.Sprintf("%v", val.FieldByName("Nic").Index(0).FieldByName("Ipaddress")))
					break
				}
				fallthrough
			default:
				v := val.FieldByName(fn)
				if v.IsValid() == false {
					break
				}
				row = append(row, fmt.Sprintf("%v", v))
			}
		}
		table.Append(row)
	}

	table.Render()

	if len(slice) > 1 {
		cosmicType = cosmicType + "s"
	}
	fmt.Printf("Found %d %s.\n", len(slice), cosmicType)
}

func printResult(outputType, cosmicType string, filter, fields []string, result interface{}) {
	for _, f := range filter {
		result = filterOutput(result, f)
	}
	switch {
	case strings.EqualFold(outputType, "table"):
		printTable(cosmicType, fields, result)
	}
}
