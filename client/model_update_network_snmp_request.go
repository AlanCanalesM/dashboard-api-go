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

// UpdateNetworkSnmpRequest struct for UpdateNetworkSnmpRequest
type UpdateNetworkSnmpRequest struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// The SNMP community string. Only relevant if 'access' is set to 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// The list of SNMP users. Only relevant if 'access' is set to 'users'.
	Users []UpdateNetworkSnmpRequestUsersInner `json:"users,omitempty"`
}

// NewUpdateNetworkSnmpRequest instantiates a new UpdateNetworkSnmpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSnmpRequest() *UpdateNetworkSnmpRequest {
	this := UpdateNetworkSnmpRequest{}
	return &this
}

// NewUpdateNetworkSnmpRequestWithDefaults instantiates a new UpdateNetworkSnmpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSnmpRequestWithDefaults() *UpdateNetworkSnmpRequest {
	this := UpdateNetworkSnmpRequest{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *UpdateNetworkSnmpRequest) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSnmpRequest) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *UpdateNetworkSnmpRequest) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *UpdateNetworkSnmpRequest) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *UpdateNetworkSnmpRequest) GetCommunityString() string {
	if o == nil || isNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSnmpRequest) GetCommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.CommunityString) {
    return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *UpdateNetworkSnmpRequest) HasCommunityString() bool {
	if o != nil && !isNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *UpdateNetworkSnmpRequest) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UpdateNetworkSnmpRequest) GetUsers() []UpdateNetworkSnmpRequestUsersInner {
	if o == nil || isNil(o.Users) {
		var ret []UpdateNetworkSnmpRequestUsersInner
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSnmpRequest) GetUsersOk() ([]UpdateNetworkSnmpRequestUsersInner, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UpdateNetworkSnmpRequest) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UpdateNetworkSnmpRequestUsersInner and assigns it to the Users field.
func (o *UpdateNetworkSnmpRequest) SetUsers(v []UpdateNetworkSnmpRequestUsersInner) {
	o.Users = v
}

func (o UpdateNetworkSnmpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.CommunityString) {
		toSerialize["communityString"] = o.CommunityString
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSnmpRequest struct {
	value *UpdateNetworkSnmpRequest
	isSet bool
}

func (v NullableUpdateNetworkSnmpRequest) Get() *UpdateNetworkSnmpRequest {
	return v.value
}

func (v *NullableUpdateNetworkSnmpRequest) Set(val *UpdateNetworkSnmpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSnmpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSnmpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSnmpRequest(val *UpdateNetworkSnmpRequest) *NullableUpdateNetworkSnmpRequest {
	return &NullableUpdateNetworkSnmpRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSnmpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSnmpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


