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


// ByDeviceApiService ByDeviceApi service
type ByDeviceApiService service

type ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct {
	ctx context.Context
	ApiService *ByDeviceApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) PerPage(perPage int32) ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) StartingAfter(startingAfter string) ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) EndingBefore(endingBefore string) ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) Execute() ([]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings

Return the devices that have a Dynamic ARP Inspection warning and their warnings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest
*/
func (a *ByDeviceApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx context.Context, networkId string) ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	return ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
func (a *ByDeviceApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r ByDeviceApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) ([]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByDeviceApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
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

type ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest struct {
	ctx context.Context
	ApiService *ByDeviceApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	serials *[]string
	tags *[]string
	tagsFilterType *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) PerPage(perPage int32) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) StartingAfter(startingAfter string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) EndingBefore(endingBefore string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) NetworkIds(networkIds []string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) ProductTypes(productTypes []string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Serials(serials []string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.serials = &serials
	return r
}

// An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Tags(tags []string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.tags = &tags
	return r
}

// An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected.
func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) TagsFilterType(tagsFilterType string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

func (r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) Execute() ([]GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesPowerModulesStatusesByDeviceExecute(r)
}

/*
GetOrganizationDevicesPowerModulesStatusesByDevice List the power status information for devices in an organization

List the power status information for devices in an organization. The data returned by this endpoint is updated every 5 minutes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest
*/
func (a *ByDeviceApiService) GetOrganizationDevicesPowerModulesStatusesByDevice(ctx context.Context, organizationId string) ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest {
	return ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner
func (a *ByDeviceApiService) GetOrganizationDevicesPowerModulesStatusesByDeviceExecute(r ByDeviceApiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest) ([]GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByDeviceApiService.GetOrganizationDevicesPowerModulesStatusesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/powerModules/statuses/byDevice"
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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	}
	if r.tagsFilterType != nil {
		localVarQueryParams.Add("tagsFilterType", parameterToString(*r.tagsFilterType, ""))
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

type ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest struct {
	ctx context.Context
	ApiService *ByDeviceApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	serials *[]string
	tags *[]string
	tagsFilterType *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) PerPage(perPage int32) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) StartingAfter(startingAfter string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) EndingBefore(endingBefore string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) NetworkIds(networkIds []string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) ProductTypes(productTypes []string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Serials(serials []string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.serials = &serials
	return r
}

// An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Tags(tags []string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.tags = &tags
	return r
}

// An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected.
func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) TagsFilterType(tagsFilterType string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

func (r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) Execute() ([]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesUplinksAddressesByDeviceExecute(r)
}

/*
GetOrganizationDevicesUplinksAddressesByDevice List the current uplink addresses for devices in an organization.

List the current uplink addresses for devices in an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest
*/
func (a *ByDeviceApiService) GetOrganizationDevicesUplinksAddressesByDevice(ctx context.Context, organizationId string) ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest {
	return ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner
func (a *ByDeviceApiService) GetOrganizationDevicesUplinksAddressesByDeviceExecute(r ByDeviceApiGetOrganizationDevicesUplinksAddressesByDeviceRequest) ([]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByDeviceApiService.GetOrganizationDevicesUplinksAddressesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/uplinks/addresses/byDevice"
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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	}
	if r.tagsFilterType != nil {
		localVarQueryParams.Add("tagsFilterType", parameterToString(*r.tagsFilterType, ""))
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

type ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest struct {
	ctx context.Context
	ApiService *ByDeviceApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	macs *[]string
	firmwareUpgradeIds *[]string
	firmwareUpgradeBatchIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) PerPage(perPage int32) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) StartingAfter(startingAfter string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) EndingBefore(endingBefore string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter by network
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) NetworkIds(networkIds []string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) Serials(serials []string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) Macs(macs []string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter by firmware upgrade ids.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeIds(firmwareUpgradeIds []string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeIds = &firmwareUpgradeIds
	return r
}

// Optional parameter to filter by firmware upgrade batch ids.
func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeBatchIds = &firmwareUpgradeBatchIds
	return r
}

func (r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) Execute() ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesByDeviceExecute(r)
}

/*
GetOrganizationFirmwareUpgradesByDevice Get firmware upgrade status for the filtered devices

Get firmware upgrade status for the filtered devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest
*/
func (a *ByDeviceApiService) GetOrganizationFirmwareUpgradesByDevice(ctx context.Context, organizationId string) ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	return ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
func (a *ByDeviceApiService) GetOrganizationFirmwareUpgradesByDeviceExecute(r ByDeviceApiGetOrganizationFirmwareUpgradesByDeviceRequest) ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByDeviceApiService.GetOrganizationFirmwareUpgradesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades/byDevice"
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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.firmwareUpgradeIds != nil {
		localVarQueryParams.Add("firmwareUpgradeIds", parameterToString(*r.firmwareUpgradeIds, "csv"))
	}
	if r.firmwareUpgradeBatchIds != nil {
		localVarQueryParams.Add("firmwareUpgradeBatchIds", parameterToString(*r.firmwareUpgradeBatchIds, "csv"))
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
