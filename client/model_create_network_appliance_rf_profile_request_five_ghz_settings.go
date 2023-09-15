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

// checks if the CreateNetworkApplianceRfProfileRequestFiveGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceRfProfileRequestFiveGhzSettings{}

// CreateNetworkApplianceRfProfileRequestFiveGhzSettings Settings related to 5Ghz band
type CreateNetworkApplianceRfProfileRequestFiveGhzSettings struct {
	// Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewCreateNetworkApplianceRfProfileRequestFiveGhzSettings instantiates a new CreateNetworkApplianceRfProfileRequestFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceRfProfileRequestFiveGhzSettings() *CreateNetworkApplianceRfProfileRequestFiveGhzSettings {
	this := CreateNetworkApplianceRfProfileRequestFiveGhzSettings{}
	return &this
}

// NewCreateNetworkApplianceRfProfileRequestFiveGhzSettingsWithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceRfProfileRequestFiveGhzSettingsWithDefaults() *CreateNetworkApplianceRfProfileRequestFiveGhzSettings {
	this := CreateNetworkApplianceRfProfileRequestFiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) GetAxEnabled() bool {
	if o == nil || IsNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AxEnabled) {
		return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !IsNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o CreateNetworkApplianceRfProfileRequestFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceRfProfileRequestFiveGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings struct {
	value *CreateNetworkApplianceRfProfileRequestFiveGhzSettings
	isSet bool
}

func (v NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) Get() *CreateNetworkApplianceRfProfileRequestFiveGhzSettings {
	return v.value
}

func (v *NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) Set(val *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings(val *CreateNetworkApplianceRfProfileRequestFiveGhzSettings) *NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings {
	return &NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceRfProfileRequestFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


