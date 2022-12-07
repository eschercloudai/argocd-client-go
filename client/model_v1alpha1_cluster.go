/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the V1alpha1Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1Cluster{}

// V1alpha1Cluster struct for V1alpha1Cluster
type V1alpha1Cluster struct {
	Annotations *map[string]string `json:"annotations,omitempty"`
	// Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode.
	ClusterResources *bool `json:"clusterResources,omitempty"`
	Config *V1alpha1ClusterConfig `json:"config,omitempty"`
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`
	Info *V1alpha1ClusterInfo `json:"info,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Name *string `json:"name,omitempty"`
	// Holds list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty.
	Namespaces []string `json:"namespaces,omitempty"`
	Project *string `json:"project,omitempty"`
	RefreshRequestedAt *V1Time `json:"refreshRequestedAt,omitempty"`
	Server *string `json:"server,omitempty"`
	ServerVersion *string `json:"serverVersion,omitempty"`
	// Shard contains optional shard number. Calculated on the fly by the application controller if not specified.
	Shard *string `json:"shard,omitempty"`
}

// NewV1alpha1Cluster instantiates a new V1alpha1Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1Cluster() *V1alpha1Cluster {
	this := V1alpha1Cluster{}
	return &this
}

// NewV1alpha1ClusterWithDefaults instantiates a new V1alpha1Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ClusterWithDefaults() *V1alpha1Cluster {
	this := V1alpha1Cluster{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetAnnotations() map[string]string {
	if o == nil || isNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *V1alpha1Cluster) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetClusterResources returns the ClusterResources field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetClusterResources() bool {
	if o == nil || isNil(o.ClusterResources) {
		var ret bool
		return ret
	}
	return *o.ClusterResources
}

// GetClusterResourcesOk returns a tuple with the ClusterResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetClusterResourcesOk() (*bool, bool) {
	if o == nil || isNil(o.ClusterResources) {
		return nil, false
	}
	return o.ClusterResources, true
}

// HasClusterResources returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasClusterResources() bool {
	if o != nil && !isNil(o.ClusterResources) {
		return true
	}

	return false
}

// SetClusterResources gets a reference to the given bool and assigns it to the ClusterResources field.
func (o *V1alpha1Cluster) SetClusterResources(v bool) {
	o.ClusterResources = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetConfig() V1alpha1ClusterConfig {
	if o == nil || isNil(o.Config) {
		var ret V1alpha1ClusterConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetConfigOk() (*V1alpha1ClusterConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given V1alpha1ClusterConfig and assigns it to the Config field.
func (o *V1alpha1Cluster) SetConfig(v V1alpha1ClusterConfig) {
	o.Config = &v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetConnectionState() V1alpha1ConnectionState {
	if o == nil || isNil(o.ConnectionState) {
		var ret V1alpha1ConnectionState
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetConnectionStateOk() (*V1alpha1ConnectionState, bool) {
	if o == nil || isNil(o.ConnectionState) {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasConnectionState() bool {
	if o != nil && !isNil(o.ConnectionState) {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given V1alpha1ConnectionState and assigns it to the ConnectionState field.
func (o *V1alpha1Cluster) SetConnectionState(v V1alpha1ConnectionState) {
	o.ConnectionState = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetInfo() V1alpha1ClusterInfo {
	if o == nil || isNil(o.Info) {
		var ret V1alpha1ClusterInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetInfoOk() (*V1alpha1ClusterInfo, bool) {
	if o == nil || isNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasInfo() bool {
	if o != nil && !isNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given V1alpha1ClusterInfo and assigns it to the Info field.
func (o *V1alpha1Cluster) SetInfo(v V1alpha1ClusterInfo) {
	o.Info = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetLabels() map[string]string {
	if o == nil || isNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasLabels() bool {
	if o != nil && !isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *V1alpha1Cluster) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1Cluster) SetName(v string) {
	o.Name = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetNamespaces() []string {
	if o == nil || isNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetNamespacesOk() ([]string, bool) {
	if o == nil || isNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasNamespaces() bool {
	if o != nil && !isNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *V1alpha1Cluster) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetProject() string {
	if o == nil || isNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetProjectOk() (*string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *V1alpha1Cluster) SetProject(v string) {
	o.Project = &v
}

// GetRefreshRequestedAt returns the RefreshRequestedAt field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetRefreshRequestedAt() V1Time {
	if o == nil || isNil(o.RefreshRequestedAt) {
		var ret V1Time
		return ret
	}
	return *o.RefreshRequestedAt
}

// GetRefreshRequestedAtOk returns a tuple with the RefreshRequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetRefreshRequestedAtOk() (*V1Time, bool) {
	if o == nil || isNil(o.RefreshRequestedAt) {
		return nil, false
	}
	return o.RefreshRequestedAt, true
}

// HasRefreshRequestedAt returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasRefreshRequestedAt() bool {
	if o != nil && !isNil(o.RefreshRequestedAt) {
		return true
	}

	return false
}

// SetRefreshRequestedAt gets a reference to the given V1Time and assigns it to the RefreshRequestedAt field.
func (o *V1alpha1Cluster) SetRefreshRequestedAt(v V1Time) {
	o.RefreshRequestedAt = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetServer() string {
	if o == nil || isNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetServerOk() (*string, bool) {
	if o == nil || isNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *V1alpha1Cluster) SetServer(v string) {
	o.Server = &v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetServerVersion() string {
	if o == nil || isNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetServerVersionOk() (*string, bool) {
	if o == nil || isNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasServerVersion() bool {
	if o != nil && !isNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *V1alpha1Cluster) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetShard returns the Shard field value if set, zero value otherwise.
func (o *V1alpha1Cluster) GetShard() string {
	if o == nil || isNil(o.Shard) {
		var ret string
		return ret
	}
	return *o.Shard
}

// GetShardOk returns a tuple with the Shard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Cluster) GetShardOk() (*string, bool) {
	if o == nil || isNil(o.Shard) {
		return nil, false
	}
	return o.Shard, true
}

// HasShard returns a boolean if a field has been set.
func (o *V1alpha1Cluster) HasShard() bool {
	if o != nil && !isNil(o.Shard) {
		return true
	}

	return false
}

// SetShard gets a reference to the given string and assigns it to the Shard field.
func (o *V1alpha1Cluster) SetShard(v string) {
	o.Shard = &v
}

func (o V1alpha1Cluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !isNil(o.ClusterResources) {
		toSerialize["clusterResources"] = o.ClusterResources
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.ConnectionState) {
		toSerialize["connectionState"] = o.ConnectionState
	}
	if !isNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !isNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.RefreshRequestedAt) {
		toSerialize["refreshRequestedAt"] = o.RefreshRequestedAt
	}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !isNil(o.ServerVersion) {
		toSerialize["serverVersion"] = o.ServerVersion
	}
	if !isNil(o.Shard) {
		toSerialize["shard"] = o.Shard
	}
	return toSerialize, nil
}

type NullableV1alpha1Cluster struct {
	value *V1alpha1Cluster
	isSet bool
}

func (v NullableV1alpha1Cluster) Get() *V1alpha1Cluster {
	return v.value
}

func (v *NullableV1alpha1Cluster) Set(val *V1alpha1Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1Cluster) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1Cluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1Cluster(val *V1alpha1Cluster) *NullableV1alpha1Cluster {
	return &NullableV1alpha1Cluster{value: val, isSet: true}
}

func (v NullableV1alpha1Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1Cluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


