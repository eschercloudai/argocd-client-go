/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the V1alpha1AppProjectSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1AppProjectSpec{}

// V1alpha1AppProjectSpec struct for V1alpha1AppProjectSpec
type V1alpha1AppProjectSpec struct {
	ClusterResourceBlacklist        []V1GroupKind                             `json:"clusterResourceBlacklist,omitempty"`
	ClusterResourceWhitelist        []V1GroupKind                             `json:"clusterResourceWhitelist,omitempty"`
	Description                     *string                                   `json:"description,omitempty"`
	Destinations                    []V1alpha1ApplicationDestination          `json:"destinations,omitempty"`
	NamespaceResourceBlacklist      []V1GroupKind                             `json:"namespaceResourceBlacklist,omitempty"`
	NamespaceResourceWhitelist      []V1GroupKind                             `json:"namespaceResourceWhitelist,omitempty"`
	OrphanedResources               *V1alpha1OrphanedResourcesMonitorSettings `json:"orphanedResources,omitempty"`
	PermitOnlyProjectScopedClusters *bool                                     `json:"permitOnlyProjectScopedClusters,omitempty"`
	Roles                           []V1alpha1ProjectRole                     `json:"roles,omitempty"`
	SignatureKeys                   []V1alpha1SignatureKey                    `json:"signatureKeys,omitempty"`
	SourceNamespaces                []string                                  `json:"sourceNamespaces,omitempty"`
	SourceRepos                     []string                                  `json:"sourceRepos,omitempty"`
	SyncWindows                     []V1alpha1SyncWindow                      `json:"syncWindows,omitempty"`
}

// NewV1alpha1AppProjectSpec instantiates a new V1alpha1AppProjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1AppProjectSpec() *V1alpha1AppProjectSpec {
	this := V1alpha1AppProjectSpec{}
	return &this
}

// NewV1alpha1AppProjectSpecWithDefaults instantiates a new V1alpha1AppProjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1AppProjectSpecWithDefaults() *V1alpha1AppProjectSpec {
	this := V1alpha1AppProjectSpec{}
	return &this
}

// GetClusterResourceBlacklist returns the ClusterResourceBlacklist field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetClusterResourceBlacklist() []V1GroupKind {
	if o == nil || isNil(o.ClusterResourceBlacklist) {
		var ret []V1GroupKind
		return ret
	}
	return o.ClusterResourceBlacklist
}

// GetClusterResourceBlacklistOk returns a tuple with the ClusterResourceBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetClusterResourceBlacklistOk() ([]V1GroupKind, bool) {
	if o == nil || isNil(o.ClusterResourceBlacklist) {
		return nil, false
	}
	return o.ClusterResourceBlacklist, true
}

// HasClusterResourceBlacklist returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasClusterResourceBlacklist() bool {
	if o != nil && !isNil(o.ClusterResourceBlacklist) {
		return true
	}

	return false
}

// SetClusterResourceBlacklist gets a reference to the given []V1GroupKind and assigns it to the ClusterResourceBlacklist field.
func (o *V1alpha1AppProjectSpec) SetClusterResourceBlacklist(v []V1GroupKind) {
	o.ClusterResourceBlacklist = v
}

// GetClusterResourceWhitelist returns the ClusterResourceWhitelist field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetClusterResourceWhitelist() []V1GroupKind {
	if o == nil || isNil(o.ClusterResourceWhitelist) {
		var ret []V1GroupKind
		return ret
	}
	return o.ClusterResourceWhitelist
}

// GetClusterResourceWhitelistOk returns a tuple with the ClusterResourceWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetClusterResourceWhitelistOk() ([]V1GroupKind, bool) {
	if o == nil || isNil(o.ClusterResourceWhitelist) {
		return nil, false
	}
	return o.ClusterResourceWhitelist, true
}

// HasClusterResourceWhitelist returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasClusterResourceWhitelist() bool {
	if o != nil && !isNil(o.ClusterResourceWhitelist) {
		return true
	}

	return false
}

