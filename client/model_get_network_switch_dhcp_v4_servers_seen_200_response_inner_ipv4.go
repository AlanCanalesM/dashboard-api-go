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

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 IPv4 attributes of the server.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 struct {
	// IPv4 address of the server.
	Address *string `json:"address,omitempty"`
	// Subnet of the server.
	Subnet *string `json:"subnet,omitempty"`
	// IPv4 gateway address of the server.
	Gateway *string `json:"gateway,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4WithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4WithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) SetAddress(v string) {
	o.Address = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) SetGateway(v string) {
	o.Gateway = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


