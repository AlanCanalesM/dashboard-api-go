/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage Energy usage of the switch
type GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage struct {
	// Total energy usage of the switch
	Total *float32 `json:"total,omitempty"`
}

// NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage instantiates a new GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage{}
	return &this
}

// NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsageWithDefaults instantiates a new GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsageWithDefaults() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) SetTotal(v float32) {
	o.Total = &v
}

func (o GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage struct {
	value *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage
	isSet bool
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) Get() *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) Set(val *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage(val *GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage {
	return &NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


