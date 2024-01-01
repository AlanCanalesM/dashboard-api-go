/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage{}

// UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage Properties for setting a new image.
type UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage struct {
	// The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
	// The file contents (a base 64 encoded string) of your new logo.
	Contents *string `json:"contents,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImageWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImageWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) SetFormat(v string) {
	o.Format = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) SetContents(v string) {
	o.Contents = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


