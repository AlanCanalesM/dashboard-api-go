/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings{}

// GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings Settings related to 5Ghz band.
type GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 5Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettingsWithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettingsWithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) GetAxEnabled() bool {
	if o == nil || IsNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AxEnabled) {
		return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !IsNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings struct {
	value *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) Get() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) Set(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings {
	return &NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


