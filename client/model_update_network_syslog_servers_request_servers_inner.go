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

// checks if the UpdateNetworkSyslogServersRequestServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSyslogServersRequestServersInner{}

// UpdateNetworkSyslogServersRequestServersInner struct for UpdateNetworkSyslogServersRequestServersInner
type UpdateNetworkSyslogServersRequestServersInner struct {
	// The IP address of the syslog server
	Host string `json:"host"`
	// The port of the syslog server
	Port int32 `json:"port"`
	// A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
	Roles []string `json:"roles"`
}

// NewUpdateNetworkSyslogServersRequestServersInner instantiates a new UpdateNetworkSyslogServersRequestServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSyslogServersRequestServersInner(host string, port int32, roles []string) *UpdateNetworkSyslogServersRequestServersInner {
	this := UpdateNetworkSyslogServersRequestServersInner{}
	this.Host = host
	this.Port = port
	this.Roles = roles
	return &this
}

// NewUpdateNetworkSyslogServersRequestServersInnerWithDefaults instantiates a new UpdateNetworkSyslogServersRequestServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSyslogServersRequestServersInnerWithDefaults() *UpdateNetworkSyslogServersRequestServersInner {
	this := UpdateNetworkSyslogServersRequestServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkSyslogServersRequestServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSyslogServersRequestServersInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkSyslogServersRequestServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *UpdateNetworkSyslogServersRequestServersInner) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSyslogServersRequestServersInner) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *UpdateNetworkSyslogServersRequestServersInner) SetPort(v int32) {
	o.Port = v
}

// GetRoles returns the Roles field value
func (o *UpdateNetworkSyslogServersRequestServersInner) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSyslogServersRequestServersInner) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UpdateNetworkSyslogServersRequestServersInner) SetRoles(v []string) {
	o.Roles = v
}

func (o UpdateNetworkSyslogServersRequestServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSyslogServersRequestServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableUpdateNetworkSyslogServersRequestServersInner struct {
	value *UpdateNetworkSyslogServersRequestServersInner
	isSet bool
}

func (v NullableUpdateNetworkSyslogServersRequestServersInner) Get() *UpdateNetworkSyslogServersRequestServersInner {
	return v.value
}

func (v *NullableUpdateNetworkSyslogServersRequestServersInner) Set(val *UpdateNetworkSyslogServersRequestServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSyslogServersRequestServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSyslogServersRequestServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSyslogServersRequestServersInner(val *UpdateNetworkSyslogServersRequestServersInner) *NullableUpdateNetworkSyslogServersRequestServersInner {
	return &NullableUpdateNetworkSyslogServersRequestServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSyslogServersRequestServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSyslogServersRequestServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


