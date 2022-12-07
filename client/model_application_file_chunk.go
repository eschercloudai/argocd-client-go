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

// checks if the ApplicationFileChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFileChunk{}

// ApplicationFileChunk struct for ApplicationFileChunk
type ApplicationFileChunk struct {
	Chunk *string `json:"chunk,omitempty"`
}

// NewApplicationFileChunk instantiates a new ApplicationFileChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFileChunk() *ApplicationFileChunk {
	this := ApplicationFileChunk{}
	return &this
}

// NewApplicationFileChunkWithDefaults instantiates a new ApplicationFileChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFileChunkWithDefaults() *ApplicationFileChunk {
	this := ApplicationFileChunk{}
	return &this
}

// GetChunk returns the Chunk field value if set, zero value otherwise.
func (o *ApplicationFileChunk) GetChunk() string {
	if o == nil || isNil(o.Chunk) {
		var ret string
		return ret
	}
	return *o.Chunk
}

// GetChunkOk returns a tuple with the Chunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFileChunk) GetChunkOk() (*string, bool) {
	if o == nil || isNil(o.Chunk) {
		return nil, false
	}
	return o.Chunk, true
}

// HasChunk returns a boolean if a field has been set.
func (o *ApplicationFileChunk) HasChunk() bool {
	if o != nil && !isNil(o.Chunk) {
		return true
	}

	return false
}

// SetChunk gets a reference to the given string and assigns it to the Chunk field.
func (o *ApplicationFileChunk) SetChunk(v string) {
	o.Chunk = &v
}

func (o ApplicationFileChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFileChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Chunk) {
		toSerialize["chunk"] = o.Chunk
	}
	return toSerialize, nil
}

type NullableApplicationFileChunk struct {
	value *ApplicationFileChunk
	isSet bool
}

func (v NullableApplicationFileChunk) Get() *ApplicationFileChunk {
	return v.value
}

func (v *NullableApplicationFileChunk) Set(val *ApplicationFileChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFileChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFileChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFileChunk(val *ApplicationFileChunk) *NullableApplicationFileChunk {
	return &NullableApplicationFileChunk{value: val, isSet: true}
}

func (v NullableApplicationFileChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFileChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


