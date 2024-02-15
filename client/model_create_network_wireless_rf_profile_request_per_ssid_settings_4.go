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

// checks if the CreateNetworkWirelessRfProfileRequestPerSsidSettings4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestPerSsidSettings4{}

// CreateNetworkWirelessRfProfileRequestPerSsidSettings4 Settings for SSID 4
type CreateNetworkWirelessRfProfileRequestPerSsidSettings4 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *CreateNetworkWirelessRfProfileRequestApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings4 instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings4() *CreateNetworkWirelessRfProfileRequestPerSsidSettings4 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings4{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings4WithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings4WithDefaults() *CreateNetworkWirelessRfProfileRequestPerSsidSettings4 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings4{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetMinBitrate() float32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetMinBitrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBands() CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	if o == nil || IsNil(o.Bands) {
		var ret CreateNetworkWirelessRfProfileRequestApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBandsOk() (*CreateNetworkWirelessRfProfileRequestApBandSettingsBands, bool) {
	if o == nil || IsNil(o.Bands) {
		return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) HasBands() bool {
	if o != nil && !IsNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given CreateNetworkWirelessRfProfileRequestApBandSettingsBands and assigns it to the Bands field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) SetBands(v CreateNetworkWirelessRfProfileRequestApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4 struct {
	value *CreateNetworkWirelessRfProfileRequestPerSsidSettings4
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) Get() *CreateNetworkWirelessRfProfileRequestPerSsidSettings4 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) Set(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings4) *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4 {
	return &NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


