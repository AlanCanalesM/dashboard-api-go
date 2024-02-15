/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner{}

// DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner struct for DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner
type DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner struct {
	// Sequence ID of the packet
	SequenceId *int32 `json:"sequenceId,omitempty"`
	// Size of the packet in bytes
	Size *int32 `json:"size,omitempty"`
	// Latency of the packet in milliseconds
	Latency *float32 `json:"latency,omitempty"`
}

// NewDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner instantiates a new DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner() *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner {
	this := DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner{}
	return &this
}

// NewDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInnerWithDefaults instantiates a new DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInnerWithDefaults() *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner {
	this := DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner{}
	return &this
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetSequenceId() int32 {
	if o == nil || IsNil(o.SequenceId) {
		var ret int32
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetSequenceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceId) {
		return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) HasSequenceId() bool {
	if o != nil && !IsNil(o.SequenceId) {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given int32 and assigns it to the SequenceId field.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) SetSequenceId(v int32) {
	o.SequenceId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) SetSize(v int32) {
	o.Size = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetLatency() float32 {
	if o == nil || IsNil(o.Latency) {
		var ret float32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) GetLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) HasLatency() bool {
	if o != nil && !IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given float32 and assigns it to the Latency field.
func (o *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) SetLatency(v float32) {
	o.Latency = &v
}

func (o DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SequenceId) {
		toSerialize["sequenceId"] = o.SequenceId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	return toSerialize, nil
}

type NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner struct {
	value *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner
	isSet bool
}

func (v NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) Get() *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) Set(val *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner(val *DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) *NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner {
	return &NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


