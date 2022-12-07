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

// checks if the V1MicroTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1MicroTime{}

// V1MicroTime MicroTime is version of Time with microsecond level precision.  +protobuf.options.marshal=false +protobuf.as=Timestamp +protobuf.options.(gogoproto.goproto_stringer)=false
type V1MicroTime struct {
	// Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.
	Nanos *int32 `json:"nanos,omitempty"`
	// Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.
	Seconds *string `json:"seconds,omitempty"`
}

// NewV1MicroTime instantiates a new V1MicroTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1MicroTime() *V1MicroTime {
	this := V1MicroTime{}
	return &this
}

// NewV1MicroTimeWithDefaults instantiates a new V1MicroTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MicroTimeWithDefaults() *V1MicroTime {
	this := V1MicroTime{}
	return &this
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *V1MicroTime) GetNanos() int32 {
	if o == nil || isNil(o.Nanos) {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MicroTime) GetNanosOk() (*int32, bool) {
	if o == nil || isNil(o.Nanos) {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *V1MicroTime) HasNanos() bool {
	if o != nil && !isNil(o.Nanos) {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *V1MicroTime) SetNanos(v int32) {
	o.Nanos = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V1MicroTime) GetSeconds() string {
	if o == nil || isNil(o.Seconds) {
		var ret string
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MicroTime) GetSecondsOk() (*string, bool) {
	if o == nil || isNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V1MicroTime) HasSeconds() bool {
	if o != nil && !isNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given string and assigns it to the Seconds field.
func (o *V1MicroTime) SetSeconds(v string) {
	o.Seconds = &v
}

func (o V1MicroTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1MicroTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Nanos) {
		toSerialize["nanos"] = o.Nanos
	}
	if !isNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	return toSerialize, nil
}

type NullableV1MicroTime struct {
	value *V1MicroTime
	isSet bool
}

func (v NullableV1MicroTime) Get() *V1MicroTime {
	return v.value
}

func (v *NullableV1MicroTime) Set(val *V1MicroTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MicroTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MicroTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MicroTime(val *V1MicroTime) *NullableV1MicroTime {
	return &NullableV1MicroTime{value: val, isSet: true}
}

func (v NullableV1MicroTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MicroTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


