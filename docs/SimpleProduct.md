# SimpleProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**PostedDate** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**FeatureImage** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewSimpleProduct

`func NewSimpleProduct() *SimpleProduct`

NewSimpleProduct instantiates a new SimpleProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleProductWithDefaults

`func NewSimpleProductWithDefaults() *SimpleProduct`

NewSimpleProductWithDefaults instantiates a new SimpleProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SimpleProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *SimpleProduct) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SimpleProduct) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SimpleProduct) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SimpleProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *SimpleProduct) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SimpleProduct) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SimpleProduct) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SimpleProduct) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCategoryId

`func (o *SimpleProduct) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *SimpleProduct) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *SimpleProduct) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *SimpleProduct) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPostedDate

`func (o *SimpleProduct) GetPostedDate() time.Time`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *SimpleProduct) GetPostedDateOk() (*time.Time, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *SimpleProduct) SetPostedDate(v time.Time)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *SimpleProduct) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImages

`func (o *SimpleProduct) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SimpleProduct) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SimpleProduct) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *SimpleProduct) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetFeatureImage

`func (o *SimpleProduct) GetFeatureImage() string`

GetFeatureImage returns the FeatureImage field if non-nil, zero value otherwise.

### GetFeatureImageOk

`func (o *SimpleProduct) GetFeatureImageOk() (*string, bool)`

GetFeatureImageOk returns a tuple with the FeatureImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureImage

`func (o *SimpleProduct) SetFeatureImage(v string)`

SetFeatureImage sets FeatureImage field to given value.

### HasFeatureImage

`func (o *SimpleProduct) HasFeatureImage() bool`

HasFeatureImage returns a boolean if a field has been set.

### GetTag

`func (o *SimpleProduct) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SimpleProduct) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SimpleProduct) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SimpleProduct) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


