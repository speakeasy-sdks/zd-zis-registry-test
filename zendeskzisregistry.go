// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package zdzisregistrytest

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/internal/hooks"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://d3v-found1259.zendesk.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// ZendeskZisRegistry - ZIS Registry: The ZIS Registry Service manages ZIS integrations and their bundles. It provides an API to manage your integration, job specs, and bundles.
type ZendeskZisRegistry struct {
	APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall    *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall
	APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration                   *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration
	APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*ZendeskZisRegistry)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *ZendeskZisRegistry) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *ZendeskZisRegistry {
	sdk := &ZendeskZisRegistry{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "3.2.3",
			GenVersion:        "2.272.7",
			UserAgent:         "speakeasy-sdk/go 3.2.3 2.272.7 1.0.0 github.com/speakeasy-sdks/zd-zis-registry-test",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.DefaultClient = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.DefaultClient)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall = newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall(sdk.sdkConfiguration)

	sdk.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration = newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration(sdk.sdkConfiguration)

	sdk.APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles = newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles(sdk.sdkConfiguration)

	return sdk
}
