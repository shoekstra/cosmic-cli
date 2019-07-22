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

package config

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Config contains cosmic-cli options.
type Config struct {
	ACLID               string   `mapstructure:"acl-id"`
	ACLName             string   `mapstructure:"acl-name"`
	Filter              []string `mapstructure:"filter"`
	InstanceID          string   `mapstructure:"instance-id"`
	InstanceName        string   `mapstructure:"instance-name"`
	NetworkID           string   `mapstructure:"network-id"`
	NetworkName         string   `mapstructure:"network-name"`
	Output              string   `mapstructure:"output"`
	Profile             string   `mapstructure:"profile"`
	ReverseSort         bool     `mapstructure:"reverse-sort"`
	ShowDescription     bool     `mapstructure:"show-description"`
	ShowHost            bool     `mapstructure:"show-host"`
	ShowID              bool     `mapstructure:"show-id"`
	ShowMACAddress      bool     `mapstructure:"show-mac-address"`
	ShowACLID           bool     `mapstructure:"show-acl-id"`
	ShowACLName         bool     `mapstructure:"show-acl-name"`
	ShowRuleNumber      bool     `mapstructure:"show-rule-number"`
	ShowNetwork         bool     `mapstructure:"show-network"`
	ShowRedundantStatus bool     `mapstructure:"show-redundant-status"`
	ShowRestartRequired bool     `mapstructure:"show-restart-required"`
	ShowSNAT            bool     `mapstructure:"show-snat"`
	ShowServiceOffering bool     `mapstructure:"show-service-offering"`
	ShowTemplate        bool     `mapstructure:"show-template"`
	ShowVersion         bool     `mapstructure:"show-version"`
	SortBy              string   `mapstructure:"sort-by"`
	VPCID               string   `mapstructure:"vpc-id"`
	VPCName             string   `mapstructure:"vpc-name"`
	Profiles            map[string]struct {
		APIURL    string `mapstructure:"api_url"`
		APIKey    string `mapstructure:"api_key"`
		SecretKey string `mapstructure:"secret_key"`
	}
}

// ValidProfile will check the profile exists.
func (c *Config) ValidProfile(p string) {
	if _, ok := c.Profiles[p]; ok == false {
		log.Fatalf("Cannot find config for specified profile \"%v\"", p)
	}
}

// CheckDuplicatedProfile will check for any profiles; if any find an error will be returned asking
// the user to fix their config before proceeding.
func (c *Config) CheckDuplicatedProfile() {
	duplicate := false

	for p, q := range c.Profiles {
		if _, ok := c.Profiles[p]; ok {
			for k, v := range c.Profiles {
				if _, ok := c.Profiles[k]; ok {
					if p == k {
						continue
					}

					if q == v {
						fmt.Printf("Duplicate profiles found: \"%s\" is a duplicate of \"%s\" \n", k, p)
						delete(c.Profiles, k)
						duplicate = true
					}
				}
			}
		}
	}

	if duplicate {
		log.Fatal("Please remove duplicate profiles before continuing.")
	}
}

// New returns an initialized Config.
func New() (*Config, error) {
	viper.AddConfigPath(configPath())
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	// Try to read in a config file, but do not check for errors
	// as a config file is not mandatory for all commands.
	_ = viper.ReadInConfig()

	// Unmarshal the resulting config into our Config struct.
	cfg := &Config{}
	err := viper.Unmarshal(cfg)

	// Check for any duplicate profiles in config file.
	cfg.CheckDuplicatedProfile()

	return cfg, err
}

func configPath() string {
	configPath, err := homedir.Expand("~/.cosmic-cli")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return configPath
}
