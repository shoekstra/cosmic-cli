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

package client

import (
	"crypto/tls"
	"sort"
	"strings"

	"github.com/MissionCriticalCloud/go-cosmic/cosmic"
	"github.com/forestgiant/sliceutil"
	"sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/internal/config"
)

// NewAsyncClientMap returns a [string]*cosmic.CosmicClient map.
func NewAsyncClientMap(cfg *config.Config) map[string]*cosmic.CosmicClient {
	profiles := getProfile(cfg)
	clientMap := make(map[string]*cosmic.CosmicClient)

	tlsConfig := &tls.Config{}
	httpTimeout := int64(60)

	for _, profile := range profiles {
		clientMap[profile] = cosmic.NewAsyncClient(
			cfg.Profiles[profile].APIURL,
			cfg.Profiles[profile].APIKey,
			cfg.Profiles[profile].SecretKey,
			tlsConfig,
			httpTimeout,
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
