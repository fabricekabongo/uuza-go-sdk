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

// checks if the VariantProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantProduct{}

// VariantProduct struct for VariantProduct
type VariantProduct struct {
	Name *string `json:"name,omitempty"`
	CategoryId *int32 `json:"category_id,omitempty"`
	Description *string `json:"description,omitempty"`
	FeatureImage *string `json:"feature_image,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Variants []Variant `json:"variants,omitempty"`
}

// NewVariantProduct instantiates a new VariantProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantProduct() *VariantProduct {
	this := VariantProduct{}
	return &this
}

// NewVariantProductWithDefaults instantiates a new VariantProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantProductWithDefaults() *VariantProduct {
	this := VariantProduct{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VariantProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VariantProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VariantProduct) SetName(v string) {
	o.Name = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *VariantProduct) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *VariantProduct) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *VariantProduct) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariantProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariantProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariantProduct) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureImage returns the FeatureImage field value if set, zero value otherwise.
func (o *VariantProduct) GetFeatureImage() string {
	if o == nil || IsNil(o.FeatureImage) {
		var ret string
		return ret
	}
	return *o.FeatureImage
}

// GetFeatureImageOk returns a tuple with the FeatureImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetFeatureImageOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureImage) {
		return nil, false
	}
	return o.FeatureImage, true
}

// HasFeatureImage returns a boolean if a field has been set.
func (o *VariantProduct) HasFeatureImage() bool {
	if o != nil && !IsNil(o.FeatureImage) {
		return true
	}

	return false
}

// SetFeatureImage gets a reference to the given string and assigns it to the FeatureImage field.
func (o *VariantProduct) SetFeatureImage(v string) {
	o.FeatureImage = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *VariantProduct) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *VariantProduct) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *VariantProduct) SetTag(v string) {
	o.Tag = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *VariantProduct) GetVariants() []Variant {
	if o == nil || IsNil(o.Variants) {
		var ret []Variant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantProduct) GetVariantsOk() ([]Variant, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *VariantProduct) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []Variant and assigns it to the Variants field.
func (o *VariantProduct) SetVariants(v []Variant) {
	o.Variants = v
}

func (o VariantProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FeatureImage) {
		toSerialize["feature_image"] = o.FeatureImage
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return toSerialize, nil
}

type NullableVariantProduct struct {
	value *VariantProduct
	isSet bool
}

func (v NullableVariantProduct) Get() *VariantProduct {
	return v.value
}

func (v *NullableVariantProduct) Set(val *VariantProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantProduct(val *VariantProduct) *NullableVariantProduct {
	return &NullableVariantProduct{value: val, isSet: true}
}

func (v NullableVariantProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

