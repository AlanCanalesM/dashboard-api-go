/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	// time when the event occurred
	OccurredAt *string `json:"occurredAt,omitempty"`
	// type of alert
	AlertTypeId *string `json:"alertTypeId,omitempty"`
	// user friendly alert type
	AlertType *string `json:"alertType,omitempty"`
	Device *NetworksNetworkIdAlertsHistoryDevice `json:"device,omitempty"`
	Destinations *NetworksNetworkIdAlertsHistoryDestinations `json:"destinations,omitempty"`
	// relevant data about the event that caused the alert
	AlertData map[string]interface{} `json:"alertData,omitempty"`
}

// NewInlineResponse20013 instantiates a new InlineResponse20013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013WithDefaults() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *InlineResponse20013) GetOccurredAt() string {
	if o == nil || isNil(o.OccurredAt) {
		var ret string
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetOccurredAtOk() (*string, bool) {
	if o == nil || isNil(o.OccurredAt) {
    return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *InlineResponse20013) HasOccurredAt() bool {
	if o != nil && !isNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given string and assigns it to the OccurredAt field.
func (o *InlineResponse20013) SetOccurredAt(v string) {
	o.OccurredAt = &v
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *InlineResponse20013) GetAlertTypeId() string {
	if o == nil || isNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.AlertTypeId) {
    return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *InlineResponse20013) HasAlertTypeId() bool {
	if o != nil && !isNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *InlineResponse20013) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *InlineResponse20013) GetAlertType() string {
	if o == nil || isNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetAlertTypeOk() (*string, bool) {
	if o == nil || isNil(o.AlertType) {
    return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *InlineResponse20013) HasAlertType() bool {
	if o != nil && !isNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *InlineResponse20013) SetAlertType(v string) {
	o.AlertType = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse20013) GetDevice() NetworksNetworkIdAlertsHistoryDevice {
	if o == nil || isNil(o.Device) {
		var ret NetworksNetworkIdAlertsHistoryDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetDeviceOk() (*NetworksNetworkIdAlertsHistoryDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse20013) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NetworksNetworkIdAlertsHistoryDevice and assigns it to the Device field.
func (o *InlineResponse20013) SetDevice(v NetworksNetworkIdAlertsHistoryDevice) {
	o.Device = &v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineResponse20013) GetDestinations() NetworksNetworkIdAlertsHistoryDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret NetworksNetworkIdAlertsHistoryDestinations
		return ret
	}
	return *o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetDestinationsOk() (*NetworksNetworkIdAlertsHistoryDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineResponse20013) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given NetworksNetworkIdAlertsHistoryDestinations and assigns it to the Destinations field.
func (o *InlineResponse20013) SetDestinations(v NetworksNetworkIdAlertsHistoryDestinations) {
	o.Destinations = &v
}

// GetAlertData returns the AlertData field value if set, zero value otherwise.
func (o *InlineResponse20013) GetAlertData() map[string]interface{} {
	if o == nil || isNil(o.AlertData) {
		var ret map[string]interface{}
		return ret
	}
	return o.AlertData
}

// GetAlertDataOk returns a tuple with the AlertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetAlertDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AlertData) {
    return map[string]interface{}{}, false
	}
	return o.AlertData, true
}

// HasAlertData returns a boolean if a field has been set.
func (o *InlineResponse20013) HasAlertData() bool {
	if o != nil && !isNil(o.AlertData) {
		return true
	}

	return false
}

// SetAlertData gets a reference to the given map[string]interface{} and assigns it to the AlertData field.
func (o *InlineResponse20013) SetAlertData(v map[string]interface{}) {
	o.AlertData = v
}

func (o InlineResponse20013) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OccurredAt) {
		toSerialize["occurredAt"] = o.OccurredAt
	}
	if !isNil(o.AlertTypeId) {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	if !isNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	if !isNil(o.AlertData) {
		toSerialize["alertData"] = o.AlertData
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20013 struct {
	value *InlineResponse20013
	isSet bool
}

func (v NullableInlineResponse20013) Get() *InlineResponse20013 {
	return v.value
}

func (v *NullableInlineResponse20013) Set(val *InlineResponse20013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013(val *InlineResponse20013) *NullableInlineResponse20013 {
	return &NullableInlineResponse20013{value: val, isSet: true}
}

func (v NullableInlineResponse20013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


