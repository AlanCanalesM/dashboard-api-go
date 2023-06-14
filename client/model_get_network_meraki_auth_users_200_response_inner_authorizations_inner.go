/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner struct for GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner
type GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner struct {
	// SSID number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Authorized zone of the user
	AuthorizedZone *string `json:"authorizedZone,omitempty"`
	// Authorization expiration time
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// User is authorized by the account name
	AuthorizedByName *string `json:"authorizedByName,omitempty"`
	// User is authorized by the account email address
	AuthorizedByEmail *string `json:"authorizedByEmail,omitempty"`
}

// NewGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner instantiates a new GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner() *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner {
	this := GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner{}
	return &this
}

// NewGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInnerWithDefaults instantiates a new GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInnerWithDefaults() *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner {
	this := GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetAuthorizedZone returns the AuthorizedZone field value if set, zero value otherwise.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedZone() string {
	if o == nil || isNil(o.AuthorizedZone) {
		var ret string
		return ret
	}
	return *o.AuthorizedZone
}

// GetAuthorizedZoneOk returns a tuple with the AuthorizedZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedZoneOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedZone) {
    return nil, false
	}
	return o.AuthorizedZone, true
}

// HasAuthorizedZone returns a boolean if a field has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) HasAuthorizedZone() bool {
	if o != nil && !isNil(o.AuthorizedZone) {
		return true
	}

	return false
}

// SetAuthorizedZone gets a reference to the given string and assigns it to the AuthorizedZone field.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) SetAuthorizedZone(v string) {
	o.AuthorizedZone = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetAuthorizedByName returns the AuthorizedByName field value if set, zero value otherwise.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedByName() string {
	if o == nil || isNil(o.AuthorizedByName) {
		var ret string
		return ret
	}
	return *o.AuthorizedByName
}

// GetAuthorizedByNameOk returns a tuple with the AuthorizedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedByNameOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedByName) {
    return nil, false
	}
	return o.AuthorizedByName, true
}

// HasAuthorizedByName returns a boolean if a field has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) HasAuthorizedByName() bool {
	if o != nil && !isNil(o.AuthorizedByName) {
		return true
	}

	return false
}

// SetAuthorizedByName gets a reference to the given string and assigns it to the AuthorizedByName field.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) SetAuthorizedByName(v string) {
	o.AuthorizedByName = &v
}

// GetAuthorizedByEmail returns the AuthorizedByEmail field value if set, zero value otherwise.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedByEmail() string {
	if o == nil || isNil(o.AuthorizedByEmail) {
		var ret string
		return ret
	}
	return *o.AuthorizedByEmail
}

// GetAuthorizedByEmailOk returns a tuple with the AuthorizedByEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) GetAuthorizedByEmailOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedByEmail) {
    return nil, false
	}
	return o.AuthorizedByEmail, true
}

// HasAuthorizedByEmail returns a boolean if a field has been set.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) HasAuthorizedByEmail() bool {
	if o != nil && !isNil(o.AuthorizedByEmail) {
		return true
	}

	return false
}

// SetAuthorizedByEmail gets a reference to the given string and assigns it to the AuthorizedByEmail field.
func (o *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) SetAuthorizedByEmail(v string) {
	o.AuthorizedByEmail = &v
}

func (o GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.AuthorizedZone) {
		toSerialize["authorizedZone"] = o.AuthorizedZone
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.AuthorizedByName) {
		toSerialize["authorizedByName"] = o.AuthorizedByName
	}
	if !isNil(o.AuthorizedByEmail) {
		toSerialize["authorizedByEmail"] = o.AuthorizedByEmail
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner struct {
	value *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner
	isSet bool
}

func (v NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) Get() *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner {
	return v.value
}

func (v *NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) Set(val *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner(val *GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) *NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner {
	return &NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner{value: val, isSet: true}
}

func (v NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


