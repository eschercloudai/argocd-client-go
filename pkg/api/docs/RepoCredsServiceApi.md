# \RepoCredsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoCredsServiceCreateRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceCreateRepositoryCredentials) | **Post** /api/v1/repocreds | CreateRepositoryCredentials creates a new repository credential set
[**RepoCredsServiceDeleteRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceDeleteRepositoryCredentials) | **Delete** /api/v1/repocreds/{url} | DeleteRepositoryCredentials deletes a repository credential set from the configuration
[**RepoCredsServiceListRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceListRepositoryCredentials) | **Get** /api/v1/repocreds | ListRepositoryCredentials gets a list of all configured repository credential sets
[**RepoCredsServiceUpdateRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceUpdateRepositoryCredentials) | **Put** /api/v1/repocreds/{creds.url} | UpdateRepositoryCredentials updates a repository credential set



## RepoCredsServiceCreateRepositoryCredentials

> V1alpha1RepoCreds RepoCredsServiceCreateRepositoryCredentials(ctx).Body(body).Upsert(upsert).Execute()

CreateRepositoryCredentials creates a new repository credential set

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
    body := *openapiclient.NewV1alpha1RepoCreds() // V1alpha1RepoCreds | Repository definition
    upsert := true // bool | Whether to create in upsert mode. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepoCredsServiceApi.RepoCredsServiceCreateRepositoryCredentials(context.Background()).Body(body).Upsert(upsert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepoCredsServiceApi.RepoCredsServiceCreateRepositoryCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepoCredsServiceCreateRepositoryCredentials`: V1alpha1RepoCreds
    fmt.Fprintf(os.Stdout, "Response from `RepoCredsServiceApi.RepoCredsServiceCreateRepositoryCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepoCredsServiceCreateRepositoryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1RepoCreds**](V1alpha1RepoCreds.md) | Repository definition | 
 **upsert** | **bool** | Whether to create in upsert mode. | 

### Return type

[**V1alpha1RepoCreds**](V1alpha1RepoCreds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoCredsServiceDeleteRepositoryCredentials

> map[string]interface{} RepoCredsServiceDeleteRepositoryCredentials(ctx, url).Execute()

DeleteRepositoryCredentials deletes a repository credential set from the configuration

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
    url := "url_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepoCredsServiceApi.RepoCredsServiceDeleteRepositoryCredentials(context.Background(), url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepoCredsServiceApi.RepoCredsServiceDeleteRepositoryCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepoCredsServiceDeleteRepositoryCredentials`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepoCredsServiceApi.RepoCredsServiceDeleteRepositoryCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepoCredsServiceDeleteRepositoryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoCredsServiceListRepositoryCredentials

> V1alpha1RepoCredsList RepoCredsServiceListRepositoryCredentials(ctx).Url(url).Execute()

ListRepositoryCredentials gets a list of all configured repository credential sets

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
    url := "url_example" // string | Repo URL for query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepoCredsServiceApi.RepoCredsServiceListRepositoryCredentials(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepoCredsServiceApi.RepoCredsServiceListRepositoryCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepoCredsServiceListRepositoryCredentials`: V1alpha1RepoCredsList
    fmt.Fprintf(os.Stdout, "Response from `RepoCredsServiceApi.RepoCredsServiceListRepositoryCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepoCredsServiceListRepositoryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | Repo URL for query. | 

### Return type

[**V1alpha1RepoCredsList**](V1alpha1RepoCredsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepoCredsServiceUpdateRepositoryCredentials

> V1alpha1RepoCreds RepoCredsServiceUpdateRepositoryCredentials(ctx, credsUrl).Body(body).Execute()

UpdateRepositoryCredentials updates a repository credential set

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
    credsUrl := "credsUrl_example" // string | URL is the URL that this credentials matches to
    body := *openapiclient.NewV1alpha1RepoCreds() // V1alpha1RepoCreds | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepoCredsServiceApi.RepoCredsServiceUpdateRepositoryCredentials(context.Background(), credsUrl).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepoCredsServiceApi.RepoCredsServiceUpdateRepositoryCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepoCredsServiceUpdateRepositoryCredentials`: V1alpha1RepoCreds
    fmt.Fprintf(os.Stdout, "Response from `RepoCredsServiceApi.RepoCredsServiceUpdateRepositoryCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credsUrl** | **string** | URL is the URL that this credentials matches to | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepoCredsServiceUpdateRepositoryCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1RepoCreds**](V1alpha1RepoCreds.md) |  | 

### Return type

[**V1alpha1RepoCreds**](V1alpha1RepoCreds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

