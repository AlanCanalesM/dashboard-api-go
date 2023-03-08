/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation Link negotiation details object for the port
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation struct {
	// The duplex mode of the port. Can be 'full' or 'half'
	Duplex *string `json:"duplex,omitempty"`
	// The speed of the port
	Speed *int32 `json:"speed,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiationWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiationWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation{}
	return &this
}

// GetDuplex returns the Duplex field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) GetDuplex() string {
	if o == nil || isNil(o.Duplex) {
		var ret string
		return ret
	}
	return *o.Duplex
}

// GetDuplexOk returns a tuple with the Duplex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) GetDuplexOk() (*string, bool) {
	if o == nil || isNil(o.Duplex) {
    return nil, false
	}
	return o.Duplex, true
}

// HasDuplex returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) HasDuplex() bool {
	if o != nil && !isNil(o.Duplex) {
		return true
	}

	return false
}

// SetDuplex gets a reference to the given string and assigns it to the Duplex field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) SetDuplex(v string) {
	o.Duplex = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) GetSpeed() int32 {
	if o == nil || isNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) GetSpeedOk() (*int32, bool) {
	if o == nil || isNil(o.Speed) {
    return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) HasSpeed() bool {
	if o != nil && !isNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) SetSpeed(v int32) {
	o.Speed = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Duplex) {
		toSerialize["duplex"] = o.Duplex
	}
	if !isNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


