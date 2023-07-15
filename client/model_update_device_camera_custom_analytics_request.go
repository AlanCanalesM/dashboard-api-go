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

// checks if the UpdateDeviceCameraCustomAnalyticsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCameraCustomAnalyticsRequest{}

// UpdateDeviceCameraCustomAnalyticsRequest struct for UpdateDeviceCameraCustomAnalyticsRequest
type UpdateDeviceCameraCustomAnalyticsRequest struct {
	// Enable custom analytics
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the custom analytics artifact
	ArtifactId *string `json:"artifactId,omitempty"`
	// Parameters for the custom analytics workload
	Parameters []UpdateDeviceCameraCustomAnalyticsRequestParametersInner `json:"parameters,omitempty"`
}

// NewUpdateDeviceCameraCustomAnalyticsRequest instantiates a new UpdateDeviceCameraCustomAnalyticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraCustomAnalyticsRequest() *UpdateDeviceCameraCustomAnalyticsRequest {
	this := UpdateDeviceCameraCustomAnalyticsRequest{}
	return &this
}

// NewUpdateDeviceCameraCustomAnalyticsRequestWithDefaults instantiates a new UpdateDeviceCameraCustomAnalyticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraCustomAnalyticsRequestWithDefaults() *UpdateDeviceCameraCustomAnalyticsRequest {
	this := UpdateDeviceCameraCustomAnalyticsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetArtifactId() string {
	if o == nil || IsNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetArtifactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactId) {
		return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasArtifactId() bool {
	if o != nil && !IsNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetParameters() []UpdateDeviceCameraCustomAnalyticsRequestParametersInner {
	if o == nil || IsNil(o.Parameters) {
		var ret []UpdateDeviceCameraCustomAnalyticsRequestParametersInner
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetParametersOk() ([]UpdateDeviceCameraCustomAnalyticsRequestParametersInner, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []UpdateDeviceCameraCustomAnalyticsRequestParametersInner and assigns it to the Parameters field.
func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetParameters(v []UpdateDeviceCameraCustomAnalyticsRequestParametersInner) {
	o.Parameters = v
}

func (o UpdateDeviceCameraCustomAnalyticsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCameraCustomAnalyticsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ArtifactId) {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCameraCustomAnalyticsRequest struct {
	value *UpdateDeviceCameraCustomAnalyticsRequest
	isSet bool
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequest) Get() *UpdateDeviceCameraCustomAnalyticsRequest {
	return v.value
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequest) Set(val *UpdateDeviceCameraCustomAnalyticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraCustomAnalyticsRequest(val *UpdateDeviceCameraCustomAnalyticsRequest) *NullableUpdateDeviceCameraCustomAnalyticsRequest {
	return &NullableUpdateDeviceCameraCustomAnalyticsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraCustomAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraCustomAnalyticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


