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

// checks if the GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup{}

// GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup The Staged Upgrade Group
type GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup struct {
	// Id of the Staged Upgrade Group
	Id *string `json:"id,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup instantiates a new GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup() *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup {
	this := GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroupWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroupWithDefaults() *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup {
	this := GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) SetDescription(v string) {
	o.Description = &v
}

func (o GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup struct {
	value *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) Get() *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) Set(val *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup(val *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup {
	return &NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


