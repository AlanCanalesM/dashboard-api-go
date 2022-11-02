/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkApplianceVlanRequest struct for CreateNetworkApplianceVlanRequest
type CreateNetworkApplianceVlanRequest struct {
	// The VLAN ID of the new VLAN (must be between 1 and 4094)
	Id string `json:"id"`
	// The name of the new VLAN
	Name string `json:"name"`
	// The subnet of the VLAN
	Subnet *string `json:"subnet,omitempty"`
	// The local IP of the appliance on the VLAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	// The id of the desired group policy to apply to the VLAN
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// Type of subnetting of the VLAN. Applicable only for template network.
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr *string `json:"cidr,omitempty"`
	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask *int32 `json:"mask,omitempty"`
	Ipv6 *UpdateNetworkApplianceSingleLanRequestIpv6 `json:"ipv6,omitempty"`
}

// NewCreateNetworkApplianceVlanRequest instantiates a new CreateNetworkApplianceVlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceVlanRequest(id string, name string) *CreateNetworkApplianceVlanRequest {
	this := CreateNetworkApplianceVlanRequest{}
	this.Id = id
	this.Name = name
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// NewCreateNetworkApplianceVlanRequestWithDefaults instantiates a new CreateNetworkApplianceVlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceVlanRequestWithDefaults() *CreateNetworkApplianceVlanRequest {
	this := CreateNetworkApplianceVlanRequest{}
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// GetId returns the Id field value
func (o *CreateNetworkApplianceVlanRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateNetworkApplianceVlanRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateNetworkApplianceVlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkApplianceVlanRequest) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *CreateNetworkApplianceVlanRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *CreateNetworkApplianceVlanRequest) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *CreateNetworkApplianceVlanRequest) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetTemplateVlanType returns the TemplateVlanType field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetTemplateVlanType() string {
	if o == nil || isNil(o.TemplateVlanType) {
		var ret string
		return ret
	}
	return *o.TemplateVlanType
}

// GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetTemplateVlanTypeOk() (*string, bool) {
	if o == nil || isNil(o.TemplateVlanType) {
    return nil, false
	}
	return o.TemplateVlanType, true
}

// HasTemplateVlanType returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasTemplateVlanType() bool {
	if o != nil && !isNil(o.TemplateVlanType) {
		return true
	}

	return false
}

// SetTemplateVlanType gets a reference to the given string and assigns it to the TemplateVlanType field.
func (o *CreateNetworkApplianceVlanRequest) SetTemplateVlanType(v string) {
	o.TemplateVlanType = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CreateNetworkApplianceVlanRequest) SetCidr(v string) {
	o.Cidr = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *CreateNetworkApplianceVlanRequest) SetMask(v int32) {
	o.Mask = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetIpv6() UpdateNetworkApplianceSingleLanRequestIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret UpdateNetworkApplianceSingleLanRequestIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetIpv6Ok() (*UpdateNetworkApplianceSingleLanRequestIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given UpdateNetworkApplianceSingleLanRequestIpv6 and assigns it to the Ipv6 field.
func (o *CreateNetworkApplianceVlanRequest) SetIpv6(v UpdateNetworkApplianceSingleLanRequestIpv6) {
	o.Ipv6 = &v
}

func (o CreateNetworkApplianceVlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.TemplateVlanType) {
		toSerialize["templateVlanType"] = o.TemplateVlanType
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkApplianceVlanRequest struct {
	value *CreateNetworkApplianceVlanRequest
	isSet bool
}

func (v NullableCreateNetworkApplianceVlanRequest) Get() *CreateNetworkApplianceVlanRequest {
	return v.value
}

func (v *NullableCreateNetworkApplianceVlanRequest) Set(val *CreateNetworkApplianceVlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceVlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceVlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceVlanRequest(val *CreateNetworkApplianceVlanRequest) *NullableCreateNetworkApplianceVlanRequest {
	return &NullableCreateNetworkApplianceVlanRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceVlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceVlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


