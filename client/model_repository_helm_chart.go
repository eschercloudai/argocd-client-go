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

// checks if the RepositoryHelmChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryHelmChart{}

// RepositoryHelmChart struct for RepositoryHelmChart
type RepositoryHelmChart struct {
	Name *string `json:"name,omitempty"`
	Versions []string `json:"versions,omitempty"`
}

// NewRepositoryHelmChart instantiates a new RepositoryHelmChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryHelmChart() *RepositoryHelmChart {
	this := RepositoryHelmChart{}
	return &this
}

// NewRepositoryHelmChartWithDefaults instantiates a new RepositoryHelmChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryHelmChartWithDefaults() *RepositoryHelmChart {
	this := RepositoryHelmChart{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryHelmChart) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHelmChart) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryHelmChart) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryHelmChart) SetName(v string) {
	o.Name = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *RepositoryHelmChart) GetVersions() []string {
	if o == nil || isNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHelmChart) GetVersionsOk() ([]string, bool) {
	if o == nil || isNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *RepositoryHelmChart) HasVersions() bool {
	if o != nil && !isNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *RepositoryHelmChart) SetVersions(v []string) {
	o.Versions = v
}

func (o RepositoryHelmChart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryHelmChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableRepositoryHelmChart struct {
	value *RepositoryHelmChart
	isSet bool
}

func (v NullableRepositoryHelmChart) Get() *RepositoryHelmChart {
	return v.value
}

func (v *NullableRepositoryHelmChart) Set(val *RepositoryHelmChart) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryHelmChart) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryHelmChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryHelmChart(val *RepositoryHelmChart) *NullableRepositoryHelmChart {
	return &NullableRepositoryHelmChart{value: val, isSet: true}
}

func (v NullableRepositoryHelmChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryHelmChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


