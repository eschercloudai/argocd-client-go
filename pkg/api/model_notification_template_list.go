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

// checks if the NotificationTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationTemplateList{}

// NotificationTemplateList struct for NotificationTemplateList
type NotificationTemplateList struct {
	Items []NotificationTemplate `json:"items,omitempty"`
}

// NewNotificationTemplateList instantiates a new NotificationTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTemplateList() *NotificationTemplateList {
	this := NotificationTemplateList{}
	return &this
}

// NewNotificationTemplateListWithDefaults instantiates a new NotificationTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTemplateListWithDefaults() *NotificationTemplateList {
	this := NotificationTemplateList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NotificationTemplateList) GetItems() []NotificationTemplate {
	if o == nil || isNil(o.Items) {
		var ret []NotificationTemplate
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTemplateList) GetItemsOk() ([]NotificationTemplate, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NotificationTemplateList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NotificationTemplate and assigns it to the Items field.
func (o *NotificationTemplateList) SetItems(v []NotificationTemplate) {
	o.Items = v
}

func (o NotificationTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationTemplateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableNotificationTemplateList struct {
	value *NotificationTemplateList
	isSet bool
}

func (v NullableNotificationTemplateList) Get() *NotificationTemplateList {
	return v.value
}

func (v *NullableNotificationTemplateList) Set(val *NotificationTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTemplateList(val *NotificationTemplateList) *NullableNotificationTemplateList {
	return &NullableNotificationTemplateList{value: val, isSet: true}
}

func (v NullableNotificationTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
