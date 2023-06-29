/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationActionBatch201ResponseActionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationActionBatch201ResponseActionsInner{}

// CreateOrganizationActionBatch201ResponseActionsInner struct for CreateOrganizationActionBatch201ResponseActionsInner
type CreateOrganizationActionBatch201ResponseActionsInner struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used by this action
	Operation string `json:"operation"`
}

// NewCreateOrganizationActionBatch201ResponseActionsInner instantiates a new CreateOrganizationActionBatch201ResponseActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatch201ResponseActionsInner(resource string, operation string) *CreateOrganizationActionBatch201ResponseActionsInner {
	this := CreateOrganizationActionBatch201ResponseActionsInner{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewCreateOrganizationActionBatch201ResponseActionsInnerWithDefaults instantiates a new CreateOrganizationActionBatch201ResponseActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatch201ResponseActionsInnerWithDefaults() *CreateOrganizationActionBatch201ResponseActionsInner {
	this := CreateOrganizationActionBatch201ResponseActionsInner{}
	return &this
}

// GetResource returns the Resource field value
func (o *CreateOrganizationActionBatch201ResponseActionsInner) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseActionsInner) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CreateOrganizationActionBatch201ResponseActionsInner) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *CreateOrganizationActionBatch201ResponseActionsInner) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseActionsInner) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *CreateOrganizationActionBatch201ResponseActionsInner) SetOperation(v string) {
	o.Operation = v
}

func (o CreateOrganizationActionBatch201ResponseActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationActionBatch201ResponseActionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["operation"] = o.Operation
	return toSerialize, nil
}

type NullableCreateOrganizationActionBatch201ResponseActionsInner struct {
	value *CreateOrganizationActionBatch201ResponseActionsInner
	isSet bool
}

func (v NullableCreateOrganizationActionBatch201ResponseActionsInner) Get() *CreateOrganizationActionBatch201ResponseActionsInner {
	return v.value
}

func (v *NullableCreateOrganizationActionBatch201ResponseActionsInner) Set(val *CreateOrganizationActionBatch201ResponseActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatch201ResponseActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatch201ResponseActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatch201ResponseActionsInner(val *CreateOrganizationActionBatch201ResponseActionsInner) *NullableCreateOrganizationActionBatch201ResponseActionsInner {
	return &NullableCreateOrganizationActionBatch201ResponseActionsInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatch201ResponseActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatch201ResponseActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