// SetClusterResourceWhitelist gets a reference to the given []V1GroupKind and assigns it to the ClusterResourceWhitelist field.
func (o *V1alpha1AppProjectSpec) SetClusterResourceWhitelist(v []V1GroupKind) {
	o.ClusterResourceWhitelist = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1alpha1AppProjectSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetDestinations() []V1alpha1ApplicationDestination {
	if o == nil || isNil(o.Destinations) {
		var ret []V1alpha1ApplicationDestination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetDestinationsOk() ([]V1alpha1ApplicationDestination, bool) {
	if o == nil || isNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []V1alpha1ApplicationDestination and assigns it to the Destinations field.
func (o *V1alpha1AppProjectSpec) SetDestinations(v []V1alpha1ApplicationDestination) {
	o.Destinations = v
}

// GetNamespaceResourceBlacklist returns the NamespaceResourceBlacklist field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetNamespaceResourceBlacklist() []V1GroupKind {
	if o == nil || isNil(o.NamespaceResourceBlacklist) {
		var ret []V1GroupKind
		return ret
	}
	return o.NamespaceResourceBlacklist
}

// GetNamespaceResourceBlacklistOk returns a tuple with the NamespaceResourceBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetNamespaceResourceBlacklistOk() ([]V1GroupKind, bool) {
	if o == nil || isNil(o.NamespaceResourceBlacklist) {
		return nil, false
	}
	return o.NamespaceResourceBlacklist, true
}

// HasNamespaceResourceBlacklist returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasNamespaceResourceBlacklist() bool {
	if o != nil && !isNil(o.NamespaceResourceBlacklist) {
		return true
	}

	return false
}

// SetNamespaceResourceBlacklist gets a reference to the given []V1GroupKind and assigns it to the NamespaceResourceBlacklist field.
func (o *V1alpha1AppProjectSpec) SetNamespaceResourceBlacklist(v []V1GroupKind) {
	o.NamespaceResourceBlacklist = v
}

// GetNamespaceResourceWhitelist returns the NamespaceResourceWhitelist field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetNamespaceResourceWhitelist() []V1GroupKind {
	if o == nil || isNil(o.NamespaceResourceWhitelist) {
		var ret []V1GroupKind
		return ret
	}
	return o.NamespaceResourceWhitelist
}

// GetNamespaceResourceWhitelistOk returns a tuple with the NamespaceResourceWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetNamespaceResourceWhitelistOk() ([]V1GroupKind, bool) {
	if o == nil || isNil(o.NamespaceResourceWhitelist) {
		return nil, false
	}
	return o.NamespaceResourceWhitelist, true
}

// HasNamespaceResourceWhitelist returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasNamespaceResourceWhitelist() bool {
	if o != nil && !isNil(o.NamespaceResourceWhitelist) {
		return true
	}

	return false
}

// SetNamespaceResourceWhitelist gets a reference to the given []V1GroupKind and assigns it to the NamespaceResourceWhitelist field.
func (o *V1alpha1AppProjectSpec) SetNamespaceResourceWhitelist(v []V1GroupKind) {
	o.NamespaceResourceWhitelist = v
}

// GetOrphanedResources returns the OrphanedResources field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetOrphanedResources() V1alpha1OrphanedResourcesMonitorSettings {
	if o == nil || isNil(o.OrphanedResources) {
		var ret V1alpha1OrphanedResourcesMonitorSettings
		return ret
	}
	return *o.OrphanedResources
}

// GetOrphanedResourcesOk returns a tuple with the OrphanedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetOrphanedResourcesOk() (*V1alpha1OrphanedResourcesMonitorSettings, bool) {
	if o == nil || isNil(o.OrphanedResources) {
		return nil, false
	}
	return o.OrphanedResources, true
}

// HasOrphanedResources returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasOrphanedResources() bool {
	if o != nil && !isNil(o.OrphanedResources) {
		return true
	}

	return false
}

