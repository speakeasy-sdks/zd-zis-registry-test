<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test/v2"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/operations"
	"log"
)

func main() {
	s := zdzisregistrytest.New()

	operationSecurity := operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity{
		Password: "",
		Username: "",
	}

	ctx := context.Background()
	res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.TwoHundredAndFourTextPlainRes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->