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

// MoveNetworkSmDevicesRequest struct for MoveNetworkSmDevicesRequest
type MoveNetworkSmDevicesRequest struct {
	// The wifiMacs of the devices to be moved.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be moved.
	Ids []string `json:"ids,omitempty"`
	// The serials of the devices to be moved.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.
	Scope []string `json:"scope,omitempty"`
	// The new network to which the devices will be moved.
	NewNetwork string `json:"newNetwork"`
}

// NewMoveNetworkSmDevicesRequest instantiates a new MoveNetworkSmDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveNetworkSmDevicesRequest(newNetwork string) *MoveNetworkSmDevicesRequest {
	this := MoveNetworkSmDevicesRequest{}
	this.NewNetwork = newNetwork
	return &this
}

// NewMoveNetworkSmDevicesRequestWithDefaults instantiates a new MoveNetworkSmDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveNetworkSmDevicesRequestWithDefaults() *MoveNetworkSmDevicesRequest {
	this := MoveNetworkSmDevicesRequest{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *MoveNetworkSmDevicesRequest) GetWifiMacs() []string {
	if o == nil || isNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevicesRequest) GetWifiMacsOk() ([]string, bool) {
	if o == nil || isNil(o.WifiMacs) {
    return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *MoveNetworkSmDevicesRequest) HasWifiMacs() bool {
	if o != nil && !isNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *MoveNetworkSmDevicesRequest) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *MoveNetworkSmDevicesRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *MoveNetworkSmDevicesRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *MoveNetworkSmDevicesRequest) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *MoveNetworkSmDevicesRequest) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevicesRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *MoveNetworkSmDevicesRequest) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *MoveNetworkSmDevicesRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MoveNetworkSmDevicesRequest) GetScope() []string {
	if o == nil || isNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevicesRequest) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MoveNetworkSmDevicesRequest) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *MoveNetworkSmDevicesRequest) SetScope(v []string) {
	o.Scope = v
}

// GetNewNetwork returns the NewNetwork field value
func (o *MoveNetworkSmDevicesRequest) GetNewNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewNetwork
}

// GetNewNetworkOk returns a tuple with the NewNetwork field value
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevicesRequest) GetNewNetworkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NewNetwork, true
}

// SetNewNetwork sets field value
func (o *MoveNetworkSmDevicesRequest) SetNewNetwork(v string) {
	o.NewNetwork = v
}

func (o MoveNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["newNetwork"] = o.NewNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableMoveNetworkSmDevicesRequest struct {
	value *MoveNetworkSmDevicesRequest
	isSet bool
}

func (v NullableMoveNetworkSmDevicesRequest) Get() *MoveNetworkSmDevicesRequest {
	return v.value
}

func (v *NullableMoveNetworkSmDevicesRequest) Set(val *MoveNetworkSmDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveNetworkSmDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveNetworkSmDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveNetworkSmDevicesRequest(val *MoveNetworkSmDevicesRequest) *NullableMoveNetworkSmDevicesRequest {
	return &NullableMoveNetworkSmDevicesRequest{value: val, isSet: true}
}

func (v NullableMoveNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveNetworkSmDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


