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

// checks if the V1alpha1RepositoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1RepositoryList{}

// V1alpha1RepositoryList RepositoryList is a collection of Repositories.
type V1alpha1RepositoryList struct {
	Items    []V1alpha1Repository `json:"items,omitempty"`
	Metadata *V1ListMeta          `json:"metadata,omitempty"`
}

// NewV1alpha1RepositoryList instantiates a new V1alpha1RepositoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1RepositoryList() *V1alpha1RepositoryList {
	this := V1alpha1RepositoryList{}
	return &this
}

// NewV1alpha1RepositoryListWithDefaults instantiates a new V1alpha1RepositoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1RepositoryListWithDefaults() *V1alpha1RepositoryList {
	this := V1alpha1RepositoryList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *V1alpha1RepositoryList) GetItems() []V1alpha1Repository {
	if o == nil || isNil(o.Items) {
		var ret []V1alpha1Repository
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryList) GetItemsOk() ([]V1alpha1Repository, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *V1alpha1RepositoryList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []V1alpha1Repository and assigns it to the Items field.
func (o *V1alpha1RepositoryList) SetItems(v []V1alpha1Repository) {
	o.Items = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1RepositoryList) GetMetadata() V1ListMeta {
	if o == nil || isNil(o.Metadata) {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1RepositoryList) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1alpha1RepositoryList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1alpha1RepositoryList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1RepositoryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableV1alpha1RepositoryList struct {
	value *V1alpha1RepositoryList
	isSet bool
}

func (v NullableV1alpha1RepositoryList) Get() *V1alpha1RepositoryList {
	return v.value
}

func (v *NullableV1alpha1RepositoryList) Set(val *V1alpha1RepositoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1RepositoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1RepositoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1RepositoryList(val *V1alpha1RepositoryList) *NullableV1alpha1RepositoryList {
	return &NullableV1alpha1RepositoryList{value: val, isSet: true}
}

func (v NullableV1alpha1RepositoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1RepositoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}