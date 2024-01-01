/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationWebhooksAlertTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksAlertTypes200Response{}

// GetOrganizationWebhooksAlertTypes200Response struct for GetOrganizationWebhooksAlertTypes200Response
type GetOrganizationWebhooksAlertTypes200Response struct {
	// The type ID of Meraki alert
	AlertTypeId *string `json:"alertTypeId,omitempty"`
	// The type of Meraki alert
	AlertType *string `json:"alertType,omitempty"`
	Example *GetOrganizationWebhooksAlertTypes200ResponseExample `json:"example,omitempty"`
}

// NewGetOrganizationWebhooksAlertTypes200Response instantiates a new GetOrganizationWebhooksAlertTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksAlertTypes200Response() *GetOrganizationWebhooksAlertTypes200Response {
	this := GetOrganizationWebhooksAlertTypes200Response{}
	return &this
}

// NewGetOrganizationWebhooksAlertTypes200ResponseWithDefaults instantiates a new GetOrganizationWebhooksAlertTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksAlertTypes200ResponseWithDefaults() *GetOrganizationWebhooksAlertTypes200Response {
	this := GetOrganizationWebhooksAlertTypes200Response{}
	return &this
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetAlertTypeId() string {
	if o == nil || IsNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertTypeId) {
		return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) HasAlertTypeId() bool {
	if o != nil && !IsNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *GetOrganizationWebhooksAlertTypes200Response) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetAlertType() string {
	if o == nil || IsNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetAlertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) HasAlertType() bool {
	if o != nil && !IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *GetOrganizationWebhooksAlertTypes200Response) SetAlertType(v string) {
	o.AlertType = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetExample() GetOrganizationWebhooksAlertTypes200ResponseExample {
	if o == nil || IsNil(o.Example) {
		var ret GetOrganizationWebhooksAlertTypes200ResponseExample
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) GetExampleOk() (*GetOrganizationWebhooksAlertTypes200ResponseExample, bool) {
	if o == nil || IsNil(o.Example) {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200Response) HasExample() bool {
	if o != nil && !IsNil(o.Example) {
		return true
	}

	return false
}

// SetExample gets a reference to the given GetOrganizationWebhooksAlertTypes200ResponseExample and assigns it to the Example field.
func (o *GetOrganizationWebhooksAlertTypes200Response) SetExample(v GetOrganizationWebhooksAlertTypes200ResponseExample) {
	o.Example = &v
}

func (o GetOrganizationWebhooksAlertTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksAlertTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertTypeId) {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	if !IsNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !IsNil(o.Example) {
		toSerialize["example"] = o.Example
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksAlertTypes200Response struct {
	value *GetOrganizationWebhooksAlertTypes200Response
	isSet bool
}

func (v NullableGetOrganizationWebhooksAlertTypes200Response) Get() *GetOrganizationWebhooksAlertTypes200Response {
	return v.value
}

func (v *NullableGetOrganizationWebhooksAlertTypes200Response) Set(val *GetOrganizationWebhooksAlertTypes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksAlertTypes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksAlertTypes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksAlertTypes200Response(val *GetOrganizationWebhooksAlertTypes200Response) *NullableGetOrganizationWebhooksAlertTypes200Response {
	return &NullableGetOrganizationWebhooksAlertTypes200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksAlertTypes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksAlertTypes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


