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

// checks if the Applicationv1alpha1EnvEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Applicationv1alpha1EnvEntry{}

// Applicationv1alpha1EnvEntry struct for Applicationv1alpha1EnvEntry
type Applicationv1alpha1EnvEntry struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApplicationv1alpha1EnvEntry instantiates a new Applicationv1alpha1EnvEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationv1alpha1EnvEntry() *Applicationv1alpha1EnvEntry {
	this := Applicationv1alpha1EnvEntry{}
	return &this
}

// NewApplicationv1alpha1EnvEntryWithDefaults instantiates a new Applicationv1alpha1EnvEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationv1alpha1EnvEntryWithDefaults() *Applicationv1alpha1EnvEntry {
	this := Applicationv1alpha1EnvEntry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Applicationv1alpha1EnvEntry) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationv1alpha1EnvEntry) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Applicationv1alpha1EnvEntry) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Applicationv1alpha1EnvEntry) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Applicationv1alpha1EnvEntry) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationv1alpha1EnvEntry) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Applicationv1alpha1EnvEntry) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Applicationv1alpha1EnvEntry) SetValue(v string) {
	o.Value = &v
}

func (o Applicationv1alpha1EnvEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Applicationv1alpha1EnvEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApplicationv1alpha1EnvEntry struct {
	value *Applicationv1alpha1EnvEntry
	isSet bool
}

func (v NullableApplicationv1alpha1EnvEntry) Get() *Applicationv1alpha1EnvEntry {
	return v.value
}

func (v *NullableApplicationv1alpha1EnvEntry) Set(val *Applicationv1alpha1EnvEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationv1alpha1EnvEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationv1alpha1EnvEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationv1alpha1EnvEntry(val *Applicationv1alpha1EnvEntry) *NullableApplicationv1alpha1EnvEntry {
	return &NullableApplicationv1alpha1EnvEntry{value: val, isSet: true}
}

func (v NullableApplicationv1alpha1EnvEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationv1alpha1EnvEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
