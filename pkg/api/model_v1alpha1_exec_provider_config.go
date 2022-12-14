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

// checks if the V1alpha1ExecProviderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ExecProviderConfig{}

// V1alpha1ExecProviderConfig struct for V1alpha1ExecProviderConfig
type V1alpha1ExecProviderConfig struct {
	ApiVersion  *string            `json:"apiVersion,omitempty"`
	Args        []string           `json:"args,omitempty"`
	Command     *string            `json:"command,omitempty"`
	Env         *map[string]string `json:"env,omitempty"`
	InstallHint *string            `json:"installHint,omitempty"`
}

// NewV1alpha1ExecProviderConfig instantiates a new V1alpha1ExecProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ExecProviderConfig() *V1alpha1ExecProviderConfig {
	this := V1alpha1ExecProviderConfig{}
	return &this
}

// NewV1alpha1ExecProviderConfigWithDefaults instantiates a new V1alpha1ExecProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ExecProviderConfigWithDefaults() *V1alpha1ExecProviderConfig {
	this := V1alpha1ExecProviderConfig{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1alpha1ExecProviderConfig) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ExecProviderConfig) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1alpha1ExecProviderConfig) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1alpha1ExecProviderConfig) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *V1alpha1ExecProviderConfig) GetArgs() []string {
	if o == nil || isNil(o.Args) {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ExecProviderConfig) GetArgsOk() ([]string, bool) {
	if o == nil || isNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *V1alpha1ExecProviderConfig) HasArgs() bool {
	if o != nil && !isNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *V1alpha1ExecProviderConfig) SetArgs(v []string) {
	o.Args = v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V1alpha1ExecProviderConfig) GetCommand() string {
	if o == nil || isNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ExecProviderConfig) GetCommandOk() (*string, bool) {
	if o == nil || isNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V1alpha1ExecProviderConfig) HasCommand() bool {
	if o != nil && !isNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *V1alpha1ExecProviderConfig) SetCommand(v string) {
	o.Command = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *V1alpha1ExecProviderConfig) GetEnv() map[string]string {
	if o == nil || isNil(o.Env) {
		var ret map[string]string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ExecProviderConfig) GetEnvOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *V1alpha1ExecProviderConfig) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given map[string]string and assigns it to the Env field.
func (o *V1alpha1ExecProviderConfig) SetEnv(v map[string]string) {
	o.Env = &v
}

// GetInstallHint returns the InstallHint field value if set, zero value otherwise.
func (o *V1alpha1ExecProviderConfig) GetInstallHint() string {
	if o == nil || isNil(o.InstallHint) {
		var ret string
		return ret
	}
	return *o.InstallHint
}

// GetInstallHintOk returns a tuple with the InstallHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ExecProviderConfig) GetInstallHintOk() (*string, bool) {
	if o == nil || isNil(o.InstallHint) {
		return nil, false
	}
	return o.InstallHint, true
}

// HasInstallHint returns a boolean if a field has been set.
func (o *V1alpha1ExecProviderConfig) HasInstallHint() bool {
	if o != nil && !isNil(o.InstallHint) {
		return true
	}

	return false
}

// SetInstallHint gets a reference to the given string and assigns it to the InstallHint field.
func (o *V1alpha1ExecProviderConfig) SetInstallHint(v string) {
	o.InstallHint = &v
}

func (o V1alpha1ExecProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ExecProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !isNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.InstallHint) {
		toSerialize["installHint"] = o.InstallHint
	}
	return toSerialize, nil
}

type NullableV1alpha1ExecProviderConfig struct {
	value *V1alpha1ExecProviderConfig
	isSet bool
}

func (v NullableV1alpha1ExecProviderConfig) Get() *V1alpha1ExecProviderConfig {
	return v.value
}

func (v *NullableV1alpha1ExecProviderConfig) Set(val *V1alpha1ExecProviderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ExecProviderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ExecProviderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ExecProviderConfig(val *V1alpha1ExecProviderConfig) *NullableV1alpha1ExecProviderConfig {
	return &NullableV1alpha1ExecProviderConfig{value: val, isSet: true}
}

func (v NullableV1alpha1ExecProviderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ExecProviderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
