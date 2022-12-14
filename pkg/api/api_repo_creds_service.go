/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RepoCredsServiceApiService RepoCredsServiceApi service
type RepoCredsServiceApiService service

type ApiRepoCredsServiceCreateRepositoryCredentialsRequest struct {
	ctx        context.Context
	ApiService *RepoCredsServiceApiService
	body       *V1alpha1RepoCreds
	upsert     *bool
}

// Repository definition
func (r ApiRepoCredsServiceCreateRepositoryCredentialsRequest) Body(body V1alpha1RepoCreds) ApiRepoCredsServiceCreateRepositoryCredentialsRequest {
	r.body = &body
	return r
}

// Whether to create in upsert mode.
func (r ApiRepoCredsServiceCreateRepositoryCredentialsRequest) Upsert(upsert bool) ApiRepoCredsServiceCreateRepositoryCredentialsRequest {
	r.upsert = &upsert
	return r
}

func (r ApiRepoCredsServiceCreateRepositoryCredentialsRequest) Execute() (*V1alpha1RepoCreds, *http.Response, error) {
	return r.ApiService.RepoCredsServiceCreateRepositoryCredentialsExecute(r)
}

/*
RepoCredsServiceCreateRepositoryCredentials CreateRepositoryCredentials creates a new repository credential set

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRepoCredsServiceCreateRepositoryCredentialsRequest
*/
func (a *RepoCredsServiceApiService) RepoCredsServiceCreateRepositoryCredentials(ctx context.Context) ApiRepoCredsServiceCreateRepositoryCredentialsRequest {
	return ApiRepoCredsServiceCreateRepositoryCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepoCreds
func (a *RepoCredsServiceApiService) RepoCredsServiceCreateRepositoryCredentialsExecute(r ApiRepoCredsServiceCreateRepositoryCredentialsRequest) (*V1alpha1RepoCreds, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepoCreds
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepoCredsServiceApiService.RepoCredsServiceCreateRepositoryCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repocreds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.upsert != nil {
		parameterAddToQuery(localVarQueryParams, "upsert", r.upsert, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRepoCredsServiceDeleteRepositoryCredentialsRequest struct {
	ctx        context.Context
	ApiService *RepoCredsServiceApiService
	url        string
}

func (r ApiRepoCredsServiceDeleteRepositoryCredentialsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RepoCredsServiceDeleteRepositoryCredentialsExecute(r)
}

/*
RepoCredsServiceDeleteRepositoryCredentials DeleteRepositoryCredentials deletes a repository credential set from the configuration

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param url
	@return ApiRepoCredsServiceDeleteRepositoryCredentialsRequest
*/
func (a *RepoCredsServiceApiService) RepoCredsServiceDeleteRepositoryCredentials(ctx context.Context, url string) ApiRepoCredsServiceDeleteRepositoryCredentialsRequest {
	return ApiRepoCredsServiceDeleteRepositoryCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		url:        url,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *RepoCredsServiceApiService) RepoCredsServiceDeleteRepositoryCredentialsExecute(r ApiRepoCredsServiceDeleteRepositoryCredentialsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepoCredsServiceApiService.RepoCredsServiceDeleteRepositoryCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repocreds/{url}"
	localVarPath = strings.Replace(localVarPath, "{"+"url"+"}", url.PathEscape(parameterValueToString(r.url, "url")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRepoCredsServiceListRepositoryCredentialsRequest struct {
	ctx        context.Context
	ApiService *RepoCredsServiceApiService
	url        *string
}

// Repo URL for query.
func (r ApiRepoCredsServiceListRepositoryCredentialsRequest) Url(url string) ApiRepoCredsServiceListRepositoryCredentialsRequest {
	r.url = &url
	return r
}

func (r ApiRepoCredsServiceListRepositoryCredentialsRequest) Execute() (*V1alpha1RepoCredsList, *http.Response, error) {
	return r.ApiService.RepoCredsServiceListRepositoryCredentialsExecute(r)
}

/*
RepoCredsServiceListRepositoryCredentials ListRepositoryCredentials gets a list of all configured repository credential sets

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRepoCredsServiceListRepositoryCredentialsRequest
*/
func (a *RepoCredsServiceApiService) RepoCredsServiceListRepositoryCredentials(ctx context.Context) ApiRepoCredsServiceListRepositoryCredentialsRequest {
	return ApiRepoCredsServiceListRepositoryCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepoCredsList
func (a *RepoCredsServiceApiService) RepoCredsServiceListRepositoryCredentialsExecute(r ApiRepoCredsServiceListRepositoryCredentialsRequest) (*V1alpha1RepoCredsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepoCredsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepoCredsServiceApiService.RepoCredsServiceListRepositoryCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repocreds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.url != nil {
		parameterAddToQuery(localVarQueryParams, "url", r.url, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRepoCredsServiceUpdateRepositoryCredentialsRequest struct {
	ctx        context.Context
	ApiService *RepoCredsServiceApiService
	credsUrl   string
	body       *V1alpha1RepoCreds
}

func (r ApiRepoCredsServiceUpdateRepositoryCredentialsRequest) Body(body V1alpha1RepoCreds) ApiRepoCredsServiceUpdateRepositoryCredentialsRequest {
	r.body = &body
	return r
}

func (r ApiRepoCredsServiceUpdateRepositoryCredentialsRequest) Execute() (*V1alpha1RepoCreds, *http.Response, error) {
	return r.ApiService.RepoCredsServiceUpdateRepositoryCredentialsExecute(r)
}

/*
RepoCredsServiceUpdateRepositoryCredentials UpdateRepositoryCredentials updates a repository credential set

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param credsUrl URL is the URL that this credentials matches to
	@return ApiRepoCredsServiceUpdateRepositoryCredentialsRequest
*/
func (a *RepoCredsServiceApiService) RepoCredsServiceUpdateRepositoryCredentials(ctx context.Context, credsUrl string) ApiRepoCredsServiceUpdateRepositoryCredentialsRequest {
	return ApiRepoCredsServiceUpdateRepositoryCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		credsUrl:   credsUrl,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepoCreds
func (a *RepoCredsServiceApiService) RepoCredsServiceUpdateRepositoryCredentialsExecute(r ApiRepoCredsServiceUpdateRepositoryCredentialsRequest) (*V1alpha1RepoCreds, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepoCreds
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepoCredsServiceApiService.RepoCredsServiceUpdateRepositoryCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repocreds/{creds.url}"
	localVarPath = strings.Replace(localVarPath, "{"+"creds.url"+"}", url.PathEscape(parameterValueToString(r.credsUrl, "credsUrl")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
