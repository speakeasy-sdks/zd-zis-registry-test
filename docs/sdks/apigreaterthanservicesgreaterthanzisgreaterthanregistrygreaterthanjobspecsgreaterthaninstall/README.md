# APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall
(*APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall*)

### Available Operations

* [DeleteAPIServicesZisRegistryJobSpecsInstall](#deleteapiserviceszisregistryjobspecsinstall) - Uninstall Job Spec
* [PostAPIServicesZisRegistryJobSpecsInstall](#postapiserviceszisregistryjobspecsinstall) - Install Job Spec

## DeleteAPIServicesZisRegistryJobSpecsInstall

Uninstalls a job spec for a ZIS integration.

#### Authentication

You can authorize requests using basic authentication, an API token, or a [ZIS OAuth access
token](/documentation/integration-services/developer-guide/developing-private-zis-integrations/#obtaining-a-zis-oauth-token).
A Zendesk app can also authorize requests to this endpoint using an
admin's browser session. See [Making API requests from a Zendesk
app](/documentation/apps/getting-started/making-api-requests-from-a-zendesk-app/#method-1-api-request-to-zendesk-api).

#### Allowed for

* Admins

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/shared"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v2"
	"context"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/operations"
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

    if res.TwoHundredAndFourTextPlainRes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest](../../pkg/models/operations/deleteapiserviceszisregistryjobspecsinstallrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.DeleteAPIServicesZisRegistryJobSpecsInstallResponse](../../pkg/models/operations/deleteapiserviceszisregistryjobspecsinstallresponse.md), error**
| Error Object                                                                                                                                                             | Status Code                                                                                                                                                              | Content Type                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallResponseBody                                                                                                        | 400                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody            | 401                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody    | 404                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponse500ResponseBody | 500                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.SDKError                                                                                                                                                       | 400-600                                                                                                                                                                  | */*                                                                                                                                                                      |

## PostAPIServicesZisRegistryJobSpecsInstall

Installs a job spec for a ZIS integration.

#### Authentication

You can authorize requests using basic authentication, an API token, or a [ZIS OAuth access
token](/documentation/integration-services/developer-guide/developing-private-zis-integrations/#obtaining-a-zis-oauth-token).
A Zendesk app can also authorize requests to this endpoint using an
admin's browser session. See [Making API requests from a Zendesk
app](/documentation/apps/getting-started/making-api-requests-from-a-zendesk-app/#method-1-api-request-to-zendesk-api).

#### Allowed for

* Admins

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/shared"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v2"
	"context"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/operations"
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
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.PostAPIServicesZisRegistryJobSpecsInstall(ctx, operations.PostAPIServicesZisRegistryJobSpecsInstallRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredTextPlainRes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PostAPIServicesZisRegistryJobSpecsInstallRequest](../../pkg/models/operations/postapiserviceszisregistryjobspecsinstallrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PostAPIServicesZisRegistryJobSpecsInstallResponse](../../pkg/models/operations/postapiserviceszisregistryjobspecsinstallresponse.md), error**
| Error Object                                                                                                                                                        | Status Code                                                                                                                                                         | Content Type                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallResponseBody                                                                                                     | 400                                                                                                                                                                 | application/json                                                                                                                                                    |
| sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody         | 401                                                                                                                                                                 | application/json                                                                                                                                                    |
| sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody | 500                                                                                                                                                                 | application/json                                                                                                                                                    |
| sdkerrors.SDKError                                                                                                                                                  | 400-600                                                                                                                                                             | */*                                                                                                                                                                 |
