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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn VTY in ACL
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn struct {
	// Name
	Name *string `json:"name,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyInWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyInWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) SetName(v string) {
	o.Name = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessListVtyIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


