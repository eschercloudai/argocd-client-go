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

// checks if the V1alpha1SyncOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncOperation{}

// V1alpha1SyncOperation SyncOperation contains details about a sync operation.
type V1alpha1SyncOperation struct {
	DryRun    *bool                           `json:"dryRun,omitempty"`
	Manifests []string                        `json:"manifests,omitempty"`
	Prune     *bool                           `json:"prune,omitempty"`
	Resources []V1alpha1SyncOperationResource `json:"resources,omitempty"`
	// Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.
	Revision     *string                    `json:"revision,omitempty"`
	Source       *V1alpha1ApplicationSource `json:"source,omitempty"`
	SyncOptions  []string                   `json:"syncOptions,omitempty"`
	SyncStrategy *V1alpha1SyncStrategy      `json:"syncStrategy,omitempty"`
}

// NewV1alpha1SyncOperation instantiates a new V1alpha1SyncOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncOperation() *V1alpha1SyncOperation {
	this := V1alpha1SyncOperation{}
	return &this
}

// NewV1alpha1SyncOperationWithDefaults instantiates a new V1alpha1SyncOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncOperationWithDefaults() *V1alpha1SyncOperation {
	this := V1alpha1SyncOperation{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetDryRun() bool {
	if o == nil || isNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetDryRunOk() (*bool, bool) {
	if o == nil || isNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasDryRun() bool {
	if o != nil && !isNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *V1alpha1SyncOperation) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetManifests returns the Manifests field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetManifests() []string {
	if o == nil || isNil(o.Manifests) {
		var ret []string
		return ret
	}
	return o.Manifests
}

// GetManifestsOk returns a tuple with the Manifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetManifestsOk() ([]string, bool) {
	if o == nil || isNil(o.Manifests) {
		return nil, false
	}
	return o.Manifests, true
}

// HasManifests returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasManifests() bool {
	if o != nil && !isNil(o.Manifests) {
		return true
	}

	return false
}

// SetManifests gets a reference to the given []string and assigns it to the Manifests field.
func (o *V1alpha1SyncOperation) SetManifests(v []string) {
	o.Manifests = v
}

// GetPrune returns the Prune field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetPrune() bool {
	if o == nil || isNil(o.Prune) {
		var ret bool
		return ret
	}
	return *o.Prune
}

// GetPruneOk returns a tuple with the Prune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetPruneOk() (*bool, bool) {
	if o == nil || isNil(o.Prune) {
		return nil, false
	}
	return o.Prune, true
}

// HasPrune returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasPrune() bool {
	if o != nil && !isNil(o.Prune) {
		return true
	}

	return false
}

// SetPrune gets a reference to the given bool and assigns it to the Prune field.
func (o *V1alpha1SyncOperation) SetPrune(v bool) {
	o.Prune = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetResources() []V1alpha1SyncOperationResource {
	if o == nil || isNil(o.Resources) {
		var ret []V1alpha1SyncOperationResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetResourcesOk() ([]V1alpha1SyncOperationResource, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []V1alpha1SyncOperationResource and assigns it to the Resources field.
func (o *V1alpha1SyncOperation) SetResources(v []V1alpha1SyncOperationResource) {
	o.Resources = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetRevision() string {
	if o == nil || isNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *V1alpha1SyncOperation) SetRevision(v string) {
	o.Revision = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetSource() V1alpha1ApplicationSource {
	if o == nil || isNil(o.Source) {
		var ret V1alpha1ApplicationSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetSourceOk() (*V1alpha1ApplicationSource, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given V1alpha1ApplicationSource and assigns it to the Source field.
func (o *V1alpha1SyncOperation) SetSource(v V1alpha1ApplicationSource) {
	o.Source = &v
}

// GetSyncOptions returns the SyncOptions field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetSyncOptions() []string {
	if o == nil || isNil(o.SyncOptions) {
		var ret []string
		return ret
	}
	return o.SyncOptions
}

// GetSyncOptionsOk returns a tuple with the SyncOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetSyncOptionsOk() ([]string, bool) {
	if o == nil || isNil(o.SyncOptions) {
		return nil, false
	}
	return o.SyncOptions, true
}

// HasSyncOptions returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasSyncOptions() bool {
	if o != nil && !isNil(o.SyncOptions) {
		return true
	}

	return false
}

// SetSyncOptions gets a reference to the given []string and assigns it to the SyncOptions field.
func (o *V1alpha1SyncOperation) SetSyncOptions(v []string) {
	o.SyncOptions = v
}

// GetSyncStrategy returns the SyncStrategy field value if set, zero value otherwise.
func (o *V1alpha1SyncOperation) GetSyncStrategy() V1alpha1SyncStrategy {
	if o == nil || isNil(o.SyncStrategy) {
		var ret V1alpha1SyncStrategy
		return ret
	}
	return *o.SyncStrategy
}

// GetSyncStrategyOk returns a tuple with the SyncStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperation) GetSyncStrategyOk() (*V1alpha1SyncStrategy, bool) {
	if o == nil || isNil(o.SyncStrategy) {
		return nil, false
	}
	return o.SyncStrategy, true
}

// HasSyncStrategy returns a boolean if a field has been set.
func (o *V1alpha1SyncOperation) HasSyncStrategy() bool {
	if o != nil && !isNil(o.SyncStrategy) {
		return true
	}

	return false
}

// SetSyncStrategy gets a reference to the given V1alpha1SyncStrategy and assigns it to the SyncStrategy field.
func (o *V1alpha1SyncOperation) SetSyncStrategy(v V1alpha1SyncStrategy) {
	o.SyncStrategy = &v
}

func (o V1alpha1SyncOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !isNil(o.Manifests) {
		toSerialize["manifests"] = o.Manifests
	}
	if !isNil(o.Prune) {
		toSerialize["prune"] = o.Prune
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.SyncOptions) {
		toSerialize["syncOptions"] = o.SyncOptions
	}
	if !isNil(o.SyncStrategy) {
		toSerialize["syncStrategy"] = o.SyncStrategy
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncOperation struct {
	value *V1alpha1SyncOperation
	isSet bool
}

func (v NullableV1alpha1SyncOperation) Get() *V1alpha1SyncOperation {
	return v.value
}

func (v *NullableV1alpha1SyncOperation) Set(val *V1alpha1SyncOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncOperation(val *V1alpha1SyncOperation) *NullableV1alpha1SyncOperation {
	return &NullableV1alpha1SyncOperation{value: val, isSet: true}
}

func (v NullableV1alpha1SyncOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
