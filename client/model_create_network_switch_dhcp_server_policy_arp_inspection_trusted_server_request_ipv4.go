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

// checks if the CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{}

// CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 The IPv4 attributes of the trusted server being added
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 struct {
	// The IPv4 address of the trusted server being added
	Address *string `json:"address,omitempty"`
}

// NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 instantiates a new CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	this := CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{}
	return &this
}

// NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4WithDefaults instantiates a new CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4WithDefaults() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	this := CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 struct {
	value *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4
	isSet bool
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Get() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	return v.value
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Set(val *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4(val *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4 {
	return &NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


