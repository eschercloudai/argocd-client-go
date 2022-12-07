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

// checks if the V1alpha1PullRequestGeneratorGitea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1PullRequestGeneratorGitea{}

// V1alpha1PullRequestGeneratorGitea PullRequestGenerator defines connection info specific to Gitea.
type V1alpha1PullRequestGeneratorGitea struct {
	Api *string `json:"api,omitempty"`
	// Allow insecure tls, for self-signed certificates; default: false.
	Insecure *bool `json:"insecure,omitempty"`
	// Gitea org or user to scan. Required.
	Owner *string `json:"owner,omitempty"`
	// Gitea repo name to scan. Required.
	Repo *string `json:"repo,omitempty"`
	TokenRef *V1alpha1SecretRef `json:"tokenRef,omitempty"`
}

// NewV1alpha1PullRequestGeneratorGitea instantiates a new V1alpha1PullRequestGeneratorGitea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1PullRequestGeneratorGitea() *V1alpha1PullRequestGeneratorGitea {
	this := V1alpha1PullRequestGeneratorGitea{}
	return &this
}

// NewV1alpha1PullRequestGeneratorGiteaWithDefaults instantiates a new V1alpha1PullRequestGeneratorGitea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1PullRequestGeneratorGiteaWithDefaults() *V1alpha1PullRequestGeneratorGitea {
	this := V1alpha1PullRequestGeneratorGitea{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorGitea) GetApi() string {
	if o == nil || isNil(o.Api) {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorGitea) GetApiOk() (*string, bool) {
	if o == nil || isNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorGitea) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *V1alpha1PullRequestGeneratorGitea) SetApi(v string) {
	o.Api = &v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorGitea) GetInsecure() bool {
	if o == nil || isNil(o.Insecure) {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorGitea) GetInsecureOk() (*bool, bool) {
	if o == nil || isNil(o.Insecure) {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorGitea) HasInsecure() bool {
	if o != nil && !isNil(o.Insecure) {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *V1alpha1PullRequestGeneratorGitea) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorGitea) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorGitea) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorGitea) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *V1alpha1PullRequestGeneratorGitea) SetOwner(v string) {
	o.Owner = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorGitea) GetRepo() string {
	if o == nil || isNil(o.Repo) {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorGitea) GetRepoOk() (*string, bool) {
	if o == nil || isNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorGitea) HasRepo() bool {
	if o != nil && !isNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *V1alpha1PullRequestGeneratorGitea) SetRepo(v string) {
	o.Repo = &v
}

// GetTokenRef returns the TokenRef field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorGitea) GetTokenRef() V1alpha1SecretRef {
	if o == nil || isNil(o.TokenRef) {
		var ret V1alpha1SecretRef
		return ret
	}
	return *o.TokenRef
}

// GetTokenRefOk returns a tuple with the TokenRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorGitea) GetTokenRefOk() (*V1alpha1SecretRef, bool) {
	if o == nil || isNil(o.TokenRef) {
		return nil, false
	}
	return o.TokenRef, true
}

// HasTokenRef returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorGitea) HasTokenRef() bool {
	if o != nil && !isNil(o.TokenRef) {
		return true
	}

	return false
}

// SetTokenRef gets a reference to the given V1alpha1SecretRef and assigns it to the TokenRef field.
func (o *V1alpha1PullRequestGeneratorGitea) SetTokenRef(v V1alpha1SecretRef) {
	o.TokenRef = &v
}

func (o V1alpha1PullRequestGeneratorGitea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1PullRequestGeneratorGitea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.Insecure) {
		toSerialize["insecure"] = o.Insecure
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !isNil(o.TokenRef) {
		toSerialize["tokenRef"] = o.TokenRef
	}
	return toSerialize, nil
}

type NullableV1alpha1PullRequestGeneratorGitea struct {
	value *V1alpha1PullRequestGeneratorGitea
	isSet bool
}

func (v NullableV1alpha1PullRequestGeneratorGitea) Get() *V1alpha1PullRequestGeneratorGitea {
	return v.value
}

func (v *NullableV1alpha1PullRequestGeneratorGitea) Set(val *V1alpha1PullRequestGeneratorGitea) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1PullRequestGeneratorGitea) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1PullRequestGeneratorGitea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1PullRequestGeneratorGitea(val *V1alpha1PullRequestGeneratorGitea) *NullableV1alpha1PullRequestGeneratorGitea {
	return &NullableV1alpha1PullRequestGeneratorGitea{value: val, isSet: true}
}

func (v NullableV1alpha1PullRequestGeneratorGitea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1PullRequestGeneratorGitea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


