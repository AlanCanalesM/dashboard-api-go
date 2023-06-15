/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20071AlertsEmail Alert settings for DHCP servers
type InlineResponse20071AlertsEmail struct {
	// When enabled, send an email if a new DHCP server is seen. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20071AlertsEmail instantiates a new InlineResponse20071AlertsEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20071AlertsEmail() *InlineResponse20071AlertsEmail {
	this := InlineResponse20071AlertsEmail{}
	return &this
}

// NewInlineResponse20071AlertsEmailWithDefaults instantiates a new InlineResponse20071AlertsEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20071AlertsEmailWithDefaults() *InlineResponse20071AlertsEmail {
	this := InlineResponse20071AlertsEmail{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20071AlertsEmail) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20071AlertsEmail) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20071AlertsEmail) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20071AlertsEmail) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20071AlertsEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20071AlertsEmail struct {
	value *InlineResponse20071AlertsEmail
	isSet bool
}

func (v NullableInlineResponse20071AlertsEmail) Get() *InlineResponse20071AlertsEmail {
	return v.value
}

func (v *NullableInlineResponse20071AlertsEmail) Set(val *InlineResponse20071AlertsEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20071AlertsEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20071AlertsEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20071AlertsEmail(val *InlineResponse20071AlertsEmail) *NullableInlineResponse20071AlertsEmail {
	return &NullableInlineResponse20071AlertsEmail{value: val, isSet: true}
}

func (v NullableInlineResponse20071AlertsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20071AlertsEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


