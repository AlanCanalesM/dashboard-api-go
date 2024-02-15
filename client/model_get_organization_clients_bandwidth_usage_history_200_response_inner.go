/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationClientsBandwidthUsageHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationClientsBandwidthUsageHistory200ResponseInner{}

// GetOrganizationClientsBandwidthUsageHistory200ResponseInner struct for GetOrganizationClientsBandwidthUsageHistory200ResponseInner
type GetOrganizationClientsBandwidthUsageHistory200ResponseInner struct {
	// Timestamp for the bandwidth usage snapshot.
	Ts *time.Time `json:"ts,omitempty"`
	// Total bandwidth usage, in mbps.
	Total *int32 `json:"total,omitempty"`
	// Uploaded data, in mbps.
	Upstream *int32 `json:"upstream,omitempty"`
	// Downloaded data, in mbps.
	Downstream *int32 `json:"downstream,omitempty"`
}

// NewGetOrganizationClientsBandwidthUsageHistory200ResponseInner instantiates a new GetOrganizationClientsBandwidthUsageHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationClientsBandwidthUsageHistory200ResponseInner() *GetOrganizationClientsBandwidthUsageHistory200ResponseInner {
	this := GetOrganizationClientsBandwidthUsageHistory200ResponseInner{}
	return &this
}

// NewGetOrganizationClientsBandwidthUsageHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationClientsBandwidthUsageHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationClientsBandwidthUsageHistory200ResponseInnerWithDefaults() *GetOrganizationClientsBandwidthUsageHistory200ResponseInner {
	this := GetOrganizationClientsBandwidthUsageHistory200ResponseInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) SetTotal(v int32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetUpstream() int32 {
	if o == nil || IsNil(o.Upstream) {
		var ret int32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetUpstreamOk() (*int32, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given int32 and assigns it to the Upstream field.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) SetUpstream(v int32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetDownstream() int32 {
	if o == nil || IsNil(o.Downstream) {
		var ret int32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) GetDownstreamOk() (*int32, bool) {
	if o == nil || IsNil(o.Downstream) {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) HasDownstream() bool {
	if o != nil && !IsNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given int32 and assigns it to the Downstream field.
func (o *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) SetDownstream(v int32) {
	o.Downstream = &v
}

func (o GetOrganizationClientsBandwidthUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationClientsBandwidthUsageHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	return toSerialize, nil
}

type NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner struct {
	value *GetOrganizationClientsBandwidthUsageHistory200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) Get() *GetOrganizationClientsBandwidthUsageHistory200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) Set(val *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner(val *GetOrganizationClientsBandwidthUsageHistory200ResponseInner) *NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner {
	return &NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationClientsBandwidthUsageHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


