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

// checks if the CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade{}

// CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade The next upgrade version for the switch network
type CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade struct {
	ToVersion *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade{}
	return &this
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) GetToVersion() CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) GetToVersionOk() (*CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) SetToVersion(v CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade struct {
	value *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) Get() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) Set(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade {
	return &NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


