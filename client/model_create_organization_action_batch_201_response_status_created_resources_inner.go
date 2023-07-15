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

// checks if the CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner{}

// CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner struct for CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner
type CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner struct {
	// ID of the created resource
	Id *string `json:"id,omitempty"`
	// URI, not including base, of the created resource
	Uri *string `json:"uri,omitempty"`
}

// NewCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner instantiates a new CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner() *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner {
	this := CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner{}
	return &this
}

// NewCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInnerWithDefaults instantiates a new CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInnerWithDefaults() *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner {
	this := CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) SetUri(v string) {
	o.Uri = &v
}

func (o CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner struct {
	value *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner
	isSet bool
}

func (v NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) Get() *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner {
	return v.value
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) Set(val *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner(val *CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) *NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner {
	return &NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


