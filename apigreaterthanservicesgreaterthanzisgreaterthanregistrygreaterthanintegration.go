// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package zdzisregistrytest

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/internal/hooks"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v3/pkg/utils"
	"io"
	"net/http"
)

type APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration struct {
	sdkConfiguration sdkConfiguration
}

func newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration(sdkConfig sdkConfiguration) *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration {
	return &APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration{
		sdkConfiguration: sdkConfig,
	}
}

// PostAPIServicesZisRegistryIntegration - Create Integration
// Creates an integration using the specified `integration` parameter.
//
// #### Authentication
//
// You can authorize requests using basic authentication or an API token.
//
// #### Allowed for
//
// * Admins
//
// #### Deleting an integration
//
// You can't delete an integration. However, you can disable it by uninstalling
// its job spec. See the [Uninstall
// Job Spec](/api-reference/integration-services/registry/jobspecs/#uninstall-job-spec)
// endpoint.
func (s *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegration) PostAPIServicesZisRegistryIntegration(ctx context.Context, request operations.PostAPIServicesZisRegistryIntegrationRequest) (*operations.PostAPIServicesZisRegistryIntegrationResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "post_/api/services/zis/registry/{integration}",
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/services/zis/registry/{integration}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "401", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIServicesZisRegistryIntegrationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.PostAPIServicesZisRegistryIntegrationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.PostAPIServicesZisRegistryIntegrationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			out.RawResponse = httpRes

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.PostAPIServicesZisRegistryIntegrationAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			out.RawResponse = httpRes

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.PostAPIServicesZisRegistryIntegrationAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			out.RawResponse = httpRes

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
