/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkSensorAlertsProfileRequestSchedule The sensor schedule to use with the alert profile.
type CreateNetworkSensorAlertsProfileRequestSchedule struct {
	// ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Id *string `json:"id,omitempty"`
}

// NewCreateNetworkSensorAlertsProfileRequestSchedule instantiates a new CreateNetworkSensorAlertsProfileRequestSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSensorAlertsProfileRequestSchedule() *CreateNetworkSensorAlertsProfileRequestSchedule {
	this := CreateNetworkSensorAlertsProfileRequestSchedule{}
	return &this
}

// NewCreateNetworkSensorAlertsProfileRequestScheduleWithDefaults instantiates a new CreateNetworkSensorAlertsProfileRequestSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSensorAlertsProfileRequestScheduleWithDefaults() *CreateNetworkSensorAlertsProfileRequestSchedule {
	this := CreateNetworkSensorAlertsProfileRequestSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkSensorAlertsProfileRequestSchedule) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequestSchedule) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateNetworkSensorAlertsProfileRequestSchedule) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateNetworkSensorAlertsProfileRequestSchedule) SetId(v string) {
	o.Id = &v
}

func (o CreateNetworkSensorAlertsProfileRequestSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkSensorAlertsProfileRequestSchedule struct {
	value *CreateNetworkSensorAlertsProfileRequestSchedule
	isSet bool
}

func (v NullableCreateNetworkSensorAlertsProfileRequestSchedule) Get() *CreateNetworkSensorAlertsProfileRequestSchedule {
	return v.value
}

func (v *NullableCreateNetworkSensorAlertsProfileRequestSchedule) Set(val *CreateNetworkSensorAlertsProfileRequestSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSensorAlertsProfileRequestSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSensorAlertsProfileRequestSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSensorAlertsProfileRequestSchedule(val *CreateNetworkSensorAlertsProfileRequestSchedule) *NullableCreateNetworkSensorAlertsProfileRequestSchedule {
	return &NullableCreateNetworkSensorAlertsProfileRequestSchedule{value: val, isSet: true}
}

func (v NullableCreateNetworkSensorAlertsProfileRequestSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSensorAlertsProfileRequestSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


