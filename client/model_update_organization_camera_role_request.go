/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateOrganizationCameraRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationCameraRoleRequest{}

// UpdateOrganizationCameraRoleRequest struct for UpdateOrganizationCameraRoleRequest
type UpdateOrganizationCameraRoleRequest struct {
	// The name of the new role. Must be unique.
	Name *string `json:"name,omitempty"`
	// Device tag on which this specified permission is applied.
	AppliedOnDevices []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner `json:"appliedOnDevices,omitempty"`
	// Network tag on which this specified permission is applied.
	AppliedOnNetworks []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner `json:"appliedOnNetworks,omitempty"`
	// Permissions to be applied org wide.
	AppliedOrgWide []CreateOrganizationCameraRoleRequestAppliedOrgWideInner `json:"appliedOrgWide,omitempty"`
}

// NewUpdateOrganizationCameraRoleRequest instantiates a new UpdateOrganizationCameraRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationCameraRoleRequest() *UpdateOrganizationCameraRoleRequest {
	this := UpdateOrganizationCameraRoleRequest{}
	return &this
}

// NewUpdateOrganizationCameraRoleRequestWithDefaults instantiates a new UpdateOrganizationCameraRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationCameraRoleRequestWithDefaults() *UpdateOrganizationCameraRoleRequest {
	this := UpdateOrganizationCameraRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationCameraRoleRequest) SetName(v string) {
	o.Name = &v
}

// GetAppliedOnDevices returns the AppliedOnDevices field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnDevices() []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner {
	if o == nil || IsNil(o.AppliedOnDevices) {
		var ret []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner
		return ret
	}
	return o.AppliedOnDevices
}

// GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnDevicesOk() ([]CreateOrganizationCameraRoleRequestAppliedOnDevicesInner, bool) {
	if o == nil || IsNil(o.AppliedOnDevices) {
		return nil, false
	}
	return o.AppliedOnDevices, true
}

// HasAppliedOnDevices returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOnDevices() bool {
	if o != nil && !IsNil(o.AppliedOnDevices) {
		return true
	}

	return false
}

// SetAppliedOnDevices gets a reference to the given []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner and assigns it to the AppliedOnDevices field.
func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOnDevices(v []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) {
	o.AppliedOnDevices = v
}

// GetAppliedOnNetworks returns the AppliedOnNetworks field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnNetworks() []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner {
	if o == nil || IsNil(o.AppliedOnNetworks) {
		var ret []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner
		return ret
	}
	return o.AppliedOnNetworks
}

// GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnNetworksOk() ([]UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner, bool) {
	if o == nil || IsNil(o.AppliedOnNetworks) {
		return nil, false
	}
	return o.AppliedOnNetworks, true
}

// HasAppliedOnNetworks returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOnNetworks() bool {
	if o != nil && !IsNil(o.AppliedOnNetworks) {
		return true
	}

	return false
}

// SetAppliedOnNetworks gets a reference to the given []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner and assigns it to the AppliedOnNetworks field.
func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOnNetworks(v []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) {
	o.AppliedOnNetworks = v
}

// GetAppliedOrgWide returns the AppliedOrgWide field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOrgWide() []CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	if o == nil || IsNil(o.AppliedOrgWide) {
		var ret []CreateOrganizationCameraRoleRequestAppliedOrgWideInner
		return ret
	}
	return o.AppliedOrgWide
}

// GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOrgWideOk() ([]CreateOrganizationCameraRoleRequestAppliedOrgWideInner, bool) {
	if o == nil || IsNil(o.AppliedOrgWide) {
		return nil, false
	}
	return o.AppliedOrgWide, true
}

// HasAppliedOrgWide returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOrgWide() bool {
	if o != nil && !IsNil(o.AppliedOrgWide) {
		return true
	}

	return false
}

// SetAppliedOrgWide gets a reference to the given []CreateOrganizationCameraRoleRequestAppliedOrgWideInner and assigns it to the AppliedOrgWide field.
func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOrgWide(v []CreateOrganizationCameraRoleRequestAppliedOrgWideInner) {
	o.AppliedOrgWide = v
}

func (o UpdateOrganizationCameraRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationCameraRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AppliedOnDevices) {
		toSerialize["appliedOnDevices"] = o.AppliedOnDevices
	}
	if !IsNil(o.AppliedOnNetworks) {
		toSerialize["appliedOnNetworks"] = o.AppliedOnNetworks
	}
	if !IsNil(o.AppliedOrgWide) {
		toSerialize["appliedOrgWide"] = o.AppliedOrgWide
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationCameraRoleRequest struct {
	value *UpdateOrganizationCameraRoleRequest
	isSet bool
}

func (v NullableUpdateOrganizationCameraRoleRequest) Get() *UpdateOrganizationCameraRoleRequest {
	return v.value
}

func (v *NullableUpdateOrganizationCameraRoleRequest) Set(val *UpdateOrganizationCameraRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationCameraRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationCameraRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationCameraRoleRequest(val *UpdateOrganizationCameraRoleRequest) *NullableUpdateOrganizationCameraRoleRequest {
	return &NullableUpdateOrganizationCameraRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationCameraRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationCameraRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


