/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork{}

// GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork 
type GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork struct {
	// Name of network
	Name *string `json:"name,omitempty"`
	// ID of network
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork instantiates a new GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork() *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork {
	this := GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetworkWithDefaults() *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork {
	this := GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork struct {
	value *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) Get() *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) Set(val *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork(val *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork {
	return &NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


