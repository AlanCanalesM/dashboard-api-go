/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkMqttBrokerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkMqttBrokerRequest{}

// CreateNetworkMqttBrokerRequest struct for CreateNetworkMqttBrokerRequest
type CreateNetworkMqttBrokerRequest struct {
	// Name of the MQTT broker.
	Name string `json:"name"`
	// Host name/IP address where the MQTT broker runs.
	Host string `json:"host"`
	// Host port though which the MQTT broker can be reached.
	Port int32 `json:"port"`
	Security *CreateNetworkMqttBrokerRequestSecurity `json:"security,omitempty"`
	// Authentication settings of the MQTT broker
	Authentication map[string]interface{} `json:"authentication,omitempty"`
}

// NewCreateNetworkMqttBrokerRequest instantiates a new CreateNetworkMqttBrokerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkMqttBrokerRequest(name string, host string, port int32) *CreateNetworkMqttBrokerRequest {
	this := CreateNetworkMqttBrokerRequest{}
	this.Name = name
	this.Host = host
	this.Port = port
	return &this
}

// NewCreateNetworkMqttBrokerRequestWithDefaults instantiates a new CreateNetworkMqttBrokerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkMqttBrokerRequestWithDefaults() *CreateNetworkMqttBrokerRequest {
	this := CreateNetworkMqttBrokerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkMqttBrokerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkMqttBrokerRequest) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *CreateNetworkMqttBrokerRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateNetworkMqttBrokerRequest) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *CreateNetworkMqttBrokerRequest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateNetworkMqttBrokerRequest) SetPort(v int32) {
	o.Port = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequest) GetSecurity() CreateNetworkMqttBrokerRequestSecurity {
	if o == nil || IsNil(o.Security) {
		var ret CreateNetworkMqttBrokerRequestSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequest) GetSecurityOk() (*CreateNetworkMqttBrokerRequestSecurity, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequest) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given CreateNetworkMqttBrokerRequestSecurity and assigns it to the Security field.
func (o *CreateNetworkMqttBrokerRequest) SetSecurity(v CreateNetworkMqttBrokerRequestSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *CreateNetworkMqttBrokerRequest) GetAuthentication() map[string]interface{} {
	if o == nil || IsNil(o.Authentication) {
		var ret map[string]interface{}
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMqttBrokerRequest) GetAuthenticationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Authentication) {
		return map[string]interface{}{}, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *CreateNetworkMqttBrokerRequest) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *CreateNetworkMqttBrokerRequest) SetAuthentication(v map[string]interface{}) {
	o.Authentication = v
}

func (o CreateNetworkMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkMqttBrokerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableCreateNetworkMqttBrokerRequest struct {
	value *CreateNetworkMqttBrokerRequest
	isSet bool
}

func (v NullableCreateNetworkMqttBrokerRequest) Get() *CreateNetworkMqttBrokerRequest {
	return v.value
}

func (v *NullableCreateNetworkMqttBrokerRequest) Set(val *CreateNetworkMqttBrokerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkMqttBrokerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkMqttBrokerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkMqttBrokerRequest(val *CreateNetworkMqttBrokerRequest) *NullableCreateNetworkMqttBrokerRequest {
	return &NullableCreateNetworkMqttBrokerRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkMqttBrokerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


