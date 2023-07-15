/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi{}

// GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi An object containing non-wifi utilization.
type GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi struct {
	// Percentage of non-wifi channel utiliation for the given band.
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifiWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifiWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) Get() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) Set(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


