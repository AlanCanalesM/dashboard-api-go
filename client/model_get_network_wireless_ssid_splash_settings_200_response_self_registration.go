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

// checks if the GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration{}

// GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration Self-registration for splash with Meraki authentication.
type GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration struct {
	// Whether or not to allow users to create their own account on the network.
	Enabled *bool `json:"enabled,omitempty"`
	// How created user accounts should be authorized.
	AuthorizationType *string `json:"authorizationType,omitempty"`
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration() *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration{}
	return &this
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistrationWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistrationWithDefaults() *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthorizationType returns the AuthorizationType field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) GetAuthorizationType() string {
	if o == nil || IsNil(o.AuthorizationType) {
		var ret string
		return ret
	}
	return *o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) GetAuthorizationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationType) {
		return nil, false
	}
	return o.AuthorizationType, true
}

// HasAuthorizationType returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) HasAuthorizationType() bool {
	if o != nil && !IsNil(o.AuthorizationType) {
		return true
	}

	return false
}

// SetAuthorizationType gets a reference to the given string and assigns it to the AuthorizationType field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) SetAuthorizationType(v string) {
	o.AuthorizationType = &v
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AuthorizationType) {
		toSerialize["authorizationType"] = o.AuthorizationType
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration struct {
	value *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration
	isSet bool
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) Get() *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) Set(val *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration(val *GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) *NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration {
	return &NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


