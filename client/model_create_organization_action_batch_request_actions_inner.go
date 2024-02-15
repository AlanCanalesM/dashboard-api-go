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

// checks if the CreateOrganizationActionBatchRequestActionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationActionBatchRequestActionsInner{}

// CreateOrganizationActionBatchRequestActionsInner struct for CreateOrganizationActionBatchRequestActionsInner
type CreateOrganizationActionBatchRequestActionsInner struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used
	Operation string `json:"operation"`
	// The body of the action
	Body map[string]interface{} `json:"body,omitempty"`
}

// NewCreateOrganizationActionBatchRequestActionsInner instantiates a new CreateOrganizationActionBatchRequestActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatchRequestActionsInner(resource string, operation string) *CreateOrganizationActionBatchRequestActionsInner {
	this := CreateOrganizationActionBatchRequestActionsInner{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewCreateOrganizationActionBatchRequestActionsInnerWithDefaults instantiates a new CreateOrganizationActionBatchRequestActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatchRequestActionsInnerWithDefaults() *CreateOrganizationActionBatchRequestActionsInner {
	this := CreateOrganizationActionBatchRequestActionsInner{}
	return &this
}

// GetResource returns the Resource field value
func (o *CreateOrganizationActionBatchRequestActionsInner) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequestActionsInner) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CreateOrganizationActionBatchRequestActionsInner) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *CreateOrganizationActionBatchRequestActionsInner) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequestActionsInner) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *CreateOrganizationActionBatchRequestActionsInner) SetOperation(v string) {
	o.Operation = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatchRequestActionsInner) GetBody() map[string]interface{} {
	if o == nil || IsNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatchRequestActionsInner) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Body) {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatchRequestActionsInner) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *CreateOrganizationActionBatchRequestActionsInner) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o CreateOrganizationActionBatchRequestActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationActionBatchRequestActionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["operation"] = o.Operation
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableCreateOrganizationActionBatchRequestActionsInner struct {
	value *CreateOrganizationActionBatchRequestActionsInner
	isSet bool
}

func (v NullableCreateOrganizationActionBatchRequestActionsInner) Get() *CreateOrganizationActionBatchRequestActionsInner {
	return v.value
}

func (v *NullableCreateOrganizationActionBatchRequestActionsInner) Set(val *CreateOrganizationActionBatchRequestActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatchRequestActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatchRequestActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatchRequestActionsInner(val *CreateOrganizationActionBatchRequestActionsInner) *NullableCreateOrganizationActionBatchRequestActionsInner {
	return &NullableCreateOrganizationActionBatchRequestActionsInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatchRequestActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatchRequestActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


