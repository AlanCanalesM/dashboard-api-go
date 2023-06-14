/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation Link negotiation details object for the port
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation struct {
	// The duplex mode of the port. Can be 'full' or 'half'
	Duplex *string `json:"duplex,omitempty"`
	// The speed of the port
	Speed *int32 `json:"speed,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiationWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiationWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation{}
	return &this
}

// GetDuplex returns the Duplex field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) GetDuplex() string {
	if o == nil || isNil(o.Duplex) {
		var ret string
		return ret
	}
	return *o.Duplex
}

// GetDuplexOk returns a tuple with the Duplex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) GetDuplexOk() (*string, bool) {
	if o == nil || isNil(o.Duplex) {
    return nil, false
	}
	return o.Duplex, true
}

// HasDuplex returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) HasDuplex() bool {
	if o != nil && !isNil(o.Duplex) {
		return true
	}

	return false
}

// SetDuplex gets a reference to the given string and assigns it to the Duplex field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) SetDuplex(v string) {
	o.Duplex = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) GetSpeed() int32 {
	if o == nil || isNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) GetSpeedOk() (*int32, bool) {
	if o == nil || isNil(o.Speed) {
    return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) HasSpeed() bool {
	if o != nil && !isNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) SetSpeed(v int32) {
	o.Speed = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Duplex) {
		toSerialize["duplex"] = o.Duplex
	}
	if !isNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


