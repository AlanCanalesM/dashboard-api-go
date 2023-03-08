/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2015Actions struct for InlineResponse2015Actions
type InlineResponse2015Actions struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used by this action
	Operation string `json:"operation"`
}

// NewInlineResponse2015Actions instantiates a new InlineResponse2015Actions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015Actions(resource string, operation string) *InlineResponse2015Actions {
	this := InlineResponse2015Actions{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewInlineResponse2015ActionsWithDefaults instantiates a new InlineResponse2015Actions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015ActionsWithDefaults() *InlineResponse2015Actions {
	this := InlineResponse2015Actions{}
	return &this
}

// GetResource returns the Resource field value
func (o *InlineResponse2015Actions) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Actions) GetResourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *InlineResponse2015Actions) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *InlineResponse2015Actions) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Actions) GetOperationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *InlineResponse2015Actions) SetOperation(v string) {
	o.Operation = v
}

func (o InlineResponse2015Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015Actions struct {
	value *InlineResponse2015Actions
	isSet bool
}

func (v NullableInlineResponse2015Actions) Get() *InlineResponse2015Actions {
	return v.value
}

func (v *NullableInlineResponse2015Actions) Set(val *InlineResponse2015Actions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015Actions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015Actions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015Actions(val *InlineResponse2015Actions) *NullableInlineResponse2015Actions {
	return &NullableInlineResponse2015Actions{value: val, isSet: true}
}

func (v NullableInlineResponse2015Actions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015Actions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


