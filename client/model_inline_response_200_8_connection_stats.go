/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2008ConnectionStats The connection stats of the device
type InlineResponse2008ConnectionStats struct {
	// The number of failed association attempts
	Assoc *int32 `json:"assoc,omitempty"`
	// The number of failed authentication attempts
	Auth *int32 `json:"auth,omitempty"`
	// The number of failed DHCP attempts
	Dhcp *int32 `json:"dhcp,omitempty"`
	// The number of failed DNS attempts
	Dns *int32 `json:"dns,omitempty"`
	// The number of successful connection attempts
	Success *int32 `json:"success,omitempty"`
}

// NewInlineResponse2008ConnectionStats instantiates a new InlineResponse2008ConnectionStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008ConnectionStats() *InlineResponse2008ConnectionStats {
	this := InlineResponse2008ConnectionStats{}
	return &this
}

// NewInlineResponse2008ConnectionStatsWithDefaults instantiates a new InlineResponse2008ConnectionStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008ConnectionStatsWithDefaults() *InlineResponse2008ConnectionStats {
	this := InlineResponse2008ConnectionStats{}
	return &this
}

// GetAssoc returns the Assoc field value if set, zero value otherwise.
func (o *InlineResponse2008ConnectionStats) GetAssoc() int32 {
	if o == nil || isNil(o.Assoc) {
		var ret int32
		return ret
	}
	return *o.Assoc
}

// GetAssocOk returns a tuple with the Assoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008ConnectionStats) GetAssocOk() (*int32, bool) {
	if o == nil || isNil(o.Assoc) {
    return nil, false
	}
	return o.Assoc, true
}

// HasAssoc returns a boolean if a field has been set.
func (o *InlineResponse2008ConnectionStats) HasAssoc() bool {
	if o != nil && !isNil(o.Assoc) {
		return true
	}

	return false
}

// SetAssoc gets a reference to the given int32 and assigns it to the Assoc field.
func (o *InlineResponse2008ConnectionStats) SetAssoc(v int32) {
	o.Assoc = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *InlineResponse2008ConnectionStats) GetAuth() int32 {
	if o == nil || isNil(o.Auth) {
		var ret int32
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008ConnectionStats) GetAuthOk() (*int32, bool) {
	if o == nil || isNil(o.Auth) {
    return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *InlineResponse2008ConnectionStats) HasAuth() bool {
	if o != nil && !isNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given int32 and assigns it to the Auth field.
func (o *InlineResponse2008ConnectionStats) SetAuth(v int32) {
	o.Auth = &v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *InlineResponse2008ConnectionStats) GetDhcp() int32 {
	if o == nil || isNil(o.Dhcp) {
		var ret int32
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008ConnectionStats) GetDhcpOk() (*int32, bool) {
	if o == nil || isNil(o.Dhcp) {
    return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *InlineResponse2008ConnectionStats) HasDhcp() bool {
	if o != nil && !isNil(o.Dhcp) {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given int32 and assigns it to the Dhcp field.
func (o *InlineResponse2008ConnectionStats) SetDhcp(v int32) {
	o.Dhcp = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *InlineResponse2008ConnectionStats) GetDns() int32 {
	if o == nil || isNil(o.Dns) {
		var ret int32
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008ConnectionStats) GetDnsOk() (*int32, bool) {
	if o == nil || isNil(o.Dns) {
    return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *InlineResponse2008ConnectionStats) HasDns() bool {
	if o != nil && !isNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given int32 and assigns it to the Dns field.
func (o *InlineResponse2008ConnectionStats) SetDns(v int32) {
	o.Dns = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse2008ConnectionStats) GetSuccess() int32 {
	if o == nil || isNil(o.Success) {
		var ret int32
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008ConnectionStats) GetSuccessOk() (*int32, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse2008ConnectionStats) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given int32 and assigns it to the Success field.
func (o *InlineResponse2008ConnectionStats) SetSuccess(v int32) {
	o.Success = &v
}

func (o InlineResponse2008ConnectionStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assoc) {
		toSerialize["assoc"] = o.Assoc
	}
	if !isNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !isNil(o.Dhcp) {
		toSerialize["dhcp"] = o.Dhcp
	}
	if !isNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008ConnectionStats struct {
	value *InlineResponse2008ConnectionStats
	isSet bool
}

func (v NullableInlineResponse2008ConnectionStats) Get() *InlineResponse2008ConnectionStats {
	return v.value
}

func (v *NullableInlineResponse2008ConnectionStats) Set(val *InlineResponse2008ConnectionStats) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008ConnectionStats) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008ConnectionStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008ConnectionStats(val *InlineResponse2008ConnectionStats) *NullableInlineResponse2008ConnectionStats {
	return &NullableInlineResponse2008ConnectionStats{value: val, isSet: true}
}

func (v NullableInlineResponse2008ConnectionStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008ConnectionStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


