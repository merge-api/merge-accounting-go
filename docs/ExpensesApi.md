# \ExpensesApi

All URIs are relative to *https://api.merge.dev/api/accounting/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpensesCreate**](ExpensesApi.md#ExpensesCreate) | **Post** /expenses | 
[**ExpensesList**](ExpensesApi.md#ExpensesList) | **Get** /expenses | 
[**ExpensesMetaPostRetrieve**](ExpensesApi.md#ExpensesMetaPostRetrieve) | **Get** /expenses/meta/post | 
[**ExpensesRetrieve**](ExpensesApi.md#ExpensesRetrieve) | **Get** /expenses/{id} | 



## ExpensesCreate

> ExpenseResponse ExpensesCreate(ctx).XAccountToken(xAccountToken).ExpenseEndpointRequest(expenseEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    expenseEndpointRequest := *openapiclient.NewExpenseEndpointRequest(*openapiclient.NewExpenseRequest()) // ExpenseEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpensesApi.ExpensesCreate(context.Background()).XAccountToken(xAccountToken).ExpenseEndpointRequest(expenseEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.ExpensesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpensesCreate`: ExpenseResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.ExpensesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpensesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **expenseEndpointRequest** | [**ExpenseEndpointRequest**](ExpenseEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**ExpenseResponse**](ExpenseResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpensesList

> PaginatedExpenseList ExpensesList(ctx).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





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
    companyId := "companyId_example" // string | If provided, will only return expenses for this company. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpensesApi.ExpensesList(context.Background()).XAccountToken(xAccountToken).CompanyId(companyId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.ExpensesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpensesList`: PaginatedExpenseList
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.ExpensesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpensesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **companyId** | **string** | If provided, will only return expenses for this company. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedExpenseList**](PaginatedExpenseList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpensesMetaPostRetrieve

> MetaResponse ExpensesMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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
    resp, r, err := api_client.ExpensesApi.ExpensesMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.ExpensesMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpensesMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.ExpensesMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpensesMetaPostRetrieveRequest struct via the builder pattern


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


## ExpensesRetrieve

> Expense ExpensesRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpensesApi.ExpensesRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.ExpensesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpensesRetrieve`: Expense
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.ExpensesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpensesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 

### Return type

[**Expense**](Expense.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

