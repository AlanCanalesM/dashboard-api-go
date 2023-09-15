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

// checks if the GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion{}

// GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion Version the device is upgrading to
type GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion struct {
	// ID of the initial firmware version
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion{}
	return &this
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersionWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersionWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return toSerialize, nil
}

type NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion struct {
	value *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) Get() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) Set(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion {
	return &NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


