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

// checks if the V1alpha1ApplicationCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ApplicationCondition{}

// V1alpha1ApplicationCondition struct for V1alpha1ApplicationCondition
type V1alpha1ApplicationCondition struct {
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	Message            *string `json:"message,omitempty"`
	Type               *string `json:"type,omitempty"`
}

// NewV1alpha1ApplicationCondition instantiates a new V1alpha1ApplicationCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationCondition() *V1alpha1ApplicationCondition {
	this := V1alpha1ApplicationCondition{}
	return &this
}

// NewV1alpha1ApplicationConditionWithDefaults instantiates a new V1alpha1ApplicationCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationConditionWithDefaults() *V1alpha1ApplicationCondition {
	this := V1alpha1ApplicationCondition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *V1alpha1ApplicationCondition) GetLastTransitionTime() string {
	if o == nil || isNil(o.LastTransitionTime) {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationCondition) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || isNil(o.LastTransitionTime) {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *V1alpha1ApplicationCondition) HasLastTransitionTime() bool {
	if o != nil && !isNil(o.LastTransitionTime) {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *V1alpha1ApplicationCondition) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1alpha1ApplicationCondition) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationCondition) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1alpha1ApplicationCondition) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1alpha1ApplicationCondition) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1alpha1ApplicationCondition) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationCondition) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1alpha1ApplicationCondition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1alpha1ApplicationCondition) SetType(v string) {
	o.Type = &v
}

func (o V1alpha1ApplicationCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ApplicationCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LastTransitionTime) {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1alpha1ApplicationCondition struct {
	value *V1alpha1ApplicationCondition
	isSet bool
}

func (v NullableV1alpha1ApplicationCondition) Get() *V1alpha1ApplicationCondition {
	return v.value
}

func (v *NullableV1alpha1ApplicationCondition) Set(val *V1alpha1ApplicationCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationCondition(val *V1alpha1ApplicationCondition) *NullableV1alpha1ApplicationCondition {
	return &NullableV1alpha1ApplicationCondition{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
