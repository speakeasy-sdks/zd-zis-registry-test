// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package zdzisregistrytest

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/zd-zis-registry-test/v2/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall struct {
	sdkConfiguration sdkConfiguration
}

func newAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall(sdkConfig sdkConfiguration) *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall {
	return &APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall{
		sdkConfiguration: sdkConfig,
	}
}

// DeleteAPIServicesZisRegistryJobSpecsInstall - Uninstall Job Spec
// Uninstalls a job spec for a ZIS integration.
//
// #### Authentication
//
// You can authorize requests using basic authentication, an API token, or a [ZIS OAuth access
// token](/documentation/integration-services/developer-guide/developing-private-zis-integrations/#obtaining-a-zis-oauth-token).
// A Zendesk app can also authorize requests to this endpoint using an
// admin's browser session. See [Making API requests from a Zendesk
// app](/documentation/apps/getting-started/making-api-requests-from-a-zendesk-app/#method-1-api-request-to-zendesk-api).
//
// #### Allowed for
//
// * Admins
func (s *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall) DeleteAPIServicesZisRegistryJobSpecsInstall(ctx context.Context, request operations.DeleteAPIServicesZisRegistryJobSpecsInstallRequest, opts ...operations.Option) (*operations.DeleteAPIServicesZisRegistryJobSpecsInstallResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/services/zis/registry/job_specs/install"

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, text/plain;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIServicesZisRegistryJobSpecsInstallResponse{
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
	case httpRes.StatusCode == 204:
		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			out := string(rawBody)
			res.Res = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallResponseBody
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
			var out sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			out.RawResponse = httpRes

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody
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
			var out sdkerrors.DeleteAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponse500ResponseBody
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

// PostAPIServicesZisRegistryJobSpecsInstall - Install Job Spec
// Installs a job spec for a ZIS integration.
//
// #### Authentication
//
// You can authorize requests using basic authentication, an API token, or a [ZIS OAuth access
// token](/documentation/integration-services/developer-guide/developing-private-zis-integrations/#obtaining-a-zis-oauth-token).
// A Zendesk app can also authorize requests to this endpoint using an
// admin's browser session. See [Making API requests from a Zendesk
// app](/documentation/apps/getting-started/making-api-requests-from-a-zendesk-app/#method-1-api-request-to-zendesk-api).
//
// #### Allowed for
//
// * Admins
func (s *APIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstall) PostAPIServicesZisRegistryJobSpecsInstall(ctx context.Context, request operations.PostAPIServicesZisRegistryJobSpecsInstallRequest, opts ...operations.Option) (*operations.PostAPIServicesZisRegistryJobSpecsInstallResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/services/zis/registry/job_specs/install"

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, text/plain;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIServicesZisRegistryJobSpecsInstallResponse{
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
		case utils.MatchContentType(contentType, `text/plain`):
			out := string(rawBody)
			res.Res = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallResponseBody
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
			var out sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseBody
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
			var out sdkerrors.PostAPIServicesZisRegistryJobSpecsInstallAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanJobSpecsGreaterThanInstallResponseResponseBody
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
