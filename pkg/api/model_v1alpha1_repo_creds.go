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

// checks if the V1alpha1RepoCreds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1RepoCreds{}

// V1alpha1RepoCreds struct for V1alpha1RepoCreds
type V1alpha1RepoCreds struct {
	EnableOCI                  *bool   `json:"enableOCI,omitempty"`
	GithubAppEnterpriseBaseUrl *string `json:"githubAppEnterpriseBaseUrl,omitempty"`
	GithubAppID                *string `json:"githubAppID,omitempty"`
	GithubAppInstallationID    *string `json:"githubAppInstallationID,omitempty"`
	GithubAppPrivateKey        *string `json:"githubAppPrivateKey,omitempty"`
	Password                   *string `json:"password,omitempty"`
	SshPrivateKey              *string `json:"sshPrivateKey,omitempty"`
	TlsClientCertData          *string `json:"tlsClientCertData,omitempty"`
	TlsClientCertKey           *string `json:"tlsClientCertKey,omitempty"`
	// Type specifies the type of the repoCreds. Can be either \"git\" or \"helm. \"git\" is assumed if empty or absent.
	Type     *string `json:"type,omitempty"`
	Url      *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewV1alpha1RepoCreds instantiates a new V1alpha1RepoCreds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1RepoCreds() *V1alpha1RepoCreds {
	this := V1alpha1RepoCreds{}
	return &this
}

// NewV1alpha1RepoCredsWithDefaults instantiates a new V1alpha1RepoCreds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1RepoCredsWithDefaults() *V1alpha1RepoCreds {
	this := V1alpha1RepoCreds{}
	return &this
}

// GetEnableOCI returns the EnableOCI field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetEnableOCI() bool {
	if o == nil || isNil(o.EnableOCI) {
		var ret bool
		return ret
	}
	return *o.EnableOCI
}

// GetEnableOCIOk returns a tuple with the EnableOCI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetEnableOCIOk() (*bool, bool) {
	if o == nil || isNil(o.EnableOCI) {
		return nil, false
	}
	return o.EnableOCI, true
}

// HasEnableOCI returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasEnableOCI() bool {
	if o != nil && !isNil(o.EnableOCI) {
		return true
	}

	return false
}

// SetEnableOCI gets a reference to the given bool and assigns it to the EnableOCI field.
func (o *V1alpha1RepoCreds) SetEnableOCI(v bool) {
	o.EnableOCI = &v
}

// GetGithubAppEnterpriseBaseUrl returns the GithubAppEnterpriseBaseUrl field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetGithubAppEnterpriseBaseUrl() string {
	if o == nil || isNil(o.GithubAppEnterpriseBaseUrl) {
		var ret string
		return ret
	}
	return *o.GithubAppEnterpriseBaseUrl
}

// GetGithubAppEnterpriseBaseUrlOk returns a tuple with the GithubAppEnterpriseBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetGithubAppEnterpriseBaseUrlOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppEnterpriseBaseUrl) {
		return nil, false
	}
	return o.GithubAppEnterpriseBaseUrl, true
}

// HasGithubAppEnterpriseBaseUrl returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasGithubAppEnterpriseBaseUrl() bool {
	if o != nil && !isNil(o.GithubAppEnterpriseBaseUrl) {
		return true
	}

	return false
}

// SetGithubAppEnterpriseBaseUrl gets a reference to the given string and assigns it to the GithubAppEnterpriseBaseUrl field.
func (o *V1alpha1RepoCreds) SetGithubAppEnterpriseBaseUrl(v string) {
	o.GithubAppEnterpriseBaseUrl = &v
}

// GetGithubAppID returns the GithubAppID field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetGithubAppID() string {
	if o == nil || isNil(o.GithubAppID) {
		var ret string
		return ret
	}
	return *o.GithubAppID
}

// GetGithubAppIDOk returns a tuple with the GithubAppID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetGithubAppIDOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppID) {
		return nil, false
	}
	return o.GithubAppID, true
}

// HasGithubAppID returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasGithubAppID() bool {
	if o != nil && !isNil(o.GithubAppID) {
		return true
	}

	return false
}

// SetGithubAppID gets a reference to the given string and assigns it to the GithubAppID field.
func (o *V1alpha1RepoCreds) SetGithubAppID(v string) {
	o.GithubAppID = &v
}

