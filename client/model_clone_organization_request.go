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

// checks if the CloneOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneOrganizationRequest{}

// CloneOrganizationRequest struct for CloneOrganizationRequest
type CloneOrganizationRequest struct {
	// The name of the new organization
	Name string `json:"name"`
}

// NewCloneOrganizationRequest instantiates a new CloneOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneOrganizationRequest(name string) *CloneOrganizationRequest {
	this := CloneOrganizationRequest{}
	this.Name = name
	return &this
}

// NewCloneOrganizationRequestWithDefaults instantiates a new CloneOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneOrganizationRequestWithDefaults() *CloneOrganizationRequest {
	this := CloneOrganizationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneOrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneOrganizationRequest) SetName(v string) {
	o.Name = v
}

func (o CloneOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCloneOrganizationRequest struct {
	value *CloneOrganizationRequest
	isSet bool
}

func (v NullableCloneOrganizationRequest) Get() *CloneOrganizationRequest {
	return v.value
}

func (v *NullableCloneOrganizationRequest) Set(val *CloneOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneOrganizationRequest(val *CloneOrganizationRequest) *NullableCloneOrganizationRequest {
	return &NullableCloneOrganizationRequest{value: val, isSet: true}
}

func (v NullableCloneOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


