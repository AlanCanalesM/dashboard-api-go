/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkFirmwareUpgradesStagedEvents200Response struct for GetNetworkFirmwareUpgradesStagedEvents200Response
type GetNetworkFirmwareUpgradesStagedEvents200Response struct {
	Products *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts `json:"products,omitempty"`
	// The ordered stages in the network
	Stages []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner `json:"stages,omitempty"`
	// Reasons for the rollback
	Reasons []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner `json:"reasons,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200Response instantiates a new GetNetworkFirmwareUpgradesStagedEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200Response() *GetNetworkFirmwareUpgradesStagedEvents200Response {
	this := GetNetworkFirmwareUpgradesStagedEvents200Response{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200Response {
	this := GetNetworkFirmwareUpgradesStagedEvents200Response{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetProducts() GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts {
	if o == nil || isNil(o.Products) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetProductsOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts and assigns it to the Products field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetProducts(v GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) {
	o.Products = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetStages() []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner {
	if o == nil || isNil(o.Stages) {
		var ret []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetStagesOk() ([]GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner, bool) {
	if o == nil || isNil(o.Stages) {
    return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasStages() bool {
	if o != nil && !isNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner and assigns it to the Stages field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetStages(v []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetReasons() []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	if o == nil || isNil(o.Reasons) {
		var ret []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetReasonsOk() ([]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner and assigns it to the Reasons field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetReasons(v []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) {
	o.Reasons = v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !isNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200Response struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200Response
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200Response) Get() *GetNetworkFirmwareUpgradesStagedEvents200Response {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200Response) Set(val *GetNetworkFirmwareUpgradesStagedEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200Response(val *GetNetworkFirmwareUpgradesStagedEvents200Response) *NullableGetNetworkFirmwareUpgradesStagedEvents200Response {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200Response{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


