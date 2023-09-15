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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication VTY AAA authentication
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication struct {
	Group *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup `json:"group,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) GetGroup() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup {
	if o == nil || IsNil(o.Group) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) GetGroupOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup and assigns it to the Group field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) SetGroup(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthenticationGroup) {
	o.Group = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


