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

// checks if the MoveOrganizationLicensesSeats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicensesSeats200Response{}

// MoveOrganizationLicensesSeats200Response struct for MoveOrganizationLicensesSeats200Response
type MoveOrganizationLicensesSeats200Response struct {
	// The ID of the organization to move the SM seats to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// The ID of the SM license to move the seats from
	LicenseId *string `json:"licenseId,omitempty"`
	// The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
	SeatCount *int32 `json:"seatCount,omitempty"`
}

// NewMoveOrganizationLicensesSeats200Response instantiates a new MoveOrganizationLicensesSeats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensesSeats200Response() *MoveOrganizationLicensesSeats200Response {
	this := MoveOrganizationLicensesSeats200Response{}
	return &this
}

// NewMoveOrganizationLicensesSeats200ResponseWithDefaults instantiates a new MoveOrganizationLicensesSeats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensesSeats200ResponseWithDefaults() *MoveOrganizationLicensesSeats200Response {
	this := MoveOrganizationLicensesSeats200Response{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *MoveOrganizationLicensesSeats200Response) GetDestOrganizationId() string {
	if o == nil || IsNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeats200Response) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestOrganizationId) {
		return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *MoveOrganizationLicensesSeats200Response) HasDestOrganizationId() bool {
	if o != nil && !IsNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *MoveOrganizationLicensesSeats200Response) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *MoveOrganizationLicensesSeats200Response) GetLicenseId() string {
	if o == nil || IsNil(o.LicenseId) {
		var ret string
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeats200Response) GetLicenseIdOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseId) {
		return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *MoveOrganizationLicensesSeats200Response) HasLicenseId() bool {
	if o != nil && !IsNil(o.LicenseId) {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given string and assigns it to the LicenseId field.
func (o *MoveOrganizationLicensesSeats200Response) SetLicenseId(v string) {
	o.LicenseId = &v
}

// GetSeatCount returns the SeatCount field value if set, zero value otherwise.
func (o *MoveOrganizationLicensesSeats200Response) GetSeatCount() int32 {
	if o == nil || IsNil(o.SeatCount) {
		var ret int32
		return ret
	}
	return *o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesSeats200Response) GetSeatCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeatCount) {
		return nil, false
	}
	return o.SeatCount, true
}

// HasSeatCount returns a boolean if a field has been set.
func (o *MoveOrganizationLicensesSeats200Response) HasSeatCount() bool {
	if o != nil && !IsNil(o.SeatCount) {
		return true
	}

	return false
}

// SetSeatCount gets a reference to the given int32 and assigns it to the SeatCount field.
func (o *MoveOrganizationLicensesSeats200Response) SetSeatCount(v int32) {
	o.SeatCount = &v
}

func (o MoveOrganizationLicensesSeats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicensesSeats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !IsNil(o.LicenseId) {
		toSerialize["licenseId"] = o.LicenseId
	}
	if !IsNil(o.SeatCount) {
		toSerialize["seatCount"] = o.SeatCount
	}
	return toSerialize, nil
}

type NullableMoveOrganizationLicensesSeats200Response struct {
	value *MoveOrganizationLicensesSeats200Response
	isSet bool
}

func (v NullableMoveOrganizationLicensesSeats200Response) Get() *MoveOrganizationLicensesSeats200Response {
	return v.value
}

func (v *NullableMoveOrganizationLicensesSeats200Response) Set(val *MoveOrganizationLicensesSeats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensesSeats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensesSeats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensesSeats200Response(val *MoveOrganizationLicensesSeats200Response) *NullableMoveOrganizationLicensesSeats200Response {
	return &NullableMoveOrganizationLicensesSeats200Response{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensesSeats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensesSeats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


