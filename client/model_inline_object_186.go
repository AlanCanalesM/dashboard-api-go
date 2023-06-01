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

// InlineObject186 struct for InlineObject186
type InlineObject186 struct {
	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies 
	BrandingPolicyIds []string `json:"brandingPolicyIds,omitempty"`
}

// NewInlineObject186 instantiates a new InlineObject186 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject186() *InlineObject186 {
	this := InlineObject186{}
	return &this
}

// NewInlineObject186WithDefaults instantiates a new InlineObject186 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject186WithDefaults() *InlineObject186 {
	this := InlineObject186{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value if set, zero value otherwise.
func (o *InlineObject186) GetBrandingPolicyIds() []string {
	if o == nil || isNil(o.BrandingPolicyIds) {
		var ret []string
		return ret
	}
	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetBrandingPolicyIdsOk() ([]string, bool) {
	if o == nil || isNil(o.BrandingPolicyIds) {
    return nil, false
	}
	return o.BrandingPolicyIds, true
}

// HasBrandingPolicyIds returns a boolean if a field has been set.
func (o *InlineObject186) HasBrandingPolicyIds() bool {
	if o != nil && !isNil(o.BrandingPolicyIds) {
		return true
	}

	return false
}

// SetBrandingPolicyIds gets a reference to the given []string and assigns it to the BrandingPolicyIds field.
func (o *InlineObject186) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o InlineObject186) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BrandingPolicyIds) {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject186 struct {
	value *InlineObject186
	isSet bool
}

func (v NullableInlineObject186) Get() *InlineObject186 {
	return v.value
}

func (v *NullableInlineObject186) Set(val *InlineObject186) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject186) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject186) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject186(val *InlineObject186) *NullableInlineObject186 {
	return &NullableInlineObject186{value: val, isSet: true}
}

func (v NullableInlineObject186) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject186) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


