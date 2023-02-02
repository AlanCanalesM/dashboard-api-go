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

// DevicesSerialCellularGatewayLanFixedIpAssignments struct for DevicesSerialCellularGatewayLanFixedIpAssignments
type DevicesSerialCellularGatewayLanFixedIpAssignments struct {
	// A descriptive name of the assignment
	Name *string `json:"name,omitempty"`
	// The IP address you want to assign to a specific server or device
	Ip string `json:"ip"`
	// The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Mac string `json:"mac"`
}

// NewDevicesSerialCellularGatewayLanFixedIpAssignments instantiates a new DevicesSerialCellularGatewayLanFixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularGatewayLanFixedIpAssignments(ip string, mac string) *DevicesSerialCellularGatewayLanFixedIpAssignments {
	this := DevicesSerialCellularGatewayLanFixedIpAssignments{}
	this.Ip = ip
	this.Mac = mac
	return &this
}

// NewDevicesSerialCellularGatewayLanFixedIpAssignmentsWithDefaults instantiates a new DevicesSerialCellularGatewayLanFixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularGatewayLanFixedIpAssignmentsWithDefaults() *DevicesSerialCellularGatewayLanFixedIpAssignments {
	this := DevicesSerialCellularGatewayLanFixedIpAssignments{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) SetIp(v string) {
	o.Ip = v
}

// GetMac returns the Mac field value
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *DevicesSerialCellularGatewayLanFixedIpAssignments) SetMac(v string) {
	o.Mac = v
}

func (o DevicesSerialCellularGatewayLanFixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if true {
		toSerialize["mac"] = o.Mac
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularGatewayLanFixedIpAssignments struct {
	value *DevicesSerialCellularGatewayLanFixedIpAssignments
	isSet bool
}

func (v NullableDevicesSerialCellularGatewayLanFixedIpAssignments) Get() *DevicesSerialCellularGatewayLanFixedIpAssignments {
	return v.value
}

func (v *NullableDevicesSerialCellularGatewayLanFixedIpAssignments) Set(val *DevicesSerialCellularGatewayLanFixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularGatewayLanFixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularGatewayLanFixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularGatewayLanFixedIpAssignments(val *DevicesSerialCellularGatewayLanFixedIpAssignments) *NullableDevicesSerialCellularGatewayLanFixedIpAssignments {
	return &NullableDevicesSerialCellularGatewayLanFixedIpAssignments{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularGatewayLanFixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularGatewayLanFixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


