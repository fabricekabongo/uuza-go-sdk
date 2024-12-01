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

// checks if the ApiProductProductIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiProductProductIdGet200Response{}

// ApiProductProductIdGet200Response struct for ApiProductProductIdGet200Response
type ApiProductProductIdGet200Response struct {
	Message *string `json:"message,omitempty"`
	Success *bool `json:"success,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewApiProductProductIdGet200Response instantiates a new ApiProductProductIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiProductProductIdGet200Response() *ApiProductProductIdGet200Response {
	this := ApiProductProductIdGet200Response{}
	return &this
}

// NewApiProductProductIdGet200ResponseWithDefaults instantiates a new ApiProductProductIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiProductProductIdGet200ResponseWithDefaults() *ApiProductProductIdGet200Response {
	this := ApiProductProductIdGet200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiProductProductIdGet200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProductProductIdGet200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiProductProductIdGet200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiProductProductIdGet200Response) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiProductProductIdGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProductProductIdGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiProductProductIdGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiProductProductIdGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiProductProductIdGet200Response) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProductProductIdGet200Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiProductProductIdGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ApiProductProductIdGet200Response) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o ApiProductProductIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiProductProductIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableApiProductProductIdGet200Response struct {
	value *ApiProductProductIdGet200Response
	isSet bool
}

func (v NullableApiProductProductIdGet200Response) Get() *ApiProductProductIdGet200Response {
	return v.value
}

func (v *NullableApiProductProductIdGet200Response) Set(val *ApiProductProductIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiProductProductIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiProductProductIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiProductProductIdGet200Response(val *ApiProductProductIdGet200Response) *NullableApiProductProductIdGet200Response {
	return &NullableApiProductProductIdGet200Response{value: val, isSet: true}
}

func (v NullableApiProductProductIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiProductProductIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


