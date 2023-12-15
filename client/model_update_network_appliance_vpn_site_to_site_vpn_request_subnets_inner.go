/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner{}

// UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner struct for UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner
type UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet string `json:"localSubnet"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner(localSubnet string) *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner{}
	this.LocalSubnet = localSubnet
	return &this
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInnerWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetLocalSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetLocalSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalSubnet, true
}

// SetLocalSubnet sets field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) SetLocalSubnet(v string) {
	o.LocalSubnet = v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetUseVpn() bool {
	if o == nil || IsNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetUseVpnOk() (*bool, bool) {
	if o == nil || IsNil(o.UseVpn) {
		return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) HasUseVpn() bool {
	if o != nil && !IsNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) SetUseVpn(v bool) {
	o.UseVpn = &v
}

func (o UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["localSubnet"] = o.LocalSubnet
	if !IsNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner struct {
	value *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) Get() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) Set(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner {
	return &NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


