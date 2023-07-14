# APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles

### Available Operations

* [PostAPIServicesZisRegistryIntegrationBundles](#postapiserviceszisregistryintegrationbundles) - Upload or Update Bundle

## PostAPIServicesZisRegistryIntegrationBundles

Uploads a new bundle or updates an existing bundle for a ZIS integration.

#### Authentication

You can authorize requests using basic authentication or an API token.

#### Allowed for

* Admins

#### Deleting a bundle

You can't delete a bundle. To disable an integration, uninstall its
job spec. See the [Uninstall
Job Spec](/api-reference/integration-services/registry/jobspecs/#uninstall-job-spec)
endpoint.

### Example Usage

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

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles.PostAPIServicesZisRegistryIntegrationBundles(ctx, operations.PostAPIServicesZisRegistryIntegrationBundlesRequest{
        RequestBody: &operations.PostAPIServicesZisRegistryIntegrationBundlesRequestBody{},
        Integration: "unde",
    }, operations.PostAPIServicesZisRegistryIntegrationBundlesSecurity{
        Password: "",
        Username: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPIServicesZisRegistryIntegrationBundles200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.PostAPIServicesZisRegistryIntegrationBundlesRequest](../../models/operations/postapiserviceszisregistryintegrationbundlesrequest.md)   | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `security`                                                                                                                                         | [operations.PostAPIServicesZisRegistryIntegrationBundlesSecurity](../../models/operations/postapiserviceszisregistryintegrationbundlessecurity.md) | :heavy_check_mark:                                                                                                                                 | The security requirements to use for the request.                                                                                                  |


### Response

**[*operations.PostAPIServicesZisRegistryIntegrationBundlesResponse](../../models/operations/postapiserviceszisregistryintegrationbundlesresponse.md), error**

