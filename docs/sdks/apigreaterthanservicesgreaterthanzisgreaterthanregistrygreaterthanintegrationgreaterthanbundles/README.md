# APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles
(*APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles*)

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
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"context"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"log"
)

func main() {
    s := zdzisregistrytest.New(
        zdzisregistrytest.WithSecurity(shared.Security{
            Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles.PostAPIServicesZisRegistryIntegrationBundles(ctx, operations.PostAPIServicesZisRegistryIntegrationBundlesRequest{
        RequestBody: &operations.PostAPIServicesZisRegistryIntegrationBundlesRequestBody{},
        Integration: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PostAPIServicesZisRegistryIntegrationBundlesRequest](../../pkg/models/operations/postapiserviceszisregistryintegrationbundlesrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.PostAPIServicesZisRegistryIntegrationBundlesResponse](../../pkg/models/operations/postapiserviceszisregistryintegrationbundlesresponse.md), error**
| Error Object                                                                                                                                                              | Status Code                                                                                                                                                               | Content Type                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| sdkerrors.PostAPIServicesZisRegistryIntegrationBundlesResponseBody                                                                                                        | 400                                                                                                                                                                       | application/json                                                                                                                                                          |
| sdkerrors.PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseBody         | 401                                                                                                                                                                       | application/json                                                                                                                                                          |
| sdkerrors.PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseResponseBody | 500                                                                                                                                                                       | application/json                                                                                                                                                          |
| sdkerrors.SDKError                                                                                                                                                        | 4xx-5xx                                                                                                                                                                   | */*                                                                                                                                                                       |
