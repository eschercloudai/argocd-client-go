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

// checks if the V1alpha1SyncWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncWindow{}

// V1alpha1SyncWindow struct for V1alpha1SyncWindow
type V1alpha1SyncWindow struct {
	Applications []string `json:"applications,omitempty"`
	Clusters     []string `json:"clusters,omitempty"`
	Duration     *string  `json:"duration,omitempty"`
	Kind         *string  `json:"kind,omitempty"`
	ManualSync   *bool    `json:"manualSync,omitempty"`
	Namespaces   []string `json:"namespaces,omitempty"`
	Schedule     *string  `json:"schedule,omitempty"`
	TimeZone     *string  `json:"timeZone,omitempty"`
}

// NewV1alpha1SyncWindow instantiates a new V1alpha1SyncWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncWindow() *V1alpha1SyncWindow {
	this := V1alpha1SyncWindow{}
	return &this
}

// NewV1alpha1SyncWindowWithDefaults instantiates a new V1alpha1SyncWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncWindowWithDefaults() *V1alpha1SyncWindow {
	this := V1alpha1SyncWindow{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetApplications() []string {
	if o == nil || isNil(o.Applications) {
		var ret []string
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetApplicationsOk() ([]string, bool) {
	if o == nil || isNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasApplications() bool {
	if o != nil && !isNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []string and assigns it to the Applications field.
func (o *V1alpha1SyncWindow) SetApplications(v []string) {
	o.Applications = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetClusters() []string {
	if o == nil || isNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetClustersOk() ([]string, bool) {
	if o == nil || isNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasClusters() bool {
	if o != nil && !isNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *V1alpha1SyncWindow) SetClusters(v []string) {
	o.Clusters = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetDuration() string {
	if o == nil || isNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetDurationOk() (*string, bool) {
	if o == nil || isNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *V1alpha1SyncWindow) SetDuration(v string) {
	o.Duration = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1SyncWindow) SetKind(v string) {
	o.Kind = &v
}

// GetManualSync returns the ManualSync field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetManualSync() bool {
	if o == nil || isNil(o.ManualSync) {
		var ret bool
		return ret
	}
	return *o.ManualSync
}

// GetManualSyncOk returns a tuple with the ManualSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetManualSyncOk() (*bool, bool) {
	if o == nil || isNil(o.ManualSync) {
		return nil, false
	}
	return o.ManualSync, true
}

// HasManualSync returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasManualSync() bool {
	if o != nil && !isNil(o.ManualSync) {
		return true
	}

	return false
}

// SetManualSync gets a reference to the given bool and assigns it to the ManualSync field.
func (o *V1alpha1SyncWindow) SetManualSync(v bool) {
	o.ManualSync = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetNamespaces() []string {
	if o == nil || isNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetNamespacesOk() ([]string, bool) {
	if o == nil || isNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasNamespaces() bool {
	if o != nil && !isNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *V1alpha1SyncWindow) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetSchedule() string {
	if o == nil || isNil(o.Schedule) {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetScheduleOk() (*string, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *V1alpha1SyncWindow) SetSchedule(v string) {
	o.Schedule = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *V1alpha1SyncWindow) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncWindow) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *V1alpha1SyncWindow) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *V1alpha1SyncWindow) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o V1alpha1SyncWindow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !isNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.ManualSync) {
		toSerialize["manualSync"] = o.ManualSync
	}
	if !isNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncWindow struct {
	value *V1alpha1SyncWindow
	isSet bool
}

func (v NullableV1alpha1SyncWindow) Get() *V1alpha1SyncWindow {
	return v.value
}

func (v *NullableV1alpha1SyncWindow) Set(val *V1alpha1SyncWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncWindow(val *V1alpha1SyncWindow) *NullableV1alpha1SyncWindow {
	return &NullableV1alpha1SyncWindow{value: val, isSet: true}
}

func (v NullableV1alpha1SyncWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
