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

// checks if the GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2{}

// GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 uplink wan2 configured limits [optional]
type GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 struct {
	// configured UP limit for the uplink (in Kbps).  Null indicated unlimited
	LimitUp *int32 `json:"limitUp,omitempty"`
	// configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 instantiates a new GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2() *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 {
	this := GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2{}
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2WithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2WithDefaults() *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 {
	this := GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 struct {
	value *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) Get() *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) Set(val *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2(val *GetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) *NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2 {
	return &NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkBandwidth200ResponseBandwidthLimitsWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


