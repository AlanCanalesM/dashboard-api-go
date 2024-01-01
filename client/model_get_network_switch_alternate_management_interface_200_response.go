/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSwitchAlternateManagementInterface200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAlternateManagementInterface200Response{}

// GetNetworkSwitchAlternateManagementInterface200Response struct for GetNetworkSwitchAlternateManagementInterface200Response
type GetNetworkSwitchAlternateManagementInterface200Response struct {
	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols,omitempty"`
	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner `json:"switches,omitempty"`
}

// NewGetNetworkSwitchAlternateManagementInterface200Response instantiates a new GetNetworkSwitchAlternateManagementInterface200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAlternateManagementInterface200Response() *GetNetworkSwitchAlternateManagementInterface200Response {
	this := GetNetworkSwitchAlternateManagementInterface200Response{}
	return &this
}

// NewGetNetworkSwitchAlternateManagementInterface200ResponseWithDefaults instantiates a new GetNetworkSwitchAlternateManagementInterface200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAlternateManagementInterface200ResponseWithDefaults() *GetNetworkSwitchAlternateManagementInterface200Response {
	this := GetNetworkSwitchAlternateManagementInterface200Response{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetProtocols() []string {
	if o == nil || IsNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Protocols) {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasProtocols() bool {
	if o != nil && !IsNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetProtocols(v []string) {
	o.Protocols = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetSwitches() []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner {
	if o == nil || IsNil(o.Switches) {
		var ret []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetSwitchesOk() ([]GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner and assigns it to the Switches field.
func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetSwitches(v []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) {
	o.Switches = v
}

func (o GetNetworkSwitchAlternateManagementInterface200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAlternateManagementInterface200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Protocols) {
		toSerialize["protocols"] = o.Protocols
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAlternateManagementInterface200Response struct {
	value *GetNetworkSwitchAlternateManagementInterface200Response
	isSet bool
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200Response) Get() *GetNetworkSwitchAlternateManagementInterface200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200Response) Set(val *GetNetworkSwitchAlternateManagementInterface200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAlternateManagementInterface200Response(val *GetNetworkSwitchAlternateManagementInterface200Response) *NullableGetNetworkSwitchAlternateManagementInterface200Response {
	return &NullableGetNetworkSwitchAlternateManagementInterface200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAlternateManagementInterface200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAlternateManagementInterface200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


