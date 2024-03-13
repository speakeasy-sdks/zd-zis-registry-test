# github.com/speakeasy-sdks/zd-zis-registry-test

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/zd-zis-registry-test
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md)

* [DeleteAPIServicesZisRegistryJobSpecsInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md#deleteapiserviceszisregistryjobspecsinstall) - Uninstall Job Spec
* [PostAPIServicesZisRegistryJobSpecsInstall](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanjobspecsgreaterthaninstall/README.md#postapiserviceszisregistryjobspecsinstall) - Install Job Spec

### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegration/README.md)

* [PostAPIServicesZisRegistryIntegration](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegration/README.md#postapiserviceszisregistryintegration) - Create Integration

### [APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegrationgreaterthanbundles/README.md)

* [PostAPIServicesZisRegistryIntegrationBundles](docs/sdks/apigreaterthanservicesgreaterthanzisgreaterthanregistrygreaterthanintegrationgreaterthanbundles/README.md#postapiserviceszisregistryintegrationbundles) - Upload or Update Bundle
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                                                                                                                                                             | Status Code                                                                                                                                                              | Content Type                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallResponseBody                                                                                                        | 400                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody            | 401                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody    | 404                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponse500ResponseBody | 500                                                                                                                                                                      | application/json                                                                                                                                                         |
| sdkerrors.SDKError                                                                                                                                                       | 4xx-5xx                                                                                                                                                                  | */*                                                                                                                                                                      |

### Example

```go
package main

import (
	"context"
	"errors"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {

		var e *sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponse500ResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `http://d3v-found1259.zendesk.com` | None |

#### Example

```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithServerIndex(0),
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithServerURL("http://d3v-found1259.zendesk.com"),
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name       | Type       | Scheme     |
| ---------- | ---------- | ---------- |
| `Password` | http       | HTTP Basic |
| `Username` | http       | HTTP Basic |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := zdzisregistrytest.New(
		zdzisregistrytest.WithSecurity(shared.Security{
			Password: zdzisregistrytest.String("<YOUR_PASSWORD_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v3"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"log"
)

func main() {
	s := zdzisregistrytest.New()

	operationSecurity := operations.PostAPIServicesZisRegistryIntegrationSecurity{
		Password: "<YOUR_PASSWORD_HERE>",
		Username: "<YOUR_USERNAME_HERE>",
	}

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration.PostAPIServicesZisRegistryIntegration(ctx, operations.PostAPIServicesZisRegistryIntegrationRequest{
		RequestBody: &operations.PostAPIServicesZisRegistryIntegrationRequestBody{},
		Integration: "<value>",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
