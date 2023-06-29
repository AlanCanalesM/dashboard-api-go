/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationBrandingPolicies200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationBrandingPolicies200ResponseInner{}

// GetOrganizationBrandingPolicies200ResponseInner struct for GetOrganizationBrandingPolicies200ResponseInner
type GetOrganizationBrandingPolicies200ResponseInner struct {
	// Name of the Dashboard branding policy.
	Name *string `json:"name,omitempty"`
	// Boolean indicating whether this policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AdminSettings *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings `json:"adminSettings,omitempty"`
	HelpSettings *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings `json:"helpSettings,omitempty"`
	CustomLogo *GetOrganizationBrandingPolicies200ResponseInnerCustomLogo `json:"customLogo,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInner instantiates a new GetOrganizationBrandingPolicies200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInner() *GetOrganizationBrandingPolicies200ResponseInner {
	this := GetOrganizationBrandingPolicies200ResponseInner{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerWithDefaults() *GetOrganizationBrandingPolicies200ResponseInner {
	this := GetOrganizationBrandingPolicies200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationBrandingPolicies200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetOrganizationBrandingPolicies200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetAdminSettings() GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	if o == nil || IsNil(o.AdminSettings) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetAdminSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerAdminSettings, bool) {
	if o == nil || IsNil(o.AdminSettings) {
		return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) HasAdminSettings() bool {
	if o != nil && !IsNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerAdminSettings and assigns it to the AdminSettings field.
func (o *GetOrganizationBrandingPolicies200ResponseInner) SetAdminSettings(v GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) {
	o.AdminSettings = &v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetHelpSettings() GetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	if o == nil || IsNil(o.HelpSettings) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerHelpSettings
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetHelpSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerHelpSettings, bool) {
	if o == nil || IsNil(o.HelpSettings) {
		return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) HasHelpSettings() bool {
	if o != nil && !IsNil(o.HelpSettings) {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerHelpSettings and assigns it to the HelpSettings field.
func (o *GetOrganizationBrandingPolicies200ResponseInner) SetHelpSettings(v GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetCustomLogo() GetOrganizationBrandingPolicies200ResponseInnerCustomLogo {
	if o == nil || IsNil(o.CustomLogo) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerCustomLogo
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) GetCustomLogoOk() (*GetOrganizationBrandingPolicies200ResponseInnerCustomLogo, bool) {
	if o == nil || IsNil(o.CustomLogo) {
		return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInner) HasCustomLogo() bool {
	if o != nil && !IsNil(o.CustomLogo) {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerCustomLogo and assigns it to the CustomLogo field.
func (o *GetOrganizationBrandingPolicies200ResponseInner) SetCustomLogo(v GetOrganizationBrandingPolicies200ResponseInnerCustomLogo) {
	o.CustomLogo = &v
}

func (o GetOrganizationBrandingPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationBrandingPolicies200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AdminSettings) {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if !IsNil(o.HelpSettings) {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if !IsNil(o.CustomLogo) {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return toSerialize, nil
}

type NullableGetOrganizationBrandingPolicies200ResponseInner struct {
	value *GetOrganizationBrandingPolicies200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInner) Get() *GetOrganizationBrandingPolicies200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInner) Set(val *GetOrganizationBrandingPolicies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInner(val *GetOrganizationBrandingPolicies200ResponseInner) *NullableGetOrganizationBrandingPolicies200ResponseInner {
	return &NullableGetOrganizationBrandingPolicies200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


