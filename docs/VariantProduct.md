# VariantProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FeatureImage** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Variants** | Pointer to [**[]Variant**](Variant.md) |  | [optional] 

## Methods

### NewVariantProduct

`func NewVariantProduct() *VariantProduct`

NewVariantProduct instantiates a new VariantProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantProductWithDefaults

`func NewVariantProductWithDefaults() *VariantProduct`

NewVariantProductWithDefaults instantiates a new VariantProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariantProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariantProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariantProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariantProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryId

`func (o *VariantProduct) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *VariantProduct) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *VariantProduct) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *VariantProduct) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *VariantProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariantProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariantProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariantProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureImage

`func (o *VariantProduct) GetFeatureImage() string`

GetFeatureImage returns the FeatureImage field if non-nil, zero value otherwise.

### GetFeatureImageOk

`func (o *VariantProduct) GetFeatureImageOk() (*string, bool)`

GetFeatureImageOk returns a tuple with the FeatureImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureImage

`func (o *VariantProduct) SetFeatureImage(v string)`

SetFeatureImage sets FeatureImage field to given value.

### HasFeatureImage

`func (o *VariantProduct) HasFeatureImage() bool`

HasFeatureImage returns a boolean if a field has been set.

### GetTag

`func (o *VariantProduct) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *VariantProduct) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *VariantProduct) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *VariantProduct) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVariants

`func (o *VariantProduct) GetVariants() []Variant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *VariantProduct) GetVariantsOk() (*[]Variant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *VariantProduct) SetVariants(v []Variant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *VariantProduct) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


