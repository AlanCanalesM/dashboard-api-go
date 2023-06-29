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

// checks if the UpdateDeviceSensorRelationshipsRequestLivestream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSensorRelationshipsRequestLivestream{}

// UpdateDeviceSensorRelationshipsRequestLivestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type UpdateDeviceSensorRelationshipsRequestLivestream struct {
	// An array of the related devices for the role
	RelatedDevices []UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner `json:"relatedDevices,omitempty"`
}

// NewUpdateDeviceSensorRelationshipsRequestLivestream instantiates a new UpdateDeviceSensorRelationshipsRequestLivestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSensorRelationshipsRequestLivestream() *UpdateDeviceSensorRelationshipsRequestLivestream {
	this := UpdateDeviceSensorRelationshipsRequestLivestream{}
	return &this
}

// NewUpdateDeviceSensorRelationshipsRequestLivestreamWithDefaults instantiates a new UpdateDeviceSensorRelationshipsRequestLivestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSensorRelationshipsRequestLivestreamWithDefaults() *UpdateDeviceSensorRelationshipsRequestLivestream {
	this := UpdateDeviceSensorRelationshipsRequestLivestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *UpdateDeviceSensorRelationshipsRequestLivestream) GetRelatedDevices() []UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner {
	if o == nil || IsNil(o.RelatedDevices) {
		var ret []UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSensorRelationshipsRequestLivestream) GetRelatedDevicesOk() ([]UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner, bool) {
	if o == nil || IsNil(o.RelatedDevices) {
		return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *UpdateDeviceSensorRelationshipsRequestLivestream) HasRelatedDevices() bool {
	if o != nil && !IsNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner and assigns it to the RelatedDevices field.
func (o *UpdateDeviceSensorRelationshipsRequestLivestream) SetRelatedDevices(v []UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) {
	o.RelatedDevices = v
}

func (o UpdateDeviceSensorRelationshipsRequestLivestream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSensorRelationshipsRequestLivestream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return toSerialize, nil
}

type NullableUpdateDeviceSensorRelationshipsRequestLivestream struct {
	value *UpdateDeviceSensorRelationshipsRequestLivestream
	isSet bool
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestream) Get() *UpdateDeviceSensorRelationshipsRequestLivestream {
	return v.value
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestream) Set(val *UpdateDeviceSensorRelationshipsRequestLivestream) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestream) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSensorRelationshipsRequestLivestream(val *UpdateDeviceSensorRelationshipsRequestLivestream) *NullableUpdateDeviceSensorRelationshipsRequestLivestream {
	return &NullableUpdateDeviceSensorRelationshipsRequestLivestream{value: val, isSet: true}
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


