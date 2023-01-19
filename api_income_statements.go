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

// IncomeStatementsApiService IncomeStatementsApi service
type IncomeStatementsApiService service

type ApiIncomeStatementsListRequest struct {
	ctx _context.Context
	ApiService *IncomeStatementsApiService
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
	remoteId *string
}

func (r ApiIncomeStatementsListRequest) XAccountToken(xAccountToken string) ApiIncomeStatementsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiIncomeStatementsListRequest) CompanyId(companyId string) ApiIncomeStatementsListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiIncomeStatementsListRequest) CreatedAfter(createdAfter time.Time) ApiIncomeStatementsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiIncomeStatementsListRequest) CreatedBefore(createdBefore time.Time) ApiIncomeStatementsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiIncomeStatementsListRequest) Cursor(cursor string) ApiIncomeStatementsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiIncomeStatementsListRequest) IncludeDeletedData(includeDeletedData bool) ApiIncomeStatementsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiIncomeStatementsListRequest) IncludeRemoteData(includeRemoteData bool) ApiIncomeStatementsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiIncomeStatementsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiIncomeStatementsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiIncomeStatementsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiIncomeStatementsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiIncomeStatementsListRequest) PageSize(pageSize int32) ApiIncomeStatementsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiIncomeStatementsListRequest) RemoteId(remoteId string) ApiIncomeStatementsListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiIncomeStatementsListRequest) Execute() (PaginatedIncomeStatementList, *_nethttp.Response, error) {
	return r.ApiService.IncomeStatementsListExecute(r)
}

/*
 * IncomeStatementsList Method for IncomeStatementsList
 * Returns a list of `IncomeStatement` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIncomeStatementsListRequest
 */
func (a *IncomeStatementsApiService) IncomeStatementsList(ctx _context.Context) ApiIncomeStatementsListRequest {
	return ApiIncomeStatementsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedIncomeStatementList
 */
func (a *IncomeStatementsApiService) IncomeStatementsListExecute(r ApiIncomeStatementsListRequest) (PaginatedIncomeStatementList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedIncomeStatementList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncomeStatementsApiService.IncomeStatementsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income-statements"

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
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
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

type ApiIncomeStatementsRetrieveRequest struct {
	ctx _context.Context
	ApiService *IncomeStatementsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiIncomeStatementsRetrieveRequest) XAccountToken(xAccountToken string) ApiIncomeStatementsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiIncomeStatementsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiIncomeStatementsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiIncomeStatementsRetrieveRequest) Execute() (IncomeStatement, *_nethttp.Response, error) {
	return r.ApiService.IncomeStatementsRetrieveExecute(r)
}

/*
 * IncomeStatementsRetrieve Method for IncomeStatementsRetrieve
 * Returns an `IncomeStatement` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIncomeStatementsRetrieveRequest
 */
func (a *IncomeStatementsApiService) IncomeStatementsRetrieve(ctx _context.Context, id string) ApiIncomeStatementsRetrieveRequest {
	return ApiIncomeStatementsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return IncomeStatement
 */
func (a *IncomeStatementsApiService) IncomeStatementsRetrieveExecute(r ApiIncomeStatementsRetrieveRequest) (IncomeStatement, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IncomeStatement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncomeStatementsApiService.IncomeStatementsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income-statements/{id}"
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
