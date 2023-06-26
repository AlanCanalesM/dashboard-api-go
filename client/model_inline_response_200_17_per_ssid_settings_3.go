/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20017PerSsidSettings3 Settings for SSID 3.
type InlineResponse20017PerSsidSettings3 struct {
	// Band mode of this SSID
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewInlineResponse20017PerSsidSettings3 instantiates a new InlineResponse20017PerSsidSettings3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017PerSsidSettings3() *InlineResponse20017PerSsidSettings3 {
	this := InlineResponse20017PerSsidSettings3{}
	return &this
}

// NewInlineResponse20017PerSsidSettings3WithDefaults instantiates a new InlineResponse20017PerSsidSettings3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017PerSsidSettings3WithDefaults() *InlineResponse20017PerSsidSettings3 {
	this := InlineResponse20017PerSsidSettings3{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *InlineResponse20017PerSsidSettings3) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017PerSsidSettings3) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *InlineResponse20017PerSsidSettings3) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *InlineResponse20017PerSsidSettings3) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *InlineResponse20017PerSsidSettings3) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017PerSsidSettings3) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *InlineResponse20017PerSsidSettings3) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *InlineResponse20017PerSsidSettings3) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o InlineResponse20017PerSsidSettings3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017PerSsidSettings3 struct {
	value *InlineResponse20017PerSsidSettings3
	isSet bool
}

func (v NullableInlineResponse20017PerSsidSettings3) Get() *InlineResponse20017PerSsidSettings3 {
	return v.value
}

func (v *NullableInlineResponse20017PerSsidSettings3) Set(val *InlineResponse20017PerSsidSettings3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017PerSsidSettings3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017PerSsidSettings3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017PerSsidSettings3(val *InlineResponse20017PerSsidSettings3) *NullableInlineResponse20017PerSsidSettings3 {
	return &NullableInlineResponse20017PerSsidSettings3{value: val, isSet: true}
}

func (v NullableInlineResponse20017PerSsidSettings3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017PerSsidSettings3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


