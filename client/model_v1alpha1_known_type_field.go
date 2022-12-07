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

// checks if the V1alpha1KnownTypeField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1KnownTypeField{}

// V1alpha1KnownTypeField struct for V1alpha1KnownTypeField
type V1alpha1KnownTypeField struct {
	Field *string `json:"field,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV1alpha1KnownTypeField instantiates a new V1alpha1KnownTypeField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1KnownTypeField() *V1alpha1KnownTypeField {
	this := V1alpha1KnownTypeField{}
	return &this
}

// NewV1alpha1KnownTypeFieldWithDefaults instantiates a new V1alpha1KnownTypeField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1KnownTypeFieldWithDefaults() *V1alpha1KnownTypeField {
	this := V1alpha1KnownTypeField{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *V1alpha1KnownTypeField) GetField() string {
	if o == nil || isNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1KnownTypeField) GetFieldOk() (*string, bool) {
	if o == nil || isNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *V1alpha1KnownTypeField) HasField() bool {
	if o != nil && !isNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *V1alpha1KnownTypeField) SetField(v string) {
	o.Field = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1KnownTypeField) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1KnownTypeField) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1KnownTypeField) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1KnownTypeField) SetType(v string) {
	o.Type = &v
}

func (o V1alpha1KnownTypeField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1KnownTypeField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1alpha1KnownTypeField struct {
	value *V1alpha1KnownTypeField
	isSet bool
}

func (v NullableV1alpha1KnownTypeField) Get() *V1alpha1KnownTypeField {
	return v.value
}

func (v *NullableV1alpha1KnownTypeField) Set(val *V1alpha1KnownTypeField) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1KnownTypeField) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1KnownTypeField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1KnownTypeField(val *V1alpha1KnownTypeField) *NullableV1alpha1KnownTypeField {
	return &NullableV1alpha1KnownTypeField{value: val, isSet: true}
}

func (v NullableV1alpha1KnownTypeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1KnownTypeField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

