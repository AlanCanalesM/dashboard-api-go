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

// checks if the GetOrganizationDevicesAvailabilities200ResponseInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesAvailabilities200ResponseInnerNetwork{}

// GetOrganizationDevicesAvailabilities200ResponseInnerNetwork Network info.
type GetOrganizationDevicesAvailabilities200ResponseInnerNetwork struct {
	// ID for the network containing the device.
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationDevicesAvailabilities200ResponseInnerNetwork instantiates a new GetOrganizationDevicesAvailabilities200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesAvailabilities200ResponseInnerNetwork() *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	this := GetOrganizationDevicesAvailabilities200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationDevicesAvailabilities200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationDevicesAvailabilities200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesAvailabilities200ResponseInnerNetworkWithDefaults() *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	this := GetOrganizationDevicesAvailabilities200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork struct {
	value *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) Get() *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) Set(val *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork(val *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) *NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	return &NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesAvailabilities200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


