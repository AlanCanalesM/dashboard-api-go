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

// checks if the GetOrganizationBrandingPolicies200ResponseInnerAdminSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationBrandingPolicies200ResponseInnerAdminSettings{}

// GetOrganizationBrandingPolicies200ResponseInnerAdminSettings Settings for describing which kinds of admins this policy applies to.
type GetOrganizationBrandingPolicies200ResponseInnerAdminSettings struct {
	// Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	AppliesTo *string `json:"appliesTo,omitempty"`
	//       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names. 
	Values []string `json:"values,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettings instantiates a new GetOrganizationBrandingPolicies200ResponseInnerAdminSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettings() *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	this := GetOrganizationBrandingPolicies200ResponseInnerAdminSettings{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettingsWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerAdminSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettingsWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	this := GetOrganizationBrandingPolicies200ResponseInnerAdminSettings{}
	return &this
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetAppliesTo() string {
	if o == nil || IsNil(o.AppliesTo) {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetAppliesToOk() (*string, bool) {
	if o == nil || IsNil(o.AppliesTo) {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) HasAppliesTo() bool {
	if o != nil && !IsNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) SetValues(v []string) {
	o.Values = v
}

func (o GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppliesTo) {
		toSerialize["appliesTo"] = o.AppliesTo
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings struct {
	value *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) Get() *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) Set(val *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings(val *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) *NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	return &NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerAdminSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


