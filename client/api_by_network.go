/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
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
	"reflect"
)


// ByNetworkApiService ByNetworkApi service
type ByNetworkApiService service

type ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest struct {
	ctx context.Context
	ApiService *ByNetworkApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest) T0(t0 string) ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 14 days after t0.
func (r ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest) T1(t1 string) ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day.
func (r ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest) Timespan(timespan float32) ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest {
	r.timespan = &timespan
	return r
}

func (r ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest) Execute() ([]GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceUplinksUsageByNetworkExecute(r)
}

/*
GetOrganizationApplianceUplinksUsageByNetwork Get the sent and received bytes for each uplink of all MX and Z networks within an organization

Get the sent and received bytes for each uplink of all MX and Z networks within an organization. If more than one device was active during the specified timespan, then the sent and received bytes will be aggregated by interface.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest
*/
func (a *ByNetworkApiService) GetOrganizationApplianceUplinksUsageByNetwork(ctx context.Context, organizationId string) ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest {
	return ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner
func (a *ByNetworkApiService) GetOrganizationApplianceUplinksUsageByNetworkExecute(r ByNetworkApiGetOrganizationApplianceUplinksUsageByNetworkRequest) ([]GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByNetworkApiService.GetOrganizationApplianceUplinksUsageByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/uplinks/usage/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

type ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest struct {
	ctx context.Context
	ApiService *ByNetworkApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
}

// Filter results by network.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) NetworkIds(networkIds []string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) Serials(serials []string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.serials = &serials
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) PerPage(perPage int32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) StartingAfter(startingAfter string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) EndingBefore(endingBefore string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) T0(t0 string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) T1(t1 string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) Timespan(timespan float32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) Interval(interval int32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	r.interval = &interval
	return r
}

func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) Execute() ([]GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesChannelUtilizationByNetworkExecute(r)
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetwork Get average channel utilization across all bands for all networks in the organization

Get average channel utilization across all bands for all networks in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest
*/
func (a *ByNetworkApiService) GetOrganizationWirelessDevicesChannelUtilizationByNetwork(ctx context.Context, organizationId string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest {
	return ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner
func (a *ByNetworkApiService) GetOrganizationWirelessDevicesChannelUtilizationByNetworkExecute(r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest) ([]GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByNetworkApiService.GetOrganizationWirelessDevicesChannelUtilizationByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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

type ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct {
	ctx context.Context
	ApiService *ByNetworkApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
}

// Filter results by network.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) NetworkIds(networkIds []string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Serials(serials []string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.serials = &serials
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) PerPage(perPage int32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) StartingAfter(startingAfter string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) EndingBefore(endingBefore string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) T0(t0 string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) T1(t1 string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Timespan(timespan float32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Interval(interval int32) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.interval = &interval
	return r
}

func (r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Execute() ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalExecute(r)
}

/*
GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval Get a time-series of average channel utilization for all bands

Get a time-series of average channel utilization for all bands

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest
*/
func (a *ByNetworkApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx context.Context, organizationId string) ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	return ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
func (a *ByNetworkApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalExecute(r ByNetworkApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByNetworkApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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
