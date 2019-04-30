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

package cosmic

import (
	"sort"
	"strings"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	"github.com/forestgiant/sliceutil"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
)

// profileError represents a profile config error.
type profileError struct {
	message string
}

// Error returns the profile error message.
func (e profileError) Error() string {
	return e.message
}

// NewAsyncClients returns a [string]*cosmic.CosmicClient map.
func NewAsyncClients(cfg *config.Config) map[string]*cosmic.CosmicClient {
	profiles := getProfile(cfg)
	clientMap := make(map[string]*cosmic.CosmicClient)

	for _, profile := range profiles {
		clientMap[profile] = cosmic.NewAsyncClient(
			cfg.Profiles[profile].APIURL,
			cfg.Profiles[profile].APIKey,
			cfg.Profiles[profile].SecretKey,
			nil,
			int64(120),
		)
	}

	return clientMap
}

func getProfile(cfg *config.Config) []string {
	result := []string{}

	for p := range cfg.Profiles {
		result = append(result, p)
	}

	if cfg.Profile != "" {
		profiles := strings.Split(cfg.Profile, ",")

		// Check that the specified profiles actually exist
		for _, p := range profiles {
			cfg.ValidProfile(p)
		}

		// Remove any profiles out of scope
		for i := len(result) - 1; i >= 0; i-- {
			r := result[i]

			if sliceutil.Contains(profiles, r) == false {
				result = append(result[:i], result[i+1:]...)
			}
		}
	}

	sort.Strings(result)

	return result
}
