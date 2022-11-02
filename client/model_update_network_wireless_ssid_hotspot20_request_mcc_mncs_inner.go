/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner struct for UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner
type UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner struct {
	// MCC value
	Mcc *string `json:"mcc,omitempty"`
	// MNC value
	Mnc *string `json:"mnc,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner instantiates a new UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner() *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestMccMncsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestMccMncsInnerWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) GetMcc() string {
	if o == nil || isNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) GetMccOk() (*string, bool) {
	if o == nil || isNil(o.Mcc) {
    return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) HasMcc() bool {
	if o != nil && !isNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) GetMnc() string {
	if o == nil || isNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) GetMncOk() (*string, bool) {
	if o == nil || isNil(o.Mnc) {
    return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) HasMnc() bool {
	if o != nil && !isNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) SetMnc(v string) {
	o.Mnc = &v
}

func (o UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !isNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner struct {
	value *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) Get() *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) Set(val *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner(val *UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) *NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner {
	return &NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


