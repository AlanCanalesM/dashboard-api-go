/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20026Milestones The Staged Upgrade Milestones for the stage
type InlineResponse20026Milestones struct {
	// Scheduled start time for the group
	ScheduledFor *time.Time `json:"scheduledFor,omitempty"`
	// Start time for the group
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Finish time for the group
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Time that the group was canceled
	CanceledAt *time.Time `json:"canceledAt,omitempty"`
}

// NewInlineResponse20026Milestones instantiates a new InlineResponse20026Milestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Milestones() *InlineResponse20026Milestones {
	this := InlineResponse20026Milestones{}
	return &this
}

// NewInlineResponse20026MilestonesWithDefaults instantiates a new InlineResponse20026Milestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026MilestonesWithDefaults() *InlineResponse20026Milestones {
	this := InlineResponse20026Milestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *InlineResponse20026Milestones) GetScheduledFor() time.Time {
	if o == nil || isNil(o.ScheduledFor) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Milestones) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledFor) {
    return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *InlineResponse20026Milestones) HasScheduledFor() bool {
	if o != nil && !isNil(o.ScheduledFor) {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *InlineResponse20026Milestones) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *InlineResponse20026Milestones) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Milestones) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *InlineResponse20026Milestones) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *InlineResponse20026Milestones) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse20026Milestones) GetCompletedAt() time.Time {
	if o == nil || isNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Milestones) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse20026Milestones) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *InlineResponse20026Milestones) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *InlineResponse20026Milestones) GetCanceledAt() time.Time {
	if o == nil || isNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Milestones) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CanceledAt) {
    return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *InlineResponse20026Milestones) HasCanceledAt() bool {
	if o != nil && !isNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *InlineResponse20026Milestones) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

func (o InlineResponse20026Milestones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScheduledFor) {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	if !isNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.CanceledAt) {
		toSerialize["canceledAt"] = o.CanceledAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026Milestones struct {
	value *InlineResponse20026Milestones
	isSet bool
}

func (v NullableInlineResponse20026Milestones) Get() *InlineResponse20026Milestones {
	return v.value
}

func (v *NullableInlineResponse20026Milestones) Set(val *InlineResponse20026Milestones) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Milestones) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Milestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Milestones(val *InlineResponse20026Milestones) *NullableInlineResponse20026Milestones {
	return &NullableInlineResponse20026Milestones{value: val, isSet: true}
}

func (v NullableInlineResponse20026Milestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Milestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


