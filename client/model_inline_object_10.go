/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	// list of all reserved IP ranges for a single MG
	ReservedIpRanges []DevicesSerialCellularGatewayLanReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// list of all fixed IP assignments for a single MG
	FixedIpAssignments []DevicesSerialCellularGatewayLanFixedIpAssignments `json:"fixedIpAssignments,omitempty"`
}

// NewInlineObject10 instantiates a new InlineObject10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// NewInlineObject10WithDefaults instantiates a new InlineObject10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10WithDefaults() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject10) GetReservedIpRanges() []DevicesSerialCellularGatewayLanReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []DevicesSerialCellularGatewayLanReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetReservedIpRangesOk() ([]DevicesSerialCellularGatewayLanReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject10) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []DevicesSerialCellularGatewayLanReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject10) SetReservedIpRanges(v []DevicesSerialCellularGatewayLanReservedIpRanges) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject10) GetFixedIpAssignments() []DevicesSerialCellularGatewayLanFixedIpAssignments {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret []DevicesSerialCellularGatewayLanFixedIpAssignments
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetFixedIpAssignmentsOk() ([]DevicesSerialCellularGatewayLanFixedIpAssignments, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject10) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []DevicesSerialCellularGatewayLanFixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineObject10) SetFixedIpAssignments(v []DevicesSerialCellularGatewayLanFixedIpAssignments) {
	o.FixedIpAssignments = v
}

func (o InlineObject10) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject10 struct {
	value *InlineObject10
	isSet bool
}

func (v NullableInlineObject10) Get() *InlineObject10 {
	return v.value
}

func (v *NullableInlineObject10) Set(val *InlineObject10) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10(val *InlineObject10) *NullableInlineObject10 {
	return &NullableInlineObject10{value: val, isSet: true}
}

func (v NullableInlineObject10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


