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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel TLS Related Parameters
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel struct {
	// Name of the configured TLS certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Name of the configured TLS tunnel
	Name *string `json:"name,omitempty"`
	// Number of the configured Loopback Interface used for TLS overlay
	LoopbackNumber *int32 `json:"loopbackNumber,omitempty"`
	// Number of the vlan expected to be used to connect to the cloud
	LocalInterface *int32 `json:"localInterface,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnelWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnelWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel{}
	return &this
}

// GetCertificateName returns the CertificateName field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetCertificateName() string {
	if o == nil || IsNil(o.CertificateName) {
		var ret string
		return ret
	}
	return *o.CertificateName
}

// GetCertificateNameOk returns a tuple with the CertificateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetCertificateNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateName) {
		return nil, false
	}
	return o.CertificateName, true
}

// HasCertificateName returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) HasCertificateName() bool {
	if o != nil && !IsNil(o.CertificateName) {
		return true
	}

	return false
}

// SetCertificateName gets a reference to the given string and assigns it to the CertificateName field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) SetCertificateName(v string) {
	o.CertificateName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) SetName(v string) {
	o.Name = &v
}

// GetLoopbackNumber returns the LoopbackNumber field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetLoopbackNumber() int32 {
	if o == nil || IsNil(o.LoopbackNumber) {
		var ret int32
		return ret
	}
	return *o.LoopbackNumber
}

// GetLoopbackNumberOk returns a tuple with the LoopbackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetLoopbackNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.LoopbackNumber) {
		return nil, false
	}
	return o.LoopbackNumber, true
}

// HasLoopbackNumber returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) HasLoopbackNumber() bool {
	if o != nil && !IsNil(o.LoopbackNumber) {
		return true
	}

	return false
}

// SetLoopbackNumber gets a reference to the given int32 and assigns it to the LoopbackNumber field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) SetLoopbackNumber(v int32) {
	o.LoopbackNumber = &v
}

// GetLocalInterface returns the LocalInterface field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetLocalInterface() int32 {
	if o == nil || IsNil(o.LocalInterface) {
		var ret int32
		return ret
	}
	return *o.LocalInterface
}

// GetLocalInterfaceOk returns a tuple with the LocalInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) GetLocalInterfaceOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalInterface) {
		return nil, false
	}
	return o.LocalInterface, true
}

// HasLocalInterface returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) HasLocalInterface() bool {
	if o != nil && !IsNil(o.LocalInterface) {
		return true
	}

	return false
}

// SetLocalInterface gets a reference to the given int32 and assigns it to the LocalInterface field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) SetLocalInterface(v int32) {
	o.LocalInterface = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateName) {
		toSerialize["certificateName"] = o.CertificateName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LoopbackNumber) {
		toSerialize["loopbackNumber"] = o.LoopbackNumber
	}
	if !IsNil(o.LocalInterface) {
		toSerialize["localInterface"] = o.LocalInterface
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


