/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkSensorMqttBrokers200ResponseInner struct for GetNetworkSensorMqttBrokers200ResponseInner
type GetNetworkSensorMqttBrokers200ResponseInner struct {
	// ID of the MQTT Broker.
	MqttBrokerId *string `json:"mqttBrokerId,omitempty"`
	// Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkSensorMqttBrokers200ResponseInner instantiates a new GetNetworkSensorMqttBrokers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorMqttBrokers200ResponseInner() *GetNetworkSensorMqttBrokers200ResponseInner {
	this := GetNetworkSensorMqttBrokers200ResponseInner{}
	return &this
}

// NewGetNetworkSensorMqttBrokers200ResponseInnerWithDefaults instantiates a new GetNetworkSensorMqttBrokers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorMqttBrokers200ResponseInnerWithDefaults() *GetNetworkSensorMqttBrokers200ResponseInner {
	this := GetNetworkSensorMqttBrokers200ResponseInner{}
	return &this
}

// GetMqttBrokerId returns the MqttBrokerId field value if set, zero value otherwise.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) GetMqttBrokerId() string {
	if o == nil || isNil(o.MqttBrokerId) {
		var ret string
		return ret
	}
	return *o.MqttBrokerId
}

// GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) GetMqttBrokerIdOk() (*string, bool) {
	if o == nil || isNil(o.MqttBrokerId) {
    return nil, false
	}
	return o.MqttBrokerId, true
}

// HasMqttBrokerId returns a boolean if a field has been set.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) HasMqttBrokerId() bool {
	if o != nil && !isNil(o.MqttBrokerId) {
		return true
	}

	return false
}

// SetMqttBrokerId gets a reference to the given string and assigns it to the MqttBrokerId field.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) SetMqttBrokerId(v string) {
	o.MqttBrokerId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSensorMqttBrokers200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkSensorMqttBrokers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MqttBrokerId) {
		toSerialize["mqttBrokerId"] = o.MqttBrokerId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorMqttBrokers200ResponseInner struct {
	value *GetNetworkSensorMqttBrokers200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSensorMqttBrokers200ResponseInner) Get() *GetNetworkSensorMqttBrokers200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSensorMqttBrokers200ResponseInner) Set(val *GetNetworkSensorMqttBrokers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorMqttBrokers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorMqttBrokers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorMqttBrokers200ResponseInner(val *GetNetworkSensorMqttBrokers200ResponseInner) *NullableGetNetworkSensorMqttBrokers200ResponseInner {
	return &NullableGetNetworkSensorMqttBrokers200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSensorMqttBrokers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorMqttBrokers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


