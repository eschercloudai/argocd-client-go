# \ApplicationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationServiceCreate**](ApplicationServiceApi.md#ApplicationServiceCreate) | **Post** /api/v1/applications | Create creates an application
[**ApplicationServiceDelete**](ApplicationServiceApi.md#ApplicationServiceDelete) | **Delete** /api/v1/applications/{name} | Delete deletes an application
[**ApplicationServiceDeleteResource**](ApplicationServiceApi.md#ApplicationServiceDeleteResource) | **Delete** /api/v1/applications/{name}/resource | DeleteResource deletes a single application resource
[**ApplicationServiceGet**](ApplicationServiceApi.md#ApplicationServiceGet) | **Get** /api/v1/applications/{name} | Get returns an application by name
[**ApplicationServiceGetApplicationSyncWindows**](ApplicationServiceApi.md#ApplicationServiceGetApplicationSyncWindows) | **Get** /api/v1/applications/{name}/syncwindows | Get returns sync windows of the application
[**ApplicationServiceGetManifests**](ApplicationServiceApi.md#ApplicationServiceGetManifests) | **Get** /api/v1/applications/{name}/manifests | GetManifests returns application manifests
[**ApplicationServiceGetManifestsWithFiles**](ApplicationServiceApi.md#ApplicationServiceGetManifestsWithFiles) | **Post** /api/v1/applications/manifestsWithFiles | GetManifestsWithFiles returns application manifests using provided files to generate them
[**ApplicationServiceGetResource**](ApplicationServiceApi.md#ApplicationServiceGetResource) | **Get** /api/v1/applications/{name}/resource | GetResource returns single application resource
[**ApplicationServiceList**](ApplicationServiceApi.md#ApplicationServiceList) | **Get** /api/v1/applications | List returns list of applications
[**ApplicationServiceListResourceActions**](ApplicationServiceApi.md#ApplicationServiceListResourceActions) | **Get** /api/v1/applications/{name}/resource/actions | ListResourceActions returns list of resource actions
[**ApplicationServiceListResourceEvents**](ApplicationServiceApi.md#ApplicationServiceListResourceEvents) | **Get** /api/v1/applications/{name}/events | ListResourceEvents returns a list of event resources
[**ApplicationServiceManagedResources**](ApplicationServiceApi.md#ApplicationServiceManagedResources) | **Get** /api/v1/applications/{applicationName}/managed-resources | ManagedResources returns list of managed resources
[**ApplicationServicePatch**](ApplicationServiceApi.md#ApplicationServicePatch) | **Patch** /api/v1/applications/{name} | Patch patch an application
[**ApplicationServicePatchResource**](ApplicationServiceApi.md#ApplicationServicePatchResource) | **Post** /api/v1/applications/{name}/resource | PatchResource patch single application resource
[**ApplicationServicePodLogs**](ApplicationServiceApi.md#ApplicationServicePodLogs) | **Get** /api/v1/applications/{name}/pods/{podName}/logs | PodLogs returns stream of log entries for the specified pod. Pod
[**ApplicationServicePodLogs2**](ApplicationServiceApi.md#ApplicationServicePodLogs2) | **Get** /api/v1/applications/{name}/logs | PodLogs returns stream of log entries for the specified pod. Pod
[**ApplicationServiceResourceTree**](ApplicationServiceApi.md#ApplicationServiceResourceTree) | **Get** /api/v1/applications/{applicationName}/resource-tree | ResourceTree returns resource tree
[**ApplicationServiceRevisionMetadata**](ApplicationServiceApi.md#ApplicationServiceRevisionMetadata) | **Get** /api/v1/applications/{name}/revisions/{revision}/metadata | Get the meta-data (author, date, tags, message) for a specific revision of the application
[**ApplicationServiceRollback**](ApplicationServiceApi.md#ApplicationServiceRollback) | **Post** /api/v1/applications/{name}/rollback | Rollback syncs an application to its target state
[**ApplicationServiceRunResourceAction**](ApplicationServiceApi.md#ApplicationServiceRunResourceAction) | **Post** /api/v1/applications/{name}/resource/actions | RunResourceAction run resource action
[**ApplicationServiceSync**](ApplicationServiceApi.md#ApplicationServiceSync) | **Post** /api/v1/applications/{name}/sync | Sync syncs an application to its target state
[**ApplicationServiceTerminateOperation**](ApplicationServiceApi.md#ApplicationServiceTerminateOperation) | **Delete** /api/v1/applications/{name}/operation | TerminateOperation terminates the currently running operation
[**ApplicationServiceUpdate**](ApplicationServiceApi.md#ApplicationServiceUpdate) | **Put** /api/v1/applications/{application.metadata.name} | Update updates an application
[**ApplicationServiceUpdateSpec**](ApplicationServiceApi.md#ApplicationServiceUpdateSpec) | **Put** /api/v1/applications/{name}/spec | UpdateSpec updates an application spec
[**ApplicationServiceWatch**](ApplicationServiceApi.md#ApplicationServiceWatch) | **Get** /api/v1/stream/applications | Watch returns stream of application change events
[**ApplicationServiceWatchResourceTree**](ApplicationServiceApi.md#ApplicationServiceWatchResourceTree) | **Get** /api/v1/stream/applications/{applicationName}/resource-tree | Watch returns stream of application resource tree



## ApplicationServiceCreate

> V1alpha1Application ApplicationServiceCreate(ctx).Body(body).Upsert(upsert).Validate(validate).Execute()

Create creates an application

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
    body := *openapiclient.NewV1alpha1Application() // V1alpha1Application | 
    upsert := true // bool |  (optional)
    validate := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceCreate(context.Background()).Body(body).Upsert(upsert).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceCreate`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1Application**](V1alpha1Application.md) |  | 
 **upsert** | **bool** |  | 
 **validate** | **bool** |  | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceDelete

> map[string]interface{} ApplicationServiceDelete(ctx, name).Cascade(cascade).PropagationPolicy(propagationPolicy).AppNamespace(appNamespace).Execute()

Delete deletes an application

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
    name := "name_example" // string | 
    cascade := true // bool |  (optional)
    propagationPolicy := "propagationPolicy_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceDelete(context.Background(), name).Cascade(cascade).PropagationPolicy(propagationPolicy).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** |  | 
 **propagationPolicy** | **string** |  | 
 **appNamespace** | **string** |  | 

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


## ApplicationServiceDeleteResource

> map[string]interface{} ApplicationServiceDeleteResource(ctx, name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).Force(force).Orphan(orphan).AppNamespace(appNamespace).Execute()

DeleteResource deletes a single application resource

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
    name := "name_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    force := true // bool |  (optional)
    orphan := true // bool |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceDeleteResource(context.Background(), name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).Force(force).Orphan(orphan).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceDeleteResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceDeleteResource`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceDeleteResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceDeleteResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **force** | **bool** |  | 
 **orphan** | **bool** |  | 
 **appNamespace** | **string** |  | 

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


## ApplicationServiceGet

> V1alpha1Application ApplicationServiceGet(ctx, name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()

Get returns an application by name

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
    name := "name_example" // string | the application's name
    refresh := "refresh_example" // string | forces application reconciliation if set to true. (optional)
    projects := []string{"Inner_example"} // []string | the project names to restrict returned list applications. (optional)
    resourceVersion := "resourceVersion_example" // string | when specified with a watch call, shows changes that occur after that particular version of a resource. (optional)
    selector := "selector_example" // string | the selector to restrict returned list to applications only with matched labels. (optional)
    repo := "repo_example" // string | the repoURL to restrict returned list applications. (optional)
    appNamespace := "appNamespace_example" // string | the application's namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceGet(context.Background(), name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceGet`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the application&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **string** | forces application reconciliation if set to true. | 
 **projects** | **[]string** | the project names to restrict returned list applications. | 
 **resourceVersion** | **string** | when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **string** | the selector to restrict returned list to applications only with matched labels. | 
 **repo** | **string** | the repoURL to restrict returned list applications. | 
 **appNamespace** | **string** | the application&#39;s namespace. | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceGetApplicationSyncWindows

> ApplicationApplicationSyncWindowsResponse ApplicationServiceGetApplicationSyncWindows(ctx, name).AppNamespace(appNamespace).Execute()

Get returns sync windows of the application

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
    name := "name_example" // string | 
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceGetApplicationSyncWindows(context.Background(), name).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceGetApplicationSyncWindows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceGetApplicationSyncWindows`: ApplicationApplicationSyncWindowsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceGetApplicationSyncWindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceGetApplicationSyncWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appNamespace** | **string** |  | 

### Return type

[**ApplicationApplicationSyncWindowsResponse**](ApplicationApplicationSyncWindowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceGetManifests

> RepositoryManifestResponse ApplicationServiceGetManifests(ctx, name).Revision(revision).AppNamespace(appNamespace).Execute()

GetManifests returns application manifests

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
    name := "name_example" // string | 
    revision := "revision_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceGetManifests(context.Background(), name).Revision(revision).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceGetManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceGetManifests`: RepositoryManifestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceGetManifests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceGetManifestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**RepositoryManifestResponse**](RepositoryManifestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceGetManifestsWithFiles

> RepositoryManifestResponse ApplicationServiceGetManifestsWithFiles(ctx).Body(body).Execute()

GetManifestsWithFiles returns application manifests using provided files to generate them

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
    body := *openapiclient.NewApplicationApplicationManifestQueryWithFilesWrapper() // ApplicationApplicationManifestQueryWithFilesWrapper |  (streaming inputs)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceGetManifestsWithFiles(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceGetManifestsWithFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceGetManifestsWithFiles`: RepositoryManifestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceGetManifestsWithFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceGetManifestsWithFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationApplicationManifestQueryWithFilesWrapper**](ApplicationApplicationManifestQueryWithFilesWrapper.md) |  (streaming inputs) | 

### Return type

[**RepositoryManifestResponse**](RepositoryManifestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceGetResource

> ApplicationApplicationResourceResponse ApplicationServiceGetResource(ctx, name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

GetResource returns single application resource

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
    name := "name_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceGetResource(context.Background(), name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceGetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceGetResource`: ApplicationApplicationResourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceGetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**ApplicationApplicationResourceResponse**](ApplicationApplicationResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceList

> V1alpha1ApplicationList ApplicationServiceList(ctx).Name(name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()

List returns list of applications

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
    name := "name_example" // string | the application's name. (optional)
    refresh := "refresh_example" // string | forces application reconciliation if set to true. (optional)
    projects := []string{"Inner_example"} // []string | the project names to restrict returned list applications. (optional)
    resourceVersion := "resourceVersion_example" // string | when specified with a watch call, shows changes that occur after that particular version of a resource. (optional)
    selector := "selector_example" // string | the selector to restrict returned list to applications only with matched labels. (optional)
    repo := "repo_example" // string | the repoURL to restrict returned list applications. (optional)
    appNamespace := "appNamespace_example" // string | the application's namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceList(context.Background()).Name(name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceList`: V1alpha1ApplicationList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | the application&#39;s name. | 
 **refresh** | **string** | forces application reconciliation if set to true. | 
 **projects** | **[]string** | the project names to restrict returned list applications. | 
 **resourceVersion** | **string** | when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **string** | the selector to restrict returned list to applications only with matched labels. | 
 **repo** | **string** | the repoURL to restrict returned list applications. | 
 **appNamespace** | **string** | the application&#39;s namespace. | 

### Return type

[**V1alpha1ApplicationList**](V1alpha1ApplicationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceListResourceActions

> ApplicationResourceActionsListResponse ApplicationServiceListResourceActions(ctx, name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

ListResourceActions returns list of resource actions

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
    name := "name_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceListResourceActions(context.Background(), name).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceListResourceActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceListResourceActions`: ApplicationResourceActionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceListResourceActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceListResourceActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**ApplicationResourceActionsListResponse**](ApplicationResourceActionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceListResourceEvents

> V1EventList ApplicationServiceListResourceEvents(ctx, name).ResourceNamespace(resourceNamespace).ResourceName(resourceName).ResourceUID(resourceUID).AppNamespace(appNamespace).Execute()

ListResourceEvents returns a list of event resources

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
    name := "name_example" // string | 
    resourceNamespace := "resourceNamespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    resourceUID := "resourceUID_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceListResourceEvents(context.Background(), name).ResourceNamespace(resourceNamespace).ResourceName(resourceName).ResourceUID(resourceUID).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceListResourceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceListResourceEvents`: V1EventList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceListResourceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceListResourceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceNamespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **resourceUID** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**V1EventList**](V1EventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceManagedResources

> ApplicationManagedResourcesResponse ApplicationServiceManagedResources(ctx, applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

ManagedResources returns list of managed resources

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
    applicationName := "applicationName_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceManagedResources(context.Background(), applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceManagedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceManagedResources`: ApplicationManagedResourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceManagedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceManagedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**ApplicationManagedResourcesResponse**](ApplicationManagedResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServicePatch

> V1alpha1Application ApplicationServicePatch(ctx, name).Body(body).Execute()

Patch patch an application

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
    name := "name_example" // string | 
    body := *openapiclient.NewApplicationApplicationPatchRequest() // ApplicationApplicationPatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServicePatch(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServicePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServicePatch`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServicePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServicePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationApplicationPatchRequest**](ApplicationApplicationPatchRequest.md) |  | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServicePatchResource

> ApplicationApplicationResourceResponse ApplicationServicePatchResource(ctx, name).Body(body).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).PatchType(patchType).AppNamespace(appNamespace).Execute()

PatchResource patch single application resource

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
    name := "name_example" // string | 
    body := "body_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    patchType := "patchType_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServicePatchResource(context.Background(), name).Body(body).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).PatchType(patchType).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServicePatchResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServicePatchResource`: ApplicationApplicationResourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServicePatchResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServicePatchResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 
 **namespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **patchType** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**ApplicationApplicationResourceResponse**](ApplicationApplicationResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServicePodLogs

> StreamResultOfApplicationLogEntry ApplicationServicePodLogs(ctx, name, podName).Namespace(namespace).Container(container).SinceSeconds(sinceSeconds).SinceTimeSeconds(sinceTimeSeconds).SinceTimeNanos(sinceTimeNanos).TailLines(tailLines).Follow(follow).UntilTime(untilTime).Filter(filter).Kind(kind).Group(group).ResourceName(resourceName).Previous(previous).AppNamespace(appNamespace).Execute()

PodLogs returns stream of log entries for the specified pod. Pod

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
    name := "name_example" // string | 
    podName := "podName_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    container := "container_example" // string |  (optional)
    sinceSeconds := "sinceSeconds_example" // string |  (optional)
    sinceTimeSeconds := "sinceTimeSeconds_example" // string | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. (optional)
    sinceTimeNanos := int32(56) // int32 | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. (optional)
    tailLines := "tailLines_example" // string |  (optional)
    follow := true // bool |  (optional)
    untilTime := "untilTime_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    previous := true // bool |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServicePodLogs(context.Background(), name, podName).Namespace(namespace).Container(container).SinceSeconds(sinceSeconds).SinceTimeSeconds(sinceTimeSeconds).SinceTimeNanos(sinceTimeNanos).TailLines(tailLines).Follow(follow).UntilTime(untilTime).Filter(filter).Kind(kind).Group(group).ResourceName(resourceName).Previous(previous).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServicePodLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServicePodLogs`: StreamResultOfApplicationLogEntry
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServicePodLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**podName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServicePodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **string** |  | 
 **container** | **string** |  | 
 **sinceSeconds** | **string** |  | 
 **sinceTimeSeconds** | **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **sinceTimeNanos** | **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **tailLines** | **string** |  | 
 **follow** | **bool** |  | 
 **untilTime** | **string** |  | 
 **filter** | **string** |  | 
 **kind** | **string** |  | 
 **group** | **string** |  | 
 **resourceName** | **string** |  | 
 **previous** | **bool** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**StreamResultOfApplicationLogEntry**](StreamResultOfApplicationLogEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServicePodLogs2

> StreamResultOfApplicationLogEntry ApplicationServicePodLogs2(ctx, name).Namespace(namespace).PodName(podName).Container(container).SinceSeconds(sinceSeconds).SinceTimeSeconds(sinceTimeSeconds).SinceTimeNanos(sinceTimeNanos).TailLines(tailLines).Follow(follow).UntilTime(untilTime).Filter(filter).Kind(kind).Group(group).ResourceName(resourceName).Previous(previous).AppNamespace(appNamespace).Execute()

PodLogs returns stream of log entries for the specified pod. Pod

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
    name := "name_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    podName := "podName_example" // string |  (optional)
    container := "container_example" // string |  (optional)
    sinceSeconds := "sinceSeconds_example" // string |  (optional)
    sinceTimeSeconds := "sinceTimeSeconds_example" // string | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. (optional)
    sinceTimeNanos := int32(56) // int32 | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. (optional)
    tailLines := "tailLines_example" // string |  (optional)
    follow := true // bool |  (optional)
    untilTime := "untilTime_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    previous := true // bool |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServicePodLogs2(context.Background(), name).Namespace(namespace).PodName(podName).Container(container).SinceSeconds(sinceSeconds).SinceTimeSeconds(sinceTimeSeconds).SinceTimeNanos(sinceTimeNanos).TailLines(tailLines).Follow(follow).UntilTime(untilTime).Filter(filter).Kind(kind).Group(group).ResourceName(resourceName).Previous(previous).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServicePodLogs2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServicePodLogs2`: StreamResultOfApplicationLogEntry
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServicePodLogs2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServicePodLogs2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **podName** | **string** |  | 
 **container** | **string** |  | 
 **sinceSeconds** | **string** |  | 
 **sinceTimeSeconds** | **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **sinceTimeNanos** | **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **tailLines** | **string** |  | 
 **follow** | **bool** |  | 
 **untilTime** | **string** |  | 
 **filter** | **string** |  | 
 **kind** | **string** |  | 
 **group** | **string** |  | 
 **resourceName** | **string** |  | 
 **previous** | **bool** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**StreamResultOfApplicationLogEntry**](StreamResultOfApplicationLogEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceResourceTree

> V1alpha1ApplicationTree ApplicationServiceResourceTree(ctx, applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

ResourceTree returns resource tree

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
    applicationName := "applicationName_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceResourceTree(context.Background(), applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceResourceTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceResourceTree`: V1alpha1ApplicationTree
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceResourceTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceResourceTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**V1alpha1ApplicationTree**](V1alpha1ApplicationTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceRevisionMetadata

> V1alpha1RevisionMetadata ApplicationServiceRevisionMetadata(ctx, name, revision).AppNamespace(appNamespace).Execute()

Get the meta-data (author, date, tags, message) for a specific revision of the application

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
    name := "name_example" // string | the application's name
    revision := "revision_example" // string | the revision of the app
    appNamespace := "appNamespace_example" // string | the application's namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceRevisionMetadata(context.Background(), name, revision).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceRevisionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceRevisionMetadata`: V1alpha1RevisionMetadata
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceRevisionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the application&#39;s name | 
**revision** | **string** | the revision of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceRevisionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appNamespace** | **string** | the application&#39;s namespace. | 

### Return type

[**V1alpha1RevisionMetadata**](V1alpha1RevisionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceRollback

> V1alpha1Application ApplicationServiceRollback(ctx, name).Body(body).Execute()

Rollback syncs an application to its target state

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
    name := "name_example" // string | 
    body := *openapiclient.NewApplicationApplicationRollbackRequest() // ApplicationApplicationRollbackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceRollback(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceRollback`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationApplicationRollbackRequest**](ApplicationApplicationRollbackRequest.md) |  | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceRunResourceAction

> map[string]interface{} ApplicationServiceRunResourceAction(ctx, name).Body(body).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

RunResourceAction run resource action

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
    name := "name_example" // string | 
    body := "body_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    resourceName := "resourceName_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceRunResourceAction(context.Background(), name).Body(body).Namespace(namespace).ResourceName(resourceName).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceRunResourceAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceRunResourceAction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceRunResourceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceRunResourceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 
 **namespace** | **string** |  | 
 **resourceName** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

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


## ApplicationServiceSync

> V1alpha1Application ApplicationServiceSync(ctx, name).Body(body).Execute()

Sync syncs an application to its target state

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
    name := "name_example" // string | 
    body := *openapiclient.NewApplicationApplicationSyncRequest() // ApplicationApplicationSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceSync(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceSync`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationApplicationSyncRequest**](ApplicationApplicationSyncRequest.md) |  | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceTerminateOperation

> map[string]interface{} ApplicationServiceTerminateOperation(ctx, name).AppNamespace(appNamespace).Execute()

TerminateOperation terminates the currently running operation

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
    name := "name_example" // string | 
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceTerminateOperation(context.Background(), name).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceTerminateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceTerminateOperation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceTerminateOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceTerminateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appNamespace** | **string** |  | 

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


## ApplicationServiceUpdate

> V1alpha1Application ApplicationServiceUpdate(ctx, applicationMetadataName).Body(body).Validate(validate).Execute()

Update updates an application

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
    applicationMetadataName := "applicationMetadataName_example" // string | Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional
    body := *openapiclient.NewV1alpha1Application() // V1alpha1Application | 
    validate := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceUpdate(context.Background(), applicationMetadataName).Body(body).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceUpdate`: V1alpha1Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationMetadataName** | **string** | Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1Application**](V1alpha1Application.md) |  | 
 **validate** | **bool** |  | 

### Return type

[**V1alpha1Application**](V1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceUpdateSpec

> V1alpha1ApplicationSpec ApplicationServiceUpdateSpec(ctx, name).Body(body).Validate(validate).AppNamespace(appNamespace).Execute()

UpdateSpec updates an application spec

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
    name := "name_example" // string | 
    body := *openapiclient.NewV1alpha1ApplicationSpec() // V1alpha1ApplicationSpec | 
    validate := true // bool |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceUpdateSpec(context.Background(), name).Body(body).Validate(validate).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceUpdateSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceUpdateSpec`: V1alpha1ApplicationSpec
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceUpdateSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceUpdateSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1ApplicationSpec**](V1alpha1ApplicationSpec.md) |  | 
 **validate** | **bool** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**V1alpha1ApplicationSpec**](V1alpha1ApplicationSpec.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceWatch

> StreamResultOfV1alpha1ApplicationWatchEvent ApplicationServiceWatch(ctx).Name(name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()

Watch returns stream of application change events

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
    name := "name_example" // string | the application's name. (optional)
    refresh := "refresh_example" // string | forces application reconciliation if set to true. (optional)
    projects := []string{"Inner_example"} // []string | the project names to restrict returned list applications. (optional)
    resourceVersion := "resourceVersion_example" // string | when specified with a watch call, shows changes that occur after that particular version of a resource. (optional)
    selector := "selector_example" // string | the selector to restrict returned list to applications only with matched labels. (optional)
    repo := "repo_example" // string | the repoURL to restrict returned list applications. (optional)
    appNamespace := "appNamespace_example" // string | the application's namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceWatch(context.Background()).Name(name).Refresh(refresh).Projects(projects).ResourceVersion(resourceVersion).Selector(selector).Repo(repo).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceWatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceWatch`: StreamResultOfV1alpha1ApplicationWatchEvent
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceWatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | the application&#39;s name. | 
 **refresh** | **string** | forces application reconciliation if set to true. | 
 **projects** | **[]string** | the project names to restrict returned list applications. | 
 **resourceVersion** | **string** | when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **string** | the selector to restrict returned list to applications only with matched labels. | 
 **repo** | **string** | the repoURL to restrict returned list applications. | 
 **appNamespace** | **string** | the application&#39;s namespace. | 

### Return type

[**StreamResultOfV1alpha1ApplicationWatchEvent**](StreamResultOfV1alpha1ApplicationWatchEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationServiceWatchResourceTree

> StreamResultOfV1alpha1ApplicationTree ApplicationServiceWatchResourceTree(ctx, applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()

Watch returns stream of application resource tree

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
    applicationName := "applicationName_example" // string | 
    namespace := "namespace_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    kind := "kind_example" // string |  (optional)
    appNamespace := "appNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationServiceApi.ApplicationServiceWatchResourceTree(context.Background(), applicationName).Namespace(namespace).Name(name).Version(version).Group(group).Kind(kind).AppNamespace(appNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServiceApi.ApplicationServiceWatchResourceTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationServiceWatchResourceTree`: StreamResultOfV1alpha1ApplicationTree
    fmt.Fprintf(os.Stdout, "Response from `ApplicationServiceApi.ApplicationServiceWatchResourceTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationServiceWatchResourceTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **group** | **string** |  | 
 **kind** | **string** |  | 
 **appNamespace** | **string** |  | 

### Return type

[**StreamResultOfV1alpha1ApplicationTree**](StreamResultOfV1alpha1ApplicationTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

