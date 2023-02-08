/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateOrganizationAdaptivePolicySettingsRequest struct for UpdateOrganizationAdaptivePolicySettingsRequest
type UpdateOrganizationAdaptivePolicySettingsRequest struct {
	// List of network IDs with adaptive policy enabled
	EnabledNetworks []string `json:"enabledNetworks,omitempty"`
}

// NewUpdateOrganizationAdaptivePolicySettingsRequest instantiates a new UpdateOrganizationAdaptivePolicySettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAdaptivePolicySettingsRequest() *UpdateOrganizationAdaptivePolicySettingsRequest {
	this := UpdateOrganizationAdaptivePolicySettingsRequest{}
	return &this
}

// NewUpdateOrganizationAdaptivePolicySettingsRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAdaptivePolicySettingsRequestWithDefaults() *UpdateOrganizationAdaptivePolicySettingsRequest {
	this := UpdateOrganizationAdaptivePolicySettingsRequest{}
	return &this
}

// GetEnabledNetworks returns the EnabledNetworks field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicySettingsRequest) GetEnabledNetworks() []string {
	if o == nil || isNil(o.EnabledNetworks) {
		var ret []string
		return ret
	}
	return o.EnabledNetworks
}

// GetEnabledNetworksOk returns a tuple with the EnabledNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicySettingsRequest) GetEnabledNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.EnabledNetworks) {
    return nil, false
	}
	return o.EnabledNetworks, true
}

// HasEnabledNetworks returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicySettingsRequest) HasEnabledNetworks() bool {
	if o != nil && !isNil(o.EnabledNetworks) {
		return true
	}

	return false
}

// SetEnabledNetworks gets a reference to the given []string and assigns it to the EnabledNetworks field.
func (o *UpdateOrganizationAdaptivePolicySettingsRequest) SetEnabledNetworks(v []string) {
	o.EnabledNetworks = v
}

func (o UpdateOrganizationAdaptivePolicySettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnabledNetworks) {
		toSerialize["enabledNetworks"] = o.EnabledNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationAdaptivePolicySettingsRequest struct {
	value *UpdateOrganizationAdaptivePolicySettingsRequest
	isSet bool
}

func (v NullableUpdateOrganizationAdaptivePolicySettingsRequest) Get() *UpdateOrganizationAdaptivePolicySettingsRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAdaptivePolicySettingsRequest) Set(val *UpdateOrganizationAdaptivePolicySettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAdaptivePolicySettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAdaptivePolicySettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAdaptivePolicySettingsRequest(val *UpdateOrganizationAdaptivePolicySettingsRequest) *NullableUpdateOrganizationAdaptivePolicySettingsRequest {
	return &NullableUpdateOrganizationAdaptivePolicySettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAdaptivePolicySettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAdaptivePolicySettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


