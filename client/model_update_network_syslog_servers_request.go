/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkSyslogServersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSyslogServersRequest{}

// UpdateNetworkSyslogServersRequest struct for UpdateNetworkSyslogServersRequest
type UpdateNetworkSyslogServersRequest struct {
	// A list of the syslog servers for this network
	Servers []UpdateNetworkSyslogServersRequestServersInner `json:"servers"`
}

// NewUpdateNetworkSyslogServersRequest instantiates a new UpdateNetworkSyslogServersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSyslogServersRequest(servers []UpdateNetworkSyslogServersRequestServersInner) *UpdateNetworkSyslogServersRequest {
	this := UpdateNetworkSyslogServersRequest{}
	this.Servers = servers
	return &this
}

// NewUpdateNetworkSyslogServersRequestWithDefaults instantiates a new UpdateNetworkSyslogServersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSyslogServersRequestWithDefaults() *UpdateNetworkSyslogServersRequest {
	this := UpdateNetworkSyslogServersRequest{}
	return &this
}

// GetServers returns the Servers field value
func (o *UpdateNetworkSyslogServersRequest) GetServers() []UpdateNetworkSyslogServersRequestServersInner {
	if o == nil {
		var ret []UpdateNetworkSyslogServersRequestServersInner
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSyslogServersRequest) GetServersOk() ([]UpdateNetworkSyslogServersRequestServersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *UpdateNetworkSyslogServersRequest) SetServers(v []UpdateNetworkSyslogServersRequestServersInner) {
	o.Servers = v
}

func (o UpdateNetworkSyslogServersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSyslogServersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["servers"] = o.Servers
	return toSerialize, nil
}

type NullableUpdateNetworkSyslogServersRequest struct {
	value *UpdateNetworkSyslogServersRequest
	isSet bool
}

func (v NullableUpdateNetworkSyslogServersRequest) Get() *UpdateNetworkSyslogServersRequest {
	return v.value
}

func (v *NullableUpdateNetworkSyslogServersRequest) Set(val *UpdateNetworkSyslogServersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSyslogServersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSyslogServersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSyslogServersRequest(val *UpdateNetworkSyslogServersRequest) *NullableUpdateNetworkSyslogServersRequest {
	return &NullableUpdateNetworkSyslogServersRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSyslogServersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSyslogServersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


