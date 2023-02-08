/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkApplianceSettingsRequest struct for UpdateNetworkApplianceSettingsRequest
type UpdateNetworkApplianceSettingsRequest struct {
	// Client tracking method of a network
	ClientTrackingMethod *string `json:"clientTrackingMethod,omitempty"`
	// Deployment mode of a network
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	DynamicDns *UpdateNetworkApplianceSettingsRequestDynamicDns `json:"dynamicDns,omitempty"`
}

// NewUpdateNetworkApplianceSettingsRequest instantiates a new UpdateNetworkApplianceSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSettingsRequest() *UpdateNetworkApplianceSettingsRequest {
	this := UpdateNetworkApplianceSettingsRequest{}
	return &this
}

// NewUpdateNetworkApplianceSettingsRequestWithDefaults instantiates a new UpdateNetworkApplianceSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSettingsRequestWithDefaults() *UpdateNetworkApplianceSettingsRequest {
	this := UpdateNetworkApplianceSettingsRequest{}
	return &this
}

// GetClientTrackingMethod returns the ClientTrackingMethod field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSettingsRequest) GetClientTrackingMethod() string {
	if o == nil || isNil(o.ClientTrackingMethod) {
		var ret string
		return ret
	}
	return *o.ClientTrackingMethod
}

// GetClientTrackingMethodOk returns a tuple with the ClientTrackingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSettingsRequest) GetClientTrackingMethodOk() (*string, bool) {
	if o == nil || isNil(o.ClientTrackingMethod) {
    return nil, false
	}
	return o.ClientTrackingMethod, true
}

// HasClientTrackingMethod returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSettingsRequest) HasClientTrackingMethod() bool {
	if o != nil && !isNil(o.ClientTrackingMethod) {
		return true
	}

	return false
}

// SetClientTrackingMethod gets a reference to the given string and assigns it to the ClientTrackingMethod field.
func (o *UpdateNetworkApplianceSettingsRequest) SetClientTrackingMethod(v string) {
	o.ClientTrackingMethod = &v
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSettingsRequest) GetDeploymentMode() string {
	if o == nil || isNil(o.DeploymentMode) {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSettingsRequest) GetDeploymentModeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentMode) {
    return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSettingsRequest) HasDeploymentMode() bool {
	if o != nil && !isNil(o.DeploymentMode) {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *UpdateNetworkApplianceSettingsRequest) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetDynamicDns returns the DynamicDns field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSettingsRequest) GetDynamicDns() UpdateNetworkApplianceSettingsRequestDynamicDns {
	if o == nil || isNil(o.DynamicDns) {
		var ret UpdateNetworkApplianceSettingsRequestDynamicDns
		return ret
	}
	return *o.DynamicDns
}

// GetDynamicDnsOk returns a tuple with the DynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSettingsRequest) GetDynamicDnsOk() (*UpdateNetworkApplianceSettingsRequestDynamicDns, bool) {
	if o == nil || isNil(o.DynamicDns) {
    return nil, false
	}
	return o.DynamicDns, true
}

// HasDynamicDns returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSettingsRequest) HasDynamicDns() bool {
	if o != nil && !isNil(o.DynamicDns) {
		return true
	}

	return false
}

// SetDynamicDns gets a reference to the given UpdateNetworkApplianceSettingsRequestDynamicDns and assigns it to the DynamicDns field.
func (o *UpdateNetworkApplianceSettingsRequest) SetDynamicDns(v UpdateNetworkApplianceSettingsRequestDynamicDns) {
	o.DynamicDns = &v
}

func (o UpdateNetworkApplianceSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientTrackingMethod) {
		toSerialize["clientTrackingMethod"] = o.ClientTrackingMethod
	}
	if !isNil(o.DeploymentMode) {
		toSerialize["deploymentMode"] = o.DeploymentMode
	}
	if !isNil(o.DynamicDns) {
		toSerialize["dynamicDns"] = o.DynamicDns
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSettingsRequest struct {
	value *UpdateNetworkApplianceSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceSettingsRequest) Get() *UpdateNetworkApplianceSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSettingsRequest) Set(val *UpdateNetworkApplianceSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSettingsRequest(val *UpdateNetworkApplianceSettingsRequest) *NullableUpdateNetworkApplianceSettingsRequest {
	return &NullableUpdateNetworkApplianceSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


