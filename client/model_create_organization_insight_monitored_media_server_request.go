/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateOrganizationInsightMonitoredMediaServerRequest struct for CreateOrganizationInsightMonitoredMediaServerRequest
type CreateOrganizationInsightMonitoredMediaServerRequest struct {
	// The name of the VoIP provider
	Name string `json:"name"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address string `json:"address"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewCreateOrganizationInsightMonitoredMediaServerRequest instantiates a new CreateOrganizationInsightMonitoredMediaServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInsightMonitoredMediaServerRequest(name string, address string) *CreateOrganizationInsightMonitoredMediaServerRequest {
	this := CreateOrganizationInsightMonitoredMediaServerRequest{}
	this.Name = name
	this.Address = address
	return &this
}

// NewCreateOrganizationInsightMonitoredMediaServerRequestWithDefaults instantiates a new CreateOrganizationInsightMonitoredMediaServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInsightMonitoredMediaServerRequestWithDefaults() *CreateOrganizationInsightMonitoredMediaServerRequest {
	this := CreateOrganizationInsightMonitoredMediaServerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetAddress(v string) {
	o.Address = v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabled() bool {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BestEffortMonitoringEnabled) {
    return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) HasBestEffortMonitoringEnabled() bool {
	if o != nil && !isNil(o.BestEffortMonitoringEnabled) {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o CreateOrganizationInsightMonitoredMediaServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.BestEffortMonitoringEnabled) {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInsightMonitoredMediaServerRequest struct {
	value *CreateOrganizationInsightMonitoredMediaServerRequest
	isSet bool
}

func (v NullableCreateOrganizationInsightMonitoredMediaServerRequest) Get() *CreateOrganizationInsightMonitoredMediaServerRequest {
	return v.value
}

func (v *NullableCreateOrganizationInsightMonitoredMediaServerRequest) Set(val *CreateOrganizationInsightMonitoredMediaServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInsightMonitoredMediaServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInsightMonitoredMediaServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInsightMonitoredMediaServerRequest(val *CreateOrganizationInsightMonitoredMediaServerRequest) *NullableCreateOrganizationInsightMonitoredMediaServerRequest {
	return &NullableCreateOrganizationInsightMonitoredMediaServerRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInsightMonitoredMediaServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInsightMonitoredMediaServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


