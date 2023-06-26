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

// InlineResponse20092 struct for InlineResponse20092
type InlineResponse20092 struct {
	// The name of the new profile. Must be unique.
	Id *string `json:"id,omitempty"`
	// The network ID of the RF Profile
	NetworkId *string `json:"networkId,omitempty"`
	// The name of the new profile. Must be unique. This param is required on creation.
	Name *string `json:"name,omitempty"`
	// Steers client to best available access point. Can be either true or false. Defaults to true.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	BandSelectionType *string `json:"bandSelectionType,omitempty"`
	ApBandSettings *InlineResponse20092ApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *InlineResponse20092TwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *InlineResponse20092FiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *InlineResponse20092SixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *InlineResponse20092Transmission `json:"transmission,omitempty"`
	PerSsidSettings *InlineResponse20092PerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineResponse20092 instantiates a new InlineResponse20092 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092WithDefaults() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20092) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20092) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20092) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20092) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20092) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20092) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20092) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20092) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20092) SetName(v string) {
	o.Name = &v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *InlineResponse20092) GetClientBalancingEnabled() bool {
	if o == nil || isNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ClientBalancingEnabled) {
    return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *InlineResponse20092) HasClientBalancingEnabled() bool {
	if o != nil && !isNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *InlineResponse20092) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *InlineResponse20092) GetMinBitrateType() string {
	if o == nil || isNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || isNil(o.MinBitrateType) {
    return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *InlineResponse20092) HasMinBitrateType() bool {
	if o != nil && !isNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *InlineResponse20092) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value if set, zero value otherwise.
func (o *InlineResponse20092) GetBandSelectionType() string {
	if o == nil || isNil(o.BandSelectionType) {
		var ret string
		return ret
	}
	return *o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.BandSelectionType) {
    return nil, false
	}
	return o.BandSelectionType, true
}

// HasBandSelectionType returns a boolean if a field has been set.
func (o *InlineResponse20092) HasBandSelectionType() bool {
	if o != nil && !isNil(o.BandSelectionType) {
		return true
	}

	return false
}

// SetBandSelectionType gets a reference to the given string and assigns it to the BandSelectionType field.
func (o *InlineResponse20092) SetBandSelectionType(v string) {
	o.BandSelectionType = &v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *InlineResponse20092) GetApBandSettings() InlineResponse20092ApBandSettings {
	if o == nil || isNil(o.ApBandSettings) {
		var ret InlineResponse20092ApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetApBandSettingsOk() (*InlineResponse20092ApBandSettings, bool) {
	if o == nil || isNil(o.ApBandSettings) {
    return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *InlineResponse20092) HasApBandSettings() bool {
	if o != nil && !isNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given InlineResponse20092ApBandSettings and assigns it to the ApBandSettings field.
func (o *InlineResponse20092) SetApBandSettings(v InlineResponse20092ApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse20092) GetTwoFourGhzSettings() InlineResponse20092TwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret InlineResponse20092TwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetTwoFourGhzSettingsOk() (*InlineResponse20092TwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse20092) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given InlineResponse20092TwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineResponse20092) SetTwoFourGhzSettings(v InlineResponse20092TwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse20092) GetFiveGhzSettings() InlineResponse20092FiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret InlineResponse20092FiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetFiveGhzSettingsOk() (*InlineResponse20092FiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse20092) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given InlineResponse20092FiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineResponse20092) SetFiveGhzSettings(v InlineResponse20092FiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse20092) GetSixGhzSettings() InlineResponse20092SixGhzSettings {
	if o == nil || isNil(o.SixGhzSettings) {
		var ret InlineResponse20092SixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetSixGhzSettingsOk() (*InlineResponse20092SixGhzSettings, bool) {
	if o == nil || isNil(o.SixGhzSettings) {
    return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse20092) HasSixGhzSettings() bool {
	if o != nil && !isNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given InlineResponse20092SixGhzSettings and assigns it to the SixGhzSettings field.
func (o *InlineResponse20092) SetSixGhzSettings(v InlineResponse20092SixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *InlineResponse20092) GetTransmission() InlineResponse20092Transmission {
	if o == nil || isNil(o.Transmission) {
		var ret InlineResponse20092Transmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetTransmissionOk() (*InlineResponse20092Transmission, bool) {
	if o == nil || isNil(o.Transmission) {
    return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *InlineResponse20092) HasTransmission() bool {
	if o != nil && !isNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given InlineResponse20092Transmission and assigns it to the Transmission field.
func (o *InlineResponse20092) SetTransmission(v InlineResponse20092Transmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineResponse20092) GetPerSsidSettings() InlineResponse20092PerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret InlineResponse20092PerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetPerSsidSettingsOk() (*InlineResponse20092PerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineResponse20092) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given InlineResponse20092PerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineResponse20092) SetPerSsidSettings(v InlineResponse20092PerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineResponse20092) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !isNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if !isNil(o.BandSelectionType) {
		toSerialize["bandSelectionType"] = o.BandSelectionType
	}
	if !isNil(o.ApBandSettings) {
		toSerialize["apBandSettings"] = o.ApBandSettings
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !isNil(o.SixGhzSettings) {
		toSerialize["sixGhzSettings"] = o.SixGhzSettings
	}
	if !isNil(o.Transmission) {
		toSerialize["transmission"] = o.Transmission
	}
	if !isNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092 struct {
	value *InlineResponse20092
	isSet bool
}

func (v NullableInlineResponse20092) Get() *InlineResponse20092 {
	return v.value
}

func (v *NullableInlineResponse20092) Set(val *InlineResponse20092) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092(val *InlineResponse20092) *NullableInlineResponse20092 {
	return &NullableInlineResponse20092{value: val, isSet: true}
}

func (v NullableInlineResponse20092) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


