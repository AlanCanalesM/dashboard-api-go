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

// checks if the UpdateOrganizationAdaptivePolicyAclRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAdaptivePolicyAclRequest{}

// UpdateOrganizationAdaptivePolicyAclRequest struct for UpdateOrganizationAdaptivePolicyAclRequest
type UpdateOrganizationAdaptivePolicyAclRequest struct {
	// Name of the adaptive policy ACL
	Name *string `json:"name,omitempty"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules. An empty array will clear the rules.
	Rules []CreateOrganizationAdaptivePolicyAclRequestRulesInner `json:"rules,omitempty"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion *string `json:"ipVersion,omitempty"`
}

// NewUpdateOrganizationAdaptivePolicyAclRequest instantiates a new UpdateOrganizationAdaptivePolicyAclRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAdaptivePolicyAclRequest() *UpdateOrganizationAdaptivePolicyAclRequest {
	this := UpdateOrganizationAdaptivePolicyAclRequest{}
	return &this
}

// NewUpdateOrganizationAdaptivePolicyAclRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyAclRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAdaptivePolicyAclRequestWithDefaults() *UpdateOrganizationAdaptivePolicyAclRequest {
	this := UpdateOrganizationAdaptivePolicyAclRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetRules() []CreateOrganizationAdaptivePolicyAclRequestRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []CreateOrganizationAdaptivePolicyAclRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetRulesOk() ([]CreateOrganizationAdaptivePolicyAclRequestRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []CreateOrganizationAdaptivePolicyAclRequestRulesInner and assigns it to the Rules field.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetRules(v []CreateOrganizationAdaptivePolicyAclRequestRulesInner) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetIpVersion() string {
	if o == nil || IsNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetIpVersionOk() (*string, bool) {
	if o == nil || IsNil(o.IpVersion) {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasIpVersion() bool {
	if o != nil && !IsNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetIpVersion(v string) {
	o.IpVersion = &v
}

func (o UpdateOrganizationAdaptivePolicyAclRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAdaptivePolicyAclRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAdaptivePolicyAclRequest struct {
	value *UpdateOrganizationAdaptivePolicyAclRequest
	isSet bool
}

func (v NullableUpdateOrganizationAdaptivePolicyAclRequest) Get() *UpdateOrganizationAdaptivePolicyAclRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAdaptivePolicyAclRequest) Set(val *UpdateOrganizationAdaptivePolicyAclRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAdaptivePolicyAclRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAdaptivePolicyAclRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAdaptivePolicyAclRequest(val *UpdateOrganizationAdaptivePolicyAclRequest) *NullableUpdateOrganizationAdaptivePolicyAclRequest {
	return &NullableUpdateOrganizationAdaptivePolicyAclRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAdaptivePolicyAclRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAdaptivePolicyAclRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


