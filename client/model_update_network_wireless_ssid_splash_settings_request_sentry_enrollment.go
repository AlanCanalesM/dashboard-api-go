/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment{}

// UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment Systems Manager sentry enrollment splash settings.
type UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment struct {
	SystemsManagerNetwork *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"`
	// The strength of the enforcement of selected system types. Must be one of: 'focused', 'click-through', and 'strict'.
	Strength *string `json:"strength,omitempty"`
	// The system types that the Sentry enforces. Must be included in: 'iOS, 'Android', 'macOS', and 'Windows'.
	EnforcedSystems []string `json:"enforcedSystems,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment() *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment{}
	return &this
}

// GetSystemsManagerNetwork returns the SystemsManagerNetwork field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetSystemsManagerNetwork() UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork {
	if o == nil || IsNil(o.SystemsManagerNetwork) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork
		return ret
	}
	return *o.SystemsManagerNetwork
}

// GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetSystemsManagerNetworkOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork, bool) {
	if o == nil || IsNil(o.SystemsManagerNetwork) {
		return nil, false
	}
	return o.SystemsManagerNetwork, true
}

// HasSystemsManagerNetwork returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasSystemsManagerNetwork() bool {
	if o != nil && !IsNil(o.SystemsManagerNetwork) {
		return true
	}

	return false
}

// SetSystemsManagerNetwork gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork and assigns it to the SystemsManagerNetwork field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetSystemsManagerNetwork(v UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork) {
	o.SystemsManagerNetwork = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetStrength() string {
	if o == nil || IsNil(o.Strength) {
		var ret string
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetStrengthOk() (*string, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given string and assigns it to the Strength field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetStrength(v string) {
	o.Strength = &v
}

// GetEnforcedSystems returns the EnforcedSystems field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetEnforcedSystems() []string {
	if o == nil || IsNil(o.EnforcedSystems) {
		var ret []string
		return ret
	}
	return o.EnforcedSystems
}

// GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetEnforcedSystemsOk() ([]string, bool) {
	if o == nil || IsNil(o.EnforcedSystems) {
		return nil, false
	}
	return o.EnforcedSystems, true
}

// HasEnforcedSystems returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasEnforcedSystems() bool {
	if o != nil && !IsNil(o.EnforcedSystems) {
		return true
	}

	return false
}

// SetEnforcedSystems gets a reference to the given []string and assigns it to the EnforcedSystems field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetEnforcedSystems(v []string) {
	o.EnforcedSystems = v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemsManagerNetwork) {
		toSerialize["systemsManagerNetwork"] = o.SystemsManagerNetwork
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !IsNil(o.EnforcedSystems) {
		toSerialize["enforcedSystems"] = o.EnforcedSystems
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment(val *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