// SetOrphanedResources gets a reference to the given V1alpha1OrphanedResourcesMonitorSettings and assigns it to the OrphanedResources field.
func (o *V1alpha1AppProjectSpec) SetOrphanedResources(v V1alpha1OrphanedResourcesMonitorSettings) {
	o.OrphanedResources = &v
}

// GetPermitOnlyProjectScopedClusters returns the PermitOnlyProjectScopedClusters field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetPermitOnlyProjectScopedClusters() bool {
	if o == nil || isNil(o.PermitOnlyProjectScopedClusters) {
		var ret bool
		return ret
	}
	return *o.PermitOnlyProjectScopedClusters
}

// GetPermitOnlyProjectScopedClustersOk returns a tuple with the PermitOnlyProjectScopedClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetPermitOnlyProjectScopedClustersOk() (*bool, bool) {
	if o == nil || isNil(o.PermitOnlyProjectScopedClusters) {
		return nil, false
	}
	return o.PermitOnlyProjectScopedClusters, true
}

// HasPermitOnlyProjectScopedClusters returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasPermitOnlyProjectScopedClusters() bool {
	if o != nil && !isNil(o.PermitOnlyProjectScopedClusters) {
		return true
	}

	return false
}

// SetPermitOnlyProjectScopedClusters gets a reference to the given bool and assigns it to the PermitOnlyProjectScopedClusters field.
func (o *V1alpha1AppProjectSpec) SetPermitOnlyProjectScopedClusters(v bool) {
	o.PermitOnlyProjectScopedClusters = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetRoles() []V1alpha1ProjectRole {
	if o == nil || isNil(o.Roles) {
		var ret []V1alpha1ProjectRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetRolesOk() ([]V1alpha1ProjectRole, bool) {
	if o == nil || isNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []V1alpha1ProjectRole and assigns it to the Roles field.
func (o *V1alpha1AppProjectSpec) SetRoles(v []V1alpha1ProjectRole) {
	o.Roles = v
}

// GetSignatureKeys returns the SignatureKeys field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetSignatureKeys() []V1alpha1SignatureKey {
	if o == nil || isNil(o.SignatureKeys) {
		var ret []V1alpha1SignatureKey
		return ret
	}
	return o.SignatureKeys
}

// GetSignatureKeysOk returns a tuple with the SignatureKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetSignatureKeysOk() ([]V1alpha1SignatureKey, bool) {
	if o == nil || isNil(o.SignatureKeys) {
		return nil, false
	}
	return o.SignatureKeys, true
}

// HasSignatureKeys returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasSignatureKeys() bool {
	if o != nil && !isNil(o.SignatureKeys) {
		return true
	}

	return false
}

// SetSignatureKeys gets a reference to the given []V1alpha1SignatureKey and assigns it to the SignatureKeys field.
func (o *V1alpha1AppProjectSpec) SetSignatureKeys(v []V1alpha1SignatureKey) {
	o.SignatureKeys = v
}

// GetSourceNamespaces returns the SourceNamespaces field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetSourceNamespaces() []string {
	if o == nil || isNil(o.SourceNamespaces) {
		var ret []string
		return ret
	}
	return o.SourceNamespaces
}

// GetSourceNamespacesOk returns a tuple with the SourceNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetSourceNamespacesOk() ([]string, bool) {
	if o == nil || isNil(o.SourceNamespaces) {
		return nil, false
	}
	return o.SourceNamespaces, true
}

// HasSourceNamespaces returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasSourceNamespaces() bool {
	if o != nil && !isNil(o.SourceNamespaces) {
		return true
	}

	return false
}

// SetSourceNamespaces gets a reference to the given []string and assigns it to the SourceNamespaces field.
func (o *V1alpha1AppProjectSpec) SetSourceNamespaces(v []string) {
	o.SourceNamespaces = v
}

