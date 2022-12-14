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

// checks if the ApplicationApplicationSyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationApplicationSyncRequest{}

// ApplicationApplicationSyncRequest struct for ApplicationApplicationSyncRequest
type ApplicationApplicationSyncRequest struct {
	AppNamespace  *string                         `json:"appNamespace,omitempty"`
	DryRun        *bool                           `json:"dryRun,omitempty"`
	Infos         []V1alpha1Info                  `json:"infos,omitempty"`
	Manifests     []string                        `json:"manifests,omitempty"`
	Name          *string                         `json:"name,omitempty"`
	Prune         *bool                           `json:"prune,omitempty"`
	Resources     []V1alpha1SyncOperationResource `json:"resources,omitempty"`
	RetryStrategy *V1alpha1RetryStrategy          `json:"retryStrategy,omitempty"`
	Revision      *string                         `json:"revision,omitempty"`
	Strategy      *V1alpha1SyncStrategy           `json:"strategy,omitempty"`
	SyncOptions   *ApplicationSyncOptions         `json:"syncOptions,omitempty"`
}

// NewApplicationApplicationSyncRequest instantiates a new ApplicationApplicationSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationApplicationSyncRequest() *ApplicationApplicationSyncRequest {
	this := ApplicationApplicationSyncRequest{}
	return &this
}

// NewApplicationApplicationSyncRequestWithDefaults instantiates a new ApplicationApplicationSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationApplicationSyncRequestWithDefaults() *ApplicationApplicationSyncRequest {
	this := ApplicationApplicationSyncRequest{}
	return &this
}

// GetAppNamespace returns the AppNamespace field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetAppNamespace() string {
	if o == nil || isNil(o.AppNamespace) {
		var ret string
		return ret
	}
	return *o.AppNamespace
}

// GetAppNamespaceOk returns a tuple with the AppNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetAppNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.AppNamespace) {
		return nil, false
	}
	return o.AppNamespace, true
}

// HasAppNamespace returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasAppNamespace() bool {
	if o != nil && !isNil(o.AppNamespace) {
		return true
	}

	return false
}

// SetAppNamespace gets a reference to the given string and assigns it to the AppNamespace field.
func (o *ApplicationApplicationSyncRequest) SetAppNamespace(v string) {
	o.AppNamespace = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetDryRun() bool {
	if o == nil || isNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || isNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasDryRun() bool {
	if o != nil && !isNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ApplicationApplicationSyncRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetInfos returns the Infos field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetInfos() []V1alpha1Info {
	if o == nil || isNil(o.Infos) {
		var ret []V1alpha1Info
		return ret
	}
	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetInfosOk() ([]V1alpha1Info, bool) {
	if o == nil || isNil(o.Infos) {
		return nil, false
	}
	return o.Infos, true
}

// HasInfos returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasInfos() bool {
	if o != nil && !isNil(o.Infos) {
		return true
	}

	return false
}

// SetInfos gets a reference to the given []V1alpha1Info and assigns it to the Infos field.
func (o *ApplicationApplicationSyncRequest) SetInfos(v []V1alpha1Info) {
	o.Infos = v
}

// GetManifests returns the Manifests field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetManifests() []string {
	if o == nil || isNil(o.Manifests) {
		var ret []string
		return ret
	}
	return o.Manifests
}

// GetManifestsOk returns a tuple with the Manifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetManifestsOk() ([]string, bool) {
	if o == nil || isNil(o.Manifests) {
		return nil, false
	}
	return o.Manifests, true
}

// HasManifests returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasManifests() bool {
	if o != nil && !isNil(o.Manifests) {
		return true
	}

	return false
}

// SetManifests gets a reference to the given []string and assigns it to the Manifests field.
func (o *ApplicationApplicationSyncRequest) SetManifests(v []string) {
	o.Manifests = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationApplicationSyncRequest) SetName(v string) {
	o.Name = &v
}

// GetPrune returns the Prune field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetPrune() bool {
	if o == nil || isNil(o.Prune) {
		var ret bool
		return ret
	}
	return *o.Prune
}

// GetPruneOk returns a tuple with the Prune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetPruneOk() (*bool, bool) {
	if o == nil || isNil(o.Prune) {
		return nil, false
	}
	return o.Prune, true
}

// HasPrune returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasPrune() bool {
	if o != nil && !isNil(o.Prune) {
		return true
	}

	return false
}

