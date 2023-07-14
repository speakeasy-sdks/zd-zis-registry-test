# APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall

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
	"context"
	"log"
	"github.com/speakeasy-sdks/zd-zis-registry-test"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
)

func main() {
    s := zendeskzisregistry.New()

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{
        JobSpecName: zendeskzisregistry.String("provident"),
    }, operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity{
        Password: "",
        Username: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAPIServicesZisRegistryJobSpecsInstall204TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest](../../models/operations/deleteapiserviceszisregistryjobspecsinstallrequest.md)   | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `security`                                                                                                                                       | [operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity](../../models/operations/deleteapiserviceszisregistryjobspecsinstallsecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |


### Response

**[*operations.DeleteAPIServicesZisRegistryJobSpecsInstallResponse](../../models/operations/deleteapiserviceszisregistryjobspecsinstallresponse.md), error**


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
	"context"
	"log"
	"github.com/speakeasy-sdks/zd-zis-registry-test"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
)

func main() {
    s := zendeskzisregistry.New()

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.PostAPIServicesZisRegistryJobSpecsInstall(ctx, operations.PostAPIServicesZisRegistryJobSpecsInstallRequest{
        JobSpecName: zendeskzisregistry.String("distinctio"),
    }, operations.PostAPIServicesZisRegistryJobSpecsInstallSecurity{
        Password: "",
        Username: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAPIServicesZisRegistryJobSpecsInstall200TextPlainString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PostAPIServicesZisRegistryJobSpecsInstallRequest](../../models/operations/postapiserviceszisregistryjobspecsinstallrequest.md)   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.PostAPIServicesZisRegistryJobSpecsInstallSecurity](../../models/operations/postapiserviceszisregistryjobspecsinstallsecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |


### Response

**[*operations.PostAPIServicesZisRegistryJobSpecsInstallResponse](../../models/operations/postapiserviceszisregistryjobspecsinstallresponse.md), error**

