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

// UpdateNetworkWirelessSsidRequestLdapServersInner struct for UpdateNetworkWirelessSsidRequestLdapServersInner
type UpdateNetworkWirelessSsidRequestLdapServersInner struct {
	// IP address of your LDAP server.
	Host string `json:"host"`
	// UDP port the LDAP server listens on.
	Port int32 `json:"port"`
}

// NewUpdateNetworkWirelessSsidRequestLdapServersInner instantiates a new UpdateNetworkWirelessSsidRequestLdapServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLdapServersInner(host string, port int32) *UpdateNetworkWirelessSsidRequestLdapServersInner {
	this := UpdateNetworkWirelessSsidRequestLdapServersInner{}
	this.Host = host
	this.Port = port
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLdapServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdapServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLdapServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestLdapServersInner {
	this := UpdateNetworkWirelessSsidRequestLdapServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *UpdateNetworkWirelessSsidRequestLdapServersInner) SetPort(v int32) {
	o.Port = v
}

func (o UpdateNetworkWirelessSsidRequestLdapServersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidRequestLdapServersInner struct {
	value *UpdateNetworkWirelessSsidRequestLdapServersInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServersInner) Get() *UpdateNetworkWirelessSsidRequestLdapServersInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServersInner) Set(val *UpdateNetworkWirelessSsidRequestLdapServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLdapServersInner(val *UpdateNetworkWirelessSsidRequestLdapServersInner) *NullableUpdateNetworkWirelessSsidRequestLdapServersInner {
	return &NullableUpdateNetworkWirelessSsidRequestLdapServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


