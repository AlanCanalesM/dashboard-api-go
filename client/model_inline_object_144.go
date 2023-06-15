/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject144 struct for InlineObject144
type InlineObject144 struct {
	//     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames). 
	Mode *string `json:"mode,omitempty"`
	// The list of items that make up the custom pie chart for traffic reporting.
	CustomPieChartItems []NetworksNetworkIdTrafficAnalysisCustomPieChartItems `json:"customPieChartItems,omitempty"`
}

// NewInlineObject144 instantiates a new InlineObject144 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject144() *InlineObject144 {
	this := InlineObject144{}
	return &this
}

// NewInlineObject144WithDefaults instantiates a new InlineObject144 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject144WithDefaults() *InlineObject144 {
	this := InlineObject144{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineObject144) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject144) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineObject144) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineObject144) SetMode(v string) {
	o.Mode = &v
}

// GetCustomPieChartItems returns the CustomPieChartItems field value if set, zero value otherwise.
func (o *InlineObject144) GetCustomPieChartItems() []NetworksNetworkIdTrafficAnalysisCustomPieChartItems {
	if o == nil || isNil(o.CustomPieChartItems) {
		var ret []NetworksNetworkIdTrafficAnalysisCustomPieChartItems
		return ret
	}
	return o.CustomPieChartItems
}

// GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject144) GetCustomPieChartItemsOk() ([]NetworksNetworkIdTrafficAnalysisCustomPieChartItems, bool) {
	if o == nil || isNil(o.CustomPieChartItems) {
    return nil, false
	}
	return o.CustomPieChartItems, true
}

// HasCustomPieChartItems returns a boolean if a field has been set.
func (o *InlineObject144) HasCustomPieChartItems() bool {
	if o != nil && !isNil(o.CustomPieChartItems) {
		return true
	}

	return false
}

// SetCustomPieChartItems gets a reference to the given []NetworksNetworkIdTrafficAnalysisCustomPieChartItems and assigns it to the CustomPieChartItems field.
func (o *InlineObject144) SetCustomPieChartItems(v []NetworksNetworkIdTrafficAnalysisCustomPieChartItems) {
	o.CustomPieChartItems = v
}

func (o InlineObject144) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.CustomPieChartItems) {
		toSerialize["customPieChartItems"] = o.CustomPieChartItems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject144 struct {
	value *InlineObject144
	isSet bool
}

func (v NullableInlineObject144) Get() *InlineObject144 {
	return v.value
}

func (v *NullableInlineObject144) Set(val *InlineObject144) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject144) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject144) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject144(val *InlineObject144) *NullableInlineObject144 {
	return &NullableInlineObject144{value: val, isSet: true}
}

func (v NullableInlineObject144) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject144) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


