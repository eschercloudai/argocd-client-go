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

// checks if the V1alpha1ApplicationSourcePlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationSourcePlugin{}

// V1alpha1ApplicationSourcePlugin struct for V1alpha1ApplicationSourcePlugin
type V1alpha1ApplicationSourcePlugin struct {
	Env  []Applicationv1alpha1EnvEntry `json:"env,omitempty"`
	Name *string                       `json:"name,omitempty"`
}

// NewV1alpha1ApplicationSourcePlugin instantiates a new V1alpha1ApplicationSourcePlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSourcePlugin() *V1alpha1ApplicationSourcePlugin {
	this := V1alpha1ApplicationSourcePlugin{}
	return &this
}

// NewV1alpha1ApplicationSourcePluginWithDefaults instantiates a new V1alpha1ApplicationSourcePlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSourcePluginWithDefaults() *V1alpha1ApplicationSourcePlugin {
	this := V1alpha1ApplicationSourcePlugin{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSourcePlugin) GetEnv() []Applicationv1alpha1EnvEntry {
	if o == nil || isNil(o.Env) {
		var ret []Applicationv1alpha1EnvEntry
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSourcePlugin) GetEnvOk() ([]Applicationv1alpha1EnvEntry, bool) {
	if o == nil || isNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSourcePlugin) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Applicationv1alpha1EnvEntry and assigns it to the Env field.
func (o *V1alpha1ApplicationSourcePlugin) SetEnv(v []Applicationv1alpha1EnvEntry) {
	o.Env = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSourcePlugin) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSourcePlugin) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSourcePlugin) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ApplicationSourcePlugin) SetName(v string) {
	o.Name = &v
}

func (o V1alpha1ApplicationSourcePlugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationSourcePlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationSourcePlugin struct {
	value *V1alpha1ApplicationSourcePlugin
	isSet bool
}

func (v NullableV1alpha1ApplicationSourcePlugin) Get() *V1alpha1ApplicationSourcePlugin {
	return v.value
}

func (v *NullableV1alpha1ApplicationSourcePlugin) Set(val *V1alpha1ApplicationSourcePlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSourcePlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSourcePlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSourcePlugin(val *V1alpha1ApplicationSourcePlugin) *NullableV1alpha1ApplicationSourcePlugin {
	return &NullableV1alpha1ApplicationSourcePlugin{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSourcePlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSourcePlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
