/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 Quality and resolution for MV52 camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52WithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52WithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52 {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV52) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


