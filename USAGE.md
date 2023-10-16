<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	zdzisregistrytest "github.com/speakeasy-sdks/zd-zis-registry-test"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
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

	if res.DeleteAPIServicesZisRegistryJobSpecsInstall204TextPlainString != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->