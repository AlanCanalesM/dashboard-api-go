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

// UpdateNetworkSwitchSettingsRequest struct for UpdateNetworkSwitchSettingsRequest
type UpdateNetworkSwitchSettingsRequest struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner `json:"powerExceptions,omitempty"`
}

// NewUpdateNetworkSwitchSettingsRequest instantiates a new UpdateNetworkSwitchSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchSettingsRequest() *UpdateNetworkSwitchSettingsRequest {
	this := UpdateNetworkSwitchSettingsRequest{}
	return &this
}

// NewUpdateNetworkSwitchSettingsRequestWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchSettingsRequestWithDefaults() *UpdateNetworkSwitchSettingsRequest {
	this := UpdateNetworkSwitchSettingsRequest{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchSettingsRequest) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequest) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchSettingsRequest) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateNetworkSwitchSettingsRequest) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchSettingsRequest) GetUseCombinedPower() bool {
	if o == nil || isNil(o.UseCombinedPower) {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequest) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || isNil(o.UseCombinedPower) {
    return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchSettingsRequest) HasUseCombinedPower() bool {
	if o != nil && !isNil(o.UseCombinedPower) {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *UpdateNetworkSwitchSettingsRequest) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchSettingsRequest) GetPowerExceptions() []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner {
	if o == nil || isNil(o.PowerExceptions) {
		var ret []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequest) GetPowerExceptionsOk() ([]UpdateNetworkSwitchSettingsRequestPowerExceptionsInner, bool) {
	if o == nil || isNil(o.PowerExceptions) {
    return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchSettingsRequest) HasPowerExceptions() bool {
	if o != nil && !isNil(o.PowerExceptions) {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner and assigns it to the PowerExceptions field.
func (o *UpdateNetworkSwitchSettingsRequest) SetPowerExceptions(v []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) {
	o.PowerExceptions = v
}

func (o UpdateNetworkSwitchSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.UseCombinedPower) {
		toSerialize["useCombinedPower"] = o.UseCombinedPower
	}
	if !isNil(o.PowerExceptions) {
		toSerialize["powerExceptions"] = o.PowerExceptions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchSettingsRequest struct {
	value *UpdateNetworkSwitchSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchSettingsRequest) Get() *UpdateNetworkSwitchSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchSettingsRequest) Set(val *UpdateNetworkSwitchSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchSettingsRequest(val *UpdateNetworkSwitchSettingsRequest) *NullableUpdateNetworkSwitchSettingsRequest {
	return &NullableUpdateNetworkSwitchSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


