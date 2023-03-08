/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject80 struct for InlineObject80
type InlineObject80 struct {
	Products *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts `json:"products,omitempty"`
	// All firmware upgrade stages in the network with their start time.
	Stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages `json:"stages"`
}

// NewInlineObject80 instantiates a new InlineObject80 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject80(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *InlineObject80 {
	this := InlineObject80{}
	this.Stages = stages
	return &this
}

// NewInlineObject80WithDefaults instantiates a new InlineObject80 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject80WithDefaults() *InlineObject80 {
	this := InlineObject80{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineObject80) GetProducts() NetworksNetworkIdFirmwareUpgradesStagedEventsProducts {
	if o == nil || isNil(o.Products) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProducts, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineObject80) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProducts and assigns it to the Products field.
func (o *InlineObject80) SetProducts(v NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) {
	o.Products = &v
}

// GetStages returns the Stages field value
func (o *InlineObject80) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedEventsStages
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InlineObject80) GetStagesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InlineObject80) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	o.Stages = v
}

func (o InlineObject80) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject80 struct {
	value *InlineObject80
	isSet bool
}

func (v NullableInlineObject80) Get() *InlineObject80 {
	return v.value
}

func (v *NullableInlineObject80) Set(val *InlineObject80) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject80) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject80) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject80(val *InlineObject80) *NullableInlineObject80 {
	return &NullableInlineObject80{value: val, isSet: true}
}

func (v NullableInlineObject80) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject80) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


