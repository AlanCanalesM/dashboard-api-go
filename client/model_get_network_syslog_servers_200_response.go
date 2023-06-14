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

// GetNetworkSyslogServers200Response struct for GetNetworkSyslogServers200Response
type GetNetworkSyslogServers200Response struct {
	// List of the syslog servers for this network
	Servers []GetNetworkSyslogServers200ResponseServersInner `json:"servers,omitempty"`
}

// NewGetNetworkSyslogServers200Response instantiates a new GetNetworkSyslogServers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSyslogServers200Response() *GetNetworkSyslogServers200Response {
	this := GetNetworkSyslogServers200Response{}
	return &this
}

// NewGetNetworkSyslogServers200ResponseWithDefaults instantiates a new GetNetworkSyslogServers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSyslogServers200ResponseWithDefaults() *GetNetworkSyslogServers200Response {
	this := GetNetworkSyslogServers200Response{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *GetNetworkSyslogServers200Response) GetServers() []GetNetworkSyslogServers200ResponseServersInner {
	if o == nil || isNil(o.Servers) {
		var ret []GetNetworkSyslogServers200ResponseServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSyslogServers200Response) GetServersOk() ([]GetNetworkSyslogServers200ResponseServersInner, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *GetNetworkSyslogServers200Response) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []GetNetworkSyslogServers200ResponseServersInner and assigns it to the Servers field.
func (o *GetNetworkSyslogServers200Response) SetServers(v []GetNetworkSyslogServers200ResponseServersInner) {
	o.Servers = v
}

func (o GetNetworkSyslogServers200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSyslogServers200Response struct {
	value *GetNetworkSyslogServers200Response
	isSet bool
}

func (v NullableGetNetworkSyslogServers200Response) Get() *GetNetworkSyslogServers200Response {
	return v.value
}

func (v *NullableGetNetworkSyslogServers200Response) Set(val *GetNetworkSyslogServers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSyslogServers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSyslogServers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSyslogServers200Response(val *GetNetworkSyslogServers200Response) *NullableGetNetworkSyslogServers200Response {
	return &NullableGetNetworkSyslogServers200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSyslogServers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSyslogServers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


