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

// checks if the GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner{}

// GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner struct for GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner
type GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner struct {
	// Switch serial number
	Serial *string `json:"serial,omitempty"`
	// Switch alternative management IP. To remove a prior IP setting, provide an empty string
	AlternateManagementIp *string `json:"alternateManagementIp,omitempty"`
	// Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Gateway *string `json:"gateway,omitempty"`
}

// NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner instantiates a new GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner() *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner {
	this := GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner{}
	return &this
}

// NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInnerWithDefaults instantiates a new GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInnerWithDefaults() *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner {
	this := GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetAlternateManagementIp() string {
	if o == nil || IsNil(o.AlternateManagementIp) {
		var ret string
		return ret
	}
	return *o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateManagementIp) {
		return nil, false
	}
	return o.AlternateManagementIp, true
}

// HasAlternateManagementIp returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasAlternateManagementIp() bool {
	if o != nil && !IsNil(o.AlternateManagementIp) {
		return true
	}

	return false
}

// SetAlternateManagementIp gets a reference to the given string and assigns it to the AlternateManagementIp field.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSubnetMask() string {
	if o == nil || IsNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetMask) {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasSubnetMask() bool {
	if o != nil && !IsNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetGateway(v string) {
	o.Gateway = &v
}

func (o GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.AlternateManagementIp) {
		toSerialize["alternateManagementIp"] = o.AlternateManagementIp
	}
	if !IsNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner struct {
	value *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner
	isSet bool
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) Get() *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner {
	return v.value
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) Set(val *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner(val *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) *NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner {
	return &NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


