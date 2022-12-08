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

// checks if the V1alpha1TLSClientConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1TLSClientConfig{}

// V1alpha1TLSClientConfig struct for V1alpha1TLSClientConfig
type V1alpha1TLSClientConfig struct {
	CaData   *string `json:"caData,omitempty"`
	CertData *string `json:"certData,omitempty"`
	// Insecure specifies that the server should be accessed without verifying the TLS certificate. For testing only.
	Insecure *bool   `json:"insecure,omitempty"`
	KeyData  *string `json:"keyData,omitempty"`
	// ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.
	ServerName *string `json:"serverName,omitempty"`
}

// NewV1alpha1TLSClientConfig instantiates a new V1alpha1TLSClientConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1TLSClientConfig() *V1alpha1TLSClientConfig {
	this := V1alpha1TLSClientConfig{}
	return &this
}

// NewV1alpha1TLSClientConfigWithDefaults instantiates a new V1alpha1TLSClientConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1TLSClientConfigWithDefaults() *V1alpha1TLSClientConfig {
	this := V1alpha1TLSClientConfig{}
	return &this
}

// GetCaData returns the CaData field value if set, zero value otherwise.
func (o *V1alpha1TLSClientConfig) GetCaData() string {
	if o == nil || isNil(o.CaData) {
		var ret string
		return ret
	}
	return *o.CaData
}

// GetCaDataOk returns a tuple with the CaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1TLSClientConfig) GetCaDataOk() (*string, bool) {
	if o == nil || isNil(o.CaData) {
		return nil, false
	}
	return o.CaData, true
}

// HasCaData returns a boolean if a field has been set.
func (o *V1alpha1TLSClientConfig) HasCaData() bool {
	if o != nil && !isNil(o.CaData) {
		return true
	}

	return false
}

// SetCaData gets a reference to the given string and assigns it to the CaData field.
func (o *V1alpha1TLSClientConfig) SetCaData(v string) {
	o.CaData = &v
}

// GetCertData returns the CertData field value if set, zero value otherwise.
func (o *V1alpha1TLSClientConfig) GetCertData() string {
	if o == nil || isNil(o.CertData) {
		var ret string
		return ret
	}
	return *o.CertData
}

// GetCertDataOk returns a tuple with the CertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1TLSClientConfig) GetCertDataOk() (*string, bool) {
	if o == nil || isNil(o.CertData) {
		return nil, false
	}
	return o.CertData, true
}

// HasCertData returns a boolean if a field has been set.
func (o *V1alpha1TLSClientConfig) HasCertData() bool {
	if o != nil && !isNil(o.CertData) {
		return true
	}

	return false
}

// SetCertData gets a reference to the given string and assigns it to the CertData field.
func (o *V1alpha1TLSClientConfig) SetCertData(v string) {
	o.CertData = &v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *V1alpha1TLSClientConfig) GetInsecure() bool {
	if o == nil || isNil(o.Insecure) {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1TLSClientConfig) GetInsecureOk() (*bool, bool) {
	if o == nil || isNil(o.Insecure) {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *V1alpha1TLSClientConfig) HasInsecure() bool {
	if o != nil && !isNil(o.Insecure) {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *V1alpha1TLSClientConfig) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *V1alpha1TLSClientConfig) GetKeyData() string {
	if o == nil || isNil(o.KeyData) {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1TLSClientConfig) GetKeyDataOk() (*string, bool) {
	if o == nil || isNil(o.KeyData) {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *V1alpha1TLSClientConfig) HasKeyData() bool {
	if o != nil && !isNil(o.KeyData) {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *V1alpha1TLSClientConfig) SetKeyData(v string) {
	o.KeyData = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *V1alpha1TLSClientConfig) GetServerName() string {
	if o == nil || isNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1TLSClientConfig) GetServerNameOk() (*string, bool) {
	if o == nil || isNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *V1alpha1TLSClientConfig) HasServerName() bool {
	if o != nil && !isNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *V1alpha1TLSClientConfig) SetServerName(v string) {
	o.ServerName = &v
}

func (o V1alpha1TLSClientConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1TLSClientConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaData) {
		toSerialize["caData"] = o.CaData
	}
	if !isNil(o.CertData) {
		toSerialize["certData"] = o.CertData
	}
	if !isNil(o.Insecure) {
		toSerialize["insecure"] = o.Insecure
	}
	if !isNil(o.KeyData) {
		toSerialize["keyData"] = o.KeyData
	}
	if !isNil(o.ServerName) {
		toSerialize["serverName"] = o.ServerName
	}
	return toSerialize, nil
}

type NullableV1alpha1TLSClientConfig struct {
	value *V1alpha1TLSClientConfig
	isSet bool
}

func (v NullableV1alpha1TLSClientConfig) Get() *V1alpha1TLSClientConfig {
	return v.value
}

func (v *NullableV1alpha1TLSClientConfig) Set(val *V1alpha1TLSClientConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1TLSClientConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1TLSClientConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1TLSClientConfig(val *V1alpha1TLSClientConfig) *NullableV1alpha1TLSClientConfig {
	return &NullableV1alpha1TLSClientConfig{value: val, isSet: true}
}

func (v NullableV1alpha1TLSClientConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1TLSClientConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}