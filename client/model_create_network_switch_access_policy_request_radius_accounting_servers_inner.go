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

// checks if the CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner{}

// CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner struct for CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner
type CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner struct {
	// Public IP address of the RADIUS accounting server
	Host string `json:"host"`
	// UDP port that the RADIUS Accounting server listens on for access requests
	Port int32 `json:"port"`
	// RADIUS client shared secret
	Secret string `json:"secret"`
}

// NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner(host string, port int32, secret string) *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner {
	this := CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner{}
	this.Host = host
	this.Port = port
	this.Secret = secret
	return &this
}

// NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults() *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner {
	this := CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetPort(v int32) {
	o.Port = v
}

// GetSecret returns the Secret field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetSecret(v string) {
	o.Secret = v
}

func (o CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

type NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner struct {
	value *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner
	isSet bool
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) Get() *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner {
	return v.value
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) Set(val *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner(val *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) *NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner {
	return &NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


