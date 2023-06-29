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

// checks if the GetNetworkFirmwareUpgradesStagedGroups200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedGroups200ResponseInner{}

// GetNetworkFirmwareUpgradesStagedGroups200ResponseInner struct for GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
type GetNetworkFirmwareUpgradesStagedGroups200ResponseInner struct {
	// Id of staged upgrade group
	GroupId *string `json:"groupId,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault *bool `json:"isDefault,omitempty"`
	AssignedDevices *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices `json:"assignedDevices,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInner instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInner() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInner{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInner{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetAssignedDevices() GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices {
	if o == nil || IsNil(o.AssignedDevices) {
		var ret GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetAssignedDevicesOk() (*GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices, bool) {
	if o == nil || IsNil(o.AssignedDevices) {
		return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasAssignedDevices() bool {
	if o != nil && !IsNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices and assigns it to the AssignedDevices field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetAssignedDevices(v GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices) {
	o.AssignedDevices = &v
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner struct {
	value *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) Get() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) Set(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner {
	return &NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


