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

// checks if the GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient{}

// GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient Client.
type GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient struct {
	// Client ID.
	Id *string `json:"id,omitempty"`
	// MAC address.
	Mac *string `json:"mac,omitempty"`
}

// NewGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient instantiates a new GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient() *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient {
	this := GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient{}
	return &this
}

// NewGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClientWithDefaults instantiates a new GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClientWithDefaults() *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient {
	this := GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) SetMac(v string) {
	o.Mac = &v
}

func (o GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient struct {
	value *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) Get() *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) Set(val *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient(val *GetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) *NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient {
	return &NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesPacketLossByClient200ResponseInnerClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


