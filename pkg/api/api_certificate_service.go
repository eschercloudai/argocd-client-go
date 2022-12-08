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
)

// CertificateServiceApiService CertificateServiceApi service
type CertificateServiceApiService service

type ApiCertificateServiceCreateCertificateRequest struct {
	ctx        context.Context
	ApiService *CertificateServiceApiService
	body       *V1alpha1RepositoryCertificateList
	upsert     *bool
}

// List of certificates to be created
func (r ApiCertificateServiceCreateCertificateRequest) Body(body V1alpha1RepositoryCertificateList) ApiCertificateServiceCreateCertificateRequest {
	r.body = &body
	return r
}

// Whether to upsert already existing certificates.
func (r ApiCertificateServiceCreateCertificateRequest) Upsert(upsert bool) ApiCertificateServiceCreateCertificateRequest {
	r.upsert = &upsert
	return r
}

func (r ApiCertificateServiceCreateCertificateRequest) Execute() (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	return r.ApiService.CertificateServiceCreateCertificateExecute(r)
}

/*
CertificateServiceCreateCertificate Creates repository certificates on the server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCertificateServiceCreateCertificateRequest
*/
func (a *CertificateServiceApiService) CertificateServiceCreateCertificate(ctx context.Context) ApiCertificateServiceCreateCertificateRequest {
	return ApiCertificateServiceCreateCertificateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepositoryCertificateList
func (a *CertificateServiceApiService) CertificateServiceCreateCertificateExecute(r ApiCertificateServiceCreateCertificateRequest) (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepositoryCertificateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateServiceApiService.CertificateServiceCreateCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates"

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

type ApiCertificateServiceDeleteCertificateRequest struct {
	ctx             context.Context
	ApiService      *CertificateServiceApiService
	hostNamePattern *string
	certType        *string
	certSubType     *string
}

// A file-glob pattern (not regular expression) the host name has to match.
func (r ApiCertificateServiceDeleteCertificateRequest) HostNamePattern(hostNamePattern string) ApiCertificateServiceDeleteCertificateRequest {
	r.hostNamePattern = &hostNamePattern
	return r
}

// The type of the certificate to match (ssh or https).
func (r ApiCertificateServiceDeleteCertificateRequest) CertType(certType string) ApiCertificateServiceDeleteCertificateRequest {
	r.certType = &certType
	return r
}

// The sub type of the certificate to match (protocol dependent, usually only used for ssh certs).
func (r ApiCertificateServiceDeleteCertificateRequest) CertSubType(certSubType string) ApiCertificateServiceDeleteCertificateRequest {
	r.certSubType = &certSubType
	return r
}

func (r ApiCertificateServiceDeleteCertificateRequest) Execute() (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	return r.ApiService.CertificateServiceDeleteCertificateExecute(r)
}

/*
CertificateServiceDeleteCertificate Delete the certificates that match the RepositoryCertificateQuery

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCertificateServiceDeleteCertificateRequest
*/
func (a *CertificateServiceApiService) CertificateServiceDeleteCertificate(ctx context.Context) ApiCertificateServiceDeleteCertificateRequest {
	return ApiCertificateServiceDeleteCertificateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepositoryCertificateList
func (a *CertificateServiceApiService) CertificateServiceDeleteCertificateExecute(r ApiCertificateServiceDeleteCertificateRequest) (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepositoryCertificateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateServiceApiService.CertificateServiceDeleteCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.hostNamePattern != nil {
		parameterAddToQuery(localVarQueryParams, "hostNamePattern", r.hostNamePattern, "")
	}
	if r.certType != nil {
		parameterAddToQuery(localVarQueryParams, "certType", r.certType, "")
	}
	if r.certSubType != nil {
		parameterAddToQuery(localVarQueryParams, "certSubType", r.certSubType, "")
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

type ApiCertificateServiceListCertificatesRequest struct {
	ctx             context.Context
	ApiService      *CertificateServiceApiService
	hostNamePattern *string
	certType        *string
	certSubType     *string
}

// A file-glob pattern (not regular expression) the host name has to match.
func (r ApiCertificateServiceListCertificatesRequest) HostNamePattern(hostNamePattern string) ApiCertificateServiceListCertificatesRequest {
	r.hostNamePattern = &hostNamePattern
	return r
}

// The type of the certificate to match (ssh or https).
func (r ApiCertificateServiceListCertificatesRequest) CertType(certType string) ApiCertificateServiceListCertificatesRequest {
	r.certType = &certType
	return r
}

// The sub type of the certificate to match (protocol dependent, usually only used for ssh certs).
func (r ApiCertificateServiceListCertificatesRequest) CertSubType(certSubType string) ApiCertificateServiceListCertificatesRequest {
	r.certSubType = &certSubType
	return r
}

func (r ApiCertificateServiceListCertificatesRequest) Execute() (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	return r.ApiService.CertificateServiceListCertificatesExecute(r)
}

/*
CertificateServiceListCertificates List all available repository certificates

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCertificateServiceListCertificatesRequest
*/
func (a *CertificateServiceApiService) CertificateServiceListCertificates(ctx context.Context) ApiCertificateServiceListCertificatesRequest {
	return ApiCertificateServiceListCertificatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1RepositoryCertificateList
func (a *CertificateServiceApiService) CertificateServiceListCertificatesExecute(r ApiCertificateServiceListCertificatesRequest) (*V1alpha1RepositoryCertificateList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1alpha1RepositoryCertificateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateServiceApiService.CertificateServiceListCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.hostNamePattern != nil {
		parameterAddToQuery(localVarQueryParams, "hostNamePattern", r.hostNamePattern, "")
	}
	if r.certType != nil {
		parameterAddToQuery(localVarQueryParams, "certType", r.certType, "")
	}
	if r.certSubType != nil {
		parameterAddToQuery(localVarQueryParams, "certSubType", r.certSubType, "")
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