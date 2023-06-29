/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidRequestLdap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLdap{}

// UpdateNetworkWirelessSsidRequestLdap The current setting for LDAP. Only valid if splashPage is 'Password-protected with LDAP'.
type UpdateNetworkWirelessSsidRequestLdap struct {
	// The LDAP servers to be used for authentication.
	Servers []UpdateNetworkWirelessSsidRequestLdapServersInner `json:"servers,omitempty"`
	Credentials *UpdateNetworkWirelessSsidRequestLdapCredentials `json:"credentials,omitempty"`
	// The base distinguished name of users on the LDAP server.
	BaseDistinguishedName *string `json:"baseDistinguishedName,omitempty"`
	ServerCaCertificate *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate `json:"serverCaCertificate,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLdap instantiates a new UpdateNetworkWirelessSsidRequestLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLdap() *UpdateNetworkWirelessSsidRequestLdap {
	this := UpdateNetworkWirelessSsidRequestLdap{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLdapWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLdapWithDefaults() *UpdateNetworkWirelessSsidRequestLdap {
	this := UpdateNetworkWirelessSsidRequestLdap{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetServers() []UpdateNetworkWirelessSsidRequestLdapServersInner {
	if o == nil || IsNil(o.Servers) {
		var ret []UpdateNetworkWirelessSsidRequestLdapServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetServersOk() ([]UpdateNetworkWirelessSsidRequestLdapServersInner, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []UpdateNetworkWirelessSsidRequestLdapServersInner and assigns it to the Servers field.
func (o *UpdateNetworkWirelessSsidRequestLdap) SetServers(v []UpdateNetworkWirelessSsidRequestLdapServersInner) {
	o.Servers = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetCredentials() UpdateNetworkWirelessSsidRequestLdapCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret UpdateNetworkWirelessSsidRequestLdapCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetCredentialsOk() (*UpdateNetworkWirelessSsidRequestLdapCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given UpdateNetworkWirelessSsidRequestLdapCredentials and assigns it to the Credentials field.
func (o *UpdateNetworkWirelessSsidRequestLdap) SetCredentials(v UpdateNetworkWirelessSsidRequestLdapCredentials) {
	o.Credentials = &v
}

// GetBaseDistinguishedName returns the BaseDistinguishedName field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetBaseDistinguishedName() string {
	if o == nil || IsNil(o.BaseDistinguishedName) {
		var ret string
		return ret
	}
	return *o.BaseDistinguishedName
}

// GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetBaseDistinguishedNameOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDistinguishedName) {
		return nil, false
	}
	return o.BaseDistinguishedName, true
}

// HasBaseDistinguishedName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) HasBaseDistinguishedName() bool {
	if o != nil && !IsNil(o.BaseDistinguishedName) {
		return true
	}

	return false
}

// SetBaseDistinguishedName gets a reference to the given string and assigns it to the BaseDistinguishedName field.
func (o *UpdateNetworkWirelessSsidRequestLdap) SetBaseDistinguishedName(v string) {
	o.BaseDistinguishedName = &v
}

// GetServerCaCertificate returns the ServerCaCertificate field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetServerCaCertificate() UpdateNetworkWirelessSsidRequestLdapServerCaCertificate {
	if o == nil || IsNil(o.ServerCaCertificate) {
		var ret UpdateNetworkWirelessSsidRequestLdapServerCaCertificate
		return ret
	}
	return *o.ServerCaCertificate
}

// GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) GetServerCaCertificateOk() (*UpdateNetworkWirelessSsidRequestLdapServerCaCertificate, bool) {
	if o == nil || IsNil(o.ServerCaCertificate) {
		return nil, false
	}
	return o.ServerCaCertificate, true
}

// HasServerCaCertificate returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdap) HasServerCaCertificate() bool {
	if o != nil && !IsNil(o.ServerCaCertificate) {
		return true
	}

	return false
}

// SetServerCaCertificate gets a reference to the given UpdateNetworkWirelessSsidRequestLdapServerCaCertificate and assigns it to the ServerCaCertificate field.
func (o *UpdateNetworkWirelessSsidRequestLdap) SetServerCaCertificate(v UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) {
	o.ServerCaCertificate = &v
}

func (o UpdateNetworkWirelessSsidRequestLdap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLdap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.BaseDistinguishedName) {
		toSerialize["baseDistinguishedName"] = o.BaseDistinguishedName
	}
	if !IsNil(o.ServerCaCertificate) {
		toSerialize["serverCaCertificate"] = o.ServerCaCertificate
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLdap struct {
	value *UpdateNetworkWirelessSsidRequestLdap
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLdap) Get() *UpdateNetworkWirelessSsidRequestLdap {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdap) Set(val *UpdateNetworkWirelessSsidRequestLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLdap(val *UpdateNetworkWirelessSsidRequestLdap) *NullableUpdateNetworkWirelessSsidRequestLdap {
	return &NullableUpdateNetworkWirelessSsidRequestLdap{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


