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

// GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged Staged upgrade
type GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged struct {
	Group *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup `json:"group,omitempty"`
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged{}
	return &this
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) GetGroup() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup {
	if o == nil || isNil(o.Group) {
		var ret GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) GetGroupOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup and assigns it to the Group field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) SetGroup(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) {
	o.Group = &v
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged struct {
	value *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) Get() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) Set(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged {
	return &NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


