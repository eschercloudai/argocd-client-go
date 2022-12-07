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

// checks if the V1alpha1MergeGenerator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1MergeGenerator{}

// V1alpha1MergeGenerator MergeGenerator merges the output of two or more generators. Where the values for all specified merge keys are equal between two sets of generated parameters, the parameter sets will be merged with the parameters from the latter generator taking precedence. Parameter sets with merge keys not present in the base generator's params will be ignored. For example, if the first generator produced [{a: '1', b: '2'}, {c: '1', d: '1'}] and the second generator produced [{'a': 'override'}], the united parameters for merge keys = ['a'] would be [{a: 'override', b: '1'}, {c: '1', d: '1'}].  MergeGenerator supports template overriding. If a MergeGenerator is one of multiple top-level generators, its template will be merged with the top-level generator before the parameters are applied.
type V1alpha1MergeGenerator struct {
	Generators []V1alpha1ApplicationSetNestedGenerator `json:"generators,omitempty"`
	MergeKeys []string `json:"mergeKeys,omitempty"`
	Template *V1alpha1ApplicationSetTemplate `json:"template,omitempty"`
}

// NewV1alpha1MergeGenerator instantiates a new V1alpha1MergeGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1MergeGenerator() *V1alpha1MergeGenerator {
	this := V1alpha1MergeGenerator{}
	return &this
}

// NewV1alpha1MergeGeneratorWithDefaults instantiates a new V1alpha1MergeGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1MergeGeneratorWithDefaults() *V1alpha1MergeGenerator {
	this := V1alpha1MergeGenerator{}
	return &this
}

// GetGenerators returns the Generators field value if set, zero value otherwise.
func (o *V1alpha1MergeGenerator) GetGenerators() []V1alpha1ApplicationSetNestedGenerator {
	if o == nil || isNil(o.Generators) {
		var ret []V1alpha1ApplicationSetNestedGenerator
		return ret
	}
	return o.Generators
}

// GetGeneratorsOk returns a tuple with the Generators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1MergeGenerator) GetGeneratorsOk() ([]V1alpha1ApplicationSetNestedGenerator, bool) {
	if o == nil || isNil(o.Generators) {
		return nil, false
	}
	return o.Generators, true
}

// HasGenerators returns a boolean if a field has been set.
func (o *V1alpha1MergeGenerator) HasGenerators() bool {
	if o != nil && !isNil(o.Generators) {
		return true
	}

	return false
}

// SetGenerators gets a reference to the given []V1alpha1ApplicationSetNestedGenerator and assigns it to the Generators field.
func (o *V1alpha1MergeGenerator) SetGenerators(v []V1alpha1ApplicationSetNestedGenerator) {
	o.Generators = v
}

// GetMergeKeys returns the MergeKeys field value if set, zero value otherwise.
func (o *V1alpha1MergeGenerator) GetMergeKeys() []string {
	if o == nil || isNil(o.MergeKeys) {
		var ret []string
		return ret
	}
	return o.MergeKeys
}

// GetMergeKeysOk returns a tuple with the MergeKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1MergeGenerator) GetMergeKeysOk() ([]string, bool) {
	if o == nil || isNil(o.MergeKeys) {
		return nil, false
	}
	return o.MergeKeys, true
}

// HasMergeKeys returns a boolean if a field has been set.
func (o *V1alpha1MergeGenerator) HasMergeKeys() bool {
	if o != nil && !isNil(o.MergeKeys) {
		return true
	}

	return false
}

// SetMergeKeys gets a reference to the given []string and assigns it to the MergeKeys field.
func (o *V1alpha1MergeGenerator) SetMergeKeys(v []string) {
	o.MergeKeys = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *V1alpha1MergeGenerator) GetTemplate() V1alpha1ApplicationSetTemplate {
	if o == nil || isNil(o.Template) {
		var ret V1alpha1ApplicationSetTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1MergeGenerator) GetTemplateOk() (*V1alpha1ApplicationSetTemplate, bool) {
	if o == nil || isNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *V1alpha1MergeGenerator) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given V1alpha1ApplicationSetTemplate and assigns it to the Template field.
func (o *V1alpha1MergeGenerator) SetTemplate(v V1alpha1ApplicationSetTemplate) {
	o.Template = &v
}

func (o V1alpha1MergeGenerator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1MergeGenerator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Generators) {
		toSerialize["generators"] = o.Generators
	}
	if !isNil(o.MergeKeys) {
		toSerialize["mergeKeys"] = o.MergeKeys
	}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableV1alpha1MergeGenerator struct {
	value *V1alpha1MergeGenerator
	isSet bool
}

func (v NullableV1alpha1MergeGenerator) Get() *V1alpha1MergeGenerator {
	return v.value
}

func (v *NullableV1alpha1MergeGenerator) Set(val *V1alpha1MergeGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1MergeGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1MergeGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1MergeGenerator(val *V1alpha1MergeGenerator) *NullableV1alpha1MergeGenerator {
	return &NullableV1alpha1MergeGenerator{value: val, isSet: true}
}

func (v NullableV1alpha1MergeGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1MergeGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


