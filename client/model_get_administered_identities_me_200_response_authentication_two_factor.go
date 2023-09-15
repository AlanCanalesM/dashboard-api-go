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

// checks if the GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor{}

// GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor TwoFactor authentication
type GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor struct {
	// If twoFactor authentication is enabled for this user
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor() *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactorWithDefaults instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactorWithDefaults() *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor struct {
	value *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) Get() *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) Set(val *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor(val *GetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor {
	return &NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationTwoFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


