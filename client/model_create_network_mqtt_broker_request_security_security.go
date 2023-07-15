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

// checks if the CreateNetworkMqttBrokerRequestSecuritySecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkMqttBrokerRequestSecuritySecurity{}

// CreateNetworkMqttBrokerRequestSecuritySecurity TLS settings of the MQTT broker.
type CreateNetworkMqttBrokerRequestSecuritySecurity struct {
	// CA Certificate of the MQTT broker.
	CaCertificate *string `json:"caCertificate,omitempty"`
	// Whether the TLS hostname verification is enabled for the MQTT broker.
	VerifyHostnames *bool `json:"verifyHostnames,omitempty"`
}

// NewCreateNetworkMqttBrokerRequestSecuritySecurity instantiates a new CreateNetworkMqttBrokerRequestSecuritySecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkMqttBrokerRequestSecuritySecurity() *CreateNetworkMqttBrokerRequestSecuritySecurity {
	this := CreateNetworkMqttBrokerRequestSecuritySecurity{}
	return &this
}

// NewCreateNetworkMqttBrokerRequestSecuritySecurityWithDefaults instantiates a new CreateNetworkMqttBrokerRequestSecuritySecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkMqttBrokerRequestSecuritySecurityWithDefaults() *CreateNetworkMqttBrokerRequestSecuritySecurity {
	this := CreateNetworkMqttBrokerRequestSecuritySecurity{}
	return &this
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetVerifyHostnames returns the VerifyHostnames field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetVerifyHostnames() bool {
	if o == nil || IsNil(o.VerifyHostnames) {
		var ret bool
		return ret
	}
	return *o.VerifyHostnames
}

// GetVerifyHostnamesOk returns a tuple with the VerifyHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetVerifyHostnamesOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyHostnames) {
		return nil, false
	}
	return o.VerifyHostnames, true
}

// HasVerifyHostnames returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) HasVerifyHostnames() bool {
	if o != nil && !IsNil(o.VerifyHostnames) {
		return true
	}

	return false
}

// SetVerifyHostnames gets a reference to the given bool and assigns it to the VerifyHostnames field.
func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) SetVerifyHostnames(v bool) {
	o.VerifyHostnames = &v
}

func (o CreateNetworkMqttBrokerRequestSecuritySecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkMqttBrokerRequestSecuritySecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	if !IsNil(o.VerifyHostnames) {
		toSerialize["verifyHostnames"] = o.VerifyHostnames
	}
	return toSerialize, nil
}

type NullableCreateNetworkMqttBrokerRequestSecuritySecurity struct {
	value *CreateNetworkMqttBrokerRequestSecuritySecurity
	isSet bool
}

func (v NullableCreateNetworkMqttBrokerRequestSecuritySecurity) Get() *CreateNetworkMqttBrokerRequestSecuritySecurity {
	return v.value
}

func (v *NullableCreateNetworkMqttBrokerRequestSecuritySecurity) Set(val *CreateNetworkMqttBrokerRequestSecuritySecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkMqttBrokerRequestSecuritySecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkMqttBrokerRequestSecuritySecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkMqttBrokerRequestSecuritySecurity(val *CreateNetworkMqttBrokerRequestSecuritySecurity) *NullableCreateNetworkMqttBrokerRequestSecuritySecurity {
	return &NullableCreateNetworkMqttBrokerRequestSecuritySecurity{value: val, isSet: true}
}

func (v NullableCreateNetworkMqttBrokerRequestSecuritySecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkMqttBrokerRequestSecuritySecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


