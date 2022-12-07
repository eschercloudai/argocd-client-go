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

// checks if the V1alpha1ApplicationSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationSource{}

// V1alpha1ApplicationSource struct for V1alpha1ApplicationSource
type V1alpha1ApplicationSource struct {
	// Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.
	Chart *string `json:"chart,omitempty"`
	Directory *V1alpha1ApplicationSourceDirectory `json:"directory,omitempty"`
	Helm *V1alpha1ApplicationSourceHelm `json:"helm,omitempty"`
	Kustomize *V1alpha1ApplicationSourceKustomize `json:"kustomize,omitempty"`
	// Path is a directory path within the Git repository, and is only valid for applications sourced from Git.
	Path *string `json:"path,omitempty"`
	Plugin *V1alpha1ApplicationSourcePlugin `json:"plugin,omitempty"`
	RepoURL *string `json:"repoURL,omitempty"`
	// TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.
	TargetRevision *string `json:"targetRevision,omitempty"`
}

// NewV1alpha1ApplicationSource instantiates a new V1alpha1ApplicationSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSource() *V1alpha1ApplicationSource {
	this := V1alpha1ApplicationSource{}
	return &this
}

// NewV1alpha1ApplicationSourceWithDefaults instantiates a new V1alpha1ApplicationSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSourceWithDefaults() *V1alpha1ApplicationSource {
	this := V1alpha1ApplicationSource{}
	return &this
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetChart() string {
	if o == nil || isNil(o.Chart) {
		var ret string
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetChartOk() (*string, bool) {
	if o == nil || isNil(o.Chart) {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasChart() bool {
	if o != nil && !isNil(o.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given string and assigns it to the Chart field.
func (o *V1alpha1ApplicationSource) SetChart(v string) {
	o.Chart = &v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetDirectory() V1alpha1ApplicationSourceDirectory {
	if o == nil || isNil(o.Directory) {
		var ret V1alpha1ApplicationSourceDirectory
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetDirectoryOk() (*V1alpha1ApplicationSourceDirectory, bool) {
	if o == nil || isNil(o.Directory) {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasDirectory() bool {
	if o != nil && !isNil(o.Directory) {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given V1alpha1ApplicationSourceDirectory and assigns it to the Directory field.
func (o *V1alpha1ApplicationSource) SetDirectory(v V1alpha1ApplicationSourceDirectory) {
	o.Directory = &v
}

// GetHelm returns the Helm field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetHelm() V1alpha1ApplicationSourceHelm {
	if o == nil || isNil(o.Helm) {
		var ret V1alpha1ApplicationSourceHelm
		return ret
	}
	return *o.Helm
}

// GetHelmOk returns a tuple with the Helm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetHelmOk() (*V1alpha1ApplicationSourceHelm, bool) {
	if o == nil || isNil(o.Helm) {
		return nil, false
	}
	return o.Helm, true
}

// HasHelm returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasHelm() bool {
	if o != nil && !isNil(o.Helm) {
		return true
	}

	return false
}

// SetHelm gets a reference to the given V1alpha1ApplicationSourceHelm and assigns it to the Helm field.
func (o *V1alpha1ApplicationSource) SetHelm(v V1alpha1ApplicationSourceHelm) {
	o.Helm = &v
}

// GetKustomize returns the Kustomize field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetKustomize() V1alpha1ApplicationSourceKustomize {
	if o == nil || isNil(o.Kustomize) {
		var ret V1alpha1ApplicationSourceKustomize
		return ret
	}
	return *o.Kustomize
}

// GetKustomizeOk returns a tuple with the Kustomize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetKustomizeOk() (*V1alpha1ApplicationSourceKustomize, bool) {
	if o == nil || isNil(o.Kustomize) {
		return nil, false
	}
	return o.Kustomize, true
}

// HasKustomize returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasKustomize() bool {
	if o != nil && !isNil(o.Kustomize) {
		return true
	}

	return false
}

// SetKustomize gets a reference to the given V1alpha1ApplicationSourceKustomize and assigns it to the Kustomize field.
func (o *V1alpha1ApplicationSource) SetKustomize(v V1alpha1ApplicationSourceKustomize) {
	o.Kustomize = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *V1alpha1ApplicationSource) SetPath(v string) {
	o.Path = &v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetPlugin() V1alpha1ApplicationSourcePlugin {
	if o == nil || isNil(o.Plugin) {
		var ret V1alpha1ApplicationSourcePlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetPluginOk() (*V1alpha1ApplicationSourcePlugin, bool) {
	if o == nil || isNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasPlugin() bool {
	if o != nil && !isNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given V1alpha1ApplicationSourcePlugin and assigns it to the Plugin field.
func (o *V1alpha1ApplicationSource) SetPlugin(v V1alpha1ApplicationSourcePlugin) {
	o.Plugin = &v
}

// GetRepoURL returns the RepoURL field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetRepoURL() string {
	if o == nil || isNil(o.RepoURL) {
		var ret string
		return ret
	}
	return *o.RepoURL
}

// GetRepoURLOk returns a tuple with the RepoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetRepoURLOk() (*string, bool) {
	if o == nil || isNil(o.RepoURL) {
		return nil, false
	}
	return o.RepoURL, true
}

// HasRepoURL returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasRepoURL() bool {
	if o != nil && !isNil(o.RepoURL) {
		return true
	}

	return false
}

// SetRepoURL gets a reference to the given string and assigns it to the RepoURL field.
func (o *V1alpha1ApplicationSource) SetRepoURL(v string) {
	o.RepoURL = &v
}

// GetTargetRevision returns the TargetRevision field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSource) GetTargetRevision() string {
	if o == nil || isNil(o.TargetRevision) {
		var ret string
		return ret
	}
	return *o.TargetRevision
}

// GetTargetRevisionOk returns a tuple with the TargetRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSource) GetTargetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.TargetRevision) {
		return nil, false
	}
	return o.TargetRevision, true
}

// HasTargetRevision returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSource) HasTargetRevision() bool {
	if o != nil && !isNil(o.TargetRevision) {
		return true
	}

	return false
}

// SetTargetRevision gets a reference to the given string and assigns it to the TargetRevision field.
func (o *V1alpha1ApplicationSource) SetTargetRevision(v string) {
	o.TargetRevision = &v
}

func (o V1alpha1ApplicationSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	if !isNil(o.Directory) {
		toSerialize["directory"] = o.Directory
	}
	if !isNil(o.Helm) {
		toSerialize["helm"] = o.Helm
	}
	if !isNil(o.Kustomize) {
		toSerialize["kustomize"] = o.Kustomize
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !isNil(o.RepoURL) {
		toSerialize["repoURL"] = o.RepoURL
	}
	if !isNil(o.TargetRevision) {
		toSerialize["targetRevision"] = o.TargetRevision
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationSource struct {
	value *V1alpha1ApplicationSource
	isSet bool
}

func (v NullableV1alpha1ApplicationSource) Get() *V1alpha1ApplicationSource {
	return v.value
}

func (v *NullableV1alpha1ApplicationSource) Set(val *V1alpha1ApplicationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSource(val *V1alpha1ApplicationSource) *NullableV1alpha1ApplicationSource {
	return &NullableV1alpha1ApplicationSource{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

