# \ClusterServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterServiceCreate**](ClusterServiceApi.md#ClusterServiceCreate) | **Post** /api/v1/clusters | Create creates a cluster
[**ClusterServiceDelete**](ClusterServiceApi.md#ClusterServiceDelete) | **Delete** /api/v1/clusters/{id.value} | Delete deletes a cluster
[**ClusterServiceGet**](ClusterServiceApi.md#ClusterServiceGet) | **Get** /api/v1/clusters/{id.value} | Get returns a cluster by server address
[**ClusterServiceInvalidateCache**](ClusterServiceApi.md#ClusterServiceInvalidateCache) | **Post** /api/v1/clusters/{id.value}/invalidate-cache | InvalidateCache invalidates cluster cache
[**ClusterServiceList**](ClusterServiceApi.md#ClusterServiceList) | **Get** /api/v1/clusters | List returns list of clusters
[**ClusterServiceRotateAuth**](ClusterServiceApi.md#ClusterServiceRotateAuth) | **Post** /api/v1/clusters/{id.value}/rotate-auth | RotateAuth rotates the bearer token used for a cluster
[**ClusterServiceUpdate**](ClusterServiceApi.md#ClusterServiceUpdate) | **Put** /api/v1/clusters/{id.value} | Update updates a cluster



## ClusterServiceCreate

> V1alpha1Cluster ClusterServiceCreate(ctx).Body(body).Upsert(upsert).Execute()

Create creates a cluster

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
    body := *openapiclient.NewV1alpha1Cluster() // V1alpha1Cluster | 
    upsert := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceCreate(context.Background()).Body(body).Upsert(upsert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceCreate`: V1alpha1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md) |  | 
 **upsert** | **bool** |  | 

### Return type

[**V1alpha1Cluster**](V1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterServiceDelete

> map[string]interface{} ClusterServiceDelete(ctx, idValue).Server(server).Name(name).IdType(idType).Execute()

Delete deletes a cluster

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
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name
    server := "server_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    idType := "idType_example" // string | type is the type of the specified cluster identifier ( \"server\" - default, \"name\" ). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceDelete(context.Background(), idValue).Server(server).Name(name).IdType(idType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idValue** | **string** | value holds the cluster server URL or cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | **string** |  | 
 **name** | **string** |  | 
 **idType** | **string** | type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 

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


## ClusterServiceGet

> V1alpha1Cluster ClusterServiceGet(ctx, idValue).Server(server).Name(name).IdType(idType).Execute()

Get returns a cluster by server address

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
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name
    server := "server_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    idType := "idType_example" // string | type is the type of the specified cluster identifier ( \"server\" - default, \"name\" ). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceGet(context.Background(), idValue).Server(server).Name(name).IdType(idType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceGet`: V1alpha1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idValue** | **string** | value holds the cluster server URL or cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | **string** |  | 
 **name** | **string** |  | 
 **idType** | **string** | type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 

### Return type

[**V1alpha1Cluster**](V1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterServiceInvalidateCache

> V1alpha1Cluster ClusterServiceInvalidateCache(ctx, idValue).Execute()

InvalidateCache invalidates cluster cache

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
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceInvalidateCache(context.Background(), idValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceInvalidateCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceInvalidateCache`: V1alpha1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceInvalidateCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idValue** | **string** | value holds the cluster server URL or cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceInvalidateCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha1Cluster**](V1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterServiceList

> V1alpha1ClusterList ClusterServiceList(ctx).Server(server).Name(name).IdType(idType).IdValue(idValue).Execute()

List returns list of clusters

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
    server := "server_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    idType := "idType_example" // string | type is the type of the specified cluster identifier ( \"server\" - default, \"name\" ). (optional)
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceList(context.Background()).Server(server).Name(name).IdType(idType).IdValue(idValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceList`: V1alpha1ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server** | **string** |  | 
 **name** | **string** |  | 
 **idType** | **string** | type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **idValue** | **string** | value holds the cluster server URL or cluster name. | 

### Return type

[**V1alpha1ClusterList**](V1alpha1ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterServiceRotateAuth

> map[string]interface{} ClusterServiceRotateAuth(ctx, idValue).Execute()

RotateAuth rotates the bearer token used for a cluster

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
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceRotateAuth(context.Background(), idValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceRotateAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceRotateAuth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceRotateAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idValue** | **string** | value holds the cluster server URL or cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceRotateAuthRequest struct via the builder pattern


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


## ClusterServiceUpdate

> V1alpha1Cluster ClusterServiceUpdate(ctx, idValue).Body(body).UpdatedFields(updatedFields).IdType(idType).Execute()

Update updates a cluster

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
    idValue := "idValue_example" // string | value holds the cluster server URL or cluster name
    body := *openapiclient.NewV1alpha1Cluster() // V1alpha1Cluster | 
    updatedFields := []string{"Inner_example"} // []string |  (optional)
    idType := "idType_example" // string | type is the type of the specified cluster identifier ( \"server\" - default, \"name\" ). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterServiceApi.ClusterServiceUpdate(context.Background(), idValue).Body(body).UpdatedFields(updatedFields).IdType(idType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterServiceApi.ClusterServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClusterServiceUpdate`: V1alpha1Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterServiceApi.ClusterServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idValue** | **string** | value holds the cluster server URL or cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md) |  | 
 **updatedFields** | **[]string** |  | 
 **idType** | **string** | type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 

### Return type

[**V1alpha1Cluster**](V1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

