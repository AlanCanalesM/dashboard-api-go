/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateOrganizationBrandingPolicyRequest struct for UpdateOrganizationBrandingPolicyRequest
type UpdateOrganizationBrandingPolicyRequest struct {
	// Name of the Dashboard branding policy.
	Name *string `json:"name,omitempty"`
	// Boolean indicating whether this policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AdminSettings *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings `json:"adminSettings,omitempty"`
	HelpSettings *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings `json:"helpSettings,omitempty"`
	CustomLogo *CreateOrganizationBrandingPolicyRequestCustomLogo `json:"customLogo,omitempty"`
}

// NewUpdateOrganizationBrandingPolicyRequest instantiates a new UpdateOrganizationBrandingPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationBrandingPolicyRequest() *UpdateOrganizationBrandingPolicyRequest {
	this := UpdateOrganizationBrandingPolicyRequest{}
	return &this
}

// NewUpdateOrganizationBrandingPolicyRequestWithDefaults instantiates a new UpdateOrganizationBrandingPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationBrandingPolicyRequestWithDefaults() *UpdateOrganizationBrandingPolicyRequest {
	this := UpdateOrganizationBrandingPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationBrandingPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateOrganizationBrandingPolicyRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequest) GetAdminSettings() GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	if o == nil || isNil(o.AdminSettings) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) GetAdminSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerAdminSettings, bool) {
	if o == nil || isNil(o.AdminSettings) {
    return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) HasAdminSettings() bool {
	if o != nil && !isNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerAdminSettings and assigns it to the AdminSettings field.
func (o *UpdateOrganizationBrandingPolicyRequest) SetAdminSettings(v GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) {
	o.AdminSettings = &v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequest) GetHelpSettings() GetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	if o == nil || isNil(o.HelpSettings) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerHelpSettings
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) GetHelpSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerHelpSettings, bool) {
	if o == nil || isNil(o.HelpSettings) {
    return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) HasHelpSettings() bool {
	if o != nil && !isNil(o.HelpSettings) {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerHelpSettings and assigns it to the HelpSettings field.
func (o *UpdateOrganizationBrandingPolicyRequest) SetHelpSettings(v GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *UpdateOrganizationBrandingPolicyRequest) GetCustomLogo() CreateOrganizationBrandingPolicyRequestCustomLogo {
	if o == nil || isNil(o.CustomLogo) {
		var ret CreateOrganizationBrandingPolicyRequestCustomLogo
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) GetCustomLogoOk() (*CreateOrganizationBrandingPolicyRequestCustomLogo, bool) {
	if o == nil || isNil(o.CustomLogo) {
    return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *UpdateOrganizationBrandingPolicyRequest) HasCustomLogo() bool {
	if o != nil && !isNil(o.CustomLogo) {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given CreateOrganizationBrandingPolicyRequestCustomLogo and assigns it to the CustomLogo field.
func (o *UpdateOrganizationBrandingPolicyRequest) SetCustomLogo(v CreateOrganizationBrandingPolicyRequestCustomLogo) {
	o.CustomLogo = &v
}

func (o UpdateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AdminSettings) {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if !isNil(o.HelpSettings) {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if !isNil(o.CustomLogo) {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationBrandingPolicyRequest struct {
	value *UpdateOrganizationBrandingPolicyRequest
	isSet bool
}

func (v NullableUpdateOrganizationBrandingPolicyRequest) Get() *UpdateOrganizationBrandingPolicyRequest {
	return v.value
}

func (v *NullableUpdateOrganizationBrandingPolicyRequest) Set(val *UpdateOrganizationBrandingPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationBrandingPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationBrandingPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationBrandingPolicyRequest(val *UpdateOrganizationBrandingPolicyRequest) *NullableUpdateOrganizationBrandingPolicyRequest {
	return &NullableUpdateOrganizationBrandingPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationBrandingPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


