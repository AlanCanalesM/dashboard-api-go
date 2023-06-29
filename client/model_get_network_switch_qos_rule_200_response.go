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

// checks if the GetNetworkSwitchQosRule200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchQosRule200Response{}

// GetNetworkSwitchQosRule200Response struct for GetNetworkSwitchQosRule200Response
type GetNetworkSwitchQosRule200Response struct {
	// Qos Rule id
	Id *string `json:"id,omitempty"`
	// The VLAN of the incoming packet. A null value will match any VLAN.
	Vlan *int32 `json:"vlan,omitempty"`
	// The protocol of the incoming packet. Can be one of \"ANY\", \"TCP\" or \"UDP\". Default value is \"ANY\"
	Protocol *string `json:"protocol,omitempty"`
	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort *int32 `json:"srcPort,omitempty"`
	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	SrcPortRange *string `json:"srcPortRange,omitempty"`
	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort *int32 `json:"dstPort,omitempty"`
	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	DstPortRange *string `json:"dstPortRange,omitempty"`
	// DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp *int32 `json:"dscp,omitempty"`
}

// NewGetNetworkSwitchQosRule200Response instantiates a new GetNetworkSwitchQosRule200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchQosRule200Response() *GetNetworkSwitchQosRule200Response {
	this := GetNetworkSwitchQosRule200Response{}
	return &this
}

// NewGetNetworkSwitchQosRule200ResponseWithDefaults instantiates a new GetNetworkSwitchQosRule200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchQosRule200ResponseWithDefaults() *GetNetworkSwitchQosRule200Response {
	this := GetNetworkSwitchQosRule200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSwitchQosRule200Response) SetId(v string) {
	o.Id = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkSwitchQosRule200Response) SetVlan(v int32) {
	o.Vlan = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkSwitchQosRule200Response) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetSrcPort() int32 {
	if o == nil || IsNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetSrcPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *GetNetworkSwitchQosRule200Response) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcPortRange returns the SrcPortRange field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetSrcPortRange() string {
	if o == nil || IsNil(o.SrcPortRange) {
		var ret string
		return ret
	}
	return *o.SrcPortRange
}

// GetSrcPortRangeOk returns a tuple with the SrcPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetSrcPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.SrcPortRange) {
		return nil, false
	}
	return o.SrcPortRange, true
}

// HasSrcPortRange returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasSrcPortRange() bool {
	if o != nil && !IsNil(o.SrcPortRange) {
		return true
	}

	return false
}

// SetSrcPortRange gets a reference to the given string and assigns it to the SrcPortRange field.
func (o *GetNetworkSwitchQosRule200Response) SetSrcPortRange(v string) {
	o.SrcPortRange = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetDstPort() int32 {
	if o == nil || IsNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetDstPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DstPort) {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasDstPort() bool {
	if o != nil && !IsNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *GetNetworkSwitchQosRule200Response) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetDstPortRange returns the DstPortRange field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetDstPortRange() string {
	if o == nil || IsNil(o.DstPortRange) {
		var ret string
		return ret
	}
	return *o.DstPortRange
}

// GetDstPortRangeOk returns a tuple with the DstPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetDstPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DstPortRange) {
		return nil, false
	}
	return o.DstPortRange, true
}

// HasDstPortRange returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasDstPortRange() bool {
	if o != nil && !IsNil(o.DstPortRange) {
		return true
	}

	return false
}

// SetDstPortRange gets a reference to the given string and assigns it to the DstPortRange field.
func (o *GetNetworkSwitchQosRule200Response) SetDstPortRange(v string) {
	o.DstPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *GetNetworkSwitchQosRule200Response) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchQosRule200Response) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *GetNetworkSwitchQosRule200Response) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *GetNetworkSwitchQosRule200Response) SetDscp(v int32) {
	o.Dscp = &v
}

func (o GetNetworkSwitchQosRule200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchQosRule200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !IsNil(o.SrcPortRange) {
		toSerialize["srcPortRange"] = o.SrcPortRange
	}
	if !IsNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !IsNil(o.DstPortRange) {
		toSerialize["dstPortRange"] = o.DstPortRange
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchQosRule200Response struct {
	value *GetNetworkSwitchQosRule200Response
	isSet bool
}

func (v NullableGetNetworkSwitchQosRule200Response) Get() *GetNetworkSwitchQosRule200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchQosRule200Response) Set(val *GetNetworkSwitchQosRule200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchQosRule200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchQosRule200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchQosRule200Response(val *GetNetworkSwitchQosRule200Response) *NullableGetNetworkSwitchQosRule200Response {
	return &NullableGetNetworkSwitchQosRule200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchQosRule200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchQosRule200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


