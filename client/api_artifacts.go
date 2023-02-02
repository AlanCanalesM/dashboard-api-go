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


// ArtifactsApiService ArtifactsApi service
type ArtifactsApiService service

type ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	organizationId string
	createOrganizationCameraCustomAnalyticsArtifact *InlineObject187
}

func (r ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact InlineObject187) ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest {
	r.createOrganizationCameraCustomAnalyticsArtifact = &createOrganizationCameraCustomAnalyticsArtifact
	return r
}

func (r ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
CreateOrganizationCameraCustomAnalyticsArtifact Create custom analytics artifact

Create custom analytics artifact. Returns an artifact upload URL with expiry time. Upload the artifact file with a put request to the returned upload URL before its expiry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *ArtifactsApiService) CreateOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string) ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest {
	return ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ArtifactsApiService) CreateOrganizationCameraCustomAnalyticsArtifactExecute(r ArtifactsApiCreateOrganizationCameraCustomAnalyticsArtifactRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.CreateOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.createOrganizationCameraCustomAnalyticsArtifact
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

type ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	organizationId string
	artifactId string
}

func (r ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
DeleteOrganizationCameraCustomAnalyticsArtifact Delete Custom Analytics Artifact

Delete Custom Analytics Artifact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param artifactId
 @return ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *ArtifactsApiService) DeleteOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string, artifactId string) ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest {
	return ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		artifactId: artifactId,
	}
}

// Execute executes the request
func (a *ArtifactsApiService) DeleteOrganizationCameraCustomAnalyticsArtifactExecute(r ArtifactsApiDeleteOrganizationCameraCustomAnalyticsArtifactRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.DeleteOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	organizationId string
	artifactId string
}

func (r ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraCustomAnalyticsArtifactExecute(r)
}

/*
GetOrganizationCameraCustomAnalyticsArtifact Get Custom Analytics Artifact

Get Custom Analytics Artifact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param artifactId
 @return ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest
*/
func (a *ArtifactsApiService) GetOrganizationCameraCustomAnalyticsArtifact(ctx context.Context, organizationId string, artifactId string) ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest {
	return ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ArtifactsApiService) GetOrganizationCameraCustomAnalyticsArtifactExecute(r ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.GetOrganizationCameraCustomAnalyticsArtifact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

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

type ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	organizationId string
}

func (r ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraCustomAnalyticsArtifactsExecute(r)
}

/*
GetOrganizationCameraCustomAnalyticsArtifacts List Custom Analytics Artifacts

List Custom Analytics Artifacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest
*/
func (a *ArtifactsApiService) GetOrganizationCameraCustomAnalyticsArtifacts(ctx context.Context, organizationId string) ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest {
	return ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ArtifactsApiService) GetOrganizationCameraCustomAnalyticsArtifactsExecute(r ArtifactsApiGetOrganizationCameraCustomAnalyticsArtifactsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.GetOrganizationCameraCustomAnalyticsArtifacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/customAnalytics/artifacts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
