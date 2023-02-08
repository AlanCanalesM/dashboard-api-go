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

// GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts Counts of the clients
type GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts struct {
	// Total counts of the clients
	Total *int32 `json:"total,omitempty"`
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts{}
	return &this
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCountsWithDefaults instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCountsWithDefaults() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) SetTotal(v int32) {
	o.Total = &v
}

func (o GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts struct {
	value *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts
	isSet bool
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) Get() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) Set(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts {
	return &NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


