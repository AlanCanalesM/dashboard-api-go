/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the BindAdministeredLicensingSubscriptionSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindAdministeredLicensingSubscriptionSubscriptionRequest{}

// BindAdministeredLicensingSubscriptionSubscriptionRequest struct for BindAdministeredLicensingSubscriptionSubscriptionRequest
type BindAdministeredLicensingSubscriptionSubscriptionRequest struct {
	// List of network ids to bind to the subscription
	NetworkIds []string `json:"networkIds"`
}

// NewBindAdministeredLicensingSubscriptionSubscriptionRequest instantiates a new BindAdministeredLicensingSubscriptionSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindAdministeredLicensingSubscriptionSubscriptionRequest(networkIds []string) *BindAdministeredLicensingSubscriptionSubscriptionRequest {
	this := BindAdministeredLicensingSubscriptionSubscriptionRequest{}
	this.NetworkIds = networkIds
	return &this
}

// NewBindAdministeredLicensingSubscriptionSubscriptionRequestWithDefaults instantiates a new BindAdministeredLicensingSubscriptionSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindAdministeredLicensingSubscriptionSubscriptionRequestWithDefaults() *BindAdministeredLicensingSubscriptionSubscriptionRequest {
	this := BindAdministeredLicensingSubscriptionSubscriptionRequest{}
	return &this
}

// GetNetworkIds returns the NetworkIds field value
func (o *BindAdministeredLicensingSubscriptionSubscriptionRequest) GetNetworkIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value
// and a boolean to check if the value has been set.
func (o *BindAdministeredLicensingSubscriptionSubscriptionRequest) GetNetworkIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkIds, true
}

// SetNetworkIds sets field value
func (o *BindAdministeredLicensingSubscriptionSubscriptionRequest) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

func (o BindAdministeredLicensingSubscriptionSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindAdministeredLicensingSubscriptionSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkIds"] = o.NetworkIds
	return toSerialize, nil
}

type NullableBindAdministeredLicensingSubscriptionSubscriptionRequest struct {
	value *BindAdministeredLicensingSubscriptionSubscriptionRequest
	isSet bool
}

func (v NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) Get() *BindAdministeredLicensingSubscriptionSubscriptionRequest {
	return v.value
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) Set(val *BindAdministeredLicensingSubscriptionSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindAdministeredLicensingSubscriptionSubscriptionRequest(val *BindAdministeredLicensingSubscriptionSubscriptionRequest) *NullableBindAdministeredLicensingSubscriptionSubscriptionRequest {
	return &NullableBindAdministeredLicensingSubscriptionSubscriptionRequest{value: val, isSet: true}
}

func (v NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


