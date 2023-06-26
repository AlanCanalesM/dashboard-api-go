/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkWirelessSsidEapOverride200ResponseEapolKey EAPOL Key settings.
type GetNetworkWirelessSsidEapOverride200ResponseEapolKey struct {
	// Maximum number of EAPOL key retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAPOL Key timeout in milliseconds.
	TimeoutInMs *int32 `json:"timeoutInMs,omitempty"`
}

// NewGetNetworkWirelessSsidEapOverride200ResponseEapolKey instantiates a new GetNetworkWirelessSsidEapOverride200ResponseEapolKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidEapOverride200ResponseEapolKey() *GetNetworkWirelessSsidEapOverride200ResponseEapolKey {
	this := GetNetworkWirelessSsidEapOverride200ResponseEapolKey{}
	return &this
}

// NewGetNetworkWirelessSsidEapOverride200ResponseEapolKeyWithDefaults instantiates a new GetNetworkWirelessSsidEapOverride200ResponseEapolKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidEapOverride200ResponseEapolKeyWithDefaults() *GetNetworkWirelessSsidEapOverride200ResponseEapolKey {
	this := GetNetworkWirelessSsidEapOverride200ResponseEapolKey{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeoutInMs returns the TimeoutInMs field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) GetTimeoutInMs() int32 {
	if o == nil || isNil(o.TimeoutInMs) {
		var ret int32
		return ret
	}
	return *o.TimeoutInMs
}

// GetTimeoutInMsOk returns a tuple with the TimeoutInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) GetTimeoutInMsOk() (*int32, bool) {
	if o == nil || isNil(o.TimeoutInMs) {
    return nil, false
	}
	return o.TimeoutInMs, true
}

// HasTimeoutInMs returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) HasTimeoutInMs() bool {
	if o != nil && !isNil(o.TimeoutInMs) {
		return true
	}

	return false
}

// SetTimeoutInMs gets a reference to the given int32 and assigns it to the TimeoutInMs field.
func (o *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) SetTimeoutInMs(v int32) {
	o.TimeoutInMs = &v
}

func (o GetNetworkWirelessSsidEapOverride200ResponseEapolKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.TimeoutInMs) {
		toSerialize["timeoutInMs"] = o.TimeoutInMs
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey struct {
	value *GetNetworkWirelessSsidEapOverride200ResponseEapolKey
	isSet bool
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) Get() *GetNetworkWirelessSsidEapOverride200ResponseEapolKey {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) Set(val *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey(val *GetNetworkWirelessSsidEapOverride200ResponseEapolKey) *NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey {
	return &NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidEapOverride200ResponseEapolKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


