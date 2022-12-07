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

// checks if the V1alpha1SCMProviderGeneratorGithub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SCMProviderGeneratorGithub{}

// V1alpha1SCMProviderGeneratorGithub SCMProviderGeneratorGithub defines connection info specific to GitHub.
type V1alpha1SCMProviderGeneratorGithub struct {
	// Scan all branches instead of just the default branch.
	AllBranches *bool `json:"allBranches,omitempty"`
	// The GitHub API URL to talk to. If blank, use https://api.github.com/.
	Api *string `json:"api,omitempty"`
	// AppSecretName is a reference to a GitHub App repo-creds secret.
	AppSecretName *string `json:"appSecretName,omitempty"`
	// GitHub org to scan. Required.
	Organization *string `json:"organization,omitempty"`
	TokenRef *V1alpha1SecretRef `json:"tokenRef,omitempty"`
}

// NewV1alpha1SCMProviderGeneratorGithub instantiates a new V1alpha1SCMProviderGeneratorGithub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SCMProviderGeneratorGithub() *V1alpha1SCMProviderGeneratorGithub {
	this := V1alpha1SCMProviderGeneratorGithub{}
	return &this
}

// NewV1alpha1SCMProviderGeneratorGithubWithDefaults instantiates a new V1alpha1SCMProviderGeneratorGithub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SCMProviderGeneratorGithubWithDefaults() *V1alpha1SCMProviderGeneratorGithub {
	this := V1alpha1SCMProviderGeneratorGithub{}
	return &this
}

// GetAllBranches returns the AllBranches field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGithub) GetAllBranches() bool {
	if o == nil || isNil(o.AllBranches) {
		var ret bool
		return ret
	}
	return *o.AllBranches
}

// GetAllBranchesOk returns a tuple with the AllBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) GetAllBranchesOk() (*bool, bool) {
	if o == nil || isNil(o.AllBranches) {
		return nil, false
	}
	return o.AllBranches, true
}

// HasAllBranches returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) HasAllBranches() bool {
	if o != nil && !isNil(o.AllBranches) {
		return true
	}

	return false
}

// SetAllBranches gets a reference to the given bool and assigns it to the AllBranches field.
func (o *V1alpha1SCMProviderGeneratorGithub) SetAllBranches(v bool) {
	o.AllBranches = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGithub) GetApi() string {
	if o == nil || isNil(o.Api) {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) GetApiOk() (*string, bool) {
	if o == nil || isNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *V1alpha1SCMProviderGeneratorGithub) SetApi(v string) {
	o.Api = &v
}

// GetAppSecretName returns the AppSecretName field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGithub) GetAppSecretName() string {
	if o == nil || isNil(o.AppSecretName) {
		var ret string
		return ret
	}
	return *o.AppSecretName
}

// GetAppSecretNameOk returns a tuple with the AppSecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) GetAppSecretNameOk() (*string, bool) {
	if o == nil || isNil(o.AppSecretName) {
		return nil, false
	}
	return o.AppSecretName, true
}

// HasAppSecretName returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) HasAppSecretName() bool {
	if o != nil && !isNil(o.AppSecretName) {
		return true
	}

	return false
}

// SetAppSecretName gets a reference to the given string and assigns it to the AppSecretName field.
func (o *V1alpha1SCMProviderGeneratorGithub) SetAppSecretName(v string) {
	o.AppSecretName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGithub) GetOrganization() string {
	if o == nil || isNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) GetOrganizationOk() (*string, bool) {
	if o == nil || isNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *V1alpha1SCMProviderGeneratorGithub) SetOrganization(v string) {
	o.Organization = &v
}

// GetTokenRef returns the TokenRef field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGithub) GetTokenRef() V1alpha1SecretRef {
	if o == nil || isNil(o.TokenRef) {
		var ret V1alpha1SecretRef
		return ret
	}
	return *o.TokenRef
}

// GetTokenRefOk returns a tuple with the TokenRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) GetTokenRefOk() (*V1alpha1SecretRef, bool) {
	if o == nil || isNil(o.TokenRef) {
		return nil, false
	}
	return o.TokenRef, true
}

// HasTokenRef returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGithub) HasTokenRef() bool {
	if o != nil && !isNil(o.TokenRef) {
		return true
	}

	return false
}

// SetTokenRef gets a reference to the given V1alpha1SecretRef and assigns it to the TokenRef field.
func (o *V1alpha1SCMProviderGeneratorGithub) SetTokenRef(v V1alpha1SecretRef) {
	o.TokenRef = &v
}

func (o V1alpha1SCMProviderGeneratorGithub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SCMProviderGeneratorGithub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllBranches) {
		toSerialize["allBranches"] = o.AllBranches
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.AppSecretName) {
		toSerialize["appSecretName"] = o.AppSecretName
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.TokenRef) {
		toSerialize["tokenRef"] = o.TokenRef
	}
	return toSerialize, nil
}

type NullableV1alpha1SCMProviderGeneratorGithub struct {
	value *V1alpha1SCMProviderGeneratorGithub
	isSet bool
}

func (v NullableV1alpha1SCMProviderGeneratorGithub) Get() *V1alpha1SCMProviderGeneratorGithub {
	return v.value
}

func (v *NullableV1alpha1SCMProviderGeneratorGithub) Set(val *V1alpha1SCMProviderGeneratorGithub) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SCMProviderGeneratorGithub) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SCMProviderGeneratorGithub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SCMProviderGeneratorGithub(val *V1alpha1SCMProviderGeneratorGithub) *NullableV1alpha1SCMProviderGeneratorGithub {
	return &NullableV1alpha1SCMProviderGeneratorGithub{value: val, isSet: true}
}

func (v NullableV1alpha1SCMProviderGeneratorGithub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SCMProviderGeneratorGithub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


