/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6{}

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 IPv6 settings for static/dynamic mode.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 struct {
	// The assignment mode for this SVI. Applies only when PPPoE is disabled.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// Static address that will override the one(s) received by SLAAC.
	Address *string `json:"address,omitempty"`
	// Static gateway that will override the one received by autoconf.
	Gateway *string `json:"gateway,omitempty"`
	Nameservers *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers `json:"nameservers,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6WithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6WithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAssignmentMode() string {
	if o == nil || IsNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentMode) {
		return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasAssignmentMode() bool {
	if o != nil && !IsNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetGateway(v string) {
	o.Gateway = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetNameservers() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers {
	if o == nil || IsNil(o.Nameservers) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetNameserversOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers and assigns it to the Nameservers field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetNameservers(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) {
	o.Nameservers = &v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


