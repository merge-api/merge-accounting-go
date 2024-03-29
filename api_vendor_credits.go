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

// VendorCreditsApiService VendorCreditsApi service
type VendorCreditsApiService service

type ApiVendorCreditsListRequest struct {
	ctx _context.Context
	ApiService *VendorCreditsApiService
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
	transactionDateAfter *time.Time
	transactionDateBefore *time.Time
}

func (r ApiVendorCreditsListRequest) XAccountToken(xAccountToken string) ApiVendorCreditsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiVendorCreditsListRequest) CompanyId(companyId string) ApiVendorCreditsListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiVendorCreditsListRequest) CreatedAfter(createdAfter time.Time) ApiVendorCreditsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiVendorCreditsListRequest) CreatedBefore(createdBefore time.Time) ApiVendorCreditsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiVendorCreditsListRequest) Cursor(cursor string) ApiVendorCreditsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiVendorCreditsListRequest) IncludeDeletedData(includeDeletedData bool) ApiVendorCreditsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiVendorCreditsListRequest) IncludeRemoteData(includeRemoteData bool) ApiVendorCreditsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiVendorCreditsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiVendorCreditsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiVendorCreditsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiVendorCreditsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiVendorCreditsListRequest) PageSize(pageSize int32) ApiVendorCreditsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiVendorCreditsListRequest) RemoteId(remoteId string) ApiVendorCreditsListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiVendorCreditsListRequest) TransactionDateAfter(transactionDateAfter time.Time) ApiVendorCreditsListRequest {
	r.transactionDateAfter = &transactionDateAfter
	return r
}
func (r ApiVendorCreditsListRequest) TransactionDateBefore(transactionDateBefore time.Time) ApiVendorCreditsListRequest {
	r.transactionDateBefore = &transactionDateBefore
	return r
}

func (r ApiVendorCreditsListRequest) Execute() (PaginatedVendorCreditList, *_nethttp.Response, error) {
	return r.ApiService.VendorCreditsListExecute(r)
}

/*
 * VendorCreditsList Method for VendorCreditsList
 * Returns a list of `VendorCredit` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiVendorCreditsListRequest
 */
func (a *VendorCreditsApiService) VendorCreditsList(ctx _context.Context) ApiVendorCreditsListRequest {
	return ApiVendorCreditsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedVendorCreditList
 */
func (a *VendorCreditsApiService) VendorCreditsListExecute(r ApiVendorCreditsListRequest) (PaginatedVendorCreditList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedVendorCreditList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorCreditsApiService.VendorCreditsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor-credits"

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
	if r.transactionDateAfter != nil {
		localVarQueryParams.Add("transaction_date_after", parameterToString(*r.transactionDateAfter, ""))
	}
	if r.transactionDateBefore != nil {
		localVarQueryParams.Add("transaction_date_before", parameterToString(*r.transactionDateBefore, ""))
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

type ApiVendorCreditsRetrieveRequest struct {
	ctx _context.Context
	ApiService *VendorCreditsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiVendorCreditsRetrieveRequest) XAccountToken(xAccountToken string) ApiVendorCreditsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiVendorCreditsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiVendorCreditsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiVendorCreditsRetrieveRequest) Execute() (VendorCredit, *_nethttp.Response, error) {
	return r.ApiService.VendorCreditsRetrieveExecute(r)
}

/*
 * VendorCreditsRetrieve Method for VendorCreditsRetrieve
 * Returns a `VendorCredit` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiVendorCreditsRetrieveRequest
 */
func (a *VendorCreditsApiService) VendorCreditsRetrieve(ctx _context.Context, id string) ApiVendorCreditsRetrieveRequest {
	return ApiVendorCreditsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return VendorCredit
 */
func (a *VendorCreditsApiService) VendorCreditsRetrieveExecute(r ApiVendorCreditsRetrieveRequest) (VendorCredit, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VendorCredit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorCreditsApiService.VendorCreditsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor-credits/{id}"
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
