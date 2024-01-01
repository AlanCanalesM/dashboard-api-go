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

// checks if the GetOrganizationClientsOverview200ResponseUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationClientsOverview200ResponseUsage{}

// GetOrganizationClientsOverview200ResponseUsage Usage information of all clients across organization
type GetOrganizationClientsOverview200ResponseUsage struct {
	Overall *GetOrganizationClientsOverview200ResponseUsageOverall `json:"overall,omitempty"`
	// Average data usage (in kb) of each client in organization
	Average *float32 `json:"average,omitempty"`
}

// NewGetOrganizationClientsOverview200ResponseUsage instantiates a new GetOrganizationClientsOverview200ResponseUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationClientsOverview200ResponseUsage() *GetOrganizationClientsOverview200ResponseUsage {
	this := GetOrganizationClientsOverview200ResponseUsage{}
	return &this
}

// NewGetOrganizationClientsOverview200ResponseUsageWithDefaults instantiates a new GetOrganizationClientsOverview200ResponseUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationClientsOverview200ResponseUsageWithDefaults() *GetOrganizationClientsOverview200ResponseUsage {
	this := GetOrganizationClientsOverview200ResponseUsage{}
	return &this
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200ResponseUsage) GetOverall() GetOrganizationClientsOverview200ResponseUsageOverall {
	if o == nil || IsNil(o.Overall) {
		var ret GetOrganizationClientsOverview200ResponseUsageOverall
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200ResponseUsage) GetOverallOk() (*GetOrganizationClientsOverview200ResponseUsageOverall, bool) {
	if o == nil || IsNil(o.Overall) {
		return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200ResponseUsage) HasOverall() bool {
	if o != nil && !IsNil(o.Overall) {
		return true
	}

	return false
}

// SetOverall gets a reference to the given GetOrganizationClientsOverview200ResponseUsageOverall and assigns it to the Overall field.
func (o *GetOrganizationClientsOverview200ResponseUsage) SetOverall(v GetOrganizationClientsOverview200ResponseUsageOverall) {
	o.Overall = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200ResponseUsage) GetAverage() float32 {
	if o == nil || IsNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200ResponseUsage) GetAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200ResponseUsage) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *GetOrganizationClientsOverview200ResponseUsage) SetAverage(v float32) {
	o.Average = &v
}

func (o GetOrganizationClientsOverview200ResponseUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationClientsOverview200ResponseUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Overall) {
		toSerialize["overall"] = o.Overall
	}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return toSerialize, nil
}

type NullableGetOrganizationClientsOverview200ResponseUsage struct {
	value *GetOrganizationClientsOverview200ResponseUsage
	isSet bool
}

func (v NullableGetOrganizationClientsOverview200ResponseUsage) Get() *GetOrganizationClientsOverview200ResponseUsage {
	return v.value
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsage) Set(val *GetOrganizationClientsOverview200ResponseUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationClientsOverview200ResponseUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationClientsOverview200ResponseUsage(val *GetOrganizationClientsOverview200ResponseUsage) *NullableGetOrganizationClientsOverview200ResponseUsage {
	return &NullableGetOrganizationClientsOverview200ResponseUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationClientsOverview200ResponseUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


