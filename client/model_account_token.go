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

// checks if the AccountToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountToken{}

// AccountToken struct for AccountToken
type AccountToken struct {
	ExpiresAt *string `json:"expiresAt,omitempty"`
	Id *string `json:"id,omitempty"`
	IssuedAt *string `json:"issuedAt,omitempty"`
}

// NewAccountToken instantiates a new AccountToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountToken() *AccountToken {
	this := AccountToken{}
	return &this
}

// NewAccountTokenWithDefaults instantiates a new AccountToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTokenWithDefaults() *AccountToken {
	this := AccountToken{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AccountToken) GetExpiresAt() string {
	if o == nil || isNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetExpiresAtOk() (*string, bool) {
	if o == nil || isNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AccountToken) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *AccountToken) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountToken) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountToken) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountToken) SetId(v string) {
	o.Id = &v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *AccountToken) GetIssuedAt() string {
	if o == nil || isNil(o.IssuedAt) {
		var ret string
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetIssuedAtOk() (*string, bool) {
	if o == nil || isNil(o.IssuedAt) {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *AccountToken) HasIssuedAt() bool {
	if o != nil && !isNil(o.IssuedAt) {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given string and assigns it to the IssuedAt field.
func (o *AccountToken) SetIssuedAt(v string) {
	o.IssuedAt = &v
}

func (o AccountToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IssuedAt) {
		toSerialize["issuedAt"] = o.IssuedAt
	}
	return toSerialize, nil
}

type NullableAccountToken struct {
	value *AccountToken
	isSet bool
}

func (v NullableAccountToken) Get() *AccountToken {
	return v.value
}

func (v *NullableAccountToken) Set(val *AccountToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountToken(val *AccountToken) *NullableAccountToken {
	return &NullableAccountToken{value: val, isSet: true}
}

func (v NullableAccountToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


