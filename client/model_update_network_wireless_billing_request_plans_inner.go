/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessBillingRequestPlansInner struct for UpdateNetworkWirelessBillingRequestPlansInner
type UpdateNetworkWirelessBillingRequestPlansInner struct {
	// The id of the pricing plan to update.
	Id *string `json:"id,omitempty"`
	// The price of the billing plan.
	Price float32 `json:"price"`
	BandwidthLimits UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits `json:"bandwidthLimits"`
	// The time limit of the pricing plan in minutes. Can be '1 hour', '1 day', '1 week', or '30 days'.
	TimeLimit string `json:"timeLimit"`
}

// NewUpdateNetworkWirelessBillingRequestPlansInner instantiates a new UpdateNetworkWirelessBillingRequestPlansInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessBillingRequestPlansInner(price float32, bandwidthLimits UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits, timeLimit string) *UpdateNetworkWirelessBillingRequestPlansInner {
	this := UpdateNetworkWirelessBillingRequestPlansInner{}
	this.Price = price
	this.BandwidthLimits = bandwidthLimits
	this.TimeLimit = timeLimit
	return &this
}

// NewUpdateNetworkWirelessBillingRequestPlansInnerWithDefaults instantiates a new UpdateNetworkWirelessBillingRequestPlansInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessBillingRequestPlansInnerWithDefaults() *UpdateNetworkWirelessBillingRequestPlansInner {
	this := UpdateNetworkWirelessBillingRequestPlansInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) SetId(v string) {
	o.Id = &v
}

// GetPrice returns the Price field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetPriceOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) SetPrice(v float32) {
	o.Price = v
}

// GetBandwidthLimits returns the BandwidthLimits field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetBandwidthLimits() UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits {
	if o == nil {
		var ret UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits
		return ret
	}

	return o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetBandwidthLimitsOk() (*UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BandwidthLimits, true
}

// SetBandwidthLimits sets field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) SetBandwidthLimits(v UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) {
	o.BandwidthLimits = v
}

// GetTimeLimit returns the TimeLimit field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetTimeLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeLimit
}

// GetTimeLimitOk returns a tuple with the TimeLimit field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInner) GetTimeLimitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeLimit, true
}

// SetTimeLimit sets field value
func (o *UpdateNetworkWirelessBillingRequestPlansInner) SetTimeLimit(v string) {
	o.TimeLimit = v
}

func (o UpdateNetworkWirelessBillingRequestPlansInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	if true {
		toSerialize["timeLimit"] = o.TimeLimit
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessBillingRequestPlansInner struct {
	value *UpdateNetworkWirelessBillingRequestPlansInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInner) Get() *UpdateNetworkWirelessBillingRequestPlansInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInner) Set(val *UpdateNetworkWirelessBillingRequestPlansInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessBillingRequestPlansInner(val *UpdateNetworkWirelessBillingRequestPlansInner) *NullableUpdateNetworkWirelessBillingRequestPlansInner {
	return &NullableUpdateNetworkWirelessBillingRequestPlansInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


