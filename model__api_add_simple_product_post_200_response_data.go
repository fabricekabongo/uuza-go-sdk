/*
Product Management API

API for managing products, including adding, updating, and retrieving products and their variants.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiAddSimpleProductPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAddSimpleProductPost200ResponseData{}

// ApiAddSimpleProductPost200ResponseData struct for ApiAddSimpleProductPost200ResponseData
type ApiAddSimpleProductPost200ResponseData struct {
	ProductId *string `json:"product_id,omitempty"`
}

// NewApiAddSimpleProductPost200ResponseData instantiates a new ApiAddSimpleProductPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAddSimpleProductPost200ResponseData() *ApiAddSimpleProductPost200ResponseData {
	this := ApiAddSimpleProductPost200ResponseData{}
	return &this
}

// NewApiAddSimpleProductPost200ResponseDataWithDefaults instantiates a new ApiAddSimpleProductPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAddSimpleProductPost200ResponseDataWithDefaults() *ApiAddSimpleProductPost200ResponseData {
	this := ApiAddSimpleProductPost200ResponseData{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ApiAddSimpleProductPost200ResponseData) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAddSimpleProductPost200ResponseData) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ApiAddSimpleProductPost200ResponseData) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ApiAddSimpleProductPost200ResponseData) SetProductId(v string) {
	o.ProductId = &v
}

func (o ApiAddSimpleProductPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAddSimpleProductPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableApiAddSimpleProductPost200ResponseData struct {
	value *ApiAddSimpleProductPost200ResponseData
	isSet bool
}

func (v NullableApiAddSimpleProductPost200ResponseData) Get() *ApiAddSimpleProductPost200ResponseData {
	return v.value
}

func (v *NullableApiAddSimpleProductPost200ResponseData) Set(val *ApiAddSimpleProductPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAddSimpleProductPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAddSimpleProductPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAddSimpleProductPost200ResponseData(val *ApiAddSimpleProductPost200ResponseData) *NullableApiAddSimpleProductPost200ResponseData {
	return &NullableApiAddSimpleProductPost200ResponseData{value: val, isSet: true}
}

func (v NullableApiAddSimpleProductPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAddSimpleProductPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


