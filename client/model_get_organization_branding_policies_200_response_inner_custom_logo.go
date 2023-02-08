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

// GetOrganizationBrandingPolicies200ResponseInnerCustomLogo Properties describing the custom logo attached to the branding policy.
type GetOrganizationBrandingPolicies200ResponseInnerCustomLogo struct {
	// Whether or not there is a custom logo enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Image *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage `json:"image,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogo instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogo() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogo{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogo{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) GetImage() GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage {
	if o == nil || isNil(o.Image) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) GetImageOk() (*GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage and assigns it to the Image field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) SetImage(v GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) {
	o.Image = &v
}

func (o GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo struct {
	value *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) Get() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) Set(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo {
	return &NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


