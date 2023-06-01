/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20092 struct for InlineResponse20092
type InlineResponse20092 struct {
	// Database ID for the admin user who made the API request.
	AdminId *string `json:"adminId,omitempty"`
	// HTTP method used in the API request.
	Method *string `json:"method,omitempty"`
	// The host which the API request was directed at.
	Host *string `json:"host,omitempty"`
	// The API request path.
	Path *string `json:"path,omitempty"`
	// The query string sent with the API request.
	QueryString *string `json:"queryString,omitempty"`
	// The API request user agent.
	UserAgent *string `json:"userAgent,omitempty"`
	// Timestamp, in iso8601 format, indicating when the API request was made.
	Ts *time.Time `json:"ts,omitempty"`
	// API request response code.
	ResponseCode *int32 `json:"responseCode,omitempty"`
	// Public IP address from which the API request was made.
	SourceIp *string `json:"sourceIp,omitempty"`
	// API version of the endpoint.
	Version *int32 `json:"version,omitempty"`
	// Operation ID for the endpoint.
	OperationId *string `json:"operationId,omitempty"`
}

// NewInlineResponse20092 instantiates a new InlineResponse20092 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092WithDefaults() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *InlineResponse20092) GetAdminId() string {
	if o == nil || isNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetAdminIdOk() (*string, bool) {
	if o == nil || isNil(o.AdminId) {
    return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *InlineResponse20092) HasAdminId() bool {
	if o != nil && !isNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *InlineResponse20092) SetAdminId(v string) {
	o.AdminId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineResponse20092) GetMethod() string {
	if o == nil || isNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetMethodOk() (*string, bool) {
	if o == nil || isNil(o.Method) {
    return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineResponse20092) HasMethod() bool {
	if o != nil && !isNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineResponse20092) SetMethod(v string) {
	o.Method = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineResponse20092) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineResponse20092) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InlineResponse20092) SetHost(v string) {
	o.Host = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *InlineResponse20092) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *InlineResponse20092) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *InlineResponse20092) SetPath(v string) {
	o.Path = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *InlineResponse20092) GetQueryString() string {
	if o == nil || isNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetQueryStringOk() (*string, bool) {
	if o == nil || isNil(o.QueryString) {
    return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *InlineResponse20092) HasQueryString() bool {
	if o != nil && !isNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *InlineResponse20092) SetQueryString(v string) {
	o.QueryString = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *InlineResponse20092) GetUserAgent() string {
	if o == nil || isNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetUserAgentOk() (*string, bool) {
	if o == nil || isNil(o.UserAgent) {
    return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *InlineResponse20092) HasUserAgent() bool {
	if o != nil && !isNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *InlineResponse20092) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20092) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20092) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20092) SetTs(v time.Time) {
	o.Ts = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *InlineResponse20092) GetResponseCode() int32 {
	if o == nil || isNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetResponseCodeOk() (*int32, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *InlineResponse20092) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *InlineResponse20092) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise.
func (o *InlineResponse20092) GetSourceIp() string {
	if o == nil || isNil(o.SourceIp) {
		var ret string
		return ret
	}
	return *o.SourceIp
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetSourceIpOk() (*string, bool) {
	if o == nil || isNil(o.SourceIp) {
    return nil, false
	}
	return o.SourceIp, true
}

// HasSourceIp returns a boolean if a field has been set.
func (o *InlineResponse20092) HasSourceIp() bool {
	if o != nil && !isNil(o.SourceIp) {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given string and assigns it to the SourceIp field.
func (o *InlineResponse20092) SetSourceIp(v string) {
	o.SourceIp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20092) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20092) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *InlineResponse20092) SetVersion(v int32) {
	o.Version = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *InlineResponse20092) GetOperationId() string {
	if o == nil || isNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetOperationIdOk() (*string, bool) {
	if o == nil || isNil(o.OperationId) {
    return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *InlineResponse20092) HasOperationId() bool {
	if o != nil && !isNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *InlineResponse20092) SetOperationId(v string) {
	o.OperationId = &v
}

func (o InlineResponse20092) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	if !isNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.QueryString) {
		toSerialize["queryString"] = o.QueryString
	}
	if !isNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.SourceIp) {
		toSerialize["sourceIp"] = o.SourceIp
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.OperationId) {
		toSerialize["operationId"] = o.OperationId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092 struct {
	value *InlineResponse20092
	isSet bool
}

func (v NullableInlineResponse20092) Get() *InlineResponse20092 {
	return v.value
}

func (v *NullableInlineResponse20092) Set(val *InlineResponse20092) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092(val *InlineResponse20092) *NullableInlineResponse20092 {
	return &NullableInlineResponse20092{value: val, isSet: true}
}

func (v NullableInlineResponse20092) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


