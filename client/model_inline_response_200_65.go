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

// InlineResponse20065 struct for InlineResponse20065
type InlineResponse20065 struct {
	// List of the syslog servers for this network
	Servers []InlineResponse20065Servers `json:"servers,omitempty"`
}

// NewInlineResponse20065 instantiates a new InlineResponse20065 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065WithDefaults() *InlineResponse20065 {
	this := InlineResponse20065{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *InlineResponse20065) GetServers() []InlineResponse20065Servers {
	if o == nil || isNil(o.Servers) {
		var ret []InlineResponse20065Servers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065) GetServersOk() ([]InlineResponse20065Servers, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *InlineResponse20065) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []InlineResponse20065Servers and assigns it to the Servers field.
func (o *InlineResponse20065) SetServers(v []InlineResponse20065Servers) {
	o.Servers = v
}

func (o InlineResponse20065) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20065 struct {
	value *InlineResponse20065
	isSet bool
}

func (v NullableInlineResponse20065) Get() *InlineResponse20065 {
	return v.value
}

func (v *NullableInlineResponse20065) Set(val *InlineResponse20065) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065(val *InlineResponse20065) *NullableInlineResponse20065 {
	return &NullableInlineResponse20065{value: val, isSet: true}
}

func (v NullableInlineResponse20065) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


