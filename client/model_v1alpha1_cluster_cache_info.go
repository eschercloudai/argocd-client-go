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

// checks if the V1alpha1ClusterCacheInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ClusterCacheInfo{}

// V1alpha1ClusterCacheInfo struct for V1alpha1ClusterCacheInfo
type V1alpha1ClusterCacheInfo struct {
	ApisCount *int32 `json:"apisCount,omitempty"`
	LastCacheSyncTime *string `json:"lastCacheSyncTime,omitempty"`
	ResourcesCount *int32 `json:"resourcesCount,omitempty"`
}

// NewV1alpha1ClusterCacheInfo instantiates a new V1alpha1ClusterCacheInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ClusterCacheInfo() *V1alpha1ClusterCacheInfo {
	this := V1alpha1ClusterCacheInfo{}
	return &this
}

// NewV1alpha1ClusterCacheInfoWithDefaults instantiates a new V1alpha1ClusterCacheInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ClusterCacheInfoWithDefaults() *V1alpha1ClusterCacheInfo {
	this := V1alpha1ClusterCacheInfo{}
	return &this
}

// GetApisCount returns the ApisCount field value if set, zero value otherwise.
func (o *V1alpha1ClusterCacheInfo) GetApisCount() int32 {
	if o == nil || isNil(o.ApisCount) {
		var ret int32
		return ret
	}
	return *o.ApisCount
}

// GetApisCountOk returns a tuple with the ApisCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterCacheInfo) GetApisCountOk() (*int32, bool) {
	if o == nil || isNil(o.ApisCount) {
		return nil, false
	}
	return o.ApisCount, true
}

// HasApisCount returns a boolean if a field has been set.
func (o *V1alpha1ClusterCacheInfo) HasApisCount() bool {
	if o != nil && !isNil(o.ApisCount) {
		return true
	}

	return false
}

// SetApisCount gets a reference to the given int32 and assigns it to the ApisCount field.
func (o *V1alpha1ClusterCacheInfo) SetApisCount(v int32) {
	o.ApisCount = &v
}

// GetLastCacheSyncTime returns the LastCacheSyncTime field value if set, zero value otherwise.
func (o *V1alpha1ClusterCacheInfo) GetLastCacheSyncTime() string {
	if o == nil || isNil(o.LastCacheSyncTime) {
		var ret string
		return ret
	}
	return *o.LastCacheSyncTime
}

// GetLastCacheSyncTimeOk returns a tuple with the LastCacheSyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterCacheInfo) GetLastCacheSyncTimeOk() (*string, bool) {
	if o == nil || isNil(o.LastCacheSyncTime) {
		return nil, false
	}
	return o.LastCacheSyncTime, true
}

// HasLastCacheSyncTime returns a boolean if a field has been set.
func (o *V1alpha1ClusterCacheInfo) HasLastCacheSyncTime() bool {
	if o != nil && !isNil(o.LastCacheSyncTime) {
		return true
	}

	return false
}

// SetLastCacheSyncTime gets a reference to the given string and assigns it to the LastCacheSyncTime field.
func (o *V1alpha1ClusterCacheInfo) SetLastCacheSyncTime(v string) {
	o.LastCacheSyncTime = &v
}

// GetResourcesCount returns the ResourcesCount field value if set, zero value otherwise.
func (o *V1alpha1ClusterCacheInfo) GetResourcesCount() int32 {
	if o == nil || isNil(o.ResourcesCount) {
		var ret int32
		return ret
	}
	return *o.ResourcesCount
}

// GetResourcesCountOk returns a tuple with the ResourcesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ClusterCacheInfo) GetResourcesCountOk() (*int32, bool) {
	if o == nil || isNil(o.ResourcesCount) {
		return nil, false
	}
	return o.ResourcesCount, true
}

// HasResourcesCount returns a boolean if a field has been set.
func (o *V1alpha1ClusterCacheInfo) HasResourcesCount() bool {
	if o != nil && !isNil(o.ResourcesCount) {
		return true
	}

	return false
}

// SetResourcesCount gets a reference to the given int32 and assigns it to the ResourcesCount field.
func (o *V1alpha1ClusterCacheInfo) SetResourcesCount(v int32) {
	o.ResourcesCount = &v
}

func (o V1alpha1ClusterCacheInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ClusterCacheInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApisCount) {
		toSerialize["apisCount"] = o.ApisCount
	}
	if !isNil(o.LastCacheSyncTime) {
		toSerialize["lastCacheSyncTime"] = o.LastCacheSyncTime
	}
	if !isNil(o.ResourcesCount) {
		toSerialize["resourcesCount"] = o.ResourcesCount
	}
	return toSerialize, nil
}

type NullableV1alpha1ClusterCacheInfo struct {
	value *V1alpha1ClusterCacheInfo
	isSet bool
}

func (v NullableV1alpha1ClusterCacheInfo) Get() *V1alpha1ClusterCacheInfo {
	return v.value
}

func (v *NullableV1alpha1ClusterCacheInfo) Set(val *V1alpha1ClusterCacheInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ClusterCacheInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ClusterCacheInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ClusterCacheInfo(val *V1alpha1ClusterCacheInfo) *NullableV1alpha1ClusterCacheInfo {
	return &NullableV1alpha1ClusterCacheInfo{value: val, isSet: true}
}

func (v NullableV1alpha1ClusterCacheInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ClusterCacheInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


