/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner{}

// GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner struct for GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner
type GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner struct {
	// The id of the category
	Id *string `json:"id,omitempty"`
	// The name of the category
	Name *string `json:"name,omitempty"`
	// Details of the associated applications
	Applications []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner `json:"applications,omitempty"`
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner{}
	return &this
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerWithDefaults instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerWithDefaults() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) SetName(v string) {
	o.Name = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetApplications() []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner {
	if o == nil || IsNil(o.Applications) {
		var ret []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) GetApplicationsOk() ([]GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner and assigns it to the Applications field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) SetApplications(v []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) {
	o.Applications = v
}

func (o GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner struct {
	value *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) Get() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) Set(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner {
	return &NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


