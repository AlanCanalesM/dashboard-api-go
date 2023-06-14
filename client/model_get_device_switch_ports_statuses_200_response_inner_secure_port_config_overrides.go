/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides The configuration overrides applied to this port when Secure Port is active.
type GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides struct {
	// The type of the  ('trunk' or 'access').
	Type *string `json:"type,omitempty"`
	// The VLAN of the . A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the . Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the . Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides{}
	return &this
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverridesWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverridesWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVoiceVlan() int32 {
	if o == nil || isNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || isNil(o.VoiceVlan) {
    return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasVoiceVlan() bool {
	if o != nil && !isNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetAllowedVlans() string {
	if o == nil || isNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetAllowedVlansOk() (*string, bool) {
	if o == nil || isNil(o.AllowedVlans) {
    return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasAllowedVlans() bool {
	if o != nil && !isNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

func (o GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.VoiceVlan) {
		toSerialize["voiceVlan"] = o.VoiceVlan
	}
	if !isNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides struct {
	value *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides
	isSet bool
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) Get() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides {
	return v.value
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) Set(val *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides(val *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) *NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides {
	return &NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


