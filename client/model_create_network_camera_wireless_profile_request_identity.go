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

// checks if the CreateNetworkCameraWirelessProfileRequestIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraWirelessProfileRequestIdentity{}

// CreateNetworkCameraWirelessProfileRequestIdentity The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
type CreateNetworkCameraWirelessProfileRequestIdentity struct {
	// The username of the identity.
	Username *string `json:"username,omitempty"`
	// The password of the identity.
	Password *string `json:"password,omitempty"`
}

// NewCreateNetworkCameraWirelessProfileRequestIdentity instantiates a new CreateNetworkCameraWirelessProfileRequestIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraWirelessProfileRequestIdentity() *CreateNetworkCameraWirelessProfileRequestIdentity {
	this := CreateNetworkCameraWirelessProfileRequestIdentity{}
	return &this
}

// NewCreateNetworkCameraWirelessProfileRequestIdentityWithDefaults instantiates a new CreateNetworkCameraWirelessProfileRequestIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraWirelessProfileRequestIdentityWithDefaults() *CreateNetworkCameraWirelessProfileRequestIdentity {
	this := CreateNetworkCameraWirelessProfileRequestIdentity{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateNetworkCameraWirelessProfileRequestIdentity) SetPassword(v string) {
	o.Password = &v
}

func (o CreateNetworkCameraWirelessProfileRequestIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraWirelessProfileRequestIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableCreateNetworkCameraWirelessProfileRequestIdentity struct {
	value *CreateNetworkCameraWirelessProfileRequestIdentity
	isSet bool
}

func (v NullableCreateNetworkCameraWirelessProfileRequestIdentity) Get() *CreateNetworkCameraWirelessProfileRequestIdentity {
	return v.value
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestIdentity) Set(val *CreateNetworkCameraWirelessProfileRequestIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraWirelessProfileRequestIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraWirelessProfileRequestIdentity(val *CreateNetworkCameraWirelessProfileRequestIdentity) *NullableCreateNetworkCameraWirelessProfileRequestIdentity {
	return &NullableCreateNetworkCameraWirelessProfileRequestIdentity{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraWirelessProfileRequestIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


