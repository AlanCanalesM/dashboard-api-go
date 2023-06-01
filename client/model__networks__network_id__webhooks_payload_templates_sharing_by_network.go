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

// NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork Information on network access to the template
type NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork struct {
	// Indicates whether network admins may modify this template
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"`
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork() *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork {
	this := NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork{}
	return &this
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetworkWithDefaults instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetworkWithDefaults() *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork {
	this := NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork{}
	return &this
}

// GetAdminsCanModify returns the AdminsCanModify field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) GetAdminsCanModify() bool {
	if o == nil || isNil(o.AdminsCanModify) {
		var ret bool
		return ret
	}
	return *o.AdminsCanModify
}

// GetAdminsCanModifyOk returns a tuple with the AdminsCanModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) GetAdminsCanModifyOk() (*bool, bool) {
	if o == nil || isNil(o.AdminsCanModify) {
    return nil, false
	}
	return o.AdminsCanModify, true
}

// HasAdminsCanModify returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) HasAdminsCanModify() bool {
	if o != nil && !isNil(o.AdminsCanModify) {
		return true
	}

	return false
}

// SetAdminsCanModify gets a reference to the given bool and assigns it to the AdminsCanModify field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) SetAdminsCanModify(v bool) {
	o.AdminsCanModify = &v
}

func (o NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminsCanModify) {
		toSerialize["adminsCanModify"] = o.AdminsCanModify
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork struct {
	value *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) Get() *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) Set(val *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork(val *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork {
	return &NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


