/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
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


// VpnFirewallRulesApiService VpnFirewallRulesApi service
type VpnFirewallRulesApiService service

type VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct {
	ctx context.Context
	ApiService *VpnFirewallRulesApiService
	organizationId string
}

func (r VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest) Execute() (*GetOrganizationApplianceVpnVpnFirewallRules200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceVpnVpnFirewallRulesExecute(r)
}

/*
GetOrganizationApplianceVpnVpnFirewallRules Return the firewall rules for an organization's site-to-site VPN

Return the firewall rules for an organization's site-to-site VPN

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest
*/
func (a *VpnFirewallRulesApiService) GetOrganizationApplianceVpnVpnFirewallRules(ctx context.Context, organizationId string) VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest {
	return VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationApplianceVpnVpnFirewallRules200Response
func (a *VpnFirewallRulesApiService) GetOrganizationApplianceVpnVpnFirewallRulesExecute(r VpnFirewallRulesApiGetOrganizationApplianceVpnVpnFirewallRulesRequest) (*GetOrganizationApplianceVpnVpnFirewallRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationApplianceVpnVpnFirewallRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VpnFirewallRulesApiService.GetOrganizationApplianceVpnVpnFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct {
	ctx context.Context
	ApiService *VpnFirewallRulesApiService
	organizationId string
	updateOrganizationApplianceVpnVpnFirewallRulesRequest *UpdateOrganizationApplianceVpnVpnFirewallRulesRequest
}

func (r VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest UpdateOrganizationApplianceVpnVpnFirewallRulesRequest) VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	r.updateOrganizationApplianceVpnVpnFirewallRulesRequest = &updateOrganizationApplianceVpnVpnFirewallRulesRequest
	return r
}

func (r VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) Execute() (*GetOrganizationApplianceVpnVpnFirewallRules200Response, *http.Response, error) {
	return r.ApiService.UpdateOrganizationApplianceVpnVpnFirewallRulesExecute(r)
}

/*
UpdateOrganizationApplianceVpnVpnFirewallRules Update the firewall rules of an organization's site-to-site VPN

Update the firewall rules of an organization's site-to-site VPN

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest
*/
func (a *VpnFirewallRulesApiService) UpdateOrganizationApplianceVpnVpnFirewallRules(ctx context.Context, organizationId string) VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest {
	return VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationApplianceVpnVpnFirewallRules200Response
func (a *VpnFirewallRulesApiService) UpdateOrganizationApplianceVpnVpnFirewallRulesExecute(r VpnFirewallRulesApiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest) (*GetOrganizationApplianceVpnVpnFirewallRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationApplianceVpnVpnFirewallRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VpnFirewallRulesApiService.UpdateOrganizationApplianceVpnVpnFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateOrganizationApplianceVpnVpnFirewallRulesRequest
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
