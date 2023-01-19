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

// ItemsApiService ItemsApi service
type ItemsApiService service

type ApiItemsListRequest struct {
	ctx _context.Context
	ApiService *ItemsApiService
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

func (r ApiItemsListRequest) XAccountToken(xAccountToken string) ApiItemsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiItemsListRequest) CompanyId(companyId string) ApiItemsListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiItemsListRequest) CreatedAfter(createdAfter time.Time) ApiItemsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiItemsListRequest) CreatedBefore(createdBefore time.Time) ApiItemsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiItemsListRequest) Cursor(cursor string) ApiItemsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiItemsListRequest) IncludeDeletedData(includeDeletedData bool) ApiItemsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiItemsListRequest) IncludeRemoteData(includeRemoteData bool) ApiItemsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiItemsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiItemsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiItemsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiItemsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiItemsListRequest) PageSize(pageSize int32) ApiItemsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiItemsListRequest) RemoteFields(remoteFields string) ApiItemsListRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiItemsListRequest) RemoteId(remoteId string) ApiItemsListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiItemsListRequest) ShowEnumOrigins(showEnumOrigins string) ApiItemsListRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}

func (r ApiItemsListRequest) Execute() (PaginatedItemList, *_nethttp.Response, error) {
	return r.ApiService.ItemsListExecute(r)
}

/*
 * ItemsList Method for ItemsList
 * Returns a list of `Item` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiItemsListRequest
 */
func (a *ItemsApiService) ItemsList(ctx _context.Context) ApiItemsListRequest {
	return ApiItemsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedItemList
 */
func (a *ItemsApiService) ItemsListExecute(r ApiItemsListRequest) (PaginatedItemList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedItemList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemsApiService.ItemsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/items"

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

type ApiItemsRetrieveRequest struct {
	ctx _context.Context
	ApiService *ItemsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
	remoteFields *string
	showEnumOrigins *string
}

func (r ApiItemsRetrieveRequest) XAccountToken(xAccountToken string) ApiItemsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiItemsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiItemsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiItemsRetrieveRequest) RemoteFields(remoteFields string) ApiItemsRetrieveRequest {
	r.remoteFields = &remoteFields
	return r
}
func (r ApiItemsRetrieveRequest) ShowEnumOrigins(showEnumOrigins string) ApiItemsRetrieveRequest {
	r.showEnumOrigins = &showEnumOrigins
	return r
}

func (r ApiItemsRetrieveRequest) Execute() (Item, *_nethttp.Response, error) {
	return r.ApiService.ItemsRetrieveExecute(r)
}

/*
 * ItemsRetrieve Method for ItemsRetrieve
 * Returns an `Item` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiItemsRetrieveRequest
 */
func (a *ItemsApiService) ItemsRetrieve(ctx _context.Context, id string) ApiItemsRetrieveRequest {
	return ApiItemsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Item
 */
func (a *ItemsApiService) ItemsRetrieveExecute(r ApiItemsRetrieveRequest) (Item, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Item
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemsApiService.ItemsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/items/{id}"
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