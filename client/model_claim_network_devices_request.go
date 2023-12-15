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

// checks if the ClaimNetworkDevicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimNetworkDevicesRequest{}

// ClaimNetworkDevicesRequest struct for ClaimNetworkDevicesRequest
type ClaimNetworkDevicesRequest struct {
	// A list of serials of devices to claim
	Serials []string `json:"serials"`
}

// NewClaimNetworkDevicesRequest instantiates a new ClaimNetworkDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimNetworkDevicesRequest(serials []string) *ClaimNetworkDevicesRequest {
	this := ClaimNetworkDevicesRequest{}
	this.Serials = serials
	return &this
}

// NewClaimNetworkDevicesRequestWithDefaults instantiates a new ClaimNetworkDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimNetworkDevicesRequestWithDefaults() *ClaimNetworkDevicesRequest {
	this := ClaimNetworkDevicesRequest{}
	return &this
}

// GetSerials returns the Serials field value
func (o *ClaimNetworkDevicesRequest) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *ClaimNetworkDevicesRequest) GetSerialsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *ClaimNetworkDevicesRequest) SetSerials(v []string) {
	o.Serials = v
}

func (o ClaimNetworkDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimNetworkDevicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serials"] = o.Serials
	return toSerialize, nil
}

type NullableClaimNetworkDevicesRequest struct {
	value *ClaimNetworkDevicesRequest
	isSet bool
}

func (v NullableClaimNetworkDevicesRequest) Get() *ClaimNetworkDevicesRequest {
	return v.value
}

func (v *NullableClaimNetworkDevicesRequest) Set(val *ClaimNetworkDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimNetworkDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimNetworkDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimNetworkDevicesRequest(val *ClaimNetworkDevicesRequest) *NullableClaimNetworkDevicesRequest {
	return &NullableClaimNetworkDevicesRequest{value: val, isSet: true}
}

func (v NullableClaimNetworkDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimNetworkDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


