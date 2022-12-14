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

// checks if the V1alpha1SyncStrategyApply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncStrategyApply{}

// V1alpha1SyncStrategyApply struct for V1alpha1SyncStrategyApply
type V1alpha1SyncStrategyApply struct {
	// Force indicates whether or not to supply the --force flag to `kubectl apply`. The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.
	Force *bool `json:"force,omitempty"`
}

// NewV1alpha1SyncStrategyApply instantiates a new V1alpha1SyncStrategyApply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncStrategyApply() *V1alpha1SyncStrategyApply {
	this := V1alpha1SyncStrategyApply{}
	return &this
}

// NewV1alpha1SyncStrategyApplyWithDefaults instantiates a new V1alpha1SyncStrategyApply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncStrategyApplyWithDefaults() *V1alpha1SyncStrategyApply {
	this := V1alpha1SyncStrategyApply{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *V1alpha1SyncStrategyApply) GetForce() bool {
	if o == nil || isNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncStrategyApply) GetForceOk() (*bool, bool) {
	if o == nil || isNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *V1alpha1SyncStrategyApply) HasForce() bool {
	if o != nil && !isNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *V1alpha1SyncStrategyApply) SetForce(v bool) {
	o.Force = &v
}

func (o V1alpha1SyncStrategyApply) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncStrategyApply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncStrategyApply struct {
	value *V1alpha1SyncStrategyApply
	isSet bool
}

func (v NullableV1alpha1SyncStrategyApply) Get() *V1alpha1SyncStrategyApply {
	return v.value
}

func (v *NullableV1alpha1SyncStrategyApply) Set(val *V1alpha1SyncStrategyApply) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncStrategyApply) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncStrategyApply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncStrategyApply(val *V1alpha1SyncStrategyApply) *NullableV1alpha1SyncStrategyApply {
	return &NullableV1alpha1SyncStrategyApply{value: val, isSet: true}
}

func (v NullableV1alpha1SyncStrategyApply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncStrategyApply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
