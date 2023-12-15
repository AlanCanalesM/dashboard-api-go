/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner struct for CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner struct {
	// Device SUDI certificate
	Sudi string `json:"sudi"`
	Tunnel *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel `json:"tunnel,omitempty"`
	User *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser `json:"user,omitempty"`
	Vty *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty `json:"vty,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner(sudi string) *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{}
	this.Sudi = sudi
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{}
	return &this
}

// GetSudi returns the Sudi field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetSudi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sudi
}

// GetSudiOk returns a tuple with the Sudi field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetSudiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sudi, true
}

// SetSudi sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) SetSudi(v string) {
	o.Sudi = v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetTunnel() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel {
	if o == nil || IsNil(o.Tunnel) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetTunnelOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel, bool) {
	if o == nil || IsNil(o.Tunnel) {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) HasTunnel() bool {
	if o != nil && !IsNil(o.Tunnel) {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel and assigns it to the Tunnel field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) SetTunnel(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) {
	o.Tunnel = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetUser() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser {
	if o == nil || IsNil(o.User) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetUserOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser and assigns it to the User field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) SetUser(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) {
	o.User = &v
}

// GetVty returns the Vty field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetVty() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty {
	if o == nil || IsNil(o.Vty) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty
		return ret
	}
	return *o.Vty
}

// GetVtyOk returns a tuple with the Vty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) GetVtyOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty, bool) {
	if o == nil || IsNil(o.Vty) {
		return nil, false
	}
	return o.Vty, true
}

// HasVty returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) HasVty() bool {
	if o != nil && !IsNil(o.Vty) {
		return true
	}

	return false
}

// SetVty gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty and assigns it to the Vty field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) SetVty(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) {
	o.Vty = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sudi"] = o.Sudi
	if !IsNil(o.Tunnel) {
		toSerialize["tunnel"] = o.Tunnel
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Vty) {
		toSerialize["vty"] = o.Vty
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


