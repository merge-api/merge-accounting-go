/*
 * Merge Accounting API
 *
 * The unified API for building rich integrations with multiple Accounting & Finance platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_accounting_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// PurchaseOrdersApiService PurchaseOrdersApi service
type PurchaseOrdersApiService service

type ApiPurchaseOrdersCreateRequest struct {
	ctx _context.Context
	ApiService *PurchaseOrdersApiService
	xAccountToken *string
	purchaseOrderEndpointRequest *PurchaseOrderEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiPurchaseOrdersCreateRequest) XAccountToken(xAccountToken string) ApiPurchaseOrdersCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiPurchaseOrdersCreateRequest) PurchaseOrderEndpointRequest(purchaseOrderEndpointRequest PurchaseOrderEndpointRequest) ApiPurchaseOrdersCreateRequest {
	r.purchaseOrderEndpointRequest = &purchaseOrderEndpointRequest
	return r
}
func (r ApiPurchaseOrdersCreateRequest) IsDebugMode(isDebugMode bool) ApiPurchaseOrdersCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiPurchaseOrdersCreateRequest) RunAsync(runAsync bool) ApiPurchaseOrdersCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiPurchaseOrdersCreateRequest) Execute() (PurchaseOrderResponse, *_nethttp.Response, error) {
	return r.ApiService.PurchaseOrdersCreateExecute(r)
}

/*
 * PurchaseOrdersCreate Method for PurchaseOrdersCreate
 * Creates a `PurchaseOrder` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPurchaseOrdersCreateRequest
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersCreate(ctx _context.Context) ApiPurchaseOrdersCreateRequest {
	return ApiPurchaseOrdersCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PurchaseOrderResponse
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersCreateExecute(r ApiPurchaseOrdersCreateRequest) (PurchaseOrderResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersApiService.PurchaseOrdersCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase-orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.purchaseOrderEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("purchaseOrderEndpointRequest is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	// body params
	localVarPostBody = r.purchaseOrderEndpointRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPurchaseOrdersListRequest struct {
	ctx _context.Context
	ApiService *PurchaseOrdersApiService
	xAccountToken *string
	companyId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteFields *string
	remoteId *string
	showEnumOrigins *string
}

func (r ApiPurchaseOrdersListRequest) XAccountToken(xAccountToken string) ApiPurchaseOrdersListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiPurchaseOrdersListRequest) CompanyId(companyId string) ApiPurchaseOrdersListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiPurchaseOrdersListRequest) CreatedAfter(createdAfter time.Time) ApiPurchaseOrdersListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiPurchaseOrdersListRequest) CreatedBefore(createdBefore time.Time) ApiPurchaseOrdersListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiPurchaseOrdersListRequest) Cursor(cursor string) ApiPurchaseOrdersListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiPurchaseOrdersListRequest) IncludeDeletedData(includeDeletedData bool) ApiPurchaseOrdersListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiPurchaseOrdersListRequest) IncludeRemoteData(includeRemoteData bool) ApiPurchaseOrdersListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiPurchaseOrdersListRequest) ModifiedAfter(modifiedAfter time.Time) ApiPurchaseOrdersListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiPurchaseOrdersListRequest) ModifiedBefore(modifiedBefore time.Time) ApiPurchaseOrdersListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiPurchaseOrdersListRequest) PageSize(pageSize int32) ApiPurchaseOrdersListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiPurchaseOrdersListRequest) RemoteFields(remoteFields string) ApiPurchaseOrdersListRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiPurchaseOrdersListRequest) RemoteId(remoteId string) ApiPurchaseOrdersListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiPurchaseOrdersListRequest) ShowEnumOrigins(showEnumOrigins string) ApiPurchaseOrdersListRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}

func (r ApiPurchaseOrdersListRequest) Execute() (PaginatedPurchaseOrderList, *_nethttp.Response, error) {
	return r.ApiService.PurchaseOrdersListExecute(r)
}

/*
 * PurchaseOrdersList Method for PurchaseOrdersList
 * Returns a list of `PurchaseOrder` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPurchaseOrdersListRequest
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersList(ctx _context.Context) ApiPurchaseOrdersListRequest {
	return ApiPurchaseOrdersListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedPurchaseOrderList
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersListExecute(r ApiPurchaseOrdersListRequest) (PaginatedPurchaseOrderList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedPurchaseOrderList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersApiService.PurchaseOrdersList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase-orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.companyId != nil {
		localVarQueryParams.Add("company_id", parameterToString(*r.companyId, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created_before", parameterToString(*r.createdBefore, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.remoteFields != nil {
		localVarQueryParams.Add("remote_fields", parameterToString(*r.remoteFields, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
	}
	if r.showEnumOrigins != nil {
		localVarQueryParams.Add("show_enum_origins", parameterToString(*r.showEnumOrigins, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPurchaseOrdersMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *PurchaseOrdersApiService
	xAccountToken *string
}

func (r ApiPurchaseOrdersMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiPurchaseOrdersMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiPurchaseOrdersMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.PurchaseOrdersMetaPostRetrieveExecute(r)
}

/*
 * PurchaseOrdersMetaPostRetrieve Method for PurchaseOrdersMetaPostRetrieve
 * Returns metadata for `PurchaseOrder` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPurchaseOrdersMetaPostRetrieveRequest
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersMetaPostRetrieve(ctx _context.Context) ApiPurchaseOrdersMetaPostRetrieveRequest {
	return ApiPurchaseOrdersMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersMetaPostRetrieveExecute(r ApiPurchaseOrdersMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersApiService.PurchaseOrdersMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase-orders/meta/post"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPurchaseOrdersRetrieveRequest struct {
	ctx _context.Context
	ApiService *PurchaseOrdersApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
	remoteFields *string
	showEnumOrigins *string
}

func (r ApiPurchaseOrdersRetrieveRequest) XAccountToken(xAccountToken string) ApiPurchaseOrdersRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiPurchaseOrdersRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiPurchaseOrdersRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiPurchaseOrdersRetrieveRequest) RemoteFields(remoteFields string) ApiPurchaseOrdersRetrieveRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiPurchaseOrdersRetrieveRequest) ShowEnumOrigins(showEnumOrigins string) ApiPurchaseOrdersRetrieveRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}

func (r ApiPurchaseOrdersRetrieveRequest) Execute() (PurchaseOrder, *_nethttp.Response, error) {
	return r.ApiService.PurchaseOrdersRetrieveExecute(r)
}

/*
 * PurchaseOrdersRetrieve Method for PurchaseOrdersRetrieve
 * Returns a `PurchaseOrder` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPurchaseOrdersRetrieveRequest
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersRetrieve(ctx _context.Context, id string) ApiPurchaseOrdersRetrieveRequest {
	return ApiPurchaseOrdersRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return PurchaseOrder
 */
func (a *PurchaseOrdersApiService) PurchaseOrdersRetrieveExecute(r ApiPurchaseOrdersRetrieveRequest) (PurchaseOrder, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersApiService.PurchaseOrdersRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase-orders/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.remoteFields != nil {
		localVarQueryParams.Add("remote_fields", parameterToString(*r.remoteFields, ""))
	}
	if r.showEnumOrigins != nil {
		localVarQueryParams.Add("show_enum_origins", parameterToString(*r.showEnumOrigins, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
