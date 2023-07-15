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

// checks if the GetOrganizationSensorReadingsLatest200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsLatest200ResponseInner{}

// GetOrganizationSensorReadingsLatest200ResponseInner struct for GetOrganizationSensorReadingsLatest200ResponseInner
type GetOrganizationSensorReadingsLatest200ResponseInner struct {
	// Serial number of the sensor that took the readings.
	Serial *string `json:"serial,omitempty"`
	Network *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork `json:"network,omitempty"`
	// Array of latest readings from the sensor. Each object represents a single reading for a single metric.
	Readings []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner `json:"readings,omitempty"`
}

// NewGetOrganizationSensorReadingsLatest200ResponseInner instantiates a new GetOrganizationSensorReadingsLatest200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsLatest200ResponseInner() *GetOrganizationSensorReadingsLatest200ResponseInner {
	this := GetOrganizationSensorReadingsLatest200ResponseInner{}
	return &this
}

// NewGetOrganizationSensorReadingsLatest200ResponseInnerWithDefaults instantiates a new GetOrganizationSensorReadingsLatest200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsLatest200ResponseInnerWithDefaults() *GetOrganizationSensorReadingsLatest200ResponseInner {
	this := GetOrganizationSensorReadingsLatest200ResponseInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetNetwork() GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetNetworkOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetNetwork(v GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) {
	o.Network = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetReadings() []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner {
	if o == nil || IsNil(o.Readings) {
		var ret []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetReadingsOk() ([]GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner, bool) {
	if o == nil || IsNil(o.Readings) {
		return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasReadings() bool {
	if o != nil && !IsNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner and assigns it to the Readings field.
func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetReadings(v []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) {
	o.Readings = v
}

func (o GetOrganizationSensorReadingsLatest200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsLatest200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsLatest200ResponseInner struct {
	value *GetOrganizationSensorReadingsLatest200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInner) Get() *GetOrganizationSensorReadingsLatest200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInner) Set(val *GetOrganizationSensorReadingsLatest200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsLatest200ResponseInner(val *GetOrganizationSensorReadingsLatest200ResponseInner) *NullableGetOrganizationSensorReadingsLatest200ResponseInner {
	return &NullableGetOrganizationSensorReadingsLatest200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsLatest200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsLatest200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


