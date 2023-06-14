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

// UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner struct for UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner
type UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner struct {
	// The name of the client which has fixed IP address
	Name string `json:"name"`
	// The MAC address of the client which has fixed IP address
	Mac string `json:"mac"`
	// The IP address of the client which has fixed IP address assigned to it
	Ip string `json:"ip"`
}

// NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner(name string, mac string, ip string) *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner {
	this := UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner{}
	this.Name = name
	this.Mac = mac
	this.Ip = ip
	return &this
}

// NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInnerWithDefaults instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInnerWithDefaults() *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner {
	this := UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) SetName(v string) {
	o.Name = v
}

// GetMac returns the Mac field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) SetMac(v string) {
	o.Mac = v
}

// GetIp returns the Ip field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) GetIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) SetIp(v string) {
	o.Ip = v
}

func (o UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner struct {
	value *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner
	isSet bool
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) Get() *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner {
	return v.value
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) Set(val *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner(val *UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner {
	return &NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


