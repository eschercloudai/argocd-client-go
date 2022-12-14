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

// checks if the V1alpha1SCMProviderGeneratorBitbucketServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SCMProviderGeneratorBitbucketServer{}

// V1alpha1SCMProviderGeneratorBitbucketServer SCMProviderGeneratorBitbucketServer defines connection info specific to Bitbucket Server.
type V1alpha1SCMProviderGeneratorBitbucketServer struct {
	// Scan all branches instead of just the default branch.
	AllBranches *bool `json:"allBranches,omitempty"`
	// The Bitbucket Server REST API URL to talk to. Required.
	Api       *string                           `json:"api,omitempty"`
	BasicAuth *V1alpha1BasicAuthBitbucketServer `json:"basicAuth,omitempty"`
	// Project to scan. Required.
	Project *string `json:"project,omitempty"`
}

// NewV1alpha1SCMProviderGeneratorBitbucketServer instantiates a new V1alpha1SCMProviderGeneratorBitbucketServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SCMProviderGeneratorBitbucketServer() *V1alpha1SCMProviderGeneratorBitbucketServer {
	this := V1alpha1SCMProviderGeneratorBitbucketServer{}
	return &this
}

// NewV1alpha1SCMProviderGeneratorBitbucketServerWithDefaults instantiates a new V1alpha1SCMProviderGeneratorBitbucketServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SCMProviderGeneratorBitbucketServerWithDefaults() *V1alpha1SCMProviderGeneratorBitbucketServer {
	this := V1alpha1SCMProviderGeneratorBitbucketServer{}
	return &this
}

// GetAllBranches returns the AllBranches field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetAllBranches() bool {
	if o == nil || isNil(o.AllBranches) {
		var ret bool
		return ret
	}
	return *o.AllBranches
}

// GetAllBranchesOk returns a tuple with the AllBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetAllBranchesOk() (*bool, bool) {
	if o == nil || isNil(o.AllBranches) {
		return nil, false
	}
	return o.AllBranches, true
}

// HasAllBranches returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasAllBranches() bool {
	if o != nil && !isNil(o.AllBranches) {
		return true
	}

	return false
}

// SetAllBranches gets a reference to the given bool and assigns it to the AllBranches field.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetAllBranches(v bool) {
	o.AllBranches = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetApi() string {
	if o == nil || isNil(o.Api) {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetApiOk() (*string, bool) {
	if o == nil || isNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetApi(v string) {
	o.Api = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetBasicAuth() V1alpha1BasicAuthBitbucketServer {
	if o == nil || isNil(o.BasicAuth) {
		var ret V1alpha1BasicAuthBitbucketServer
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetBasicAuthOk() (*V1alpha1BasicAuthBitbucketServer, bool) {
	if o == nil || isNil(o.BasicAuth) {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasBasicAuth() bool {
	if o != nil && !isNil(o.BasicAuth) {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given V1alpha1BasicAuthBitbucketServer and assigns it to the BasicAuth field.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetBasicAuth(v V1alpha1BasicAuthBitbucketServer) {
	o.BasicAuth = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetProject() string {
	if o == nil || isNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) GetProjectOk() (*string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *V1alpha1SCMProviderGeneratorBitbucketServer) SetProject(v string) {
	o.Project = &v
}

func (o V1alpha1SCMProviderGeneratorBitbucketServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SCMProviderGeneratorBitbucketServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllBranches) {
		toSerialize["allBranches"] = o.AllBranches
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.BasicAuth) {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableV1alpha1SCMProviderGeneratorBitbucketServer struct {
	value *V1alpha1SCMProviderGeneratorBitbucketServer
	isSet bool
}

func (v NullableV1alpha1SCMProviderGeneratorBitbucketServer) Get() *V1alpha1SCMProviderGeneratorBitbucketServer {
	return v.value
}

func (v *NullableV1alpha1SCMProviderGeneratorBitbucketServer) Set(val *V1alpha1SCMProviderGeneratorBitbucketServer) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SCMProviderGeneratorBitbucketServer) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SCMProviderGeneratorBitbucketServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SCMProviderGeneratorBitbucketServer(val *V1alpha1SCMProviderGeneratorBitbucketServer) *NullableV1alpha1SCMProviderGeneratorBitbucketServer {
	return &NullableV1alpha1SCMProviderGeneratorBitbucketServer{value: val, isSet: true}
}

func (v NullableV1alpha1SCMProviderGeneratorBitbucketServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SCMProviderGeneratorBitbucketServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
