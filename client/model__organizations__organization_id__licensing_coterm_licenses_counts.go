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

// OrganizationsOrganizationIdLicensingCotermLicensesCounts struct for OrganizationsOrganizationIdLicensingCotermLicensesCounts
type OrganizationsOrganizationIdLicensingCotermLicensesCounts struct {
	// The license model type
	Model *string `json:"model,omitempty"`
	// The number of counts the license contains of this model
	Count *int32 `json:"count,omitempty"`
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesCounts instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensingCotermLicensesCounts() *OrganizationsOrganizationIdLicensingCotermLicensesCounts {
	this := OrganizationsOrganizationIdLicensingCotermLicensesCounts{}
	return &this
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesCountsWithDefaults instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensingCotermLicensesCountsWithDefaults() *OrganizationsOrganizationIdLicensingCotermLicensesCounts {
	this := OrganizationsOrganizationIdLicensingCotermLicensesCounts{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) SetModel(v string) {
	o.Model = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesCounts) SetCount(v int32) {
	o.Count = &v
}

func (o OrganizationsOrganizationIdLicensingCotermLicensesCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts struct {
	value *OrganizationsOrganizationIdLicensingCotermLicensesCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) Get() *OrganizationsOrganizationIdLicensingCotermLicensesCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) Set(val *OrganizationsOrganizationIdLicensingCotermLicensesCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensingCotermLicensesCounts(val *OrganizationsOrganizationIdLicensingCotermLicensesCounts) *NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts {
	return &NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


