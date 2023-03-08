/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20076 struct for InlineResponse20076
type InlineResponse20076 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Average data rate in kilobytes-per-second
	AverageKbps *int32 `json:"averageKbps,omitempty"`
	// Download rate in kilobytes-per-second
	DownloadKbps *int32 `json:"downloadKbps,omitempty"`
	// Upload rate in kilobytes-per-second
	UploadKbps *int32 `json:"uploadKbps,omitempty"`
}

// NewInlineResponse20076 instantiates a new InlineResponse20076 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20076() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20076WithDefaults() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20076) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20076) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20076) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20076) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20076) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20076) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAverageKbps returns the AverageKbps field value if set, zero value otherwise.
func (o *InlineResponse20076) GetAverageKbps() int32 {
	if o == nil || isNil(o.AverageKbps) {
		var ret int32
		return ret
	}
	return *o.AverageKbps
}

// GetAverageKbpsOk returns a tuple with the AverageKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetAverageKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.AverageKbps) {
    return nil, false
	}
	return o.AverageKbps, true
}

// HasAverageKbps returns a boolean if a field has been set.
func (o *InlineResponse20076) HasAverageKbps() bool {
	if o != nil && !isNil(o.AverageKbps) {
		return true
	}

	return false
}

// SetAverageKbps gets a reference to the given int32 and assigns it to the AverageKbps field.
func (o *InlineResponse20076) SetAverageKbps(v int32) {
	o.AverageKbps = &v
}

// GetDownloadKbps returns the DownloadKbps field value if set, zero value otherwise.
func (o *InlineResponse20076) GetDownloadKbps() int32 {
	if o == nil || isNil(o.DownloadKbps) {
		var ret int32
		return ret
	}
	return *o.DownloadKbps
}

// GetDownloadKbpsOk returns a tuple with the DownloadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetDownloadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.DownloadKbps) {
    return nil, false
	}
	return o.DownloadKbps, true
}

// HasDownloadKbps returns a boolean if a field has been set.
func (o *InlineResponse20076) HasDownloadKbps() bool {
	if o != nil && !isNil(o.DownloadKbps) {
		return true
	}

	return false
}

// SetDownloadKbps gets a reference to the given int32 and assigns it to the DownloadKbps field.
func (o *InlineResponse20076) SetDownloadKbps(v int32) {
	o.DownloadKbps = &v
}

// GetUploadKbps returns the UploadKbps field value if set, zero value otherwise.
func (o *InlineResponse20076) GetUploadKbps() int32 {
	if o == nil || isNil(o.UploadKbps) {
		var ret int32
		return ret
	}
	return *o.UploadKbps
}

// GetUploadKbpsOk returns a tuple with the UploadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetUploadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.UploadKbps) {
    return nil, false
	}
	return o.UploadKbps, true
}

// HasUploadKbps returns a boolean if a field has been set.
func (o *InlineResponse20076) HasUploadKbps() bool {
	if o != nil && !isNil(o.UploadKbps) {
		return true
	}

	return false
}

// SetUploadKbps gets a reference to the given int32 and assigns it to the UploadKbps field.
func (o *InlineResponse20076) SetUploadKbps(v int32) {
	o.UploadKbps = &v
}

func (o InlineResponse20076) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.AverageKbps) {
		toSerialize["averageKbps"] = o.AverageKbps
	}
	if !isNil(o.DownloadKbps) {
		toSerialize["downloadKbps"] = o.DownloadKbps
	}
	if !isNil(o.UploadKbps) {
		toSerialize["uploadKbps"] = o.UploadKbps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20076 struct {
	value *InlineResponse20076
	isSet bool
}

func (v NullableInlineResponse20076) Get() *InlineResponse20076 {
	return v.value
}

func (v *NullableInlineResponse20076) Set(val *InlineResponse20076) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20076) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20076) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20076(val *InlineResponse20076) *NullableInlineResponse20076 {
	return &NullableInlineResponse20076{value: val, isSet: true}
}

func (v NullableInlineResponse20076) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20076) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


