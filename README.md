# go-nautobot

## :warning: Disclaimer :warning:

This project is in **beta** development stage, and it's suitable to change before being released as generally available. Use it at your own discretion.

## Introduction

`go-nautobot` package provides the Go Bindings to interact with [Nautobot Source of Truth](https://nautobot.readthedocs.io/en/stable/) API. Nautobot provides OpenAPI 3.0 specs and an `api_version` query parameter to specify the `major.minor` version to generate the schema from.

This package is auto-generated from Nautobot, and it comes with its own versioning schema, independent of Nautobot. For more details about versioning, check the [Release Versioning](#release-versioning) section.

## Customization

This package only generates the bindings for the Nautobot Core application, and not for the rich [apps ecosystem](https://docs.nautobot.com/projects/core/en/stable/apps/) around it. It's likely that the bindings in this package are not 100% correspond to your Nautobot environment, as you may have installed some public apps or your own homegrown ones.

Being aware of it, most often than not, you would need to generate your own bindings using the OpenAPI schema provided by your Nautobot deployment (with the installed apps). These are the steps to reproduce it:

1. Install `openapi-codegen`.

   ```bash
   $ wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar -O openapi-generator-cli.jar
   ```

2. Define your `openapi-codegen` config file. You can use the [one in this repository](development/oapi-config.yaml).

3. Download the `swagger.yaml` from the API, using the proper minor version:

   ```bash
   $ wget --header="Authorization: Token <your-nautobot-token>" http://<your_nautobot>/api/swagger.yaml?api_version=<major.minor_version> -O swagger.yaml
   ```

4. Generate the Go bindings

   ```bash
   $ export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
   $ java ./openapi-generator-cli.jar generate --config oapi-config.yaml \
    --input-spec swagger.yaml \
    --output ./ \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --http-user-agent go-nautobot/$(<major.minor_version>)
   ```

## Release Versioning

`go-nautobot` uses semantic versioning with a loose dependency on the Nautobot major-minor versioning used to generate the bindings.

The minor version in `go-nautobot` will always match the minor version from Nautobot, but the patch version will evolve independently.

For example, a fictitious release process would be:

| Nautobot | `go-nautobot`          |
| -------- | ---------------------- |
| 5.1.0    | 5.1.0                  |
| 5.1.1    | 5.1.1                  |
| 5.1.1\*  | 5.1.2 (some local fix) |
| 5.1.2    | 5.1.3                  |

> the "\*" case shows that when regenerating the bindings for the same Nautobot version, a increment on the patch release of `go-nautobot` occurs.

This way the project versioning will keep a relationship with the original minor Nautobot version while it also enables its own release lifecycle.

## How to use `go-nautobot` package

### Go Get Nautobot package

Simply point your `get` for the `nautobot` package to the version you require, in this example version `v2.3.0-beta`:

```bash
$ go get github.com/nautobot/go-nautobot/v2@v2.3.0-beta
```

### Go main example

```go
package main

import (
	"context"
	"fmt"

	nb "github.com/nautobot/go-nautobot/v2"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	token := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	
	config := nb.NewConfiguration()
	config.Servers[0].URL = "demo.nautobot.com/api"
	c := nb.NewAPIClient(config)
	
	ctx := context.Background()
	auth := context.WithValue(
		ctx,
		nb.ContextAPIKeys,
		map[string]nb.APIKey{
			"tokenAuth": {
				Key:    token,
				Prefix: "Token",
			},
		},
	)

	resp, _, err := c.DcimAPI.DcimManufacturersList(auth).Execute()
	check(err)

	fmt.Printf("%v", string(resp.Results))

}

```

### Other usage examples

- [Terraform Provider for Nautobot](https://github.com/nautobot/terraform-provider-nautobot/)

### Documentation
See the `docs` folder for documentation

## Local Development

> See the `Development` folder's README for instructions.

### Run tests locally

```
go test -v -gcflags="-e" ./...
```

> Hint: If you get a build fail during testing, check that you are not limiting container memory to 2GB. Upgrade to 4GB.

### Trigger the release manually

The release process can be triggered manually from GitHub Actions CLI (with the proper permissions):

```
gh workflow run release.yml -f tag=1.3.2
```

## TODO
- get rid of the openAPI spec patcher, this means either solving issues with the code generator or with the Nautobot API (see `development/scripts/fix-spec.py`)
- create tests for API functions

## Author
- @chadell Maintainer of go-nautobot
- @TobiPeterG (Tobias Görgens) brought up go-nautobot for Nautobot 2.x
