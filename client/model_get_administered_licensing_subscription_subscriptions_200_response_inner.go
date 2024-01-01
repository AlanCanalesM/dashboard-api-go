/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner{}

// GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner struct for GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
type GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner struct {
	// Subscription's ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Subscription name
	Name *string `json:"name,omitempty"`
	// Subscription description
	Description *string `json:"description,omitempty"`
	// Subscription status
	Status *string `json:"status,omitempty"`
	// Subscription start date
	StartDate *time.Time `json:"startDate,omitempty"`
	// Subscription expiration date
	EndDate *time.Time `json:"endDate,omitempty"`
	// Web order id
	WebOrderId *string `json:"webOrderId,omitempty"`
	// Products the subscription has entitlements for
	ProductTypes []string `json:"productTypes,omitempty"`
	// Entitlement info
	Entitlements []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner `json:"entitlements,omitempty"`
	Counts *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts `json:"counts,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetWebOrderId returns the WebOrderId field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetWebOrderId() string {
	if o == nil || IsNil(o.WebOrderId) {
		var ret string
		return ret
	}
	return *o.WebOrderId
}

// GetWebOrderIdOk returns a tuple with the WebOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetWebOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebOrderId) {
		return nil, false
	}
	return o.WebOrderId, true
}

// HasWebOrderId returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasWebOrderId() bool {
	if o != nil && !IsNil(o.WebOrderId) {
		return true
	}

	return false
}

// SetWebOrderId gets a reference to the given string and assigns it to the WebOrderId field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetWebOrderId(v string) {
	o.WebOrderId = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetProductTypes() []string {
	if o == nil || IsNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasProductTypes() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEntitlements() []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner {
	if o == nil || IsNil(o.Entitlements) {
		var ret []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEntitlementsOk() ([]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner and assigns it to the Entitlements field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetEntitlements(v []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner) {
	o.Entitlements = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetCounts() GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetCountsOk() (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts and assigns it to the Counts field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetCounts(v GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) {
	o.Counts = &v
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.WebOrderId) {
		toSerialize["webOrderId"] = o.WebOrderId
	}
	if !IsNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner struct {
	value *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) Get() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) Set(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner {
	return &NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