// SetPrune gets a reference to the given bool and assigns it to the Prune field.
func (o *ApplicationApplicationSyncRequest) SetPrune(v bool) {
	o.Prune = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetResources() []V1alpha1SyncOperationResource {
	if o == nil || isNil(o.Resources) {
		var ret []V1alpha1SyncOperationResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetResourcesOk() ([]V1alpha1SyncOperationResource, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []V1alpha1SyncOperationResource and assigns it to the Resources field.
func (o *ApplicationApplicationSyncRequest) SetResources(v []V1alpha1SyncOperationResource) {
	o.Resources = v
}

// GetRetryStrategy returns the RetryStrategy field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetRetryStrategy() V1alpha1RetryStrategy {
	if o == nil || isNil(o.RetryStrategy) {
		var ret V1alpha1RetryStrategy
		return ret
	}
	return *o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetRetryStrategyOk() (*V1alpha1RetryStrategy, bool) {
	if o == nil || isNil(o.RetryStrategy) {
		return nil, false
	}
	return o.RetryStrategy, true
}

// HasRetryStrategy returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasRetryStrategy() bool {
	if o != nil && !isNil(o.RetryStrategy) {
		return true
	}

	return false
}

// SetRetryStrategy gets a reference to the given V1alpha1RetryStrategy and assigns it to the RetryStrategy field.
func (o *ApplicationApplicationSyncRequest) SetRetryStrategy(v V1alpha1RetryStrategy) {
	o.RetryStrategy = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetRevision() string {
	if o == nil || isNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *ApplicationApplicationSyncRequest) SetRevision(v string) {
	o.Revision = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetStrategy() V1alpha1SyncStrategy {
	if o == nil || isNil(o.Strategy) {
		var ret V1alpha1SyncStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetStrategyOk() (*V1alpha1SyncStrategy, bool) {
	if o == nil || isNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasStrategy() bool {
	if o != nil && !isNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given V1alpha1SyncStrategy and assigns it to the Strategy field.
func (o *ApplicationApplicationSyncRequest) SetStrategy(v V1alpha1SyncStrategy) {
	o.Strategy = &v
}

// GetSyncOptions returns the SyncOptions field value if set, zero value otherwise.
func (o *ApplicationApplicationSyncRequest) GetSyncOptions() ApplicationSyncOptions {
	if o == nil || isNil(o.SyncOptions) {
		var ret ApplicationSyncOptions
		return ret
	}
	return *o.SyncOptions
}

// GetSyncOptionsOk returns a tuple with the SyncOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationSyncRequest) GetSyncOptionsOk() (*ApplicationSyncOptions, bool) {
	if o == nil || isNil(o.SyncOptions) {
		return nil, false
	}
	return o.SyncOptions, true
}

// HasSyncOptions returns a boolean if a field has been set.
func (o *ApplicationApplicationSyncRequest) HasSyncOptions() bool {
	if o != nil && !isNil(o.SyncOptions) {
		return true
	}

	return false
}

// SetSyncOptions gets a reference to the given ApplicationSyncOptions and assigns it to the SyncOptions field.
func (o *ApplicationApplicationSyncRequest) SetSyncOptions(v ApplicationSyncOptions) {
	o.SyncOptions = &v
}

func (o ApplicationApplicationSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationApplicationSyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppNamespace) {
		toSerialize["appNamespace"] = o.AppNamespace
	}
	if !isNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !isNil(o.Infos) {
		toSerialize["infos"] = o.Infos
	}
	if !isNil(o.Manifests) {
		toSerialize["manifests"] = o.Manifests
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Prune) {
		toSerialize["prune"] = o.Prune
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.RetryStrategy) {
		toSerialize["retryStrategy"] = o.RetryStrategy
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !isNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !isNil(o.SyncOptions) {
		toSerialize["syncOptions"] = o.SyncOptions
	}
	return toSerialize, nil
}

type NullableApplicationApplicationSyncRequest struct {
	value *ApplicationApplicationSyncRequest
	isSet bool
}

func (v NullableApplicationApplicationSyncRequest) Get() *ApplicationApplicationSyncRequest {
	return v.value
}

func (v *NullableApplicationApplicationSyncRequest) Set(val *ApplicationApplicationSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationApplicationSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationApplicationSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationApplicationSyncRequest(val *ApplicationApplicationSyncRequest) *NullableApplicationApplicationSyncRequest {
	return &NullableApplicationApplicationSyncRequest{value: val, isSet: true}
}

func (v NullableApplicationApplicationSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationApplicationSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
