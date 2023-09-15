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

// checks if the UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard{}

// UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard IP source address spoofing settings
type UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard struct {
	// Mode of protection
	Mode *string `json:"mode,omitempty"`
}

// NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard instantiates a new UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard {
	this := UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard{}
	return &this
}

// NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuardWithDefaults instantiates a new UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuardWithDefaults() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard {
	this := UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) SetMode(v string) {
	o.Mode = &v
}

func (o UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard struct {
	value *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) Get() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) Set(val *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard(val *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard {
	return &NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


