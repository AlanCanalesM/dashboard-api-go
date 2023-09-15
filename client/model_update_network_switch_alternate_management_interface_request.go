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

// checks if the UpdateNetworkSwitchAlternateManagementInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchAlternateManagementInterfaceRequest{}

// UpdateNetworkSwitchAlternateManagementInterfaceRequest struct for UpdateNetworkSwitchAlternateManagementInterfaceRequest
type UpdateNetworkSwitchAlternateManagementInterfaceRequest struct {
	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols,omitempty"`
	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner `json:"switches,omitempty"`
}

// NewUpdateNetworkSwitchAlternateManagementInterfaceRequest instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchAlternateManagementInterfaceRequest() *UpdateNetworkSwitchAlternateManagementInterfaceRequest {
	this := UpdateNetworkSwitchAlternateManagementInterfaceRequest{}
	return &this
}

// NewUpdateNetworkSwitchAlternateManagementInterfaceRequestWithDefaults instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestWithDefaults() *UpdateNetworkSwitchAlternateManagementInterfaceRequest {
	this := UpdateNetworkSwitchAlternateManagementInterfaceRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetProtocols() []string {
	if o == nil || IsNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.Protocols) {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasProtocols() bool {
	if o != nil && !IsNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetProtocols(v []string) {
	o.Protocols = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetSwitches() []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner {
	if o == nil || IsNil(o.Switches) {
		var ret []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetSwitchesOk() ([]UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner and assigns it to the Switches field.
func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetSwitches(v []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) {
	o.Switches = v
}

func (o UpdateNetworkSwitchAlternateManagementInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchAlternateManagementInterfaceRequest) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest struct {
	value *UpdateNetworkSwitchAlternateManagementInterfaceRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) Get() *UpdateNetworkSwitchAlternateManagementInterfaceRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) Set(val *UpdateNetworkSwitchAlternateManagementInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchAlternateManagementInterfaceRequest(val *UpdateNetworkSwitchAlternateManagementInterfaceRequest) *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest {
	return &NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchAlternateManagementInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


