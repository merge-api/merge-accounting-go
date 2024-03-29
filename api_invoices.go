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

// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type ApiInvoicesCreateRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	xAccountToken *string
	invoiceEndpointRequest *InvoiceEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiInvoicesCreateRequest) XAccountToken(xAccountToken string) ApiInvoicesCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiInvoicesCreateRequest) InvoiceEndpointRequest(invoiceEndpointRequest InvoiceEndpointRequest) ApiInvoicesCreateRequest {
	r.invoiceEndpointRequest = &invoiceEndpointRequest
	return r
}
func (r ApiInvoicesCreateRequest) IsDebugMode(isDebugMode bool) ApiInvoicesCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiInvoicesCreateRequest) RunAsync(runAsync bool) ApiInvoicesCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiInvoicesCreateRequest) Execute() (InvoiceResponse, *_nethttp.Response, error) {
	return r.ApiService.InvoicesCreateExecute(r)
}

/*
 * InvoicesCreate Method for InvoicesCreate
 * Creates an `Invoice` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInvoicesCreateRequest
 */
func (a *InvoicesApiService) InvoicesCreate(ctx _context.Context) ApiInvoicesCreateRequest {
	return ApiInvoicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InvoiceResponse
 */
func (a *InvoicesApiService) InvoicesCreateExecute(r ApiInvoicesCreateRequest) (InvoiceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.InvoicesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.invoiceEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("invoiceEndpointRequest is required and must be specified")
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
	localVarPostBody = r.invoiceEndpointRequest
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

type ApiInvoicesListRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	xAccountToken *string
	companyId *string
	contactId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	issueDateAfter *time.Time
	issueDateBefore *time.Time
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteFields *string
	remoteId *string
	showEnumOrigins *string
	type_ *string
}

func (r ApiInvoicesListRequest) XAccountToken(xAccountToken string) ApiInvoicesListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiInvoicesListRequest) CompanyId(companyId string) ApiInvoicesListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiInvoicesListRequest) ContactId(contactId string) ApiInvoicesListRequest {
	r.contactId = &contactId
	return r
}
func (r ApiInvoicesListRequest) CreatedAfter(createdAfter time.Time) ApiInvoicesListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiInvoicesListRequest) CreatedBefore(createdBefore time.Time) ApiInvoicesListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiInvoicesListRequest) Cursor(cursor string) ApiInvoicesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiInvoicesListRequest) IncludeDeletedData(includeDeletedData bool) ApiInvoicesListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiInvoicesListRequest) IncludeRemoteData(includeRemoteData bool) ApiInvoicesListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiInvoicesListRequest) IssueDateAfter(issueDateAfter time.Time) ApiInvoicesListRequest {
	r.issueDateAfter = &issueDateAfter
	return r
}
func (r ApiInvoicesListRequest) IssueDateBefore(issueDateBefore time.Time) ApiInvoicesListRequest {
	r.issueDateBefore = &issueDateBefore
	return r
}
func (r ApiInvoicesListRequest) ModifiedAfter(modifiedAfter time.Time) ApiInvoicesListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiInvoicesListRequest) ModifiedBefore(modifiedBefore time.Time) ApiInvoicesListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiInvoicesListRequest) PageSize(pageSize int32) ApiInvoicesListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiInvoicesListRequest) RemoteFields(remoteFields string) ApiInvoicesListRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiInvoicesListRequest) RemoteId(remoteId string) ApiInvoicesListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiInvoicesListRequest) ShowEnumOrigins(showEnumOrigins string) ApiInvoicesListRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}
func (r ApiInvoicesListRequest) Type_(type_ string) ApiInvoicesListRequest {
	r.type_ = &type_
	return r
}

func (r ApiInvoicesListRequest) Execute() (PaginatedInvoiceList, *_nethttp.Response, error) {
	return r.ApiService.InvoicesListExecute(r)
}

/*
 * InvoicesList Method for InvoicesList
 * Returns a list of `Invoice` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInvoicesListRequest
 */
func (a *InvoicesApiService) InvoicesList(ctx _context.Context) ApiInvoicesListRequest {
	return ApiInvoicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedInvoiceList
 */
func (a *InvoicesApiService) InvoicesListExecute(r ApiInvoicesListRequest) (PaginatedInvoiceList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedInvoiceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.InvoicesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.companyId != nil {
		localVarQueryParams.Add("company_id", parameterToString(*r.companyId, ""))
	}
	if r.contactId != nil {
		localVarQueryParams.Add("contact_id", parameterToString(*r.contactId, ""))
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
	if r.issueDateAfter != nil {
		localVarQueryParams.Add("issue_date_after", parameterToString(*r.issueDateAfter, ""))
	}
	if r.issueDateBefore != nil {
		localVarQueryParams.Add("issue_date_before", parameterToString(*r.issueDateBefore, ""))
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
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
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

type ApiInvoicesMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	xAccountToken *string
}

func (r ApiInvoicesMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiInvoicesMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiInvoicesMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.InvoicesMetaPostRetrieveExecute(r)
}

/*
 * InvoicesMetaPostRetrieve Method for InvoicesMetaPostRetrieve
 * Returns metadata for `Invoice` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInvoicesMetaPostRetrieveRequest
 */
func (a *InvoicesApiService) InvoicesMetaPostRetrieve(ctx _context.Context) ApiInvoicesMetaPostRetrieveRequest {
	return ApiInvoicesMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *InvoicesApiService) InvoicesMetaPostRetrieveExecute(r ApiInvoicesMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.InvoicesMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/meta/post"

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

type ApiInvoicesRetrieveRequest struct {
	ctx _context.Context
	ApiService *InvoicesApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
	remoteFields *string
	showEnumOrigins *string
}

func (r ApiInvoicesRetrieveRequest) XAccountToken(xAccountToken string) ApiInvoicesRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiInvoicesRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiInvoicesRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiInvoicesRetrieveRequest) RemoteFields(remoteFields string) ApiInvoicesRetrieveRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiInvoicesRetrieveRequest) ShowEnumOrigins(showEnumOrigins string) ApiInvoicesRetrieveRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}

func (r ApiInvoicesRetrieveRequest) Execute() (Invoice, *_nethttp.Response, error) {
	return r.ApiService.InvoicesRetrieveExecute(r)
}

/*
 * InvoicesRetrieve Method for InvoicesRetrieve
 * Returns an `Invoice` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiInvoicesRetrieveRequest
 */
func (a *InvoicesApiService) InvoicesRetrieve(ctx _context.Context, id string) ApiInvoicesRetrieveRequest {
	return ApiInvoicesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Invoice
 */
func (a *InvoicesApiService) InvoicesRetrieveExecute(r ApiInvoicesRetrieveRequest) (Invoice, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Invoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.InvoicesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}"
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
