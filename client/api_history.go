/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
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


// HistoryApiService HistoryApi service
type HistoryApiService service

type HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	serial string
	zoneId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	objectType *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) T0(t0 string) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.
func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) T1(t1 string) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.
func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) Timespan(timespan float32) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.
func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) Resolution(resolution int32) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.resolution = &resolution
	return r
}

// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) ObjectType(objectType string) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.objectType = &objectType
	return r
}

func (r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsZoneHistoryExecute(r)
}

/*
GetDeviceCameraAnalyticsZoneHistory Return historical records for analytic zones

Return historical records for analytic zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @param zoneId
 @return HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest
*/
func (a *HistoryApiService) GetDeviceCameraAnalyticsZoneHistory(ctx context.Context, serial string, zoneId string) HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	return HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		zoneId: zoneId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *HistoryApiService) GetDeviceCameraAnalyticsZoneHistoryExecute(r HistoryApiGetDeviceCameraAnalyticsZoneHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetDeviceCameraAnalyticsZoneHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/zones/{zoneId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", url.PathEscape(parameterToString(r.zoneId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
	}
	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
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

type HistoryApiGetNetworkAlertsHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 30.
func (r HistoryApiGetNetworkAlertsHistoryRequest) PerPage(perPage int32) HistoryApiGetNetworkAlertsHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r HistoryApiGetNetworkAlertsHistoryRequest) StartingAfter(startingAfter string) HistoryApiGetNetworkAlertsHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r HistoryApiGetNetworkAlertsHistoryRequest) EndingBefore(endingBefore string) HistoryApiGetNetworkAlertsHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r HistoryApiGetNetworkAlertsHistoryRequest) Execute() ([]GetNetworkAlertsHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAlertsHistoryExecute(r)
}

/*
GetNetworkAlertsHistory Return the alert history for this network

Return the alert history for this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return HistoryApiGetNetworkAlertsHistoryRequest
*/
func (a *HistoryApiService) GetNetworkAlertsHistory(ctx context.Context, networkId string) HistoryApiGetNetworkAlertsHistoryRequest {
	return HistoryApiGetNetworkAlertsHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkAlertsHistory200ResponseInner
func (a *HistoryApiService) GetNetworkAlertsHistoryExecute(r HistoryApiGetNetworkAlertsHistoryRequest) ([]GetNetworkAlertsHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkAlertsHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetNetworkAlertsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/alerts/history"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type HistoryApiGetOrganizationSensorReadingsHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	networkIds *[]string
	serials *[]string
	metrics *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) PerPage(perPage int32) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) StartingAfter(startingAfter string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) EndingBefore(endingBefore string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) T0(t0 string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) T1(t1 string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) Timespan(timespan float32) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.timespan = &timespan
	return r
}

// Optional parameter to filter readings by network.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) NetworkIds(networkIds []string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter readings by sensor.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) Serials(serials []string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.serials = &serials
	return r
}

// Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) Metrics(metrics []string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	r.metrics = &metrics
	return r
}

func (r HistoryApiGetOrganizationSensorReadingsHistoryRequest) Execute() ([]GetOrganizationSensorReadingsHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSensorReadingsHistoryExecute(r)
}

/*
GetOrganizationSensorReadingsHistory Return all reported readings from sensors in a given timespan, sorted by timestamp

Return all reported readings from sensors in a given timespan, sorted by timestamp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return HistoryApiGetOrganizationSensorReadingsHistoryRequest
*/
func (a *HistoryApiService) GetOrganizationSensorReadingsHistory(ctx context.Context, organizationId string) HistoryApiGetOrganizationSensorReadingsHistoryRequest {
	return HistoryApiGetOrganizationSensorReadingsHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSensorReadingsHistory200ResponseInner
func (a *HistoryApiService) GetOrganizationSensorReadingsHistoryExecute(r HistoryApiGetOrganizationSensorReadingsHistoryRequest) ([]GetOrganizationSensorReadingsHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSensorReadingsHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetOrganizationSensorReadingsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sensor/readings/history"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.metrics != nil {
		localVarQueryParams.Add("metrics", parameterToString(*r.metrics, "csv"))
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
