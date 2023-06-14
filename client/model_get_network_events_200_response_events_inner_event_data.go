/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkEvents200ResponseEventsInnerEventData An object containing more data related to the event.
type GetNetworkEvents200ResponseEventsInnerEventData struct {
	// The radio band number the client is trying to connect to.
	Radio *string `json:"radio,omitempty"`
	// The virtual access point (VAP) number the client is connecting to.
	Vap *string `json:"vap,omitempty"`
	// The client's MAC address
	ClientMac *string `json:"client_mac,omitempty"`
	// The client's IP address
	ClientIp *string `json:"client_ip,omitempty"`
	// The radio channel the client is connecting to.
	Channel *string `json:"channel,omitempty"`
	// The current received signal strength indication (RSSI) of the client connected to an AP.
	Rssi *string `json:"rssi,omitempty"`
	// The association ID of the client.
	Aid *string `json:"aid,omitempty"`
}

// NewGetNetworkEvents200ResponseEventsInnerEventData instantiates a new GetNetworkEvents200ResponseEventsInnerEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEvents200ResponseEventsInnerEventData() *GetNetworkEvents200ResponseEventsInnerEventData {
	this := GetNetworkEvents200ResponseEventsInnerEventData{}
	return &this
}

// NewGetNetworkEvents200ResponseEventsInnerEventDataWithDefaults instantiates a new GetNetworkEvents200ResponseEventsInnerEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEvents200ResponseEventsInnerEventDataWithDefaults() *GetNetworkEvents200ResponseEventsInnerEventData {
	this := GetNetworkEvents200ResponseEventsInnerEventData{}
	return &this
}

// GetRadio returns the Radio field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetRadio() string {
	if o == nil || isNil(o.Radio) {
		var ret string
		return ret
	}
	return *o.Radio
}

// GetRadioOk returns a tuple with the Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetRadioOk() (*string, bool) {
	if o == nil || isNil(o.Radio) {
    return nil, false
	}
	return o.Radio, true
}

// HasRadio returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasRadio() bool {
	if o != nil && !isNil(o.Radio) {
		return true
	}

	return false
}

// SetRadio gets a reference to the given string and assigns it to the Radio field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetRadio(v string) {
	o.Radio = &v
}

// GetVap returns the Vap field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetVap() string {
	if o == nil || isNil(o.Vap) {
		var ret string
		return ret
	}
	return *o.Vap
}

// GetVapOk returns a tuple with the Vap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetVapOk() (*string, bool) {
	if o == nil || isNil(o.Vap) {
    return nil, false
	}
	return o.Vap, true
}

// HasVap returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasVap() bool {
	if o != nil && !isNil(o.Vap) {
		return true
	}

	return false
}

// SetVap gets a reference to the given string and assigns it to the Vap field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetVap(v string) {
	o.Vap = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetClientMac() string {
	if o == nil || isNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetClientMacOk() (*string, bool) {
	if o == nil || isNil(o.ClientMac) {
    return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasClientMac() bool {
	if o != nil && !isNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetClientIp() string {
	if o == nil || isNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetClientIpOk() (*string, bool) {
	if o == nil || isNil(o.ClientIp) {
    return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasClientIp() bool {
	if o != nil && !isNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetChannel() string {
	if o == nil || isNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetChannelOk() (*string, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetChannel(v string) {
	o.Channel = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetRssi() string {
	if o == nil || isNil(o.Rssi) {
		var ret string
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetRssiOk() (*string, bool) {
	if o == nil || isNil(o.Rssi) {
    return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasRssi() bool {
	if o != nil && !isNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given string and assigns it to the Rssi field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetRssi(v string) {
	o.Rssi = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetAid() string {
	if o == nil || isNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) GetAidOk() (*string, bool) {
	if o == nil || isNil(o.Aid) {
    return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) HasAid() bool {
	if o != nil && !isNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *GetNetworkEvents200ResponseEventsInnerEventData) SetAid(v string) {
	o.Aid = &v
}

func (o GetNetworkEvents200ResponseEventsInnerEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Radio) {
		toSerialize["radio"] = o.Radio
	}
	if !isNil(o.Vap) {
		toSerialize["vap"] = o.Vap
	}
	if !isNil(o.ClientMac) {
		toSerialize["client_mac"] = o.ClientMac
	}
	if !isNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	if !isNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkEvents200ResponseEventsInnerEventData struct {
	value *GetNetworkEvents200ResponseEventsInnerEventData
	isSet bool
}

func (v NullableGetNetworkEvents200ResponseEventsInnerEventData) Get() *GetNetworkEvents200ResponseEventsInnerEventData {
	return v.value
}

func (v *NullableGetNetworkEvents200ResponseEventsInnerEventData) Set(val *GetNetworkEvents200ResponseEventsInnerEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEvents200ResponseEventsInnerEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEvents200ResponseEventsInnerEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEvents200ResponseEventsInnerEventData(val *GetNetworkEvents200ResponseEventsInnerEventData) *NullableGetNetworkEvents200ResponseEventsInnerEventData {
	return &NullableGetNetworkEvents200ResponseEventsInnerEventData{value: val, isSet: true}
}

func (v NullableGetNetworkEvents200ResponseEventsInnerEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEvents200ResponseEventsInnerEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


