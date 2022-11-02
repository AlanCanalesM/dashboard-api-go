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

// GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients Clients
type GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients struct {
	Counts *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts `json:"counts,omitempty"`
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients{}
	return &this
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClientsWithDefaults instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClientsWithDefaults() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) GetCounts() GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	if o == nil || isNil(o.Counts) {
		var ret GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) GetCountsOk() (*GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts and assigns it to the Counts field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) SetCounts(v GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) {
	o.Counts = &v
}

func (o GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients struct {
	value *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients
	isSet bool
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) Get() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) Set(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients {
	return &NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


