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

// checks if the GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner{}

// GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner struct for GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner
type GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner struct {
	// The name of the VPN peer
	Name *string `json:"name,omitempty"`
	// [optional] The public IP of the VPN peer
	PublicIp *string `json:"publicIp,omitempty"`
	// [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	RemoteId *string `json:"remoteId,omitempty"`
	// [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	LocalId *string `json:"localId,omitempty"`
	// The shared secret with the VPN peer
	Secret *string `json:"secret,omitempty"`
	// The list of the private subnets of the VPN peer
	PrivateSubnets []string `json:"privateSubnets,omitempty"`
	IpsecPolicies *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies `json:"ipsecPolicies,omitempty"`
	// One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	IpsecPoliciesPreset *string `json:"ipsecPoliciesPreset,omitempty"`
	// [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IkeVersion *string `json:"ikeVersion,omitempty"`
	// A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	NetworkTags []string `json:"networkTags,omitempty"`
}

// NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner {
	this := GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerWithDefaults instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerWithDefaults() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner {
	this := GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetName(v string) {
	o.Name = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetLocalId returns the LocalId field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetLocalId() string {
	if o == nil || IsNil(o.LocalId) {
		var ret string
		return ret
	}
	return *o.LocalId
}

// GetLocalIdOk returns a tuple with the LocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetLocalIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalId) {
		return nil, false
	}
	return o.LocalId, true
}

// HasLocalId returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasLocalId() bool {
	if o != nil && !IsNil(o.LocalId) {
		return true
	}

	return false
}

// SetLocalId gets a reference to the given string and assigns it to the LocalId field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetLocalId(v string) {
	o.LocalId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetSecret(v string) {
	o.Secret = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPrivateSubnets() []string {
	if o == nil || IsNil(o.PrivateSubnets) {
		var ret []string
		return ret
	}
	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivateSubnets) {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// HasPrivateSubnets returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasPrivateSubnets() bool {
	if o != nil && !IsNil(o.PrivateSubnets) {
		return true
	}

	return false
}

// SetPrivateSubnets gets a reference to the given []string and assigns it to the PrivateSubnets field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

// GetIpsecPolicies returns the IpsecPolicies field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPolicies() GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies {
	if o == nil || IsNil(o.IpsecPolicies) {
		var ret GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies
		return ret
	}
	return *o.IpsecPolicies
}

// GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesOk() (*GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies, bool) {
	if o == nil || IsNil(o.IpsecPolicies) {
		return nil, false
	}
	return o.IpsecPolicies, true
}

// HasIpsecPolicies returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIpsecPolicies() bool {
	if o != nil && !IsNil(o.IpsecPolicies) {
		return true
	}

	return false
}

// SetIpsecPolicies gets a reference to the given GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies and assigns it to the IpsecPolicies field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIpsecPolicies(v GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) {
	o.IpsecPolicies = &v
}

// GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesPreset() string {
	if o == nil || IsNil(o.IpsecPoliciesPreset) {
		var ret string
		return ret
	}
	return *o.IpsecPoliciesPreset
}

// GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesPresetOk() (*string, bool) {
	if o == nil || IsNil(o.IpsecPoliciesPreset) {
		return nil, false
	}
	return o.IpsecPoliciesPreset, true
}

// HasIpsecPoliciesPreset returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIpsecPoliciesPreset() bool {
	if o != nil && !IsNil(o.IpsecPoliciesPreset) {
		return true
	}

	return false
}

// SetIpsecPoliciesPreset gets a reference to the given string and assigns it to the IpsecPoliciesPreset field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIpsecPoliciesPreset(v string) {
	o.IpsecPoliciesPreset = &v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIkeVersion() string {
	if o == nil || IsNil(o.IkeVersion) {
		var ret string
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIkeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.IkeVersion) {
		return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIkeVersion() bool {
	if o != nil && !IsNil(o.IkeVersion) {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given string and assigns it to the IkeVersion field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIkeVersion(v string) {
	o.IkeVersion = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNetworkTags() []string {
	if o == nil || IsNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkTags) {
		return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasNetworkTags() bool {
	if o != nil && !IsNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

func (o GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !IsNil(o.LocalId) {
		toSerialize["localId"] = o.LocalId
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.PrivateSubnets) {
		toSerialize["privateSubnets"] = o.PrivateSubnets
	}
	if !IsNil(o.IpsecPolicies) {
		toSerialize["ipsecPolicies"] = o.IpsecPolicies
	}
	if !IsNil(o.IpsecPoliciesPreset) {
		toSerialize["ipsecPoliciesPreset"] = o.IpsecPoliciesPreset
	}
	if !IsNil(o.IkeVersion) {
		toSerialize["ikeVersion"] = o.IkeVersion
	}
	if !IsNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner struct {
	value *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner
	isSet bool
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) Get() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner {
	return v.value
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) Set(val *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner(val *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner {
	return &NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


