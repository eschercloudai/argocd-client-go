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

// checks if the V1alpha1ApplicationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationList{}

// V1alpha1ApplicationList struct for V1alpha1ApplicationList
type V1alpha1ApplicationList struct {
	Items []V1alpha1Application `json:"items,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewV1alpha1ApplicationList instantiates a new V1alpha1ApplicationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationList() *V1alpha1ApplicationList {
	this := V1alpha1ApplicationList{}
	return &this
}

// NewV1alpha1ApplicationListWithDefaults instantiates a new V1alpha1ApplicationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationListWithDefaults() *V1alpha1ApplicationList {
	this := V1alpha1ApplicationList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *V1alpha1ApplicationList) GetItems() []V1alpha1Application {
	if o == nil || isNil(o.Items) {
		var ret []V1alpha1Application
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationList) GetItemsOk() ([]V1alpha1Application, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *V1alpha1ApplicationList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []V1alpha1Application and assigns it to the Items field.
func (o *V1alpha1ApplicationList) SetItems(v []V1alpha1Application) {
	o.Items = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1ApplicationList) GetMetadata() V1ListMeta {
	if o == nil || isNil(o.Metadata) {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1ApplicationList) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1alpha1ApplicationList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1alpha1ApplicationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationList struct {
	value *V1alpha1ApplicationList
	isSet bool
}

func (v NullableV1alpha1ApplicationList) Get() *V1alpha1ApplicationList {
	return v.value
}

func (v *NullableV1alpha1ApplicationList) Set(val *V1alpha1ApplicationList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationList(val *V1alpha1ApplicationList) *NullableV1alpha1ApplicationList {
	return &NullableV1alpha1ApplicationList{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


