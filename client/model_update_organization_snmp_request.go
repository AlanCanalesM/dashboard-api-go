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

// checks if the UpdateOrganizationSnmpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSnmpRequest{}

// UpdateOrganizationSnmpRequest struct for UpdateOrganizationSnmpRequest
type UpdateOrganizationSnmpRequest struct {
	// Boolean indicating whether SNMP version 2c is enabled for the organization.
	V2cEnabled *bool `json:"v2cEnabled,omitempty"`
	// Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3Enabled *bool `json:"v3Enabled,omitempty"`
	// The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
	V3AuthMode *string `json:"v3AuthMode,omitempty"`
	// The SNMP version 3 authentication password. Must be at least 8 characters if specified.
	V3AuthPass *string `json:"v3AuthPass,omitempty"`
	// The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
	V3PrivMode *string `json:"v3PrivMode,omitempty"`
	// The SNMP version 3 privacy password. Must be at least 8 characters if specified.
	V3PrivPass *string `json:"v3PrivPass,omitempty"`
	// The list of IPv4 addresses that are allowed to access the SNMP server.
	PeerIps []string `json:"peerIps,omitempty"`
}

// NewUpdateOrganizationSnmpRequest instantiates a new UpdateOrganizationSnmpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSnmpRequest() *UpdateOrganizationSnmpRequest {
	this := UpdateOrganizationSnmpRequest{}
	return &this
}

// NewUpdateOrganizationSnmpRequestWithDefaults instantiates a new UpdateOrganizationSnmpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSnmpRequestWithDefaults() *UpdateOrganizationSnmpRequest {
	this := UpdateOrganizationSnmpRequest{}
	return &this
}

// GetV2cEnabled returns the V2cEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV2cEnabled() bool {
	if o == nil || IsNil(o.V2cEnabled) {
		var ret bool
		return ret
	}
	return *o.V2cEnabled
}

// GetV2cEnabledOk returns a tuple with the V2cEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV2cEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V2cEnabled) {
		return nil, false
	}
	return o.V2cEnabled, true
}

// HasV2cEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV2cEnabled() bool {
	if o != nil && !IsNil(o.V2cEnabled) {
		return true
	}

	return false
}

// SetV2cEnabled gets a reference to the given bool and assigns it to the V2cEnabled field.
func (o *UpdateOrganizationSnmpRequest) SetV2cEnabled(v bool) {
	o.V2cEnabled = &v
}

// GetV3Enabled returns the V3Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV3Enabled() bool {
	if o == nil || IsNil(o.V3Enabled) {
		var ret bool
		return ret
	}
	return *o.V3Enabled
}

// GetV3EnabledOk returns a tuple with the V3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V3Enabled) {
		return nil, false
	}
	return o.V3Enabled, true
}

// HasV3Enabled returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV3Enabled() bool {
	if o != nil && !IsNil(o.V3Enabled) {
		return true
	}

	return false
}

// SetV3Enabled gets a reference to the given bool and assigns it to the V3Enabled field.
func (o *UpdateOrganizationSnmpRequest) SetV3Enabled(v bool) {
	o.V3Enabled = &v
}

// GetV3AuthMode returns the V3AuthMode field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV3AuthMode() string {
	if o == nil || IsNil(o.V3AuthMode) {
		var ret string
		return ret
	}
	return *o.V3AuthMode
}

// GetV3AuthModeOk returns a tuple with the V3AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV3AuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.V3AuthMode) {
		return nil, false
	}
	return o.V3AuthMode, true
}

// HasV3AuthMode returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV3AuthMode() bool {
	if o != nil && !IsNil(o.V3AuthMode) {
		return true
	}

	return false
}

// SetV3AuthMode gets a reference to the given string and assigns it to the V3AuthMode field.
func (o *UpdateOrganizationSnmpRequest) SetV3AuthMode(v string) {
	o.V3AuthMode = &v
}

// GetV3AuthPass returns the V3AuthPass field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV3AuthPass() string {
	if o == nil || IsNil(o.V3AuthPass) {
		var ret string
		return ret
	}
	return *o.V3AuthPass
}

// GetV3AuthPassOk returns a tuple with the V3AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV3AuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.V3AuthPass) {
		return nil, false
	}
	return o.V3AuthPass, true
}

