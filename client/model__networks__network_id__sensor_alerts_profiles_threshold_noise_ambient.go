/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient Ambient noise threshold. One of 'level' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient struct {
	// Alerting threshold as adjusted decibels.
	Level *int32 `json:"level,omitempty"`
	// Alerting threshold as a qualitative ambient noise level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient() *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbientWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbientWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) GetLevel() int32 {
	if o == nil || isNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) GetLevelOk() (*int32, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) SetLevel(v int32) {
	o.Level = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient(val *NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


