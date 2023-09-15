/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the BlinkDeviceLeds202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlinkDeviceLeds202Response{}

// BlinkDeviceLeds202Response struct for BlinkDeviceLeds202Response
type BlinkDeviceLeds202Response struct {
	// The duration in seconds. Will be between 5 and 120. Default is 20 seconds
	Duration *int32 `json:"duration,omitempty"`
	// The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds
	Period *int32 `json:"period,omitempty"`
	// The duty cycle as the percent active. Will be between 10 and 90. Default is 50
	Duty *int32 `json:"duty,omitempty"`
}

// NewBlinkDeviceLeds202Response instantiates a new BlinkDeviceLeds202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlinkDeviceLeds202Response() *BlinkDeviceLeds202Response {
	this := BlinkDeviceLeds202Response{}
	return &this
}

// NewBlinkDeviceLeds202ResponseWithDefaults instantiates a new BlinkDeviceLeds202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlinkDeviceLeds202ResponseWithDefaults() *BlinkDeviceLeds202Response {
	this := BlinkDeviceLeds202Response{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BlinkDeviceLeds202Response) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlinkDeviceLeds202Response) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BlinkDeviceLeds202Response) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *BlinkDeviceLeds202Response) SetDuration(v int32) {
	o.Duration = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *BlinkDeviceLeds202Response) GetPeriod() int32 {
	if o == nil || IsNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlinkDeviceLeds202Response) GetPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *BlinkDeviceLeds202Response) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *BlinkDeviceLeds202Response) SetPeriod(v int32) {
	o.Period = &v
}

// GetDuty returns the Duty field value if set, zero value otherwise.
func (o *BlinkDeviceLeds202Response) GetDuty() int32 {
	if o == nil || IsNil(o.Duty) {
		var ret int32
		return ret
	}
	return *o.Duty
}

// GetDutyOk returns a tuple with the Duty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlinkDeviceLeds202Response) GetDutyOk() (*int32, bool) {
	if o == nil || IsNil(o.Duty) {
		return nil, false
	}
	return o.Duty, true
}

// HasDuty returns a boolean if a field has been set.
func (o *BlinkDeviceLeds202Response) HasDuty() bool {
	if o != nil && !IsNil(o.Duty) {
		return true
	}

	return false
}

// SetDuty gets a reference to the given int32 and assigns it to the Duty field.
func (o *BlinkDeviceLeds202Response) SetDuty(v int32) {
	o.Duty = &v
}

func (o BlinkDeviceLeds202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlinkDeviceLeds202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Duty) {
		toSerialize["duty"] = o.Duty
	}
	return toSerialize, nil
}

type NullableBlinkDeviceLeds202Response struct {
	value *BlinkDeviceLeds202Response
	isSet bool
}

func (v NullableBlinkDeviceLeds202Response) Get() *BlinkDeviceLeds202Response {
	return v.value
}

func (v *NullableBlinkDeviceLeds202Response) Set(val *BlinkDeviceLeds202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBlinkDeviceLeds202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBlinkDeviceLeds202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlinkDeviceLeds202Response(val *BlinkDeviceLeds202Response) *NullableBlinkDeviceLeds202Response {
	return &NullableBlinkDeviceLeds202Response{value: val, isSet: true}
}

func (v NullableBlinkDeviceLeds202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlinkDeviceLeds202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


