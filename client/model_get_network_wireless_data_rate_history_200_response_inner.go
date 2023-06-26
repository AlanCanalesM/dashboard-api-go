/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetNetworkWirelessDataRateHistory200ResponseInner struct for GetNetworkWirelessDataRateHistory200ResponseInner
type GetNetworkWirelessDataRateHistory200ResponseInner struct {
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

// NewGetNetworkWirelessDataRateHistory200ResponseInner instantiates a new GetNetworkWirelessDataRateHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessDataRateHistory200ResponseInner() *GetNetworkWirelessDataRateHistory200ResponseInner {
	this := GetNetworkWirelessDataRateHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessDataRateHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessDataRateHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessDataRateHistory200ResponseInnerWithDefaults() *GetNetworkWirelessDataRateHistory200ResponseInner {
	this := GetNetworkWirelessDataRateHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAverageKbps returns the AverageKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetAverageKbps() int32 {
	if o == nil || isNil(o.AverageKbps) {
		var ret int32
		return ret
	}
	return *o.AverageKbps
}

// GetAverageKbpsOk returns a tuple with the AverageKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetAverageKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.AverageKbps) {
    return nil, false
	}
	return o.AverageKbps, true
}

// HasAverageKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) HasAverageKbps() bool {
	if o != nil && !isNil(o.AverageKbps) {
		return true
	}

	return false
}

// SetAverageKbps gets a reference to the given int32 and assigns it to the AverageKbps field.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) SetAverageKbps(v int32) {
	o.AverageKbps = &v
}

// GetDownloadKbps returns the DownloadKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetDownloadKbps() int32 {
	if o == nil || isNil(o.DownloadKbps) {
		var ret int32
		return ret
	}
	return *o.DownloadKbps
}

// GetDownloadKbpsOk returns a tuple with the DownloadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetDownloadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.DownloadKbps) {
    return nil, false
	}
	return o.DownloadKbps, true
}

// HasDownloadKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) HasDownloadKbps() bool {
	if o != nil && !isNil(o.DownloadKbps) {
		return true
	}

	return false
}

// SetDownloadKbps gets a reference to the given int32 and assigns it to the DownloadKbps field.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) SetDownloadKbps(v int32) {
	o.DownloadKbps = &v
}

// GetUploadKbps returns the UploadKbps field value if set, zero value otherwise.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetUploadKbps() int32 {
	if o == nil || isNil(o.UploadKbps) {
		var ret int32
		return ret
	}
	return *o.UploadKbps
}

// GetUploadKbpsOk returns a tuple with the UploadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) GetUploadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.UploadKbps) {
    return nil, false
	}
	return o.UploadKbps, true
}

// HasUploadKbps returns a boolean if a field has been set.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) HasUploadKbps() bool {
	if o != nil && !isNil(o.UploadKbps) {
		return true
	}

	return false
}

// SetUploadKbps gets a reference to the given int32 and assigns it to the UploadKbps field.
func (o *GetNetworkWirelessDataRateHistory200ResponseInner) SetUploadKbps(v int32) {
	o.UploadKbps = &v
}

func (o GetNetworkWirelessDataRateHistory200ResponseInner) MarshalJSON() ([]byte, error) {
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

type NullableGetNetworkWirelessDataRateHistory200ResponseInner struct {
	value *GetNetworkWirelessDataRateHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessDataRateHistory200ResponseInner) Get() *GetNetworkWirelessDataRateHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessDataRateHistory200ResponseInner) Set(val *GetNetworkWirelessDataRateHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessDataRateHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessDataRateHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessDataRateHistory200ResponseInner(val *GetNetworkWirelessDataRateHistory200ResponseInner) *NullableGetNetworkWirelessDataRateHistory200ResponseInner {
	return &NullableGetNetworkWirelessDataRateHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessDataRateHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessDataRateHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


