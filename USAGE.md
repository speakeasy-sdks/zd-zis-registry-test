<!-- Start SDK Example Usage -->


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
    operationSecurity := operations.DeleteAPIServicesZisRegistryJobSpecsInstallSecurity{
            Password: "",
            Username: "",
        }

    ctx := context.Background()
    res, err := s.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall.DeleteAPIServicesZisRegistryJobSpecsInstall(ctx, operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest{
        JobSpecName: zendeskzisregistry.String("corrupti"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAPIServicesZisRegistryJobSpecsInstall204TextPlainString != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->