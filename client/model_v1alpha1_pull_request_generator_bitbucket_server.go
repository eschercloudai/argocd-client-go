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

// checks if the V1alpha1PullRequestGeneratorBitbucketServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1PullRequestGeneratorBitbucketServer{}

// V1alpha1PullRequestGeneratorBitbucketServer PullRequestGenerator defines connection info specific to BitbucketServer.
type V1alpha1PullRequestGeneratorBitbucketServer struct {
	// The Bitbucket REST API URL to talk to e.g. https://bitbucket.org/rest Required.
	Api *string `json:"api,omitempty"`
	BasicAuth *V1alpha1BasicAuthBitbucketServer `json:"basicAuth,omitempty"`
	// Project to scan. Required.
	Project *string `json:"project,omitempty"`
	// Repo name to scan. Required.
	Repo *string `json:"repo,omitempty"`
}

// NewV1alpha1PullRequestGeneratorBitbucketServer instantiates a new V1alpha1PullRequestGeneratorBitbucketServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1PullRequestGeneratorBitbucketServer() *V1alpha1PullRequestGeneratorBitbucketServer {
	this := V1alpha1PullRequestGeneratorBitbucketServer{}
	return &this
}

// NewV1alpha1PullRequestGeneratorBitbucketServerWithDefaults instantiates a new V1alpha1PullRequestGeneratorBitbucketServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1PullRequestGeneratorBitbucketServerWithDefaults() *V1alpha1PullRequestGeneratorBitbucketServer {
	this := V1alpha1PullRequestGeneratorBitbucketServer{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetApi() string {
	if o == nil || isNil(o.Api) {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetApiOk() (*string, bool) {
	if o == nil || isNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetApi(v string) {
	o.Api = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetBasicAuth() V1alpha1BasicAuthBitbucketServer {
	if o == nil || isNil(o.BasicAuth) {
		var ret V1alpha1BasicAuthBitbucketServer
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetBasicAuthOk() (*V1alpha1BasicAuthBitbucketServer, bool) {
	if o == nil || isNil(o.BasicAuth) {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasBasicAuth() bool {
	if o != nil && !isNil(o.BasicAuth) {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given V1alpha1BasicAuthBitbucketServer and assigns it to the BasicAuth field.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetBasicAuth(v V1alpha1BasicAuthBitbucketServer) {
	o.BasicAuth = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetProject() string {
	if o == nil || isNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetProjectOk() (*string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetProject(v string) {
	o.Project = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetRepo() string {
	if o == nil || isNil(o.Repo) {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) GetRepoOk() (*string, bool) {
	if o == nil || isNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) HasRepo() bool {
	if o != nil && !isNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *V1alpha1PullRequestGeneratorBitbucketServer) SetRepo(v string) {
	o.Repo = &v
}

func (o V1alpha1PullRequestGeneratorBitbucketServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1PullRequestGeneratorBitbucketServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.BasicAuth) {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	return toSerialize, nil
}

type NullableV1alpha1PullRequestGeneratorBitbucketServer struct {
	value *V1alpha1PullRequestGeneratorBitbucketServer
	isSet bool
}

func (v NullableV1alpha1PullRequestGeneratorBitbucketServer) Get() *V1alpha1PullRequestGeneratorBitbucketServer {
	return v.value
}

func (v *NullableV1alpha1PullRequestGeneratorBitbucketServer) Set(val *V1alpha1PullRequestGeneratorBitbucketServer) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1PullRequestGeneratorBitbucketServer) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1PullRequestGeneratorBitbucketServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1PullRequestGeneratorBitbucketServer(val *V1alpha1PullRequestGeneratorBitbucketServer) *NullableV1alpha1PullRequestGeneratorBitbucketServer {
	return &NullableV1alpha1PullRequestGeneratorBitbucketServer{value: val, isSet: true}
}

func (v NullableV1alpha1PullRequestGeneratorBitbucketServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1PullRequestGeneratorBitbucketServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


