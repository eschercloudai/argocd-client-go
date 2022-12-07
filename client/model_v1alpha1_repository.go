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

// checks if the V1alpha1Repository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1Repository{}

// V1alpha1Repository struct for V1alpha1Repository
type V1alpha1Repository struct {
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`
	// EnableLFS specifies whether git-lfs support should be enabled for this repo. Only valid for Git repositories.
	EnableLfs *bool `json:"enableLfs,omitempty"`
	EnableOCI *bool `json:"enableOCI,omitempty"`
	GithubAppEnterpriseBaseUrl *string `json:"githubAppEnterpriseBaseUrl,omitempty"`
	GithubAppID *string `json:"githubAppID,omitempty"`
	GithubAppInstallationID *string `json:"githubAppInstallationID,omitempty"`
	GithubAppPrivateKey *string `json:"githubAppPrivateKey,omitempty"`
	InheritedCreds *bool `json:"inheritedCreds,omitempty"`
	Insecure *bool `json:"insecure,omitempty"`
	InsecureIgnoreHostKey *bool `json:"insecureIgnoreHostKey,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Project *string `json:"project,omitempty"`
	Proxy *string `json:"proxy,omitempty"`
	Repo *string `json:"repo,omitempty"`
	// SSHPrivateKey contains the PEM data for authenticating at the repo server. Only used with Git repos.
	SshPrivateKey *string `json:"sshPrivateKey,omitempty"`
	TlsClientCertData *string `json:"tlsClientCertData,omitempty"`
	TlsClientCertKey *string `json:"tlsClientCertKey,omitempty"`
	// Type specifies the type of the repo. Can be either \"git\" or \"helm. \"git\" is assumed if empty or absent.
	Type *string `json:"type,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewV1alpha1Repository instantiates a new V1alpha1Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1Repository() *V1alpha1Repository {
	this := V1alpha1Repository{}
	return &this
}

// NewV1alpha1RepositoryWithDefaults instantiates a new V1alpha1Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1RepositoryWithDefaults() *V1alpha1Repository {
	this := V1alpha1Repository{}
	return &this
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetConnectionState() V1alpha1ConnectionState {
	if o == nil || isNil(o.ConnectionState) {
		var ret V1alpha1ConnectionState
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetConnectionStateOk() (*V1alpha1ConnectionState, bool) {
	if o == nil || isNil(o.ConnectionState) {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasConnectionState() bool {
	if o != nil && !isNil(o.ConnectionState) {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given V1alpha1ConnectionState and assigns it to the ConnectionState field.
func (o *V1alpha1Repository) SetConnectionState(v V1alpha1ConnectionState) {
	o.ConnectionState = &v
}

// GetEnableLfs returns the EnableLfs field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetEnableLfs() bool {
	if o == nil || isNil(o.EnableLfs) {
		var ret bool
		return ret
	}
	return *o.EnableLfs
}

// GetEnableLfsOk returns a tuple with the EnableLfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetEnableLfsOk() (*bool, bool) {
	if o == nil || isNil(o.EnableLfs) {
		return nil, false
	}
	return o.EnableLfs, true
}

// HasEnableLfs returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasEnableLfs() bool {
	if o != nil && !isNil(o.EnableLfs) {
		return true
	}

	return false
}

// SetEnableLfs gets a reference to the given bool and assigns it to the EnableLfs field.
func (o *V1alpha1Repository) SetEnableLfs(v bool) {
	o.EnableLfs = &v
}

// GetEnableOCI returns the EnableOCI field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetEnableOCI() bool {
	if o == nil || isNil(o.EnableOCI) {
		var ret bool
		return ret
	}
	return *o.EnableOCI
}

// GetEnableOCIOk returns a tuple with the EnableOCI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetEnableOCIOk() (*bool, bool) {
	if o == nil || isNil(o.EnableOCI) {
		return nil, false
	}
	return o.EnableOCI, true
}

// HasEnableOCI returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasEnableOCI() bool {
	if o != nil && !isNil(o.EnableOCI) {
		return true
	}

	return false
}

// SetEnableOCI gets a reference to the given bool and assigns it to the EnableOCI field.
func (o *V1alpha1Repository) SetEnableOCI(v bool) {
	o.EnableOCI = &v
}

// GetGithubAppEnterpriseBaseUrl returns the GithubAppEnterpriseBaseUrl field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetGithubAppEnterpriseBaseUrl() string {
	if o == nil || isNil(o.GithubAppEnterpriseBaseUrl) {
		var ret string
		return ret
	}
	return *o.GithubAppEnterpriseBaseUrl
}

// GetGithubAppEnterpriseBaseUrlOk returns a tuple with the GithubAppEnterpriseBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetGithubAppEnterpriseBaseUrlOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppEnterpriseBaseUrl) {
		return nil, false
	}
	return o.GithubAppEnterpriseBaseUrl, true
}

// HasGithubAppEnterpriseBaseUrl returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasGithubAppEnterpriseBaseUrl() bool {
	if o != nil && !isNil(o.GithubAppEnterpriseBaseUrl) {
		return true
	}

	return false
}

// SetGithubAppEnterpriseBaseUrl gets a reference to the given string and assigns it to the GithubAppEnterpriseBaseUrl field.
func (o *V1alpha1Repository) SetGithubAppEnterpriseBaseUrl(v string) {
	o.GithubAppEnterpriseBaseUrl = &v
}

// GetGithubAppID returns the GithubAppID field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetGithubAppID() string {
	if o == nil || isNil(o.GithubAppID) {
		var ret string
		return ret
	}
	return *o.GithubAppID
}

// GetGithubAppIDOk returns a tuple with the GithubAppID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetGithubAppIDOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppID) {
		return nil, false
	}
	return o.GithubAppID, true
}

// HasGithubAppID returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasGithubAppID() bool {
	if o != nil && !isNil(o.GithubAppID) {
		return true
	}

	return false
}

// SetGithubAppID gets a reference to the given string and assigns it to the GithubAppID field.
func (o *V1alpha1Repository) SetGithubAppID(v string) {
	o.GithubAppID = &v
}

// GetGithubAppInstallationID returns the GithubAppInstallationID field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetGithubAppInstallationID() string {
	if o == nil || isNil(o.GithubAppInstallationID) {
		var ret string
		return ret
	}
	return *o.GithubAppInstallationID
}

// GetGithubAppInstallationIDOk returns a tuple with the GithubAppInstallationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetGithubAppInstallationIDOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppInstallationID) {
		return nil, false
	}
	return o.GithubAppInstallationID, true
}

// HasGithubAppInstallationID returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasGithubAppInstallationID() bool {
	if o != nil && !isNil(o.GithubAppInstallationID) {
		return true
	}

	return false
}

// SetGithubAppInstallationID gets a reference to the given string and assigns it to the GithubAppInstallationID field.
func (o *V1alpha1Repository) SetGithubAppInstallationID(v string) {
	o.GithubAppInstallationID = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetGithubAppPrivateKey() string {
	if o == nil || isNil(o.GithubAppPrivateKey) {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppPrivateKey) {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasGithubAppPrivateKey() bool {
	if o != nil && !isNil(o.GithubAppPrivateKey) {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *V1alpha1Repository) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetInheritedCreds returns the InheritedCreds field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetInheritedCreds() bool {
	if o == nil || isNil(o.InheritedCreds) {
		var ret bool
		return ret
	}
	return *o.InheritedCreds
}

// GetInheritedCredsOk returns a tuple with the InheritedCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetInheritedCredsOk() (*bool, bool) {
	if o == nil || isNil(o.InheritedCreds) {
		return nil, false
	}
	return o.InheritedCreds, true
}

// HasInheritedCreds returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasInheritedCreds() bool {
	if o != nil && !isNil(o.InheritedCreds) {
		return true
	}

	return false
}

// SetInheritedCreds gets a reference to the given bool and assigns it to the InheritedCreds field.
func (o *V1alpha1Repository) SetInheritedCreds(v bool) {
	o.InheritedCreds = &v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetInsecure() bool {
	if o == nil || isNil(o.Insecure) {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetInsecureOk() (*bool, bool) {
	if o == nil || isNil(o.Insecure) {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasInsecure() bool {
	if o != nil && !isNil(o.Insecure) {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *V1alpha1Repository) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetInsecureIgnoreHostKey returns the InsecureIgnoreHostKey field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetInsecureIgnoreHostKey() bool {
	if o == nil || isNil(o.InsecureIgnoreHostKey) {
		var ret bool
		return ret
	}
	return *o.InsecureIgnoreHostKey
}

// GetInsecureIgnoreHostKeyOk returns a tuple with the InsecureIgnoreHostKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetInsecureIgnoreHostKeyOk() (*bool, bool) {
	if o == nil || isNil(o.InsecureIgnoreHostKey) {
		return nil, false
	}
	return o.InsecureIgnoreHostKey, true
}

// HasInsecureIgnoreHostKey returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasInsecureIgnoreHostKey() bool {
	if o != nil && !isNil(o.InsecureIgnoreHostKey) {
		return true
	}

	return false
}

// SetInsecureIgnoreHostKey gets a reference to the given bool and assigns it to the InsecureIgnoreHostKey field.
func (o *V1alpha1Repository) SetInsecureIgnoreHostKey(v bool) {
	o.InsecureIgnoreHostKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1Repository) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *V1alpha1Repository) SetPassword(v string) {
	o.Password = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetProject() string {
	if o == nil || isNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetProjectOk() (*string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *V1alpha1Repository) SetProject(v string) {
	o.Project = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetProxy() string {
	if o == nil || isNil(o.Proxy) {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetProxyOk() (*string, bool) {
	if o == nil || isNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasProxy() bool {
	if o != nil && !isNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *V1alpha1Repository) SetProxy(v string) {
	o.Proxy = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetRepo() string {
	if o == nil || isNil(o.Repo) {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetRepoOk() (*string, bool) {
	if o == nil || isNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasRepo() bool {
	if o != nil && !isNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *V1alpha1Repository) SetRepo(v string) {
	o.Repo = &v
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetSshPrivateKey() string {
	if o == nil || isNil(o.SshPrivateKey) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.SshPrivateKey) {
		return nil, false
	}
	return o.SshPrivateKey, true
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasSshPrivateKey() bool {
	if o != nil && !isNil(o.SshPrivateKey) {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given string and assigns it to the SshPrivateKey field.
func (o *V1alpha1Repository) SetSshPrivateKey(v string) {
	o.SshPrivateKey = &v
}

// GetTlsClientCertData returns the TlsClientCertData field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetTlsClientCertData() string {
	if o == nil || isNil(o.TlsClientCertData) {
		var ret string
		return ret
	}
	return *o.TlsClientCertData
}

// GetTlsClientCertDataOk returns a tuple with the TlsClientCertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetTlsClientCertDataOk() (*string, bool) {
	if o == nil || isNil(o.TlsClientCertData) {
		return nil, false
	}
	return o.TlsClientCertData, true
}

// HasTlsClientCertData returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasTlsClientCertData() bool {
	if o != nil && !isNil(o.TlsClientCertData) {
		return true
	}

	return false
}

// SetTlsClientCertData gets a reference to the given string and assigns it to the TlsClientCertData field.
func (o *V1alpha1Repository) SetTlsClientCertData(v string) {
	o.TlsClientCertData = &v
}

// GetTlsClientCertKey returns the TlsClientCertKey field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetTlsClientCertKey() string {
	if o == nil || isNil(o.TlsClientCertKey) {
		var ret string
		return ret
	}
	return *o.TlsClientCertKey
}

// GetTlsClientCertKeyOk returns a tuple with the TlsClientCertKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetTlsClientCertKeyOk() (*string, bool) {
	if o == nil || isNil(o.TlsClientCertKey) {
		return nil, false
	}
	return o.TlsClientCertKey, true
}

// HasTlsClientCertKey returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasTlsClientCertKey() bool {
	if o != nil && !isNil(o.TlsClientCertKey) {
		return true
	}

	return false
}

// SetTlsClientCertKey gets a reference to the given string and assigns it to the TlsClientCertKey field.
func (o *V1alpha1Repository) SetTlsClientCertKey(v string) {
	o.TlsClientCertKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1Repository) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *V1alpha1Repository) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Repository) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *V1alpha1Repository) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *V1alpha1Repository) SetUsername(v string) {
	o.Username = &v
}

func (o V1alpha1Repository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1Repository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConnectionState) {
		toSerialize["connectionState"] = o.ConnectionState
	}
	if !isNil(o.EnableLfs) {
		toSerialize["enableLfs"] = o.EnableLfs
	}
	if !isNil(o.EnableOCI) {
		toSerialize["enableOCI"] = o.EnableOCI
	}
	if !isNil(o.GithubAppEnterpriseBaseUrl) {
		toSerialize["githubAppEnterpriseBaseUrl"] = o.GithubAppEnterpriseBaseUrl
	}
	if !isNil(o.GithubAppID) {
		toSerialize["githubAppID"] = o.GithubAppID
	}
	if !isNil(o.GithubAppInstallationID) {
		toSerialize["githubAppInstallationID"] = o.GithubAppInstallationID
	}
	if !isNil(o.GithubAppPrivateKey) {
		toSerialize["githubAppPrivateKey"] = o.GithubAppPrivateKey
	}
	if !isNil(o.InheritedCreds) {
		toSerialize["inheritedCreds"] = o.InheritedCreds
	}
	if !isNil(o.Insecure) {
		toSerialize["insecure"] = o.Insecure
	}
	if !isNil(o.InsecureIgnoreHostKey) {
		toSerialize["insecureIgnoreHostKey"] = o.InsecureIgnoreHostKey
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !isNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !isNil(o.SshPrivateKey) {
		toSerialize["sshPrivateKey"] = o.SshPrivateKey
	}
	if !isNil(o.TlsClientCertData) {
		toSerialize["tlsClientCertData"] = o.TlsClientCertData
	}
	if !isNil(o.TlsClientCertKey) {
		toSerialize["tlsClientCertKey"] = o.TlsClientCertKey
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableV1alpha1Repository struct {
	value *V1alpha1Repository
	isSet bool
}

func (v NullableV1alpha1Repository) Get() *V1alpha1Repository {
	return v.value
}

func (v *NullableV1alpha1Repository) Set(val *V1alpha1Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1Repository) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1Repository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1Repository(val *V1alpha1Repository) *NullableV1alpha1Repository {
	return &NullableV1alpha1Repository{value: val, isSet: true}
}

func (v NullableV1alpha1Repository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1Repository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


