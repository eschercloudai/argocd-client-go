# \ProjectServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectServiceCreate**](ProjectServiceApi.md#ProjectServiceCreate) | **Post** /api/v1/projects | Create a new project
[**ProjectServiceCreateToken**](ProjectServiceApi.md#ProjectServiceCreateToken) | **Post** /api/v1/projects/{project}/roles/{role}/token | Create a new project token
[**ProjectServiceDelete**](ProjectServiceApi.md#ProjectServiceDelete) | **Delete** /api/v1/projects/{name} | Delete deletes a project
[**ProjectServiceDeleteToken**](ProjectServiceApi.md#ProjectServiceDeleteToken) | **Delete** /api/v1/projects/{project}/roles/{role}/token/{iat} | Delete a new project token
[**ProjectServiceGet**](ProjectServiceApi.md#ProjectServiceGet) | **Get** /api/v1/projects/{name} | Get returns a project by name
[**ProjectServiceGetDetailedProject**](ProjectServiceApi.md#ProjectServiceGetDetailedProject) | **Get** /api/v1/projects/{name}/detailed | GetDetailedProject returns a project that include project, global project and scoped resources by name
[**ProjectServiceGetGlobalProjects**](ProjectServiceApi.md#ProjectServiceGetGlobalProjects) | **Get** /api/v1/projects/{name}/globalprojects | Get returns a virtual project by name
[**ProjectServiceGetSyncWindowsState**](ProjectServiceApi.md#ProjectServiceGetSyncWindowsState) | **Get** /api/v1/projects/{name}/syncwindows | GetSchedulesState returns true if there are any active sync syncWindows
[**ProjectServiceList**](ProjectServiceApi.md#ProjectServiceList) | **Get** /api/v1/projects | List returns list of projects
[**ProjectServiceListEvents**](ProjectServiceApi.md#ProjectServiceListEvents) | **Get** /api/v1/projects/{name}/events | ListEvents returns a list of project events
[**ProjectServiceUpdate**](ProjectServiceApi.md#ProjectServiceUpdate) | **Put** /api/v1/projects/{project.metadata.name} | Update updates a project



## ProjectServiceCreate

> V1alpha1AppProject ProjectServiceCreate(ctx).Body(body).Execute()

Create a new project

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
    body := *openapiclient.NewProjectProjectCreateRequest() // ProjectProjectCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceCreate`: V1alpha1AppProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProjectProjectCreateRequest**](ProjectProjectCreateRequest.md) |  | 

### Return type

[**V1alpha1AppProject**](V1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceCreateToken

> ProjectProjectTokenResponse ProjectServiceCreateToken(ctx, project, role).Body(body).Execute()

Create a new project token

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
    project := "project_example" // string | 
    role := "role_example" // string | 
    body := *openapiclient.NewProjectProjectTokenCreateRequest() // ProjectProjectTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceCreateToken(context.Background(), project, role).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceCreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceCreateToken`: ProjectProjectTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceCreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**role** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProjectProjectTokenCreateRequest**](ProjectProjectTokenCreateRequest.md) |  | 

### Return type

[**ProjectProjectTokenResponse**](ProjectProjectTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceDelete

> map[string]interface{} ProjectServiceDelete(ctx, name).Execute()

Delete deletes a project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceDeleteRequest struct via the builder pattern


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


## ProjectServiceDeleteToken

> map[string]interface{} ProjectServiceDeleteToken(ctx, project, role, iat).Id(id).Execute()

Delete a new project token

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
    project := "project_example" // string | 
    role := "role_example" // string | 
    iat := "iat_example" // string | 
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceDeleteToken(context.Background(), project, role, iat).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceDeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceDeleteToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceDeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**role** | **string** |  | 
**iat** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **id** | **string** |  | 

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


## ProjectServiceGet

> V1alpha1AppProject ProjectServiceGet(ctx, name).Execute()

Get returns a project by name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGet`: V1alpha1AppProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha1AppProject**](V1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGetDetailedProject

> ProjectDetailedProjectsResponse ProjectServiceGetDetailedProject(ctx, name).Execute()

GetDetailedProject returns a project that include project, global project and scoped resources by name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGetDetailedProject(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGetDetailedProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGetDetailedProject`: ProjectDetailedProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGetDetailedProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetDetailedProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectDetailedProjectsResponse**](ProjectDetailedProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGetGlobalProjects

> ProjectGlobalProjectsResponse ProjectServiceGetGlobalProjects(ctx, name).Execute()

Get returns a virtual project by name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGetGlobalProjects(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGetGlobalProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGetGlobalProjects`: ProjectGlobalProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGetGlobalProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetGlobalProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectGlobalProjectsResponse**](ProjectGlobalProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGetSyncWindowsState

> ProjectSyncWindowsResponse ProjectServiceGetSyncWindowsState(ctx, name).Execute()

GetSchedulesState returns true if there are any active sync syncWindows

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGetSyncWindowsState(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGetSyncWindowsState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGetSyncWindowsState`: ProjectSyncWindowsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGetSyncWindowsState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetSyncWindowsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectSyncWindowsResponse**](ProjectSyncWindowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceList

> V1alpha1AppProjectList ProjectServiceList(ctx).Name(name).Execute()

List returns list of projects

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
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceList(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceList`: V1alpha1AppProjectList
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**V1alpha1AppProjectList**](V1alpha1AppProjectList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceListEvents

> V1EventList ProjectServiceListEvents(ctx, name).Execute()

ListEvents returns a list of project events

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceListEvents(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceListEvents`: V1EventList
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceListEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ProjectServiceUpdate

> V1alpha1AppProject ProjectServiceUpdate(ctx, projectMetadataName).Body(body).Execute()

Update updates a project

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
    projectMetadataName := "projectMetadataName_example" // string | Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional
    body := *openapiclient.NewProjectProjectUpdateRequest() // ProjectProjectUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceUpdate(context.Background(), projectMetadataName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceUpdate`: V1alpha1AppProject
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectMetadataName** | **string** | Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProjectProjectUpdateRequest**](ProjectProjectUpdateRequest.md) |  | 

### Return type

[**V1alpha1AppProject**](V1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

