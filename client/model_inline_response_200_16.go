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

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	// The subnet of the single LAN
	Subnet *string `json:"subnet,omitempty"`
	// The local IP of the appliance on the single LAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	MandatoryDhcp *InlineResponse20016MandatoryDhcp `json:"mandatoryDhcp,omitempty"`
	Ipv6 *InlineResponse20016Ipv6 `json:"ipv6,omitempty"`
}

// NewInlineResponse20016 instantiates a new InlineResponse20016 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016WithDefaults() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20016) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20016) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20016) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *InlineResponse20016) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *InlineResponse20016) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *InlineResponse20016) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetMandatoryDhcp returns the MandatoryDhcp field value if set, zero value otherwise.
func (o *InlineResponse20016) GetMandatoryDhcp() InlineResponse20016MandatoryDhcp {
	if o == nil || isNil(o.MandatoryDhcp) {
		var ret InlineResponse20016MandatoryDhcp
		return ret
	}
	return *o.MandatoryDhcp
}

// GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetMandatoryDhcpOk() (*InlineResponse20016MandatoryDhcp, bool) {
	if o == nil || isNil(o.MandatoryDhcp) {
    return nil, false
	}
	return o.MandatoryDhcp, true
}

// HasMandatoryDhcp returns a boolean if a field has been set.
func (o *InlineResponse20016) HasMandatoryDhcp() bool {
	if o != nil && !isNil(o.MandatoryDhcp) {
		return true
	}

	return false
}

// SetMandatoryDhcp gets a reference to the given InlineResponse20016MandatoryDhcp and assigns it to the MandatoryDhcp field.
func (o *InlineResponse20016) SetMandatoryDhcp(v InlineResponse20016MandatoryDhcp) {
	o.MandatoryDhcp = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineResponse20016) GetIpv6() InlineResponse20016Ipv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret InlineResponse20016Ipv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetIpv6Ok() (*InlineResponse20016Ipv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineResponse20016) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given InlineResponse20016Ipv6 and assigns it to the Ipv6 field.
func (o *InlineResponse20016) SetIpv6(v InlineResponse20016Ipv6) {
	o.Ipv6 = &v
}

func (o InlineResponse20016) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.MandatoryDhcp) {
		toSerialize["mandatoryDhcp"] = o.MandatoryDhcp
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016 struct {
	value *InlineResponse20016
	isSet bool
}

func (v NullableInlineResponse20016) Get() *InlineResponse20016 {
	return v.value
}

func (v *NullableInlineResponse20016) Set(val *InlineResponse20016) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016(val *InlineResponse20016) *NullableInlineResponse20016 {
	return &NullableInlineResponse20016{value: val, isSet: true}
}

func (v NullableInlineResponse20016) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


