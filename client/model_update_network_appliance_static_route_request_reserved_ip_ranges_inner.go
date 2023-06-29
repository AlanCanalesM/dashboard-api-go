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

// checks if the UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner{}

// UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner struct for UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner
type UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner struct {
	// The first IP in the reserved range
	Start string `json:"start"`
	// The last IP in the reserved range
	End string `json:"end"`
	// A text comment for the reserved range
	Comment string `json:"comment"`
}

// NewUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner instantiates a new UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner(start string, end string, comment string) *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner {
	this := UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner{}
	this.Start = start
	this.End = end
	this.Comment = comment
	return &this
}

// NewUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInnerWithDefaults instantiates a new UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInnerWithDefaults() *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner {
	this := UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner{}
	return &this
}

// GetStart returns the Start field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) SetComment(v string) {
	o.Comment = v
}

func (o UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["comment"] = o.Comment
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner struct {
	value *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) Get() *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) Set(val *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner(val *UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) *NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner {
	return &NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


