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

// checks if the NotificationTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationTrigger{}

// NotificationTrigger struct for NotificationTrigger
type NotificationTrigger struct {
	Name *string `json:"name,omitempty"`
}

// NewNotificationTrigger instantiates a new NotificationTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTrigger() *NotificationTrigger {
	this := NotificationTrigger{}
	return &this
}

// NewNotificationTriggerWithDefaults instantiates a new NotificationTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTriggerWithDefaults() *NotificationTrigger {
	this := NotificationTrigger{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationTrigger) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTrigger) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationTrigger) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationTrigger) SetName(v string) {
	o.Name = &v
}

func (o NotificationTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableNotificationTrigger struct {
	value *NotificationTrigger
	isSet bool
}

func (v NullableNotificationTrigger) Get() *NotificationTrigger {
	return v.value
}

func (v *NullableNotificationTrigger) Set(val *NotificationTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTrigger(val *NotificationTrigger) *NullableNotificationTrigger {
	return &NullableNotificationTrigger{value: val, isSet: true}
}

func (v NullableNotificationTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


