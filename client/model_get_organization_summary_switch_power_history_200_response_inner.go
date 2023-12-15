/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationSummarySwitchPowerHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummarySwitchPowerHistory200ResponseInner{}

// GetOrganizationSummarySwitchPowerHistory200ResponseInner struct for GetOrganizationSummarySwitchPowerHistory200ResponseInner
type GetOrganizationSummarySwitchPowerHistory200ResponseInner struct {
	// Timestamp of the start of the interval.
	Ts *time.Time `json:"ts,omitempty"`
	// The PoE power draw in watts for all switch ports in the organization for the given interval.
	Draw *float32 `json:"draw,omitempty"`
}

// NewGetOrganizationSummarySwitchPowerHistory200ResponseInner instantiates a new GetOrganizationSummarySwitchPowerHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummarySwitchPowerHistory200ResponseInner() *GetOrganizationSummarySwitchPowerHistory200ResponseInner {
	this := GetOrganizationSummarySwitchPowerHistory200ResponseInner{}
	return &this
}

// NewGetOrganizationSummarySwitchPowerHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationSummarySwitchPowerHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummarySwitchPowerHistory200ResponseInnerWithDefaults() *GetOrganizationSummarySwitchPowerHistory200ResponseInner {
	this := GetOrganizationSummarySwitchPowerHistory200ResponseInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) GetDraw() float32 {
	if o == nil || IsNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) GetDrawOk() (*float32, bool) {
	if o == nil || IsNil(o.Draw) {
		return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) HasDraw() bool {
	if o != nil && !IsNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *GetOrganizationSummarySwitchPowerHistory200ResponseInner) SetDraw(v float32) {
	o.Draw = &v
}

func (o GetOrganizationSummarySwitchPowerHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummarySwitchPowerHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner struct {
	value *GetOrganizationSummarySwitchPowerHistory200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) Get() *GetOrganizationSummarySwitchPowerHistory200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) Set(val *GetOrganizationSummarySwitchPowerHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummarySwitchPowerHistory200ResponseInner(val *GetOrganizationSummarySwitchPowerHistory200ResponseInner) *NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner {
	return &NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummarySwitchPowerHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


