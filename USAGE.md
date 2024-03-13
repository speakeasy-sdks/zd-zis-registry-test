<!-- Start SDK Example Usage [usage] -->
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