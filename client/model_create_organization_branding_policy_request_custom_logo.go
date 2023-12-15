/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationBrandingPolicyRequestCustomLogo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationBrandingPolicyRequestCustomLogo{}

// CreateOrganizationBrandingPolicyRequestCustomLogo Properties describing the custom logo attached to the branding policy.
type CreateOrganizationBrandingPolicyRequestCustomLogo struct {
	// Whether or not there is a custom logo enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Image *CreateOrganizationBrandingPolicyRequestCustomLogoImage `json:"image,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogo instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequestCustomLogo() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	this := CreateOrganizationBrandingPolicyRequestCustomLogo{}
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	this := CreateOrganizationBrandingPolicyRequestCustomLogo{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImage() CreateOrganizationBrandingPolicyRequestCustomLogoImage {
	if o == nil || IsNil(o.Image) {
		var ret CreateOrganizationBrandingPolicyRequestCustomLogoImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImageOk() (*CreateOrganizationBrandingPolicyRequestCustomLogoImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given CreateOrganizationBrandingPolicyRequestCustomLogoImage and assigns it to the Image field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetImage(v CreateOrganizationBrandingPolicyRequestCustomLogoImage) {
	o.Image = &v
}

func (o CreateOrganizationBrandingPolicyRequestCustomLogo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationBrandingPolicyRequestCustomLogo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableCreateOrganizationBrandingPolicyRequestCustomLogo struct {
	value *CreateOrganizationBrandingPolicyRequestCustomLogo
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Get() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Set(val *CreateOrganizationBrandingPolicyRequestCustomLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequestCustomLogo(val *CreateOrganizationBrandingPolicyRequestCustomLogo) *NullableCreateOrganizationBrandingPolicyRequestCustomLogo {
	return &NullableCreateOrganizationBrandingPolicyRequestCustomLogo{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


