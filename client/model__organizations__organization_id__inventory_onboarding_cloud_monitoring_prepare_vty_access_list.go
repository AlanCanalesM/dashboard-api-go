/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList AccessList details
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList struct {
	VtyIn *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn `json:"vtyIn,omitempty"`
	VtyOut *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut `json:"vtyOut,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList{}
	return &this
}

// GetVtyIn returns the VtyIn field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) GetVtyIn() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn {
	if o == nil || isNil(o.VtyIn) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn
		return ret
	}
	return *o.VtyIn
}

// GetVtyInOk returns a tuple with the VtyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) GetVtyInOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn, bool) {
	if o == nil || isNil(o.VtyIn) {
    return nil, false
	}
	return o.VtyIn, true
}

// HasVtyIn returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) HasVtyIn() bool {
	if o != nil && !isNil(o.VtyIn) {
		return true
	}

	return false
}

// SetVtyIn gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn and assigns it to the VtyIn field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) SetVtyIn(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyIn) {
	o.VtyIn = &v
}

// GetVtyOut returns the VtyOut field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) GetVtyOut() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut {
	if o == nil || isNil(o.VtyOut) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut
		return ret
	}
	return *o.VtyOut
}

// GetVtyOutOk returns a tuple with the VtyOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) GetVtyOutOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut, bool) {
	if o == nil || isNil(o.VtyOut) {
    return nil, false
	}
	return o.VtyOut, true
}

// HasVtyOut returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) HasVtyOut() bool {
	if o != nil && !isNil(o.VtyOut) {
		return true
	}

	return false
}

// SetVtyOut gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut and assigns it to the VtyOut field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) SetVtyOut(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) {
	o.VtyOut = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VtyIn) {
		toSerialize["vtyIn"] = o.VtyIn
	}
	if !isNil(o.VtyOut) {
		toSerialize["vtyOut"] = o.VtyOut
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


