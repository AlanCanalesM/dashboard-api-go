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

// OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers struct for OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers
type OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers struct {
	// The name of the VPN peer
	Name string `json:"name"`
	// [optional] The public IP of the VPN peer
	PublicIp *string `json:"publicIp,omitempty"`
	// The list of the private subnets of the VPN peer
	PrivateSubnets []string `json:"privateSubnets"`
	// [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	LocalId *string `json:"localId,omitempty"`
	// [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	RemoteId *string `json:"remoteId,omitempty"`
	IpsecPolicies *InlineResponse20084IpsecPolicies `json:"ipsecPolicies,omitempty"`
	// One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	IpsecPoliciesPreset *string `json:"ipsecPoliciesPreset,omitempty"`
	// The shared secret with the VPN peer
	Secret string `json:"secret"`
	// [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IkeVersion *string `json:"ikeVersion,omitempty"`
	// A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	NetworkTags []string `json:"networkTags,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers(name string, privateSubnets []string, secret string) *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	this := OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers{}
	this.Name = name
	this.PrivateSubnets = privateSubnets
	this.Secret = secret
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeersWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeersWithDefaults() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	this := OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// GetName returns the Name field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetName(v string) {
	o.Name = v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetPrivateSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrivateSubnets, true
}

// SetPrivateSubnets sets field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

// GetLocalId returns the LocalId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetLocalId() string {
	if o == nil || isNil(o.LocalId) {
		var ret string
		return ret
	}
	return *o.LocalId
}

// GetLocalIdOk returns a tuple with the LocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetLocalIdOk() (*string, bool) {
	if o == nil || isNil(o.LocalId) {
    return nil, false
	}
	return o.LocalId, true
}

// HasLocalId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasLocalId() bool {
	if o != nil && !isNil(o.LocalId) {
		return true
	}

	return false
}

// SetLocalId gets a reference to the given string and assigns it to the LocalId field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetLocalId(v string) {
	o.LocalId = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetRemoteId() string {
	if o == nil || isNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetRemoteIdOk() (*string, bool) {
	if o == nil || isNil(o.RemoteId) {
    return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasRemoteId() bool {
	if o != nil && !isNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetIpsecPolicies returns the IpsecPolicies field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIpsecPolicies() InlineResponse20084IpsecPolicies {
	if o == nil || isNil(o.IpsecPolicies) {
		var ret InlineResponse20084IpsecPolicies
		return ret
	}
	return *o.IpsecPolicies
}

// GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIpsecPoliciesOk() (*InlineResponse20084IpsecPolicies, bool) {
	if o == nil || isNil(o.IpsecPolicies) {
    return nil, false
	}
	return o.IpsecPolicies, true
}

// HasIpsecPolicies returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasIpsecPolicies() bool {
	if o != nil && !isNil(o.IpsecPolicies) {
		return true
	}

	return false
}

// SetIpsecPolicies gets a reference to the given InlineResponse20084IpsecPolicies and assigns it to the IpsecPolicies field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetIpsecPolicies(v InlineResponse20084IpsecPolicies) {
	o.IpsecPolicies = &v
}

// GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIpsecPoliciesPreset() string {
	if o == nil || isNil(o.IpsecPoliciesPreset) {
		var ret string
		return ret
	}
	return *o.IpsecPoliciesPreset
}

// GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIpsecPoliciesPresetOk() (*string, bool) {
	if o == nil || isNil(o.IpsecPoliciesPreset) {
    return nil, false
	}
	return o.IpsecPoliciesPreset, true
}

// HasIpsecPoliciesPreset returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasIpsecPoliciesPreset() bool {
	if o != nil && !isNil(o.IpsecPoliciesPreset) {
		return true
	}

	return false
}

// SetIpsecPoliciesPreset gets a reference to the given string and assigns it to the IpsecPoliciesPreset field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetIpsecPoliciesPreset(v string) {
	o.IpsecPoliciesPreset = &v
}

// GetSecret returns the Secret field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetSecret(v string) {
	o.Secret = v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIkeVersion() string {
	if o == nil || isNil(o.IkeVersion) {
		var ret string
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetIkeVersionOk() (*string, bool) {
	if o == nil || isNil(o.IkeVersion) {
    return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasIkeVersion() bool {
	if o != nil && !isNil(o.IkeVersion) {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given string and assigns it to the IkeVersion field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetIkeVersion(v string) {
	o.IkeVersion = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetNetworkTags() []string {
	if o == nil || isNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkTags) {
    return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) HasNetworkTags() bool {
	if o != nil && !isNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

func (o OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if true {
		toSerialize["privateSubnets"] = o.PrivateSubnets
	}
	if !isNil(o.LocalId) {
		toSerialize["localId"] = o.LocalId
	}
	if !isNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !isNil(o.IpsecPolicies) {
		toSerialize["ipsecPolicies"] = o.IpsecPolicies
	}
	if !isNil(o.IpsecPoliciesPreset) {
		toSerialize["ipsecPoliciesPreset"] = o.IpsecPoliciesPreset
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.IkeVersion) {
		toSerialize["ikeVersion"] = o.IkeVersion
	}
	if !isNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers struct {
	value *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) Get() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) Set(val *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers(val *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	return &NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


