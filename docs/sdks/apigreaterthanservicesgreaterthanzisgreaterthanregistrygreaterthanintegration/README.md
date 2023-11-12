# APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration
(*APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration*)

### Available Operations

* [PostAPIServicesZisRegistryIntegration](#postapiserviceszisregistryintegration) - Create Integration

## PostAPIServicesZisRegistryIntegration

Creates an integration using the specified `integration` parameter.

#### Authentication

You can authorize requests using basic authentication or an API token.

#### Allowed for

* Admins

#### Deleting an integration

You can't delete an integration. However, you can disable it by uninstalling
its job spec. See the [Uninstall
Job Spec](/api-reference/integration-services/registry/jobspecs/#uninstall-job-spec)
endpoint.

### Example Usage

```go
package main

import(
	"context"
	"log"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v2"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/operations"
)

func main() {
    s := zdzisregistrytest.New(
        zdzisregistrytest.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration.PostAPIServicesZisRegistryIntegration(ctx, operations.PostAPIServicesZisRegistryIntegrationRequest{
        RequestBody: &operations.PostAPIServicesZisRegistryIntegrationRequestBody{},
        Integration: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PostAPIServicesZisRegistryIntegrationRequest](../../pkg/models/operations/postapiserviceszisregistryintegrationrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.PostAPIServicesZisRegistryIntegrationResponse](../../pkg/models/operations/postapiserviceszisregistryintegrationresponse.md), error**
| Error Object                                                                                                                                     | Status Code                                                                                                                                      | Content Type                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| sdkerrors.PostAPIServicesZisRegistryIntegrationResponseBody                                                                                      | 400                                                                                                                                              | application/json                                                                                                                                 |
| sdkerrors.PostAPIServicesZisRegistryIntegrationAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationResponseBody         | 401                                                                                                                                              | application/json                                                                                                                                 |
| sdkerrors.PostAPIServicesZisRegistryIntegrationAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationResponseResponseBody | 500                                                                                                                                              | application/json                                                                                                                                 |
| sdkerrors.SDKError                                                                                                                               | 400-600                                                                                                                                          | */*                                                                                                                                              |
