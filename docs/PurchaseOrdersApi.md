# \PurchaseOrdersApi

All URIs are relative to *https://api.merge.dev/api/accounting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseOrdersCreate**](PurchaseOrdersApi.md#PurchaseOrdersCreate) | **Post** /purchase-orders | 
[**PurchaseOrdersList**](PurchaseOrdersApi.md#PurchaseOrdersList) | **Get** /purchase-orders | 
[**PurchaseOrdersMetaPostRetrieve**](PurchaseOrdersApi.md#PurchaseOrdersMetaPostRetrieve) | **Get** /purchase-orders/meta/post | 
[**PurchaseOrdersRetrieve**](PurchaseOrdersApi.md#PurchaseOrdersRetrieve) | **Get** /purchase-orders/{id} | 



## PurchaseOrdersCreate

> PurchaseOrderResponse PurchaseOrdersCreate(ctx).XAccountToken(xAccountToken).PurchaseOrderEndpointRequest(purchaseOrderEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    purchaseOrderEndpointRequest := *openapiclient.NewPurchaseOrderEndpointRequest(*openapiclient.NewPurchaseOrderRequest()) // PurchaseOrderEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseOrdersApi.PurchaseOrdersCreate(context.Background()).XAccountToken(xAccountToken).PurchaseOrderEndpointRequest(purchaseOrderEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersApi.PurchaseOrdersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseOrdersCreate`: PurchaseOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersApi.PurchaseOrdersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseOrdersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **purchaseOrderEndpointRequest** | [**PurchaseOrderEndpointRequest**](PurchaseOrderEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**PurchaseOrderResponse**](PurchaseOrderResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseOrdersList

> PaginatedPurchaseOrderList PurchaseOrdersList(ctx).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    companyId := "companyId_example" // string | If provided, will only return purchase orders for this company. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteFields := "status" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    showEnumOrigins := "status" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseOrdersApi.PurchaseOrdersList(context.Background()).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersApi.PurchaseOrdersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseOrdersList`: PaginatedPurchaseOrderList
    fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersApi.PurchaseOrdersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseOrdersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **companyId** | **string** | If provided, will only return purchase orders for this company. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**PaginatedPurchaseOrderList**](PaginatedPurchaseOrderList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseOrdersMetaPostRetrieve

> MetaResponse PurchaseOrdersMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseOrdersApi.PurchaseOrdersMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersApi.PurchaseOrdersMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseOrdersMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersApi.PurchaseOrdersMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseOrdersMetaPostRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseOrdersRetrieve

> PurchaseOrder PurchaseOrdersRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    id := TODO // string | 
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    remoteFields := "status" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "status" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseOrdersApi.PurchaseOrdersRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersApi.PurchaseOrdersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseOrdersRetrieve`: PurchaseOrder
    fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersApi.PurchaseOrdersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseOrdersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**PurchaseOrder**](PurchaseOrder.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

