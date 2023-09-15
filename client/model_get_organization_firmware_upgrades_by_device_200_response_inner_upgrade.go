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

// checks if the GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade{}

// GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade The devices upgrade details and status
type GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade struct {
	// Start time of the upgrade
	Time *string `json:"time,omitempty"`
	FromVersion *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion `json:"fromVersion,omitempty"`
	ToVersion *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion `json:"toVersion,omitempty"`
	// Status of the upgrade
	Status *string `json:"status,omitempty"`
	// ID of the upgrade
	Id *string `json:"id,omitempty"`
	// ID of the upgrade batch
	UpgradeBatchId *string `json:"upgradeBatchId,omitempty"`
	Staged *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged `json:"staged,omitempty"`
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade{}
	return &this
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetTime(v string) {
	o.Time = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetFromVersion() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion {
	if o == nil || IsNil(o.FromVersion) {
		var ret GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetFromVersionOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion, bool) {
	if o == nil || IsNil(o.FromVersion) {
		return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasFromVersion() bool {
	if o != nil && !IsNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion and assigns it to the FromVersion field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetFromVersion(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion) {
	o.FromVersion = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetToVersion() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetToVersionOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion and assigns it to the ToVersion field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetToVersion(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion) {
	o.ToVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetStatus(v string) {
	o.Status = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetId(v string) {
	o.Id = &v
}

// GetUpgradeBatchId returns the UpgradeBatchId field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetUpgradeBatchId() string {
	if o == nil || IsNil(o.UpgradeBatchId) {
		var ret string
		return ret
	}
	return *o.UpgradeBatchId
}

// GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetUpgradeBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeBatchId) {
		return nil, false
	}
	return o.UpgradeBatchId, true
}

// HasUpgradeBatchId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasUpgradeBatchId() bool {
	if o != nil && !IsNil(o.UpgradeBatchId) {
		return true
	}

	return false
}

// SetUpgradeBatchId gets a reference to the given string and assigns it to the UpgradeBatchId field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetUpgradeBatchId(v string) {
	o.UpgradeBatchId = &v
}

// GetStaged returns the Staged field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStaged() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged {
	if o == nil || IsNil(o.Staged) {
		var ret GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged
		return ret
	}
	return *o.Staged
}

// GetStagedOk returns a tuple with the Staged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStagedOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged, bool) {
	if o == nil || IsNil(o.Staged) {
		return nil, false
	}
	return o.Staged, true
}

// HasStaged returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasStaged() bool {
	if o != nil && !IsNil(o.Staged) {
		return true
	}

	return false
}

// SetStaged gets a reference to the given GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged and assigns it to the Staged field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetStaged(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged) {
	o.Staged = &v
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
	}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UpgradeBatchId) {
		toSerialize["upgradeBatchId"] = o.UpgradeBatchId
	}
	if !IsNil(o.Staged) {
		toSerialize["staged"] = o.Staged
	}
	return toSerialize, nil
}

type NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade struct {
	value *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) Get() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) Set(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade {
	return &NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


