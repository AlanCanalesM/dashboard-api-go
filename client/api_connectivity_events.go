/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
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


// ConnectivityEventsApiService ConnectivityEventsApi service
type ConnectivityEventsApiService service

type ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest struct {
	ctx context.Context
	ApiService *ConnectivityEventsApiService
	networkId string
	clientId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	types *[]string
	includedSeverities *[]string
	band *string
	ssidNumber *int32
	deviceSerial *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) PerPage(perPage int32) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) StartingAfter(startingAfter string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) EndingBefore(endingBefore string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) T0(t0 string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) T1(t1 string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) Timespan(timespan float32) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.timespan = &timespan
	return r
}

// A list of event types to include. If not specified, events of all types will be returned. Valid types are &#39;assoc&#39;, &#39;disassoc&#39;, &#39;auth&#39;, &#39;deauth&#39;, &#39;dns&#39;, &#39;dhcp&#39;, &#39;roam&#39;, &#39;connection&#39; and/or &#39;sticky&#39;.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) Types(types []string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.types = &types
	return r
}

// A list of severities to include. If not specified, events of all severities will be returned. Valid severities are &#39;good&#39;, &#39;info&#39;, &#39;warn&#39; and/or &#39;bad&#39;.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) IncludedSeverities(includedSeverities []string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.includedSeverities = &includedSeverities
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39;, &#39;6&#39;).
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) Band(band string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.band = &band
	return r
}

// An SSID number to include. If not specified, events for all SSIDs will be returned.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) SsidNumber(ssidNumber int32) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.ssidNumber = &ssidNumber
	return r
}

// Filter results by an AP&#39;s serial number.
func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) DeviceSerial(deviceSerial string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	r.deviceSerial = &deviceSerial
	return r
}

func (r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientConnectivityEventsExecute(r)
}

/*
GetNetworkWirelessClientConnectivityEvents List the wireless connectivity events for a client within a network in the timespan.

List the wireless connectivity events for a client within a network in the timespan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest
*/
func (a *ConnectivityEventsApiService) GetNetworkWirelessClientConnectivityEvents(ctx context.Context, networkId string, clientId string) ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest {
	return ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConnectivityEventsApiService) GetNetworkWirelessClientConnectivityEventsExecute(r ConnectivityEventsApiGetNetworkWirelessClientConnectivityEventsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityEventsApiService.GetNetworkWirelessClientConnectivityEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/connectivityEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.types != nil {
		localVarQueryParams.Add("types", parameterToString(*r.types, "csv"))
	}
	if r.includedSeverities != nil {
		localVarQueryParams.Add("includedSeverities", parameterToString(*r.includedSeverities, "csv"))
	}
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssidNumber != nil {
		localVarQueryParams.Add("ssidNumber", parameterToString(*r.ssidNumber, ""))
	}
	if r.deviceSerial != nil {
		localVarQueryParams.Add("deviceSerial", parameterToString(*r.deviceSerial, ""))
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
