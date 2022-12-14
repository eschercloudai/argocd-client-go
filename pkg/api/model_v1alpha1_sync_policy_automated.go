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

// checks if the V1alpha1SyncPolicyAutomated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncPolicyAutomated{}

// V1alpha1SyncPolicyAutomated struct for V1alpha1SyncPolicyAutomated
type V1alpha1SyncPolicyAutomated struct {
	AllowEmpty *bool `json:"allowEmpty,omitempty"`
	Prune      *bool `json:"prune,omitempty"`
	SelfHeal   *bool `json:"selfHeal,omitempty"`
}

// NewV1alpha1SyncPolicyAutomated instantiates a new V1alpha1SyncPolicyAutomated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncPolicyAutomated() *V1alpha1SyncPolicyAutomated {
	this := V1alpha1SyncPolicyAutomated{}
	return &this
}

// NewV1alpha1SyncPolicyAutomatedWithDefaults instantiates a new V1alpha1SyncPolicyAutomated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncPolicyAutomatedWithDefaults() *V1alpha1SyncPolicyAutomated {
	this := V1alpha1SyncPolicyAutomated{}
	return &this
}

// GetAllowEmpty returns the AllowEmpty field value if set, zero value otherwise.
func (o *V1alpha1SyncPolicyAutomated) GetAllowEmpty() bool {
	if o == nil || isNil(o.AllowEmpty) {
		var ret bool
		return ret
	}
	return *o.AllowEmpty
}

// GetAllowEmptyOk returns a tuple with the AllowEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncPolicyAutomated) GetAllowEmptyOk() (*bool, bool) {
	if o == nil || isNil(o.AllowEmpty) {
		return nil, false
	}
	return o.AllowEmpty, true
}

// HasAllowEmpty returns a boolean if a field has been set.
func (o *V1alpha1SyncPolicyAutomated) HasAllowEmpty() bool {
	if o != nil && !isNil(o.AllowEmpty) {
		return true
	}

	return false
}

// SetAllowEmpty gets a reference to the given bool and assigns it to the AllowEmpty field.
func (o *V1alpha1SyncPolicyAutomated) SetAllowEmpty(v bool) {
	o.AllowEmpty = &v
}

// GetPrune returns the Prune field value if set, zero value otherwise.
func (o *V1alpha1SyncPolicyAutomated) GetPrune() bool {
	if o == nil || isNil(o.Prune) {
		var ret bool
		return ret
	}
	return *o.Prune
}

// GetPruneOk returns a tuple with the Prune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncPolicyAutomated) GetPruneOk() (*bool, bool) {
	if o == nil || isNil(o.Prune) {
		return nil, false
	}
	return o.Prune, true
}

// HasPrune returns a boolean if a field has been set.
func (o *V1alpha1SyncPolicyAutomated) HasPrune() bool {
	if o != nil && !isNil(o.Prune) {
		return true
	}

	return false
}

// SetPrune gets a reference to the given bool and assigns it to the Prune field.
func (o *V1alpha1SyncPolicyAutomated) SetPrune(v bool) {
	o.Prune = &v
}

// GetSelfHeal returns the SelfHeal field value if set, zero value otherwise.
func (o *V1alpha1SyncPolicyAutomated) GetSelfHeal() bool {
	if o == nil || isNil(o.SelfHeal) {
		var ret bool
		return ret
	}
	return *o.SelfHeal
}

// GetSelfHealOk returns a tuple with the SelfHeal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncPolicyAutomated) GetSelfHealOk() (*bool, bool) {
	if o == nil || isNil(o.SelfHeal) {
		return nil, false
	}
	return o.SelfHeal, true
}

// HasSelfHeal returns a boolean if a field has been set.
func (o *V1alpha1SyncPolicyAutomated) HasSelfHeal() bool {
	if o != nil && !isNil(o.SelfHeal) {
		return true
	}

	return false
}

// SetSelfHeal gets a reference to the given bool and assigns it to the SelfHeal field.
func (o *V1alpha1SyncPolicyAutomated) SetSelfHeal(v bool) {
	o.SelfHeal = &v
}

func (o V1alpha1SyncPolicyAutomated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncPolicyAutomated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowEmpty) {
		toSerialize["allowEmpty"] = o.AllowEmpty
	}
	if !isNil(o.Prune) {
		toSerialize["prune"] = o.Prune
	}
	if !isNil(o.SelfHeal) {
		toSerialize["selfHeal"] = o.SelfHeal
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncPolicyAutomated struct {
	value *V1alpha1SyncPolicyAutomated
	isSet bool
}

func (v NullableV1alpha1SyncPolicyAutomated) Get() *V1alpha1SyncPolicyAutomated {
	return v.value
}

func (v *NullableV1alpha1SyncPolicyAutomated) Set(val *V1alpha1SyncPolicyAutomated) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncPolicyAutomated) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncPolicyAutomated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncPolicyAutomated(val *V1alpha1SyncPolicyAutomated) *NullableV1alpha1SyncPolicyAutomated {
	return &NullableV1alpha1SyncPolicyAutomated{value: val, isSet: true}
}

func (v NullableV1alpha1SyncPolicyAutomated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncPolicyAutomated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
