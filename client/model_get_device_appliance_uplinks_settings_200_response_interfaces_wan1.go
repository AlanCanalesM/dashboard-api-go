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

// checks if the GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1{}

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 WAN 1 settings.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 struct {
	// Enable or disable the interface.
	Enabled *bool `json:"enabled,omitempty"`
	VlanTagging *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging `json:"vlanTagging,omitempty"`
	Svis *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis `json:"svis,omitempty"`
	Pppoe *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe `json:"pppoe,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1WithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1WithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetVlanTagging() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	if o == nil || IsNil(o.VlanTagging) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetVlanTaggingOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging, bool) {
	if o == nil || IsNil(o.VlanTagging) {
		return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) HasVlanTagging() bool {
	if o != nil && !IsNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging and assigns it to the VlanTagging field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) SetVlanTagging(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) {
	o.VlanTagging = &v
}

// GetSvis returns the Svis field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetSvis() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis {
	if o == nil || IsNil(o.Svis) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis
		return ret
	}
	return *o.Svis
}

// GetSvisOk returns a tuple with the Svis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetSvisOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis, bool) {
	if o == nil || IsNil(o.Svis) {
		return nil, false
	}
	return o.Svis, true
}

// HasSvis returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) HasSvis() bool {
	if o != nil && !IsNil(o.Svis) {
		return true
	}

	return false
}

// SetSvis gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis and assigns it to the Svis field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) SetSvis(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis) {
	o.Svis = &v
}

// GetPppoe returns the Pppoe field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetPppoe() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe {
	if o == nil || IsNil(o.Pppoe) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe
		return ret
	}
	return *o.Pppoe
}

// GetPppoeOk returns a tuple with the Pppoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) GetPppoeOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe, bool) {
	if o == nil || IsNil(o.Pppoe) {
		return nil, false
	}
	return o.Pppoe, true
}

// HasPppoe returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) HasPppoe() bool {
	if o != nil && !IsNil(o.Pppoe) {
		return true
	}

	return false
}

// SetPppoe gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe and assigns it to the Pppoe field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) SetPppoe(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) {
	o.Pppoe = &v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.VlanTagging) {
		toSerialize["vlanTagging"] = o.VlanTagging
	}
	if !IsNil(o.Svis) {
		toSerialize["svis"] = o.Svis
	}
	if !IsNil(o.Pppoe) {
		toSerialize["pppoe"] = o.Pppoe
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1 {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


