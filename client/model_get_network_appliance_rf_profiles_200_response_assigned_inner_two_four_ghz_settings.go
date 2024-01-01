/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings{}

// GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings Settings related to 2.4Ghz band.
type GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 2.4Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettingsWithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettingsWithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || IsNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AxEnabled) {
		return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !IsNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings struct {
	value *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) Get() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) Set(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings {
	return &NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


