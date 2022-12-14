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

// checks if the V1alpha1SCMProviderGeneratorFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SCMProviderGeneratorFilter{}

// V1alpha1SCMProviderGeneratorFilter SCMProviderGeneratorFilter is a single repository filter. If multiple filter types are set on a single struct, they will be AND'd together. All filters must pass for a repo to be included.
type V1alpha1SCMProviderGeneratorFilter struct {
	// A regex which must match the branch name.
	BranchMatch *string `json:"branchMatch,omitempty"`
	// A regex which must match at least one label.
	LabelMatch *string `json:"labelMatch,omitempty"`
	// An array of paths, all of which must not exist.
	PathsDoNotExist []string `json:"pathsDoNotExist,omitempty"`
	// An array of paths, all of which must exist.
	PathsExist []string `json:"pathsExist,omitempty"`
	// A regex for repo names.
	RepositoryMatch *string `json:"repositoryMatch,omitempty"`
}

// NewV1alpha1SCMProviderGeneratorFilter instantiates a new V1alpha1SCMProviderGeneratorFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SCMProviderGeneratorFilter() *V1alpha1SCMProviderGeneratorFilter {
	this := V1alpha1SCMProviderGeneratorFilter{}
	return &this
}

// NewV1alpha1SCMProviderGeneratorFilterWithDefaults instantiates a new V1alpha1SCMProviderGeneratorFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SCMProviderGeneratorFilterWithDefaults() *V1alpha1SCMProviderGeneratorFilter {
	this := V1alpha1SCMProviderGeneratorFilter{}
	return &this
}

// GetBranchMatch returns the BranchMatch field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorFilter) GetBranchMatch() string {
	if o == nil || isNil(o.BranchMatch) {
		var ret string
		return ret
	}
	return *o.BranchMatch
}

// GetBranchMatchOk returns a tuple with the BranchMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) GetBranchMatchOk() (*string, bool) {
	if o == nil || isNil(o.BranchMatch) {
		return nil, false
	}
	return o.BranchMatch, true
}

// HasBranchMatch returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) HasBranchMatch() bool {
	if o != nil && !isNil(o.BranchMatch) {
		return true
	}

	return false
}

// SetBranchMatch gets a reference to the given string and assigns it to the BranchMatch field.
func (o *V1alpha1SCMProviderGeneratorFilter) SetBranchMatch(v string) {
	o.BranchMatch = &v
}

// GetLabelMatch returns the LabelMatch field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorFilter) GetLabelMatch() string {
	if o == nil || isNil(o.LabelMatch) {
		var ret string
		return ret
	}
	return *o.LabelMatch
}

// GetLabelMatchOk returns a tuple with the LabelMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) GetLabelMatchOk() (*string, bool) {
	if o == nil || isNil(o.LabelMatch) {
		return nil, false
	}
	return o.LabelMatch, true
}

// HasLabelMatch returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) HasLabelMatch() bool {
	if o != nil && !isNil(o.LabelMatch) {
		return true
	}

	return false
}

// SetLabelMatch gets a reference to the given string and assigns it to the LabelMatch field.
func (o *V1alpha1SCMProviderGeneratorFilter) SetLabelMatch(v string) {
	o.LabelMatch = &v
}

// GetPathsDoNotExist returns the PathsDoNotExist field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsDoNotExist() []string {
	if o == nil || isNil(o.PathsDoNotExist) {
		var ret []string
		return ret
	}
	return o.PathsDoNotExist
}

// GetPathsDoNotExistOk returns a tuple with the PathsDoNotExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsDoNotExistOk() ([]string, bool) {
	if o == nil || isNil(o.PathsDoNotExist) {
		return nil, false
	}
	return o.PathsDoNotExist, true
}

// HasPathsDoNotExist returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) HasPathsDoNotExist() bool {
	if o != nil && !isNil(o.PathsDoNotExist) {
		return true
	}

	return false
}

// SetPathsDoNotExist gets a reference to the given []string and assigns it to the PathsDoNotExist field.
func (o *V1alpha1SCMProviderGeneratorFilter) SetPathsDoNotExist(v []string) {
	o.PathsDoNotExist = v
}

// GetPathsExist returns the PathsExist field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsExist() []string {
	if o == nil || isNil(o.PathsExist) {
		var ret []string
		return ret
	}
	return o.PathsExist
}

// GetPathsExistOk returns a tuple with the PathsExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) GetPathsExistOk() ([]string, bool) {
	if o == nil || isNil(o.PathsExist) {
		return nil, false
	}
	return o.PathsExist, true
}

// HasPathsExist returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) HasPathsExist() bool {
	if o != nil && !isNil(o.PathsExist) {
		return true
	}

	return false
}

// SetPathsExist gets a reference to the given []string and assigns it to the PathsExist field.
func (o *V1alpha1SCMProviderGeneratorFilter) SetPathsExist(v []string) {
	o.PathsExist = v
}

// GetRepositoryMatch returns the RepositoryMatch field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorFilter) GetRepositoryMatch() string {
	if o == nil || isNil(o.RepositoryMatch) {
		var ret string
		return ret
	}
	return *o.RepositoryMatch
}

// GetRepositoryMatchOk returns a tuple with the RepositoryMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) GetRepositoryMatchOk() (*string, bool) {
	if o == nil || isNil(o.RepositoryMatch) {
		return nil, false
	}
	return o.RepositoryMatch, true
}

// HasRepositoryMatch returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorFilter) HasRepositoryMatch() bool {
	if o != nil && !isNil(o.RepositoryMatch) {
		return true
	}

	return false
}

// SetRepositoryMatch gets a reference to the given string and assigns it to the RepositoryMatch field.
func (o *V1alpha1SCMProviderGeneratorFilter) SetRepositoryMatch(v string) {
	o.RepositoryMatch = &v
}

func (o V1alpha1SCMProviderGeneratorFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SCMProviderGeneratorFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BranchMatch) {
		toSerialize["branchMatch"] = o.BranchMatch
	}
	if !isNil(o.LabelMatch) {
		toSerialize["labelMatch"] = o.LabelMatch
	}
	if !isNil(o.PathsDoNotExist) {
		toSerialize["pathsDoNotExist"] = o.PathsDoNotExist
	}
	if !isNil(o.PathsExist) {
		toSerialize["pathsExist"] = o.PathsExist
	}
	if !isNil(o.RepositoryMatch) {
		toSerialize["repositoryMatch"] = o.RepositoryMatch
	}
	return toSerialize, nil
}

type NullableV1alpha1SCMProviderGeneratorFilter struct {
	value *V1alpha1SCMProviderGeneratorFilter
	isSet bool
}

func (v NullableV1alpha1SCMProviderGeneratorFilter) Get() *V1alpha1SCMProviderGeneratorFilter {
	return v.value
}

func (v *NullableV1alpha1SCMProviderGeneratorFilter) Set(val *V1alpha1SCMProviderGeneratorFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SCMProviderGeneratorFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SCMProviderGeneratorFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SCMProviderGeneratorFilter(val *V1alpha1SCMProviderGeneratorFilter) *NullableV1alpha1SCMProviderGeneratorFilter {
	return &NullableV1alpha1SCMProviderGeneratorFilter{value: val, isSet: true}
}

func (v NullableV1alpha1SCMProviderGeneratorFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SCMProviderGeneratorFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
