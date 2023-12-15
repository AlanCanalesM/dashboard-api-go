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

// checks if the UpdateNetworkCellularGatewaySubnetPoolRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkCellularGatewaySubnetPoolRequest{}

// UpdateNetworkCellularGatewaySubnetPoolRequest struct for UpdateNetworkCellularGatewaySubnetPoolRequest
type UpdateNetworkCellularGatewaySubnetPoolRequest struct {
	// Mask used for the subnet of all MGs in  this network.
	Mask *int32 `json:"mask,omitempty"`
	// CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Cidr *string `json:"cidr,omitempty"`
}

// NewUpdateNetworkCellularGatewaySubnetPoolRequest instantiates a new UpdateNetworkCellularGatewaySubnetPoolRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkCellularGatewaySubnetPoolRequest() *UpdateNetworkCellularGatewaySubnetPoolRequest {
	this := UpdateNetworkCellularGatewaySubnetPoolRequest{}
	return &this
}

// NewUpdateNetworkCellularGatewaySubnetPoolRequestWithDefaults instantiates a new UpdateNetworkCellularGatewaySubnetPoolRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkCellularGatewaySubnetPoolRequestWithDefaults() *UpdateNetworkCellularGatewaySubnetPoolRequest {
	this := UpdateNetworkCellularGatewaySubnetPoolRequest{}
	return &this
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) GetMask() int32 {
	if o == nil || IsNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) GetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) SetMask(v int32) {
	o.Mask = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *UpdateNetworkCellularGatewaySubnetPoolRequest) SetCidr(v string) {
	o.Cidr = &v
}

func (o UpdateNetworkCellularGatewaySubnetPoolRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkCellularGatewaySubnetPoolRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return toSerialize, nil
}

type NullableUpdateNetworkCellularGatewaySubnetPoolRequest struct {
	value *UpdateNetworkCellularGatewaySubnetPoolRequest
	isSet bool
}

func (v NullableUpdateNetworkCellularGatewaySubnetPoolRequest) Get() *UpdateNetworkCellularGatewaySubnetPoolRequest {
	return v.value
}

func (v *NullableUpdateNetworkCellularGatewaySubnetPoolRequest) Set(val *UpdateNetworkCellularGatewaySubnetPoolRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkCellularGatewaySubnetPoolRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkCellularGatewaySubnetPoolRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkCellularGatewaySubnetPoolRequest(val *UpdateNetworkCellularGatewaySubnetPoolRequest) *NullableUpdateNetworkCellularGatewaySubnetPoolRequest {
	return &NullableUpdateNetworkCellularGatewaySubnetPoolRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkCellularGatewaySubnetPoolRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkCellularGatewaySubnetPoolRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


