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

// checks if the V1alpha1ClusterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ClusterInfo{}

// V1alpha1ClusterInfo struct for V1alpha1ClusterInfo
type V1alpha1ClusterInfo struct {
	ApiVersions []string `json:"apiVersions,omitempty"`
	ApplicationsCount *string `json:"applicationsCount,omitempty"`
	CacheInfo *V1alpha1ClusterCacheInfo `json:"cacheInfo,omitempty"`
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// NewV1alpha1ClusterInfo instantiates a new V1alpha1ClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ClusterInfo() *V1alpha1ClusterInfo {
	this := V1alpha1ClusterInfo{}
	return &this
}

// NewV1alpha1ClusterInfoWithDefaults instantiates a new V1alpha1ClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ClusterInfoWithDefaults() *V1alpha1ClusterInfo {
	this := V1alpha1ClusterInfo{}
	return &this
}

// GetApiVersions returns the ApiVersions field value if set, zero value otherwise.
func (o *V1alpha1ClusterInfo) GetApiVersions() []string {
	if o == nil || isNil(o.ApiVersions) {
		var ret []string
		return ret
	}
	return o.ApiVersions
}

// GetApiVersionsOk returns a tuple with the ApiVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterInfo) GetApiVersionsOk() ([]string, bool) {
	if o == nil || isNil(o.ApiVersions) {
		return nil, false
	}
	return o.ApiVersions, true
}

// HasApiVersions returns a boolean if a field has been set.
func (o *V1alpha1ClusterInfo) HasApiVersions() bool {
	if o != nil && !isNil(o.ApiVersions) {
		return true
	}

	return false
}

// SetApiVersions gets a reference to the given []string and assigns it to the ApiVersions field.
func (o *V1alpha1ClusterInfo) SetApiVersions(v []string) {
	o.ApiVersions = v
}

// GetApplicationsCount returns the ApplicationsCount field value if set, zero value otherwise.
func (o *V1alpha1ClusterInfo) GetApplicationsCount() string {
	if o == nil || isNil(o.ApplicationsCount) {
		var ret string
		return ret
	}
	return *o.ApplicationsCount
}

// GetApplicationsCountOk returns a tuple with the ApplicationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterInfo) GetApplicationsCountOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationsCount) {
		return nil, false
	}
	return o.ApplicationsCount, true
}

// HasApplicationsCount returns a boolean if a field has been set.
func (o *V1alpha1ClusterInfo) HasApplicationsCount() bool {
	if o != nil && !isNil(o.ApplicationsCount) {
		return true
	}

	return false
}

// SetApplicationsCount gets a reference to the given string and assigns it to the ApplicationsCount field.
func (o *V1alpha1ClusterInfo) SetApplicationsCount(v string) {
	o.ApplicationsCount = &v
}

// GetCacheInfo returns the CacheInfo field value if set, zero value otherwise.
func (o *V1alpha1ClusterInfo) GetCacheInfo() V1alpha1ClusterCacheInfo {
	if o == nil || isNil(o.CacheInfo) {
		var ret V1alpha1ClusterCacheInfo
		return ret
	}
	return *o.CacheInfo
}

// GetCacheInfoOk returns a tuple with the CacheInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterInfo) GetCacheInfoOk() (*V1alpha1ClusterCacheInfo, bool) {
	if o == nil || isNil(o.CacheInfo) {
		return nil, false
	}
	return o.CacheInfo, true
}

// HasCacheInfo returns a boolean if a field has been set.
func (o *V1alpha1ClusterInfo) HasCacheInfo() bool {
	if o != nil && !isNil(o.CacheInfo) {
		return true
	}

	return false
}

// SetCacheInfo gets a reference to the given V1alpha1ClusterCacheInfo and assigns it to the CacheInfo field.
func (o *V1alpha1ClusterInfo) SetCacheInfo(v V1alpha1ClusterCacheInfo) {
	o.CacheInfo = &v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *V1alpha1ClusterInfo) GetConnectionState() V1alpha1ConnectionState {
	if o == nil || isNil(o.ConnectionState) {
		var ret V1alpha1ConnectionState
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterInfo) GetConnectionStateOk() (*V1alpha1ConnectionState, bool) {
	if o == nil || isNil(o.ConnectionState) {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *V1alpha1ClusterInfo) HasConnectionState() bool {
	if o != nil && !isNil(o.ConnectionState) {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given V1alpha1ConnectionState and assigns it to the ConnectionState field.
func (o *V1alpha1ClusterInfo) SetConnectionState(v V1alpha1ConnectionState) {
	o.ConnectionState = &v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *V1alpha1ClusterInfo) GetServerVersion() string {
	if o == nil || isNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterInfo) GetServerVersionOk() (*string, bool) {
	if o == nil || isNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *V1alpha1ClusterInfo) HasServerVersion() bool {
	if o != nil && !isNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *V1alpha1ClusterInfo) SetServerVersion(v string) {
	o.ServerVersion = &v
}

func (o V1alpha1ClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ClusterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersions) {
		toSerialize["apiVersions"] = o.ApiVersions
	}
	if !isNil(o.ApplicationsCount) {
		toSerialize["applicationsCount"] = o.ApplicationsCount
	}
	if !isNil(o.CacheInfo) {
		toSerialize["cacheInfo"] = o.CacheInfo
	}
	if !isNil(o.ConnectionState) {
		toSerialize["connectionState"] = o.ConnectionState
	}
	if !isNil(o.ServerVersion) {
		toSerialize["serverVersion"] = o.ServerVersion
	}
	return toSerialize, nil
}

type NullableV1alpha1ClusterInfo struct {
	value *V1alpha1ClusterInfo
	isSet bool
}

func (v NullableV1alpha1ClusterInfo) Get() *V1alpha1ClusterInfo {
	return v.value
}

func (v *NullableV1alpha1ClusterInfo) Set(val *V1alpha1ClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ClusterInfo(val *V1alpha1ClusterInfo) *NullableV1alpha1ClusterInfo {
	return &NullableV1alpha1ClusterInfo{value: val, isSet: true}
}

func (v NullableV1alpha1ClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