// HasV3AuthPass returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV3AuthPass() bool {
	if o != nil && !IsNil(o.V3AuthPass) {
		return true
	}

	return false
}

// SetV3AuthPass gets a reference to the given string and assigns it to the V3AuthPass field.
func (o *UpdateOrganizationSnmpRequest) SetV3AuthPass(v string) {
	o.V3AuthPass = &v
}

// GetV3PrivMode returns the V3PrivMode field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV3PrivMode() string {
	if o == nil || IsNil(o.V3PrivMode) {
		var ret string
		return ret
	}
	return *o.V3PrivMode
}

// GetV3PrivModeOk returns a tuple with the V3PrivMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV3PrivModeOk() (*string, bool) {
	if o == nil || IsNil(o.V3PrivMode) {
		return nil, false
	}
	return o.V3PrivMode, true
}

// HasV3PrivMode returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV3PrivMode() bool {
	if o != nil && !IsNil(o.V3PrivMode) {
		return true
	}

	return false
}

// SetV3PrivMode gets a reference to the given string and assigns it to the V3PrivMode field.
func (o *UpdateOrganizationSnmpRequest) SetV3PrivMode(v string) {
	o.V3PrivMode = &v
}

// GetV3PrivPass returns the V3PrivPass field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetV3PrivPass() string {
	if o == nil || IsNil(o.V3PrivPass) {
		var ret string
		return ret
	}
	return *o.V3PrivPass
}

// GetV3PrivPassOk returns a tuple with the V3PrivPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetV3PrivPassOk() (*string, bool) {
	if o == nil || IsNil(o.V3PrivPass) {
		return nil, false
	}
	return o.V3PrivPass, true
}

// HasV3PrivPass returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasV3PrivPass() bool {
	if o != nil && !IsNil(o.V3PrivPass) {
		return true
	}

	return false
}

// SetV3PrivPass gets a reference to the given string and assigns it to the V3PrivPass field.
func (o *UpdateOrganizationSnmpRequest) SetV3PrivPass(v string) {
	o.V3PrivPass = &v
}

// GetPeerIps returns the PeerIps field value if set, zero value otherwise.
func (o *UpdateOrganizationSnmpRequest) GetPeerIps() []string {
	if o == nil || IsNil(o.PeerIps) {
		var ret []string
		return ret
	}
	return o.PeerIps
}

// GetPeerIpsOk returns a tuple with the PeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSnmpRequest) GetPeerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.PeerIps) {
		return nil, false
	}
	return o.PeerIps, true
}

// HasPeerIps returns a boolean if a field has been set.
func (o *UpdateOrganizationSnmpRequest) HasPeerIps() bool {
	if o != nil && !IsNil(o.PeerIps) {
		return true
	}

	return false
}

// SetPeerIps gets a reference to the given []string and assigns it to the PeerIps field.
func (o *UpdateOrganizationSnmpRequest) SetPeerIps(v []string) {
	o.PeerIps = v
}

func (o UpdateOrganizationSnmpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSnmpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.V2cEnabled) {
		toSerialize["v2cEnabled"] = o.V2cEnabled
	}
	if !IsNil(o.V3Enabled) {
		toSerialize["v3Enabled"] = o.V3Enabled
	}
	if !IsNil(o.V3AuthMode) {
		toSerialize["v3AuthMode"] = o.V3AuthMode
	}
	if !IsNil(o.V3AuthPass) {
		toSerialize["v3AuthPass"] = o.V3AuthPass
	}
	if !IsNil(o.V3PrivMode) {
		toSerialize["v3PrivMode"] = o.V3PrivMode
	}
	if !IsNil(o.V3PrivPass) {
		toSerialize["v3PrivPass"] = o.V3PrivPass
	}
	if !IsNil(o.PeerIps) {
		toSerialize["peerIps"] = o.PeerIps
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationSnmpRequest struct {
	value *UpdateOrganizationSnmpRequest
	isSet bool
}

func (v NullableUpdateOrganizationSnmpRequest) Get() *UpdateOrganizationSnmpRequest {
	return v.value
}

func (v *NullableUpdateOrganizationSnmpRequest) Set(val *UpdateOrganizationSnmpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSnmpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSnmpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSnmpRequest(val *UpdateOrganizationSnmpRequest) *NullableUpdateOrganizationSnmpRequest {
	return &NullableUpdateOrganizationSnmpRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSnmpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSnmpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


