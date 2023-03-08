/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject177 struct for InlineObject177
type InlineObject177 struct {
	// The email of the dashboard administrator. This attribute can not be updated.
	Email string `json:"email"`
	// The name of the dashboard administrator
	Name string `json:"name"`
	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the dashboard administrator has privileges on
	Tags []OrganizationsOrganizationIdAdminsTags `json:"tags,omitempty"`
	// The list of networks that the dashboard administrator has privileges on
	Networks []OrganizationsOrganizationIdAdminsNetworks `json:"networks,omitempty"`
	// The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of 'Email' or 'Cisco SecureX Sign-On'. The default is Email authentication
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewInlineObject177 instantiates a new InlineObject177 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject177(email string, name string, orgAccess string) *InlineObject177 {
	this := InlineObject177{}
	this.Email = email
	this.Name = name
	this.OrgAccess = orgAccess
	return &this
}

// NewInlineObject177WithDefaults instantiates a new InlineObject177 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject177WithDefaults() *InlineObject177 {
	this := InlineObject177{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineObject177) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineObject177) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *InlineObject177) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject177) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *InlineObject177) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetOrgAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *InlineObject177) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject177) GetTags() []OrganizationsOrganizationIdAdminsTags {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdAdminsTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetTagsOk() ([]OrganizationsOrganizationIdAdminsTags, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject177) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdAdminsTags and assigns it to the Tags field.
func (o *InlineObject177) SetTags(v []OrganizationsOrganizationIdAdminsTags) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject177) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdAdminsNetworks
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetNetworksOk() ([]OrganizationsOrganizationIdAdminsNetworks, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject177) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdAdminsNetworks and assigns it to the Networks field.
func (o *InlineObject177) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *InlineObject177) GetAuthenticationMethod() string {
	if o == nil || isNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationMethod) {
    return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *InlineObject177) HasAuthenticationMethod() bool {
	if o != nil && !isNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *InlineObject177) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o InlineObject177) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !isNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject177 struct {
	value *InlineObject177
	isSet bool
}

func (v NullableInlineObject177) Get() *InlineObject177 {
	return v.value
}

func (v *NullableInlineObject177) Set(val *InlineObject177) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject177) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject177) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject177(val *InlineObject177) *NullableInlineObject177 {
	return &NullableInlineObject177{value: val, isSet: true}
}

func (v NullableInlineObject177) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject177) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


