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

// OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi An object containing wifi utilization.
type OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi struct {
	// Percentage of wifi channel utiliation for the given band.
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifiWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifiWithDefaults() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi {
	this := OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi struct {
	value *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) Get() *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) Set(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi(val *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi {
	return &NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


