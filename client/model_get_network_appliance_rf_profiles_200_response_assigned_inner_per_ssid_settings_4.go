/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4{}

// GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 Settings for SSID 4.
type GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 struct {
	// Band mode of this SSID
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4WithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4WithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 struct {
	value *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) Get() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) Set(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 {
	return &NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


