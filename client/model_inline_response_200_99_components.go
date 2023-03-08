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

// InlineResponse20099Components Components
type InlineResponse20099Components struct {
	// Power Supplies
	PowerSupplies []string `json:"powerSupplies,omitempty"`
}

// NewInlineResponse20099Components instantiates a new InlineResponse20099Components object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20099Components() *InlineResponse20099Components {
	this := InlineResponse20099Components{}
	return &this
}

// NewInlineResponse20099ComponentsWithDefaults instantiates a new InlineResponse20099Components object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20099ComponentsWithDefaults() *InlineResponse20099Components {
	this := InlineResponse20099Components{}
	return &this
}

// GetPowerSupplies returns the PowerSupplies field value if set, zero value otherwise.
func (o *InlineResponse20099Components) GetPowerSupplies() []string {
	if o == nil || isNil(o.PowerSupplies) {
		var ret []string
		return ret
	}
	return o.PowerSupplies
}

// GetPowerSuppliesOk returns a tuple with the PowerSupplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20099Components) GetPowerSuppliesOk() ([]string, bool) {
	if o == nil || isNil(o.PowerSupplies) {
    return nil, false
	}
	return o.PowerSupplies, true
}

// HasPowerSupplies returns a boolean if a field has been set.
func (o *InlineResponse20099Components) HasPowerSupplies() bool {
	if o != nil && !isNil(o.PowerSupplies) {
		return true
	}

	return false
}

// SetPowerSupplies gets a reference to the given []string and assigns it to the PowerSupplies field.
func (o *InlineResponse20099Components) SetPowerSupplies(v []string) {
	o.PowerSupplies = v
}

func (o InlineResponse20099Components) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PowerSupplies) {
		toSerialize["powerSupplies"] = o.PowerSupplies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20099Components struct {
	value *InlineResponse20099Components
	isSet bool
}

func (v NullableInlineResponse20099Components) Get() *InlineResponse20099Components {
	return v.value
}

func (v *NullableInlineResponse20099Components) Set(val *InlineResponse20099Components) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20099Components) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20099Components) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20099Components(val *InlineResponse20099Components) *NullableInlineResponse20099Components {
	return &NullableInlineResponse20099Components{value: val, isSet: true}
}

func (v NullableInlineResponse20099Components) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20099Components) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


