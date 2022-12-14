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

// checks if the V1alpha1ApplicationSetNestedGenerator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationSetNestedGenerator{}

// V1alpha1ApplicationSetNestedGenerator ApplicationSetNestedGenerator represents a generator nested within a combination-type generator (MatrixGenerator or MergeGenerator).
type V1alpha1ApplicationSetNestedGenerator struct {
	ClusterDecisionResource *V1alpha1DuckTypeGenerator    `json:"clusterDecisionResource,omitempty"`
	Clusters                *V1alpha1ClusterGenerator     `json:"clusters,omitempty"`
	Git                     *V1alpha1GitGenerator         `json:"git,omitempty"`
	List                    *V1alpha1ListGenerator        `json:"list,omitempty"`
	Matrix                  *V1JSON                       `json:"matrix,omitempty"`
	Merge                   *V1JSON                       `json:"merge,omitempty"`
	PullRequest             *V1alpha1PullRequestGenerator `json:"pullRequest,omitempty"`
	ScmProvider             *V1alpha1SCMProviderGenerator `json:"scmProvider,omitempty"`
	Selector                *V1LabelSelector              `json:"selector,omitempty"`
}

// NewV1alpha1ApplicationSetNestedGenerator instantiates a new V1alpha1ApplicationSetNestedGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSetNestedGenerator() *V1alpha1ApplicationSetNestedGenerator {
	this := V1alpha1ApplicationSetNestedGenerator{}
	return &this
}

// NewV1alpha1ApplicationSetNestedGeneratorWithDefaults instantiates a new V1alpha1ApplicationSetNestedGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSetNestedGeneratorWithDefaults() *V1alpha1ApplicationSetNestedGenerator {
	this := V1alpha1ApplicationSetNestedGenerator{}
	return &this
}

// GetClusterDecisionResource returns the ClusterDecisionResource field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetClusterDecisionResource() V1alpha1DuckTypeGenerator {
	if o == nil || isNil(o.ClusterDecisionResource) {
		var ret V1alpha1DuckTypeGenerator
		return ret
	}
	return *o.ClusterDecisionResource
}

// GetClusterDecisionResourceOk returns a tuple with the ClusterDecisionResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetClusterDecisionResourceOk() (*V1alpha1DuckTypeGenerator, bool) {
	if o == nil || isNil(o.ClusterDecisionResource) {
		return nil, false
	}
	return o.ClusterDecisionResource, true
}

// HasClusterDecisionResource returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasClusterDecisionResource() bool {
	if o != nil && !isNil(o.ClusterDecisionResource) {
		return true
	}

	return false
}

