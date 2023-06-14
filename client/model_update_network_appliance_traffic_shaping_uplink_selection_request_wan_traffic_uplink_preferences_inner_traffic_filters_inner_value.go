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

// UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue Value object of this traffic filter
type UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	// Protocol of this custom type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource `json:"source"`
	Destination UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination `json:"destination"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue(source UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource, destination UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSource() UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	if o == nil {
		var ret UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSourceOk() (*UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetSource(v UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestination() UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination {
	if o == nil {
		var ret UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestinationOk() (*UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetDestination(v UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) {
	o.Destination = v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Get() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Set(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


