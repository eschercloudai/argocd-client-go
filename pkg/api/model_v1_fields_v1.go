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

// checks if the V1FieldsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1FieldsV1{}

// V1FieldsV1 FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.  Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.  The exact format is defined in sigs.k8s.io/structured-merge-diff +protobuf.options.(gogoproto.goproto_stringer)=false
type V1FieldsV1 struct {
	// Raw is the underlying serialization of this object.
	Raw *string `json:"Raw,omitempty"`
}

// NewV1FieldsV1 instantiates a new V1FieldsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1FieldsV1() *V1FieldsV1 {
	this := V1FieldsV1{}
	return &this
}

// NewV1FieldsV1WithDefaults instantiates a new V1FieldsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1FieldsV1WithDefaults() *V1FieldsV1 {
	this := V1FieldsV1{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *V1FieldsV1) GetRaw() string {
	if o == nil || isNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FieldsV1) GetRawOk() (*string, bool) {
	if o == nil || isNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *V1FieldsV1) HasRaw() bool {
	if o != nil && !isNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *V1FieldsV1) SetRaw(v string) {
	o.Raw = &v
}

func (o V1FieldsV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1FieldsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Raw) {
		toSerialize["Raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableV1FieldsV1 struct {
	value *V1FieldsV1
	isSet bool
}

func (v NullableV1FieldsV1) Get() *V1FieldsV1 {
	return v.value
}

func (v *NullableV1FieldsV1) Set(val *V1FieldsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableV1FieldsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableV1FieldsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1FieldsV1(val *V1FieldsV1) *NullableV1FieldsV1 {
	return &NullableV1FieldsV1{value: val, isSet: true}
}

func (v NullableV1FieldsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1FieldsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
