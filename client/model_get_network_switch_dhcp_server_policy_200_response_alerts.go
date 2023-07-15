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

// checks if the GetNetworkSwitchDhcpServerPolicy200ResponseAlerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpServerPolicy200ResponseAlerts{}

// GetNetworkSwitchDhcpServerPolicy200ResponseAlerts Email alert settings for DHCP servers
type GetNetworkSwitchDhcpServerPolicy200ResponseAlerts struct {
	Email *GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail `json:"email,omitempty"`
}

// NewGetNetworkSwitchDhcpServerPolicy200ResponseAlerts instantiates a new GetNetworkSwitchDhcpServerPolicy200ResponseAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpServerPolicy200ResponseAlerts() *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts {
	this := GetNetworkSwitchDhcpServerPolicy200ResponseAlerts{}
	return &this
}

// NewGetNetworkSwitchDhcpServerPolicy200ResponseAlertsWithDefaults instantiates a new GetNetworkSwitchDhcpServerPolicy200ResponseAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpServerPolicy200ResponseAlertsWithDefaults() *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts {
	this := GetNetworkSwitchDhcpServerPolicy200ResponseAlerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) GetEmail() GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail {
	if o == nil || IsNil(o.Email) {
		var ret GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) GetEmailOk() (*GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail and assigns it to the Email field.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) SetEmail(v GetNetworkSwitchDhcpServerPolicy200ResponseAlertsEmail) {
	o.Email = &v
}

func (o GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts struct {
	value *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) Get() *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) Set(val *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts(val *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) *NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts {
	return &NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


