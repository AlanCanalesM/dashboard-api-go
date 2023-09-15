/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner{}

// GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner struct for GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner
type GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner struct {
	// URL to the application
	Url *string `json:"url,omitempty"`
	// Name of the application
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner() *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner struct {
	value *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) Get() *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) Set(val *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner(val *GetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) *NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner {
	return &NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopeApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


