/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject167 struct for InlineObject167
type InlineObject167 struct {
	// The name of the organization
	Name string `json:"name"`
	Management *OrganizationsManagement `json:"management,omitempty"`
}

// NewInlineObject167 instantiates a new InlineObject167 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject167(name string) *InlineObject167 {
	this := InlineObject167{}
	this.Name = name
	return &this
}

// NewInlineObject167WithDefaults instantiates a new InlineObject167 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject167WithDefaults() *InlineObject167 {
	this := InlineObject167{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject167) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject167) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject167) SetName(v string) {
	o.Name = v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *InlineObject167) GetManagement() OrganizationsManagement {
	if o == nil || isNil(o.Management) {
		var ret OrganizationsManagement
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject167) GetManagementOk() (*OrganizationsManagement, bool) {
	if o == nil || isNil(o.Management) {
    return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *InlineObject167) HasManagement() bool {
	if o != nil && !isNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given OrganizationsManagement and assigns it to the Management field.
func (o *InlineObject167) SetManagement(v OrganizationsManagement) {
	o.Management = &v
}

func (o InlineObject167) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject167 struct {
	value *InlineObject167
	isSet bool
}

func (v NullableInlineObject167) Get() *InlineObject167 {
	return v.value
}

func (v *NullableInlineObject167) Set(val *InlineObject167) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject167) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject167) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject167(val *InlineObject167) *NullableInlineObject167 {
	return &NullableInlineObject167{value: val, isSet: true}
}

func (v NullableInlineObject167) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject167) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


