/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkSwitchStormControl200Response struct for GetNetworkSwitchStormControl200Response
type GetNetworkSwitchStormControl200Response struct {
	// Broadcast threshold.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Multicast threshold.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Unknown Unicast threshold.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
}

// NewGetNetworkSwitchStormControl200Response instantiates a new GetNetworkSwitchStormControl200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchStormControl200Response() *GetNetworkSwitchStormControl200Response {
	this := GetNetworkSwitchStormControl200Response{}
	return &this
}

// NewGetNetworkSwitchStormControl200ResponseWithDefaults instantiates a new GetNetworkSwitchStormControl200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchStormControl200ResponseWithDefaults() *GetNetworkSwitchStormControl200Response {
	this := GetNetworkSwitchStormControl200Response{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *GetNetworkSwitchStormControl200Response) GetBroadcastThreshold() int32 {
	if o == nil || isNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStormControl200Response) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.BroadcastThreshold) {
    return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *GetNetworkSwitchStormControl200Response) HasBroadcastThreshold() bool {
	if o != nil && !isNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *GetNetworkSwitchStormControl200Response) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *GetNetworkSwitchStormControl200Response) GetMulticastThreshold() int32 {
	if o == nil || isNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStormControl200Response) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.MulticastThreshold) {
    return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *GetNetworkSwitchStormControl200Response) HasMulticastThreshold() bool {
	if o != nil && !isNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *GetNetworkSwitchStormControl200Response) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *GetNetworkSwitchStormControl200Response) GetUnknownUnicastThreshold() int32 {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStormControl200Response) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
    return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *GetNetworkSwitchStormControl200Response) HasUnknownUnicastThreshold() bool {
	if o != nil && !isNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *GetNetworkSwitchStormControl200Response) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

func (o GetNetworkSwitchStormControl200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BroadcastThreshold) {
		toSerialize["broadcastThreshold"] = o.BroadcastThreshold
	}
	if !isNil(o.MulticastThreshold) {
		toSerialize["multicastThreshold"] = o.MulticastThreshold
	}
	if !isNil(o.UnknownUnicastThreshold) {
		toSerialize["unknownUnicastThreshold"] = o.UnknownUnicastThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchStormControl200Response struct {
	value *GetNetworkSwitchStormControl200Response
	isSet bool
}

func (v NullableGetNetworkSwitchStormControl200Response) Get() *GetNetworkSwitchStormControl200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchStormControl200Response) Set(val *GetNetworkSwitchStormControl200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchStormControl200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchStormControl200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchStormControl200Response(val *GetNetworkSwitchStormControl200Response) *NullableGetNetworkSwitchStormControl200Response {
	return &NullableGetNetworkSwitchStormControl200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchStormControl200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchStormControl200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


