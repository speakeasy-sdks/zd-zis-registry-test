<!-- Start SDK Example Usage [usage] -->
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

	operationSecurity := operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity{
		Password: "<YOUR_PASSWORD_HERE>",
		Username: "<YOUR_USERNAME_HERE>",
	}

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->