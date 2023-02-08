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

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket Last packet the server received.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket struct {
	Source *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource `json:"source,omitempty"`
	Destination *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination `json:"destination,omitempty"`
	// Packet type.
	Type *string `json:"type,omitempty"`
	Ethernet *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet `json:"ethernet,omitempty"`
	Ip *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp `json:"ip,omitempty"`
	Udp *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp `json:"udp,omitempty"`
	Fields *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields `json:"fields,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetSource() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource {
	if o == nil || isNil(o.Source) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetSourceOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource and assigns it to the Source field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetSource(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetDestination() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination {
	if o == nil || isNil(o.Destination) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetDestinationOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination and assigns it to the Destination field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetDestination(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketDestination) {
	o.Destination = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetType(v string) {
	o.Type = &v
}

// GetEthernet returns the Ethernet field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetEthernet() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet {
	if o == nil || isNil(o.Ethernet) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet
		return ret
	}
	return *o.Ethernet
}

// GetEthernetOk returns a tuple with the Ethernet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetEthernetOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet, bool) {
	if o == nil || isNil(o.Ethernet) {
    return nil, false
	}
	return o.Ethernet, true
}

// HasEthernet returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasEthernet() bool {
	if o != nil && !isNil(o.Ethernet) {
		return true
	}

	return false
}

// SetEthernet gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet and assigns it to the Ethernet field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetEthernet(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketEthernet) {
	o.Ethernet = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetIp() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp {
	if o == nil || isNil(o.Ip) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetIpOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp and assigns it to the Ip field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetIp(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketIp) {
	o.Ip = &v
}

// GetUdp returns the Udp field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetUdp() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp {
	if o == nil || isNil(o.Udp) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp
		return ret
	}
	return *o.Udp
}

// GetUdpOk returns a tuple with the Udp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetUdpOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp, bool) {
	if o == nil || isNil(o.Udp) {
    return nil, false
	}
	return o.Udp, true
}

// HasUdp returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasUdp() bool {
	if o != nil && !isNil(o.Udp) {
		return true
	}

	return false
}

// SetUdp gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp and assigns it to the Udp field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetUdp(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketUdp) {
	o.Udp = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetFields() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields {
	if o == nil || isNil(o.Fields) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) GetFieldsOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields and assigns it to the Fields field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) SetFields(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) {
	o.Fields = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Ethernet) {
		toSerialize["ethernet"] = o.Ethernet
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Udp) {
		toSerialize["udp"] = o.Udp
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


