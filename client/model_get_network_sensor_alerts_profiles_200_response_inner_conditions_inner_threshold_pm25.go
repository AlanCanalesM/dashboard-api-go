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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 struct {
	// Alerting threshold as PM2.5 parts per million.
	Concentration *int32 `json:"concentration,omitempty"`
	// Alerting threshold as a qualitative PM2.5 level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25WithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25WithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) GetConcentration() int32 {
	if o == nil || IsNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) GetConcentrationOk() (*int32, bool) {
	if o == nil || IsNil(o.Concentration) {
		return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) HasConcentration() bool {
	if o != nil && !IsNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) SetConcentration(v int32) {
	o.Concentration = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25 {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPm25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


