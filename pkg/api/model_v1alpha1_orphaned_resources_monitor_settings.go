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

// checks if the V1alpha1OrphanedResourcesMonitorSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1OrphanedResourcesMonitorSettings{}

// V1alpha1OrphanedResourcesMonitorSettings struct for V1alpha1OrphanedResourcesMonitorSettings
type V1alpha1OrphanedResourcesMonitorSettings struct {
	Ignore []V1alpha1OrphanedResourceKey `json:"ignore,omitempty"`
	Warn   *bool                         `json:"warn,omitempty"`
}

// NewV1alpha1OrphanedResourcesMonitorSettings instantiates a new V1alpha1OrphanedResourcesMonitorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1OrphanedResourcesMonitorSettings() *V1alpha1OrphanedResourcesMonitorSettings {
	this := V1alpha1OrphanedResourcesMonitorSettings{}
	return &this
}

// NewV1alpha1OrphanedResourcesMonitorSettingsWithDefaults instantiates a new V1alpha1OrphanedResourcesMonitorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1OrphanedResourcesMonitorSettingsWithDefaults() *V1alpha1OrphanedResourcesMonitorSettings {
	this := V1alpha1OrphanedResourcesMonitorSettings{}
	return &this
}

// GetIgnore returns the Ignore field value if set, zero value otherwise.
func (o *V1alpha1OrphanedResourcesMonitorSettings) GetIgnore() []V1alpha1OrphanedResourceKey {
	if o == nil || isNil(o.Ignore) {
		var ret []V1alpha1OrphanedResourceKey
		return ret
	}
	return o.Ignore
}

// GetIgnoreOk returns a tuple with the Ignore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1OrphanedResourcesMonitorSettings) GetIgnoreOk() ([]V1alpha1OrphanedResourceKey, bool) {
	if o == nil || isNil(o.Ignore) {
		return nil, false
	}
	return o.Ignore, true
}

// HasIgnore returns a boolean if a field has been set.
func (o *V1alpha1OrphanedResourcesMonitorSettings) HasIgnore() bool {
	if o != nil && !isNil(o.Ignore) {
		return true
	}

	return false
}

// SetIgnore gets a reference to the given []V1alpha1OrphanedResourceKey and assigns it to the Ignore field.
func (o *V1alpha1OrphanedResourcesMonitorSettings) SetIgnore(v []V1alpha1OrphanedResourceKey) {
	o.Ignore = v
}

// GetWarn returns the Warn field value if set, zero value otherwise.
func (o *V1alpha1OrphanedResourcesMonitorSettings) GetWarn() bool {
	if o == nil || isNil(o.Warn) {
		var ret bool
		return ret
	}
	return *o.Warn
}

// GetWarnOk returns a tuple with the Warn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1OrphanedResourcesMonitorSettings) GetWarnOk() (*bool, bool) {
	if o == nil || isNil(o.Warn) {
		return nil, false
	}
	return o.Warn, true
}

// HasWarn returns a boolean if a field has been set.
func (o *V1alpha1OrphanedResourcesMonitorSettings) HasWarn() bool {
	if o != nil && !isNil(o.Warn) {
		return true
	}

	return false
}

// SetWarn gets a reference to the given bool and assigns it to the Warn field.
func (o *V1alpha1OrphanedResourcesMonitorSettings) SetWarn(v bool) {
	o.Warn = &v
}

func (o V1alpha1OrphanedResourcesMonitorSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1OrphanedResourcesMonitorSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ignore) {
		toSerialize["ignore"] = o.Ignore
	}
	if !isNil(o.Warn) {
		toSerialize["warn"] = o.Warn
	}
	return toSerialize, nil
}

type NullableV1alpha1OrphanedResourcesMonitorSettings struct {
	value *V1alpha1OrphanedResourcesMonitorSettings
	isSet bool
}

func (v NullableV1alpha1OrphanedResourcesMonitorSettings) Get() *V1alpha1OrphanedResourcesMonitorSettings {
	return v.value
}

func (v *NullableV1alpha1OrphanedResourcesMonitorSettings) Set(val *V1alpha1OrphanedResourcesMonitorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1OrphanedResourcesMonitorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1OrphanedResourcesMonitorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1OrphanedResourcesMonitorSettings(val *V1alpha1OrphanedResourcesMonitorSettings) *NullableV1alpha1OrphanedResourcesMonitorSettings {
	return &NullableV1alpha1OrphanedResourcesMonitorSettings{value: val, isSet: true}
}

func (v NullableV1alpha1OrphanedResourcesMonitorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1OrphanedResourcesMonitorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}