/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkSettings200ResponseClientPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseClientPrivacy{}

// GetNetworkSettings200ResponseClientPrivacy Privacy settings
type GetNetworkSettings200ResponseClientPrivacy struct {
	// The number of days, weeks, or months in Epoch time to expire the data before
	ExpireDataOlderThan *int32 `json:"expireDataOlderThan,omitempty"`
	// The date to expire the data before
	ExpireDataBefore *time.Time `json:"expireDataBefore,omitempty"`
}

// NewGetNetworkSettings200ResponseClientPrivacy instantiates a new GetNetworkSettings200ResponseClientPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseClientPrivacy() *GetNetworkSettings200ResponseClientPrivacy {
	this := GetNetworkSettings200ResponseClientPrivacy{}
	return &this
}

// NewGetNetworkSettings200ResponseClientPrivacyWithDefaults instantiates a new GetNetworkSettings200ResponseClientPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseClientPrivacyWithDefaults() *GetNetworkSettings200ResponseClientPrivacy {
	this := GetNetworkSettings200ResponseClientPrivacy{}
	return &this
}

// GetExpireDataOlderThan returns the ExpireDataOlderThan field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseClientPrivacy) GetExpireDataOlderThan() int32 {
	if o == nil || IsNil(o.ExpireDataOlderThan) {
		var ret int32
		return ret
	}
	return *o.ExpireDataOlderThan
}

// GetExpireDataOlderThanOk returns a tuple with the ExpireDataOlderThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseClientPrivacy) GetExpireDataOlderThanOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireDataOlderThan) {
		return nil, false
	}
	return o.ExpireDataOlderThan, true
}

// HasExpireDataOlderThan returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseClientPrivacy) HasExpireDataOlderThan() bool {
	if o != nil && !IsNil(o.ExpireDataOlderThan) {
		return true
	}

	return false
}

// SetExpireDataOlderThan gets a reference to the given int32 and assigns it to the ExpireDataOlderThan field.
func (o *GetNetworkSettings200ResponseClientPrivacy) SetExpireDataOlderThan(v int32) {
	o.ExpireDataOlderThan = &v
}

// GetExpireDataBefore returns the ExpireDataBefore field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseClientPrivacy) GetExpireDataBefore() time.Time {
	if o == nil || IsNil(o.ExpireDataBefore) {
		var ret time.Time
		return ret
	}
	return *o.ExpireDataBefore
}

// GetExpireDataBeforeOk returns a tuple with the ExpireDataBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseClientPrivacy) GetExpireDataBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpireDataBefore) {
		return nil, false
	}
	return o.ExpireDataBefore, true
}

// HasExpireDataBefore returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseClientPrivacy) HasExpireDataBefore() bool {
	if o != nil && !IsNil(o.ExpireDataBefore) {
		return true
	}

	return false
}

// SetExpireDataBefore gets a reference to the given time.Time and assigns it to the ExpireDataBefore field.
func (o *GetNetworkSettings200ResponseClientPrivacy) SetExpireDataBefore(v time.Time) {
	o.ExpireDataBefore = &v
}

func (o GetNetworkSettings200ResponseClientPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseClientPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpireDataOlderThan) {
		toSerialize["expireDataOlderThan"] = o.ExpireDataOlderThan
	}
	if !IsNil(o.ExpireDataBefore) {
		toSerialize["expireDataBefore"] = o.ExpireDataBefore
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseClientPrivacy struct {
	value *GetNetworkSettings200ResponseClientPrivacy
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseClientPrivacy) Get() *GetNetworkSettings200ResponseClientPrivacy {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseClientPrivacy) Set(val *GetNetworkSettings200ResponseClientPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseClientPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseClientPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseClientPrivacy(val *GetNetworkSettings200ResponseClientPrivacy) *NullableGetNetworkSettings200ResponseClientPrivacy {
	return &NullableGetNetworkSettings200ResponseClientPrivacy{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseClientPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseClientPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


