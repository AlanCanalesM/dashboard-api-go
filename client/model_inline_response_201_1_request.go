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

// InlineResponse2011Request Ping request parameters
type InlineResponse2011Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// IP address or FQDN to ping
	Target *string `json:"target,omitempty"`
	// Number of pings to send
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse2011Request instantiates a new InlineResponse2011Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2011Request() *InlineResponse2011Request {
	this := InlineResponse2011Request{}
	return &this
}

// NewInlineResponse2011RequestWithDefaults instantiates a new InlineResponse2011Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2011RequestWithDefaults() *InlineResponse2011Request {
	this := InlineResponse2011Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2011Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2011Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2011Request) SetSerial(v string) {
	o.Serial = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *InlineResponse2011Request) GetTarget() string {
	if o == nil || isNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Request) GetTargetOk() (*string, bool) {
	if o == nil || isNil(o.Target) {
    return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *InlineResponse2011Request) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *InlineResponse2011Request) SetTarget(v string) {
	o.Target = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse2011Request) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Request) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse2011Request) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse2011Request) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse2011Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2011Request struct {
	value *InlineResponse2011Request
	isSet bool
}

func (v NullableInlineResponse2011Request) Get() *InlineResponse2011Request {
	return v.value
}

func (v *NullableInlineResponse2011Request) Set(val *InlineResponse2011Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2011Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2011Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2011Request(val *InlineResponse2011Request) *NullableInlineResponse2011Request {
	return &NullableInlineResponse2011Request{value: val, isSet: true}
}

func (v NullableInlineResponse2011Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2011Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


