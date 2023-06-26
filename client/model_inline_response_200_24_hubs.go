/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20024Hubs struct for InlineResponse20024Hubs
type InlineResponse20024Hubs struct {
	// The network ID of the hub.
	HubId *string `json:"hubId,omitempty"`
	// Indicates whether default route traffic should be sent to this hub.
	UseDefaultRoute *bool `json:"useDefaultRoute,omitempty"`
}

// NewInlineResponse20024Hubs instantiates a new InlineResponse20024Hubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20024Hubs() *InlineResponse20024Hubs {
	this := InlineResponse20024Hubs{}
	return &this
}

// NewInlineResponse20024HubsWithDefaults instantiates a new InlineResponse20024Hubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20024HubsWithDefaults() *InlineResponse20024Hubs {
	this := InlineResponse20024Hubs{}
	return &this
}

// GetHubId returns the HubId field value if set, zero value otherwise.
func (o *InlineResponse20024Hubs) GetHubId() string {
	if o == nil || isNil(o.HubId) {
		var ret string
		return ret
	}
	return *o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024Hubs) GetHubIdOk() (*string, bool) {
	if o == nil || isNil(o.HubId) {
    return nil, false
	}
	return o.HubId, true
}

// HasHubId returns a boolean if a field has been set.
func (o *InlineResponse20024Hubs) HasHubId() bool {
	if o != nil && !isNil(o.HubId) {
		return true
	}

	return false
}

// SetHubId gets a reference to the given string and assigns it to the HubId field.
func (o *InlineResponse20024Hubs) SetHubId(v string) {
	o.HubId = &v
}

// GetUseDefaultRoute returns the UseDefaultRoute field value if set, zero value otherwise.
func (o *InlineResponse20024Hubs) GetUseDefaultRoute() bool {
	if o == nil || isNil(o.UseDefaultRoute) {
		var ret bool
		return ret
	}
	return *o.UseDefaultRoute
}

// GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024Hubs) GetUseDefaultRouteOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefaultRoute) {
    return nil, false
	}
	return o.UseDefaultRoute, true
}

// HasUseDefaultRoute returns a boolean if a field has been set.
func (o *InlineResponse20024Hubs) HasUseDefaultRoute() bool {
	if o != nil && !isNil(o.UseDefaultRoute) {
		return true
	}

	return false
}

// SetUseDefaultRoute gets a reference to the given bool and assigns it to the UseDefaultRoute field.
func (o *InlineResponse20024Hubs) SetUseDefaultRoute(v bool) {
	o.UseDefaultRoute = &v
}

func (o InlineResponse20024Hubs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HubId) {
		toSerialize["hubId"] = o.HubId
	}
	if !isNil(o.UseDefaultRoute) {
		toSerialize["useDefaultRoute"] = o.UseDefaultRoute
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20024Hubs struct {
	value *InlineResponse20024Hubs
	isSet bool
}

func (v NullableInlineResponse20024Hubs) Get() *InlineResponse20024Hubs {
	return v.value
}

func (v *NullableInlineResponse20024Hubs) Set(val *InlineResponse20024Hubs) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20024Hubs) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20024Hubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20024Hubs(val *InlineResponse20024Hubs) *NullableInlineResponse20024Hubs {
	return &NullableInlineResponse20024Hubs{value: val, isSet: true}
}

func (v NullableInlineResponse20024Hubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20024Hubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


