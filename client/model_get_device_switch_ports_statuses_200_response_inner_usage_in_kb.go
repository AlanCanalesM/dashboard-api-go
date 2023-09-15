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

// checks if the GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb{}

// GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb A breakdown of how many kilobytes have passed through this port during the timespan.
type GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb struct {
	// The total amount of data sent and received (in kilobytes).
	Total *int32 `json:"total,omitempty"`
	// The amount of data sent (in kilobytes).
	Sent *int32 `json:"sent,omitempty"`
	// The amount of data received (in kilobytes).
	Recv *int32 `json:"recv,omitempty"`
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb() *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb{}
	return &this
}

// NewGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKbWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKbWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb {
	this := GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) SetTotal(v int32) {
	o.Total = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetSent() int32 {
	if o == nil || IsNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetSentOk() (*int32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) SetSent(v int32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetRecv() int32 {
	if o == nil || IsNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) GetRecvOk() (*int32, bool) {
	if o == nil || IsNil(o.Recv) {
		return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) HasRecv() bool {
	if o != nil && !IsNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) SetRecv(v int32) {
	o.Recv = &v
}

func (o GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb struct {
	value *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb
	isSet bool
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) Get() *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb {
	return v.value
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) Set(val *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb(val *GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) *NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb {
	return &NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


