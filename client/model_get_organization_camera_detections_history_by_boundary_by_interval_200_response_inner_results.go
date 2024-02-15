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

// checks if the GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults{}

// GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults The analytics data
type GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults struct {
	// The period start time
	StartTime *string `json:"startTime,omitempty"`
	// The period end time
	EndTime *string `json:"endTime,omitempty"`
	// The detection type
	ObjectType *string `json:"objectType,omitempty"`
	// The number of detections entered
	In *int32 `json:"in,omitempty"`
	// The number of detections exited
	Out *int32 `json:"out,omitempty"`
}

// NewGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults instantiates a new GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults() *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults {
	this := GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults{}
	return &this
}

// NewGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResultsWithDefaults instantiates a new GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResultsWithDefaults() *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults {
	this := GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) SetEndTime(v string) {
	o.EndTime = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetIn() int32 {
	if o == nil || IsNil(o.In) {
		var ret int32
		return ret
	}
	return *o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetInOk() (*int32, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given int32 and assigns it to the In field.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) SetIn(v int32) {
	o.In = &v
}

// GetOut returns the Out field value if set, zero value otherwise.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetOut() int32 {
	if o == nil || IsNil(o.Out) {
		var ret int32
		return ret
	}
	return *o.Out
}

// GetOutOk returns a tuple with the Out field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) GetOutOk() (*int32, bool) {
	if o == nil || IsNil(o.Out) {
		return nil, false
	}
	return o.Out, true
}

// HasOut returns a boolean if a field has been set.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) HasOut() bool {
	if o != nil && !IsNil(o.Out) {
		return true
	}

	return false
}

// SetOut gets a reference to the given int32 and assigns it to the Out field.
func (o *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) SetOut(v int32) {
	o.Out = &v
}

func (o GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !IsNil(o.Out) {
		toSerialize["out"] = o.Out
	}
	return toSerialize, nil
}

type NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults struct {
	value *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults
	isSet bool
}

func (v NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) Get() *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults {
	return v.value
}

func (v *NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) Set(val *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults(val *GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) *NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults {
	return &NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults{value: val, isSet: true}
}

func (v NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInnerResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


