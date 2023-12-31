// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseResponseBody - Internal Server Error
type PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseResponseBody{}

func (e *PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseBody - Unauthorized
type PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseBody{}

func (e *PostAPIServicesZisRegistryIntegrationBundlesAPIGreaterThanServicesGreaterThanZisGreaterThanRegistryGreaterThanIntegrationGreaterThanBundlesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// PostAPIServicesZisRegistryIntegrationBundlesResponseBody - Bad Request
type PostAPIServicesZisRegistryIntegrationBundlesResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PostAPIServicesZisRegistryIntegrationBundlesResponseBody{}

func (e *PostAPIServicesZisRegistryIntegrationBundlesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