// GetGithubAppInstallationID returns the GithubAppInstallationID field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetGithubAppInstallationID() string {
	if o == nil || isNil(o.GithubAppInstallationID) {
		var ret string
		return ret
	}
	return *o.GithubAppInstallationID
}

// GetGithubAppInstallationIDOk returns a tuple with the GithubAppInstallationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetGithubAppInstallationIDOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppInstallationID) {
		return nil, false
	}
	return o.GithubAppInstallationID, true
}

// HasGithubAppInstallationID returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasGithubAppInstallationID() bool {
	if o != nil && !isNil(o.GithubAppInstallationID) {
		return true
	}

	return false
}

// SetGithubAppInstallationID gets a reference to the given string and assigns it to the GithubAppInstallationID field.
func (o *V1alpha1RepoCreds) SetGithubAppInstallationID(v string) {
	o.GithubAppInstallationID = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetGithubAppPrivateKey() string {
	if o == nil || isNil(o.GithubAppPrivateKey) {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.GithubAppPrivateKey) {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasGithubAppPrivateKey() bool {
	if o != nil && !isNil(o.GithubAppPrivateKey) {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *V1alpha1RepoCreds) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *V1alpha1RepoCreds) SetPassword(v string) {
	o.Password = &v
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetSshPrivateKey() string {
	if o == nil || isNil(o.SshPrivateKey) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.SshPrivateKey) {
		return nil, false
	}
	return o.SshPrivateKey, true
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasSshPrivateKey() bool {
	if o != nil && !isNil(o.SshPrivateKey) {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given string and assigns it to the SshPrivateKey field.
func (o *V1alpha1RepoCreds) SetSshPrivateKey(v string) {
	o.SshPrivateKey = &v
}

// GetTlsClientCertData returns the TlsClientCertData field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetTlsClientCertData() string {
	if o == nil || isNil(o.TlsClientCertData) {
		var ret string
		return ret
	}
	return *o.TlsClientCertData
}

// GetTlsClientCertDataOk returns a tuple with the TlsClientCertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetTlsClientCertDataOk() (*string, bool) {
	if o == nil || isNil(o.TlsClientCertData) {
		return nil, false
	}
	return o.TlsClientCertData, true
}

// HasTlsClientCertData returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasTlsClientCertData() bool {
	if o != nil && !isNil(o.TlsClientCertData) {
		return true
	}

	return false
}

// SetTlsClientCertData gets a reference to the given string and assigns it to the TlsClientCertData field.
func (o *V1alpha1RepoCreds) SetTlsClientCertData(v string) {
	o.TlsClientCertData = &v
}

// GetTlsClientCertKey returns the TlsClientCertKey field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetTlsClientCertKey() string {
	if o == nil || isNil(o.TlsClientCertKey) {
		var ret string
		return ret
	}
	return *o.TlsClientCertKey
}

// GetTlsClientCertKeyOk returns a tuple with the TlsClientCertKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetTlsClientCertKeyOk() (*string, bool) {
	if o == nil || isNil(o.TlsClientCertKey) {
		return nil, false
	}
	return o.TlsClientCertKey, true
}

// HasTlsClientCertKey returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasTlsClientCertKey() bool {
	if o != nil && !isNil(o.TlsClientCertKey) {
		return true
	}

	return false
}

// SetTlsClientCertKey gets a reference to the given string and assigns it to the TlsClientCertKey field.
func (o *V1alpha1RepoCreds) SetTlsClientCertKey(v string) {
	o.TlsClientCertKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1RepoCreds) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V1alpha1RepoCreds) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *V1alpha1RepoCreds) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepoCreds) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *V1alpha1RepoCreds) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *V1alpha1RepoCreds) SetUsername(v string) {
	o.Username = &v
}

func (o V1alpha1RepoCreds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1RepoCreds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
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
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableV1alpha1RepoCreds struct {
	value *V1alpha1RepoCreds
	isSet bool
}

func (v NullableV1alpha1RepoCreds) Get() *V1alpha1RepoCreds {
	return v.value
}

func (v *NullableV1alpha1RepoCreds) Set(val *V1alpha1RepoCreds) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1RepoCreds) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1RepoCreds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1RepoCreds(val *V1alpha1RepoCreds) *NullableV1alpha1RepoCreds {
	return &NullableV1alpha1RepoCreds{value: val, isSet: true}
}

func (v NullableV1alpha1RepoCreds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1RepoCreds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
