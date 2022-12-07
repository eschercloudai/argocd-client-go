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

// checks if the V1alpha1ResourceActionParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ResourceActionParam{}

// V1alpha1ResourceActionParam struct for V1alpha1ResourceActionParam
type V1alpha1ResourceActionParam struct {
	Default *string `json:"default,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewV1alpha1ResourceActionParam instantiates a new V1alpha1ResourceActionParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceActionParam() *V1alpha1ResourceActionParam {
	this := V1alpha1ResourceActionParam{}
	return &this
}

// NewV1alpha1ResourceActionParamWithDefaults instantiates a new V1alpha1ResourceActionParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceActionParamWithDefaults() *V1alpha1ResourceActionParam {
	this := V1alpha1ResourceActionParam{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V1alpha1ResourceActionParam) GetDefault() string {
	if o == nil || isNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceActionParam) GetDefaultOk() (*string, bool) {
	if o == nil || isNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V1alpha1ResourceActionParam) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *V1alpha1ResourceActionParam) SetDefault(v string) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceActionParam) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceActionParam) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceActionParam) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceActionParam) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1ResourceActionParam) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceActionParam) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1ResourceActionParam) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1ResourceActionParam) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1alpha1ResourceActionParam) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceActionParam) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1alpha1ResourceActionParam) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1alpha1ResourceActionParam) SetValue(v string) {
	o.Value = &v
}

func (o V1alpha1ResourceActionParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ResourceActionParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableV1alpha1ResourceActionParam struct {
	value *V1alpha1ResourceActionParam
	isSet bool
}

func (v NullableV1alpha1ResourceActionParam) Get() *V1alpha1ResourceActionParam {
	return v.value
}

func (v *NullableV1alpha1ResourceActionParam) Set(val *V1alpha1ResourceActionParam) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceActionParam) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceActionParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceActionParam(val *V1alpha1ResourceActionParam) *NullableV1alpha1ResourceActionParam {
	return &NullableV1alpha1ResourceActionParam{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceActionParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceActionParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

