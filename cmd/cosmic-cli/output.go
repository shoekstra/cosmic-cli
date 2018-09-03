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
	. "sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/helper"
)

func filterMatch(obj interface{}, filter string) bool {
	if filter == "" {
		return true
	}

	if validFilter, _ := regexp.MatchString("=", filter); validFilter == false {
		fmt.Println("Invalid filter string passed, filters should be in the form of \"field=string\".")
		os.Exit(1)
	}

	split := strings.Split(filter, "=")
	filterField := strings.TrimSpace(split[0])
	filterString := strings.TrimSpace(split[1])

	// bval represents the base type if val is a nested type, this is only needed when we embed
	// an existing cosmic type with one of our own to add additional fields (e.g. cosmic.VPC is
	// nested within vpc.VPC)
	// var bval reflect.Value
	var val reflect.Value
	switch fmt.Sprintf("%s", reflect.TypeOf(obj)) {
	case "*vpc.VPC":
		val = reflect.Indirect(reflect.ValueOf(obj).Elem().FieldByName("VPC"))
	default:
		val = reflect.Indirect(reflect.ValueOf(obj))
	}

	if filterField == "ipaddress" {
		v := val.FieldByName("Nic")
		if v.IsValid() == false {
			return false
		}
		ip := fmt.Sprintf("%v", v.Index(0).FieldByName("Ipaddress"))
		match, _ := regexp.MatchString(strings.ToLower(filterString), ip)
		if match == true {
			return true
		}

		return false
	}

	f := strings.Title(strings.ToLower(filterField))
	v := val.FieldByName(f)
	if v.IsValid() == false {
		return false
	}
	value := fmt.Sprintf("%v", v.Interface())
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

	slice := InterfaceSlice(result)

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

func printTable(cosmicType string, fields []string, result interface{}) {
	slice := InterfaceSlice(result)

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
			fns := fmt.Sprintf("%s", strings.Replace(f, " ", "", -1))
			fns = strings.Title(strings.ToLower(fns))
			// We have some exceptions where the field name does not exist on the reflected object.
			switch fns {
			// *cosmic.VirtualMachine does not contain a "ipaddress" field so we need to manually
			// add the primary NIC IP to our table.
			case "Ipaddress":
				row = append(row, fmt.Sprintf("%v", val.FieldByName("Nic").Index(0).FieldByName("Ipaddress")))
			default:
				row = append(row, fmt.Sprintf("%v", val.FieldByName(fns).Interface()))
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

func printResult(outputType, filter, cosmicType string, fields []string, result interface{}) {
	result = filterOutput(result, filter)
	switch {
	case strings.EqualFold(outputType, "table"):
		printTable(cosmicType, fields, result)
	}
}