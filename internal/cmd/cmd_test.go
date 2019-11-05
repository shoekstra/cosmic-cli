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
	"errors"
)

func ExampleprintErr() {
	err := errors.New("Error returned using profile \"profile1\": Get https://api.cosmic.local/client/api/?apiKey=jDCMCLD8GGeSupR8rFyBRBRKX3AffGKVtycc6B6hjjFNb5D4-ThsU-KrnVJKxzBccTKLx2qArrymxT4xDevr6J&command=listVirtualMachines&response=json&signature=nx963U5Qv08Wm5ey2nRV0U%2B02m4%3D: dial tcp: lookup api.cosmic.local: no such host")
	printErr(err)

	// Output:
	// Error returned using profile "profile1": Get https://api.cosmic.local/client/api/?apiKey=**redacted**&command=listVirtualMachines&response=json&signature=**redacted**: dial tcp: lookup api.cosmic.local: no such host
}
