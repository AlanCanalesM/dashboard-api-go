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

// GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup The staged upgrade group
type GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup struct {
	// Id of the staged upgrade group
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup{}
	return &this
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroupWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroupWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup struct {
	value *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) Get() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) Set(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup {
	return &NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStagedGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


