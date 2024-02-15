/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ConnectionStatsApiService ConnectionStatsApi service
type ConnectionStatsApiService service

type ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest struct {
	ctx context.Context
	ApiService *ConnectionStatsApiService
	serial string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) T0(t0 string) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) T1(t1 string) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) Timespan(timespan float32) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) Band(band string) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) Ssid(ssid int32) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) Vlan(vlan int32) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) ApTag(apTag string) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	r.apTag = &apTag
	return r
}

func (r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) Execute() (*GetDeviceWirelessConnectionStats200Response, *http.Response, error) {
	return r.ApiService.GetDeviceWirelessConnectionStatsExecute(r)
}

/*
GetDeviceWirelessConnectionStats Aggregated connectivity info for a given AP on this network

Aggregated connectivity info for a given AP on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest
*/
func (a *ConnectionStatsApiService) GetDeviceWirelessConnectionStats(ctx context.Context, serial string) ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest {
	return ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return GetDeviceWirelessConnectionStats200Response
func (a *ConnectionStatsApiService) GetDeviceWirelessConnectionStatsExecute(r ConnectionStatsApiGetDeviceWirelessConnectionStatsRequest) (*GetDeviceWirelessConnectionStats200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDeviceWirelessConnectionStats200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionStatsApiService.GetDeviceWirelessConnectionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/wireless/connectionStats"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest struct {
	ctx context.Context
	ApiService *ConnectionStatsApiService
	networkId string
	clientId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) T0(t0 string) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) T1(t1 string) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) Timespan(timespan float32) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) Band(band string) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) Ssid(ssid int32) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) Vlan(vlan int32) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) ApTag(apTag string) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	r.apTag = &apTag
	return r
}

func (r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientConnectionStatsExecute(r)
}

/*
GetNetworkWirelessClientConnectionStats Aggregated connectivity info for a given client on this network

Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest
*/
func (a *ConnectionStatsApiService) GetNetworkWirelessClientConnectionStats(ctx context.Context, networkId string, clientId string) ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest {
	return ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConnectionStatsApiService) GetNetworkWirelessClientConnectionStatsExecute(r ConnectionStatsApiGetNetworkWirelessClientConnectionStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionStatsApiService.GetNetworkWirelessClientConnectionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/connectionStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest struct {
	ctx context.Context
	ApiService *ConnectionStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) T0(t0 string) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) T1(t1 string) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) Timespan(timespan float32) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) Band(band string) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) Ssid(ssid int32) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) Vlan(vlan int32) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) ApTag(apTag string) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	r.apTag = &apTag
	return r
}

func (r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientsConnectionStatsExecute(r)
}

/*
GetNetworkWirelessClientsConnectionStats Aggregated connectivity info for this network, grouped by clients

Aggregated connectivity info for this network, grouped by clients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest
*/
func (a *ConnectionStatsApiService) GetNetworkWirelessClientsConnectionStats(ctx context.Context, networkId string) ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest {
	return ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConnectionStatsApiService) GetNetworkWirelessClientsConnectionStatsExecute(r ConnectionStatsApiGetNetworkWirelessClientsConnectionStatsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionStatsApiService.GetNetworkWirelessClientsConnectionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/connectionStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest struct {
	ctx context.Context
	ApiService *ConnectionStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) T0(t0 string) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) T1(t1 string) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) Timespan(timespan float32) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) Band(band string) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) Ssid(ssid int32) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) Vlan(vlan int32) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) ApTag(apTag string) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	r.apTag = &apTag
	return r
}

func (r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) Execute() (*GetNetworkWirelessConnectionStats200Response, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessConnectionStatsExecute(r)
}

/*
GetNetworkWirelessConnectionStats Aggregated connectivity info for this network

Aggregated connectivity info for this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest
*/
func (a *ConnectionStatsApiService) GetNetworkWirelessConnectionStats(ctx context.Context, networkId string) ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest {
	return ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkWirelessConnectionStats200Response
func (a *ConnectionStatsApiService) GetNetworkWirelessConnectionStatsExecute(r ConnectionStatsApiGetNetworkWirelessConnectionStatsRequest) (*GetNetworkWirelessConnectionStats200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkWirelessConnectionStats200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionStatsApiService.GetNetworkWirelessConnectionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/connectionStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest struct {
	ctx context.Context
	ApiService *ConnectionStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) T0(t0 string) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) T1(t1 string) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) Timespan(timespan float32) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) Band(band string) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) Ssid(ssid int32) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) Vlan(vlan int32) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) ApTag(apTag string) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	r.apTag = &apTag
	return r
}

func (r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) Execute() ([]GetDeviceWirelessConnectionStats200Response, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessDevicesConnectionStatsExecute(r)
}

/*
GetNetworkWirelessDevicesConnectionStats Aggregated connectivity info for this network, grouped by node

Aggregated connectivity info for this network, grouped by node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest
*/
func (a *ConnectionStatsApiService) GetNetworkWirelessDevicesConnectionStats(ctx context.Context, networkId string) ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest {
	return ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetDeviceWirelessConnectionStats200Response
func (a *ConnectionStatsApiService) GetNetworkWirelessDevicesConnectionStatsExecute(r ConnectionStatsApiGetNetworkWirelessDevicesConnectionStatsRequest) ([]GetDeviceWirelessConnectionStats200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetDeviceWirelessConnectionStats200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionStatsApiService.GetNetworkWirelessDevicesConnectionStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/devices/connectionStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
