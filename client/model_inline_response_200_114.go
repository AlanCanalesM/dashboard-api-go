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

// InlineResponse200114 struct for InlineResponse200114
type InlineResponse200114 struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []InlineResponse200113 `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []InlineResponse200113 `json:"movedLicenses,omitempty"`
}

// NewInlineResponse200114 instantiates a new InlineResponse200114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200114() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200114WithDefaults() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *InlineResponse200114) GetRemainderLicenses() []InlineResponse200113 {
	if o == nil || isNil(o.RemainderLicenses) {
		var ret []InlineResponse200113
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetRemainderLicensesOk() ([]InlineResponse200113, bool) {
	if o == nil || isNil(o.RemainderLicenses) {
    return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *InlineResponse200114) HasRemainderLicenses() bool {
	if o != nil && !isNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []InlineResponse200113 and assigns it to the RemainderLicenses field.
func (o *InlineResponse200114) SetRemainderLicenses(v []InlineResponse200113) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *InlineResponse200114) GetMovedLicenses() []InlineResponse200113 {
	if o == nil || isNil(o.MovedLicenses) {
		var ret []InlineResponse200113
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetMovedLicensesOk() ([]InlineResponse200113, bool) {
	if o == nil || isNil(o.MovedLicenses) {
    return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *InlineResponse200114) HasMovedLicenses() bool {
	if o != nil && !isNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []InlineResponse200113 and assigns it to the MovedLicenses field.
func (o *InlineResponse200114) SetMovedLicenses(v []InlineResponse200113) {
	o.MovedLicenses = v
}

func (o InlineResponse200114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !isNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200114 struct {
	value *InlineResponse200114
	isSet bool
}

func (v NullableInlineResponse200114) Get() *InlineResponse200114 {
	return v.value
}

func (v *NullableInlineResponse200114) Set(val *InlineResponse200114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200114(val *InlineResponse200114) *NullableInlineResponse200114 {
	return &NullableInlineResponse200114{value: val, isSet: true}
}

func (v NullableInlineResponse200114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


