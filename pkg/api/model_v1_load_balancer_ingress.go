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

// checks if the V1LoadBalancerIngress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1LoadBalancerIngress{}

// V1LoadBalancerIngress LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
type V1LoadBalancerIngress struct {
	Hostname *string        `json:"hostname,omitempty"`
	Ip       *string        `json:"ip,omitempty"`
	Ports    []V1PortStatus `json:"ports,omitempty"`
}

// NewV1LoadBalancerIngress instantiates a new V1LoadBalancerIngress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LoadBalancerIngress() *V1LoadBalancerIngress {
	this := V1LoadBalancerIngress{}
	return &this
}

// NewV1LoadBalancerIngressWithDefaults instantiates a new V1LoadBalancerIngress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LoadBalancerIngressWithDefaults() *V1LoadBalancerIngress {
	this := V1LoadBalancerIngress{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V1LoadBalancerIngress) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoadBalancerIngress) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V1LoadBalancerIngress) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *V1LoadBalancerIngress) SetHostname(v string) {
	o.Hostname = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *V1LoadBalancerIngress) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoadBalancerIngress) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *V1LoadBalancerIngress) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *V1LoadBalancerIngress) SetIp(v string) {
	o.Ip = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *V1LoadBalancerIngress) GetPorts() []V1PortStatus {
	if o == nil || isNil(o.Ports) {
		var ret []V1PortStatus
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoadBalancerIngress) GetPortsOk() ([]V1PortStatus, bool) {
	if o == nil || isNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *V1LoadBalancerIngress) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []V1PortStatus and assigns it to the Ports field.
func (o *V1LoadBalancerIngress) SetPorts(v []V1PortStatus) {
	o.Ports = v
}

func (o V1LoadBalancerIngress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1LoadBalancerIngress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableV1LoadBalancerIngress struct {
	value *V1LoadBalancerIngress
	isSet bool
}

func (v NullableV1LoadBalancerIngress) Get() *V1LoadBalancerIngress {
	return v.value
}

func (v *NullableV1LoadBalancerIngress) Set(val *V1LoadBalancerIngress) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LoadBalancerIngress) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LoadBalancerIngress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LoadBalancerIngress(val *V1LoadBalancerIngress) *NullableV1LoadBalancerIngress {
	return &NullableV1LoadBalancerIngress{value: val, isSet: true}
}

func (v NullableV1LoadBalancerIngress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LoadBalancerIngress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}