/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkWirelessSsidSplashSettings200ResponseBilling Details associated with billing splash
type GetNetworkWirelessSsidSplashSettings200ResponseBilling struct {
	FreeAccess *GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess `json:"freeAccess,omitempty"`
	// Whether or not billing uses the fast login prepaid access option.
	PrepaidAccessFastLoginEnabled *bool `json:"prepaidAccessFastLoginEnabled,omitempty"`
	// The email address that reeceives replies from clients
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseBilling instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidSplashSettings200ResponseBilling() *GetNetworkWirelessSsidSplashSettings200ResponseBilling {
	this := GetNetworkWirelessSsidSplashSettings200ResponseBilling{}
	return &this
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseBillingWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidSplashSettings200ResponseBillingWithDefaults() *GetNetworkWirelessSsidSplashSettings200ResponseBilling {
	this := GetNetworkWirelessSsidSplashSettings200ResponseBilling{}
	return &this
}

// GetFreeAccess returns the FreeAccess field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetFreeAccess() GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess {
	if o == nil || isNil(o.FreeAccess) {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess
		return ret
	}
	return *o.FreeAccess
}

// GetFreeAccessOk returns a tuple with the FreeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetFreeAccessOk() (*GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess, bool) {
	if o == nil || isNil(o.FreeAccess) {
    return nil, false
	}
	return o.FreeAccess, true
}

// HasFreeAccess returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) HasFreeAccess() bool {
	if o != nil && !isNil(o.FreeAccess) {
		return true
	}

	return false
}

// SetFreeAccess gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess and assigns it to the FreeAccess field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) SetFreeAccess(v GetNetworkWirelessSsidSplashSettings200ResponseBillingFreeAccess) {
	o.FreeAccess = &v
}

// GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetPrepaidAccessFastLoginEnabled() bool {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.PrepaidAccessFastLoginEnabled
}

// GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
    return nil, false
	}
	return o.PrepaidAccessFastLoginEnabled, true
}

// HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) HasPrepaidAccessFastLoginEnabled() bool {
	if o != nil && !isNil(o.PrepaidAccessFastLoginEnabled) {
		return true
	}

	return false
}

// SetPrepaidAccessFastLoginEnabled gets a reference to the given bool and assigns it to the PrepaidAccessFastLoginEnabled field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) SetPrepaidAccessFastLoginEnabled(v bool) {
	o.PrepaidAccessFastLoginEnabled = &v
}

// GetReplyToEmailAddress returns the ReplyToEmailAddress field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetReplyToEmailAddress() string {
	if o == nil || isNil(o.ReplyToEmailAddress) {
		var ret string
		return ret
	}
	return *o.ReplyToEmailAddress
}

// GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) GetReplyToEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.ReplyToEmailAddress) {
    return nil, false
	}
	return o.ReplyToEmailAddress, true
}

// HasReplyToEmailAddress returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) HasReplyToEmailAddress() bool {
	if o != nil && !isNil(o.ReplyToEmailAddress) {
		return true
	}

	return false
}

// SetReplyToEmailAddress gets a reference to the given string and assigns it to the ReplyToEmailAddress field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseBilling) SetReplyToEmailAddress(v string) {
	o.ReplyToEmailAddress = &v
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FreeAccess) {
		toSerialize["freeAccess"] = o.FreeAccess
	}
	if !isNil(o.PrepaidAccessFastLoginEnabled) {
		toSerialize["prepaidAccessFastLoginEnabled"] = o.PrepaidAccessFastLoginEnabled
	}
	if !isNil(o.ReplyToEmailAddress) {
		toSerialize["replyToEmailAddress"] = o.ReplyToEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling struct {
	value *GetNetworkWirelessSsidSplashSettings200ResponseBilling
	isSet bool
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) Get() *GetNetworkWirelessSsidSplashSettings200ResponseBilling {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) Set(val *GetNetworkWirelessSsidSplashSettings200ResponseBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidSplashSettings200ResponseBilling(val *GetNetworkWirelessSsidSplashSettings200ResponseBilling) *NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling {
	return &NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


