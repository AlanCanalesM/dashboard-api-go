/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessLatencyHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessLatencyHistory200ResponseInner{}

// GetNetworkWirelessLatencyHistory200ResponseInner struct for GetNetworkWirelessLatencyHistory200ResponseInner
type GetNetworkWirelessLatencyHistory200ResponseInner struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Average latency in milliseconds
	AvgLatencyMs *int32 `json:"avgLatencyMs,omitempty"`
}

// NewGetNetworkWirelessLatencyHistory200ResponseInner instantiates a new GetNetworkWirelessLatencyHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessLatencyHistory200ResponseInner() *GetNetworkWirelessLatencyHistory200ResponseInner {
	this := GetNetworkWirelessLatencyHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessLatencyHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessLatencyHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessLatencyHistory200ResponseInnerWithDefaults() *GetNetworkWirelessLatencyHistory200ResponseInner {
	this := GetNetworkWirelessLatencyHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAvgLatencyMs returns the AvgLatencyMs field value if set, zero value otherwise.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetAvgLatencyMs() int32 {
	if o == nil || IsNil(o.AvgLatencyMs) {
		var ret int32
		return ret
	}
	return *o.AvgLatencyMs
}

// GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetAvgLatencyMsOk() (*int32, bool) {
	if o == nil || IsNil(o.AvgLatencyMs) {
		return nil, false
	}
	return o.AvgLatencyMs, true
}

// HasAvgLatencyMs returns a boolean if a field has been set.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasAvgLatencyMs() bool {
	if o != nil && !IsNil(o.AvgLatencyMs) {
		return true
	}

	return false
}

// SetAvgLatencyMs gets a reference to the given int32 and assigns it to the AvgLatencyMs field.
func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetAvgLatencyMs(v int32) {
	o.AvgLatencyMs = &v
}

func (o GetNetworkWirelessLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessLatencyHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.AvgLatencyMs) {
		toSerialize["avgLatencyMs"] = o.AvgLatencyMs
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessLatencyHistory200ResponseInner struct {
	value *GetNetworkWirelessLatencyHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessLatencyHistory200ResponseInner) Get() *GetNetworkWirelessLatencyHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessLatencyHistory200ResponseInner) Set(val *GetNetworkWirelessLatencyHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessLatencyHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessLatencyHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessLatencyHistory200ResponseInner(val *GetNetworkWirelessLatencyHistory200ResponseInner) *NullableGetNetworkWirelessLatencyHistory200ResponseInner {
	return &NullableGetNetworkWirelessLatencyHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessLatencyHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


