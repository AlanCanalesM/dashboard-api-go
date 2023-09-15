/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidHotspot20Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidHotspot20Request{}

// UpdateNetworkWirelessSsidHotspot20Request struct for UpdateNetworkWirelessSsidHotspot20Request
type UpdateNetworkWirelessSsidHotspot20Request struct {
	// Whether or not Hotspot 2.0 for this SSID is enabled
	Enabled *bool `json:"enabled,omitempty"`
	Operator *UpdateNetworkWirelessSsidHotspot20RequestOperator `json:"operator,omitempty"`
	Venue *UpdateNetworkWirelessSsidHotspot20RequestVenue `json:"venue,omitempty"`
	// The network type of this SSID ('Private network', 'Private network with guest access', 'Chargeable public network', 'Free public network', 'Personal device network', 'Emergency services only network', 'Test or experimental', 'Wildcard')
	NetworkAccessType *string `json:"networkAccessType,omitempty"`
	// An array of domain names
	Domains []string `json:"domains,omitempty"`
	// An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)
	RoamConsortOis []string `json:"roamConsortOis,omitempty"`
	// An array of MCC/MNC pairs
	MccMncs []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner `json:"mccMncs,omitempty"`
	// An array of NAI realms
	NaiRealms []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner `json:"naiRealms,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20Request instantiates a new UpdateNetworkWirelessSsidHotspot20Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20Request() *UpdateNetworkWirelessSsidHotspot20Request {
	this := UpdateNetworkWirelessSsidHotspot20Request{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestWithDefaults() *UpdateNetworkWirelessSsidHotspot20Request {
	this := UpdateNetworkWirelessSsidHotspot20Request{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetOperator() UpdateNetworkWirelessSsidHotspot20RequestOperator {
	if o == nil || IsNil(o.Operator) {
		var ret UpdateNetworkWirelessSsidHotspot20RequestOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetOperatorOk() (*UpdateNetworkWirelessSsidHotspot20RequestOperator, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given UpdateNetworkWirelessSsidHotspot20RequestOperator and assigns it to the Operator field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetOperator(v UpdateNetworkWirelessSsidHotspot20RequestOperator) {
	o.Operator = &v
}

// GetVenue returns the Venue field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetVenue() UpdateNetworkWirelessSsidHotspot20RequestVenue {
	if o == nil || IsNil(o.Venue) {
		var ret UpdateNetworkWirelessSsidHotspot20RequestVenue
		return ret
	}
	return *o.Venue
}

// GetVenueOk returns a tuple with the Venue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetVenueOk() (*UpdateNetworkWirelessSsidHotspot20RequestVenue, bool) {
	if o == nil || IsNil(o.Venue) {
		return nil, false
	}
	return o.Venue, true
}

// HasVenue returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasVenue() bool {
	if o != nil && !IsNil(o.Venue) {
		return true
	}

	return false
}

// SetVenue gets a reference to the given UpdateNetworkWirelessSsidHotspot20RequestVenue and assigns it to the Venue field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetVenue(v UpdateNetworkWirelessSsidHotspot20RequestVenue) {
	o.Venue = &v
}

// GetNetworkAccessType returns the NetworkAccessType field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNetworkAccessType() string {
	if o == nil || IsNil(o.NetworkAccessType) {
		var ret string
		return ret
	}
	return *o.NetworkAccessType
}

// GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNetworkAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkAccessType) {
		return nil, false
	}
	return o.NetworkAccessType, true
}

// HasNetworkAccessType returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasNetworkAccessType() bool {
	if o != nil && !IsNil(o.NetworkAccessType) {
		return true
	}

	return false
}

// SetNetworkAccessType gets a reference to the given string and assigns it to the NetworkAccessType field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetNetworkAccessType(v string) {
	o.NetworkAccessType = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetDomains(v []string) {
	o.Domains = v
}

// GetRoamConsortOis returns the RoamConsortOis field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetRoamConsortOis() []string {
	if o == nil || IsNil(o.RoamConsortOis) {
		var ret []string
		return ret
	}
	return o.RoamConsortOis
}

// GetRoamConsortOisOk returns a tuple with the RoamConsortOis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetRoamConsortOisOk() ([]string, bool) {
	if o == nil || IsNil(o.RoamConsortOis) {
		return nil, false
	}
	return o.RoamConsortOis, true
}

// HasRoamConsortOis returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasRoamConsortOis() bool {
	if o != nil && !IsNil(o.RoamConsortOis) {
		return true
	}

	return false
}

// SetRoamConsortOis gets a reference to the given []string and assigns it to the RoamConsortOis field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetRoamConsortOis(v []string) {
	o.RoamConsortOis = v
}

// GetMccMncs returns the MccMncs field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetMccMncs() []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner {
	if o == nil || IsNil(o.MccMncs) {
		var ret []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner
		return ret
	}
	return o.MccMncs
}

// GetMccMncsOk returns a tuple with the MccMncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetMccMncsOk() ([]UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner, bool) {
	if o == nil || IsNil(o.MccMncs) {
		return nil, false
	}
	return o.MccMncs, true
}

// HasMccMncs returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasMccMncs() bool {
	if o != nil && !IsNil(o.MccMncs) {
		return true
	}

	return false
}

// SetMccMncs gets a reference to the given []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner and assigns it to the MccMncs field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetMccMncs(v []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner) {
	o.MccMncs = v
}

// GetNaiRealms returns the NaiRealms field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNaiRealms() []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner {
	if o == nil || IsNil(o.NaiRealms) {
		var ret []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner
		return ret
	}
	return o.NaiRealms
}

// GetNaiRealmsOk returns a tuple with the NaiRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNaiRealmsOk() ([]UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner, bool) {
	if o == nil || IsNil(o.NaiRealms) {
		return nil, false
	}
	return o.NaiRealms, true
}

// HasNaiRealms returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20Request) HasNaiRealms() bool {
	if o != nil && !IsNil(o.NaiRealms) {
		return true
	}

	return false
}

// SetNaiRealms gets a reference to the given []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner and assigns it to the NaiRealms field.
func (o *UpdateNetworkWirelessSsidHotspot20Request) SetNaiRealms(v []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) {
	o.NaiRealms = v
}

func (o UpdateNetworkWirelessSsidHotspot20Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidHotspot20Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Venue) {
		toSerialize["venue"] = o.Venue
	}
	if !IsNil(o.NetworkAccessType) {
		toSerialize["networkAccessType"] = o.NetworkAccessType
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.RoamConsortOis) {
		toSerialize["roamConsortOis"] = o.RoamConsortOis
	}
	if !IsNil(o.MccMncs) {
		toSerialize["mccMncs"] = o.MccMncs
	}
	if !IsNil(o.NaiRealms) {
		toSerialize["naiRealms"] = o.NaiRealms
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidHotspot20Request struct {
	value *UpdateNetworkWirelessSsidHotspot20Request
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20Request) Get() *UpdateNetworkWirelessSsidHotspot20Request {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20Request) Set(val *UpdateNetworkWirelessSsidHotspot20Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20Request(val *UpdateNetworkWirelessSsidHotspot20Request) *NullableUpdateNetworkWirelessSsidHotspot20Request {
	return &NullableUpdateNetworkWirelessSsidHotspot20Request{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


