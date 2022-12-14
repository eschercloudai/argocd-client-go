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

// checks if the V1alpha1ClusterGenerator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ClusterGenerator{}

// V1alpha1ClusterGenerator ClusterGenerator defines a generator to match against clusters registered with ArgoCD.
type V1alpha1ClusterGenerator struct {
	Selector *V1LabelSelector                `json:"selector,omitempty"`
	Template *V1alpha1ApplicationSetTemplate `json:"template,omitempty"`
	Values   *map[string]string              `json:"values,omitempty"`
}

// NewV1alpha1ClusterGenerator instantiates a new V1alpha1ClusterGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ClusterGenerator() *V1alpha1ClusterGenerator {
	this := V1alpha1ClusterGenerator{}
	return &this
}

// NewV1alpha1ClusterGeneratorWithDefaults instantiates a new V1alpha1ClusterGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ClusterGeneratorWithDefaults() *V1alpha1ClusterGenerator {
	this := V1alpha1ClusterGenerator{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V1alpha1ClusterGenerator) GetSelector() V1LabelSelector {
	if o == nil || isNil(o.Selector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterGenerator) GetSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || isNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V1alpha1ClusterGenerator) HasSelector() bool {
	if o != nil && !isNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given V1LabelSelector and assigns it to the Selector field.
func (o *V1alpha1ClusterGenerator) SetSelector(v V1LabelSelector) {
	o.Selector = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *V1alpha1ClusterGenerator) GetTemplate() V1alpha1ApplicationSetTemplate {
	if o == nil || isNil(o.Template) {
		var ret V1alpha1ApplicationSetTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool) {
	if o == nil || isNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *V1alpha1ClusterGenerator) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given V1alpha1ApplicationSetTemplate and assigns it to the Template field.
func (o *V1alpha1ClusterGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate) {
	o.Template = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *V1alpha1ClusterGenerator) GetValues() map[string]string {
	if o == nil || isNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterGenerator) GetValuesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *V1alpha1ClusterGenerator) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *V1alpha1ClusterGenerator) SetValues(v map[string]string) {
	o.Values = &v
}

func (o V1alpha1ClusterGenerator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ClusterGenerator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableV1alpha1ClusterGenerator struct {
	value *V1alpha1ClusterGenerator
	isSet bool
}

func (v NullableV1alpha1ClusterGenerator) Get() *V1alpha1ClusterGenerator {
	return v.value
}

func (v *NullableV1alpha1ClusterGenerator) Set(val *V1alpha1ClusterGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ClusterGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ClusterGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ClusterGenerator(val *V1alpha1ClusterGenerator) *NullableV1alpha1ClusterGenerator {
	return &NullableV1alpha1ClusterGenerator{value: val, isSet: true}
}

func (v NullableV1alpha1ClusterGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ClusterGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
