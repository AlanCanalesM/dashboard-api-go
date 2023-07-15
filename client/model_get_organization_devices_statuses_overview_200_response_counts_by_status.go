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

// checks if the GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus{}

// GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus byStatus
type GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus struct {
	// online count
	Online *int32 `json:"online,omitempty"`
	// alerting count
	Alerting *int32 `json:"alerting,omitempty"`
	// offline count
	Offline *int32 `json:"offline,omitempty"`
	// dormant count
	Dormant *int32 `json:"dormant,omitempty"`
}

// NewGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus instantiates a new GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus() *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus {
	this := GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus{}
	return &this
}

// NewGetOrganizationDevicesStatusesOverview200ResponseCountsByStatusWithDefaults instantiates a new GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesStatusesOverview200ResponseCountsByStatusWithDefaults() *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus {
	this := GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus{}
	return &this
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetOnline() int32 {
	if o == nil || IsNil(o.Online) {
		var ret int32
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetOnlineOk() (*int32, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) HasOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given int32 and assigns it to the Online field.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) SetOnline(v int32) {
	o.Online = &v
}

// GetAlerting returns the Alerting field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetAlerting() int32 {
	if o == nil || IsNil(o.Alerting) {
		var ret int32
		return ret
	}
	return *o.Alerting
}

// GetAlertingOk returns a tuple with the Alerting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetAlertingOk() (*int32, bool) {
	if o == nil || IsNil(o.Alerting) {
		return nil, false
	}
	return o.Alerting, true
}

// HasAlerting returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) HasAlerting() bool {
	if o != nil && !IsNil(o.Alerting) {
		return true
	}

	return false
}

// SetAlerting gets a reference to the given int32 and assigns it to the Alerting field.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) SetAlerting(v int32) {
	o.Alerting = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetOffline() int32 {
	if o == nil || IsNil(o.Offline) {
		var ret int32
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetOfflineOk() (*int32, bool) {
	if o == nil || IsNil(o.Offline) {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) HasOffline() bool {
	if o != nil && !IsNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given int32 and assigns it to the Offline field.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) SetOffline(v int32) {
	o.Offline = &v
}

// GetDormant returns the Dormant field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetDormant() int32 {
	if o == nil || IsNil(o.Dormant) {
		var ret int32
		return ret
	}
	return *o.Dormant
}

// GetDormantOk returns a tuple with the Dormant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) GetDormantOk() (*int32, bool) {
	if o == nil || IsNil(o.Dormant) {
		return nil, false
	}
	return o.Dormant, true
}

// HasDormant returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) HasDormant() bool {
	if o != nil && !IsNil(o.Dormant) {
		return true
	}

	return false
}

// SetDormant gets a reference to the given int32 and assigns it to the Dormant field.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) SetDormant(v int32) {
	o.Dormant = &v
}

func (o GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Alerting) {
		toSerialize["alerting"] = o.Alerting
	}
	if !IsNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !IsNil(o.Dormant) {
		toSerialize["dormant"] = o.Dormant
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus struct {
	value *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus
	isSet bool
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) Get() *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus {
	return v.value
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) Set(val *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus(val *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) *NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus {
	return &NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


