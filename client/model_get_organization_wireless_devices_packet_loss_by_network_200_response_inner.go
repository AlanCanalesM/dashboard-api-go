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

// checks if the GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner{}

// GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner struct for GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner
type GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner struct {
	Downstream *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream `json:"downstream,omitempty"`
	Upstream *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream `json:"upstream,omitempty"`
	Network *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork `json:"network,omitempty"`
}

// NewGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner instantiates a new GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner() *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner {
	this := GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner {
	this := GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner{}
	return &this
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetDownstream() GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream {
	if o == nil || IsNil(o.Downstream) {
		var ret GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetDownstreamOk() (*GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream, bool) {
	if o == nil || IsNil(o.Downstream) {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) HasDownstream() bool {
	if o != nil && !IsNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream and assigns it to the Downstream field.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) SetDownstream(v GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerDownstream) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetUpstream() GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream {
	if o == nil || IsNil(o.Upstream) {
		var ret GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetUpstreamOk() (*GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream and assigns it to the Upstream field.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) SetUpstream(v GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerUpstream) {
	o.Upstream = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerNetwork) {
	o.Network = &v
}

func (o GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner struct {
	value *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) Get() *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) Set(val *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner(val *GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) *NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner {
	return &NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


