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

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 Destination ipv4 attributes of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 struct {
	// Destination ipv4 address of the packet.
	Address *string `json:"address,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4WithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4WithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4 {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestinationIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


