/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner{}

// CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner struct for CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner
type CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner struct {
	// Serial of the device
	Serial string `json:"serial"`
	// Name of the device
	Name *string `json:"name,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner(serial string) *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner{}
	this.Serial = serial
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInnerWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInnerWithDefaults() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) SetSerial(v string) {
	o.Serial = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) SetName(v string) {
	o.Name = &v
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner struct {
	value *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) Get() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) Set(val *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner(val *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner {
	return &NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


