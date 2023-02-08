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

// CheckinNetworkSmDevicesRequest struct for CheckinNetworkSmDevicesRequest
type CheckinNetworkSmDevicesRequest struct {
	// The wifiMacs of the devices to be checked-in.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be checked-in.
	Ids []string `json:"ids,omitempty"`
	// The serials of the devices to be checked-in.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in.
	Scope []string `json:"scope,omitempty"`
}

// NewCheckinNetworkSmDevicesRequest instantiates a new CheckinNetworkSmDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckinNetworkSmDevicesRequest() *CheckinNetworkSmDevicesRequest {
	this := CheckinNetworkSmDevicesRequest{}
	return &this
}

// NewCheckinNetworkSmDevicesRequestWithDefaults instantiates a new CheckinNetworkSmDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckinNetworkSmDevicesRequestWithDefaults() *CheckinNetworkSmDevicesRequest {
	this := CheckinNetworkSmDevicesRequest{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *CheckinNetworkSmDevicesRequest) GetWifiMacs() []string {
	if o == nil || isNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckinNetworkSmDevicesRequest) GetWifiMacsOk() ([]string, bool) {
	if o == nil || isNil(o.WifiMacs) {
    return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *CheckinNetworkSmDevicesRequest) HasWifiMacs() bool {
	if o != nil && !isNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *CheckinNetworkSmDevicesRequest) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CheckinNetworkSmDevicesRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckinNetworkSmDevicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CheckinNetworkSmDevicesRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CheckinNetworkSmDevicesRequest) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *CheckinNetworkSmDevicesRequest) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckinNetworkSmDevicesRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *CheckinNetworkSmDevicesRequest) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *CheckinNetworkSmDevicesRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CheckinNetworkSmDevicesRequest) GetScope() []string {
	if o == nil || isNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckinNetworkSmDevicesRequest) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CheckinNetworkSmDevicesRequest) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *CheckinNetworkSmDevicesRequest) SetScope(v []string) {
	o.Scope = v
}

func (o CheckinNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMacs) {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableCheckinNetworkSmDevicesRequest struct {
	value *CheckinNetworkSmDevicesRequest
	isSet bool
}

func (v NullableCheckinNetworkSmDevicesRequest) Get() *CheckinNetworkSmDevicesRequest {
	return v.value
}

func (v *NullableCheckinNetworkSmDevicesRequest) Set(val *CheckinNetworkSmDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckinNetworkSmDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckinNetworkSmDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckinNetworkSmDevicesRequest(val *CheckinNetworkSmDevicesRequest) *NullableCheckinNetworkSmDevicesRequest {
	return &NullableCheckinNetworkSmDevicesRequest{value: val, isSet: true}
}

func (v NullableCheckinNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckinNetworkSmDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


