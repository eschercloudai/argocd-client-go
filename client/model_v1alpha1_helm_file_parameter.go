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

// checks if the V1alpha1HelmFileParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1HelmFileParameter{}

// V1alpha1HelmFileParameter struct for V1alpha1HelmFileParameter
type V1alpha1HelmFileParameter struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewV1alpha1HelmFileParameter instantiates a new V1alpha1HelmFileParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1HelmFileParameter() *V1alpha1HelmFileParameter {
	this := V1alpha1HelmFileParameter{}
	return &this
}

// NewV1alpha1HelmFileParameterWithDefaults instantiates a new V1alpha1HelmFileParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1HelmFileParameterWithDefaults() *V1alpha1HelmFileParameter {
	this := V1alpha1HelmFileParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1HelmFileParameter) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HelmFileParameter) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1HelmFileParameter) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1HelmFileParameter) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *V1alpha1HelmFileParameter) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1HelmFileParameter) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *V1alpha1HelmFileParameter) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *V1alpha1HelmFileParameter) SetPath(v string) {
	o.Path = &v
}

func (o V1alpha1HelmFileParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1HelmFileParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableV1alpha1HelmFileParameter struct {
	value *V1alpha1HelmFileParameter
	isSet bool
}

func (v NullableV1alpha1HelmFileParameter) Get() *V1alpha1HelmFileParameter {
	return v.value
}

func (v *NullableV1alpha1HelmFileParameter) Set(val *V1alpha1HelmFileParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1HelmFileParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1HelmFileParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1HelmFileParameter(val *V1alpha1HelmFileParameter) *NullableV1alpha1HelmFileParameter {
	return &NullableV1alpha1HelmFileParameter{value: val, isSet: true}
}

func (v NullableV1alpha1HelmFileParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1HelmFileParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


