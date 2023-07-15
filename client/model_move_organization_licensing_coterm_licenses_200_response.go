/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MoveOrganizationLicensingCotermLicenses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicensingCotermLicenses200Response{}

// MoveOrganizationLicensingCotermLicenses200Response struct for MoveOrganizationLicensingCotermLicenses200Response
type MoveOrganizationLicensingCotermLicenses200Response struct {
	// Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
	RemainderLicenses []GetOrganizationLicensingCotermLicenses200ResponseInner `json:"remainderLicenses,omitempty"`
	// Newly moved licenses created in the destination organization of the license move operation
	MovedLicenses []GetOrganizationLicensingCotermLicenses200ResponseInner `json:"movedLicenses,omitempty"`
}

// NewMoveOrganizationLicensingCotermLicenses200Response instantiates a new MoveOrganizationLicensingCotermLicenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensingCotermLicenses200Response() *MoveOrganizationLicensingCotermLicenses200Response {
	this := MoveOrganizationLicensingCotermLicenses200Response{}
	return &this
}

// NewMoveOrganizationLicensingCotermLicenses200ResponseWithDefaults instantiates a new MoveOrganizationLicensingCotermLicenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensingCotermLicenses200ResponseWithDefaults() *MoveOrganizationLicensingCotermLicenses200Response {
	this := MoveOrganizationLicensingCotermLicenses200Response{}
	return &this
}

// GetRemainderLicenses returns the RemainderLicenses field value if set, zero value otherwise.
func (o *MoveOrganizationLicensingCotermLicenses200Response) GetRemainderLicenses() []GetOrganizationLicensingCotermLicenses200ResponseInner {
	if o == nil || IsNil(o.RemainderLicenses) {
		var ret []GetOrganizationLicensingCotermLicenses200ResponseInner
		return ret
	}
	return o.RemainderLicenses
}

// GetRemainderLicensesOk returns a tuple with the RemainderLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicenses200Response) GetRemainderLicensesOk() ([]GetOrganizationLicensingCotermLicenses200ResponseInner, bool) {
	if o == nil || IsNil(o.RemainderLicenses) {
		return nil, false
	}
	return o.RemainderLicenses, true
}

// HasRemainderLicenses returns a boolean if a field has been set.
func (o *MoveOrganizationLicensingCotermLicenses200Response) HasRemainderLicenses() bool {
	if o != nil && !IsNil(o.RemainderLicenses) {
		return true
	}

	return false
}

// SetRemainderLicenses gets a reference to the given []GetOrganizationLicensingCotermLicenses200ResponseInner and assigns it to the RemainderLicenses field.
func (o *MoveOrganizationLicensingCotermLicenses200Response) SetRemainderLicenses(v []GetOrganizationLicensingCotermLicenses200ResponseInner) {
	o.RemainderLicenses = v
}

// GetMovedLicenses returns the MovedLicenses field value if set, zero value otherwise.
func (o *MoveOrganizationLicensingCotermLicenses200Response) GetMovedLicenses() []GetOrganizationLicensingCotermLicenses200ResponseInner {
	if o == nil || IsNil(o.MovedLicenses) {
		var ret []GetOrganizationLicensingCotermLicenses200ResponseInner
		return ret
	}
	return o.MovedLicenses
}

// GetMovedLicensesOk returns a tuple with the MovedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicenses200Response) GetMovedLicensesOk() ([]GetOrganizationLicensingCotermLicenses200ResponseInner, bool) {
	if o == nil || IsNil(o.MovedLicenses) {
		return nil, false
	}
	return o.MovedLicenses, true
}

// HasMovedLicenses returns a boolean if a field has been set.
func (o *MoveOrganizationLicensingCotermLicenses200Response) HasMovedLicenses() bool {
	if o != nil && !IsNil(o.MovedLicenses) {
		return true
	}

	return false
}

// SetMovedLicenses gets a reference to the given []GetOrganizationLicensingCotermLicenses200ResponseInner and assigns it to the MovedLicenses field.
func (o *MoveOrganizationLicensingCotermLicenses200Response) SetMovedLicenses(v []GetOrganizationLicensingCotermLicenses200ResponseInner) {
	o.MovedLicenses = v
}

func (o MoveOrganizationLicensingCotermLicenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicensingCotermLicenses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemainderLicenses) {
		toSerialize["remainderLicenses"] = o.RemainderLicenses
	}
	if !IsNil(o.MovedLicenses) {
		toSerialize["movedLicenses"] = o.MovedLicenses
	}
	return toSerialize, nil
}

type NullableMoveOrganizationLicensingCotermLicenses200Response struct {
	value *MoveOrganizationLicensingCotermLicenses200Response
	isSet bool
}

func (v NullableMoveOrganizationLicensingCotermLicenses200Response) Get() *MoveOrganizationLicensingCotermLicenses200Response {
	return v.value
}

func (v *NullableMoveOrganizationLicensingCotermLicenses200Response) Set(val *MoveOrganizationLicensingCotermLicenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensingCotermLicenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensingCotermLicenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensingCotermLicenses200Response(val *MoveOrganizationLicensingCotermLicenses200Response) *NullableMoveOrganizationLicensingCotermLicenses200Response {
	return &NullableMoveOrganizationLicensingCotermLicenses200Response{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensingCotermLicenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensingCotermLicenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


