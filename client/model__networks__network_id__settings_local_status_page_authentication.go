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

// NetworksNetworkIdSettingsLocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
type NetworksNetworkIdSettingsLocalStatusPageAuthentication struct {
	// Enables / disables the authentication on Local Status page(s).
	Enabled *bool `json:"enabled,omitempty"`
	// The password used for Local Status Page(s). Set this to null to clear the password.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdSettingsLocalStatusPageAuthentication instantiates a new NetworksNetworkIdSettingsLocalStatusPageAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSettingsLocalStatusPageAuthentication() *NetworksNetworkIdSettingsLocalStatusPageAuthentication {
	this := NetworksNetworkIdSettingsLocalStatusPageAuthentication{}
	return &this
}

// NewNetworksNetworkIdSettingsLocalStatusPageAuthenticationWithDefaults instantiates a new NetworksNetworkIdSettingsLocalStatusPageAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSettingsLocalStatusPageAuthenticationWithDefaults() *NetworksNetworkIdSettingsLocalStatusPageAuthentication {
	this := NetworksNetworkIdSettingsLocalStatusPageAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdSettingsLocalStatusPageAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdSettingsLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication struct {
	value *NetworksNetworkIdSettingsLocalStatusPageAuthentication
	isSet bool
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) Get() *NetworksNetworkIdSettingsLocalStatusPageAuthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) Set(val *NetworksNetworkIdSettingsLocalStatusPageAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSettingsLocalStatusPageAuthentication(val *NetworksNetworkIdSettingsLocalStatusPageAuthentication) *NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication {
	return &NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPageAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


