/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSmDeviceWlanLists200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDeviceWlanLists200ResponseInner{}

// GetNetworkSmDeviceWlanLists200ResponseInner struct for GetNetworkSmDeviceWlanLists200ResponseInner
type GetNetworkSmDeviceWlanLists200ResponseInner struct {
	// When the Meraki record for the wlanList was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Meraki managed Id of the wlanList record.
	Id *string `json:"id,omitempty"`
	// An XML string containing the WLAN List for the device.
	Xml *string `json:"xml,omitempty"`
}

// NewGetNetworkSmDeviceWlanLists200ResponseInner instantiates a new GetNetworkSmDeviceWlanLists200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceWlanLists200ResponseInner() *GetNetworkSmDeviceWlanLists200ResponseInner {
	this := GetNetworkSmDeviceWlanLists200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceWlanLists200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceWlanLists200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceWlanLists200ResponseInnerWithDefaults() *GetNetworkSmDeviceWlanLists200ResponseInner {
	this := GetNetworkSmDeviceWlanLists200ResponseInner{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetXml() string {
	if o == nil || IsNil(o.Xml) {
		var ret string
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) GetXmlOk() (*string, bool) {
	if o == nil || IsNil(o.Xml) {
		return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) HasXml() bool {
	if o != nil && !IsNil(o.Xml) {
		return true
	}

	return false
}

// SetXml gets a reference to the given string and assigns it to the Xml field.
func (o *GetNetworkSmDeviceWlanLists200ResponseInner) SetXml(v string) {
	o.Xml = &v
}

func (o GetNetworkSmDeviceWlanLists200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDeviceWlanLists200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Xml) {
		toSerialize["xml"] = o.Xml
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDeviceWlanLists200ResponseInner struct {
	value *GetNetworkSmDeviceWlanLists200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceWlanLists200ResponseInner) Get() *GetNetworkSmDeviceWlanLists200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceWlanLists200ResponseInner) Set(val *GetNetworkSmDeviceWlanLists200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceWlanLists200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceWlanLists200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceWlanLists200ResponseInner(val *GetNetworkSmDeviceWlanLists200ResponseInner) *NullableGetNetworkSmDeviceWlanLists200ResponseInner {
	return &NullableGetNetworkSmDeviceWlanLists200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceWlanLists200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceWlanLists200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


