/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject111 struct for InlineObject111
type InlineObject111 struct {
	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols,omitempty"`
	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`
}

// NewInlineObject111 instantiates a new InlineObject111 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject111() *InlineObject111 {
	this := InlineObject111{}
	return &this
}

// NewInlineObject111WithDefaults instantiates a new InlineObject111 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject111WithDefaults() *InlineObject111 {
	this := InlineObject111{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject111) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject111) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject111) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject111) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineObject111) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject111) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject111) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineObject111) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *InlineObject111) GetProtocols() []string {
	if o == nil || isNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject111) GetProtocolsOk() ([]string, bool) {
	if o == nil || isNil(o.Protocols) {
    return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *InlineObject111) HasProtocols() bool {
	if o != nil && !isNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *InlineObject111) SetProtocols(v []string) {
	o.Protocols = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineObject111) GetSwitches() []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	if o == nil || isNil(o.Switches) {
		var ret []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject111) GetSwitchesOk() ([]NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineObject111) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches and assigns it to the Switches field.
func (o *InlineObject111) SetSwitches(v []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) {
	o.Switches = v
}

func (o InlineObject111) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.Protocols) {
		toSerialize["protocols"] = o.Protocols
	}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject111 struct {
	value *InlineObject111
	isSet bool
}

func (v NullableInlineObject111) Get() *InlineObject111 {
	return v.value
}

func (v *NullableInlineObject111) Set(val *InlineObject111) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject111) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject111) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject111(val *InlineObject111) *NullableInlineObject111 {
	return &NullableInlineObject111{value: val, isSet: true}
}

func (v NullableInlineObject111) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject111) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


