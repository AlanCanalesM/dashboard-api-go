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

// checks if the UpdateNetworkCellularGatewayDhcpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkCellularGatewayDhcpRequest{}

// UpdateNetworkCellularGatewayDhcpRequest struct for UpdateNetworkCellularGatewayDhcpRequest
type UpdateNetworkCellularGatewayDhcpRequest struct {
	// DHCP Lease time for all MG of the network. Possible values are '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// DNS name servers mode for all MG of the network. Possible values are: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// list of fixed IPs representing the the DNS Name servers when the mode is 'custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewUpdateNetworkCellularGatewayDhcpRequest instantiates a new UpdateNetworkCellularGatewayDhcpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkCellularGatewayDhcpRequest() *UpdateNetworkCellularGatewayDhcpRequest {
	this := UpdateNetworkCellularGatewayDhcpRequest{}
	return &this
}

// NewUpdateNetworkCellularGatewayDhcpRequestWithDefaults instantiates a new UpdateNetworkCellularGatewayDhcpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkCellularGatewayDhcpRequestWithDefaults() *UpdateNetworkCellularGatewayDhcpRequest {
	this := UpdateNetworkCellularGatewayDhcpRequest{}
	return &this
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDhcpLeaseTime() string {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDhcpLeaseTime() bool {
	if o != nil && !IsNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsNameservers() string {
	if o == nil || IsNil(o.DnsNameservers) {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsNameserversOk() (*string, bool) {
	if o == nil || IsNil(o.DnsNameservers) {
		return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDnsNameservers() bool {
	if o != nil && !IsNil(o.DnsNameservers) {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsCustomNameservers() []string {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDnsCustomNameservers() bool {
	if o != nil && !IsNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o UpdateNetworkCellularGatewayDhcpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkCellularGatewayDhcpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !IsNil(o.DnsNameservers) {
		toSerialize["dnsNameservers"] = o.DnsNameservers
	}
	if !IsNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	return toSerialize, nil
}

type NullableUpdateNetworkCellularGatewayDhcpRequest struct {
	value *UpdateNetworkCellularGatewayDhcpRequest
	isSet bool
}

func (v NullableUpdateNetworkCellularGatewayDhcpRequest) Get() *UpdateNetworkCellularGatewayDhcpRequest {
	return v.value
}

func (v *NullableUpdateNetworkCellularGatewayDhcpRequest) Set(val *UpdateNetworkCellularGatewayDhcpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkCellularGatewayDhcpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkCellularGatewayDhcpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkCellularGatewayDhcpRequest(val *UpdateNetworkCellularGatewayDhcpRequest) *NullableUpdateNetworkCellularGatewayDhcpRequest {
	return &NullableUpdateNetworkCellularGatewayDhcpRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkCellularGatewayDhcpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkCellularGatewayDhcpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


