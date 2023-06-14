/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkWebhooksWebhookTest201Response struct for CreateNetworkWebhooksWebhookTest201Response
type CreateNetworkWebhooksWebhookTest201Response struct {
	// Webhook delivery identifier
	Id *string `json:"id,omitempty"`
	// URL where the webhook was delivered
	Url *string `json:"url,omitempty"`
	// Current status of the webhook delivery
	Status *string `json:"status,omitempty"`
}

// NewCreateNetworkWebhooksWebhookTest201Response instantiates a new CreateNetworkWebhooksWebhookTest201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksWebhookTest201Response() *CreateNetworkWebhooksWebhookTest201Response {
	this := CreateNetworkWebhooksWebhookTest201Response{}
	return &this
}

// NewCreateNetworkWebhooksWebhookTest201ResponseWithDefaults instantiates a new CreateNetworkWebhooksWebhookTest201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksWebhookTest201ResponseWithDefaults() *CreateNetworkWebhooksWebhookTest201Response {
	this := CreateNetworkWebhooksWebhookTest201Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateNetworkWebhooksWebhookTest201Response) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateNetworkWebhooksWebhookTest201Response) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTest201Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateNetworkWebhooksWebhookTest201Response) SetStatus(v string) {
	o.Status = &v
}

func (o CreateNetworkWebhooksWebhookTest201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWebhooksWebhookTest201Response struct {
	value *CreateNetworkWebhooksWebhookTest201Response
	isSet bool
}

func (v NullableCreateNetworkWebhooksWebhookTest201Response) Get() *CreateNetworkWebhooksWebhookTest201Response {
	return v.value
}

func (v *NullableCreateNetworkWebhooksWebhookTest201Response) Set(val *CreateNetworkWebhooksWebhookTest201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksWebhookTest201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksWebhookTest201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksWebhookTest201Response(val *CreateNetworkWebhooksWebhookTest201Response) *NullableCreateNetworkWebhooksWebhookTest201Response {
	return &NullableCreateNetworkWebhooksWebhookTest201Response{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksWebhookTest201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksWebhookTest201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


