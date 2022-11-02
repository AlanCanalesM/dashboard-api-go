/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet Additional ethernet attributes of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet struct {
	// Ethernet type of the packet.
	Type *string `json:"type,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernetWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernetWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) SetType(v string) {
	o.Type = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


