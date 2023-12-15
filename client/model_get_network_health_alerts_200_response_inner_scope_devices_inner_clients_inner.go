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

// checks if the GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner{}

// GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner struct for GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner
type GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner struct {
	// Mac address of the client
	Mac *string `json:"mac,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner() *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) SetMac(v string) {
	o.Mac = &v
}

func (o GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner struct {
	value *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) Get() *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) Set(val *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner(val *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) *NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner {
	return &NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


