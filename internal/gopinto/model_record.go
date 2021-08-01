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

// Record struct for Record
type Record struct {
	Name  string      `json:"name"`
	Type  RecordType  `json:"type"`
	Class RecordClass `json:"class"`
	Ttl   *int32      `json:"ttl,omitempty"`
	Data  string      `json:"data"`
}

// NewRecord instantiates a new Record object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecord(name string, type_ RecordType, class RecordClass, data string) *Record {
	this := Record{}
	this.Name = name
	this.Type = type_
	this.Class = class
	this.Data = data
	return &this
}

// NewRecordWithDefaults instantiates a new Record object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordWithDefaults() *Record {
	this := Record{}
	return &this
}

// GetName returns the Name field value
func (o *Record) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Record) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Record) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Record) GetType() RecordType {
	if o == nil {
		var ret RecordType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Record) GetTypeOk() (*RecordType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Record) SetType(v RecordType) {
	o.Type = v
}

// GetClass returns the Class field value
func (o *Record) GetClass() RecordClass {
	if o == nil {
		var ret RecordClass
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *Record) GetClassOk() (*RecordClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *Record) SetClass(v RecordClass) {
	o.Class = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *Record) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *Record) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *Record) SetTtl(v int32) {
	o.Ttl = &v
}

// GetData returns the Data field value
func (o *Record) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Record) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Record) SetData(v string) {
	o.Data = v
}

func (o Record) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["class"] = o.Class
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRecord struct {
	value *Record
	isSet bool
}

func (v NullableRecord) Get() *Record {
	return v.value
}

func (v *NullableRecord) Set(val *Record) {
	v.value = val
	v.isSet = true
}

func (v NullableRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecord(val *Record) *NullableRecord {
	return &NullableRecord{value: val, isSet: true}
}

func (v NullableRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
