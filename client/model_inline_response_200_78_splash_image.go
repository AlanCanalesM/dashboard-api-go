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

// InlineResponse20078SplashImage The image used in the splash page.
type InlineResponse20078SplashImage struct {
	// The MD5 value of the image file.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the image file.
	Extension *string `json:"extension,omitempty"`
}

// NewInlineResponse20078SplashImage instantiates a new InlineResponse20078SplashImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20078SplashImage() *InlineResponse20078SplashImage {
	this := InlineResponse20078SplashImage{}
	return &this
}

// NewInlineResponse20078SplashImageWithDefaults instantiates a new InlineResponse20078SplashImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20078SplashImageWithDefaults() *InlineResponse20078SplashImage {
	this := InlineResponse20078SplashImage{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *InlineResponse20078SplashImage) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20078SplashImage) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *InlineResponse20078SplashImage) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *InlineResponse20078SplashImage) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *InlineResponse20078SplashImage) GetExtension() string {
	if o == nil || isNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20078SplashImage) GetExtensionOk() (*string, bool) {
	if o == nil || isNil(o.Extension) {
    return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *InlineResponse20078SplashImage) HasExtension() bool {
	if o != nil && !isNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *InlineResponse20078SplashImage) SetExtension(v string) {
	o.Extension = &v
}

func (o InlineResponse20078SplashImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !isNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20078SplashImage struct {
	value *InlineResponse20078SplashImage
	isSet bool
}

func (v NullableInlineResponse20078SplashImage) Get() *InlineResponse20078SplashImage {
	return v.value
}

func (v *NullableInlineResponse20078SplashImage) Set(val *InlineResponse20078SplashImage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20078SplashImage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20078SplashImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20078SplashImage(val *InlineResponse20078SplashImage) *NullableInlineResponse20078SplashImage {
	return &NullableInlineResponse20078SplashImage{value: val, isSet: true}
}

func (v NullableInlineResponse20078SplashImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20078SplashImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


