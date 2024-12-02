# \DefaultApi

All URIs are relative to *https://uuza.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAddSimpleProductPost**](DefaultApi.md#ApiAddSimpleProductPost) | **Post** /api/add-simple-product | Add a simple product
[**ApiAddVariantProductPost**](DefaultApi.md#ApiAddVariantProductPost) | **Post** /api/add-variant-product | Add a variant product
[**ApiProductProductIdGet**](DefaultApi.md#ApiProductProductIdGet) | **Get** /api/product/{product_id} | Get a product
[**ApiProductsGet**](DefaultApi.md#ApiProductsGet) | **Get** /api/products | Get all products
[**ApiUpdateSimpleProductProductIdPost**](DefaultApi.md#ApiUpdateSimpleProductProductIdPost) | **Post** /api/update-simple-product/{product_id} | Update a simple product
[**ApiUpdateVariantProductProductIdPost**](DefaultApi.md#ApiUpdateVariantProductProductIdPost) | **Post** /api/update-variant-product/{product_id} | Update a variant product



## ApiAddSimpleProductPost

> ApiAddSimpleProductPost200Response ApiAddSimpleProductPost(ctx).SimpleProduct(simpleProduct).Execute()

Add a simple product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    simpleProduct := *openapiclient.NewSimpleProduct() // SimpleProduct | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiAddSimpleProductPost(context.Background()).SimpleProduct(simpleProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAddSimpleProductPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAddSimpleProductPost`: ApiAddSimpleProductPost200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAddSimpleProductPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAddSimpleProductPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simpleProduct** | [**SimpleProduct**](SimpleProduct.md) |  | 

### Return type

[**ApiAddSimpleProductPost200Response**](ApiAddSimpleProductPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAddVariantProductPost

> ApiAddSimpleProductPost200Response ApiAddVariantProductPost(ctx).VariantProduct(variantProduct).Execute()

Add a variant product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    variantProduct := *openapiclient.NewVariantProduct() // VariantProduct | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiAddVariantProductPost(context.Background()).VariantProduct(variantProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAddVariantProductPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAddVariantProductPost`: ApiAddSimpleProductPost200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAddVariantProductPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAddVariantProductPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variantProduct** | [**VariantProduct**](VariantProduct.md) |  | 

### Return type

[**ApiAddSimpleProductPost200Response**](ApiAddSimpleProductPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductProductIdGet

> ApiProductProductIdGet200Response ApiProductProductIdGet(ctx, productId).Execute()

Get a product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    productId := 356 // uint64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiProductProductIdGet(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiProductProductIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductProductIdGet`: ApiProductProductIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiProductProductIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProductProductIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiProductProductIdGet200Response**](ApiProductProductIdGet200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductsGet

> ApiProductsGet200Response ApiProductsGet(ctx).PerPage(perPage).Page(page).Search(search).Execute()

Get all products



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiProductsGet(context.Background()).PerPage(perPage).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiProductsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsGet`: ApiProductsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiProductsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**ApiProductsGet200Response**](ApiProductsGet200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUpdateSimpleProductProductIdPost

> ApiUpdateSimpleProductProductIdPost(ctx, productId).SimpleProduct(simpleProduct).Execute()

Update a simple product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    productId := 358 // uint64 | 
    simpleProduct := *openapiclient.NewSimpleProduct() // SimpleProduct | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ApiUpdateSimpleProductProductIdPost(context.Background(), productId).SimpleProduct(simpleProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiUpdateSimpleProductProductIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUpdateSimpleProductProductIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simpleProduct** | [**SimpleProduct**](SimpleProduct.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUpdateVariantProductProductIdPost

> ApiUpdateVariantProductProductIdPost(ctx, productId).VariantProduct(variantProduct).Execute()

Update a variant product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    productId :=  245// string | 
    variantProduct := *openapiclient.NewVariantProduct() // VariantProduct | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ApiUpdateVariantProductProductIdPost(context.Background(), productId).VariantProduct(variantProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiUpdateVariantProductProductIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUpdateVariantProductProductIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variantProduct** | [**VariantProduct**](VariantProduct.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

