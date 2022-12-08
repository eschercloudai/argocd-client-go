# \RepositoryServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoryServiceCreateRepository**](RepositoryServiceApi.md#RepositoryServiceCreateRepository) | **Post** /api/v1/repositories | CreateRepository creates a new repository configuration
[**RepositoryServiceDeleteRepository**](RepositoryServiceApi.md#RepositoryServiceDeleteRepository) | **Delete** /api/v1/repositories/{repo} | DeleteRepository deletes a repository from the configuration
[**RepositoryServiceGet**](RepositoryServiceApi.md#RepositoryServiceGet) | **Get** /api/v1/repositories/{repo} | Get returns a repository or its credentials
[**RepositoryServiceGetAppDetails**](RepositoryServiceApi.md#RepositoryServiceGetAppDetails) | **Post** /api/v1/repositories/{source.repoURL}/appdetails | GetAppDetails returns application details by given path
[**RepositoryServiceGetHelmCharts**](RepositoryServiceApi.md#RepositoryServiceGetHelmCharts) | **Get** /api/v1/repositories/{repo}/helmcharts | GetHelmCharts returns list of helm charts in the specified repository
[**RepositoryServiceListApps**](RepositoryServiceApi.md#RepositoryServiceListApps) | **Get** /api/v1/repositories/{repo}/apps | ListApps returns list of apps in the repo
[**RepositoryServiceListRefs**](RepositoryServiceApi.md#RepositoryServiceListRefs) | **Get** /api/v1/repositories/{repo}/refs | 
[**RepositoryServiceListRepositories**](RepositoryServiceApi.md#RepositoryServiceListRepositories) | **Get** /api/v1/repositories | ListRepositories gets a list of all configured repositories
[**RepositoryServiceUpdateRepository**](RepositoryServiceApi.md#RepositoryServiceUpdateRepository) | **Put** /api/v1/repositories/{repo.repo} | UpdateRepository updates a repository configuration
[**RepositoryServiceValidateAccess**](RepositoryServiceApi.md#RepositoryServiceValidateAccess) | **Post** /api/v1/repositories/{repo}/validate | ValidateAccess validates access to a repository with given parameters



## RepositoryServiceCreateRepository

> V1alpha1Repository RepositoryServiceCreateRepository(ctx).Body(body).Upsert(upsert).CredsOnly(credsOnly).Execute()

CreateRepository creates a new repository configuration

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
    body := *openapiclient.NewV1alpha1Repository() // V1alpha1Repository | Repository definition
    upsert := true // bool | Whether to create in upsert mode. (optional)
    credsOnly := true // bool | Whether to operate on credential set instead of repository. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceCreateRepository(context.Background()).Body(body).Upsert(upsert).CredsOnly(credsOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceCreateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceCreateRepository`: V1alpha1Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceCreateRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1Repository**](V1alpha1Repository.md) | Repository definition | 
 **upsert** | **bool** | Whether to create in upsert mode. | 
 **credsOnly** | **bool** | Whether to operate on credential set instead of repository. | 

### Return type

[**V1alpha1Repository**](V1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceDeleteRepository

> map[string]interface{} RepositoryServiceDeleteRepository(ctx, repo).ForceRefresh(forceRefresh).Execute()

DeleteRepository deletes a repository from the configuration

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
    repo := "repo_example" // string | Repo URL for query
    forceRefresh := true // bool | Whether to force a cache refresh on repo's connection state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceDeleteRepository(context.Background(), repo).ForceRefresh(forceRefresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceDeleteRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceDeleteRepository`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceDeleteRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo URL for query | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **bool** | Whether to force a cache refresh on repo&#39;s connection state. | 

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


## RepositoryServiceGet

> V1alpha1Repository RepositoryServiceGet(ctx, repo).ForceRefresh(forceRefresh).Execute()

Get returns a repository or its credentials

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
    repo := "repo_example" // string | Repo URL for query
    forceRefresh := true // bool | Whether to force a cache refresh on repo's connection state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceGet(context.Background(), repo).ForceRefresh(forceRefresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceGet`: V1alpha1Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo URL for query | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **bool** | Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**V1alpha1Repository**](V1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceGetAppDetails

> RepositoryRepoAppDetailsResponse RepositoryServiceGetAppDetails(ctx, sourceRepoURL).Body(body).Execute()

GetAppDetails returns application details by given path

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
    sourceRepoURL := "sourceRepoURL_example" // string | RepoURL is the URL to the repository (Git or Helm) that contains the application manifests
    body := *openapiclient.NewRepositoryRepoAppDetailsQuery() // RepositoryRepoAppDetailsQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceGetAppDetails(context.Background(), sourceRepoURL).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceGetAppDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceGetAppDetails`: RepositoryRepoAppDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceGetAppDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceRepoURL** | **string** | RepoURL is the URL to the repository (Git or Helm) that contains the application manifests | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceGetAppDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RepositoryRepoAppDetailsQuery**](RepositoryRepoAppDetailsQuery.md) |  | 

### Return type

[**RepositoryRepoAppDetailsResponse**](RepositoryRepoAppDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceGetHelmCharts

> RepositoryHelmChartsResponse RepositoryServiceGetHelmCharts(ctx, repo).ForceRefresh(forceRefresh).Execute()

GetHelmCharts returns list of helm charts in the specified repository

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
    repo := "repo_example" // string | Repo URL for query
    forceRefresh := true // bool | Whether to force a cache refresh on repo's connection state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceGetHelmCharts(context.Background(), repo).ForceRefresh(forceRefresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceGetHelmCharts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceGetHelmCharts`: RepositoryHelmChartsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceGetHelmCharts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo URL for query | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceGetHelmChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **bool** | Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**RepositoryHelmChartsResponse**](RepositoryHelmChartsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceListApps

> RepositoryRepoAppsResponse RepositoryServiceListApps(ctx, repo).Revision(revision).AppName(appName).AppProject(appProject).Execute()

ListApps returns list of apps in the repo

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
    repo := "repo_example" // string | 
    revision := "revision_example" // string |  (optional)
    appName := "appName_example" // string |  (optional)
    appProject := "appProject_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceListApps(context.Background(), repo).Revision(revision).AppName(appName).AppProject(appProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceListApps`: RepositoryRepoAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** |  | 
 **appName** | **string** |  | 
 **appProject** | **string** |  | 

### Return type

[**RepositoryRepoAppsResponse**](RepositoryRepoAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceListRefs

> RepositoryRefs RepositoryServiceListRefs(ctx, repo).ForceRefresh(forceRefresh).Execute()



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
    repo := "repo_example" // string | Repo URL for query
    forceRefresh := true // bool | Whether to force a cache refresh on repo's connection state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceListRefs(context.Background(), repo).ForceRefresh(forceRefresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceListRefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceListRefs`: RepositoryRefs
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceListRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repo URL for query | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceListRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **bool** | Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**RepositoryRefs**](RepositoryRefs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceListRepositories

> V1alpha1RepositoryList RepositoryServiceListRepositories(ctx).Repo(repo).ForceRefresh(forceRefresh).Execute()

ListRepositories gets a list of all configured repositories

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
    repo := "repo_example" // string | Repo URL for query. (optional)
    forceRefresh := true // bool | Whether to force a cache refresh on repo's connection state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceListRepositories(context.Background()).Repo(repo).ForceRefresh(forceRefresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceListRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceListRepositories`: V1alpha1RepositoryList
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceListRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceListRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | Repo URL for query. | 
 **forceRefresh** | **bool** | Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**V1alpha1RepositoryList**](V1alpha1RepositoryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceUpdateRepository

> V1alpha1Repository RepositoryServiceUpdateRepository(ctx, repoRepo).Body(body).Execute()

UpdateRepository updates a repository configuration

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
    repoRepo := "repoRepo_example" // string | Repo contains the URL to the remote repository
    body := *openapiclient.NewV1alpha1Repository() // V1alpha1Repository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceUpdateRepository(context.Background(), repoRepo).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceUpdateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceUpdateRepository`: V1alpha1Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceUpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoRepo** | **string** | Repo contains the URL to the remote repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1Repository**](V1alpha1Repository.md) |  | 

### Return type

[**V1alpha1Repository**](V1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryServiceValidateAccess

> map[string]interface{} RepositoryServiceValidateAccess(ctx, repo).Body(body).Username(username).Password(password).SshPrivateKey(sshPrivateKey).Insecure(insecure).TlsClientCertData(tlsClientCertData).TlsClientCertKey(tlsClientCertKey).Type_(type_).Name(name).EnableOci(enableOci).GithubAppPrivateKey(githubAppPrivateKey).GithubAppID(githubAppID).GithubAppInstallationID(githubAppInstallationID).GithubAppEnterpriseBaseUrl(githubAppEnterpriseBaseUrl).Proxy(proxy).Project(project).Execute()

ValidateAccess validates access to a repository with given parameters

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
    repo := "repo_example" // string | The URL to the repo
    body := "body_example" // string | The URL to the repo
    username := "username_example" // string | Username for accessing repo. (optional)
    password := "password_example" // string | Password for accessing repo. (optional)
    sshPrivateKey := "sshPrivateKey_example" // string | Private key data for accessing SSH repository. (optional)
    insecure := true // bool | Whether to skip certificate or host key validation. (optional)
    tlsClientCertData := "tlsClientCertData_example" // string | TLS client cert data for accessing HTTPS repository. (optional)
    tlsClientCertKey := "tlsClientCertKey_example" // string | TLS client cert key for accessing HTTPS repository. (optional)
    type_ := "type__example" // string | The type of the repo. (optional)
    name := "name_example" // string | The name of the repo. (optional)
    enableOci := true // bool | Whether helm-oci support should be enabled for this repo. (optional)
    githubAppPrivateKey := "githubAppPrivateKey_example" // string | Github App Private Key PEM data. (optional)
    githubAppID := "githubAppID_example" // string | Github App ID of the app used to access the repo. (optional)
    githubAppInstallationID := "githubAppInstallationID_example" // string | Github App Installation ID of the installed GitHub App. (optional)
    githubAppEnterpriseBaseUrl := "githubAppEnterpriseBaseUrl_example" // string | Github App Enterprise base url if empty will default to https://api.github.com. (optional)
    proxy := "proxy_example" // string | HTTP/HTTPS proxy to access the repository. (optional)
    project := "project_example" // string | Reference between project and repository that allow you automatically to be added as item inside SourceRepos project entity. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryServiceApi.RepositoryServiceValidateAccess(context.Background(), repo).Body(body).Username(username).Password(password).SshPrivateKey(sshPrivateKey).Insecure(insecure).TlsClientCertData(tlsClientCertData).TlsClientCertKey(tlsClientCertKey).Type_(type_).Name(name).EnableOci(enableOci).GithubAppPrivateKey(githubAppPrivateKey).GithubAppID(githubAppID).GithubAppInstallationID(githubAppInstallationID).GithubAppEnterpriseBaseUrl(githubAppEnterpriseBaseUrl).Proxy(proxy).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryServiceApi.RepositoryServiceValidateAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryServiceValidateAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepositoryServiceApi.RepositoryServiceValidateAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The URL to the repo | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryServiceValidateAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The URL to the repo | 
 **username** | **string** | Username for accessing repo. | 
 **password** | **string** | Password for accessing repo. | 
 **sshPrivateKey** | **string** | Private key data for accessing SSH repository. | 
 **insecure** | **bool** | Whether to skip certificate or host key validation. | 
 **tlsClientCertData** | **string** | TLS client cert data for accessing HTTPS repository. | 
 **tlsClientCertKey** | **string** | TLS client cert key for accessing HTTPS repository. | 
 **type_** | **string** | The type of the repo. | 
 **name** | **string** | The name of the repo. | 
 **enableOci** | **bool** | Whether helm-oci support should be enabled for this repo. | 
 **githubAppPrivateKey** | **string** | Github App Private Key PEM data. | 
 **githubAppID** | **string** | Github App ID of the app used to access the repo. | 
 **githubAppInstallationID** | **string** | Github App Installation ID of the installed GitHub App. | 
 **githubAppEnterpriseBaseUrl** | **string** | Github App Enterprise base url if empty will default to https://api.github.com. | 
 **proxy** | **string** | HTTP/HTTPS proxy to access the repository. | 
 **project** | **string** | Reference between project and repository that allow you automatically to be added as item inside SourceRepos project entity. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

