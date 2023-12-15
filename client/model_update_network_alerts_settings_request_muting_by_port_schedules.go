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

// checks if the UpdateNetworkAlertsSettingsRequestMutingByPortSchedules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkAlertsSettingsRequestMutingByPortSchedules{}

// UpdateNetworkAlertsSettingsRequestMutingByPortSchedules Mute wireless unreachable alerts based on switch port schedules
type UpdateNetworkAlertsSettingsRequestMutingByPortSchedules struct {
	// If true, then wireless unreachable alerts will be muted when caused by a port schedule
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkAlertsSettingsRequestMutingByPortSchedules instantiates a new UpdateNetworkAlertsSettingsRequestMutingByPortSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAlertsSettingsRequestMutingByPortSchedules() *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules {
	this := UpdateNetworkAlertsSettingsRequestMutingByPortSchedules{}
	return &this
}

// NewUpdateNetworkAlertsSettingsRequestMutingByPortSchedulesWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequestMutingByPortSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAlertsSettingsRequestMutingByPortSchedulesWithDefaults() *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules {
	this := UpdateNetworkAlertsSettingsRequestMutingByPortSchedules{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules struct {
	value *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules
	isSet bool
}

func (v NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) Get() *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules {
	return v.value
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) Set(val *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules(val *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) *NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules {
	return &NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules{value: val, isSet: true}
}

func (v NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMutingByPortSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


