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

// NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials (Optional) The credentials of the user account to be used by the AP to bind to your Active Directory server. The Active Directory account should have permissions on all your Active Directory servers. Only valid if the splashPage is 'Password-protected with Active Directory'.
type NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials struct {
	// The logon name of the Active Directory account.
	LogonName *string `json:"logonName,omitempty"`
	// The password to the Active Directory user account.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials() *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentialsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentialsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials{}
	return &this
}

// GetLogonName returns the LogonName field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) GetLogonName() string {
	if o == nil || isNil(o.LogonName) {
		var ret string
		return ret
	}
	return *o.LogonName
}

// GetLogonNameOk returns a tuple with the LogonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) GetLogonNameOk() (*string, bool) {
	if o == nil || isNil(o.LogonName) {
    return nil, false
	}
	return o.LogonName, true
}

// HasLogonName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) HasLogonName() bool {
	if o != nil && !isNil(o.LogonName) {
		return true
	}

	return false
}

// SetLogonName gets a reference to the given string and assigns it to the LogonName field.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) SetLogonName(v string) {
	o.LogonName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LogonName) {
		toSerialize["logonName"] = o.LogonName
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials struct {
	value *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) Get() *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) Set(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials {
	return &NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


