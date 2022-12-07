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

// checks if the V1alpha1SyncOperationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1SyncOperationResource{}

// V1alpha1SyncOperationResource SyncOperationResource contains resources to sync.
type V1alpha1SyncOperationResource struct {
	Group *string `json:"group,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewV1alpha1SyncOperationResource instantiates a new V1alpha1SyncOperationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SyncOperationResource() *V1alpha1SyncOperationResource {
	this := V1alpha1SyncOperationResource{}
	return &this
}

// NewV1alpha1SyncOperationResourceWithDefaults instantiates a new V1alpha1SyncOperationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SyncOperationResourceWithDefaults() *V1alpha1SyncOperationResource {
	this := V1alpha1SyncOperationResource{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResource) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResource) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResource) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1SyncOperationResource) SetGroup(v string) {
	o.Group = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResource) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResource) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResource) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1SyncOperationResource) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResource) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResource) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResource) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1SyncOperationResource) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1SyncOperationResource) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SyncOperationResource) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1SyncOperationResource) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1SyncOperationResource) SetNamespace(v string) {
	o.Namespace = &v
}

func (o V1alpha1SyncOperationResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1SyncOperationResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableV1alpha1SyncOperationResource struct {
	value *V1alpha1SyncOperationResource
	isSet bool
}

func (v NullableV1alpha1SyncOperationResource) Get() *V1alpha1SyncOperationResource {
	return v.value
}

func (v *NullableV1alpha1SyncOperationResource) Set(val *V1alpha1SyncOperationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SyncOperationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SyncOperationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SyncOperationResource(val *V1alpha1SyncOperationResource) *NullableV1alpha1SyncOperationResource {
	return &NullableV1alpha1SyncOperationResource{value: val, isSet: true}
}

func (v NullableV1alpha1SyncOperationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SyncOperationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


