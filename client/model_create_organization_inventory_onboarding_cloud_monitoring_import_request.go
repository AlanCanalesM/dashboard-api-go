/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{}

// CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct for CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
type CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	// A set of device imports to commit
	Devices []CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner `json:"devices"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(devices []CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{}
	this.Devices = devices
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{}
	return &this
}

// GetDevices returns the Devices field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) GetDevices() []CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner {
	if o == nil {
		var ret []CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) GetDevicesOk() ([]CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) SetDevices(v []CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) {
	o.Devices = v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(val *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


