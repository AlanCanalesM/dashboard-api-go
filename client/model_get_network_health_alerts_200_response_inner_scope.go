/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkHealthAlerts200ResponseInnerScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInnerScope{}

// GetNetworkHealthAlerts200ResponseInnerScope The scope of the alert
type GetNetworkHealthAlerts200ResponseInnerScope struct {
	// Devices related to the alert
	Devices []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner `json:"devices,omitempty"`
	// Applications related to the alert
	Applications []GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner `json:"applications,omitempty"`
	// Peers related to the alert
	Peers []GetNetworkHealthAlerts200ResponseInnerScopePeersInner `json:"peers,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInnerScope instantiates a new GetNetworkHealthAlerts200ResponseInnerScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInnerScope() *GetNetworkHealthAlerts200ResponseInnerScope {
	this := GetNetworkHealthAlerts200ResponseInnerScope{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerScopeWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerScopeWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScope {
	this := GetNetworkHealthAlerts200ResponseInnerScope{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetDevices() []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner {
	if o == nil || IsNil(o.Devices) {
		var ret []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetDevicesOk() ([]GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner and assigns it to the Devices field.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) SetDevices(v []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) {
	o.Devices = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetApplications() []GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner {
	if o == nil || IsNil(o.Applications) {
		var ret []GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetApplicationsOk() ([]GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner and assigns it to the Applications field.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) SetApplications(v []GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) {
	o.Applications = v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetPeers() []GetNetworkHealthAlerts200ResponseInnerScopePeersInner {
	if o == nil || IsNil(o.Peers) {
		var ret []GetNetworkHealthAlerts200ResponseInnerScopePeersInner
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) GetPeersOk() ([]GetNetworkHealthAlerts200ResponseInnerScopePeersInner, bool) {
	if o == nil || IsNil(o.Peers) {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) HasPeers() bool {
	if o != nil && !IsNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []GetNetworkHealthAlerts200ResponseInnerScopePeersInner and assigns it to the Peers field.
func (o *GetNetworkHealthAlerts200ResponseInnerScope) SetPeers(v []GetNetworkHealthAlerts200ResponseInnerScopePeersInner) {
	o.Peers = v
}

func (o GetNetworkHealthAlerts200ResponseInnerScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInnerScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInnerScope struct {
	value *GetNetworkHealthAlerts200ResponseInnerScope
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScope) Get() *GetNetworkHealthAlerts200ResponseInnerScope {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScope) Set(val *GetNetworkHealthAlerts200ResponseInnerScope) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScope) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInnerScope(val *GetNetworkHealthAlerts200ResponseInnerScope) *NullableGetNetworkHealthAlerts200ResponseInnerScope {
	return &NullableGetNetworkHealthAlerts200ResponseInnerScope{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


