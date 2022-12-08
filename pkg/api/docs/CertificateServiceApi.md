# \CertificateServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateServiceCreateCertificate**](CertificateServiceApi.md#CertificateServiceCreateCertificate) | **Post** /api/v1/certificates | Creates repository certificates on the server
[**CertificateServiceDeleteCertificate**](CertificateServiceApi.md#CertificateServiceDeleteCertificate) | **Delete** /api/v1/certificates | Delete the certificates that match the RepositoryCertificateQuery
[**CertificateServiceListCertificates**](CertificateServiceApi.md#CertificateServiceListCertificates) | **Get** /api/v1/certificates | List all available repository certificates



## CertificateServiceCreateCertificate

> V1alpha1RepositoryCertificateList CertificateServiceCreateCertificate(ctx).Body(body).Upsert(upsert).Execute()

Creates repository certificates on the server

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
    body := *openapiclient.NewV1alpha1RepositoryCertificateList() // V1alpha1RepositoryCertificateList | List of certificates to be created
    upsert := true // bool | Whether to upsert already existing certificates. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateServiceApi.CertificateServiceCreateCertificate(context.Background()).Body(body).Upsert(upsert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.CertificateServiceCreateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateServiceCreateCertificate`: V1alpha1RepositoryCertificateList
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.CertificateServiceCreateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateServiceCreateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1RepositoryCertificateList**](V1alpha1RepositoryCertificateList.md) | List of certificates to be created | 
 **upsert** | **bool** | Whether to upsert already existing certificates. | 

### Return type

[**V1alpha1RepositoryCertificateList**](V1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateServiceDeleteCertificate

> V1alpha1RepositoryCertificateList CertificateServiceDeleteCertificate(ctx).HostNamePattern(hostNamePattern).CertType(certType).CertSubType(certSubType).Execute()

Delete the certificates that match the RepositoryCertificateQuery

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
    hostNamePattern := "hostNamePattern_example" // string | A file-glob pattern (not regular expression) the host name has to match. (optional)
    certType := "certType_example" // string | The type of the certificate to match (ssh or https). (optional)
    certSubType := "certSubType_example" // string | The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateServiceApi.CertificateServiceDeleteCertificate(context.Background()).HostNamePattern(hostNamePattern).CertType(certType).CertSubType(certSubType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.CertificateServiceDeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateServiceDeleteCertificate`: V1alpha1RepositoryCertificateList
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.CertificateServiceDeleteCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateServiceDeleteCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostNamePattern** | **string** | A file-glob pattern (not regular expression) the host name has to match. | 
 **certType** | **string** | The type of the certificate to match (ssh or https). | 
 **certSubType** | **string** | The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**V1alpha1RepositoryCertificateList**](V1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateServiceListCertificates

> V1alpha1RepositoryCertificateList CertificateServiceListCertificates(ctx).HostNamePattern(hostNamePattern).CertType(certType).CertSubType(certSubType).Execute()

List all available repository certificates

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
    hostNamePattern := "hostNamePattern_example" // string | A file-glob pattern (not regular expression) the host name has to match. (optional)
    certType := "certType_example" // string | The type of the certificate to match (ssh or https). (optional)
    certSubType := "certSubType_example" // string | The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateServiceApi.CertificateServiceListCertificates(context.Background()).HostNamePattern(hostNamePattern).CertType(certType).CertSubType(certSubType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.CertificateServiceListCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateServiceListCertificates`: V1alpha1RepositoryCertificateList
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.CertificateServiceListCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateServiceListCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostNamePattern** | **string** | A file-glob pattern (not regular expression) the host name has to match. | 
 **certType** | **string** | The type of the certificate to match (ssh or https). | 
 **certSubType** | **string** | The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**V1alpha1RepositoryCertificateList**](V1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

