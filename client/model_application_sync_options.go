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

// checks if the ApplicationSyncOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSyncOptions{}

// ApplicationSyncOptions struct for ApplicationSyncOptions
type ApplicationSyncOptions struct {
	Items []string `json:"items,omitempty"`
}

// NewApplicationSyncOptions instantiates a new ApplicationSyncOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSyncOptions() *ApplicationSyncOptions {
	this := ApplicationSyncOptions{}
	return &this
}

// NewApplicationSyncOptionsWithDefaults instantiates a new ApplicationSyncOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSyncOptionsWithDefaults() *ApplicationSyncOptions {
	this := ApplicationSyncOptions{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ApplicationSyncOptions) GetItems() []string {
	if o == nil || isNil(o.Items) {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSyncOptions) GetItemsOk() ([]string, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApplicationSyncOptions) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *ApplicationSyncOptions) SetItems(v []string) {
	o.Items = v
}

func (o ApplicationSyncOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSyncOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableApplicationSyncOptions struct {
	value *ApplicationSyncOptions
	isSet bool
}

func (v NullableApplicationSyncOptions) Get() *ApplicationSyncOptions {
	return v.value
}

func (v *NullableApplicationSyncOptions) Set(val *ApplicationSyncOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSyncOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSyncOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSyncOptions(val *ApplicationSyncOptions) *NullableApplicationSyncOptions {
	return &NullableApplicationSyncOptions{value: val, isSet: true}
}

func (v NullableApplicationSyncOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSyncOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


