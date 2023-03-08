/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdHealthAlertsScopeNetwork Network of the peer
type NetworksNetworkIdHealthAlertsScopeNetwork struct {
	// Name of the network
	Name *string `json:"name,omitempty"`
	// Id of the network
	Id *string `json:"id,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopeNetwork instantiates a new NetworksNetworkIdHealthAlertsScopeNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopeNetwork() *NetworksNetworkIdHealthAlertsScopeNetwork {
	this := NetworksNetworkIdHealthAlertsScopeNetwork{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeNetworkWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeNetworkWithDefaults() *NetworksNetworkIdHealthAlertsScopeNetwork {
	this := NetworksNetworkIdHealthAlertsScopeNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdHealthAlertsScopeNetwork) SetId(v string) {
	o.Id = &v
}

func (o NetworksNetworkIdHealthAlertsScopeNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopeNetwork struct {
	value *NetworksNetworkIdHealthAlertsScopeNetwork
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopeNetwork) Get() *NetworksNetworkIdHealthAlertsScopeNetwork {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeNetwork) Set(val *NetworksNetworkIdHealthAlertsScopeNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopeNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopeNetwork(val *NetworksNetworkIdHealthAlertsScopeNetwork) *NullableNetworksNetworkIdHealthAlertsScopeNetwork {
	return &NullableNetworksNetworkIdHealthAlertsScopeNetwork{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopeNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


