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

// checks if the V1alpha1SyncStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncStrategy{}

// V1alpha1SyncStrategy struct for V1alpha1SyncStrategy
type V1alpha1SyncStrategy struct {
	Apply *V1alpha1SyncStrategyApply `json:"apply,omitempty"`
	Hook *V1alpha1SyncStrategyHook `json:"hook,omitempty"`
}

// NewV1alpha1SyncStrategy instantiates a new V1alpha1SyncStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncStrategy() *V1alpha1SyncStrategy {
	this := V1alpha1SyncStrategy{}
	return &this
}

// NewV1alpha1SyncStrategyWithDefaults instantiates a new V1alpha1SyncStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncStrategyWithDefaults() *V1alpha1SyncStrategy {
	this := V1alpha1SyncStrategy{}
	return &this
}

// GetApply returns the Apply field value if set, zero value otherwise.
func (o *V1alpha1SyncStrategy) GetApply() V1alpha1SyncStrategyApply {
	if o == nil || isNil(o.Apply) {
		var ret V1alpha1SyncStrategyApply
		return ret
	}
	return *o.Apply
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncStrategy) GetApplyOk() (*V1alpha1SyncStrategyApply, bool) {
	if o == nil || isNil(o.Apply) {
		return nil, false
	}
	return o.Apply, true
}

// HasApply returns a boolean if a field has been set.
func (o *V1alpha1SyncStrategy) HasApply() bool {
	if o != nil && !isNil(o.Apply) {
		return true
	}

	return false
}

// SetApply gets a reference to the given V1alpha1SyncStrategyApply and assigns it to the Apply field.
func (o *V1alpha1SyncStrategy) SetApply(v V1alpha1SyncStrategyApply) {
	o.Apply = &v
}

// GetHook returns the Hook field value if set, zero value otherwise.
func (o *V1alpha1SyncStrategy) GetHook() V1alpha1SyncStrategyHook {
	if o == nil || isNil(o.Hook) {
		var ret V1alpha1SyncStrategyHook
		return ret
	}
	return *o.Hook
}

// GetHookOk returns a tuple with the Hook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncStrategy) GetHookOk() (*V1alpha1SyncStrategyHook, bool) {
	if o == nil || isNil(o.Hook) {
		return nil, false
	}
	return o.Hook, true
}

// HasHook returns a boolean if a field has been set.
func (o *V1alpha1SyncStrategy) HasHook() bool {
	if o != nil && !isNil(o.Hook) {
		return true
	}

	return false
}

// SetHook gets a reference to the given V1alpha1SyncStrategyHook and assigns it to the Hook field.
func (o *V1alpha1SyncStrategy) SetHook(v V1alpha1SyncStrategyHook) {
	o.Hook = &v
}

func (o V1alpha1SyncStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Apply) {
		toSerialize["apply"] = o.Apply
	}
	if !isNil(o.Hook) {
		toSerialize["hook"] = o.Hook
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncStrategy struct {
	value *V1alpha1SyncStrategy
	isSet bool
}

func (v NullableV1alpha1SyncStrategy) Get() *V1alpha1SyncStrategy {
	return v.value
}

func (v *NullableV1alpha1SyncStrategy) Set(val *V1alpha1SyncStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncStrategy(val *V1alpha1SyncStrategy) *NullableV1alpha1SyncStrategy {
	return &NullableV1alpha1SyncStrategy{value: val, isSet: true}
}

func (v NullableV1alpha1SyncStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


