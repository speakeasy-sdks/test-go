# github.com/speakeasy-sdks/test-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	testgo "github.com/speakeasy-sdks/test-go"
	"github.com/speakeasy-sdks/test-go/pkg/models/shared"
	"log"
)

func main() {
	s := testgo.New()

	ctx := context.Background()
	res, err := s.AcmeGo.CreateUserv1(ctx, shared.UserInput{
		Country:   "Benin",
		Email:     "Della67@yahoo.com",
		Firstname: "Enrique",
		Lastname:  "Ernser",
		Nickname:  "panel",
		Password:  "OXx1B29WwlhtAAe",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.User != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [AcmeGo SDK](docs/sdks/acmego/README.md)

* [CreateUserv1](docs/sdks/acmego/README.md#createuserv1) - Create user
* [DeleteUserv1](docs/sdks/acmego/README.md#deleteuserv1) - Delete a user by ID
* [GetHealth](docs/sdks/acmego/README.md#gethealth) - Healthcheck
* [GetUserv1](docs/sdks/acmego/README.md#getuserv1) - Get a user by ID
* [SearchUsersv1](docs/sdks/acmego/README.md#searchusersv1) - Search users
* [UpdateUserv1](docs/sdks/acmego/README.md#updateuserv1)
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
