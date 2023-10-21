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