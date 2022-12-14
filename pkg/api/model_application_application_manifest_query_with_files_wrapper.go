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

// checks if the ApplicationApplicationManifestQueryWithFilesWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationApplicationManifestQueryWithFilesWrapper{}

// ApplicationApplicationManifestQueryWithFilesWrapper struct for ApplicationApplicationManifestQueryWithFilesWrapper
type ApplicationApplicationManifestQueryWithFilesWrapper struct {
	Chunk *ApplicationFileChunk                         `json:"chunk,omitempty"`
	Query *ApplicationApplicationManifestQueryWithFiles `json:"query,omitempty"`
}

// NewApplicationApplicationManifestQueryWithFilesWrapper instantiates a new ApplicationApplicationManifestQueryWithFilesWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationApplicationManifestQueryWithFilesWrapper() *ApplicationApplicationManifestQueryWithFilesWrapper {
	this := ApplicationApplicationManifestQueryWithFilesWrapper{}
	return &this
}

// NewApplicationApplicationManifestQueryWithFilesWrapperWithDefaults instantiates a new ApplicationApplicationManifestQueryWithFilesWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationApplicationManifestQueryWithFilesWrapperWithDefaults() *ApplicationApplicationManifestQueryWithFilesWrapper {
	this := ApplicationApplicationManifestQueryWithFilesWrapper{}
	return &this
}

// GetChunk returns the Chunk field value if set, zero value otherwise.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) GetChunk() ApplicationFileChunk {
	if o == nil || isNil(o.Chunk) {
		var ret ApplicationFileChunk
		return ret
	}
	return *o.Chunk
}

// GetChunkOk returns a tuple with the Chunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) GetChunkOk() (*ApplicationFileChunk, bool) {
	if o == nil || isNil(o.Chunk) {
		return nil, false
	}
	return o.Chunk, true
}

// HasChunk returns a boolean if a field has been set.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) HasChunk() bool {
	if o != nil && !isNil(o.Chunk) {
		return true
	}

	return false
}

// SetChunk gets a reference to the given ApplicationFileChunk and assigns it to the Chunk field.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) SetChunk(v ApplicationFileChunk) {
	o.Chunk = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) GetQuery() ApplicationApplicationManifestQueryWithFiles {
	if o == nil || isNil(o.Query) {
		var ret ApplicationApplicationManifestQueryWithFiles
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) GetQueryOk() (*ApplicationApplicationManifestQueryWithFiles, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given ApplicationApplicationManifestQueryWithFiles and assigns it to the Query field.
func (o *ApplicationApplicationManifestQueryWithFilesWrapper) SetQuery(v ApplicationApplicationManifestQueryWithFiles) {
	o.Query = &v
}

func (o ApplicationApplicationManifestQueryWithFilesWrapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationApplicationManifestQueryWithFilesWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Chunk) {
		toSerialize["chunk"] = o.Chunk
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableApplicationApplicationManifestQueryWithFilesWrapper struct {
	value *ApplicationApplicationManifestQueryWithFilesWrapper
	isSet bool
}

func (v NullableApplicationApplicationManifestQueryWithFilesWrapper) Get() *ApplicationApplicationManifestQueryWithFilesWrapper {
	return v.value
}

func (v *NullableApplicationApplicationManifestQueryWithFilesWrapper) Set(val *ApplicationApplicationManifestQueryWithFilesWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationApplicationManifestQueryWithFilesWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationApplicationManifestQueryWithFilesWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationApplicationManifestQueryWithFilesWrapper(val *ApplicationApplicationManifestQueryWithFilesWrapper) *NullableApplicationApplicationManifestQueryWithFilesWrapper {
	return &NullableApplicationApplicationManifestQueryWithFilesWrapper{value: val, isSet: true}
}

func (v NullableApplicationApplicationManifestQueryWithFilesWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationApplicationManifestQueryWithFilesWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
