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

// checks if the V1alpha1ApplicationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationStatus{}

// V1alpha1ApplicationStatus struct for V1alpha1ApplicationStatus
type V1alpha1ApplicationStatus struct {
	Conditions []V1alpha1ApplicationCondition `json:"conditions,omitempty"`
	Health *V1alpha1HealthStatus `json:"health,omitempty"`
	History []V1alpha1RevisionHistory `json:"history,omitempty"`
	ObservedAt *V1Time `json:"observedAt,omitempty"`
	OperationState *V1alpha1OperationState `json:"operationState,omitempty"`
	ReconciledAt *V1Time `json:"reconciledAt,omitempty"`
	ResourceHealthSource *string `json:"resourceHealthSource,omitempty"`
	Resources []V1alpha1ResourceStatus `json:"resources,omitempty"`
	SourceType *string `json:"sourceType,omitempty"`
	Summary *V1alpha1ApplicationSummary `json:"summary,omitempty"`
	Sync *V1alpha1SyncStatus `json:"sync,omitempty"`
}

// NewV1alpha1ApplicationStatus instantiates a new V1alpha1ApplicationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationStatus() *V1alpha1ApplicationStatus {
	this := V1alpha1ApplicationStatus{}
	return &this
}

// NewV1alpha1ApplicationStatusWithDefaults instantiates a new V1alpha1ApplicationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationStatusWithDefaults() *V1alpha1ApplicationStatus {
	this := V1alpha1ApplicationStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetConditions() []V1alpha1ApplicationCondition {
	if o == nil || isNil(o.Conditions) {
		var ret []V1alpha1ApplicationCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetConditionsOk() ([]V1alpha1ApplicationCondition, bool) {
	if o == nil || isNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasConditions() bool {
	if o != nil && !isNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []V1alpha1ApplicationCondition and assigns it to the Conditions field.
func (o *V1alpha1ApplicationStatus) SetConditions(v []V1alpha1ApplicationCondition) {
	o.Conditions = v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetHealth() V1alpha1HealthStatus {
	if o == nil || isNil(o.Health) {
		var ret V1alpha1HealthStatus
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetHealthOk() (*V1alpha1HealthStatus, bool) {
	if o == nil || isNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasHealth() bool {
	if o != nil && !isNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given V1alpha1HealthStatus and assigns it to the Health field.
func (o *V1alpha1ApplicationStatus) SetHealth(v V1alpha1HealthStatus) {
	o.Health = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetHistory() []V1alpha1RevisionHistory {
	if o == nil || isNil(o.History) {
		var ret []V1alpha1RevisionHistory
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetHistoryOk() ([]V1alpha1RevisionHistory, bool) {
	if o == nil || isNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasHistory() bool {
	if o != nil && !isNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []V1alpha1RevisionHistory and assigns it to the History field.
func (o *V1alpha1ApplicationStatus) SetHistory(v []V1alpha1RevisionHistory) {
	o.History = v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetObservedAt() V1Time {
	if o == nil || isNil(o.ObservedAt) {
		var ret V1Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetObservedAtOk() (*V1Time, bool) {
	if o == nil || isNil(o.ObservedAt) {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasObservedAt() bool {
	if o != nil && !isNil(o.ObservedAt) {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given V1Time and assigns it to the ObservedAt field.
func (o *V1alpha1ApplicationStatus) SetObservedAt(v V1Time) {
	o.ObservedAt = &v
}

// GetOperationState returns the OperationState field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetOperationState() V1alpha1OperationState {
	if o == nil || isNil(o.OperationState) {
		var ret V1alpha1OperationState
		return ret
	}
	return *o.OperationState
}

// GetOperationStateOk returns a tuple with the OperationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetOperationStateOk() (*V1alpha1OperationState, bool) {
	if o == nil || isNil(o.OperationState) {
		return nil, false
	}
	return o.OperationState, true
}

// HasOperationState returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasOperationState() bool {
	if o != nil && !isNil(o.OperationState) {
		return true
	}

	return false
}

// SetOperationState gets a reference to the given V1alpha1OperationState and assigns it to the OperationState field.
func (o *V1alpha1ApplicationStatus) SetOperationState(v V1alpha1OperationState) {
	o.OperationState = &v
}

// GetReconciledAt returns the ReconciledAt field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetReconciledAt() V1Time {
	if o == nil || isNil(o.ReconciledAt) {
		var ret V1Time
		return ret
	}
	return *o.ReconciledAt
}

// GetReconciledAtOk returns a tuple with the ReconciledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetReconciledAtOk() (*V1Time, bool) {
	if o == nil || isNil(o.ReconciledAt) {
		return nil, false
	}
	return o.ReconciledAt, true
}

// HasReconciledAt returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasReconciledAt() bool {
	if o != nil && !isNil(o.ReconciledAt) {
		return true
	}

	return false
}

// SetReconciledAt gets a reference to the given V1Time and assigns it to the ReconciledAt field.
func (o *V1alpha1ApplicationStatus) SetReconciledAt(v V1Time) {
	o.ReconciledAt = &v
}

// GetResourceHealthSource returns the ResourceHealthSource field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetResourceHealthSource() string {
	if o == nil || isNil(o.ResourceHealthSource) {
		var ret string
		return ret
	}
	return *o.ResourceHealthSource
}

// GetResourceHealthSourceOk returns a tuple with the ResourceHealthSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetResourceHealthSourceOk() (*string, bool) {
	if o == nil || isNil(o.ResourceHealthSource) {
		return nil, false
	}
	return o.ResourceHealthSource, true
}

// HasResourceHealthSource returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasResourceHealthSource() bool {
	if o != nil && !isNil(o.ResourceHealthSource) {
		return true
	}

	return false
}

// SetResourceHealthSource gets a reference to the given string and assigns it to the ResourceHealthSource field.
func (o *V1alpha1ApplicationStatus) SetResourceHealthSource(v string) {
	o.ResourceHealthSource = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetResources() []V1alpha1ResourceStatus {
	if o == nil || isNil(o.Resources) {
		var ret []V1alpha1ResourceStatus
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetResourcesOk() ([]V1alpha1ResourceStatus, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []V1alpha1ResourceStatus and assigns it to the Resources field.
func (o *V1alpha1ApplicationStatus) SetResources(v []V1alpha1ResourceStatus) {
	o.Resources = v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetSourceType() string {
	if o == nil || isNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetSourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasSourceType() bool {
	if o != nil && !isNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *V1alpha1ApplicationStatus) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetSummary() V1alpha1ApplicationSummary {
	if o == nil || isNil(o.Summary) {
		var ret V1alpha1ApplicationSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetSummaryOk() (*V1alpha1ApplicationSummary, bool) {
	if o == nil || isNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasSummary() bool {
	if o != nil && !isNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given V1alpha1ApplicationSummary and assigns it to the Summary field.
func (o *V1alpha1ApplicationStatus) SetSummary(v V1alpha1ApplicationSummary) {
	o.Summary = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *V1alpha1ApplicationStatus) GetSync() V1alpha1SyncStatus {
	if o == nil || isNil(o.Sync) {
		var ret V1alpha1SyncStatus
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationStatus) GetSyncOk() (*V1alpha1SyncStatus, bool) {
	if o == nil || isNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *V1alpha1ApplicationStatus) HasSync() bool {
	if o != nil && !isNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given V1alpha1SyncStatus and assigns it to the Sync field.
func (o *V1alpha1ApplicationStatus) SetSync(v V1alpha1SyncStatus) {
	o.Sync = &v
}

func (o V1alpha1ApplicationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !isNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !isNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !isNil(o.ObservedAt) {
		toSerialize["observedAt"] = o.ObservedAt
	}
	if !isNil(o.OperationState) {
		toSerialize["operationState"] = o.OperationState
	}
	if !isNil(o.ReconciledAt) {
		toSerialize["reconciledAt"] = o.ReconciledAt
	}
	if !isNil(o.ResourceHealthSource) {
		toSerialize["resourceHealthSource"] = o.ResourceHealthSource
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !isNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !isNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationStatus struct {
	value *V1alpha1ApplicationStatus
	isSet bool
}

func (v NullableV1alpha1ApplicationStatus) Get() *V1alpha1ApplicationStatus {
	return v.value
}

func (v *NullableV1alpha1ApplicationStatus) Set(val *V1alpha1ApplicationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationStatus(val *V1alpha1ApplicationStatus) *NullableV1alpha1ApplicationStatus {
	return &NullableV1alpha1ApplicationStatus{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


