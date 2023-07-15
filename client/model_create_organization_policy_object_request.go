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

// checks if the CreateOrganizationPolicyObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationPolicyObjectRequest{}

// CreateOrganizationPolicyObjectRequest struct for CreateOrganizationPolicyObjectRequest
type CreateOrganizationPolicyObjectRequest struct {
	// Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name"`
	// Category of a policy object (one of: adaptivePolicy, network)
	Category string `json:"category"`
	// Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)
	Type string `json:"type"`
	// CIDR Value of a policy object (e.g. 10.11.12.1/24\")
	Cidr *string `json:"cidr,omitempty"`
	// Fully qualified domain name of policy object (e.g. \"example.com\")
	Fqdn *string `json:"fqdn,omitempty"`
	// Mask of a policy object (e.g. \"255.255.0.0\")
	Mask *string `json:"mask,omitempty"`
	// IP Address of a policy object (e.g. \"1.2.3.4\")
	Ip *string `json:"ip,omitempty"`
	// The IDs of policy object groups the policy object belongs to
	GroupIds []int32 `json:"groupIds,omitempty"`
}

// NewCreateOrganizationPolicyObjectRequest instantiates a new CreateOrganizationPolicyObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationPolicyObjectRequest(name string, category string, type_ string) *CreateOrganizationPolicyObjectRequest {
	this := CreateOrganizationPolicyObjectRequest{}
	this.Name = name
	this.Category = category
	this.Type = type_
	return &this
}

// NewCreateOrganizationPolicyObjectRequestWithDefaults instantiates a new CreateOrganizationPolicyObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationPolicyObjectRequestWithDefaults() *CreateOrganizationPolicyObjectRequest {
	this := CreateOrganizationPolicyObjectRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationPolicyObjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationPolicyObjectRequest) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value
func (o *CreateOrganizationPolicyObjectRequest) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CreateOrganizationPolicyObjectRequest) SetCategory(v string) {
	o.Category = v
}

// GetType returns the Type field value
func (o *CreateOrganizationPolicyObjectRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrganizationPolicyObjectRequest) SetType(v string) {
	o.Type = v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CreateOrganizationPolicyObjectRequest) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CreateOrganizationPolicyObjectRequest) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CreateOrganizationPolicyObjectRequest) SetCidr(v string) {
	o.Cidr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *CreateOrganizationPolicyObjectRequest) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *CreateOrganizationPolicyObjectRequest) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *CreateOrganizationPolicyObjectRequest) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *CreateOrganizationPolicyObjectRequest) GetMask() string {
	if o == nil || IsNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateOrganizationPolicyObjectRequest) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *CreateOrganizationPolicyObjectRequest) SetMask(v string) {
	o.Mask = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CreateOrganizationPolicyObjectRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreateOrganizationPolicyObjectRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CreateOrganizationPolicyObjectRequest) SetIp(v string) {
	o.Ip = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *CreateOrganizationPolicyObjectRequest) GetGroupIds() []int32 {
	if o == nil || IsNil(o.GroupIds) {
		var ret []int32
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationPolicyObjectRequest) GetGroupIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *CreateOrganizationPolicyObjectRequest) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []int32 and assigns it to the GroupIds field.
func (o *CreateOrganizationPolicyObjectRequest) SetGroupIds(v []int32) {
	o.GroupIds = v
}

func (o CreateOrganizationPolicyObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationPolicyObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["category"] = o.Category
	toSerialize["type"] = o.Type
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	return toSerialize, nil
}

type NullableCreateOrganizationPolicyObjectRequest struct {
	value *CreateOrganizationPolicyObjectRequest
	isSet bool
}

func (v NullableCreateOrganizationPolicyObjectRequest) Get() *CreateOrganizationPolicyObjectRequest {
	return v.value
}

func (v *NullableCreateOrganizationPolicyObjectRequest) Set(val *CreateOrganizationPolicyObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationPolicyObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationPolicyObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationPolicyObjectRequest(val *CreateOrganizationPolicyObjectRequest) *NullableCreateOrganizationPolicyObjectRequest {
	return &NullableCreateOrganizationPolicyObjectRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationPolicyObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationPolicyObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


