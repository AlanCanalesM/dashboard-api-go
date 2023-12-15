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

// checks if the GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner{}

// GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner struct for GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner
type GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner struct {
	// Response status code of the API response
	Code *int32 `json:"code,omitempty"`
	// Number of records that match the status code
	Count *int32 `json:"count,omitempty"`
}

// NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner {
	this := GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner{}
	return &this
}

// NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInnerWithDefaults instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInnerWithDefaults() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner {
	this := GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) SetCode(v int32) {
	o.Code = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) SetCount(v int32) {
	o.Count = &v
}

func (o GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner struct {
	value *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner
	isSet bool
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) Get() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner {
	return v.value
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) Set(val *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner(val *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner {
	return &NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


