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

// checks if the GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource{}

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource Source of 'custom' type traffic filter
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address (e.g.\"192.168.10.1\", which is the same as \"192.168.10.1/32\"), or \"any\". Cannot be used in combination with the \"vlan\" property
	Cidr *string `json:"cidr,omitempty"`
	// VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the \"cidr\" property and is currently only available under a template network.
	Vlan *int32 `json:"vlan,omitempty"`
	// Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the \"vlan\" property and is currently only available under a template network.
	Host *int32 `json:"host,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource{}
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSourceWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSourceWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) SetCidr(v string) {
	o.Cidr = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) SetVlan(v int32) {
	o.Vlan = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetHost() int32 {
	if o == nil || IsNil(o.Host) {
		var ret int32
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) GetHostOk() (*int32, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given int32 and assigns it to the Host field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) SetHost(v int32) {
	o.Host = &v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


