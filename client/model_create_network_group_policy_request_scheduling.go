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

// checks if the CreateNetworkGroupPolicyRequestScheduling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequestScheduling{}

// CreateNetworkGroupPolicyRequestScheduling     The schedule for the group policy. Schedules are applied to days of the week. 
type CreateNetworkGroupPolicyRequestScheduling struct {
	// Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.
	Enabled *bool `json:"enabled,omitempty"`
	Monday *CreateNetworkGroupPolicyRequestSchedulingMonday `json:"monday,omitempty"`
	Tuesday *CreateNetworkGroupPolicyRequestSchedulingTuesday `json:"tuesday,omitempty"`
	Wednesday *CreateNetworkGroupPolicyRequestSchedulingWednesday `json:"wednesday,omitempty"`
	Thursday *CreateNetworkGroupPolicyRequestSchedulingThursday `json:"thursday,omitempty"`
	Friday *CreateNetworkGroupPolicyRequestSchedulingFriday `json:"friday,omitempty"`
	Saturday *CreateNetworkGroupPolicyRequestSchedulingSaturday `json:"saturday,omitempty"`
	Sunday *CreateNetworkGroupPolicyRequestSchedulingSunday `json:"sunday,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestScheduling instantiates a new CreateNetworkGroupPolicyRequestScheduling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestScheduling() *CreateNetworkGroupPolicyRequestScheduling {
	this := CreateNetworkGroupPolicyRequestScheduling{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestSchedulingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestScheduling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestSchedulingWithDefaults() *CreateNetworkGroupPolicyRequestScheduling {
	this := CreateNetworkGroupPolicyRequestScheduling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetMonday() CreateNetworkGroupPolicyRequestSchedulingMonday {
	if o == nil || IsNil(o.Monday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetMondayOk() (*CreateNetworkGroupPolicyRequestSchedulingMonday, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingMonday and assigns it to the Monday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetMonday(v CreateNetworkGroupPolicyRequestSchedulingMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetTuesday() CreateNetworkGroupPolicyRequestSchedulingTuesday {
	if o == nil || IsNil(o.Tuesday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetTuesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingTuesday, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingTuesday and assigns it to the Tuesday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetTuesday(v CreateNetworkGroupPolicyRequestSchedulingTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetWednesday() CreateNetworkGroupPolicyRequestSchedulingWednesday {
	if o == nil || IsNil(o.Wednesday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetWednesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingWednesday, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingWednesday and assigns it to the Wednesday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetWednesday(v CreateNetworkGroupPolicyRequestSchedulingWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetThursday() CreateNetworkGroupPolicyRequestSchedulingThursday {
	if o == nil || IsNil(o.Thursday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetThursdayOk() (*CreateNetworkGroupPolicyRequestSchedulingThursday, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingThursday and assigns it to the Thursday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetThursday(v CreateNetworkGroupPolicyRequestSchedulingThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetFriday() CreateNetworkGroupPolicyRequestSchedulingFriday {
	if o == nil || IsNil(o.Friday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetFridayOk() (*CreateNetworkGroupPolicyRequestSchedulingFriday, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingFriday and assigns it to the Friday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetFriday(v CreateNetworkGroupPolicyRequestSchedulingFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetSaturday() CreateNetworkGroupPolicyRequestSchedulingSaturday {
	if o == nil || IsNil(o.Saturday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetSaturdayOk() (*CreateNetworkGroupPolicyRequestSchedulingSaturday, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingSaturday and assigns it to the Saturday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetSaturday(v CreateNetworkGroupPolicyRequestSchedulingSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetSunday() CreateNetworkGroupPolicyRequestSchedulingSunday {
	if o == nil || IsNil(o.Sunday) {
		var ret CreateNetworkGroupPolicyRequestSchedulingSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) GetSundayOk() (*CreateNetworkGroupPolicyRequestSchedulingSunday, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestScheduling) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given CreateNetworkGroupPolicyRequestSchedulingSunday and assigns it to the Sunday field.
func (o *CreateNetworkGroupPolicyRequestScheduling) SetSunday(v CreateNetworkGroupPolicyRequestSchedulingSunday) {
	o.Sunday = &v
}

func (o CreateNetworkGroupPolicyRequestScheduling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequestScheduling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !IsNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !IsNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !IsNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !IsNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !IsNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	if !IsNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	return toSerialize, nil
}

type NullableCreateNetworkGroupPolicyRequestScheduling struct {
	value *CreateNetworkGroupPolicyRequestScheduling
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestScheduling) Get() *CreateNetworkGroupPolicyRequestScheduling {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestScheduling) Set(val *CreateNetworkGroupPolicyRequestScheduling) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestScheduling) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestScheduling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestScheduling(val *CreateNetworkGroupPolicyRequestScheduling) *NullableCreateNetworkGroupPolicyRequestScheduling {
	return &NullableCreateNetworkGroupPolicyRequestScheduling{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestScheduling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestScheduling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


