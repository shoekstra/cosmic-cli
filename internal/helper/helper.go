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
	"reflect"
	"strings"
)

// Contains return true if the slice contains an element matching the provided string.
func Contains(slice []string, str string) bool {
	for _, a := range slice {
		if strings.EqualFold(a, str) {
			return true
		}
	}

	return false
}

// ContainsNoSpaces return true if the slice contains an element matching the provided string, first
// removing any spaces from string elements in the slice.
func ContainsNoSpaces(slice []string, str string) bool {
	var sliceNoSpaces []string
	for _, a := range slice {
		sliceNoSpaces = append(sliceNoSpaces, fmt.Sprintf("%s", strings.Replace(a, " ", "", -1)))
	}

	return Contains(sliceNoSpaces, str)
}

// InterfaceSlice returns a slice of interfaces.
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
