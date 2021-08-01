/*
 * Fava - OpenApi Gateway - DNS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopinto

import (
	"encoding/json"
)

// ApiOptions struct for ApiOptions
type ApiOptions struct {
	AccessOptions *AccessOptions     `json:"access_options,omitempty"`
	Meta          *map[string]string `json:"meta,omitempty"`
}

// NewApiOptions instantiates a new ApiOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOptions() *ApiOptions {
	this := ApiOptions{}
	return &this
}

// NewApiOptionsWithDefaults instantiates a new ApiOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOptionsWithDefaults() *ApiOptions {
	this := ApiOptions{}
	return &this
}

// GetAccessOptions returns the AccessOptions field value if set, zero value otherwise.
func (o *ApiOptions) GetAccessOptions() AccessOptions {
	if o == nil || o.AccessOptions == nil {
		var ret AccessOptions
		return ret
	}
	return *o.AccessOptions
}

// GetAccessOptionsOk returns a tuple with the AccessOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptions) GetAccessOptionsOk() (*AccessOptions, bool) {
	if o == nil || o.AccessOptions == nil {
		return nil, false
	}
	return o.AccessOptions, true
}

// HasAccessOptions returns a boolean if a field has been set.
func (o *ApiOptions) HasAccessOptions() bool {
	if o != nil && o.AccessOptions != nil {
		return true
	}

	return false
}

// SetAccessOptions gets a reference to the given AccessOptions and assigns it to the AccessOptions field.
func (o *ApiOptions) SetAccessOptions(v AccessOptions) {
	o.AccessOptions = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApiOptions) GetMeta() map[string]string {
	if o == nil || o.Meta == nil {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptions) GetMetaOk() (*map[string]string, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApiOptions) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *ApiOptions) SetMeta(v map[string]string) {
	o.Meta = &v
}

func (o ApiOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessOptions != nil {
		toSerialize["access_options"] = o.AccessOptions
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableApiOptions struct {
	value *ApiOptions
	isSet bool
}

func (v NullableApiOptions) Get() *ApiOptions {
	return v.value
}

func (v *NullableApiOptions) Set(val *ApiOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOptions(val *ApiOptions) *NullableApiOptions {
	return &NullableApiOptions{value: val, isSet: true}
}

func (v NullableApiOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
