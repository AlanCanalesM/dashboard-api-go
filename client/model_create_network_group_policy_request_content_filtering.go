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

// checks if the CreateNetworkGroupPolicyRequestContentFiltering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequestContentFiltering{}

// CreateNetworkGroupPolicyRequestContentFiltering The content filtering settings for your group policy
type CreateNetworkGroupPolicyRequestContentFiltering struct {
	AllowedUrlPatterns *CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns `json:"allowedUrlPatterns,omitempty"`
	BlockedUrlPatterns *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns `json:"blockedUrlPatterns,omitempty"`
	BlockedUrlCategories *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories `json:"blockedUrlCategories,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestContentFiltering instantiates a new CreateNetworkGroupPolicyRequestContentFiltering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestContentFiltering() *CreateNetworkGroupPolicyRequestContentFiltering {
	this := CreateNetworkGroupPolicyRequestContentFiltering{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestContentFilteringWithDefaults instantiates a new CreateNetworkGroupPolicyRequestContentFiltering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestContentFilteringWithDefaults() *CreateNetworkGroupPolicyRequestContentFiltering {
	this := CreateNetworkGroupPolicyRequestContentFiltering{}
	return &this
}

// GetAllowedUrlPatterns returns the AllowedUrlPatterns field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetAllowedUrlPatterns() CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns {
	if o == nil || IsNil(o.AllowedUrlPatterns) {
		var ret CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns
		return ret
	}
	return *o.AllowedUrlPatterns
}

// GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetAllowedUrlPatternsOk() (*CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns, bool) {
	if o == nil || IsNil(o.AllowedUrlPatterns) {
		return nil, false
	}
	return o.AllowedUrlPatterns, true
}

// HasAllowedUrlPatterns returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) HasAllowedUrlPatterns() bool {
	if o != nil && !IsNil(o.AllowedUrlPatterns) {
		return true
	}

	return false
}

// SetAllowedUrlPatterns gets a reference to the given CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns and assigns it to the AllowedUrlPatterns field.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) SetAllowedUrlPatterns(v CreateNetworkGroupPolicyRequestContentFilteringAllowedUrlPatterns) {
	o.AllowedUrlPatterns = &v
}

// GetBlockedUrlPatterns returns the BlockedUrlPatterns field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetBlockedUrlPatterns() CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns {
	if o == nil || IsNil(o.BlockedUrlPatterns) {
		var ret CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns
		return ret
	}
	return *o.BlockedUrlPatterns
}

// GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetBlockedUrlPatternsOk() (*CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns, bool) {
	if o == nil || IsNil(o.BlockedUrlPatterns) {
		return nil, false
	}
	return o.BlockedUrlPatterns, true
}

// HasBlockedUrlPatterns returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) HasBlockedUrlPatterns() bool {
	if o != nil && !IsNil(o.BlockedUrlPatterns) {
		return true
	}

	return false
}

// SetBlockedUrlPatterns gets a reference to the given CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns and assigns it to the BlockedUrlPatterns field.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) SetBlockedUrlPatterns(v CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) {
	o.BlockedUrlPatterns = &v
}

// GetBlockedUrlCategories returns the BlockedUrlCategories field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetBlockedUrlCategories() CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories {
	if o == nil || IsNil(o.BlockedUrlCategories) {
		var ret CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories
		return ret
	}
	return *o.BlockedUrlCategories
}

// GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) GetBlockedUrlCategoriesOk() (*CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories, bool) {
	if o == nil || IsNil(o.BlockedUrlCategories) {
		return nil, false
	}
	return o.BlockedUrlCategories, true
}

// HasBlockedUrlCategories returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) HasBlockedUrlCategories() bool {
	if o != nil && !IsNil(o.BlockedUrlCategories) {
		return true
	}

	return false
}

// SetBlockedUrlCategories gets a reference to the given CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories and assigns it to the BlockedUrlCategories field.
func (o *CreateNetworkGroupPolicyRequestContentFiltering) SetBlockedUrlCategories(v CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) {
	o.BlockedUrlCategories = &v
}

func (o CreateNetworkGroupPolicyRequestContentFiltering) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequestContentFiltering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedUrlPatterns) {
		toSerialize["allowedUrlPatterns"] = o.AllowedUrlPatterns
	}
	if !IsNil(o.BlockedUrlPatterns) {
		toSerialize["blockedUrlPatterns"] = o.BlockedUrlPatterns
	}
	if !IsNil(o.BlockedUrlCategories) {
		toSerialize["blockedUrlCategories"] = o.BlockedUrlCategories
	}
	return toSerialize, nil
}

type NullableCreateNetworkGroupPolicyRequestContentFiltering struct {
	value *CreateNetworkGroupPolicyRequestContentFiltering
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestContentFiltering) Get() *CreateNetworkGroupPolicyRequestContentFiltering {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFiltering) Set(val *CreateNetworkGroupPolicyRequestContentFiltering) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestContentFiltering) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFiltering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestContentFiltering(val *CreateNetworkGroupPolicyRequestContentFiltering) *NullableCreateNetworkGroupPolicyRequestContentFiltering {
	return &NullableCreateNetworkGroupPolicyRequestContentFiltering{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestContentFiltering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFiltering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


