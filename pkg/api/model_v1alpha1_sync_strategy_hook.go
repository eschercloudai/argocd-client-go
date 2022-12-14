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

// checks if the V1alpha1SyncStrategyHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncStrategyHook{}

// V1alpha1SyncStrategyHook SyncStrategyHook will perform a sync using hooks annotations. If no hook annotation is specified falls back to `kubectl apply`.
type V1alpha1SyncStrategyHook struct {
	SyncStrategyApply *V1alpha1SyncStrategyApply `json:"syncStrategyApply,omitempty"`
}

// NewV1alpha1SyncStrategyHook instantiates a new V1alpha1SyncStrategyHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncStrategyHook() *V1alpha1SyncStrategyHook {
	this := V1alpha1SyncStrategyHook{}
	return &this
}

// NewV1alpha1SyncStrategyHookWithDefaults instantiates a new V1alpha1SyncStrategyHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncStrategyHookWithDefaults() *V1alpha1SyncStrategyHook {
	this := V1alpha1SyncStrategyHook{}
	return &this
}

// GetSyncStrategyApply returns the SyncStrategyApply field value if set, zero value otherwise.
func (o *V1alpha1SyncStrategyHook) GetSyncStrategyApply() V1alpha1SyncStrategyApply {
	if o == nil || isNil(o.SyncStrategyApply) {
		var ret V1alpha1SyncStrategyApply
		return ret
	}
	return *o.SyncStrategyApply
}

// GetSyncStrategyApplyOk returns a tuple with the SyncStrategyApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncStrategyHook) GetSyncStrategyApplyOk() (*V1alpha1SyncStrategyApply, bool) {
	if o == nil || isNil(o.SyncStrategyApply) {
		return nil, false
	}
	return o.SyncStrategyApply, true
}

// HasSyncStrategyApply returns a boolean if a field has been set.
func (o *V1alpha1SyncStrategyHook) HasSyncStrategyApply() bool {
	if o != nil && !isNil(o.SyncStrategyApply) {
		return true
	}

	return false
}

// SetSyncStrategyApply gets a reference to the given V1alpha1SyncStrategyApply and assigns it to the SyncStrategyApply field.
func (o *V1alpha1SyncStrategyHook) SetSyncStrategyApply(v V1alpha1SyncStrategyApply) {
	o.SyncStrategyApply = &v
}

func (o V1alpha1SyncStrategyHook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncStrategyHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SyncStrategyApply) {
		toSerialize["syncStrategyApply"] = o.SyncStrategyApply
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncStrategyHook struct {
	value *V1alpha1SyncStrategyHook
	isSet bool
}

func (v NullableV1alpha1SyncStrategyHook) Get() *V1alpha1SyncStrategyHook {
	return v.value
}

func (v *NullableV1alpha1SyncStrategyHook) Set(val *V1alpha1SyncStrategyHook) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncStrategyHook) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncStrategyHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncStrategyHook(val *V1alpha1SyncStrategyHook) *NullableV1alpha1SyncStrategyHook {
	return &NullableV1alpha1SyncStrategyHook{value: val, isSet: true}
}

func (v NullableV1alpha1SyncStrategyHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncStrategyHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