// GetSourceRepos returns the SourceRepos field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetSourceRepos() []string {
	if o == nil || isNil(o.SourceRepos) {
		var ret []string
		return ret
	}
	return o.SourceRepos
}

// GetSourceReposOk returns a tuple with the SourceRepos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetSourceReposOk() ([]string, bool) {
	if o == nil || isNil(o.SourceRepos) {
		return nil, false
	}
	return o.SourceRepos, true
}

// HasSourceRepos returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasSourceRepos() bool {
	if o != nil && !isNil(o.SourceRepos) {
		return true
	}

	return false
}

// SetSourceRepos gets a reference to the given []string and assigns it to the SourceRepos field.
func (o *V1alpha1AppProjectSpec) SetSourceRepos(v []string) {
	o.SourceRepos = v
}

// GetSyncWindows returns the SyncWindows field value if set, zero value otherwise.
func (o *V1alpha1AppProjectSpec) GetSyncWindows() []V1alpha1SyncWindow {
	if o == nil || isNil(o.SyncWindows) {
		var ret []V1alpha1SyncWindow
		return ret
	}
	return o.SyncWindows
}

// GetSyncWindowsOk returns a tuple with the SyncWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectSpec) GetSyncWindowsOk() ([]V1alpha1SyncWindow, bool) {
	if o == nil || isNil(o.SyncWindows) {
		return nil, false
	}
	return o.SyncWindows, true
}

// HasSyncWindows returns a boolean if a field has been set.
func (o *V1alpha1AppProjectSpec) HasSyncWindows() bool {
	if o != nil && !isNil(o.SyncWindows) {
		return true
	}

	return false
}

// SetSyncWindows gets a reference to the given []V1alpha1SyncWindow and assigns it to the SyncWindows field.
func (o *V1alpha1AppProjectSpec) SetSyncWindows(v []V1alpha1SyncWindow) {
	o.SyncWindows = v
}

func (o V1alpha1AppProjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1AppProjectSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterResourceBlacklist) {
		toSerialize["clusterResourceBlacklist"] = o.ClusterResourceBlacklist
	}
	if !isNil(o.ClusterResourceWhitelist) {
		toSerialize["clusterResourceWhitelist"] = o.ClusterResourceWhitelist
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	if !isNil(o.NamespaceResourceBlacklist) {
		toSerialize["namespaceResourceBlacklist"] = o.NamespaceResourceBlacklist
	}
	if !isNil(o.NamespaceResourceWhitelist) {
		toSerialize["namespaceResourceWhitelist"] = o.NamespaceResourceWhitelist
	}
	if !isNil(o.OrphanedResources) {
		toSerialize["orphanedResources"] = o.OrphanedResources
	}
	if !isNil(o.PermitOnlyProjectScopedClusters) {
		toSerialize["permitOnlyProjectScopedClusters"] = o.PermitOnlyProjectScopedClusters
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.SignatureKeys) {
		toSerialize["signatureKeys"] = o.SignatureKeys
	}
	if !isNil(o.SourceNamespaces) {
		toSerialize["sourceNamespaces"] = o.SourceNamespaces
	}
	if !isNil(o.SourceRepos) {
		toSerialize["sourceRepos"] = o.SourceRepos
	}
	if !isNil(o.SyncWindows) {
		toSerialize["syncWindows"] = o.SyncWindows
	}
	return toSerialize, nil
}

type NullableV1alpha1AppProjectSpec struct {
	value *V1alpha1AppProjectSpec
	isSet bool
}

func (v NullableV1alpha1AppProjectSpec) Get() *V1alpha1AppProjectSpec {
	return v.value
}

func (v *NullableV1alpha1AppProjectSpec) Set(val *V1alpha1AppProjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1AppProjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1AppProjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1AppProjectSpec(val *V1alpha1AppProjectSpec) *NullableV1alpha1AppProjectSpec {
	return &NullableV1alpha1AppProjectSpec{value: val, isSet: true}
}

func (v NullableV1alpha1AppProjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1AppProjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
