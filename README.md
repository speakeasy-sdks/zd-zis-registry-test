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

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/zd-zis-registry-test"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
)

func main() {
    s := zendeskzisregistry.New()
    operationSecurity := operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity{
            Password: "",
            Username: "",
        }

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{
        JobSpecName: zendeskzisregistry.String("corrupti"),
    }, operationSecurity)
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

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
