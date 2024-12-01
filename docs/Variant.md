# Variant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**PostedDate** | Pointer to **time.Time** |  | [optional] 
**ColorId** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **[]string** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVariant

`func NewVariant() *Variant`

NewVariant instantiates a new Variant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithDefaults

`func NewVariantWithDefaults() *Variant`

NewVariantWithDefaults instantiates a new Variant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Variant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Variant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *Variant) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Variant) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Variant) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Variant) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Variant) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Variant) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Variant) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Variant) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPostedDate

`func (o *Variant) GetPostedDate() time.Time`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *Variant) GetPostedDateOk() (*time.Time, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *Variant) SetPostedDate(v time.Time)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *Variant) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetColorId

`func (o *Variant) GetColorId() int32`

GetColorId returns the ColorId field if non-nil, zero value otherwise.

### GetColorIdOk

`func (o *Variant) GetColorIdOk() (*int32, bool)`

GetColorIdOk returns a tuple with the ColorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorId

`func (o *Variant) SetColorId(v int32)`

SetColorId sets ColorId field to given value.

### HasColorId

`func (o *Variant) HasColorId() bool`

HasColorId returns a boolean if a field has been set.

### GetSize

`func (o *Variant) GetSize() []string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Variant) GetSizeOk() (*[]string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Variant) SetSize(v []string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Variant) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetImages

`func (o *Variant) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Variant) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Variant) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *Variant) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


