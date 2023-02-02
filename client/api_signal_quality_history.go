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


// SignalQualityHistoryApiService SignalQualityHistoryApi service
type SignalQualityHistoryApiService service

type SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest struct {
	ctx context.Context
	ApiService *SignalQualityHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	autoResolution *bool
	clientId *string
	deviceSerial *string
	apTag *string
	band *string
	ssid *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) T0(t0 string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) T1(t1 string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) Timespan(timespan float32) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) Resolution(resolution int32) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.resolution = &resolution
	return r
}

// Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) AutoResolution(autoResolution bool) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.autoResolution = &autoResolution
	return r
}

// Filter results by network client.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) ClientId(clientId string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.clientId = &clientId
	return r
}

// Filter results by device.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) DeviceSerial(deviceSerial string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.deviceSerial = &deviceSerial
	return r
}

// Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) ApTag(apTag string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.apTag = &apTag
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;).
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) Band(band string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.band = &band
	return r
}

// Filter results by SSID number.
func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) Ssid(ssid int32) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	r.ssid = &ssid
	return r
}

func (r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) Execute() ([]InlineResponse20076, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessSignalQualityHistoryExecute(r)
}

/*
GetNetworkWirelessSignalQualityHistory Return signal quality (SNR/RSSI) over time for a device or network client

Return signal quality (SNR/RSSI) over time for a device or network client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest
*/
func (a *SignalQualityHistoryApiService) GetNetworkWirelessSignalQualityHistory(ctx context.Context, networkId string) SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest {
	return SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20076
func (a *SignalQualityHistoryApiService) GetNetworkWirelessSignalQualityHistoryExecute(r SignalQualityHistoryApiGetNetworkWirelessSignalQualityHistoryRequest) ([]InlineResponse20076, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20076
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SignalQualityHistoryApiService.GetNetworkWirelessSignalQualityHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/signalQualityHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	if r.autoResolution != nil {
		localVarQueryParams.Add("autoResolution", parameterToString(*r.autoResolution, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("clientId", parameterToString(*r.clientId, ""))
	}
	if r.deviceSerial != nil {
		localVarQueryParams.Add("deviceSerial", parameterToString(*r.deviceSerial, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
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
