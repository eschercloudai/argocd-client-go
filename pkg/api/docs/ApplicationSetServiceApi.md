# \ApplicationSetServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationSetServiceCreate**](ApplicationSetServiceApi.md#ApplicationSetServiceCreate) | **Post** /api/v1/applicationsets | Create creates an applicationset
[**ApplicationSetServiceDelete**](ApplicationSetServiceApi.md#ApplicationSetServiceDelete) | **Delete** /api/v1/applicationsets/{name} | Delete deletes an application set
[**ApplicationSetServiceGet**](ApplicationSetServiceApi.md#ApplicationSetServiceGet) | **Get** /api/v1/applicationsets/{name} | Get returns an applicationset by name
[**ApplicationSetServiceList**](ApplicationSetServiceApi.md#ApplicationSetServiceList) | **Get** /api/v1/applicationsets | List returns list of applicationset



## ApplicationSetServiceCreate

> V1alpha1ApplicationSet ApplicationSetServiceCreate(ctx).Body(body).Upsert(upsert).Execute()

Create creates an applicationset

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
    body := *openapiclient.NewV1alpha1ApplicationSet() // V1alpha1ApplicationSet | 
    upsert := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSetServiceApi.ApplicationSetServiceCreate(context.Background()).Body(body).Upsert(upsert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSetServiceApi.ApplicationSetServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationSetServiceCreate`: V1alpha1ApplicationSet
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSetServiceApi.ApplicationSetServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationSetServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1ApplicationSet**](V1alpha1ApplicationSet.md) |  | 
 **upsert** | **bool** |  | 

### Return type

[**V1alpha1ApplicationSet**](V1alpha1ApplicationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSetServiceDelete

> ApplicationsetApplicationSetResponse ApplicationSetServiceDelete(ctx, name).Execute()

Delete deletes an application set

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
    resp, r, err := apiClient.ApplicationSetServiceApi.ApplicationSetServiceDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSetServiceApi.ApplicationSetServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationSetServiceDelete`: ApplicationsetApplicationSetResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSetServiceApi.ApplicationSetServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationSetServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationsetApplicationSetResponse**](ApplicationsetApplicationSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSetServiceGet

> V1alpha1ApplicationSet ApplicationSetServiceGet(ctx, name).Execute()

Get returns an applicationset by name

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
    name := "name_example" // string | the applicationsets's name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSetServiceApi.ApplicationSetServiceGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSetServiceApi.ApplicationSetServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationSetServiceGet`: V1alpha1ApplicationSet
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSetServiceApi.ApplicationSetServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the applicationsets&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationSetServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha1ApplicationSet**](V1alpha1ApplicationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSetServiceList

> V1alpha1ApplicationSetList ApplicationSetServiceList(ctx).Projects(projects).Selector(selector).Execute()

List returns list of applicationset

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
    projects := []string{"Inner_example"} // []string | the project names to restrict returned list applicationsets. (optional)
    selector := "selector_example" // string | the selector to restrict returned list to applications only with matched labels. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSetServiceApi.ApplicationSetServiceList(context.Background()).Projects(projects).Selector(selector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSetServiceApi.ApplicationSetServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationSetServiceList`: V1alpha1ApplicationSetList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSetServiceApi.ApplicationSetServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationSetServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projects** | **[]string** | the project names to restrict returned list applicationsets. | 
 **selector** | **string** | the selector to restrict returned list to applications only with matched labels. | 

### Return type

[**V1alpha1ApplicationSetList**](V1alpha1ApplicationSetList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

