/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage{}

// GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage Usage info in megabytes
type GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage struct {
	// Total usage in megabytes
	Total *float32 `json:"total,omitempty"`
	// Average usage in megabytes
	Average *float32 `json:"average,omitempty"`
}

// NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage instantiates a new GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage{}
	return &this
}

// NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsageWithDefaults instantiates a new GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsageWithDefaults() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) GetAverage() float32 {
	if o == nil || IsNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) GetAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) SetAverage(v float32) {
	o.Average = &v
}

func (o GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage struct {
	value *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage
	isSet bool
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) Get() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) Set(val *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage(val *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage {
	return &NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


