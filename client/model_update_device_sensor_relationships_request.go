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

// UpdateDeviceSensorRelationshipsRequest struct for UpdateDeviceSensorRelationshipsRequest
type UpdateDeviceSensorRelationshipsRequest struct {
	Livestream *UpdateDeviceSensorRelationshipsRequestLivestream `json:"livestream,omitempty"`
}

// NewUpdateDeviceSensorRelationshipsRequest instantiates a new UpdateDeviceSensorRelationshipsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSensorRelationshipsRequest() *UpdateDeviceSensorRelationshipsRequest {
	this := UpdateDeviceSensorRelationshipsRequest{}
	return &this
}

// NewUpdateDeviceSensorRelationshipsRequestWithDefaults instantiates a new UpdateDeviceSensorRelationshipsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSensorRelationshipsRequestWithDefaults() *UpdateDeviceSensorRelationshipsRequest {
	this := UpdateDeviceSensorRelationshipsRequest{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *UpdateDeviceSensorRelationshipsRequest) GetLivestream() UpdateDeviceSensorRelationshipsRequestLivestream {
	if o == nil || isNil(o.Livestream) {
		var ret UpdateDeviceSensorRelationshipsRequestLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSensorRelationshipsRequest) GetLivestreamOk() (*UpdateDeviceSensorRelationshipsRequestLivestream, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *UpdateDeviceSensorRelationshipsRequest) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given UpdateDeviceSensorRelationshipsRequestLivestream and assigns it to the Livestream field.
func (o *UpdateDeviceSensorRelationshipsRequest) SetLivestream(v UpdateDeviceSensorRelationshipsRequestLivestream) {
	o.Livestream = &v
}

func (o UpdateDeviceSensorRelationshipsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceSensorRelationshipsRequest struct {
	value *UpdateDeviceSensorRelationshipsRequest
	isSet bool
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) Get() *UpdateDeviceSensorRelationshipsRequest {
	return v.value
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) Set(val *UpdateDeviceSensorRelationshipsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSensorRelationshipsRequest(val *UpdateDeviceSensorRelationshipsRequest) *NullableUpdateDeviceSensorRelationshipsRequest {
	return &NullableUpdateDeviceSensorRelationshipsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


