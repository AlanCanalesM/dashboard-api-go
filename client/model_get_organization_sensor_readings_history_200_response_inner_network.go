/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationSensorReadingsHistory200ResponseInnerNetwork Network to which the sensor belongs.
type GetOrganizationSensorReadingsHistory200ResponseInnerNetwork struct {
	// ID of the network.
	Id *string `json:"id,omitempty"`
	// Name of the network.
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNetwork instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNetwork() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNetworkWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork(val *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


