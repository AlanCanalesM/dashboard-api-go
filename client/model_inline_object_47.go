/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject47 struct for InlineObject47
type InlineObject47 struct {
	// Set mode to 'enabled' to enable malware prevention, otherwise 'disabled'
	Mode string `json:"mode"`
	// The urls that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedUrls []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls `json:"allowedUrls,omitempty"`
	// The sha256 digests of files that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedFiles []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"`
}

// NewInlineObject47 instantiates a new InlineObject47 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject47(mode string) *InlineObject47 {
	this := InlineObject47{}
	this.Mode = mode
	return &this
}

// NewInlineObject47WithDefaults instantiates a new InlineObject47 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject47WithDefaults() *InlineObject47 {
	this := InlineObject47{}
	return &this
}

// GetMode returns the Mode field value
func (o *InlineObject47) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InlineObject47) SetMode(v string) {
	o.Mode = v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise.
func (o *InlineObject47) GetAllowedUrls() []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls {
	if o == nil || isNil(o.AllowedUrls) {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls
		return ret
	}
	return o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetAllowedUrlsOk() ([]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls, bool) {
	if o == nil || isNil(o.AllowedUrls) {
    return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *InlineObject47) HasAllowedUrls() bool {
	if o != nil && !isNil(o.AllowedUrls) {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls and assigns it to the AllowedUrls field.
func (o *InlineObject47) SetAllowedUrls(v []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls) {
	o.AllowedUrls = v
}

// GetAllowedFiles returns the AllowedFiles field value if set, zero value otherwise.
func (o *InlineObject47) GetAllowedFiles() []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles {
	if o == nil || isNil(o.AllowedFiles) {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles
		return ret
	}
	return o.AllowedFiles
}

// GetAllowedFilesOk returns a tuple with the AllowedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetAllowedFilesOk() ([]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles, bool) {
	if o == nil || isNil(o.AllowedFiles) {
    return nil, false
	}
	return o.AllowedFiles, true
}

// HasAllowedFiles returns a boolean if a field has been set.
func (o *InlineObject47) HasAllowedFiles() bool {
	if o != nil && !isNil(o.AllowedFiles) {
		return true
	}

	return false
}

// SetAllowedFiles gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles and assigns it to the AllowedFiles field.
func (o *InlineObject47) SetAllowedFiles(v []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles) {
	o.AllowedFiles = v
}

func (o InlineObject47) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.AllowedUrls) {
		toSerialize["allowedUrls"] = o.AllowedUrls
	}
	if !isNil(o.AllowedFiles) {
		toSerialize["allowedFiles"] = o.AllowedFiles
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject47 struct {
	value *InlineObject47
	isSet bool
}

func (v NullableInlineObject47) Get() *InlineObject47 {
	return v.value
}

func (v *NullableInlineObject47) Set(val *InlineObject47) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject47) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject47) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject47(val *InlineObject47) *NullableInlineObject47 {
	return &NullableInlineObject47{value: val, isSet: true}
}

func (v NullableInlineObject47) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject47) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


