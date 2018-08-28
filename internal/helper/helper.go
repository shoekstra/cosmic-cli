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

package helper

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func Contains(slice []string, str string) bool {
	for _, a := range slice {
		if strings.EqualFold(a, str) {
			return true
		}
	}

	return false
}

func ContainsNoSpaces(slice []string, str string) bool {
	var sliceNoSpaces []string
	for _, a := range slice {
		sliceNoSpaces = append(sliceNoSpaces, fmt.Sprintf("%s", strings.Replace(a, " ", "", -1)))
	}

	return Contains(sliceNoSpaces, str)
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func MatchFilter(obj interface{}, filter string) bool {
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

	valueOf := reflect.Indirect(reflect.ValueOf(obj))
	for i := 0; i < valueOf.NumField(); i++ {
		if strings.EqualFold(filterField, valueOf.Type().Field(i).Name) {
			match, _ := regexp.MatchString(
				strings.ToLower(filterString),
				strings.ToLower(fmt.Sprintf("%v", valueOf.Field(i).Interface())),
			)
			if match == true {
				return true
			}
		}
	}

	return false
}
