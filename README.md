# github.com/speakeasy-sdks/zd-zis-registry-test

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/zd-zis-registry-test
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: "",
			Username: "",
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.DeleteAPIServicesZisRegistryJobSpecsInstall204TextPlainString != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md)

* [DeleteAPIServicesZisRegistryJobSpecsInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md#deleteapiserviceszisregistryjobspecsinstall) - Uninstall Job Spec
* [PostAPIServicesZisRegistryJobSpecsInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md#postapiserviceszisregistryjobspecsinstall) - Install Job Spec

### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegration/README.md)

* [PostAPIServicesZisRegistryIntegration](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegration/README.md#postapiserviceszisregistryintegration) - Create Integration

### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegrationgreaterthanbundles/README.md)

* [PostAPIServicesZisRegistryIntegrationBundles](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegrationgreaterthanbundles/README.md#postapiserviceszisregistryintegrationbundles) - Upload or Update Bundle
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
