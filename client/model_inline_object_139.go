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

// InlineObject139 struct for InlineObject139
type InlineObject139 struct {
	// A list of the syslog servers for this network
	Servers []NetworksNetworkIdSyslogServersServers `json:"servers"`
}

// NewInlineObject139 instantiates a new InlineObject139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject139(servers []NetworksNetworkIdSyslogServersServers) *InlineObject139 {
	this := InlineObject139{}
	this.Servers = servers
	return &this
}

// NewInlineObject139WithDefaults instantiates a new InlineObject139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject139WithDefaults() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// GetServers returns the Servers field value
func (o *InlineObject139) GetServers() []NetworksNetworkIdSyslogServersServers {
	if o == nil {
		var ret []NetworksNetworkIdSyslogServersServers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetServersOk() ([]NetworksNetworkIdSyslogServersServers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *InlineObject139) SetServers(v []NetworksNetworkIdSyslogServersServers) {
	o.Servers = v
}

func (o InlineObject139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject139 struct {
	value *InlineObject139
	isSet bool
}

func (v NullableInlineObject139) Get() *InlineObject139 {
	return v.value
}

func (v *NullableInlineObject139) Set(val *InlineObject139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject139(val *InlineObject139) *NullableInlineObject139 {
	return &NullableInlineObject139{value: val, isSet: true}
}

func (v NullableInlineObject139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


