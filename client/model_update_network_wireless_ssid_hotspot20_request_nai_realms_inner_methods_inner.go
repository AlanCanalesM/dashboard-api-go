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

// UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner struct for UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner
type UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner struct {
	// ID of method
	Id *string `json:"id,omitempty"`
	// The authentication types for the method. These should be formatted as an object with the EAP method category in camelcase as the key and the list of types as the value (nonEapInnerAuthentication: Reserved, PAP, CHAP, MSCHAP, MSCHAPV2; eapInnerAuthentication: EAP-TLS, EAP-SIM, EAP-AKA, EAP-TTLS with MSCHAPv2; credentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, none, Reserved, Vendor Specific; tunneledEapMethodCredentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, Reserved, Anonymous, Vendor Specific)
	AuthenticationTypes map[string]interface{} `json:"authenticationTypes,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInnerWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) SetId(v string) {
	o.Id = &v
}

// GetAuthenticationTypes returns the AuthenticationTypes field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetAuthenticationTypes() map[string]interface{} {
	if o == nil || isNil(o.AuthenticationTypes) {
		var ret map[string]interface{}
		return ret
	}
	return o.AuthenticationTypes
}

// GetAuthenticationTypesOk returns a tuple with the AuthenticationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetAuthenticationTypesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AuthenticationTypes) {
    return map[string]interface{}{}, false
	}
	return o.AuthenticationTypes, true
}

// HasAuthenticationTypes returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) HasAuthenticationTypes() bool {
	if o != nil && !isNil(o.AuthenticationTypes) {
		return true
	}

	return false
}

// SetAuthenticationTypes gets a reference to the given map[string]interface{} and assigns it to the AuthenticationTypes field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) SetAuthenticationTypes(v map[string]interface{}) {
	o.AuthenticationTypes = v
}

func (o UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AuthenticationTypes) {
		toSerialize["authenticationTypes"] = o.AuthenticationTypes
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner struct {
	value *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) Get() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) Set(val *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner(val *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner {
	return &NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


