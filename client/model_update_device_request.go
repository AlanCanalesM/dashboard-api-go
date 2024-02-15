/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceRequest{}

// UpdateDeviceRequest struct for UpdateDeviceRequest
type UpdateDeviceRequest struct {
	// The name of a device
	Name *string `json:"name,omitempty"`
	// The list of tags of a device
	Tags []string `json:"tags,omitempty"`
	// The latitude of a device
	Lat *float32 `json:"lat,omitempty"`
	// The longitude of a device
	Lng *float32 `json:"lng,omitempty"`
	// The address of a device
	Address *string `json:"address,omitempty"`
	// The notes for the device. String. Limited to 255 characters.
	Notes *string `json:"notes,omitempty"`
	// Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	MoveMapMarker *bool `json:"moveMapMarker,omitempty"`
	// The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanId *string `json:"floorPlanId,omitempty"`
}

// NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceRequest() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDeviceRequest) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateDeviceRequest) SetTags(v []string) {
	o.Tags = v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *UpdateDeviceRequest) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetLng() float32 {
	if o == nil || IsNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetLngOk() (*float32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *UpdateDeviceRequest) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateDeviceRequest) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateDeviceRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetMoveMapMarker returns the MoveMapMarker field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetMoveMapMarker() bool {
	if o == nil || IsNil(o.MoveMapMarker) {
		var ret bool
		return ret
	}
	return *o.MoveMapMarker
}

// GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetMoveMapMarkerOk() (*bool, bool) {
	if o == nil || IsNil(o.MoveMapMarker) {
		return nil, false
	}
	return o.MoveMapMarker, true
}

// HasMoveMapMarker returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasMoveMapMarker() bool {
	if o != nil && !IsNil(o.MoveMapMarker) {
		return true
	}

	return false
}

// SetMoveMapMarker gets a reference to the given bool and assigns it to the MoveMapMarker field.
func (o *UpdateDeviceRequest) SetMoveMapMarker(v bool) {
	o.MoveMapMarker = &v
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetSwitchProfileId() string {
	if o == nil || IsNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchProfileId) {
		return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasSwitchProfileId() bool {
	if o != nil && !IsNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *UpdateDeviceRequest) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetFloorPlanId() string {
	if o == nil || IsNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.FloorPlanId) {
		return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasFloorPlanId() bool {
	if o != nil && !IsNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *UpdateDeviceRequest) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

func (o UpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.MoveMapMarker) {
		toSerialize["moveMapMarker"] = o.MoveMapMarker
	}
	if !IsNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !IsNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	return toSerialize, nil
}

type NullableUpdateDeviceRequest struct {
	value *UpdateDeviceRequest
	isSet bool
}

func (v NullableUpdateDeviceRequest) Get() *UpdateDeviceRequest {
	return v.value
}

func (v *NullableUpdateDeviceRequest) Set(val *UpdateDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceRequest(val *UpdateDeviceRequest) *NullableUpdateDeviceRequest {
	return &NullableUpdateDeviceRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


