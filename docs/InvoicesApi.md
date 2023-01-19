# \InvoicesApi

All URIs are relative to *https://api.merge.dev/api/accounting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicesCreate**](InvoicesApi.md#InvoicesCreate) | **Post** /invoices | 
[**InvoicesList**](InvoicesApi.md#InvoicesList) | **Get** /invoices | 
[**InvoicesMetaPostRetrieve**](InvoicesApi.md#InvoicesMetaPostRetrieve) | **Get** /invoices/meta/post | 
[**InvoicesRetrieve**](InvoicesApi.md#InvoicesRetrieve) | **Get** /invoices/{id} | 



## InvoicesCreate

> InvoiceResponse InvoicesCreate(ctx).XAccountToken(xAccountToken).InvoiceEndpointRequest(invoiceEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    invoiceEndpointRequest := *openapiclient.NewInvoiceEndpointRequest(*openapiclient.NewInvoiceRequest()) // InvoiceEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.InvoicesCreate(context.Background()).XAccountToken(xAccountToken).InvoiceEndpointRequest(invoiceEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesCreate`: InvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **invoiceEndpointRequest** | [**InvoiceEndpointRequest**](InvoiceEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesList

> PaginatedInvoiceList InvoicesList(ctx).XAccountToken(xAccountToken).CompanyId(companyId).ContactId(contactId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).Type_(type_).Execute()





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
    companyId := "companyId_example" // string | If provided, will only return invoices for this company. (optional)
    contactId := "contactId_example" // string | If provided, will only return invoices for this contact. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteFields := "type" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    showEnumOrigins := "type" // string | Which fields should be returned in non-normalized form. (optional)
    type_ := "type__example" // string | If provided, will only return Invoices with this type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.InvoicesList(context.Background()).XAccountToken(xAccountToken).CompanyId(companyId).ContactId(contactId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteFields(remoteFields).RemoteId(remoteId).ShowEnumOrigins(showEnumOrigins).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesList`: PaginatedInvoiceList
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **companyId** | **string** | If provided, will only return invoices for this company. | 
 **contactId** | **string** | If provided, will only return invoices for this contact. | 
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
 **type_** | **string** | If provided, will only return Invoices with this type | 

### Return type

[**PaginatedInvoiceList**](PaginatedInvoiceList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesMetaPostRetrieve

> MetaResponse InvoicesMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.InvoicesApi.InvoicesMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesMetaPostRetrieveRequest struct via the builder pattern


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


## InvoicesRetrieve

> Invoice InvoicesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





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
    remoteFields := "type" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "type" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.InvoicesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoicesRetrieve`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

