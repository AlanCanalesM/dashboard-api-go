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

// GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice Represents the details of an imported device.
type GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice struct {
	// The url to the device details page within dashboard.
	Url *string `json:"url,omitempty"`
	// Whether or not the device was successfully created in dashboard.
	Created *bool `json:"created,omitempty"`
	// Represents the current state of importing the device.
	Status *string `json:"status,omitempty"`
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice instantiates a new GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice {
	this := GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice{}
	return &this
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDeviceWithDefaults instantiates a new GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDeviceWithDefaults() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice {
	this := GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) SetUrl(v string) {
	o.Url = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetCreated() bool {
	if o == nil || isNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetCreatedOk() (*bool, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) SetCreated(v bool) {
	o.Created = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) SetStatus(v string) {
	o.Status = &v
}

func (o GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice struct {
	value *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice
	isSet bool
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) Get() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice {
	return v.value
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) Set(val *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice(val *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice {
	return &NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice{value: val, isSet: true}
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


