/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup The Staged Upgrade Group containing the name and ID
type UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup struct {
	// ID of the Staged Upgrade Group
	Id string `json:"id"`
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup(id string) *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup{}
	this.Id = id
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroupWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroupWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) SetId(v string) {
	o.Id = v
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup struct {
	value *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) Get() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) Set(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup {
	return &NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


