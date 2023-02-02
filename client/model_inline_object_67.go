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

// InlineObject67 struct for InlineObject67
type InlineObject67 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject67 instantiates a new InlineObject67 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject67() *InlineObject67 {
	this := InlineObject67{}
	return &this
}

// NewInlineObject67WithDefaults instantiates a new InlineObject67 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject67WithDefaults() *InlineObject67 {
	this := InlineObject67{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject67) GetDestinations() []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject67) GetDestinationsOk() ([]NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject67) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject67) SetDestinations(v []NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject67) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject67 struct {
	value *InlineObject67
	isSet bool
}

func (v NullableInlineObject67) Get() *InlineObject67 {
	return v.value
}

func (v *NullableInlineObject67) Set(val *InlineObject67) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject67) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject67) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject67(val *InlineObject67) *NullableInlineObject67 {
	return &NullableInlineObject67{value: val, isSet: true}
}

func (v NullableInlineObject67) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject67) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


