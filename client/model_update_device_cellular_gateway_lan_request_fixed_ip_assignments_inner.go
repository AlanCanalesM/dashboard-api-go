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

// checks if the UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner{}

// UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner struct for UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner
type UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner struct {
	// A descriptive name of the assignment
	Name *string `json:"name,omitempty"`
	// The IP address you want to assign to a specific server or device
	Ip string `json:"ip"`
	// The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Mac string `json:"mac"`
}

// NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner instantiates a new UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner(ip string, mac string) *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner {
	this := UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner{}
	this.Ip = ip
	this.Mac = mac
	return &this
}

// NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInnerWithDefaults instantiates a new UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInnerWithDefaults() *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner {
	this := UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetIp(v string) {
	o.Ip = v
}

// GetMac returns the Mac field value
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetMac(v string) {
	o.Mac = v
}

func (o UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["ip"] = o.Ip
	toSerialize["mac"] = o.Mac
	return toSerialize, nil
}

type NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner struct {
	value *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner
	isSet bool
}

func (v NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) Get() *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner {
	return v.value
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) Set(val *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner(val *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) *NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner {
	return &NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


