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

// checks if the ProjectSyncWindowsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSyncWindowsResponse{}

// ProjectSyncWindowsResponse struct for ProjectSyncWindowsResponse
type ProjectSyncWindowsResponse struct {
	Windows []V1alpha1SyncWindow `json:"windows,omitempty"`
}

// NewProjectSyncWindowsResponse instantiates a new ProjectSyncWindowsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSyncWindowsResponse() *ProjectSyncWindowsResponse {
	this := ProjectSyncWindowsResponse{}
	return &this
}

// NewProjectSyncWindowsResponseWithDefaults instantiates a new ProjectSyncWindowsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSyncWindowsResponseWithDefaults() *ProjectSyncWindowsResponse {
	this := ProjectSyncWindowsResponse{}
	return &this
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *ProjectSyncWindowsResponse) GetWindows() []V1alpha1SyncWindow {
	if o == nil || isNil(o.Windows) {
		var ret []V1alpha1SyncWindow
		return ret
	}
	return o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSyncWindowsResponse) GetWindowsOk() ([]V1alpha1SyncWindow, bool) {
	if o == nil || isNil(o.Windows) {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *ProjectSyncWindowsResponse) HasWindows() bool {
	if o != nil && !isNil(o.Windows) {
		return true
	}

	return false
}

// SetWindows gets a reference to the given []V1alpha1SyncWindow and assigns it to the Windows field.
func (o *ProjectSyncWindowsResponse) SetWindows(v []V1alpha1SyncWindow) {
	o.Windows = v
}

func (o ProjectSyncWindowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSyncWindowsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Windows) {
		toSerialize["windows"] = o.Windows
	}
	return toSerialize, nil
}

type NullableProjectSyncWindowsResponse struct {
	value *ProjectSyncWindowsResponse
	isSet bool
}

func (v NullableProjectSyncWindowsResponse) Get() *ProjectSyncWindowsResponse {
	return v.value
}

func (v *NullableProjectSyncWindowsResponse) Set(val *ProjectSyncWindowsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSyncWindowsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSyncWindowsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSyncWindowsResponse(val *ProjectSyncWindowsResponse) *NullableProjectSyncWindowsResponse {
	return &NullableProjectSyncWindowsResponse{value: val, isSet: true}
}

func (v NullableProjectSyncWindowsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSyncWindowsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
