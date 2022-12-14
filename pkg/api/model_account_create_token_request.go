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

// checks if the AccountCreateTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreateTokenRequest{}

// AccountCreateTokenRequest struct for AccountCreateTokenRequest
type AccountCreateTokenRequest struct {
	ExpiresIn *string `json:"expiresIn,omitempty"`
	Id        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// NewAccountCreateTokenRequest instantiates a new AccountCreateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateTokenRequest() *AccountCreateTokenRequest {
	this := AccountCreateTokenRequest{}
	return &this
}

// NewAccountCreateTokenRequestWithDefaults instantiates a new AccountCreateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateTokenRequestWithDefaults() *AccountCreateTokenRequest {
	this := AccountCreateTokenRequest{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AccountCreateTokenRequest) GetExpiresIn() string {
	if o == nil || isNil(o.ExpiresIn) {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateTokenRequest) GetExpiresInOk() (*string, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AccountCreateTokenRequest) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *AccountCreateTokenRequest) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountCreateTokenRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateTokenRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountCreateTokenRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountCreateTokenRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountCreateTokenRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateTokenRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountCreateTokenRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountCreateTokenRequest) SetName(v string) {
	o.Name = &v
}

func (o AccountCreateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreateTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAccountCreateTokenRequest struct {
	value *AccountCreateTokenRequest
	isSet bool
}

func (v NullableAccountCreateTokenRequest) Get() *AccountCreateTokenRequest {
	return v.value
}

func (v *NullableAccountCreateTokenRequest) Set(val *AccountCreateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateTokenRequest(val *AccountCreateTokenRequest) *NullableAccountCreateTokenRequest {
	return &NullableAccountCreateTokenRequest{value: val, isSet: true}
}

func (v NullableAccountCreateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
