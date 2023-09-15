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

// checks if the UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication{}

// UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication The current setting for password-based authentication.
type UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication struct {
	// Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication() *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication {
	this := UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthenticationWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthenticationWithDefaults() *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication {
	this := UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication struct {
	value *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) Get() *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) Set(val *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication(val *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) *NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication {
	return &NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


