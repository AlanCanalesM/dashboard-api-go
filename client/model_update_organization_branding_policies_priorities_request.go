/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateOrganizationBrandingPoliciesPrioritiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationBrandingPoliciesPrioritiesRequest{}

// UpdateOrganizationBrandingPoliciesPrioritiesRequest struct for UpdateOrganizationBrandingPoliciesPrioritiesRequest
type UpdateOrganizationBrandingPoliciesPrioritiesRequest struct {
	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies 
	BrandingPolicyIds []string `json:"brandingPolicyIds,omitempty"`
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesRequest instantiates a new UpdateOrganizationBrandingPoliciesPrioritiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationBrandingPoliciesPrioritiesRequest() *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	this := UpdateOrganizationBrandingPoliciesPrioritiesRequest{}
	return &this
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesRequestWithDefaults instantiates a new UpdateOrganizationBrandingPoliciesPrioritiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationBrandingPoliciesPrioritiesRequestWithDefaults() *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	this := UpdateOrganizationBrandingPoliciesPrioritiesRequest{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) GetBrandingPolicyIds() []string {
	if o == nil || IsNil(o.BrandingPolicyIds) {
		var ret []string
		return ret
	}
	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) GetBrandingPolicyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BrandingPolicyIds) {
		return nil, false
	}
	return o.BrandingPolicyIds, true
}

// HasBrandingPolicyIds returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) HasBrandingPolicyIds() bool {
	if o != nil && !IsNil(o.BrandingPolicyIds) {
		return true
	}

	return false
}

// SetBrandingPolicyIds gets a reference to the given []string and assigns it to the BrandingPolicyIds field.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o UpdateOrganizationBrandingPoliciesPrioritiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationBrandingPoliciesPrioritiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandingPolicyIds) {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest struct {
	value *UpdateOrganizationBrandingPoliciesPrioritiesRequest
	isSet bool
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Get() *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	return v.value
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Set(val *UpdateOrganizationBrandingPoliciesPrioritiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationBrandingPoliciesPrioritiesRequest(val *UpdateOrganizationBrandingPoliciesPrioritiesRequest) *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest {
	return &NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


