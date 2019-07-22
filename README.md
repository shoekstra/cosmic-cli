# cosmic-cli

A CLI interface to manage Cosmic Cloud resources written in Golang.

## Installation

A CI job uploads a new binary each time a branch is created or updated, be warned that these are bleeding edge and may also include bugs.

The following binaries are available from the master branch:

* `cosmic-cli-darwin-amd64` ([master/unreleased](https://sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/-/jobs/artifacts/master/download?job=build+darwin-amd64))
* `cosmic-cli-linux-amd64` ([master/unreleased](https://sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/-/jobs/artifacts/master/download?job=build+linux-amd64))
* `cosmic-cli-windows-amd64.exe` ([master/unreleased](https://sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/-/jobs/artifacts/master/download?job=build+windows-amd64))

Alternatively you can browse [published tags](https://sbp.gitlab.schubergphilis.com/shoekstra/cosmic-cli/tags) to download a released version that will adhere to SemVer.

## Usage

### Configuring profiles

`cosmic-cli` runs it's subcommands across multiple Cosmic zones, to use it you'll need to create a configuration file at `$HOME/.cosmic-cli/config.toml`.

For example:

```
[profiles.sbp-nl1]
api_url    = "https://nl1.mcc.schubergphilis.com/client/api"
api_key    = "CIB9_t..."
secret_key = "RnD3Kl..."

[profiles.sbp-nl1-admin]
api_url    = "https://admin-nl1.mcc.schubergphilis.com/client/api"
api_key    = "CIB9_t..."
secret_key = "RnD3Kl..."

[profiles.sbp-nl2]
api_url    = "https://nl2.mcc.schubergphilis.com/client/api"
api_key    = "BXpvK0..."
secret_key = "e1zI9w..."
```

### Using the filter

The format when using the filter is `field=value`, where `field` can be any table header. `value` is parsed as a regex for more flexible filtering.

For example: `-f name=vdi` will match any results where the name field contains "vdi". If you are searching for values that contain spaces, put the whole filter in quotes, e.g. `-f 'name=vdi 2'`.

You can use multiple filters by calling `-f` multiple times, e.g.

```
-f name=vdi -f vpcofferingname=default
```

Or you can pass a list, e.g.

```
-f name=vdi,vpcofferingname=default
```

### Command help

Help for a specific subcommand is available by running `cosmic-cli help <subcommand>` or `cosmic-cli <subcommand> -h`.

## License

```
Copyright 2019 Stephen Hoekstra <stephenhoekstra@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
