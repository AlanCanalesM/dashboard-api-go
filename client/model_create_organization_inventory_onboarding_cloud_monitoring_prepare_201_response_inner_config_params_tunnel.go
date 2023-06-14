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

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel Configuration options used to connect to the device
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel struct {
	// 
	Mode *string `json:"mode,omitempty"`
	// The port used for the ssh tunnel.
	Port *string `json:"port,omitempty"`
	// SSH tunnel URL used to connect to the device
	Host *string `json:"host,omitempty"`
	// The name of the tunnel we are attempting to connect to
	Name *string `json:"name,omitempty"`
	RootCertificate *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate `json:"rootCertificate,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) SetMode(v string) {
	o.Mode = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) SetPort(v string) {
	o.Port = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) SetHost(v string) {
	o.Host = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) SetName(v string) {
	o.Name = &v
}

// GetRootCertificate returns the RootCertificate field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetRootCertificate() CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate {
	if o == nil || isNil(o.RootCertificate) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate
		return ret
	}
	return *o.RootCertificate
}

// GetRootCertificateOk returns a tuple with the RootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) GetRootCertificateOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate, bool) {
	if o == nil || isNil(o.RootCertificate) {
    return nil, false
	}
	return o.RootCertificate, true
}

// HasRootCertificate returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) HasRootCertificate() bool {
	if o != nil && !isNil(o.RootCertificate) {
		return true
	}

	return false
}

// SetRootCertificate gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate and assigns it to the RootCertificate field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) SetRootCertificate(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) {
	o.RootCertificate = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.RootCertificate) {
		toSerialize["rootCertificate"] = o.RootCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


