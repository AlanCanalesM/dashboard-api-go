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

// checks if the GetNetworkSwitchDscpToCosMappings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDscpToCosMappings200Response{}

// GetNetworkSwitchDscpToCosMappings200Response struct for GetNetworkSwitchDscpToCosMappings200Response
type GetNetworkSwitchDscpToCosMappings200Response struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner `json:"mappings,omitempty"`
}

// NewGetNetworkSwitchDscpToCosMappings200Response instantiates a new GetNetworkSwitchDscpToCosMappings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDscpToCosMappings200Response() *GetNetworkSwitchDscpToCosMappings200Response {
	this := GetNetworkSwitchDscpToCosMappings200Response{}
	return &this
}

// NewGetNetworkSwitchDscpToCosMappings200ResponseWithDefaults instantiates a new GetNetworkSwitchDscpToCosMappings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDscpToCosMappings200ResponseWithDefaults() *GetNetworkSwitchDscpToCosMappings200Response {
	this := GetNetworkSwitchDscpToCosMappings200Response{}
	return &this
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *GetNetworkSwitchDscpToCosMappings200Response) GetMappings() []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner {
	if o == nil || IsNil(o.Mappings) {
		var ret []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDscpToCosMappings200Response) GetMappingsOk() ([]GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *GetNetworkSwitchDscpToCosMappings200Response) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner and assigns it to the Mappings field.
func (o *GetNetworkSwitchDscpToCosMappings200Response) SetMappings(v []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) {
	o.Mappings = v
}

func (o GetNetworkSwitchDscpToCosMappings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDscpToCosMappings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDscpToCosMappings200Response struct {
	value *GetNetworkSwitchDscpToCosMappings200Response
	isSet bool
}

func (v NullableGetNetworkSwitchDscpToCosMappings200Response) Get() *GetNetworkSwitchDscpToCosMappings200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200Response) Set(val *GetNetworkSwitchDscpToCosMappings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDscpToCosMappings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDscpToCosMappings200Response(val *GetNetworkSwitchDscpToCosMappings200Response) *NullableGetNetworkSwitchDscpToCosMappings200Response {
	return &NullableGetNetworkSwitchDscpToCosMappings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDscpToCosMappings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


