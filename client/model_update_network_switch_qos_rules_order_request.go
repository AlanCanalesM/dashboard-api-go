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

// UpdateNetworkSwitchQosRulesOrderRequest struct for UpdateNetworkSwitchQosRulesOrderRequest
type UpdateNetworkSwitchQosRulesOrderRequest struct {
	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	RuleIds []string `json:"ruleIds"`
}

// NewUpdateNetworkSwitchQosRulesOrderRequest instantiates a new UpdateNetworkSwitchQosRulesOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchQosRulesOrderRequest(ruleIds []string) *UpdateNetworkSwitchQosRulesOrderRequest {
	this := UpdateNetworkSwitchQosRulesOrderRequest{}
	this.RuleIds = ruleIds
	return &this
}

// NewUpdateNetworkSwitchQosRulesOrderRequestWithDefaults instantiates a new UpdateNetworkSwitchQosRulesOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchQosRulesOrderRequestWithDefaults() *UpdateNetworkSwitchQosRulesOrderRequest {
	this := UpdateNetworkSwitchQosRulesOrderRequest{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *UpdateNetworkSwitchQosRulesOrderRequest) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRulesOrderRequest) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *UpdateNetworkSwitchQosRulesOrderRequest) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o UpdateNetworkSwitchQosRulesOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchQosRulesOrderRequest struct {
	value *UpdateNetworkSwitchQosRulesOrderRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchQosRulesOrderRequest) Get() *UpdateNetworkSwitchQosRulesOrderRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchQosRulesOrderRequest) Set(val *UpdateNetworkSwitchQosRulesOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchQosRulesOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchQosRulesOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchQosRulesOrderRequest(val *UpdateNetworkSwitchQosRulesOrderRequest) *NullableUpdateNetworkSwitchQosRulesOrderRequest {
	return &NullableUpdateNetworkSwitchQosRulesOrderRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchQosRulesOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchQosRulesOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


