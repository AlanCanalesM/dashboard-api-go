/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{}

// CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct for CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	// The mac address of the trusted server being added
	Mac string `json:"mac"`
	// The VLAN of the trusted server being added. It must be between 1 and 4094
	Vlan int32 `json:"vlan"`
	Ipv4 CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 `json:"ipv4"`
}

// NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest instantiates a new CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(mac string, vlan int32, ipv4 CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	this := CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{}
	this.Mac = mac
	this.Vlan = vlan
	this.Ipv4 = ipv4
	return &this
}

// NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestWithDefaults instantiates a new CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestWithDefaults() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	this := CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{}
	return &this
}

// GetMac returns the Mac field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) SetMac(v string) {
	o.Mac = v
}

// GetVlan returns the Vlan field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) SetVlan(v int32) {
	o.Vlan = v
}

// GetIpv4 returns the Ipv4 field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetIpv4() CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	if o == nil {
		var ret CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) GetIpv4Ok() (*CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) SetIpv4(v CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) {
	o.Ipv4 = v
}

func (o CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	toSerialize["vlan"] = o.Vlan
	toSerialize["ipv4"] = o.Ipv4
	return toSerialize, nil
}

type NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	value *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
	isSet bool
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Get() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return v.value
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Set(val *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(val *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return &NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


