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

// checks if the V1alpha1ResourceIgnoreDifferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ResourceIgnoreDifferences{}

// V1alpha1ResourceIgnoreDifferences ResourceIgnoreDifferences contains resource filter and list of json paths which should be ignored during comparison with live state.
type V1alpha1ResourceIgnoreDifferences struct {
	Group *string `json:"group,omitempty"`
	JqPathExpressions []string `json:"jqPathExpressions,omitempty"`
	JsonPointers []string `json:"jsonPointers,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ManagedFieldsManagers []string `json:"managedFieldsManagers,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewV1alpha1ResourceIgnoreDifferences instantiates a new V1alpha1ResourceIgnoreDifferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceIgnoreDifferences() *V1alpha1ResourceIgnoreDifferences {
	this := V1alpha1ResourceIgnoreDifferences{}
	return &this
}

// NewV1alpha1ResourceIgnoreDifferencesWithDefaults instantiates a new V1alpha1ResourceIgnoreDifferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceIgnoreDifferencesWithDefaults() *V1alpha1ResourceIgnoreDifferences {
	this := V1alpha1ResourceIgnoreDifferences{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1ResourceIgnoreDifferences) SetGroup(v string) {
	o.Group = &v
}

// GetJqPathExpressions returns the JqPathExpressions field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetJqPathExpressions() []string {
	if o == nil || isNil(o.JqPathExpressions) {
		var ret []string
		return ret
	}
	return o.JqPathExpressions
}

// GetJqPathExpressionsOk returns a tuple with the JqPathExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetJqPathExpressionsOk() ([]string, bool) {
	if o == nil || isNil(o.JqPathExpressions) {
		return nil, false
	}
	return o.JqPathExpressions, true
}

// HasJqPathExpressions returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasJqPathExpressions() bool {
	if o != nil && !isNil(o.JqPathExpressions) {
		return true
	}

	return false
}

// SetJqPathExpressions gets a reference to the given []string and assigns it to the JqPathExpressions field.
func (o *V1alpha1ResourceIgnoreDifferences) SetJqPathExpressions(v []string) {
	o.JqPathExpressions = v
}

// GetJsonPointers returns the JsonPointers field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetJsonPointers() []string {
	if o == nil || isNil(o.JsonPointers) {
		var ret []string
		return ret
	}
	return o.JsonPointers
}

// GetJsonPointersOk returns a tuple with the JsonPointers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetJsonPointersOk() ([]string, bool) {
	if o == nil || isNil(o.JsonPointers) {
		return nil, false
	}
	return o.JsonPointers, true
}

// HasJsonPointers returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasJsonPointers() bool {
	if o != nil && !isNil(o.JsonPointers) {
		return true
	}

	return false
}

// SetJsonPointers gets a reference to the given []string and assigns it to the JsonPointers field.
func (o *V1alpha1ResourceIgnoreDifferences) SetJsonPointers(v []string) {
	o.JsonPointers = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1ResourceIgnoreDifferences) SetKind(v string) {
	o.Kind = &v
}

// GetManagedFieldsManagers returns the ManagedFieldsManagers field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetManagedFieldsManagers() []string {
	if o == nil || isNil(o.ManagedFieldsManagers) {
		var ret []string
		return ret
	}
	return o.ManagedFieldsManagers
}

// GetManagedFieldsManagersOk returns a tuple with the ManagedFieldsManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetManagedFieldsManagersOk() ([]string, bool) {
	if o == nil || isNil(o.ManagedFieldsManagers) {
		return nil, false
	}
	return o.ManagedFieldsManagers, true
}

// HasManagedFieldsManagers returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasManagedFieldsManagers() bool {
	if o != nil && !isNil(o.ManagedFieldsManagers) {
		return true
	}

	return false
}

// SetManagedFieldsManagers gets a reference to the given []string and assigns it to the ManagedFieldsManagers field.
func (o *V1alpha1ResourceIgnoreDifferences) SetManagedFieldsManagers(v []string) {
	o.ManagedFieldsManagers = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceIgnoreDifferences) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1ResourceIgnoreDifferences) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceIgnoreDifferences) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1ResourceIgnoreDifferences) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1ResourceIgnoreDifferences) SetNamespace(v string) {
	o.Namespace = &v
}

func (o V1alpha1ResourceIgnoreDifferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ResourceIgnoreDifferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.JqPathExpressions) {
		toSerialize["jqPathExpressions"] = o.JqPathExpressions
	}
	if !isNil(o.JsonPointers) {
		toSerialize["jsonPointers"] = o.JsonPointers
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.ManagedFieldsManagers) {
		toSerialize["managedFieldsManagers"] = o.ManagedFieldsManagers
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableV1alpha1ResourceIgnoreDifferences struct {
	value *V1alpha1ResourceIgnoreDifferences
	isSet bool
}

func (v NullableV1alpha1ResourceIgnoreDifferences) Get() *V1alpha1ResourceIgnoreDifferences {
	return v.value
}

func (v *NullableV1alpha1ResourceIgnoreDifferences) Set(val *V1alpha1ResourceIgnoreDifferences) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceIgnoreDifferences) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceIgnoreDifferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceIgnoreDifferences(val *V1alpha1ResourceIgnoreDifferences) *NullableV1alpha1ResourceIgnoreDifferences {
	return &NullableV1alpha1ResourceIgnoreDifferences{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceIgnoreDifferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceIgnoreDifferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

