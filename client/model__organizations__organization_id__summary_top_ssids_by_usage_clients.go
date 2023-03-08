/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopSsidsByUsageClients Clients info of the SSID
type OrganizationsOrganizationIdSummaryTopSsidsByUsageClients struct {
	Counts *OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClients instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClients() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageClients{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopSsidsByUsageClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopSsidsByUsageClientsWithDefaults() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients {
	this := OrganizationsOrganizationIdSummaryTopSsidsByUsageClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) GetCounts() OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) GetCountsOk() (*OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) SetCounts(v OrganizationsOrganizationIdSummaryTopSsidsByUsageClientsCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients struct {
	value *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) Get() *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) Set(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients(val *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients {
	return &NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSsidsByUsageClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


