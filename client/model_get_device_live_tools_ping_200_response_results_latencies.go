/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetDeviceLiveToolsPing200ResponseResultsLatencies Packet latency stats
type GetDeviceLiveToolsPing200ResponseResultsLatencies struct {
	// Minimum latency
	Minimum *float32 `json:"minimum,omitempty"`
	// Average latency
	Average *float32 `json:"average,omitempty"`
	// Maximum latency
	Maximum *float32 `json:"maximum,omitempty"`
}

// NewGetDeviceLiveToolsPing200ResponseResultsLatencies instantiates a new GetDeviceLiveToolsPing200ResponseResultsLatencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceLiveToolsPing200ResponseResultsLatencies() *GetDeviceLiveToolsPing200ResponseResultsLatencies {
	this := GetDeviceLiveToolsPing200ResponseResultsLatencies{}
	return &this
}

// NewGetDeviceLiveToolsPing200ResponseResultsLatenciesWithDefaults instantiates a new GetDeviceLiveToolsPing200ResponseResultsLatencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceLiveToolsPing200ResponseResultsLatenciesWithDefaults() *GetDeviceLiveToolsPing200ResponseResultsLatencies {
	this := GetDeviceLiveToolsPing200ResponseResultsLatencies{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetMinimum() float32 {
	if o == nil || isNil(o.Minimum) {
		var ret float32
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetMinimumOk() (*float32, bool) {
	if o == nil || isNil(o.Minimum) {
    return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) HasMinimum() bool {
	if o != nil && !isNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given float32 and assigns it to the Minimum field.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) SetMinimum(v float32) {
	o.Minimum = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetAverage() float32 {
	if o == nil || isNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetAverageOk() (*float32, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) SetAverage(v float32) {
	o.Average = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetMaximum() float32 {
	if o == nil || isNil(o.Maximum) {
		var ret float32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) GetMaximumOk() (*float32, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given float32 and assigns it to the Maximum field.
func (o *GetDeviceLiveToolsPing200ResponseResultsLatencies) SetMaximum(v float32) {
	o.Maximum = &v
}

func (o GetDeviceLiveToolsPing200ResponseResultsLatencies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeviceLiveToolsPing200ResponseResultsLatencies struct {
	value *GetDeviceLiveToolsPing200ResponseResultsLatencies
	isSet bool
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) Get() *GetDeviceLiveToolsPing200ResponseResultsLatencies {
	return v.value
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) Set(val *GetDeviceLiveToolsPing200ResponseResultsLatencies) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceLiveToolsPing200ResponseResultsLatencies(val *GetDeviceLiveToolsPing200ResponseResultsLatencies) *NullableGetDeviceLiveToolsPing200ResponseResultsLatencies {
	return &NullableGetDeviceLiveToolsPing200ResponseResultsLatencies{value: val, isSet: true}
}

func (v NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceLiveToolsPing200ResponseResultsLatencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


