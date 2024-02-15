/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationBrandingPolicies200ResponseInnerHelpSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationBrandingPolicies200ResponseInnerHelpSettings{}

// GetOrganizationBrandingPolicies200ResponseInnerHelpSettings       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values. 
type GetOrganizationBrandingPolicies200ResponseInnerHelpSettings struct {
	//       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpTab *string `json:"helpTab,omitempty"`
	//       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'. 
	GetHelpSubtab *string `json:"getHelpSubtab,omitempty"`
	//       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'. 
	CommunitySubtab *string `json:"communitySubtab,omitempty"`
	//       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'. 
	CasesSubtab *string `json:"casesSubtab,omitempty"`
	//       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'. 
	DataProtectionRequestsSubtab *string `json:"dataProtectionRequestsSubtab,omitempty"`
	//       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	GetHelpSubtabKnowledgeBaseSearch *string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`
	//       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'. 
	UniversalSearchKnowledgeBaseSearch *string `json:"universalSearchKnowledgeBaseSearch,omitempty"`
	//       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	CiscoMerakiProductDocumentation *string `json:"ciscoMerakiProductDocumentation,omitempty"`
	//       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string. 
	SupportContactInfo *string `json:"supportContactInfo,omitempty"`
	//       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'. 
	NewFeaturesSubtab *string `json:"newFeaturesSubtab,omitempty"`
	//       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'. 
	FirewallInfoSubtab *string `json:"firewallInfoSubtab,omitempty"`
	//       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'. 
	ApiDocsSubtab *string `json:"apiDocsSubtab,omitempty"`
	//       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'. 
	HardwareReplacementsSubtab *string `json:"hardwareReplacementsSubtab,omitempty"`
	//       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'. 
	SmForums *string `json:"smForums,omitempty"`
	//       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'. 
	HelpWidget *string `json:"helpWidget,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInnerHelpSettings instantiates a new GetOrganizationBrandingPolicies200ResponseInnerHelpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInnerHelpSettings() *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	this := GetOrganizationBrandingPolicies200ResponseInnerHelpSettings{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerHelpSettingsWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerHelpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerHelpSettingsWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	this := GetOrganizationBrandingPolicies200ResponseInnerHelpSettings{}
	return &this
}

// GetHelpTab returns the HelpTab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHelpTab() string {
	if o == nil || IsNil(o.HelpTab) {
		var ret string
		return ret
	}
	return *o.HelpTab
}

// GetHelpTabOk returns a tuple with the HelpTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHelpTabOk() (*string, bool) {
	if o == nil || IsNil(o.HelpTab) {
		return nil, false
	}
	return o.HelpTab, true
}

// HasHelpTab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasHelpTab() bool {
	if o != nil && !IsNil(o.HelpTab) {
		return true
	}

	return false
}

// SetHelpTab gets a reference to the given string and assigns it to the HelpTab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetHelpTab(v string) {
	o.HelpTab = &v
}

// GetGetHelpSubtab returns the GetHelpSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetGetHelpSubtab() string {
	if o == nil || IsNil(o.GetHelpSubtab) {
		var ret string
		return ret
	}
	return *o.GetHelpSubtab
}

// GetGetHelpSubtabOk returns a tuple with the GetHelpSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetGetHelpSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.GetHelpSubtab) {
		return nil, false
	}
	return o.GetHelpSubtab, true
}

// HasGetHelpSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasGetHelpSubtab() bool {
	if o != nil && !IsNil(o.GetHelpSubtab) {
		return true
	}

	return false
}

// SetGetHelpSubtab gets a reference to the given string and assigns it to the GetHelpSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetGetHelpSubtab(v string) {
	o.GetHelpSubtab = &v
}

// GetCommunitySubtab returns the CommunitySubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCommunitySubtab() string {
	if o == nil || IsNil(o.CommunitySubtab) {
		var ret string
		return ret
	}
	return *o.CommunitySubtab
}

// GetCommunitySubtabOk returns a tuple with the CommunitySubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCommunitySubtabOk() (*string, bool) {
	if o == nil || IsNil(o.CommunitySubtab) {
		return nil, false
	}
	return o.CommunitySubtab, true
}

// HasCommunitySubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasCommunitySubtab() bool {
	if o != nil && !IsNil(o.CommunitySubtab) {
		return true
	}

	return false
}

// SetCommunitySubtab gets a reference to the given string and assigns it to the CommunitySubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetCommunitySubtab(v string) {
	o.CommunitySubtab = &v
}

// GetCasesSubtab returns the CasesSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCasesSubtab() string {
	if o == nil || IsNil(o.CasesSubtab) {
		var ret string
		return ret
	}
	return *o.CasesSubtab
}

// GetCasesSubtabOk returns a tuple with the CasesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCasesSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.CasesSubtab) {
		return nil, false
	}
	return o.CasesSubtab, true
}

// HasCasesSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasCasesSubtab() bool {
	if o != nil && !IsNil(o.CasesSubtab) {
		return true
	}

	return false
}

// SetCasesSubtab gets a reference to the given string and assigns it to the CasesSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetCasesSubtab(v string) {
	o.CasesSubtab = &v
}

// GetDataProtectionRequestsSubtab returns the DataProtectionRequestsSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetDataProtectionRequestsSubtab() string {
	if o == nil || IsNil(o.DataProtectionRequestsSubtab) {
		var ret string
		return ret
	}
	return *o.DataProtectionRequestsSubtab
}

// GetDataProtectionRequestsSubtabOk returns a tuple with the DataProtectionRequestsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetDataProtectionRequestsSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.DataProtectionRequestsSubtab) {
		return nil, false
	}
	return o.DataProtectionRequestsSubtab, true
}

// HasDataProtectionRequestsSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasDataProtectionRequestsSubtab() bool {
	if o != nil && !IsNil(o.DataProtectionRequestsSubtab) {
		return true
	}

	return false
}

// SetDataProtectionRequestsSubtab gets a reference to the given string and assigns it to the DataProtectionRequestsSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetDataProtectionRequestsSubtab(v string) {
	o.DataProtectionRequestsSubtab = &v
}

// GetGetHelpSubtabKnowledgeBaseSearch returns the GetHelpSubtabKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetGetHelpSubtabKnowledgeBaseSearch() string {
	if o == nil || IsNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		var ret string
		return ret
	}
	return *o.GetHelpSubtabKnowledgeBaseSearch
}

// GetGetHelpSubtabKnowledgeBaseSearchOk returns a tuple with the GetHelpSubtabKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetGetHelpSubtabKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || IsNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		return nil, false
	}
	return o.GetHelpSubtabKnowledgeBaseSearch, true
}

// HasGetHelpSubtabKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasGetHelpSubtabKnowledgeBaseSearch() bool {
	if o != nil && !IsNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		return true
	}

	return false
}

// SetGetHelpSubtabKnowledgeBaseSearch gets a reference to the given string and assigns it to the GetHelpSubtabKnowledgeBaseSearch field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetGetHelpSubtabKnowledgeBaseSearch(v string) {
	o.GetHelpSubtabKnowledgeBaseSearch = &v
}

// GetUniversalSearchKnowledgeBaseSearch returns the UniversalSearchKnowledgeBaseSearch field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetUniversalSearchKnowledgeBaseSearch() string {
	if o == nil || IsNil(o.UniversalSearchKnowledgeBaseSearch) {
		var ret string
		return ret
	}
	return *o.UniversalSearchKnowledgeBaseSearch
}

// GetUniversalSearchKnowledgeBaseSearchOk returns a tuple with the UniversalSearchKnowledgeBaseSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetUniversalSearchKnowledgeBaseSearchOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalSearchKnowledgeBaseSearch) {
		return nil, false
	}
	return o.UniversalSearchKnowledgeBaseSearch, true
}

// HasUniversalSearchKnowledgeBaseSearch returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasUniversalSearchKnowledgeBaseSearch() bool {
	if o != nil && !IsNil(o.UniversalSearchKnowledgeBaseSearch) {
		return true
	}

	return false
}

// SetUniversalSearchKnowledgeBaseSearch gets a reference to the given string and assigns it to the UniversalSearchKnowledgeBaseSearch field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetUniversalSearchKnowledgeBaseSearch(v string) {
	o.UniversalSearchKnowledgeBaseSearch = &v
}

// GetCiscoMerakiProductDocumentation returns the CiscoMerakiProductDocumentation field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCiscoMerakiProductDocumentation() string {
	if o == nil || IsNil(o.CiscoMerakiProductDocumentation) {
		var ret string
		return ret
	}
	return *o.CiscoMerakiProductDocumentation
}

// GetCiscoMerakiProductDocumentationOk returns a tuple with the CiscoMerakiProductDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetCiscoMerakiProductDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.CiscoMerakiProductDocumentation) {
		return nil, false
	}
	return o.CiscoMerakiProductDocumentation, true
}

// HasCiscoMerakiProductDocumentation returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasCiscoMerakiProductDocumentation() bool {
	if o != nil && !IsNil(o.CiscoMerakiProductDocumentation) {
		return true
	}

	return false
}

// SetCiscoMerakiProductDocumentation gets a reference to the given string and assigns it to the CiscoMerakiProductDocumentation field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetCiscoMerakiProductDocumentation(v string) {
	o.CiscoMerakiProductDocumentation = &v
}

// GetSupportContactInfo returns the SupportContactInfo field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetSupportContactInfo() string {
	if o == nil || IsNil(o.SupportContactInfo) {
		var ret string
		return ret
	}
	return *o.SupportContactInfo
}

// GetSupportContactInfoOk returns a tuple with the SupportContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetSupportContactInfoOk() (*string, bool) {
	if o == nil || IsNil(o.SupportContactInfo) {
		return nil, false
	}
	return o.SupportContactInfo, true
}

// HasSupportContactInfo returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasSupportContactInfo() bool {
	if o != nil && !IsNil(o.SupportContactInfo) {
		return true
	}

	return false
}

// SetSupportContactInfo gets a reference to the given string and assigns it to the SupportContactInfo field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetSupportContactInfo(v string) {
	o.SupportContactInfo = &v
}

// GetNewFeaturesSubtab returns the NewFeaturesSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetNewFeaturesSubtab() string {
	if o == nil || IsNil(o.NewFeaturesSubtab) {
		var ret string
		return ret
	}
	return *o.NewFeaturesSubtab
}

// GetNewFeaturesSubtabOk returns a tuple with the NewFeaturesSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetNewFeaturesSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.NewFeaturesSubtab) {
		return nil, false
	}
	return o.NewFeaturesSubtab, true
}

// HasNewFeaturesSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasNewFeaturesSubtab() bool {
	if o != nil && !IsNil(o.NewFeaturesSubtab) {
		return true
	}

	return false
}

// SetNewFeaturesSubtab gets a reference to the given string and assigns it to the NewFeaturesSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetNewFeaturesSubtab(v string) {
	o.NewFeaturesSubtab = &v
}

// GetFirewallInfoSubtab returns the FirewallInfoSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetFirewallInfoSubtab() string {
	if o == nil || IsNil(o.FirewallInfoSubtab) {
		var ret string
		return ret
	}
	return *o.FirewallInfoSubtab
}

// GetFirewallInfoSubtabOk returns a tuple with the FirewallInfoSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetFirewallInfoSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.FirewallInfoSubtab) {
		return nil, false
	}
	return o.FirewallInfoSubtab, true
}

// HasFirewallInfoSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasFirewallInfoSubtab() bool {
	if o != nil && !IsNil(o.FirewallInfoSubtab) {
		return true
	}

	return false
}

// SetFirewallInfoSubtab gets a reference to the given string and assigns it to the FirewallInfoSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetFirewallInfoSubtab(v string) {
	o.FirewallInfoSubtab = &v
}

// GetApiDocsSubtab returns the ApiDocsSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetApiDocsSubtab() string {
	if o == nil || IsNil(o.ApiDocsSubtab) {
		var ret string
		return ret
	}
	return *o.ApiDocsSubtab
}

// GetApiDocsSubtabOk returns a tuple with the ApiDocsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetApiDocsSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.ApiDocsSubtab) {
		return nil, false
	}
	return o.ApiDocsSubtab, true
}

// HasApiDocsSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasApiDocsSubtab() bool {
	if o != nil && !IsNil(o.ApiDocsSubtab) {
		return true
	}

	return false
}

// SetApiDocsSubtab gets a reference to the given string and assigns it to the ApiDocsSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetApiDocsSubtab(v string) {
	o.ApiDocsSubtab = &v
}

// GetHardwareReplacementsSubtab returns the HardwareReplacementsSubtab field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHardwareReplacementsSubtab() string {
	if o == nil || IsNil(o.HardwareReplacementsSubtab) {
		var ret string
		return ret
	}
	return *o.HardwareReplacementsSubtab
}

// GetHardwareReplacementsSubtabOk returns a tuple with the HardwareReplacementsSubtab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHardwareReplacementsSubtabOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareReplacementsSubtab) {
		return nil, false
	}
	return o.HardwareReplacementsSubtab, true
}

// HasHardwareReplacementsSubtab returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasHardwareReplacementsSubtab() bool {
	if o != nil && !IsNil(o.HardwareReplacementsSubtab) {
		return true
	}

	return false
}

// SetHardwareReplacementsSubtab gets a reference to the given string and assigns it to the HardwareReplacementsSubtab field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetHardwareReplacementsSubtab(v string) {
	o.HardwareReplacementsSubtab = &v
}

// GetSmForums returns the SmForums field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetSmForums() string {
	if o == nil || IsNil(o.SmForums) {
		var ret string
		return ret
	}
	return *o.SmForums
}

// GetSmForumsOk returns a tuple with the SmForums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetSmForumsOk() (*string, bool) {
	if o == nil || IsNil(o.SmForums) {
		return nil, false
	}
	return o.SmForums, true
}

// HasSmForums returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasSmForums() bool {
	if o != nil && !IsNil(o.SmForums) {
		return true
	}

	return false
}

// SetSmForums gets a reference to the given string and assigns it to the SmForums field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetSmForums(v string) {
	o.SmForums = &v
}

// GetHelpWidget returns the HelpWidget field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHelpWidget() string {
	if o == nil || IsNil(o.HelpWidget) {
		var ret string
		return ret
	}
	return *o.HelpWidget
}

// GetHelpWidgetOk returns a tuple with the HelpWidget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) GetHelpWidgetOk() (*string, bool) {
	if o == nil || IsNil(o.HelpWidget) {
		return nil, false
	}
	return o.HelpWidget, true
}

// HasHelpWidget returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) HasHelpWidget() bool {
	if o != nil && !IsNil(o.HelpWidget) {
		return true
	}

	return false
}

// SetHelpWidget gets a reference to the given string and assigns it to the HelpWidget field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) SetHelpWidget(v string) {
	o.HelpWidget = &v
}

func (o GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HelpTab) {
		toSerialize["helpTab"] = o.HelpTab
	}
	if !IsNil(o.GetHelpSubtab) {
		toSerialize["getHelpSubtab"] = o.GetHelpSubtab
	}
	if !IsNil(o.CommunitySubtab) {
		toSerialize["communitySubtab"] = o.CommunitySubtab
	}
	if !IsNil(o.CasesSubtab) {
		toSerialize["casesSubtab"] = o.CasesSubtab
	}
	if !IsNil(o.DataProtectionRequestsSubtab) {
		toSerialize["dataProtectionRequestsSubtab"] = o.DataProtectionRequestsSubtab
	}
	if !IsNil(o.GetHelpSubtabKnowledgeBaseSearch) {
		toSerialize["getHelpSubtabKnowledgeBaseSearch"] = o.GetHelpSubtabKnowledgeBaseSearch
	}
	if !IsNil(o.UniversalSearchKnowledgeBaseSearch) {
		toSerialize["universalSearchKnowledgeBaseSearch"] = o.UniversalSearchKnowledgeBaseSearch
	}
	if !IsNil(o.CiscoMerakiProductDocumentation) {
		toSerialize["ciscoMerakiProductDocumentation"] = o.CiscoMerakiProductDocumentation
	}
	if !IsNil(o.SupportContactInfo) {
		toSerialize["supportContactInfo"] = o.SupportContactInfo
	}
	if !IsNil(o.NewFeaturesSubtab) {
		toSerialize["newFeaturesSubtab"] = o.NewFeaturesSubtab
	}
	if !IsNil(o.FirewallInfoSubtab) {
		toSerialize["firewallInfoSubtab"] = o.FirewallInfoSubtab
	}
	if !IsNil(o.ApiDocsSubtab) {
		toSerialize["apiDocsSubtab"] = o.ApiDocsSubtab
	}
	if !IsNil(o.HardwareReplacementsSubtab) {
		toSerialize["hardwareReplacementsSubtab"] = o.HardwareReplacementsSubtab
	}
	if !IsNil(o.SmForums) {
		toSerialize["smForums"] = o.SmForums
	}
	if !IsNil(o.HelpWidget) {
		toSerialize["helpWidget"] = o.HelpWidget
	}
	return toSerialize, nil
}

type NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings struct {
	value *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) Get() *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) Set(val *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings(val *GetOrganizationBrandingPolicies200ResponseInnerHelpSettings) *NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings {
	return &NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerHelpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


