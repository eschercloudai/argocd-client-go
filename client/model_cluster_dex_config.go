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

// checks if the ClusterDexConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterDexConfig{}

// ClusterDexConfig struct for ClusterDexConfig
type ClusterDexConfig struct {
	Connectors []ClusterConnector `json:"connectors,omitempty"`
}

// NewClusterDexConfig instantiates a new ClusterDexConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDexConfig() *ClusterDexConfig {
	this := ClusterDexConfig{}
	return &this
}

// NewClusterDexConfigWithDefaults instantiates a new ClusterDexConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDexConfigWithDefaults() *ClusterDexConfig {
	this := ClusterDexConfig{}
	return &this
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *ClusterDexConfig) GetConnectors() []ClusterConnector {
	if o == nil || isNil(o.Connectors) {
		var ret []ClusterConnector
		return ret
	}
	return o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDexConfig) GetConnectorsOk() ([]ClusterConnector, bool) {
	if o == nil || isNil(o.Connectors) {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *ClusterDexConfig) HasConnectors() bool {
	if o != nil && !isNil(o.Connectors) {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []ClusterConnector and assigns it to the Connectors field.
func (o *ClusterDexConfig) SetConnectors(v []ClusterConnector) {
	o.Connectors = v
}

func (o ClusterDexConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterDexConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Connectors) {
		toSerialize["connectors"] = o.Connectors
	}
	return toSerialize, nil
}

type NullableClusterDexConfig struct {
	value *ClusterDexConfig
	isSet bool
}

func (v NullableClusterDexConfig) Get() *ClusterDexConfig {
	return v.value
}

func (v *NullableClusterDexConfig) Set(val *ClusterDexConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDexConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDexConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDexConfig(val *ClusterDexConfig) *NullableClusterDexConfig {
	return &NullableClusterDexConfig{value: val, isSet: true}
}

func (v NullableClusterDexConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDexConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