// SetClusterDecisionResource gets a reference to the given V1alpha1DuckTypeGenerator and assigns it to the ClusterDecisionResource field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetClusterDecisionResource(v V1alpha1DuckTypeGenerator) {
	o.ClusterDecisionResource = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetClusters() V1alpha1ClusterGenerator {
	if o == nil || isNil(o.Clusters) {
		var ret V1alpha1ClusterGenerator
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetClustersOk() (*V1alpha1ClusterGenerator, bool) {
	if o == nil || isNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasClusters() bool {
	if o != nil && !isNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given V1alpha1ClusterGenerator and assigns it to the Clusters field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetClusters(v V1alpha1ClusterGenerator) {
	o.Clusters = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetGit() V1alpha1GitGenerator {
	if o == nil || isNil(o.Git) {
		var ret V1alpha1GitGenerator
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetGitOk() (*V1alpha1GitGenerator, bool) {
	if o == nil || isNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given V1alpha1GitGenerator and assigns it to the Git field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetGit(v V1alpha1GitGenerator) {
	o.Git = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetList() V1alpha1ListGenerator {
	if o == nil || isNil(o.List) {
		var ret V1alpha1ListGenerator
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetListOk() (*V1alpha1ListGenerator, bool) {
	if o == nil || isNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasList() bool {
	if o != nil && !isNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given V1alpha1ListGenerator and assigns it to the List field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetList(v V1alpha1ListGenerator) {
	o.List = &v
}

// GetMatrix returns the Matrix field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetMatrix() V1JSON {
	if o == nil || isNil(o.Matrix) {
		var ret V1JSON
		return ret
	}
	return *o.Matrix
}

// GetMatrixOk returns a tuple with the Matrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetMatrixOk() (*V1JSON, bool) {
	if o == nil || isNil(o.Matrix) {
		return nil, false
	}
	return o.Matrix, true
}

// HasMatrix returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasMatrix() bool {
	if o != nil && !isNil(o.Matrix) {
		return true
	}

	return false
}

// SetMatrix gets a reference to the given V1JSON and assigns it to the Matrix field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetMatrix(v V1JSON) {
	o.Matrix = &v
}

// GetMerge returns the Merge field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetMerge() V1JSON {
	if o == nil || isNil(o.Merge) {
		var ret V1JSON
		return ret
	}
	return *o.Merge
}

// GetMergeOk returns a tuple with the Merge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetMergeOk() (*V1JSON, bool) {
	if o == nil || isNil(o.Merge) {
		return nil, false
	}
	return o.Merge, true
}

// HasMerge returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasMerge() bool {
	if o != nil && !isNil(o.Merge) {
		return true
	}

	return false
}

// SetMerge gets a reference to the given V1JSON and assigns it to the Merge field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetMerge(v V1JSON) {
	o.Merge = &v
}

// GetPullRequest returns the PullRequest field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetPullRequest() V1alpha1PullRequestGenerator {
	if o == nil || isNil(o.PullRequest) {
		var ret V1alpha1PullRequestGenerator
		return ret
	}
	return *o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetPullRequestOk() (*V1alpha1PullRequestGenerator, bool) {
	if o == nil || isNil(o.PullRequest) {
		return nil, false
	}
	return o.PullRequest, true
}

// HasPullRequest returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasPullRequest() bool {
	if o != nil && !isNil(o.PullRequest) {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given V1alpha1PullRequestGenerator and assigns it to the PullRequest field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetPullRequest(v V1alpha1PullRequestGenerator) {
	o.PullRequest = &v
}

// GetScmProvider returns the ScmProvider field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetScmProvider() V1alpha1SCMProviderGenerator {
	if o == nil || isNil(o.ScmProvider) {
		var ret V1alpha1SCMProviderGenerator
		return ret
	}
	return *o.ScmProvider
}

// GetScmProviderOk returns a tuple with the ScmProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetScmProviderOk() (*V1alpha1SCMProviderGenerator, bool) {
	if o == nil || isNil(o.ScmProvider) {
		return nil, false
	}
	return o.ScmProvider, true
}

// HasScmProvider returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasScmProvider() bool {
	if o != nil && !isNil(o.ScmProvider) {
		return true
	}

	return false
}

// SetScmProvider gets a reference to the given V1alpha1SCMProviderGenerator and assigns it to the ScmProvider field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetScmProvider(v V1alpha1SCMProviderGenerator) {
	o.ScmProvider = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSetNestedGenerator) GetSelector() V1LabelSelector {
	if o == nil || isNil(o.Selector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) GetSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || isNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSetNestedGenerator) HasSelector() bool {
	if o != nil && !isNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given V1LabelSelector and assigns it to the Selector field.
func (o *V1alpha1ApplicationSetNestedGenerator) SetSelector(v V1LabelSelector) {
	o.Selector = &v
}

func (o V1alpha1ApplicationSetNestedGenerator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationSetNestedGenerator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterDecisionResource) {
		toSerialize["clusterDecisionResource"] = o.ClusterDecisionResource
	}
	if !isNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	if !isNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !isNil(o.Matrix) {
		toSerialize["matrix"] = o.Matrix
	}
	if !isNil(o.Merge) {
		toSerialize["merge"] = o.Merge
	}
	if !isNil(o.PullRequest) {
		toSerialize["pullRequest"] = o.PullRequest
	}
	if !isNil(o.ScmProvider) {
		toSerialize["scmProvider"] = o.ScmProvider
	}
	if !isNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationSetNestedGenerator struct {
	value *V1alpha1ApplicationSetNestedGenerator
	isSet bool
}

func (v NullableV1alpha1ApplicationSetNestedGenerator) Get() *V1alpha1ApplicationSetNestedGenerator {
	return v.value
}

func (v *NullableV1alpha1ApplicationSetNestedGenerator) Set(val *V1alpha1ApplicationSetNestedGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSetNestedGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSetNestedGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSetNestedGenerator(val *V1alpha1ApplicationSetNestedGenerator) *NullableV1alpha1ApplicationSetNestedGenerator {
	return &NullableV1alpha1ApplicationSetNestedGenerator{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSetNestedGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSetNestedGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
