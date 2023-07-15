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

// checks if the GetDeviceSwitchRoutingStaticRoute200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingStaticRoute200Response{}

// GetDeviceSwitchRoutingStaticRoute200Response struct for GetDeviceSwitchRoutingStaticRoute200Response
type GetDeviceSwitchRoutingStaticRoute200Response struct {
	// The identifier of a layer 3 static route
	StaticRouteId *string `json:"staticRouteId,omitempty"`
	// The name or description of the layer 3 static route
	Name *string `json:"name,omitempty"`
	// The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet string `json:"subnet"`
	//  The IP address of the router to which traffic for this destination network should be sent
	NextHopIp string `json:"nextHopIp"`
	// Option to advertise static routes via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static routes over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewGetDeviceSwitchRoutingStaticRoute200Response instantiates a new GetDeviceSwitchRoutingStaticRoute200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingStaticRoute200Response(subnet string, nextHopIp string) *GetDeviceSwitchRoutingStaticRoute200Response {
	this := GetDeviceSwitchRoutingStaticRoute200Response{}
	this.Subnet = subnet
	this.NextHopIp = nextHopIp
	return &this
}

// NewGetDeviceSwitchRoutingStaticRoute200ResponseWithDefaults instantiates a new GetDeviceSwitchRoutingStaticRoute200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingStaticRoute200ResponseWithDefaults() *GetDeviceSwitchRoutingStaticRoute200Response {
	this := GetDeviceSwitchRoutingStaticRoute200Response{}
	return &this
}

// GetStaticRouteId returns the StaticRouteId field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetStaticRouteId() string {
	if o == nil || IsNil(o.StaticRouteId) {
		var ret string
		return ret
	}
	return *o.StaticRouteId
}

// GetStaticRouteIdOk returns a tuple with the StaticRouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetStaticRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.StaticRouteId) {
		return nil, false
	}
	return o.StaticRouteId, true
}

// HasStaticRouteId returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasStaticRouteId() bool {
	if o != nil && !IsNil(o.StaticRouteId) {
		return true
	}

	return false
}

// SetStaticRouteId gets a reference to the given string and assigns it to the StaticRouteId field.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetStaticRouteId(v string) {
	o.StaticRouteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetSubnet(v string) {
	o.Subnet = v
}

// GetNextHopIp returns the NextHopIp field value
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNextHopIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNextHopIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextHopIp, true
}

// SetNextHopIp sets field value
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetNextHopIp(v string) {
	o.NextHopIp = v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || IsNil(o.AdvertiseViaOspfEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvertiseViaOspfEnabled) {
		return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && !IsNil(o.AdvertiseViaOspfEnabled) {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || IsNil(o.PreferOverOspfRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferOverOspfRoutesEnabled) {
		return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && !IsNil(o.PreferOverOspfRoutesEnabled) {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o GetDeviceSwitchRoutingStaticRoute200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingStaticRoute200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StaticRouteId) {
		toSerialize["staticRouteId"] = o.StaticRouteId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["subnet"] = o.Subnet
	toSerialize["nextHopIp"] = o.NextHopIp
	if !IsNil(o.AdvertiseViaOspfEnabled) {
		toSerialize["advertiseViaOspfEnabled"] = o.AdvertiseViaOspfEnabled
	}
	if !IsNil(o.PreferOverOspfRoutesEnabled) {
		toSerialize["preferOverOspfRoutesEnabled"] = o.PreferOverOspfRoutesEnabled
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchRoutingStaticRoute200Response struct {
	value *GetDeviceSwitchRoutingStaticRoute200Response
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingStaticRoute200Response) Get() *GetDeviceSwitchRoutingStaticRoute200Response {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingStaticRoute200Response) Set(val *GetDeviceSwitchRoutingStaticRoute200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingStaticRoute200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingStaticRoute200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingStaticRoute200Response(val *GetDeviceSwitchRoutingStaticRoute200Response) *NullableGetDeviceSwitchRoutingStaticRoute200Response {
	return &NullableGetDeviceSwitchRoutingStaticRoute200Response{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingStaticRoute200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingStaticRoute200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


