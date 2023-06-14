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

// UpdateNetworkApplianceFirewallSettingsRequest struct for UpdateNetworkApplianceFirewallSettingsRequest
type UpdateNetworkApplianceFirewallSettingsRequest struct {
	SpoofingProtection *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection `json:"spoofingProtection,omitempty"`
}

// NewUpdateNetworkApplianceFirewallSettingsRequest instantiates a new UpdateNetworkApplianceFirewallSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallSettingsRequest() *UpdateNetworkApplianceFirewallSettingsRequest {
	this := UpdateNetworkApplianceFirewallSettingsRequest{}
	return &this
}

// NewUpdateNetworkApplianceFirewallSettingsRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallSettingsRequestWithDefaults() *UpdateNetworkApplianceFirewallSettingsRequest {
	this := UpdateNetworkApplianceFirewallSettingsRequest{}
	return &this
}

// GetSpoofingProtection returns the SpoofingProtection field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallSettingsRequest) GetSpoofingProtection() UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection {
	if o == nil || isNil(o.SpoofingProtection) {
		var ret UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection
		return ret
	}
	return *o.SpoofingProtection
}

// GetSpoofingProtectionOk returns a tuple with the SpoofingProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequest) GetSpoofingProtectionOk() (*UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection, bool) {
	if o == nil || isNil(o.SpoofingProtection) {
    return nil, false
	}
	return o.SpoofingProtection, true
}

// HasSpoofingProtection returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequest) HasSpoofingProtection() bool {
	if o != nil && !isNil(o.SpoofingProtection) {
		return true
	}

	return false
}

// SetSpoofingProtection gets a reference to the given UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection and assigns it to the SpoofingProtection field.
func (o *UpdateNetworkApplianceFirewallSettingsRequest) SetSpoofingProtection(v UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) {
	o.SpoofingProtection = &v
}

func (o UpdateNetworkApplianceFirewallSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SpoofingProtection) {
		toSerialize["spoofingProtection"] = o.SpoofingProtection
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceFirewallSettingsRequest struct {
	value *UpdateNetworkApplianceFirewallSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequest) Get() *UpdateNetworkApplianceFirewallSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequest) Set(val *UpdateNetworkApplianceFirewallSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallSettingsRequest(val *UpdateNetworkApplianceFirewallSettingsRequest) *NullableUpdateNetworkApplianceFirewallSettingsRequest {
	return &NullableUpdateNetworkApplianceFirewallSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


