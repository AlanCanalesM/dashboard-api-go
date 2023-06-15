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

// InlineResponse20071Alerts Email alert settings for DHCP servers
type InlineResponse20071Alerts struct {
	Email *InlineResponse20071AlertsEmail `json:"email,omitempty"`
}

// NewInlineResponse20071Alerts instantiates a new InlineResponse20071Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20071Alerts() *InlineResponse20071Alerts {
	this := InlineResponse20071Alerts{}
	return &this
}

// NewInlineResponse20071AlertsWithDefaults instantiates a new InlineResponse20071Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20071AlertsWithDefaults() *InlineResponse20071Alerts {
	this := InlineResponse20071Alerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse20071Alerts) GetEmail() InlineResponse20071AlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret InlineResponse20071AlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20071Alerts) GetEmailOk() (*InlineResponse20071AlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse20071Alerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given InlineResponse20071AlertsEmail and assigns it to the Email field.
func (o *InlineResponse20071Alerts) SetEmail(v InlineResponse20071AlertsEmail) {
	o.Email = &v
}

func (o InlineResponse20071Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20071Alerts struct {
	value *InlineResponse20071Alerts
	isSet bool
}

func (v NullableInlineResponse20071Alerts) Get() *InlineResponse20071Alerts {
	return v.value
}

func (v *NullableInlineResponse20071Alerts) Set(val *InlineResponse20071Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20071Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20071Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20071Alerts(val *InlineResponse20071Alerts) *NullableInlineResponse20071Alerts {
	return &NullableInlineResponse20071Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse20071Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20071Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


