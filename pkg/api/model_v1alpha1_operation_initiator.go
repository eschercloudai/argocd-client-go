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

// checks if the V1alpha1OperationInitiator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1OperationInitiator{}

// V1alpha1OperationInitiator struct for V1alpha1OperationInitiator
type V1alpha1OperationInitiator struct {
	// Automated is set to true if operation was initiated automatically by the application controller.
	Automated *bool   `json:"automated,omitempty"`
	Username  *string `json:"username,omitempty"`
}

// NewV1alpha1OperationInitiator instantiates a new V1alpha1OperationInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1OperationInitiator() *V1alpha1OperationInitiator {
	this := V1alpha1OperationInitiator{}
	return &this
}

// NewV1alpha1OperationInitiatorWithDefaults instantiates a new V1alpha1OperationInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1OperationInitiatorWithDefaults() *V1alpha1OperationInitiator {
	this := V1alpha1OperationInitiator{}
	return &this
}

// GetAutomated returns the Automated field value if set, zero value otherwise.
func (o *V1alpha1OperationInitiator) GetAutomated() bool {
	if o == nil || isNil(o.Automated) {
		var ret bool
		return ret
	}
	return *o.Automated
}

// GetAutomatedOk returns a tuple with the Automated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1OperationInitiator) GetAutomatedOk() (*bool, bool) {
	if o == nil || isNil(o.Automated) {
		return nil, false
	}
	return o.Automated, true
}

// HasAutomated returns a boolean if a field has been set.
func (o *V1alpha1OperationInitiator) HasAutomated() bool {
	if o != nil && !isNil(o.Automated) {
		return true
	}

	return false
}

// SetAutomated gets a reference to the given bool and assigns it to the Automated field.
func (o *V1alpha1OperationInitiator) SetAutomated(v bool) {
	o.Automated = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *V1alpha1OperationInitiator) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1OperationInitiator) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *V1alpha1OperationInitiator) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *V1alpha1OperationInitiator) SetUsername(v string) {
	o.Username = &v
}

func (o V1alpha1OperationInitiator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1OperationInitiator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Automated) {
		toSerialize["automated"] = o.Automated
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableV1alpha1OperationInitiator struct {
	value *V1alpha1OperationInitiator
	isSet bool
}

func (v NullableV1alpha1OperationInitiator) Get() *V1alpha1OperationInitiator {
	return v.value
}

func (v *NullableV1alpha1OperationInitiator) Set(val *V1alpha1OperationInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1OperationInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1OperationInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1OperationInitiator(val *V1alpha1OperationInitiator) *NullableV1alpha1OperationInitiator {
	return &NullableV1alpha1OperationInitiator{value: val, isSet: true}
}

func (v NullableV1alpha1OperationInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1OperationInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
