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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInner{}

// GetNetworkSensorAlertsProfiles200ResponseInner struct for GetNetworkSensorAlertsProfiles200ResponseInner
type GetNetworkSensorAlertsProfiles200ResponseInner struct {
	// ID of the sensor alert profile.
	ProfileId *string `json:"profileId,omitempty"`
	// Name of the sensor alert profile.
	Name *string `json:"name,omitempty"`
	Schedule *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule `json:"schedule,omitempty"`
	// List of conditions that will cause the profile to send an alert.
	Conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner `json:"conditions"`
	Recipients *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients `json:"recipients,omitempty"`
	// List of device serials assigned to this sensor alert profile.
	Serials []string `json:"serials,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInner instantiates a new GetNetworkSensorAlertsProfiles200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInner(conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) *GetNetworkSensorAlertsProfiles200ResponseInner {
	this := GetNetworkSensorAlertsProfiles200ResponseInner{}
	this.Conditions = conditions
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInner {
	this := GetNetworkSensorAlertsProfiles200ResponseInner{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSchedule() GetNetworkSensorAlertsProfiles200ResponseInnerSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetScheduleOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given GetNetworkSensorAlertsProfiles200ResponseInnerSchedule and assigns it to the Schedule field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetSchedule(v GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) {
	o.Schedule = &v
}

// GetConditions returns the Conditions field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	if o == nil {
		var ret []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetConditionsOk() ([]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) {
	o.Conditions = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	if o == nil || IsNil(o.Recipients) {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given GetNetworkSensorAlertsProfiles200ResponseInnerRecipients and assigns it to the Recipients field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) {
	o.Recipients = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetSerials(v []string) {
	o.Serials = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	toSerialize["conditions"] = o.Conditions
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInner struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInner) Get() *GetNetworkSensorAlertsProfiles200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInner) Set(val *GetNetworkSensorAlertsProfiles200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInner(val *GetNetworkSensorAlertsProfiles200ResponseInner) *NullableGetNetworkSensorAlertsProfiles200ResponseInner {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


