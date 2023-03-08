/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ApplicationUsageApiService ApplicationUsageApi service
type ApplicationUsageApiService service

type ApplicationUsageApiGetNetworkClientsApplicationUsageRequest struct {
	ctx context.Context
	ApiService *ApplicationUsageApiService
	networkId string
	clients *string
	ssidNumber *int32
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
}

// A list of client keys, MACs or IPs separated by comma.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) Clients(clients string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.clients = &clients
	return r
}

// An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) SsidNumber(ssidNumber int32) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.ssidNumber = &ssidNumber
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) PerPage(perPage int32) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) StartingAfter(startingAfter string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) EndingBefore(endingBefore string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) T0(t0 string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) T1(t1 string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) Timespan(timespan float32) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkClientsApplicationUsageExecute(r)
}

/*
GetNetworkClientsApplicationUsage Return the application usage data for clients

Return the application usage data for clients. Usage data is in kilobytes. Clients can be identified by client keys or either the MACs or IPs depending on whether the network uses Track-by-IP.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return ApplicationUsageApiGetNetworkClientsApplicationUsageRequest
*/
func (a *ApplicationUsageApiService) GetNetworkClientsApplicationUsage(ctx context.Context, networkId string) ApplicationUsageApiGetNetworkClientsApplicationUsageRequest {
	return ApplicationUsageApiGetNetworkClientsApplicationUsageRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ApplicationUsageApiService) GetNetworkClientsApplicationUsageExecute(r ApplicationUsageApiGetNetworkClientsApplicationUsageRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationUsageApiService.GetNetworkClientsApplicationUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/clients/applicationUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clients == nil {
		return localVarReturnValue, nil, reportError("clients is required and must be specified")
	}

	localVarQueryParams.Add("clients", parameterToString(*r.clients, ""))
	if r.ssidNumber != nil {
		localVarQueryParams.Add("ssidNumber", parameterToString(*r.ssidNumber, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
