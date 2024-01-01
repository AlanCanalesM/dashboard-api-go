/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceFirewallL7FirewallRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallL7FirewallRulesRequest{}

// UpdateNetworkApplianceFirewallL7FirewallRulesRequest struct for UpdateNetworkApplianceFirewallL7FirewallRulesRequest
type UpdateNetworkApplianceFirewallL7FirewallRulesRequest struct {
	// An ordered array of the MX L7 firewall rules
	Rules []UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkApplianceFirewallL7FirewallRulesRequest instantiates a new UpdateNetworkApplianceFirewallL7FirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallL7FirewallRulesRequest() *UpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallL7FirewallRulesRequest{}
	return &this
}

// NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallL7FirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestWithDefaults() *UpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallL7FirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) GetRules() []UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) SetRules(v []UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkApplianceFirewallL7FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallL7FirewallRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest struct {
	value *UpdateNetworkApplianceFirewallL7FirewallRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) Get() *UpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) Set(val *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest(val *UpdateNetworkApplianceFirewallL7FirewallRulesRequest) *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest {
	return &NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


