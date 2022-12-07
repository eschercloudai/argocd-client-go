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

// checks if the V1alpha1ResourceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1ResourceResult{}

// V1alpha1ResourceResult struct for V1alpha1ResourceResult
type V1alpha1ResourceResult struct {
	Group *string `json:"group,omitempty"`
	// HookPhase contains the state of any operation associated with this resource OR hook This can also contain values for non-hook resources.
	HookPhase *string `json:"hookPhase,omitempty"`
	HookType *string `json:"hookType,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Message *string `json:"message,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Status *string `json:"status,omitempty"`
	SyncPhase *string `json:"syncPhase,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewV1alpha1ResourceResult instantiates a new V1alpha1ResourceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceResult() *V1alpha1ResourceResult {
	this := V1alpha1ResourceResult{}
	return &this
}

// NewV1alpha1ResourceResultWithDefaults instantiates a new V1alpha1ResourceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceResultWithDefaults() *V1alpha1ResourceResult {
	this := V1alpha1ResourceResult{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1ResourceResult) SetGroup(v string) {
	o.Group = &v
}

// GetHookPhase returns the HookPhase field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetHookPhase() string {
	if o == nil || isNil(o.HookPhase) {
		var ret string
		return ret
	}
	return *o.HookPhase
}

// GetHookPhaseOk returns a tuple with the HookPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetHookPhaseOk() (*string, bool) {
	if o == nil || isNil(o.HookPhase) {
		return nil, false
	}
	return o.HookPhase, true
}

// HasHookPhase returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasHookPhase() bool {
	if o != nil && !isNil(o.HookPhase) {
		return true
	}

	return false
}

// SetHookPhase gets a reference to the given string and assigns it to the HookPhase field.
func (o *V1alpha1ResourceResult) SetHookPhase(v string) {
	o.HookPhase = &v
}

// GetHookType returns the HookType field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetHookType() string {
	if o == nil || isNil(o.HookType) {
		var ret string
		return ret
	}
	return *o.HookType
}

// GetHookTypeOk returns a tuple with the HookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetHookTypeOk() (*string, bool) {
	if o == nil || isNil(o.HookType) {
		return nil, false
	}
	return o.HookType, true
}

// HasHookType returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasHookType() bool {
	if o != nil && !isNil(o.HookType) {
		return true
	}

	return false
}

// SetHookType gets a reference to the given string and assigns it to the HookType field.
func (o *V1alpha1ResourceResult) SetHookType(v string) {
	o.HookType = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1ResourceResult) SetKind(v string) {
	o.Kind = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1alpha1ResourceResult) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceResult) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1ResourceResult) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V1alpha1ResourceResult) SetStatus(v string) {
	o.Status = &v
}

// GetSyncPhase returns the SyncPhase field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetSyncPhase() string {
	if o == nil || isNil(o.SyncPhase) {
		var ret string
		return ret
	}
	return *o.SyncPhase
}

// GetSyncPhaseOk returns a tuple with the SyncPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetSyncPhaseOk() (*string, bool) {
	if o == nil || isNil(o.SyncPhase) {
		return nil, false
	}
	return o.SyncPhase, true
}

// HasSyncPhase returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasSyncPhase() bool {
	if o != nil && !isNil(o.SyncPhase) {
		return true
	}

	return false
}

// SetSyncPhase gets a reference to the given string and assigns it to the SyncPhase field.
func (o *V1alpha1ResourceResult) SetSyncPhase(v string) {
	o.SyncPhase = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1alpha1ResourceResult) SetVersion(v string) {
	o.Version = &v
}

func (o V1alpha1ResourceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1ResourceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.HookPhase) {
		toSerialize["hookPhase"] = o.HookPhase
	}
	if !isNil(o.HookType) {
		toSerialize["hookType"] = o.HookType
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.SyncPhase) {
		toSerialize["syncPhase"] = o.SyncPhase
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableV1alpha1ResourceResult struct {
	value *V1alpha1ResourceResult
	isSet bool
}

func (v NullableV1alpha1ResourceResult) Get() *V1alpha1ResourceResult {
	return v.value
}

func (v *NullableV1alpha1ResourceResult) Set(val *V1alpha1ResourceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceResult(val *V1alpha1ResourceResult) *NullableV1alpha1ResourceResult {
	return &NullableV1alpha1ResourceResult{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


