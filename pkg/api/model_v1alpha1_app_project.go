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

// checks if the V1alpha1AppProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1AppProject{}

// V1alpha1AppProject struct for V1alpha1AppProject
type V1alpha1AppProject struct {
	Metadata *V1ObjectMeta             `json:"metadata,omitempty"`
	Spec     *V1alpha1AppProjectSpec   `json:"spec,omitempty"`
	Status   *V1alpha1AppProjectStatus `json:"status,omitempty"`
}

// NewV1alpha1AppProject instantiates a new V1alpha1AppProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1AppProject() *V1alpha1AppProject {
	this := V1alpha1AppProject{}
	return &this
}

// NewV1alpha1AppProjectWithDefaults instantiates a new V1alpha1AppProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1AppProjectWithDefaults() *V1alpha1AppProject {
	this := V1alpha1AppProject{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1AppProject) GetMetadata() V1ObjectMeta {
	if o == nil || isNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProject) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1AppProject) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1alpha1AppProject) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *V1alpha1AppProject) GetSpec() V1alpha1AppProjectSpec {
	if o == nil || isNil(o.Spec) {
		var ret V1alpha1AppProjectSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProject) GetSpecOk() (*V1alpha1AppProjectSpec, bool) {
	if o == nil || isNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *V1alpha1AppProject) HasSpec() bool {
	if o != nil && !isNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given V1alpha1AppProjectSpec and assigns it to the Spec field.
func (o *V1alpha1AppProject) SetSpec(v V1alpha1AppProjectSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha1AppProject) GetStatus() V1alpha1AppProjectStatus {
	if o == nil || isNil(o.Status) {
		var ret V1alpha1AppProjectStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProject) GetStatusOk() (*V1alpha1AppProjectStatus, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha1AppProject) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1alpha1AppProjectStatus and assigns it to the Status field.
func (o *V1alpha1AppProject) SetStatus(v V1alpha1AppProjectStatus) {
	o.Status = &v
}

func (o V1alpha1AppProject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1AppProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV1alpha1AppProject struct {
	value *V1alpha1AppProject
	isSet bool
}

func (v NullableV1alpha1AppProject) Get() *V1alpha1AppProject {
	return v.value
}

func (v *NullableV1alpha1AppProject) Set(val *V1alpha1AppProject) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1AppProject) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1AppProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1AppProject(val *V1alpha1AppProject) *NullableV1alpha1AppProject {
	return &NullableV1alpha1AppProject{value: val, isSet: true}
}

func (v NullableV1alpha1AppProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1AppProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
