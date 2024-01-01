/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkWirelessRfProfiles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessRfProfiles200Response{}

// GetNetworkWirelessRfProfiles200Response struct for GetNetworkWirelessRfProfiles200Response
type GetNetworkWirelessRfProfiles200Response struct {
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
	ApBandSettings *GetNetworkWirelessRfProfiles200ResponseApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *GetNetworkWirelessRfProfiles200ResponseSixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *GetNetworkWirelessRfProfiles200ResponseTransmission `json:"transmission,omitempty"`
	PerSsidSettings *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200Response instantiates a new GetNetworkWirelessRfProfiles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200Response() *GetNetworkWirelessRfProfiles200Response {
	this := GetNetworkWirelessRfProfiles200Response{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponseWithDefaults instantiates a new GetNetworkWirelessRfProfiles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponseWithDefaults() *GetNetworkWirelessRfProfiles200Response {
	this := GetNetworkWirelessRfProfiles200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkWirelessRfProfiles200Response) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetNetworkWirelessRfProfiles200Response) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWirelessRfProfiles200Response) SetName(v string) {
	o.Name = &v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetClientBalancingEnabled() bool {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasClientBalancingEnabled() bool {
	if o != nil && !IsNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *GetNetworkWirelessRfProfiles200Response) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetMinBitrateType() string {
	if o == nil || IsNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MinBitrateType) {
		return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasMinBitrateType() bool {
	if o != nil && !IsNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *GetNetworkWirelessRfProfiles200Response) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetBandSelectionType() string {
	if o == nil || IsNil(o.BandSelectionType) {
		var ret string
		return ret
	}
	return *o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BandSelectionType) {
		return nil, false
	}
	return o.BandSelectionType, true
}

// HasBandSelectionType returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasBandSelectionType() bool {
	if o != nil && !IsNil(o.BandSelectionType) {
		return true
	}

	return false
}

// SetBandSelectionType gets a reference to the given string and assigns it to the BandSelectionType field.
func (o *GetNetworkWirelessRfProfiles200Response) SetBandSelectionType(v string) {
	o.BandSelectionType = &v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetApBandSettings() GetNetworkWirelessRfProfiles200ResponseApBandSettings {
	if o == nil || IsNil(o.ApBandSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetApBandSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseApBandSettings, bool) {
	if o == nil || IsNil(o.ApBandSettings) {
		return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasApBandSettings() bool {
	if o != nil && !IsNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseApBandSettings and assigns it to the ApBandSettings field.
func (o *GetNetworkWirelessRfProfiles200Response) SetApBandSettings(v GetNetworkWirelessRfProfiles200ResponseApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetTwoFourGhzSettings() GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetTwoFourGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *GetNetworkWirelessRfProfiles200Response) SetTwoFourGhzSettings(v GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetFiveGhzSettings() GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetFiveGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *GetNetworkWirelessRfProfiles200Response) SetFiveGhzSettings(v GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetSixGhzSettings() GetNetworkWirelessRfProfiles200ResponseSixGhzSettings {
	if o == nil || IsNil(o.SixGhzSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseSixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetSixGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseSixGhzSettings, bool) {
	if o == nil || IsNil(o.SixGhzSettings) {
		return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasSixGhzSettings() bool {
	if o != nil && !IsNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseSixGhzSettings and assigns it to the SixGhzSettings field.
func (o *GetNetworkWirelessRfProfiles200Response) SetSixGhzSettings(v GetNetworkWirelessRfProfiles200ResponseSixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetTransmission() GetNetworkWirelessRfProfiles200ResponseTransmission {
	if o == nil || IsNil(o.Transmission) {
		var ret GetNetworkWirelessRfProfiles200ResponseTransmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetTransmissionOk() (*GetNetworkWirelessRfProfiles200ResponseTransmission, bool) {
	if o == nil || IsNil(o.Transmission) {
		return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasTransmission() bool {
	if o != nil && !IsNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given GetNetworkWirelessRfProfiles200ResponseTransmission and assigns it to the Transmission field.
func (o *GetNetworkWirelessRfProfiles200Response) SetTransmission(v GetNetworkWirelessRfProfiles200ResponseTransmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200Response) GetPerSsidSettings() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings {
	if o == nil || IsNil(o.PerSsidSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200Response) GetPerSsidSettingsOk() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings, bool) {
	if o == nil || IsNil(o.PerSsidSettings) {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200Response) HasPerSsidSettings() bool {
	if o != nil && !IsNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings and assigns it to the PerSsidSettings field.
func (o *GetNetworkWirelessRfProfiles200Response) SetPerSsidSettings(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o GetNetworkWirelessRfProfiles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessRfProfiles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !IsNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if !IsNil(o.BandSelectionType) {
		toSerialize["bandSelectionType"] = o.BandSelectionType
	}
	if !IsNil(o.ApBandSettings) {
		toSerialize["apBandSettings"] = o.ApBandSettings
	}
	if !IsNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !IsNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !IsNil(o.SixGhzSettings) {
		toSerialize["sixGhzSettings"] = o.SixGhzSettings
	}
	if !IsNil(o.Transmission) {
		toSerialize["transmission"] = o.Transmission
	}
	if !IsNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessRfProfiles200Response struct {
	value *GetNetworkWirelessRfProfiles200Response
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200Response) Get() *GetNetworkWirelessRfProfiles200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200Response) Set(val *GetNetworkWirelessRfProfiles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200Response(val *GetNetworkWirelessRfProfiles200Response) *NullableGetNetworkWirelessRfProfiles200Response {
	return &NullableGetNetworkWirelessRfProfiles200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


