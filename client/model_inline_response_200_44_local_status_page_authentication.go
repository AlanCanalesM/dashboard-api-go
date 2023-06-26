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

// InlineResponse20044LocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse20044LocalStatusPageAuthentication struct {
	// Enables / disables the authentication on Local Status page(s).
	Enabled *bool `json:"enabled,omitempty"`
	// The username used for Local Status Page(s).
	Username *string `json:"username,omitempty"`
}

// NewInlineResponse20044LocalStatusPageAuthentication instantiates a new InlineResponse20044LocalStatusPageAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044LocalStatusPageAuthentication() *InlineResponse20044LocalStatusPageAuthentication {
	this := InlineResponse20044LocalStatusPageAuthentication{}
	return &this
}

// NewInlineResponse20044LocalStatusPageAuthenticationWithDefaults instantiates a new InlineResponse20044LocalStatusPageAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044LocalStatusPageAuthenticationWithDefaults() *InlineResponse20044LocalStatusPageAuthentication {
	this := InlineResponse20044LocalStatusPageAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20044LocalStatusPageAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044LocalStatusPageAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20044LocalStatusPageAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20044LocalStatusPageAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse20044LocalStatusPageAuthentication) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044LocalStatusPageAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse20044LocalStatusPageAuthentication) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse20044LocalStatusPageAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o InlineResponse20044LocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044LocalStatusPageAuthentication struct {
	value *InlineResponse20044LocalStatusPageAuthentication
	isSet bool
}

func (v NullableInlineResponse20044LocalStatusPageAuthentication) Get() *InlineResponse20044LocalStatusPageAuthentication {
	return v.value
}

func (v *NullableInlineResponse20044LocalStatusPageAuthentication) Set(val *InlineResponse20044LocalStatusPageAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044LocalStatusPageAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044LocalStatusPageAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044LocalStatusPageAuthentication(val *InlineResponse20044LocalStatusPageAuthentication) *NullableInlineResponse20044LocalStatusPageAuthentication {
	return &NullableInlineResponse20044LocalStatusPageAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse20044LocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044LocalStatusPageAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


