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

// checks if the GetNetworkApplianceSettings200ResponseDynamicDns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSettings200ResponseDynamicDns{}

// GetNetworkApplianceSettings200ResponseDynamicDns Dynamic DNS settings for a network
type GetNetworkApplianceSettings200ResponseDynamicDns struct {
	// Dynamic DNS enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Dynamic DNS url prefix. DDNS must be enabled to update
	Prefix *string `json:"prefix,omitempty"`
	// Dynamic DNS url. DDNS must be enabled to update
	Url *string `json:"url,omitempty"`
}

// NewGetNetworkApplianceSettings200ResponseDynamicDns instantiates a new GetNetworkApplianceSettings200ResponseDynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSettings200ResponseDynamicDns() *GetNetworkApplianceSettings200ResponseDynamicDns {
	this := GetNetworkApplianceSettings200ResponseDynamicDns{}
	return &this
}

// NewGetNetworkApplianceSettings200ResponseDynamicDnsWithDefaults instantiates a new GetNetworkApplianceSettings200ResponseDynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSettings200ResponseDynamicDnsWithDefaults() *GetNetworkApplianceSettings200ResponseDynamicDns {
	this := GetNetworkApplianceSettings200ResponseDynamicDns{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkApplianceSettings200ResponseDynamicDns) SetUrl(v string) {
	o.Url = &v
}

func (o GetNetworkApplianceSettings200ResponseDynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSettings200ResponseDynamicDns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSettings200ResponseDynamicDns struct {
	value *GetNetworkApplianceSettings200ResponseDynamicDns
	isSet bool
}

func (v NullableGetNetworkApplianceSettings200ResponseDynamicDns) Get() *GetNetworkApplianceSettings200ResponseDynamicDns {
	return v.value
}

func (v *NullableGetNetworkApplianceSettings200ResponseDynamicDns) Set(val *GetNetworkApplianceSettings200ResponseDynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSettings200ResponseDynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSettings200ResponseDynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSettings200ResponseDynamicDns(val *GetNetworkApplianceSettings200ResponseDynamicDns) *NullableGetNetworkApplianceSettings200ResponseDynamicDns {
	return &NullableGetNetworkApplianceSettings200ResponseDynamicDns{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSettings200ResponseDynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSettings200ResponseDynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


