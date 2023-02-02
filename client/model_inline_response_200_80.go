/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20080 struct for InlineResponse20080
type InlineResponse20080 struct {
	// Organization ID
	Id *string `json:"id,omitempty"`
	// Organization name
	Name *string `json:"name,omitempty"`
	// Organization URL
	Url *string `json:"url,omitempty"`
	Api *OrganizationsApi `json:"api,omitempty"`
	Licensing *OrganizationsLicensing `json:"licensing,omitempty"`
	Cloud *OrganizationsCloud `json:"cloud,omitempty"`
	Management *OrganizationsManagement `json:"management,omitempty"`
}

// NewInlineResponse20080 instantiates a new InlineResponse20080 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20080() *InlineResponse20080 {
	this := InlineResponse20080{}
	return &this
}

// NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20080WithDefaults() *InlineResponse20080 {
	this := InlineResponse20080{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20080) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20080) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20080) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20080) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20080) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20080) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20080) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20080) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20080) SetUrl(v string) {
	o.Url = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *InlineResponse20080) GetApi() OrganizationsApi {
	if o == nil || isNil(o.Api) {
		var ret OrganizationsApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetApiOk() (*OrganizationsApi, bool) {
	if o == nil || isNil(o.Api) {
    return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *InlineResponse20080) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given OrganizationsApi and assigns it to the Api field.
func (o *InlineResponse20080) SetApi(v OrganizationsApi) {
	o.Api = &v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *InlineResponse20080) GetLicensing() OrganizationsLicensing {
	if o == nil || isNil(o.Licensing) {
		var ret OrganizationsLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetLicensingOk() (*OrganizationsLicensing, bool) {
	if o == nil || isNil(o.Licensing) {
    return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *InlineResponse20080) HasLicensing() bool {
	if o != nil && !isNil(o.Licensing) {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given OrganizationsLicensing and assigns it to the Licensing field.
func (o *InlineResponse20080) SetLicensing(v OrganizationsLicensing) {
	o.Licensing = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *InlineResponse20080) GetCloud() OrganizationsCloud {
	if o == nil || isNil(o.Cloud) {
		var ret OrganizationsCloud
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetCloudOk() (*OrganizationsCloud, bool) {
	if o == nil || isNil(o.Cloud) {
    return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *InlineResponse20080) HasCloud() bool {
	if o != nil && !isNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given OrganizationsCloud and assigns it to the Cloud field.
func (o *InlineResponse20080) SetCloud(v OrganizationsCloud) {
	o.Cloud = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *InlineResponse20080) GetManagement() OrganizationsManagement {
	if o == nil || isNil(o.Management) {
		var ret OrganizationsManagement
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetManagementOk() (*OrganizationsManagement, bool) {
	if o == nil || isNil(o.Management) {
    return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *InlineResponse20080) HasManagement() bool {
	if o != nil && !isNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given OrganizationsManagement and assigns it to the Management field.
func (o *InlineResponse20080) SetManagement(v OrganizationsManagement) {
	o.Management = &v
}

func (o InlineResponse20080) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.Licensing) {
		toSerialize["licensing"] = o.Licensing
	}
	if !isNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !isNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20080 struct {
	value *InlineResponse20080
	isSet bool
}

func (v NullableInlineResponse20080) Get() *InlineResponse20080 {
	return v.value
}

func (v *NullableInlineResponse20080) Set(val *InlineResponse20080) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20080) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20080) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20080(val *InlineResponse20080) *NullableInlineResponse20080 {
	return &NullableInlineResponse20080{value: val, isSet: true}
}

func (v NullableInlineResponse20080) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20080) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


