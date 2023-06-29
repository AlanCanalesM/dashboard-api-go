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

// checks if the GetOrganizationSamlIdps200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSamlIdps200ResponseInner{}

// GetOrganizationSamlIdps200ResponseInner struct for GetOrganizationSamlIdps200ResponseInner
type GetOrganizationSamlIdps200ResponseInner struct {
	// ID associated with the SAML Identity Provider (IdP)
	IdpId *string `json:"idpId,omitempty"`
	// URL that is consuming SAML Identity Provider (IdP)
	ConsumerUrl *string `json:"consumerUrl,omitempty"`
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewGetOrganizationSamlIdps200ResponseInner instantiates a new GetOrganizationSamlIdps200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSamlIdps200ResponseInner() *GetOrganizationSamlIdps200ResponseInner {
	this := GetOrganizationSamlIdps200ResponseInner{}
	return &this
}

// NewGetOrganizationSamlIdps200ResponseInnerWithDefaults instantiates a new GetOrganizationSamlIdps200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSamlIdps200ResponseInnerWithDefaults() *GetOrganizationSamlIdps200ResponseInner {
	this := GetOrganizationSamlIdps200ResponseInner{}
	return &this
}

// GetIdpId returns the IdpId field value if set, zero value otherwise.
func (o *GetOrganizationSamlIdps200ResponseInner) GetIdpId() string {
	if o == nil || IsNil(o.IdpId) {
		var ret string
		return ret
	}
	return *o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) GetIdpIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdpId) {
		return nil, false
	}
	return o.IdpId, true
}

// HasIdpId returns a boolean if a field has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) HasIdpId() bool {
	if o != nil && !IsNil(o.IdpId) {
		return true
	}

	return false
}

// SetIdpId gets a reference to the given string and assigns it to the IdpId field.
func (o *GetOrganizationSamlIdps200ResponseInner) SetIdpId(v string) {
	o.IdpId = &v
}

// GetConsumerUrl returns the ConsumerUrl field value if set, zero value otherwise.
func (o *GetOrganizationSamlIdps200ResponseInner) GetConsumerUrl() string {
	if o == nil || IsNil(o.ConsumerUrl) {
		var ret string
		return ret
	}
	return *o.ConsumerUrl
}

// GetConsumerUrlOk returns a tuple with the ConsumerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) GetConsumerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerUrl) {
		return nil, false
	}
	return o.ConsumerUrl, true
}

// HasConsumerUrl returns a boolean if a field has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) HasConsumerUrl() bool {
	if o != nil && !IsNil(o.ConsumerUrl) {
		return true
	}

	return false
}

// SetConsumerUrl gets a reference to the given string and assigns it to the ConsumerUrl field.
func (o *GetOrganizationSamlIdps200ResponseInner) SetConsumerUrl(v string) {
	o.ConsumerUrl = &v
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *GetOrganizationSamlIdps200ResponseInner) GetX509certSha1Fingerprint() string {
	if o == nil || IsNil(o.X509certSha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.X509certSha1Fingerprint) {
		return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) HasX509certSha1Fingerprint() bool {
	if o != nil && !IsNil(o.X509certSha1Fingerprint) {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *GetOrganizationSamlIdps200ResponseInner) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *GetOrganizationSamlIdps200ResponseInner) GetSloLogoutUrl() string {
	if o == nil || IsNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SloLogoutUrl) {
		return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *GetOrganizationSamlIdps200ResponseInner) HasSloLogoutUrl() bool {
	if o != nil && !IsNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *GetOrganizationSamlIdps200ResponseInner) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o GetOrganizationSamlIdps200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSamlIdps200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdpId) {
		toSerialize["idpId"] = o.IdpId
	}
	if !IsNil(o.ConsumerUrl) {
		toSerialize["consumerUrl"] = o.ConsumerUrl
	}
	if !IsNil(o.X509certSha1Fingerprint) {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if !IsNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return toSerialize, nil
}

type NullableGetOrganizationSamlIdps200ResponseInner struct {
	value *GetOrganizationSamlIdps200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSamlIdps200ResponseInner) Get() *GetOrganizationSamlIdps200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSamlIdps200ResponseInner) Set(val *GetOrganizationSamlIdps200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSamlIdps200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSamlIdps200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSamlIdps200ResponseInner(val *GetOrganizationSamlIdps200ResponseInner) *NullableGetOrganizationSamlIdps200ResponseInner {
	return &NullableGetOrganizationSamlIdps200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSamlIdps200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSamlIdps200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


