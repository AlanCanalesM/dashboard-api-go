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

// UpdateNetworkSmDevicesFieldsRequest struct for UpdateNetworkSmDevicesFieldsRequest
type UpdateNetworkSmDevicesFieldsRequest struct {
	// The wifiMac of the device to be modified.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The id of the device to be modified.
	Id *string `json:"id,omitempty"`
	// The serial of the device to be modified.
	Serial *string `json:"serial,omitempty"`
	DeviceFields UpdateNetworkSmDevicesFieldsRequestDeviceFields `json:"deviceFields"`
}

// NewUpdateNetworkSmDevicesFieldsRequest instantiates a new UpdateNetworkSmDevicesFieldsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSmDevicesFieldsRequest(deviceFields UpdateNetworkSmDevicesFieldsRequestDeviceFields) *UpdateNetworkSmDevicesFieldsRequest {
	this := UpdateNetworkSmDevicesFieldsRequest{}
	this.DeviceFields = deviceFields
	return &this
}

// NewUpdateNetworkSmDevicesFieldsRequestWithDefaults instantiates a new UpdateNetworkSmDevicesFieldsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSmDevicesFieldsRequestWithDefaults() *UpdateNetworkSmDevicesFieldsRequest {
	this := UpdateNetworkSmDevicesFieldsRequest{}
	return &this
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *UpdateNetworkSmDevicesFieldsRequest) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkSmDevicesFieldsRequest) SetId(v string) {
	o.Id = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *UpdateNetworkSmDevicesFieldsRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetDeviceFields returns the DeviceFields field value
func (o *UpdateNetworkSmDevicesFieldsRequest) GetDeviceFields() UpdateNetworkSmDevicesFieldsRequestDeviceFields {
	if o == nil {
		var ret UpdateNetworkSmDevicesFieldsRequestDeviceFields
		return ret
	}

	return o.DeviceFields
}

// GetDeviceFieldsOk returns a tuple with the DeviceFields field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequest) GetDeviceFieldsOk() (*UpdateNetworkSmDevicesFieldsRequestDeviceFields, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeviceFields, true
}

// SetDeviceFields sets field value
func (o *UpdateNetworkSmDevicesFieldsRequest) SetDeviceFields(v UpdateNetworkSmDevicesFieldsRequestDeviceFields) {
	o.DeviceFields = v
}

func (o UpdateNetworkSmDevicesFieldsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["deviceFields"] = o.DeviceFields
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSmDevicesFieldsRequest struct {
	value *UpdateNetworkSmDevicesFieldsRequest
	isSet bool
}

func (v NullableUpdateNetworkSmDevicesFieldsRequest) Get() *UpdateNetworkSmDevicesFieldsRequest {
	return v.value
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequest) Set(val *UpdateNetworkSmDevicesFieldsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSmDevicesFieldsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSmDevicesFieldsRequest(val *UpdateNetworkSmDevicesFieldsRequest) *NullableUpdateNetworkSmDevicesFieldsRequest {
	return &NullableUpdateNetworkSmDevicesFieldsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSmDevicesFieldsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


