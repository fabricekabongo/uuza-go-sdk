/*
Product Management API

API for managing products, including adding, updating, and retrieving products and their variants.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Variant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Variant{}

// Variant struct for Variant
type Variant struct {
	Name *string `json:"name,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	PostedDate *time.Time `json:"posted_date,omitempty"`
	ColorId *int32 `json:"color_id,omitempty"`
	Size []string `json:"size,omitempty"`
	Images []string `json:"images,omitempty"`
}

// NewVariant instantiates a new Variant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariant() *Variant {
	this := Variant{}
	return &this
}

// NewVariantWithDefaults instantiates a new Variant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantWithDefaults() *Variant {
	this := Variant{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Variant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Variant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Variant) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Variant) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Variant) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *Variant) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Variant) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Variant) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Variant) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPostedDate returns the PostedDate field value if set, zero value otherwise.
func (o *Variant) GetPostedDate() time.Time {
	if o == nil || IsNil(o.PostedDate) {
		var ret time.Time
		return ret
	}
	return *o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetPostedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PostedDate) {
		return nil, false
	}
	return o.PostedDate, true
}

// HasPostedDate returns a boolean if a field has been set.
func (o *Variant) HasPostedDate() bool {
	if o != nil && !IsNil(o.PostedDate) {
		return true
	}

	return false
}

// SetPostedDate gets a reference to the given time.Time and assigns it to the PostedDate field.
func (o *Variant) SetPostedDate(v time.Time) {
	o.PostedDate = &v
}

// GetColorId returns the ColorId field value if set, zero value otherwise.
func (o *Variant) GetColorId() int32 {
	if o == nil || IsNil(o.ColorId) {
		var ret int32
		return ret
	}
	return *o.ColorId
}

// GetColorIdOk returns a tuple with the ColorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetColorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ColorId) {
		return nil, false
	}
	return o.ColorId, true
}

// HasColorId returns a boolean if a field has been set.
func (o *Variant) HasColorId() bool {
	if o != nil && !IsNil(o.ColorId) {
		return true
	}

	return false
}

// SetColorId gets a reference to the given int32 and assigns it to the ColorId field.
func (o *Variant) SetColorId(v int32) {
	o.ColorId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Variant) GetSize() []string {
	if o == nil || IsNil(o.Size) {
		var ret []string
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetSizeOk() ([]string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Variant) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given []string and assigns it to the Size field.
func (o *Variant) SetSize(v []string) {
	o.Size = v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Variant) GetImages() []string {
	if o == nil || IsNil(o.Images) {
		var ret []string
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variant) GetImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Variant) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []string and assigns it to the Images field.
func (o *Variant) SetImages(v []string) {
	o.Images = v
}

func (o Variant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Variant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.PostedDate) {
		toSerialize["posted_date"] = o.PostedDate
	}
	if !IsNil(o.ColorId) {
		toSerialize["color_id"] = o.ColorId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableVariant struct {
	value *Variant
	isSet bool
}

func (v NullableVariant) Get() *Variant {
	return v.value
}

func (v *NullableVariant) Set(val *Variant) {
	v.value = val
	v.isSet = true
}

func (v NullableVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariant(val *Variant) *NullableVariant {
	return &NullableVariant{value: val, isSet: true}
}

func (v NullableVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


