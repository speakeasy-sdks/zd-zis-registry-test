// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIServicesZisRegistryIntegrationSecurity struct {
	Password string `security:"scheme,type=http,subtype=basic,name=password"`
	Username string `security:"scheme,type=http,subtype=basic,name=username"`
}

func (o *PostAPIServicesZisRegistryIntegrationSecurity) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *PostAPIServicesZisRegistryIntegrationSecurity) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type PostAPIServicesZisRegistryIntegrationRequestBody struct {
}

type PostAPIServicesZisRegistryIntegrationRequest struct {
	RequestBody *PostAPIServicesZisRegistryIntegrationRequestBody `request:"mediaType=application/json"`
	Integration string                                            `pathParam:"style=simple,explode=false,name=integration"`
}

func (o *PostAPIServicesZisRegistryIntegrationRequest) GetRequestBody() *PostAPIServicesZisRegistryIntegrationRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *PostAPIServicesZisRegistryIntegrationRequest) GetIntegration() string {
	if o == nil {
		return ""
	}
	return o.Integration
}

// PostAPIServicesZisRegistryIntegration500ApplicationJSON - Internal Server Error
type PostAPIServicesZisRegistryIntegration500ApplicationJSON struct {
}

// PostAPIServicesZisRegistryIntegration401ApplicationJSON - Unauthorized
type PostAPIServicesZisRegistryIntegration401ApplicationJSON struct {
}

// PostAPIServicesZisRegistryIntegration400ApplicationJSON - Bad Request
type PostAPIServicesZisRegistryIntegration400ApplicationJSON struct {
}

// PostAPIServicesZisRegistryIntegration200ApplicationJSON - OK
type PostAPIServicesZisRegistryIntegration200ApplicationJSON struct {
}

type PostAPIServicesZisRegistryIntegrationResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PostAPIServicesZisRegistryIntegration200ApplicationJSONObject *PostAPIServicesZisRegistryIntegration200ApplicationJSON
	// Bad Request
	PostAPIServicesZisRegistryIntegration400ApplicationJSONObject *PostAPIServicesZisRegistryIntegration400ApplicationJSON
	// Unauthorized
	PostAPIServicesZisRegistryIntegration401ApplicationJSONObject *PostAPIServicesZisRegistryIntegration401ApplicationJSON
	// Internal Server Error
	PostAPIServicesZisRegistryIntegration500ApplicationJSONObject *PostAPIServicesZisRegistryIntegration500ApplicationJSON
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetPostAPIServicesZisRegistryIntegration200ApplicationJSONObject() *PostAPIServicesZisRegistryIntegration200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PostAPIServicesZisRegistryIntegration200ApplicationJSONObject
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetPostAPIServicesZisRegistryIntegration400ApplicationJSONObject() *PostAPIServicesZisRegistryIntegration400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PostAPIServicesZisRegistryIntegration400ApplicationJSONObject
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetPostAPIServicesZisRegistryIntegration401ApplicationJSONObject() *PostAPIServicesZisRegistryIntegration401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PostAPIServicesZisRegistryIntegration401ApplicationJSONObject
}

func (o *PostAPIServicesZisRegistryIntegrationResponse) GetPostAPIServicesZisRegistryIntegration500ApplicationJSONObject() *PostAPIServicesZisRegistryIntegration500ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PostAPIServicesZisRegistryIntegration500ApplicationJSONObject
}
