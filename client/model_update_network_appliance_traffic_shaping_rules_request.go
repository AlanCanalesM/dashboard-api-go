/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkApplianceTrafficShapingRulesRequest struct for UpdateNetworkApplianceTrafficShapingRulesRequest
type UpdateNetworkApplianceTrafficShapingRulesRequest struct {
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequest instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRulesRequest() *UpdateNetworkApplianceTrafficShapingRulesRequest {
	this := UpdateNetworkApplianceTrafficShapingRulesRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRulesRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequest {
	this := UpdateNetworkApplianceTrafficShapingRulesRequest{}
	return &this
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetDefaultRulesEnabled() bool {
	if o == nil || isNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRulesEnabled) {
    return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) HasDefaultRulesEnabled() bool {
	if o != nil && !isNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetRules() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) SetRules(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultRulesEnabled) {
		toSerialize["defaultRulesEnabled"] = o.DefaultRulesEnabled
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingRulesRequest struct {
	value *UpdateNetworkApplianceTrafficShapingRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequest) Get() *UpdateNetworkApplianceTrafficShapingRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequest) Set(val *UpdateNetworkApplianceTrafficShapingRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRulesRequest(val *UpdateNetworkApplianceTrafficShapingRulesRequest) *NullableUpdateNetworkApplianceTrafficShapingRulesRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


