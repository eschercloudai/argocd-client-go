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

// checks if the V1alpha1ResourceRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ResourceRef{}

// V1alpha1ResourceRef struct for V1alpha1ResourceRef
type V1alpha1ResourceRef struct {
	Group     *string `json:"group,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Uid       *string `json:"uid,omitempty"`
	Version   *string `json:"version,omitempty"`
}

// NewV1alpha1ResourceRef instantiates a new V1alpha1ResourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceRef() *V1alpha1ResourceRef {
	this := V1alpha1ResourceRef{}
	return &this
}

// NewV1alpha1ResourceRefWithDefaults instantiates a new V1alpha1ResourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceRefWithDefaults() *V1alpha1ResourceRef {
	this := V1alpha1ResourceRef{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1ResourceRef) SetGroup(v string) {
	o.Group = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1ResourceRef) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceRef) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1ResourceRef) SetNamespace(v string) {
	o.Namespace = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetUid() string {
	if o == nil || isNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetUidOk() (*string, bool) {
	if o == nil || isNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasUid() bool {
	if o != nil && !isNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *V1alpha1ResourceRef) SetUid(v string) {
	o.Uid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1alpha1ResourceRef) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceRef) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1alpha1ResourceRef) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1alpha1ResourceRef) SetVersion(v string) {
	o.Version = &v
}

func (o V1alpha1ResourceRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ResourceRef) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableV1alpha1ResourceRef struct {
	value *V1alpha1ResourceRef
	isSet bool
}

func (v NullableV1alpha1ResourceRef) Get() *V1alpha1ResourceRef {
	return v.value
}

func (v *NullableV1alpha1ResourceRef) Set(val *V1alpha1ResourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceRef(val *V1alpha1ResourceRef) *NullableV1alpha1ResourceRef {
	return &NullableV1alpha1ResourceRef{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
