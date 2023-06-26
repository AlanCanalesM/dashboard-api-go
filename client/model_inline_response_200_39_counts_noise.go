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

// InlineResponse20039CountsNoise Object containing the number of sensors that are currently alerting due to noise readings
type InlineResponse20039CountsNoise struct {
	// Number of sensors that are currently alerting due to ambient noise readings
	Ambient *int32 `json:"ambient,omitempty"`
}

// NewInlineResponse20039CountsNoise instantiates a new InlineResponse20039CountsNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20039CountsNoise() *InlineResponse20039CountsNoise {
	this := InlineResponse20039CountsNoise{}
	return &this
}

// NewInlineResponse20039CountsNoiseWithDefaults instantiates a new InlineResponse20039CountsNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20039CountsNoiseWithDefaults() *InlineResponse20039CountsNoise {
	this := InlineResponse20039CountsNoise{}
	return &this
}

// GetAmbient returns the Ambient field value if set, zero value otherwise.
func (o *InlineResponse20039CountsNoise) GetAmbient() int32 {
	if o == nil || isNil(o.Ambient) {
		var ret int32
		return ret
	}
	return *o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039CountsNoise) GetAmbientOk() (*int32, bool) {
	if o == nil || isNil(o.Ambient) {
    return nil, false
	}
	return o.Ambient, true
}

// HasAmbient returns a boolean if a field has been set.
func (o *InlineResponse20039CountsNoise) HasAmbient() bool {
	if o != nil && !isNil(o.Ambient) {
		return true
	}

	return false
}

// SetAmbient gets a reference to the given int32 and assigns it to the Ambient field.
func (o *InlineResponse20039CountsNoise) SetAmbient(v int32) {
	o.Ambient = &v
}

func (o InlineResponse20039CountsNoise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ambient) {
		toSerialize["ambient"] = o.Ambient
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20039CountsNoise struct {
	value *InlineResponse20039CountsNoise
	isSet bool
}

func (v NullableInlineResponse20039CountsNoise) Get() *InlineResponse20039CountsNoise {
	return v.value
}

func (v *NullableInlineResponse20039CountsNoise) Set(val *InlineResponse20039CountsNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20039CountsNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20039CountsNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20039CountsNoise(val *InlineResponse20039CountsNoise) *NullableInlineResponse20039CountsNoise {
	return &NullableInlineResponse20039CountsNoise{value: val, isSet: true}
}

func (v NullableInlineResponse20039CountsNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20039CountsNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


