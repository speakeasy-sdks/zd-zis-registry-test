// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package zendeskzisregistry

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/zd-zis-registry-test/pkg/utils"
	"io"
	"net/http"
)

type apiGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles struct {
	sdkConfiguration sdkConfiguration
}

func newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles(sdkConfig sdkConfiguration) *apiGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles {
	return &apiGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles{
		sdkConfiguration: sdkConfig,
	}
}

// PostAPIServicesZisRegistryIntegrationBundles - Upload or Update Bundle
// Uploads a new bundle or updates an existing bundle for a ZIS integration.
//
// #### Authentication
//
// You can authorize requests using basic authentication or an API token.
//
// #### Allowed for
//
// * Admins
//
// #### Deleting a bundle
//
// You can't delete a bundle. To disable an integration, uninstall its
// job spec. See the [Uninstall
// Job Spec](/api-reference/integration-services/registry/jobspecs/#uninstall-job-spec)
// endpoint.
func (s *apiGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundles) PostAPIServicesZisRegistryIntegrationBundles(ctx context.Context, request operations.PostAPIServicesZisRegistryIntegrationBundlesRequest, security operations.PostAPIServicesZisRegistryIntegrationBundlesSecurity) (*operations.PostAPIServicesZisRegistryIntegrationBundlesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/services/zis/registry/{integration}/bundles", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0.8, application/json;q=0.5, text/plain;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIServicesZisRegistryIntegrationBundlesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			out := string(rawBody)
			res.PostAPIServicesZisRegistryIntegrationBundles200TextPlainString = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostAPIServicesZisRegistryIntegrationBundles400ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostAPIServicesZisRegistryIntegrationBundles400ApplicationJSONObject = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostAPIServicesZisRegistryIntegrationBundles401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostAPIServicesZisRegistryIntegrationBundles401ApplicationJSONObject = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostAPIServicesZisRegistryIntegrationBundles500ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostAPIServicesZisRegistryIntegrationBundles500ApplicationJSONObject = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
