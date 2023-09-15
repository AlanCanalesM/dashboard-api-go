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

// checks if the UpdateOrganizationApplianceSecurityIntrusionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationApplianceSecurityIntrusionRequest{}

// UpdateOrganizationApplianceSecurityIntrusionRequest struct for UpdateOrganizationApplianceSecurityIntrusionRequest
type UpdateOrganizationApplianceSecurityIntrusionRequest struct {
	// Sets a list of specific SNORT signatures to allow
	AllowedRules []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner `json:"allowedRules"`
}

// NewUpdateOrganizationApplianceSecurityIntrusionRequest instantiates a new UpdateOrganizationApplianceSecurityIntrusionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationApplianceSecurityIntrusionRequest(allowedRules []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner) *UpdateOrganizationApplianceSecurityIntrusionRequest {
	this := UpdateOrganizationApplianceSecurityIntrusionRequest{}
	this.AllowedRules = allowedRules
	return &this
}

// NewUpdateOrganizationApplianceSecurityIntrusionRequestWithDefaults instantiates a new UpdateOrganizationApplianceSecurityIntrusionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationApplianceSecurityIntrusionRequestWithDefaults() *UpdateOrganizationApplianceSecurityIntrusionRequest {
	this := UpdateOrganizationApplianceSecurityIntrusionRequest{}
	return &this
}

// GetAllowedRules returns the AllowedRules field value
func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) GetAllowedRules() []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner {
	if o == nil {
		var ret []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner
		return ret
	}

	return o.AllowedRules
}

// GetAllowedRulesOk returns a tuple with the AllowedRules field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) GetAllowedRulesOk() ([]UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedRules, true
}

// SetAllowedRules sets field value
func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) SetAllowedRules(v []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner) {
	o.AllowedRules = v
}

func (o UpdateOrganizationApplianceSecurityIntrusionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationApplianceSecurityIntrusionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowedRules"] = o.AllowedRules
	return toSerialize, nil
}

type NullableUpdateOrganizationApplianceSecurityIntrusionRequest struct {
	value *UpdateOrganizationApplianceSecurityIntrusionRequest
	isSet bool
}

func (v NullableUpdateOrganizationApplianceSecurityIntrusionRequest) Get() *UpdateOrganizationApplianceSecurityIntrusionRequest {
	return v.value
}

func (v *NullableUpdateOrganizationApplianceSecurityIntrusionRequest) Set(val *UpdateOrganizationApplianceSecurityIntrusionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationApplianceSecurityIntrusionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationApplianceSecurityIntrusionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationApplianceSecurityIntrusionRequest(val *UpdateOrganizationApplianceSecurityIntrusionRequest) *NullableUpdateOrganizationApplianceSecurityIntrusionRequest {
	return &NullableUpdateOrganizationApplianceSecurityIntrusionRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationApplianceSecurityIntrusionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationApplianceSecurityIntrusionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


